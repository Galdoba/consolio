package v2

import (
	"fmt"
	"strings"

	"github.com/charmbracelet/bubbles/key"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/huh"
	"github.com/charmbracelet/lipgloss"
)

// cursor manages the visual cursor position and appearance in the search list.
// (ai generated comment)
type cursor struct {
	index      int    // Current selected item index
	offset     int    // Scroll offset for displaying long lists
	selected   string // Symbol for selected item (e.g., "> ")
	unselected string // Symbol for unselected items (e.g., "  ")
}

// defaultCursor provides the standard cursor configuration.
// (ai generated comment)
var defaultCursor = cursor{
	index:      0,
	offset:     0,
	selected:   "> ",
	unselected: "  ",
}

// searchModel represents the Bubble Tea model for the interactive search prompt.
// It manages the search state, filtering, rendering, and user interactions.
// (ai generated comment)
type searchModel struct {
	title       string // Main title displayed at the top
	description string // Additional description text
	filter      string // Current search filter text
	body        string // Rendered list body content
	summary     string // Summary text showing filtered results
	cursor      cursor // Cursor management for selection

	lg            *lipgloss.Renderer // Lipgloss renderer for styling
	theme         *huh.Theme         // Visual theme for consistent styling
	fullList      []*Item            // Complete unfiltered item list
	filteredList  []*Item            // Currently filtered item list
	selectedItem  *Item              // Currently selected item
	width         int                // Prompt width
	height        int                // Prompt height
	caseSensitive bool               // Whether search is case sensitive
	done          bool               // Whether search is completed
	err           error              // Error state if search fails
}

// newSearch creates and initializes a new search model with the provided options.
// Sets up the search state with default values and applies configuration options.
// Returns the search model or an error if initialization fails.
// (ai generated comment)
func newSearch(opts ...PromptOption) (*searchModel, error) {
	pb := &promptBuilder{
		promptType:       ptSearch,
		settings:         map[any]any{},
		defaultsRegistry: defaultRegistry(),
	}
	for _, modify := range opts {
		modify(pb)
	}
	if err := validateRequiredFields(pb); err != nil {
		return nil, fmt.Errorf("failed field validation: %v", err)
	}
	sm := searchModel{
		filter:        "",
		cursor:        defaultCursor,
		lg:            lipgloss.DefaultRenderer(),
		fullList:      pb.getItems(),
		filteredList:  pb.getItems(),
		title:         pb.getTitle(),
		theme:         pb.getTheme(),
		description:   pb.getDescription(),
		body:          "",
		summary:       "",
		width:         pb.getWidth(),
		height:        max(pb.getHeight(), 20),
		caseSensitive: pb.getCaseSensitive(),
	}
	if len(sm.fullList) == 0 {
		return nil, fmt.Errorf("search-items pool is empty")
	}
	copy(sm.filteredList, sm.fullList)
	return &sm, nil
}

// Init initializes the search model as required by the Bubble Tea Model interface.
// (ai generated comment)
func (sm searchModel) Init() tea.Cmd {
	return nil
}

// updateFilter refreshes the filtered list based on the current search term.
// If the filter is empty, shows all items. Otherwise, filters items that match the search.
// (ai generated comment)
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

// normalizeFilter processes the filter text based on case sensitivity setting.
// Converts to lowercase if case insensitive search is enabled.
// (ai generated comment)
func normalizeFilter(filter string, caseSensitive bool) string {
	if !caseSensitive {
		return strings.ToLower(filter)
	}
	return filter
}

// matchSearch checks if an item matches the search term.
// Handles case sensitivity and performs substring matching.
// (ai generated comment)
func (sm *searchModel) matchSearch(item *Item, searchTerm string) bool {
	itemKey := item.key
	if !sm.caseSensitive {
		itemKey = strings.ToLower(itemKey)
	}
	return strings.Contains(itemKey, searchTerm)
}

