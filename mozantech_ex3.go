//source: https://www.linkedin.com/posts/mozantech_wevegotyoucovered-mozantech-becomeamozer-activity-6808377208567074816-jbKV
/*
New week = New challenge!
Spend some of your time thinking outside the box and solve our coding challenge.
Let's see if you got what it takes to become the next mozer!

We will post 12 challenges during the next weeks.

3/12:
Write a function that takes a string, and returns the character that is most commonly used in the string.

Post your answer in the comments. We will post the solution on Monday!
Till then, good luck everyone 😄 🙌
*/

package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	read := ""
	fmt.Print("String Input: ")
	read, _ = reader.ReadString('\n')
	read = strings.Replace(read, "\n", "", 1)
	charMap := make(map[string]int)
	for _, r := range read {
		c := string(r)
		_, exists := charMap[c]
		if !exists {
			charMap[c] = 0
		}
		charMap[c]++
		fmt.Println(c)
	}
	maxNumber := math.MinInt16
	maxKey := ""
	for iterator := range charMap {
		if maxNumber < charMap[iterator] {
			maxKey = iterator
			maxNumber = charMap[iterator]
		}
	}
	fmt.Printf("The most frequent character is '%s' with %d occurrences\n", maxKey, maxNumber)

}
