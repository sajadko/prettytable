package main

import (
	"fmt"
	"sort"
)

type Table struct {
	values       [][]string
	headers      []string
	lockedHeader bool
	bodyTheme    string
	color        string
	headerSplit  bool
	paddingRight int
	paddingLeft  int
}

func (t *Table) pushValue(values []string) {
	t.values = append(t.values, values)
}

func (t *Table) print() {

	//Calculate Maximum String Length
	var maxChars []int
	var valuesAndHeaders = t.values
	valuesAndHeaders = append(valuesAndHeaders, t.headers)
	for index := range valuesAndHeaders {
		var theChars []int
		for index2 := range valuesAndHeaders {
			if index > (len(valuesAndHeaders[index2]) - 1) {
				break
			} else {
				theChars = append(theChars, len(valuesAndHeaders[index2][index]))
			}
		}
		sort.Ints(theChars)
		lastIndexOfTheCharsSlice := len(theChars)
		if len(theChars) > 0 {
			maxChars = append(maxChars, theChars[lastIndexOfTheCharsSlice-1])
		}
	}

	//fmt.Println(maxChars)

	//Add Space To Values
	for index, value := range t.values {
		for index2, value2 := range value {
			diff := len(value2) - maxChars[index2]
			if -(diff) != 0 {
				var spaces = ""
				for i := 1; i <= -(diff); i++ {
					spaces += " "
				}

				t.values[index][index2] = value2 + spaces
			}
		}
	}

	//Add Space To Values
	for index, value := range t.headers {
		diff := len(value) - maxChars[index]
		if -(diff) != 0 {
			var spaces = ""
			for i := 1; i <= -(diff); i++ {
				spaces = spaces + " "
			}
			t.headers[index] = value + spaces
		}
	}

	//fmt.Println(t.values)

	var theme map[string]string
	switch t.bodyTheme {
	case "t1":
		theme = bodyCharsT1
		break
	case "t2":
		theme = bodyCharsT2
		break
	default:
		theme = bodyCharsT1
	}

	//fmt.Print(theme["top_left"], theme["top"], theme["top"], theme["top"], theme["top"], theme["top"], theme["top"], theme["top"], theme["top"], theme["top"], theme["top"], theme["top_mid"], theme["top"], theme["top"], theme["top"], theme["top"], theme["top"], theme["top"], theme["top"], theme["top"], theme["top"], theme["top_mid"], theme["top"], theme["top"], theme["top"], theme["top"], theme["top_right"], "\n")

	fmt.Print(theme["top_left"])
	for index, length := range maxChars {
		for i := 1; i <= length; i++ {
			fmt.Print(theme["top"])
		}
		if index != (len(maxChars) - 1) {
			fmt.Print(theme["top_mid"])
		}
	}
	fmt.Print(theme["top_right"], "\n")

	//Headers
	if t.headers != nil {

		fmt.Print(theme["left"])
		for index, header := range t.headers {
			fmt.Print(header)
			if index != (len(t.headers) - 1) {
				fmt.Print(theme["middle"])
			}
		}
		fmt.Print(theme["right"])
		fmt.Print("\n")
		fmt.Print(theme["left_mid"])
		for index, length := range maxChars {
			for i := 1; i <= length; i++ {
				fmt.Print(theme["mid"])
			}
			if index != (len(maxChars) - 1) {
				fmt.Print(theme["mid_mid"])
			}
		}
		fmt.Print(theme["right_mid"], "\n")
	}

	//Values
	for index, value := range t.values {
		fmt.Print(theme["left"])
		for index2, value2 := range value {
			fmt.Print(value2)
			if index2 != (len(value) - 1) {
				fmt.Print(theme["middle"])
			}
		}
		fmt.Print(theme["right"])
		fmt.Print("\n")

		if len(value) > 1 && index != (len(value)) {
			fmt.Print(theme["left_mid"])
			for index, length := range maxChars {
				for i := 1; i <= length; i++ {
					fmt.Print(theme["mid"])
				}
				if index != (len(maxChars) - 1) {
					fmt.Print(theme["mid_mid"])
				}
			}
			fmt.Print(theme["right_mid"], "\n")
		}
	}

	//fmt.Print(theme["left_mid"], theme["mid"], theme["mid"], theme["mid"], theme["mid"], theme["mid_mid"], theme["mid"], theme["mid"], theme["mid"], theme["mid"], theme["mid_mid"], theme["mid"], theme["mid"], theme["mid"], theme["mid"], theme["right_mid"], "\n")
	//fmt.Print(theme["left"], "    ", theme["middle"], "    ", theme["middle"], "    ", theme["right"], "\n")

	//fmt.Print(theme["bottom_left"], theme["bottom"], theme["bottom"], theme["bottom"], theme["bottom"], theme["bottom_mid"], theme["bottom"], theme["bottom"], theme["bottom"], theme["bottom"], theme["bottom_mid"], theme["bottom"], theme["bottom"], theme["bottom"], theme["bottom"], theme["bottom_right"], "\n")

	fmt.Print(theme["bottom_left"])
	for index, length := range maxChars {
		for i := 1; i <= length; i++ {
			fmt.Print(theme["bottom"])
		}
		if index != (len(maxChars) - 1) {
			fmt.Print(theme["bottom_mid"])
		}
	}
	fmt.Print(theme["bottom_right"], "\n")

}

var bodyCharsT1 = map[string]string{
	"top":          "─",
	"top_mid":      "┬",
	"top_left":     "┌",
	"top_right":    "┐",
	"bottom":       "─",
	"bottom_mid":   "┴",
	"bottom_left":  "└",
	"bottom_right": "┘",
	"left":         "│",
	"left_mid":     "├",
	"mid":          "─",
	"mid_mid":      "┼",
	"right":        "│",
	"right_mid":    "┤",
	"middle":       "│",
}
var bodyCharsT2 = map[string]string{
	"top":          "═",
	"top_mid":      "╤",
	"top_left":     "╔",
	"top_right":    "╗",
	"bottom":       "═",
	"bottom_mid":   "╧",
	"bottom_left":  "╚",
	"bottom_right": "╝",
	"left":         "║",
	"left_mid":     "╟",
	"mid":          "─",
	"mid_mid":      "┼",
	"right":        "║",
	"right_mid":    "╢",
	"middle":       "│",
}

func main() {

	myTable := Table{
		values:       [][]string{{"Sajad", "K", "Developer"}, {"VSCode", "Microsoft", "IDE"}},
		headers:      []string{"First Name", "Last Name", "Role"},
		lockedHeader: false,
		bodyTheme:    "t1",
		color:        "#FFFFFF",
		headerSplit:  false,
	}
	myTable.pushValue([]string{"Mr.X", "Anonymous", "Stranger"})
	myTable.pushValue([]string{"Mr.Y", "Anonymous", "Stranger"})
	myTable.print()

	clients := Table{
		values: [][]string{
			{"112.83.68.215", "be:ec:ad:ea:b4:d6", "USA"},
			{"168.31.13.241", "fc:ee:bd:21:eb:e2", "Canada"},
			{"137.114.50.162", "3a:55:06:c8:e3:4b", "England"},
			{"32.255.101.12", "93:31:80:fd:42:b7", "Germany"},
		},
		headers:      []string{"IP", "Mac Address", "Country"},
		lockedHeader: false,
		bodyTheme:    "t2",
		color:        "#FFFFFF",
		headerSplit:  false,
	}

	clients.print()

	fmt.Scanf("h")

	//myInts := []int{4, 2, 6, 1, 8, 7}
	//fmt.Println(myInts)
	//sort.Ints(myInts)
	//fmt.Println(myInts)

}
