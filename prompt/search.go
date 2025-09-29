package prompt

import (
	"fmt"
	"strings"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/huh"
	"github.com/charmbracelet/lipgloss"
)

type filter string

type searchStyle struct {
	Title,
	Descr,
	Filter,
	Cursor,
	List,
	Summary,
	Help lipgloss.Style
}

func newSearchStyle(theme *huh.Theme) *searchStyle {
	ss := searchStyle{}
	ss.Title = theme.Focused.Title
	ss.Descr = theme.Focused.Description
	ss.Filter = theme.Focused.File
	ss.Cursor = theme.Focused.SelectedOption
	ss.List = theme.Focused.UnselectedOption
	ss.Summary = theme.Help.ShortKey
	ss.Help = theme.Help.FullKey
	return &ss
}

type cursor struct {
	index      int
	offset     int
	selected   string
	unselected string
}

var defaultCursor = cursor{
	index:      0,
	offset:     0,
	selected:   "> ",
	unselected: "  ",
}

type searchModel struct {
	title       string
	description string
	filter      string
	body        string
	summary     string
	help        string
	cursor      cursor

	lg           *lipgloss.Renderer
	style        *searchStyle
	fullList     []*Item
	filteredList []*Item
	selectedItem *Item
	width        int
	height       int
}

var searchResult = &Item{}

func newSearch(opts ...PromptOption) (*searchModel, error) {
	pb, err := newPromptBuilder(formSearch)
	if err != nil {
		return nil, fmt.Errorf("failed to create prompt builder: %v", err)
	}
	for _, modify := range opts {
		modify(pb)
	}
	sm := searchModel{
		filter:       "",
		cursor:       defaultCursor,
		lg:           lipgloss.DefaultRenderer(),
		fullList:     pb.items,
		filteredList: pb.items,
		// selectedItem: &Item{},
		title:       pb.title,
		style:       newSearchStyle(pb.theme),
		description: pb.description,
		body:        "",
		summary:     "",
		help:        defaultHelpSearch,
		width:       pb.width,
		height:      pb.height,
	}
	copy(sm.filteredList, sm.fullList)
	return &sm, nil
}

func (sm searchModel) Init() tea.Cmd {
	return nil
}

func (sm *searchModel) updateFilter() {
	switch sm.filter {
	case "":
		sm.filteredList = sm.fullList
	default:
		list := []*Item{}
		for _, item := range sm.fullList {
			if strings.Contains(item.GetKey(), sm.filter) {
				list = append(list, item)
			}
		}
		sm.filteredList = list
		sm.cursor.index = 0
	}
}

func (sm searchModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmds []tea.Cmd
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c":
			return sm, tea.Interrupt
		case "esc":
			return sm, tea.Quit
		case "backspace":
			switch glyphsLen(sm.filter) {
			case 0:
			default:
				letters := strings.Split(sm.filter, "")
				sm.filter = strings.Join(letters[:len(letters)-1], "")
			}
			sm.updateFilter()
			sm.cursorReset()
		case "up":
			sm.moveCursor(-1)
		case "down":
			sm.moveCursor(1)
		case "pgdown":
			sm.moveCursor(sm.maxListHeight())
		case "pgup":
			sm.moveCursor(sm.maxListHeight() * -1)
		case "enter":
			sm.selectedItem = sm.getSelectedItem()
			cmds = append(cmds, tea.Quit)
		default:
			if glyphsLen(msg.String()) == 1 {
				sm.filter += msg.String()
			}
			sm.updateFilter()
			sm.cursorReset()
		}
	}
	if len(sm.filteredList) == 1 {
		sm.selectedItem = sm.filteredList[0]
		cmds = append(cmds, tea.Quit)
	}

	return sm, tea.Batch(cmds...)
}

func (sm *searchModel) cursorReset() {
	sm.cursor.index = 0
	sm.cursor.offset = 0
}

func (sm *searchModel) moveCursor(direction int) {
	if len(sm.filteredList) > 0 && direction != 0 {
		switch direction > 0 {
		case true:
			sm.cursor.index = min(len(sm.filteredList)-1, sm.cursor.index+direction)
		case false:
			sm.cursor.index = max(sm.cursor.offset+direction, sm.cursor.index+direction, 0)
		}
		if sm.cursor.index >= sm.maxCursorIndexAllowed() {
			sm.cursor.index = sm.maxCursorIndexAllowed()
			sm.cursor.offset = sm.cursor.index - sm.maxListHeight() + direction
		}
		if sm.cursor.index < sm.cursor.offset {
			sm.cursor.offset = sm.cursor.index
		}
	}
}

