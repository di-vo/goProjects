package main

import (
	"fmt"
	"math/rand/v2"
	"os"
	"slices"
	"strings"
	"time"

	"github.com/charmbracelet/bubbles/timer"
	tea "github.com/charmbracelet/bubbletea"
)

const (
	rows           = 10
	cols           = 10
	timeout        = time.Second
	playerInterval = time.Second / 4
	itemInterval   = 5
)

var (
	itemTimerCount int = itemInterval
)

type vec struct {
	x, y int
}

func (v *vec) add(v1 vec) vec {
	return vec{
		x: v.x + v1.x,
		y: v.y + v1.y,
	}
}

func (v *vec) subtract(v1 vec) vec {
	return vec{
		x: v.x - v1.x,
		y: v.y - v1.y,
	}
}

type playerElem struct {
	pos vec
	dir vec
}

type playerStruct struct {
	head playerElem
	body []playerElem
}

type model struct {
	timer  timer.Model
	board  [][]int
	player playerStruct
	items  []vec
	score  int
	won    bool
}

func randRange(min, max int) int {
	return rand.IntN(max-min) + min
}

func initialModel() model {
	board := make([][]int, rows)

	for i := range board {
		board[i] = make([]int, cols)
	}

	player := playerStruct{
		head: playerElem{pos: vec{x: 0, y: 0}, dir: vec{x: 1, y: 0}},
		body: make([]playerElem, 0),
	}

	return model{
		timer:  timer.NewWithInterval(timeout, playerInterval),
		board:  board,
		player: player,
		items:  make([]vec, 0),
		score:  0,
        won: false,
	}
}

func (m model) Init() tea.Cmd {
	return m.timer.Init()
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q":
			return m, tea.Quit
		case "left", "h":
			if m.player.head.dir.x == 0 {
				m.player.head.dir = vec{x: -1, y: 0}
			}
		case "down", "j":
			if m.player.head.dir.y == 0 {
				m.player.head.dir = vec{x: 0, y: 1}
			}
		case "up", "k":
			if m.player.head.dir.y == 0 {
				m.player.head.dir = vec{x: 0, y: -1}
			}
		case "right", "l":
			if m.player.head.dir.x == 0 {
				m.player.head.dir = vec{x: 1, y: 0}
			}
		}
	case timer.TickMsg:
		if (m.player.head.dir.x < 0 && m.player.head.pos.x > 0) || (m.player.head.dir.x > 0 && m.player.head.pos.x < rows-1) {
			// horizontal check
			m.player.head.pos.x += m.player.head.dir.x
		} else if (m.player.head.dir.y < 0 && m.player.head.pos.y > 0) || (m.player.head.dir.y > 0 && m.player.head.pos.y < cols-1) {
			// vertical check
			m.player.head.pos.y += m.player.head.dir.y
		} else {
			return m, tea.Quit
		}

		overlap := slices.IndexFunc(m.player.body, func(v playerElem) bool { return m.player.head.pos.x == v.pos.x && m.player.head.pos.y == v.pos.y })

		if overlap != -1 {
			return m, tea.Quit
		}

		for i := range m.player.body {
			m.player.body[i].pos = m.player.body[i].pos.add(m.player.body[i].dir)
		}

		for i := len(m.player.body) - 1; i >= 0; i-- {
			if i == 0 {
				m.player.body[i].dir = m.player.head.dir
			} else {
				m.player.body[i].dir = m.player.body[i-1].dir
			}
		}

		idx := slices.IndexFunc(m.items, func(v vec) bool { return m.player.head.pos.x == v.x && m.player.head.pos.y == v.y })

		if idx != -1 {
			var pos, dir vec

			if len(m.player.body) == 0 {
				pos = m.player.head.pos.subtract(m.player.head.dir)
				dir = m.player.head.dir
			} else {
				pos = m.player.body[len(m.player.body)-1].pos.subtract(m.player.body[len(m.player.body)-1].dir)
				dir = m.player.body[len(m.player.body)-1].dir
			}

			m.player.body = append(m.player.body, playerElem{pos: pos, dir: dir})
			m.items = slices.Delete(m.items, idx, idx+1)
			m.score++

            if len(m.player.body) == rows * cols - 1 {
                m.won = true
                return m, tea.Quit
            }
		}

		var cmd tea.Cmd
		m.timer, cmd = m.timer.Update(msg)
		return m, cmd
	case timer.TimeoutMsg:
		itemTimerCount++

		if itemTimerCount >= itemInterval {
			spawn := vec{x: randRange(0, rows), y: randRange(0, cols)}
			m.items = append(m.items, spawn)
			itemTimerCount = 0
		}
		m.timer.Timeout = timeout
	}

	return m, nil
}

func (m model) View() string {
	s := "Term Snake!\n\n"

	s += fmt.Sprintf("Score: %d\n\n", m.score)

	s += strings.Repeat("-", rows+2) + "\n"

	for i := range m.board {
		for j := range m.board[i] {
			if j == 0 {
				s += "|"
			}

			c := " "

			for _, e := range m.items {
				if e.x == j && e.y == i {
					c = "x"
				}
			}

			// head
			if m.player.head.pos.x == j && m.player.head.pos.y == i {
				c = "o"
			}

			// body
			for _, e := range m.player.body {
				if e.pos.x == j && e.pos.y == i {
					c = "e"
				}
			}

			s += c

			if j == len(m.board[i])-1 {
				s += "|"
			}
		}
		s += "\n"
	}

	s += strings.Repeat("-", rows+2) + "\n"

    if m.won {
        s += "You won!\n"
    }

	return s
}

func main() {
	p := tea.NewProgram(initialModel())

	if _, err := p.Run(); err != nil {
		fmt.Printf("Alas, there's been an error: %v", err)
		os.Exit(1)
	}
}
