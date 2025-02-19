package sol

import (
	"log"

	"oddstream.games/gosol/sound"
)

func FindCardOwner(card *Card) Pile {
	for _, pile := range TheBaize.piles {
		for _, c := range pile.Cards() {
			if c == card {
				return pile
			}
		}
	}
	// log.Panic("Cannot find card")
	return nil
}

func AnyCardsProne(cards []*Card) bool {
	for _, c := range cards {
		if c.Prone() {
			return true
		}
	}
	return false
}

func FlipUpExposedCard(p Pile) {
	if _, isStock := (p).(*Stock); !isStock {
		if c := p.Peek(); c != nil {
			c.FlipUp()
		}
	}
}

// MoveCard is an optimized, single card version of MoveCards
func MoveCard(src Pile, dst Pile) *Card {
	if c := src.Pop(); c != nil {
		sound.Play("Place")
		dst.Push(c)
		FlipUpExposedCard(src)
		TheBaize.setFlag(dirtyCardPositions)
		return c
	}
	return nil
}

func MoveNamedCard(src Pile, suit, ordinal int, dst Pile) {

	// 1. find the card in the src Pile
	var ID CardID = NewCardID(0, suit, ordinal)
	var card *Card
	for _, c := range src.Cards() {
		if SameCard(ID, c.ID) {
			card = c
			break
		}
	}
	if card == nil {
		println("Could not find card in src Pile")
		return
	}

	// 2.find the card in it's owning pile
	var index int = src.IndexOf(card)
	if index == -1 {
		println("Could not find card", card.String(), "in pile")
		return
	}

	// 3. extract the card from it's owning pile
	src.Delete(index)

	// 4. push the card onto the dst pile
	sound.Play("Place")
	card.FlipUp()
	dst.Push(card)
	FlipUpExposedCard(src)
	TheBaize.setFlag(dirtyCardPositions)
}

// MoveCards is used when dragging a tail from ome pile to another
func MoveCards(src Pile, moveFromIndex int, dst Pile) {

	oldSrcLen := src.Len()

	tmp := make([]*Card, 0, oldSrcLen)

	// pop the tail off the source and push onto temp stack
	for i := oldSrcLen - 1; i >= moveFromIndex; i-- {
		tmp = append(tmp, src.Pop())
	}

	sound.Play("Slide")

	// pop all cards off the temp stack and onto the destination
	for i := len(tmp) - 1; i >= 0; i-- {
		dst.Push(tmp[i])
	}

	FlipUpExposedCard(src)

	if oldSrcLen == src.Len() {
		log.Println("nothing happened in MoveCards")
	}

	TheBaize.setFlag(dirtyCardPositions)
}

func MoveAllCards(src Pile, dst Pile) {
	if src.Empty() {
		return
	}
	for i := 0; i < src.Len(); i++ {
		dst.Push(src.Get(i))
	}
	src.Reset()
	TheBaize.setFlag(dirtyCardPositions)
}