// Update handles messages and updates the search model state.
// Processes keyboard input for navigation, filtering, and selection.
// Returns the updated model and any commands to execute.
// (ai generated comment)
func (sm searchModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmds []tea.Cmd
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c":
			return sm, tea.Interrupt
		case "esc":
			sm.done = true
			sm.err = fmt.Errorf("search canceled")
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
			if sm.selectedItem == nil {
				sm.err = fmt.Errorf("no item selected")
			}
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

// cursorReset resets the cursor to the top of the filtered list.
// Called when the filter changes to ensure selection starts from the beginning.
// (ai generated comment)
func (sm *searchModel) cursorReset() {
	sm.cursor.index = 0
	sm.cursor.offset = 0
}

// moveCursor moves the cursor by the specified direction (positive = down, negative = up).
// Handles bounds checking and adjusts the scroll offset when needed.
// (ai generated comment)
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

// glyphsLen calculates the number of glyphs (visible characters) in a string.
// (ai generated comment)
func glyphsLen(s string) int {
	return len([]rune(s))
}

// getSelectedItem returns the currently selected item based on cursor position.
// Returns nil if no item is selected or the list is empty.
// (ai generated comment)
func (sm *searchModel) getSelectedItem() *Item {
	for i, item := range sm.filteredList {
		if i == sm.cursor.index {
			return item
		}
	}
	return nil
}

// startLine returns the styled string that begins each line in the view.
// Uses box-drawing characters for consistent visual structure.
// (ai generated comment)
func startLine() string {
	return lipgloss.NewStyle().Render("\n┃ ")
}

// viewTitle renders the title section of the search prompt.
// Uses the theme's focused style for consistent appearance.
// (ai generated comment)
func (sm *searchModel) viewTitle() string {
	if sm.title == "" {
		return ""
	}
	return sm.theme.Focused.TextInput.Prompt.Render("┃ ") + sm.theme.Focused.Title.Render(sm.title)
}

// viewDescription renders the description section of the search prompt.
// Returns empty string if no description is set.
// (ai generated comment)
func (sm *searchModel) viewDescription() string {
	if sm.description == "" {
		return ""
	}
	return startLine() + sm.theme.Focused.Description.Render(sm.description)
}

// viewFilter renders the current filter input section.
// Shows the filter label and the current filter text.
// (ai generated comment)
func (sm *searchModel) viewFilter() string {
	s := startLine() + "filter: " + sm.theme.Focused.TextInput.Prompt.Render(sm.filter)
	s += startLine()
	return s
}

// viewBody renders the main list body with filtered items.
// Only displays items within the current scroll viewport.
// (ai generated comment)
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

// renderCursor renders the cursor symbol for a given index.
// Shows the selected symbol for the current index, unselected for others.
// (ai generated comment)
func (sm *searchModel) renderCursor(index int) string {
	cursStr := sm.cursor.unselected
	if index == sm.cursor.index {
		cursStr = sm.cursor.selected
	}
	return sm.theme.Focused.SelectedOption.Render(cursStr)
}

// renderItem renders an individual item with search term highlighting.
// Highlights the matching portion of the item key using theme colors.
// (ai generated comment)
func (sm *searchModel) renderItem(item *Item) string {
	/*
		ai generated func
	*/
	s := item.key
	lowerItem := strings.ToLower(item.key)
	lowerInput := strings.ToLower(sm.filter)
	if idx := strings.Index(lowerItem, lowerInput); idx != -1 {
		before := item.key[:idx]
		match := item.key[idx : idx+len(sm.filter)]
		after := item.key[idx+len(sm.filter):]
		s = before + sm.theme.Focused.SelectedOption.Render(match) + after
	}
	return s
}

// maxCursorIndexAllowed calculates the maximum cursor index for the current view.
// Ensures cursor doesn't go beyond the visible portion of the filtered list.
// (ai generated comment)
func (sm *searchModel) maxCursorIndexAllowed() int {
	return min(sm.cursor.offset+sm.maxListHeight(), len(sm.filteredList))
}

// maxListHeight calculates the maximum available height for displaying items.
// (ai generated comment)
func (sm *searchModel) maxListHeight() int {
	n := 0
	for _, data := range []string{sm.title, sm.description, sm.viewFilter(), sm.summary} {
		if data == "" {
			continue
		}
		lines := strings.Split(data, "\n")
		n += len(lines)
	}
	return max(sm.height-n-2, 0)
}

// viewSummary renders the summary section showing filtering statistics.
// Displays counts of filtered items and current viewport range.
// (ai generated comment)
func (sm *searchModel) viewSummary() string {
	s := ""
	if len(sm.filteredList) != 0 || len(sm.filteredList) != len(sm.fullList) {
		s += startLine() + startLine() + sm.theme.Focused.Option.Render(fmt.Sprintf("%v/%v items filtered", len(sm.filteredList), len(sm.fullList)))
	}
	if sm.maxCursorIndexAllowed()-sm.cursor.offset != len(sm.filteredList) {
		s += "\n" + sm.theme.Help.ShortKey.Render(fmt.Sprintf("show items [%v-%v] of %v filtered", sm.cursor.offset, sm.maxCursorIndexAllowed(), len(sm.filteredList)))

	}

	return s
}

// viewHelp renders the help section with key binding instructions.
// Shows available keyboard shortcuts for navigation and actions.
// (ai generated comment)
func (sm *searchModel) viewHelp() string {
	return renderHelp(sm.theme, []key.Binding{
		key.NewBinding(key.WithHelp("↑/↓", "move cursor")),
		key.NewBinding(key.WithHelp("enter", "submit")),
		key.NewBinding(key.WithHelp("esc", "cancel")),
	})
}

// View renders the complete search prompt interface.
// Returns empty string if the search is completed or an item is selected.
// (ai generated comment)
func (sm searchModel) View() string {
	if sm.done {
		return ""
	}
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

// renderHelp generates a formatted help section with key bindings.
// Uses the provided theme for consistent styling of keys and descriptions.
// (ai generated comment)
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
