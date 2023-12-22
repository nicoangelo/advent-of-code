package camelcards

type HandValuator interface {
	getCardValues() map[Card]int
	getCardValue(Card) int
	getHandValue(map[Card]int) int
}

type ClassicHandValuator struct {
}

var defaultCardValues = map[Card]int{
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

func (hv *ClassicHandValuator) getCardValues() map[Card]int {
	return defaultCardValues
}

func (hv *ClassicHandValuator) getCardValue(c Card) int {
	return defaultCardValues[c]
}

func (hv *ClassicHandValuator) getHandValue(cv map[Card]int) int {
	highestCount := 0
	for _, v := range cv {
		if v > highestCount {
			highestCount = v
		}
	}
	if len(cv) == 1 {
		return FiveOfAKind
	} else if len(cv) == 2 && highestCount == 4 {
		// Four of a kind: 4+1
		return FourOfAKind
	} else if len(cv) == 2 {
		// Full House: 3+2
		return FullHouse
	} else if len(cv) == 3 && highestCount == 3 {
		// Three of a kind: 3+1+1
		return ThreeOfAKind
	} else if len(cv) == 3 {
		// Two pair: 2+2+1
		return TwoPair
	} else if len(cv) == 4 {
		// One pair: 2+1+1+1
		return OnePair
	}
	return HighCard
}

type JokerHandValuator struct {
}