func glyphsLen(s string) int {
	return len(strings.Split(s, ""))
}

func (sm *searchModel) getSelectedItem() *Item {
	for i, item := range sm.filteredList {
		if i == sm.cursor.index {
			return item
		}
	}
	return nil
}

func (sm *searchModel) viewTitle() string {
	if sm.title == "" {
		return ""
	}
	return sm.style.Title.Render(sm.title) + "\n"
}

func (sm *searchModel) viewDescription() string {
	if sm.description == "" {
		return ""
	}
	return sm.style.Descr.Render(sm.description) + "\n"
}

func (sm *searchModel) viewFilter() string {
	s := sm.style.Filter.Render("filter: "+sm.filter) + "\n"
	s += " " + "\n"
	return s
}

func (sm *searchModel) viewBody() string {
	start := sm.cursor.offset
	end := sm.maxCursorIndexAllowed()
	s := ""
	for i := start; i < end; i++ {
		item := sm.filteredList[i]
		s += sm.renderCursor(i) + sm.renderItem(item) + "\n"
	}
	return s
}

func (sm *searchModel) renderCursor(index int) string {
	cursStr := sm.cursor.unselected
	if index == sm.cursor.index {
		cursStr = sm.cursor.selected
	}
	return sm.style.Cursor.Render(cursStr)
}

func (sm *searchModel) renderItem(item *Item) string {
	clean := strings.Split(item.Key, sm.filter)
	s := ""
	for i := 0; i < len(clean); i++ {
		s += clean[i]
		if i == len(clean)-1 {
			continue
		}
		s += sm.style.Cursor.Render(sm.filter)
	}
	return s
}

func (sm *searchModel) renderListed(i int) string {
	s := ""
	item := sm.filteredList[i]
	switch i {
	case sm.cursor.index:
		s += sm.style.Cursor.Render(sm.cursor.selected) + renderItemKey(sm.style.Filter, item.GetKey(), sm.filter) + "\n"
	default:
		s += sm.style.Cursor.Render(sm.cursor.unselected) + renderItemKey(sm.style.Filter, item.GetKey(), sm.filter) + "\n"
	}
	return s
}

func renderItemKey(style lipgloss.Style, key string, filter string) string {
	parts := strings.Split(key, filter)
	switch len(parts) {
	case 1:
		return key
	default:
		s := ""
		connectorStr := style.Render(filter)
		for i := range parts {
			s += parts[i]
			if i == len(parts)-1 {
				continue
			}
			s += connectorStr
		}
		return s
	}
}

func (sm *searchModel) maxCursorIndexAllowed() int {
	return min(sm.cursor.offset+sm.maxListHeight(), len(sm.filteredList))
}

func (sm *searchModel) maxListHeight() int {
	n := 0
	for _, data := range []string{sm.title, sm.description, sm.viewFilter(), sm.summary, sm.help} {
		if data == "" {
			continue
		}
		lines := strings.Split(data, "\n")
		n += len(lines)
	}
	return max(sm.height-n-3, 0)
}

func (sm *searchModel) viewSummary() string {
	s := "\n"
	if len(sm.filteredList) != sm.maxCursorIndexAllowed() || len(sm.filteredList) != len(sm.fullList) {
		s += sm.style.Summary.Render(fmt.Sprintf(" %v/%v items filtered", len(sm.filteredList), len(sm.fullList))) + "\n"

	}
	s += sm.style.Summary.Render(fmt.Sprintf(" show items [%v-%v] of %v filtered", sm.cursor.offset, sm.maxCursorIndexAllowed(), len(sm.filteredList))) + "\n"
	return s
}

func (sm *searchModel) viewHelp() string {
	return sm.style.Help.Render("\n"+fmt.Sprintf("%v", sm.help)) + "\n"
}

func (sm searchModel) View() string {
	if sm.selectedItem != nil {
		return ""
	}
	s := sm.viewTitle()
	s += sm.viewDescription()
	s += sm.viewFilter()
	s += sm.viewBody()
	s += sm.viewSummary()
	s += sm.viewHelp()
	return s
}
