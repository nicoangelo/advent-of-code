package camelcards

import (
	"fmt"
)

type Card rune

var cardValues = map[Card]int{
	'A': 14,
	'K': 13,
	'Q': 12,
	'J': 11,
	'T': 10,
	'9': 9,
	'8': 8,
	'7': 7,
	'6': 6,
	'5': 5,
	'4': 4,
	'3': 3,
	'2': 2,
}

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
	cardMap    map[Card]int
	cards      [CardCount]Card
	cardValues [CardCount]int
	bid        int
	handType   int
}

func (h Hand) String() string {
	return fmt.Sprintf("%c / handType: %s", h.cards, getHandName(h.handType))
}

func (h *Hand) GetBid() int {
	return h.bid
}

func (h *Hand) FromLine(l string) {
	h.cardMap = make(map[Card]int)
	for i, r := range l {
		if i < CardCount {
			c := Card(r)
			h.cardMap[c]++
			h.cards[i] = c
			h.cardValues[i] = cardValues[c]
		}
		if i > CardCount {
			h.bid *= 10
			h.bid += int(r - '0')
		}
	}
	h.handType = h.getHandType()
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

func (h *Hand) getHandType() int {
	highestCount := 0
	for _, v := range h.cardMap {
		if v > highestCount {
			highestCount = v
		}
	}
	if len(h.cardMap) == 1 {
		return FiveOfAKind
	} else if len(h.cardMap) == 2 && highestCount == 4 {
		// Four of a kind: 4+1
		return FourOfAKind
	} else if len(h.cardMap) == 2 {
		// Full House: 3+2
		return FullHouse
	} else if len(h.cardMap) == 3 && highestCount == 3 {
		// Three of a kind: 3+1+1
		return ThreeOfAKind
	} else if len(h.cardMap) == 3 {
		// Two pair: 2+2+1
		return TwoPair
	} else if len(h.cardMap) == 4 {
		// One pair: 2+1+1+1
		return OnePair
	}
	return HighCard
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
