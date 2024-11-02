package main

import (
	"fmt"
	"os"
)

const N = 9

func main() {
	if len(os.Args) != 10 {
		fmt.Println("Error")
		return
	}

	var grid [N][N]int

	for i := 0; i < N; i++ {
		row := os.Args[i+1]
		if len(row) != N {
			fmt.Println("Error")
			return
		}
		for j := 0; j < N; j++ {
			c := row[j]
			if c == '.' {
				grid[i][j] = 0
			} else if c >= '1' && c <= '9' {
				grid[i][j] = int(c - '0')
			} else {
				fmt.Println("Error")
				return
			}
		}
	}

	if solveSudoku(&grid, 0, 0) {
		printGrid(&grid)
	} else {
		fmt.Println("Error")
	}
}

func printGrid(grid *[N][N]int) {
	for i := 0; i < N; i++ {
		for j := 0; j < N; j++ {
			fmt.Printf("%d", grid[i][j])
			if j != N-1 {
				fmt.Print(" ")
			}
		}
		fmt.Println()
	}
}

func solveSudoku(grid *[N][N]int, row, col int) bool {
	if row == N {
		return true
	}
	if col == N {
		return solveSudoku(grid, row+1, 0)
	}
	if grid[row][col] != 0 {
		return solveSudoku(grid, row, col+1)
	}
	for num := 1; num <= 9; num++ {
		if isSafe(grid, row, col, num) {
			grid[row][col] = num
			if solveSudoku(grid, row, col+1) {
				return true
			}
			grid[row][col] = 0
		}
	}
	return false
}

func isSafe(grid *[N][N]int, row, col, num int) bool {
	// Check row and column
	for i := 0; i < N; i++ {
		if grid[row][i] == num || grid[i][col] == num {
			return false
		}
	}
	// Check 3x3 box
	startRow := row - row%3
	startCol := col - col%3
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if grid[startRow+i][startCol+j] == num {
				return false
			}
		}
	}
	return true
}
