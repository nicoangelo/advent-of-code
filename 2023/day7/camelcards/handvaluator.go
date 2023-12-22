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
	JokerCard Card
}

func (hv *JokerHandValuator) getCardValue(c Card) int {
	if c == hv.JokerCard {
		return 1
	}
	return defaultCardValues[c]
}

func (hv *JokerHandValuator) getHandValue(cv map[Card]int) int {
	highestCount := 0
	for _, v := range cv {
		if v > highestCount {
			highestCount = v
		}
	}
	if len(cv) == 1 {
		// nothing to improve
		return FiveOfAKind
	} else if len(cv) == 2 && highestCount == 4 {
		// Four of a kind: 4+1
		if cv[hv.JokerCard] > 0 { // remaining card type is a Joker
			return FiveOfAKind
		}
		return FourOfAKind
	} else if len(cv) == 2 {
		// Full House: 3+2
		if cv[hv.JokerCard] >= 2 { // remaining two cards are a Joker
			return FiveOfAKind
		}
		return FullHouse
	} else if len(cv) == 3 && highestCount == 3 {
		// Three of a kind: 3+1+1
		if cv[hv.JokerCard] == 1 || cv[hv.JokerCard] == 3 { // one of the two remaining card types are a Joker
			return FourOfAKind
		}
		return ThreeOfAKind
	} else if len(cv) == 3 {
		// Two pair: 2+2+1
		if cv[hv.JokerCard] == 2 { // one of the two pairs is of type Joker
			return FourOfAKind
		}
		if cv[hv.JokerCard] == 1 { // the single remaining card is a Joker
			return FullHouse
		}
		return TwoPair
	} else if len(cv) == 4 {
		// One pair: 2+1+1+1
		// one of the remaining cards OR the pair can be a joker
		// splitting 2 jokers to make a TwoPair is less powerful than ThreeOfAKind
		if cv[hv.JokerCard] == 1 || cv[hv.JokerCard] == 2 {
			return ThreeOfAKind
		}
		return OnePair
	}
	return HighCard
}
