package controllers

import (
	"bufio"
	"net/http"
	"strconv"
	"strings"
)

func Day1(w http.ResponseWriter, aocInput string) {

	caloriesPerElf := make([]int, 1)
	elf := 0

	//sum calories per elf

	scanner := bufio.NewScanner(strings.NewReader(aocInput))

	for scanner.Scan() {
		if len(scanner.Text()) == 0 {
			caloriesPerElf = append(caloriesPerElf, 0)
			//fmt.Println(caloriesPerElf[elf])
			elf++
			continue

		}
		calories, _ := strconv.Atoi(scanner.Text())
		caloriesPerElf[elf] += calories

	}

	//Day 1 part 1

	// fmt.Println("Day 1 Part 1:", FindMax(caloriesPerElf))

	WriteCookie(w, "day1a", strconv.Itoa(FindMax(caloriesPerElf)))

	//Day 1 part 2

	threeFatestElvesSum := 0
	for i := 0; i < 3; i++ {
		threeFatestElvesSum += FindMax(caloriesPerElf)
		caloriesPerElf = Remove(caloriesPerElf, FindMaxIndex(caloriesPerElf))
	}
	// fmt.Println("Day 1 Part 2:", threeFatestElvesSum)
	WriteCookie(w, "day1b", strconv.Itoa(threeFatestElvesSum))

}
func FindMax(intSlice []int) (max int) {
	max = 0
	for _, v := range intSlice {
		if v > max {
			max = v
		}
	}
	return max
}
func FindMaxIndex(intSlice []int) (index int) {
	maxSize := 0
	index = 0
	for i, v := range intSlice {
		if v > maxSize {
			maxSize = v
			index = i
		}
	}
	return index
}

func Remove(s []int, i int) []int {
	s[i] = s[len(s)-1]
	return s[:len(s)-1]
}
