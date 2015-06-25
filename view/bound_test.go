package view

import (
	"testing"
)

func TestPoint(t *testing.T) {
	p := Point{100, 200}
	if p.X != 100 {
		t.Errorf("Point X value do not match")
	}
	if p.Y != 200 {
		t.Errorf("Point Y value do not match")
	}
}

func TestRect(t *testing.T) {
	r := Rect{Point{1, 2}, Point{3, 4}}
	if r.TopLeft.X != 1 {
		t.Errorf("Rect X1 value do not match")
	}
	if r.TopLeft.Y != 2 {
		t.Errorf("Rect Y1 value do not match")
	}
	if r.BottomRight.X != 3 {
		t.Errorf("Rect X2 value do not match")
	}
	if r.BottomRight.Y != 4 {
		t.Errorf("Rect Y2 value do not match")
	}
}
