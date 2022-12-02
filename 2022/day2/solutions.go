package day2

// Rock > Scissors, Scissors > Paper, Paper > Rock

const (
	BONUS_LOOSE = 0
	BONUS_DRAW  = 3
	BONUS_WIN   = 6
	ROCK        = 1
	PAPER       = 2
	SCISSOR     = 3
)

// PART1
// A|X=Rock     1
// B|Y=Paper    2
// C|Z=Scissor  3
// Bonus: loose = 0, draw = 3, win = 6
var rspWinMap = map[string]int{
	"A X": ROCK + BONUS_DRAW,
	"A Y": PAPER + BONUS_WIN,
	"A Z": SCISSOR + BONUS_LOOSE,
	"B X": ROCK + BONUS_LOOSE,
	"B Y": PAPER + BONUS_DRAW,
	"B Z": SCISSOR + BONUS_WIN,
	"C X": ROCK + BONUS_WIN,
	"C Y": PAPER + BONUS_LOOSE,
	"C Z": SCISSOR + BONUS_DRAW,
}

func part1(lines []string) int {
	return getWinSum(lines, rspWinMap)
}

func part2(lines []string) int {
	return getWinSum(lines, rspWinOutcomeMap)
}

func getWinSum(lines []string, winMap map[string]int) (sum int) {
	for _, v := range lines {
		sum += winMap[v]
	}
	return sum
}

// PART2
// A=Rock     1
// B=Paper    2
// C=Scissor  3
// X=Loose
// Y=Draw
// Z=Win
// Bonus: loose = 0, draw = 3, win = 6
var rspWinOutcomeMap = map[string]int{
	"A X": SCISSOR + BONUS_LOOSE,
	"A Y": ROCK + BONUS_DRAW,
	"A Z": PAPER + BONUS_WIN,
	"B X": ROCK + BONUS_LOOSE,
	"B Y": PAPER + BONUS_DRAW,
	"B Z": SCISSOR + BONUS_WIN,
	"C X": PAPER + BONUS_LOOSE,
	"C Y": SCISSOR + BONUS_DRAW,
	"C Z": ROCK + BONUS_WIN,
}
