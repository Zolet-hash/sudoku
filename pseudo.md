if len(arg) != 9: Error
board = parse args into 9 * 9 ints ('.' -> 0)

if any row/col/box invalid: Error

solution = 0
solutionBpoard = empty

func backtrack(): 
	if no empty cells: 
		solution += 1
		if solution == 1: save board to solutionboard
		return

	choose cell (prefer less candidates)
			if val valid:
			place val, update trackers backtrack
			undo place
			if solutions >= 2 : return


if solutions == 1: print solutionboard in required format
	else print Error

	sudoku pseudocode
