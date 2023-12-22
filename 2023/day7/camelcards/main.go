package camelcards

import (
	"fmt"
)

type Card rune

const (
	CardCount    = 5
	HighCard     = 1
	OnePair      = 2
	TwoPair      = 3
	ThreeOfAKind = 4
	FullHouse    = 5
	FourOfAKind  = 6
	FiveOfAKind  = 7
)

type Hand struct {
	cardTypeCounts map[Card]int
	cards          [CardCount]Card
	cardValues     [CardCount]int
	bid            int
	handType       int
}

func (h Hand) String() string {
	return fmt.Sprintf("%c / handType: %s", h.cards, getHandName(h.handType))
}

func (h *Hand) GetBid() int {
	return h.bid
}

func (h *Hand) FromLine(l string, vl HandValuator) {
	h.cardTypeCounts = make(map[Card]int)
	for i, r := range l {
		if i < CardCount {
			c := Card(r)
			h.cardTypeCounts[c]++
			h.cards[i] = c
			h.cardValues[i] = vl.getCardValue(c)
		}
		if i > CardCount {
			h.bid *= 10
			h.bid += int(r - '0')
		}
	}
	h.handType = vl.getHandValue(h.cardTypeCounts)
}

func (h1 *Hand) Compare(h2 *Hand) int {
	if h1.handType > h2.handType {
		return 1
	}
	if h1.handType < h2.handType {
		return -1
	}
	return h1.compareCards(h2)
}

func (h1 *Hand) compareCards(h2 *Hand) int {
	for i := 0; i < CardCount; i++ {
		if h1.cardValues[i] > h2.cardValues[i] {
			return 1
		}
		if h1.cardValues[i] < h2.cardValues[i] {
			return -1
		}
	}
	return 0
}

func getHandName(handType int) string {
	switch handType {
	case 1:
		return "High Card"
	case 2:
		return "One Pair"
	case 3:
		return "Two Pair"
	case 4:
		return "Three Of A Kind"
	case 5:
		return "Full House"
	case 6:
		return "Four Of A Kind"
	case 7:
		return "Five Of A Kind"
	}
	return ""
}
