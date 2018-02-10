package main

import "fmt"

func main() {
	const (
		January   = iota + 1
		February
		March1
		April
		May
		June
		July
		August
		September
		October
		November1
		December
	)

	fmt.Println(January)
	months := [...]string{
		January:   "1 月",
		February:  "2 月",
		March1:    "3 月",
		April:     "4 月",
		May:       "5 月",
		June:      "6 月",
		July:      "7 月",
		August:    "8 月",
		September: "9 月",
		October:   "10 月",
		November1: "11 月",
		December:  "12 月",
	}
	fmt.Println(months)
	quarter := [...][]string{
		{months[January], months[February], months[March1]},
		{months[April], months[May], months[June]},
		{months[July], months[August], months[September]},
		{months[October], months[November1], months[December]},
	}
	for index := 0; index < len(quarter); index++ {
		fmt.Printf("第%d季度%s\n", index+1, quarter[index])
	}

	quarter1 := months[:4]
	fmt.Println(quarter1)
	quarter2 := months[4:7]
	fmt.Println(quarter2)
	quarter3 := months[7:10]
	fmt.Println(quarter3)

	months[July] = "Seventh Months"
	fmt.Println(quarter2)
	fmt.Println(months)

	a := [...]int{0, 1, 2, 3, 4, 5}
	reverse(a[:])
	fmt.Println(a) // "[5 4 3 2 1 0]"

}

// reverse reverses a slice of ints in place.
func reverse(s []int) {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		fmt.Println("i => ", i, "  j => ", j)
		s[i], s[j] = s[j], s[i]
	}
}
