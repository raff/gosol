package ui

import (
	"github.com/fogleman/gg"
	"github.com/hajimehoshi/ebiten/v2"
	"golang.org/x/image/font"
	"oddstream.games/gosol/input"
	"oddstream.games/gosol/util"
)

// Label is a button that displays a single rune
type Label struct {
	WidgetBase
	text        string
	fontFace    font.Face
	requestType string
}

func (l *Label) createImg() *ebiten.Image {
	// println("Label createImg", l.x, l.y, l.width, l.height, l.text)
	if l.width == 0 || l.height == 0 {
		return nil // prevent ebiten.NewImageFromImage panic
	}
	dc := gg.NewContext(l.width, l.height)
	dc.SetRGBA(1, 1, 1, 1)
	dc.SetFontFace(l.fontFace)
	// nota bene - text is drawn with y as a baseline, descenders may be clipped
	dc.DrawString(l.text, 0, float64(l.height)*0.8)

	// uncomment this to show the area we expect the text to occupy
	// dc.DrawLine(0, float64(0), float64(l.width), float64(0))
	// dc.DrawLine(0, float64(l.height), float64(l.width), float64(l.height))
	// dc.DrawLine(0, float64(0), float64(l.width), float64(l.height))
	// dc.Stroke()

	return ebiten.NewImageFromImage(dc.Image())
}

func measureText(text string, fontFace font.Face) (int, int) {
	dc := gg.NewContext(8, 8)
	dc.SetFontFace(fontFace)
	width, height := dc.MeasureString(text)
	return int(width), int(height)
}

// NewLabel creates a new Label
func NewLabel(parent Container, align int, text string, fontFace font.Face, requestType string) *Label {

	width, height := measureText(text, fontFace)

	l := &Label{
		// widget x, y will be set by LayoutWidgets
		WidgetBase: WidgetBase{parent: parent, img: nil, width: int(width), height: int(height), align: align},
		text:       text, fontFace: fontFace, requestType: requestType}
	l.Activate()
	return l
}

// Activate tells the input we need notifications
func (l *Label) Activate() {
	l.disabled = false
	l.img = l.createImg()
}

// Deactivate tells the input we no longer need notofications
func (l *Label) Deactivate() {
	l.disabled = true
	l.img = l.createImg()
}

// NotifyCallback is called by the Subject (Input/Stroke) when something interesting happens
func (l *Label) NotifyCallback(v input.StrokeEvent) {
	if l.disabled {
		return
	}
	switch v.Event {
	case input.Tap:
		if l.requestType != "" && util.InRect(v.X, v.Y, l.OffsetRect) {
			// println("label notify", l.requestType, ":=", l.text)
			cmdFn(ChangeRequest{ChangeRequested: l.requestType, Data: l.text})
		}
	}
}

func (l *Label) UpdateText(text string) {
	if l.text != text {
		l.text = text
		l.width, l.height = measureText(l.text, l.fontFace)
		// println("text:", text, "width:", l.width, "height:", l.height)
		l.img = l.createImg()
	}
}
