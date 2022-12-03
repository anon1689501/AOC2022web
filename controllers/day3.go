package controllers

import (
	"bufio"
	"net/http"
	"strconv"
	"strings"
)

func Day3(w http.ResponseWriter, aocInput string) {
	//read input

	scanner := bufio.NewScanner(strings.NewReader(aocInput))

	//runeArray := make([]rune, 1)
	sum := 0

	stringArray := make([]string, 0)

	for scanner.Scan() {
		mylen := len(scanner.Text()) / 2
		left := scanner.Text()[:mylen]
		right := scanner.Text()[mylen:]
		//fmt.Println(scanner.Text(), left, right)
		stringArray = append(stringArray, scanner.Text())

	out:
		for _, v := range left {
			for _, c := range right {
				if v == c {
					//runeArray = append(runeArray, v)
					if v < 96 {
						sum += int(v) - 38
					} else {
						sum += int(v) - 96
					}

					break out
				}
			}
		}

	}

	//part 1 answer

	WriteCookie(w, "day3a", strconv.Itoa(sum))

	sumB := 0
	for i := 0; i < len(stringArray); i += 3 {
	outter:
		for _, v := range stringArray[i] {
			for _, b := range stringArray[i+1] {
				if v == b {
					for _, n := range stringArray[i+2] {
						if v == n {
							//fmt.Println(v)
							if v < 96 {
								sumB += int(v) - 38
							} else {
								sumB += int(v) - 96
							}
							break outter
						}
					}
				}
			}
		}

	}

	//part 2 answer

	WriteCookie(w, "day3b", strconv.Itoa(sumB))

}
