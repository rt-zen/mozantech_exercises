//source https://www.linkedin.com/feed/update/urn:li:activity:6802933766923931648/
/*
Once a week we will challenge you to think outside of the box and solve some coding challenges.
Let's see if you got what it takes to become the next mozer!

We will post 12 challenges during the next weeks.

1/12:
Create a function that takes in two strings as two parameters and returns a boolean that indicates whether or not the first string is an anagram of the second string.

Post your answer in the comments. We will post the solution on Friday!
Till then, good luck everyone 😄 🙌
*/

package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strings"
)

//var m1 = regexp.MustCompile()

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("1st expression: ")
	strA, _ := reader.ReadString('\n')
	fmt.Print("2nd expression: ")
	strB, _ := reader.ReadString('\n')
	if strA == strB {
		fmt.Println("The expressions are equal")
	} else {
		replacer := strings.NewReplacer(
			":", "",
			".", "",
			",", "",
			" ", "",
			";", "",
			"!", "",
			"?", "",
			"-", "",
			"\"", "",
			"'", "")
		strA = replacer.Replace(strA)
		strB = replacer.Replace(strB)
		SortString(strA)
		SortString(strB)
		if strA == strB {
			fmt.Print("Anagram detected.")
		} else {
			fmt.Print("Anagram not detected.")
		}
	}
}

func SortString(str string) string {
	strOutA := strings.Split(strings.ToLower(str), "")
	sort.Strings(strOutA)
	strOut := strings.Join(strOutA, "")
	result := strings.ReplaceAll(strOut, " ", "")
	return result
}
