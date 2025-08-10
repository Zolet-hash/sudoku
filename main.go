package main

import (
	"fmt"
	"os"
)

func boxIndex(r, c int) int { return (r/3)*3 + c/3 }

func printErrorAndExit() {
	fmt.Println("Error")
	os.Exit(0)
}

func main() {
	args := os.Args[1:]
	if len(args) != 9 {
		printErrorAndExit()
	}

	var board [9][9]int
	var rows [9][10]bool
	var cols [9][10]bool
	var boxes [9][10]bool

	// parse and validate
	for r := 0; r < 9; r++ {
		row := args[r]
		if len(row) != 9 {
			printErrorAndExit()
		}
		for c := 0; c < 9; c++ {
			ch := row[c]
			if ch == '.' {
				board[r][c] = 0
				continue
			}
			if ch < '1' || ch > '9' {
				printErrorAndExit()
			}
			n := int(ch - '0')
			b := boxIndex(r, c)
			if rows[r][n] || cols[c][n] || boxes[b][n] {
				// duplicate in initial board
				printErrorAndExit()
			}
			board[r][c] = n
			rows[r][n] = true
			cols[c][n] = true
			boxes[b][n] = true
		}
	}

	// solver with uniqueness check (stop at 2 solutions)
	var solutions int
	var solutionBoard [9][9]int

	// helper to copy current board
	copySolution := func() {
		for i := 0; i < 9; i++ {
			for j := 0; j < 9; j++ {
				solutionBoard[i][j] = board[i][j]
			}
		}
	}

	// find next empty using MRV (minimum remaining values) Ai algorithms
	var dfs func()
	dfs = func() {
		if solutions >= 2 {
			return
		}

		// find empty cell; use MRV heuristic
		minCandidates := 10
		rSel, cSel := -1, -1
		for r := 0; r < 9; r++ {
			for c := 0; c < 9; c++ {
				if board[r][c] != 0 {
					continue
				}
				count := 0
				b := boxIndex(r, c)
				for n := 1; n <= 9; n++ {
					if !rows[r][n] && !cols[c][n] && !boxes[b][n] {
						count++
					}
				}
				if count == 0 {
					// dead end
					return
				}
				if count < minCandidates {
					minCandidates = count
					rSel = r
					cSel = c
					if count == 1 {
						// can't get better than 1, but continue scanning to check for 0
					}
				}
			}
		}

		if rSel == -1 {
			// no empty found -> solution
			solutions++
			if solutions == 1 {
				copySolution()
			}
			return
		}

		bSel := boxIndex(rSel, cSel)
		for n := 1; n <= 9; n++ {
			if rows[rSel][n] || cols[cSel][n] || boxes[bSel][n] {
				continue
			}
			// place
			board[rSel][cSel] = n
			rows[rSel][n] = true
			cols[cSel][n] = true
			boxes[bSel][n] = true

			dfs()

			// undo
			board[rSel][cSel] = 0
			rows[rSel][n] = false
			cols[cSel][n] = false
			boxes[bSel][n] = false

			if solutions >= 2 {
				return
			}
		}
	}

	dfs()

	if solutions != 1 {
		printErrorAndExit()
	}

	// print solved board (stored in solutionBoard)
	for r := 0; r < 9; r++ {
		for c := 0; c < 9; c++ {
			if c > 0 {
				fmt.Print(" ")
			}
			fmt.Print(solutionBoard[r][c])
		}
		fmt.Println()
	}
}
