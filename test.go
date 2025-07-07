package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strings"
)

func main() {
	r := bufio.NewReader(os.Stdin)
	a := ""

	a, _ = r.ReadString('\n')
	a = strings.TrimSpace(a)

	b := strings.Fields(a)
	sort.Strings(b)

	fmt.Println(b)
}
