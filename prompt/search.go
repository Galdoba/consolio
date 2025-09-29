package prompt

import (
	"fmt"
	"strings"

	"github.com/charmbracelet/bubbles/key"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/huh"
	"github.com/charmbracelet/lipgloss"
)

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

	lg            *lipgloss.Renderer
	theme         *huh.Theme
	fullList      []*Item
	filteredList  []*Item
	selectedItem  *Item
	width         int
	height        int
	caseSensitive bool
}

func newSearch(opts ...PromptOption) (*searchModel, error) {
	pb, err := newPromptBuilder(formSearch)
	if err != nil {
		return nil, fmt.Errorf("failed to create prompt builder: %v", err)
	}
	for _, modify := range opts {
		modify(pb)
	}
	sm := searchModel{
		filter:        "",
		cursor:        defaultCursor,
		lg:            lipgloss.DefaultRenderer(),
		fullList:      pb.items,
		filteredList:  pb.items,
		title:         pb.title,
		theme:         pb.theme,
		description:   pb.description,
		body:          "",
		summary:       "",
		help:          defaultHelpSearch,
		width:         pb.width,
		height:        pb.height,
		caseSensitive: pb.searchCaseSensitive,
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
		searchTerm := normalizeFilter(sm.filter, sm.caseSensitive)

		for _, item := range sm.fullList {
			if sm.matchSearch(item, searchTerm) {
				list = append(list, item)
			}
		}
		sm.filteredList = list
	}
}

func normalizeFilter(filter string, caseSensitive bool) string {
	if !caseSensitive {
		return strings.ToLower(filter)
	}
	return filter
}

func (sm *searchModel) matchSearch(item *Item, searchTerm string) bool {
	itemKey := item.GetKey()
	if !sm.caseSensitive {
		itemKey = strings.ToLower(itemKey)
	}
	return strings.Contains(itemKey, searchTerm)
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

func startLine() string {
	return lipgloss.NewStyle().Render("\n┃ ")
}

func (sm *searchModel) viewTitle() string {
	if sm.title == "" {
		return ""
	}
	return sm.theme.Focused.TextInput.Prompt.Render("┃ ") + sm.theme.Focused.Title.Render(sm.title)
}

func (sm *searchModel) viewDescription() string {
	if sm.description == "" {
		return ""
	}
	return startLine() + sm.theme.Focused.Description.Render(sm.description)
}

func (sm *searchModel) viewFilter() string {
	s := startLine() + "filter: " + sm.theme.Focused.TextInput.Prompt.Render(sm.filter)
	s += startLine()
	return s
}

func (sm *searchModel) viewBody() string {
	start := sm.cursor.offset
	end := sm.maxCursorIndexAllowed()
	s := ""
	for i := start; i < end; i++ {
		item := sm.filteredList[i]
		s += startLine() + sm.renderCursor(i) + sm.renderItem(item)
	}
	return s
}

func (sm *searchModel) renderCursor(index int) string {
	cursStr := sm.cursor.unselected
	if index == sm.cursor.index {
		cursStr = sm.cursor.selected
	}
	return sm.theme.Focused.SelectedOption.Render(cursStr)
}

func (sm *searchModel) renderItem(item *Item) string {
	/*
		ai generated func
	*/
	s := item.GetKey()
	lowerItem := strings.ToLower(item.GetKey())
	lowerInput := strings.ToLower(sm.filter)
	if idx := strings.Index(lowerItem, lowerInput); idx != -1 {
		before := item.GetKey()[:idx]
		match := item.GetKey()[idx : idx+len(sm.filter)]
		after := item.GetKey()[idx+len(sm.filter):]
		s = before + sm.theme.Focused.SelectedOption.Render(match) + after
	}
	return s
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
	s := ""
	if len(sm.filteredList) != sm.maxCursorIndexAllowed() || len(sm.filteredList) != len(sm.fullList) {
		s += startLine() + startLine() + sm.theme.Focused.Option.Render(fmt.Sprintf("%v/%v items filtered", len(sm.filteredList), len(sm.fullList)))
	}
	if sm.maxCursorIndexAllowed()-sm.cursor.offset != len(sm.filteredList) {
		s += "\n" + sm.theme.Help.ShortKey.Render(fmt.Sprintf("show items [%v-%v] of %v filtered", sm.cursor.offset, sm.maxCursorIndexAllowed(), len(sm.filteredList)))

	}

	return s
}

func (sm *searchModel) viewHelp() string {
	return renderHelp(sm.theme, []key.Binding{
		key.NewBinding(key.WithHelp("↑/↓", "move cursor")),
		key.NewBinding(key.WithHelp("enter", "submit")),
		key.NewBinding(key.WithHelp("esc", "cancel")),
	})
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

func renderHelp(theme *huh.Theme, binds []key.Binding) string {
	s := "\n\n"
	for i, bind := range binds {
		s += theme.Help.FullKey.Render(bind.Help().Key) + " " + theme.Help.FullDesc.Render(bind.Help().Desc)
		if i != len(binds)-1 {
			s += theme.Help.FullSeparator.Render(" • ")
		}
	}
	return s
}
