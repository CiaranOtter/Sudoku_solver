package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
)

var board = make([][]int, 9)

func read_file(name string) string {
	file, err := os.Open(name)

	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	var line string
	for scanner.Scan() {
		line = scanner.Text()
		format_string(line)
		print_borad()
		backTrack(0, 0)
		print_borad()
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return line
}

func format_string(line string) {
	for i, c := range line {

		if i >= 81 {
			break
		}

		var row float64

		row = float64(i / 9)

		if i%9 == 0 {
			board[int(math.Round(math.Floor(row)))] = make([]int, 9)
		}

		if c == '.' {
			board[int(math.Round(math.Floor(row)))][i%9] = -1
		} else {
			board[int(math.Round(math.Floor(row)))][i%9] = int(c - '0')
		}
	}
}

func print_borad() {
	fmt.Print("+-+-+-+-+-+-+-+-+-+\n")
	for _, row := range board {
		fmt.Printf("|")
		for _, col := range row {

			if col == -1 {
				fmt.Print(" |")
			} else {
				fmt.Printf("%d|", col)
			}
		}
		fmt.Print("\n+-+-+-+-+-+-+-+-+-+\n")
	}
}

func pretty_print(line string) {
	for i, c := range line {

		if c == '\n' {
			continue
		}
		if i%9 == 0 {
			fmt.Printf("\n")
			fmt.Printf("--------------------\n|")
		}

		if c == '.' {
			fmt.Printf(" |")
		} else {
			fmt.Printf("%c|", c)
		}
	}
}

func check_row(row int, poss []int) []int {
	for _, cell := range board[row] {
		for i, val := range poss {
			if cell == val {
				poss = append(poss[:i], poss[i+1:]...)
			}
		}
	}

	return poss
}

func check_col(col int, poss []int) []int {
	for _, cell := range board {
		for i, val := range poss {
			if cell[col] == val {
				poss = append(poss[:i], poss[i+1:]...)
			}
		}
	}

	return poss
}

func check_square(col int, row int, poss []int) []int {
	start_col := int(math.Floor(float64(col/3))) * 3
	start_row := int(math.Floor(float64(row/3))) * 3

	// fmt.Printf("start row: %d, start col: %d\n", start_row, start_col)

	for i := start_row; i < (start_row)+3; i++ {
		for j := start_col; j < (start_col)+3; j++ {
			// fmt.Printf("Checking %d %d\n", i, j)
			for c, val := range poss {
				if val == board[i][j] {
					poss = append(poss[:c], poss[c+1:]...)
				}
			}
		}
	}

	return poss
}

func find_possible(row int, col int) []int {
	poss := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}

	poss = check_row(row, poss)
	// fmt.Printf("For %d %d: ", row, col)
	// for _, p := range poss {
	// 	fmt.Print(p)
	// }

	// fmt.Print("\n")

	poss = check_col(col, poss)
	// fmt.Printf("For %d %d: ", row, col)

	// for _, p := range poss {
	// 	fmt.Print(p)
	// }

	// fmt.Print("\n")
	poss = check_square(col, row, poss)
	// fmt.Printf("For %d %d: ", row, col)

	// for _, p := range poss {
	// 	fmt.Print(p)
	// }

	// fmt.Print("\n")
	return poss
}

func find_next(row int, col int) (int, int) {
	if col == 8 {
		// fmt.Printf("HEading to next row\n")
		col = 0
		row++
	} else {
		col++
	}

	return row, col
}

func backTrack(row int, col int) bool {

	if board[row][col] != -1 {
		// fmt.Printf("There is a value at %d %d\n", row, col)

		if row == 8 && col == 8 {
			return true
		}
		next_row, next_col := find_next(row, col)
		return backTrack(next_row, next_col)
	}
	poss := find_possible(row, col)

	if row == 8 && col == 8 && len(poss) == 1 {
		board[row][col] = poss[0]
		return true
	}

	next_row, next_col := find_next(row, col)

	// fmt.Printf("Next row: %d, Next col: %d\n", next_row, next_col)

	for _, p := range poss {
		board[row][col] = p
		// fmt.Printf("Testing %d at position %d %d with possible numbers %d\n", p, row, col, len(poss))

		if backTrack(next_row, next_col) == true {
			return true
		}

	}

	board[row][col] = -1
	return false
	// for i := 0; i < 9; i++ {
	// 	for j := 0; j < 9; j++ {
	// 		find_possible(i, j)
	// 		fmt.Print("\n")
	// 	}
	// }
}

func main() {
	// pretty_print(read_file("../data/easy.txt"))
	read_file("../data/easy.txt")
	read_file("../data/intermediate.txt")
	read_file("../data/hard.txt")
	read_file("../data/insane.txt")
	// print_borad()

	// backTrack(0, 0)

	// print_borad()
}
