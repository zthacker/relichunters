package renderer

import (
	"github.com/gdamore/tcell/v2"
	"relichunters/internal/models"
)

type TCellRenderer struct {
	screen tcell.Screen
}

func (t *TCellRenderer) DrawTextStyled(x, y int, text string, s *models.Style) error {
	if t.screen == nil {
		return nil
	}
	style := tcell.StyleDefault.Foreground(tcell.Color(s.ForegroundColor)).Background(tcell.Color(s.BackgroundColor))
	for i, ch := range text {
		t.screen.SetContent(x+i, y, ch, nil, style)
	}
	return nil
}

func (t *TCellRenderer) DrawBox(x, y, w, h int, style *models.Style) error {
	// Top and bottom
	s := tcell.StyleDefault.Foreground(tcell.Color(style.ForegroundColor)).Background(tcell.Color(style.BackgroundColor))
	for i := 0; i < w; i++ {
		t.screen.SetContent(x+i, y, '─', nil, s)
		t.screen.SetContent(x+i, y+h-1, '─', nil, s)
	}
	// Left and right
	for j := 0; j < h; j++ {
		t.screen.SetContent(x, y+j, '│', nil, s)
		t.screen.SetContent(x+w-1, y+j, '│', nil, s)
	}
	// Corners
	t.screen.SetContent(x, y, '┌', nil, s)
	t.screen.SetContent(x+w-1, y, '┐', nil, s)
	t.screen.SetContent(x, y+h-1, '└', nil, s)
	t.screen.SetContent(x+w-1, y+h-1, '┘', nil, s)

	return nil
}

func (t *TCellRenderer) DrawLine(x1, y1, x2, y2 int, style *models.Style) error {
	//TODO implement me
	panic("implement me")
}

func (t *TCellRenderer) DrawImage(img string, x, y int) error {
	//TODO implement me
	panic("implement me")
}

func NewTCellRenderer(screen tcell.Screen) *TCellRenderer {
	return &TCellRenderer{screen}
}

func (t *TCellRenderer) Init() error {
	return nil
}

func (t *TCellRenderer) Clear() error {
	if t.screen == nil {
		return nil
	}
	t.screen.Clear()
	return nil
}

func (t *TCellRenderer) DrawText(x, y int, text string) error {
	if t.screen == nil {
		return nil
	}

	style := tcell.StyleDefault
	for i, ch := range text {
		t.screen.SetContent(x+i, y, ch, nil, style)
	}
	return nil
}

func (t *TCellRenderer) DrawSprites() error {
	//TODO implement me
	panic("no sprites to draw -- implement me")
}

func (t *TCellRenderer) Present() error {
	if t.screen == nil {
		return nil
	}
	t.screen.Show()
	return nil
}

func (t *TCellRenderer) Stop() error {
	if t.screen != nil {
		t.screen.Fini()
	}
	return nil
}

func (t *TCellRenderer) GetSize() (int, int) {
	return t.screen.Size()
}
