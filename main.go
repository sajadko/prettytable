package main

import (
	"fmt"
	"github.com/gookit/color"
	"sort"
)

//Table structure
type Table struct {
	values       [][]string
	headers      []string
	lockedHeader bool
	bodyTheme    string
	bodyColor    string
	headerSplit  bool
	paddingRight int
	paddingLeft  int
}

//PushValue pushes the value to the table
func (t *Table) PushValue(values []string) {
	t.values = append(t.values, values)
}

//Print prints the table
func (t *Table) Print() {

	//Set Theme
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

	//Change Body Color if exists
	if t.bodyColor != "" {
		rgbcolor := color.HEX(t.bodyColor, false)
		for key, value := range theme {
			text := rgbcolor.Sprint(value)
			theme[key] = text
		}
	}

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
			diff := len(color.ClearCode(value2)) - maxChars[index2]

			var spaces = ""
			var paddingRight = ""
			var paddingLeft = ""
			if -(diff) != 0 {
				for i := 1; i <= -(diff); i++ {
					spaces += " "
				}
			}
			if t.paddingRight != 0 {
				for i := 1; i <= t.paddingRight; i++ {
					paddingRight += " "
				}
			}
			if t.paddingLeft != 0 {
				for i := 1; i <= t.paddingLeft; i++ {
					paddingLeft += " "
				}
			}
			t.values[index][index2] = paddingLeft + value2 + paddingRight + spaces
		}
	}

	//Add Space To Headers
	for index, value := range t.headers {
		diff := len(color.ClearCode(value)) - maxChars[index]
		var spaces = ""
		var paddingRight = ""
		var paddingLeft = ""
		if -(diff) != 0 {
			for i := 1; i <= -(diff); i++ {
				spaces += " "
			}
		}
		if t.paddingRight != 0 {
			for i := 1; i <= t.paddingRight; i++ {
				paddingRight += " "
			}
		}
		if t.paddingLeft != 0 {
			for i := 1; i <= t.paddingLeft; i++ {
				paddingLeft += " "
			}
		}
		t.headers[index] = paddingLeft + value + paddingRight + spaces
	}

	//fmt.Println(t.values)

	//Add paddings to maxChars
	for index, max := range maxChars {
		maxChars[index] = max + t.paddingRight + t.paddingLeft
	}

	//TOP
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

	//Bottom
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

//Bolder Bolds the text
func Bolder(text string) string {
	boldPattern := "\u001B[1m%s\u001B[0m"
	return fmt.Sprintf(boldPattern, text)
}

func main() {

	myTable := Table{
		values:       [][]string{{"Sajad", "K", "Developer"}, {"VSCode", "Microsoft", "IDE"}},
		headers:      []string{"First Name", "Last Name", "Role"},
		lockedHeader: false,
		bodyTheme:    "t1",
		bodyColor:    "#FFFFFF",
		headerSplit:  false,
	}
	myTable.PushValue([]string{"Mr.X", "Anonymous", "Stranger"})
	myTable.PushValue([]string{"Mr.Y", "Anonymous", "Stranger"})
	myTable.Print()

	//aqua := color.RGBColor{0, 255, 255, 0}
	clients := Table{
		values: [][]string{
			{"112.83.68.215", "be:ec:ad:ea:b4:d6", "USA", color.HEX("#76FF03").Sprint("Online"), "Test"},
			{"168.31.13.241", "fc:ee:bd:21:eb:e2", "Canada", color.HEX("#76FF03").Sprint("Online"), "Test"},
			{"137.114.50.162", "3a:55:06:c8:e3:4b", "England", color.HEX("#f44336").Sprint("Offline"), "Test"},
			{"32.255.101.12", "93:31:80:fd:42:b7", "Germany", color.HEX("#76FF03").Sprint("Online"), "Test"},
		},
		headers:      []string{Bolder("IP"), Bolder("Mac Address"), Bolder("Country"), Bolder("Status"), Bolder("Test")},
		lockedHeader: false,
		bodyTheme:    "t2",
		bodyColor:    "",
		headerSplit:  false,
		paddingLeft:  2,
		paddingRight: 2,
	}

	clients.Print()

	fmt.Scanf("h")
	//const PrintColor = "\033[38;5;%dm%s\033[39;49m\n"
	//for i := 0; i <= 256; i++ {
	//	fmt.Printf(PrintColor, i, "Sajad")
	//}

}
