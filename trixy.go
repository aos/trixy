package main

import (
	"math/rand"
	"time"

	"github.com/gdamore/tcell"
)

var currentY = 0
var lineLength = 10
var numLines = 50
var leadingGlyphs = []rune{
	'一', '二', '三', '四', '五', '六', '七', '八', '九', '十',
}

func main() {
	s, err := tcell.NewScreen()
	if err != nil {
		panic(err)
	}

	if err = s.Init(); err != nil {
		panic(err)
	}

	s.SetStyle(tcell.StyleDefault.
		Foreground(tcell.ColorWhite).
		Background(tcell.ColorBlack))
	s.Clear()

	quit := make(chan struct{})
	go func() {
		for {
			ev := s.PollEvent()
			switch ev := ev.(type) {
			case *tcell.EventKey:
				switch ev.Key() {
				case tcell.KeyEscape, tcell.KeyEnter:
					close(quit)
					return
				case tcell.KeyCtrlL:
					s.Sync()
				}
			case *tcell.EventResize:
				s.Sync()
			}
		}
	}()

loop:
	for {
		select {
		case <-quit:
			break loop
		case <-time.After(time.Millisecond * 250):
		}
		Lines(s, numLines)
	}

	s.Fini()
}

func Lines(s tcell.Screen, num int) {
	rand.Seed(time.Now().UnixNano())
	w, h := s.Size()
	s.Clear()

	for i := 0; i < num; i++ {
		DrawLine(s, rand.Intn(w), h)
		DrawLine(s, rand.Intn(w), h)
	}
}

func DrawLine(s tcell.Screen, xLocation, h int) {
	st := tcell.StyleDefault

	for i := 0; i < lineLength; i++ {
		if i == 0 {
			g := leadingGlyphs[rand.Intn(len(leadingGlyphs))]
			s.SetCell(xLocation, currentY-i, st.Foreground(tcell.ColorWhite), g)
		} else {
			s.SetCell(xLocation, currentY-i, st.Foreground(tcell.Color40), rune('A'+rand.Intn(26)))
		}
	}

	currentY++
	if currentY-lineLength > h {
		currentY = 0
	}

	s.Show()
}
