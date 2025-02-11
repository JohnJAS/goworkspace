package main

import "fmt"

var table = [][]string{
	{"", "M", "MM", "MMM"},
	{"", "C", "CC", "CCC", "CD", "D", "DC", "DCC", "DCCC", "CM"},
	{"", "X", "XX", "XXX", "XL", "L", "LX", "LXX", "LXXX", "XC"},
	{"", "I", "II", "III", "IV", "V", "VI", "VII", "VIII", "IX"},
}

func intToRoman(num int) string {
	return table[0][num/1000] + table[1][num%1000/100] + table[2][num%100/10] + table[3][num%10]
}

func main() {
	fmt.Println(intToRoman(35))
	fmt.Println(intToRoman(3749))
	fmt.Println(intToRoman(58))
	fmt.Println(intToRoman(1994))
}
