package controllers

import (
	"bufio"
	"net/http"
	"strconv"
	"strings"
)

func Day2pA(w http.ResponseWriter, aocInput string) {

	//read input

	scanner := bufio.NewScanner(strings.NewReader(aocInput))

	game := ""
	points := 0

	for scanner.Scan() {
		game = scanner.Text()
		if game[0] == 65 && game[2] == 88 || game[0] == 66 && game[2] == 89 || game[0] == 67 && game[2] == 90 {
			points += 3
		}
		if game[0] == 65 && game[2] == 89 || game[0] == 66 && game[2] == 90 || game[0] == 67 && game[2] == 88 {
			points += 6
		}
		points += int(game[2]) - 87

	}

	// fmt.Println("Day 2 Part 1:", points)
	WriteCookie(w, "day2a", strconv.Itoa(points))

}

func Day2pB(w http.ResponseWriter, aocInput string) {

	//read input

	scanner := bufio.NewScanner(strings.NewReader(aocInput))

	game := ""
	points := 0

	for scanner.Scan() {
		game = scanner.Text()
		if game[2] == 89 {
			points += 3
		}
		if game[2] == 90 {
			points += 6
		}

		//rock conditions
		if game[0] == 65 && game[2] == 88 {
			points += 3
		}
		if game[0] == 65 && game[2] == 89 {
			points += 1
		}
		if game[0] == 65 && game[2] == 90 {
			points += 2
		}

		// paper conditions
		if game[0] == 66 && game[2] == 88 {
			points += 1
		}
		if game[0] == 66 && game[2] == 89 {
			points += 2
		}
		if game[0] == 66 && game[2] == 90 {
			points += 3
		}

		// scissor conditions
		if game[0] == 67 && game[2] == 88 {
			points += 2
		}
		if game[0] == 67 && game[2] == 89 {
			points += 3
		}
		if game[0] == 67 && game[2] == 90 {
			points += 1
		}
	}

	// fmt.Println("Day 2 Part 2:", points)
	WriteCookie(w, "day2b", strconv.Itoa(points))

}
