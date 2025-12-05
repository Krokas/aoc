package main

import (
	aoc2021 "aoc/2021"
	aoc2025 "aoc/start/2025"
	"strconv"
	"time"

	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/bubbles/list"
	"github.com/charmbracelet/bubbles/progress"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type Model struct {
	list        list.Model // items on the to-do list
	progress    progress.Model
	yearCursor  int
	cursor      int // which to-do list item our cursor is pointing at
	level       int
	calculating bool
	hasResult   bool
	result      int
}

const (
	padding  = 2
	maxWidth = 80
)

type tickMsg time.Time

var docStyle = lipgloss.NewStyle().Margin(1, 2)
var helpStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("#626262")).Render

func initialModel(items []list.Item) Model {

	m := Model{
		// Our to-do list is a grocery list
		list:        list.New(items, list.NewDefaultDelegate(), 0, 0),
		progress:    progress.New(progress.WithDefaultGradient()),
		level:       0,
		calculating: false,
		hasResult:   false,
		result:      0,
	}

	m.list.Title = "Advent of code"

	return m
}

func (m Model) Init() tea.Cmd {
	// Just return `nil`, which means "no I/O right now, please."
	return nil
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	switch msg := msg.(type) {
	case tea.KeyMsg:
		if msg.String() == "ctrl+c" {
			return m, cmd
		}

		if msg.String() == "up" {
			m.cursor--

			if m.cursor < 0 {
				m.cursor = 0
			}
		}

		if msg.String() == "down" {
			m.cursor++

			if m.cursor >= len(m.list.Items()) {
				m.cursor = len(m.list.Items()) - 1
			}
		}

		if msg.String() == "backspace" {
			if m.level > 0 {
				m.level = 0
				m.cursor = 0
				m.list.SetItems(initialList())
			}

			return m, nil
		}

		if msg.String() == "enter" {
			switch m.level {
			case 1:
				m.hasResult = false
				m.result = 0
				m.calculating = true
				return m, tickCmd()
			default:
				title := m.list.Items()[m.cursor].FilterValue()
				m.list.Title = title
				m.yearCursor = m.cursor
				m.list.AdditionalFullHelpKeys = func() []key.Binding {
					return []key.Binding{
						key.NewBinding(
							key.WithKeys("backspace"),
							key.WithHelp("backspace", "Back"),
						),
					}
				}
				switch m.cursor {
				case 0: // 2021
					m.list.SetItems(aoc2021.GetList())
					m.level = 1
					m.cursor = 0
					break
				case 4: // 2025
					m.list.SetItems(aoc2025.GetList())
					m.level = 1
					m.cursor = 0
					break
				default:
					m.list.SetItems(initialList())
					m.level = 0
					break
				}
				_, cmd = m.list.Update(msg)
				return m, cmd
			}
		}
	case tea.WindowSizeMsg:
		h, v := docStyle.GetFrameSize()
		m.list.SetSize(msg.Width-h, msg.Height-v)

		m.progress.Width = min(msg.Width-padding*2-4, maxWidth)

	case tickMsg:
		if m.progress.Percent() == 1.0 {
			m.calculating = false
			m.hasResult = true
			m.result = execute(m.yearCursor, m.cursor)
			return m, nil
		}
		cmd = m.progress.IncrPercent(0.25)
		return m, tea.Batch(tickCmd(), cmd)
	case progress.FrameMsg:
		progressModel, cmd := m.progress.Update(msg)
		m.progress = progressModel.(progress.Model)
		return m, cmd
	default:
		return m, nil
	}

	m.list, cmd = m.list.Update(msg)
	return m, cmd
}

func tickCmd() tea.Cmd {
	return tea.Tick(time.Second*1, func(t time.Time) tea.Msg {
		return tickMsg(t)
	})
}

func (m Model) View() string {

	var anotherStyle = lipgloss.NewStyle().
		BorderStyle(lipgloss.RoundedBorder()).
		BorderForeground(lipgloss.Color("228")).
		BorderBackground(lipgloss.Color("63")).
		BorderTop(true).
		BorderLeft(true).
		BorderRight(true).
		BorderBottom(true)

	var views []string

	listStyle := lipgloss.NewStyle().
		Padding(1, 2)

	listResult := listStyle.Render(m.list.View())

	views = append(views, listResult)

	resultStyle := lipgloss.NewStyle().
		Padding(1, 2).
		BorderStyle(lipgloss.NormalBorder()).
		BorderForeground(lipgloss.Color("#ffffff")).
		BorderLeft(true)

	if m.calculating && !m.hasResult {
		styledProgress := resultStyle.Render(m.progress.View())
		views = append(views, styledProgress)
	}

	if !m.calculating && m.hasResult {
		resultString := "Result: " + strconv.Itoa(m.result)
		view := resultStyle.Render(resultString)
		views = append(views, view)
	}

	return anotherStyle.Render(lipgloss.JoinHorizontal(lipgloss.Top, views...))
}
