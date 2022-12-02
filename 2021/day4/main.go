package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

const ColumnCount = 5
const RowCount = 5

type BingoBoard struct {
	Items         [RowCount][ColumnCount]BingoItem
	HasWon        bool // the number of iterations for the first "Bingo!"
	WinningNumber int
}
type BingoItem struct {
	Value   int
	Matched bool
}

func (bingoItem BingoItem) String() string {
	return fmt.Sprintf("%d", bingoItem.Value)
}

func main() {
	input_path := flag.String("input", "./input", "The input data")
	numbers, boards := getFileContents(*input_path)
	scores := getScores(numbers, boards)
	fmt.Printf("First board has won with score %d\n", scores[0])
	fmt.Printf("Last board has won with score %d\n", scores[len(scores)-1])
}

func getFileContents(filepath string) (numbers_drawn []int, boards []*BingoBoard) {
	file, err := os.Open(filepath)

	if err != nil {
		log.Fatalln(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)

	// first row --> numbers drawn
	scanner.Scan()
	strings_drawn := strings.Split(scanner.Text(), ",")
	numbers_drawn = make([]int, len(strings_drawn))
	for i, v := range strings_drawn {
		numbers_drawn[i], _ = strconv.Atoi(v)
	}
	var current_board *BingoBoard

	current_board_row := 0

	// remaining rows --> boards
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			continue
		}
		if current_board_row == 0 {
			current_board = new(BingoBoard)
		}
		column_tokens := strings.Split(line, " ")
		column := 0
		for _, v := range column_tokens {
			number, err := strconv.Atoi(v)
			if err == nil {
				current_board.Items[current_board_row][column] = BingoItem{
					Value: number}
				column++
			}
		}
		current_board_row++
		if current_board_row == RowCount {
			current_board_row = 0
			boards = append(boards, current_board)
		}
	}
	if err := scanner.Err(); err != nil {
		log.Fatalln(err)
	}
	return numbers_drawn, boards
}

func getScores(numbers []int, boards []*BingoBoard) []int {
	scores := make([]int, 0)
	for _, number := range numbers {
		for _, board := range boards {
			if board.HasWon {
				continue
			}
			board.findMatchingItemsForNumber(number)
			board.checkBoardIsBingo(number)
			if board.HasWon {
				scores = append(scores, board.calculateScore())
			}
		}
	}
	return scores
}

func (board *BingoBoard) checkBoardIsBingo(number int) {
rows:
	for row := 0; row < RowCount; row++ {
		for col := 0; col < ColumnCount; col++ {
			if !board.Items[row][col].Matched {
				continue rows
			}
		}
		board.markWinningBoard(number)
	}
columns:
	for col := 0; col < ColumnCount; col++ {
		for row := 0; row < RowCount; row++ {
			if !board.Items[row][col].Matched {
				continue columns
			}
		}
		board.markWinningBoard(number)
	}
}

func (board *BingoBoard) findMatchingItemsForNumber(number int) {
	for row := 0; row < RowCount; row++ {
		for col := 0; col < ColumnCount; col++ {
			item := &board.Items[row][col]
			if item.Value == number {
				item.Matched = true
			}
		}
	}
}

func (board *BingoBoard) markWinningBoard(number int) {
	board.HasWon = true
	board.WinningNumber = number
}

func (board *BingoBoard) calculateScore() int {
	unmarkedSum := 0
	for row := 0; row < RowCount; row++ {
		for col := 0; col < ColumnCount; col++ {
			item := &board.Items[row][col]
			if !item.Matched {
				unmarkedSum += item.Value
			}
		}
	}
	score := board.WinningNumber * unmarkedSum
	fmt.Printf("Winning score: %d\n", score)
	return score
}
