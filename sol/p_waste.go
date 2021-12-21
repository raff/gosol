package sol

//lint:file-ignore ST1005 Error messages are toasted, so need to be capitalized

import (
	"errors"
	"image"
)

type Waste struct {
	pile *Pile
}

func NewWaste(slot image.Point, fanType FanType) *Pile {
	p := &Pile{}
	p.Ctor(&Waste{pile: p}, "Waste", slot, fanType)
	return p
}

func (*Waste) CanMoveTail(tail []*Card) (bool, error) {
	if len(tail) > 1 {
		return false, errors.New("Can only move a single card from Waste")
	}
	if AnyCardsProne(tail) {
		return false, errors.New("Cannot move a face down card")
	}
	return true, nil
}

func (w *Waste) CanAcceptCard(card *Card) (bool, error) {
	if _, ok := (card.owner.subtype).(*Stock); !ok {
		return false, errors.New("Waste can only accept cards from the Stock")
	}
	return true, nil
}

func (w *Waste) CanAcceptTail(tail []*Card) (bool, error) {
	if _, ok := (tail[0].owner.subtype).(*Stock); !ok {
		return false, errors.New("Waste can only accept cards from the Stock")
	}
	return true, nil
}

func (w *Waste) TailTapped(tail []*Card) {
	w.pile.GenericTailTapped(tail)
}

func (w *Waste) Collect() {
	w.pile.GenericCollect()
}

func (w *Waste) Conformant() bool {
	// zero or one cards would be conformant
	return w.pile.Len() < 2
}

func (w *Waste) Complete() bool {
	return w.pile.Empty()
}

func (w *Waste) UnsortedPairs() int {
	// Waste (like Stock, Reserve) is always considered unsorted
	if w.pile.Empty() {
		return 0
	}
	return w.pile.Len() - 1
}
