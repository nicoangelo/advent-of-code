package camelcards

type HandValuator interface {
	getCardValue(Card) int
	getHandValue(map[Card]int) int
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

type ClassicHandValuator struct {
}

func (hv *ClassicHandValuator) getCardValue(c Card) int {
	return defaultCardValues[c]
}

func (hv *ClassicHandValuator) getHandValue(cardCounts map[Card]int) int {
	highestCount := 0
	for _, v := range cardCounts {
		if v > highestCount {
			highestCount = v
		}
	}
	if len(cardCounts) == 1 {
		return FiveOfAKind
	} else if len(cardCounts) == 2 && highestCount == 4 {
		// Four of a kind: 4+1
		return FourOfAKind
	} else if len(cardCounts) == 2 {
		// Full House: 3+2
		return FullHouse
	} else if len(cardCounts) == 3 && highestCount == 3 {
		// Three of a kind: 3+1+1
		return ThreeOfAKind
	} else if len(cardCounts) == 3 {
		// Two pair: 2+2+1
		return TwoPair
	} else if len(cardCounts) == 4 {
		// One pair: 2+1+1+1
		return OnePair
	}
	return HighCard
}

type JokerHandValuator struct {
	JokerCard Card
}

func (hv *JokerHandValuator) getCardValue(c Card) int {
	if c == hv.JokerCard {
		return 1
	}
	return defaultCardValues[c]
}

func (hv *JokerHandValuator) getHandValue(cardCounts map[Card]int) int {
	highestCount := 0
	for _, v := range cardCounts {
		if v > highestCount {
			highestCount = v
		}
	}
	if len(cardCounts) == 1 {
		// nothing to improve
		return FiveOfAKind
	} else if len(cardCounts) == 2 && highestCount == 4 {
		// Four of a kind: 4+1
		// either one of the card types is a Joker
		if cardCounts[hv.JokerCard] > 0 {
			return FiveOfAKind
		}
		return FourOfAKind
	} else if len(cardCounts) == 2 {
		// Full House: 3+2
		// either one of the card types is a Joker
		if cardCounts[hv.JokerCard] >= 2 {
			return FiveOfAKind
		}
		return FullHouse
	} else if len(cardCounts) == 3 && highestCount == 3 {
		// Three of a kind: 3+1+1
		// either one of the card types is a Joker
		if cardCounts[hv.JokerCard] > 0 {
			return FourOfAKind
		}
		return ThreeOfAKind
	} else if len(cardCounts) == 3 {
		// Two pair: 2+2+1
		// one of the two pairs is of type Joker
		if cardCounts[hv.JokerCard] == 2 {
			return FourOfAKind
		}
		// ..OR the single remaining card is a Joker
		if cardCounts[hv.JokerCard] == 1 {
			return FullHouse
		}
		return TwoPair
	} else if len(cardCounts) == 4 {
		// One pair: 2+1+1+1
		// either one of the card types is a Joker
		// (note: splitting 2 jokers to make a TwoPair is less powerful than ThreeOfAKind)
		if cardCounts[hv.JokerCard] > 0 {
			return ThreeOfAKind
		}
		return OnePair
	}
	// if all cards are different, but there's one Joker, we can make it a Pair
	if cardCounts[hv.JokerCard] == 1 {
		return OnePair
	}
	return HighCard
}
