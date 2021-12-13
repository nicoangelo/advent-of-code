package main

type Entry struct {
	Signals             []*Signal
	OutputValuePatterns []*Signal
}

func (entry Entry) String() string {
	res := ""
	for _, v := range entry.Signals {
		res += v.String()
	}
	res += "|"
	for _, v := range entry.OutputValuePatterns {
		res += v.String()
	}
	return res
}

func (entry *Entry) GetSignalOfDigit(digit int) *Signal {
	for _, signal := range entry.Signals {
		if signal.Digit == digit {
			return signal
		}
	}
	return nil
}
