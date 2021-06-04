//source: https://www.linkedin.com/feed/update/urn:li:activity:6805873728812130304/
/*
Once a week we will challenge you to think outside of the box and solve some coding challenges.
Let's see if you got what it takes to become the next mozer!

We will post 12 challenges during the next weeks.

2/12:
Create a function that accepts an array of strings. Return the longest string.

Post your answer in the comments. We will post the solution on Monday!
Till then, good luck everyone 😄 🙌
*/

package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	fmt.Println("Type the strings you wish and type 'quit' when you wish to stop")
	var stringArray []string
	reader := bufio.NewReader(os.Stdin)
	cnt := 0
	biggestStringIndex := 0
	biggestStringCount := 0
	read := ""
	for true {
		fmt.Print("String ", cnt+1, ": ")
		read, _ = reader.ReadString('\n')
		if read == "quit\n" {
			break
		}
		stringArray = append(stringArray, read)
		if biggestStringCount < len(stringArray[cnt]) {
			biggestStringCount = len(stringArray[cnt])
			biggestStringIndex = cnt
		}
		cnt++
	}
	if len(stringArray) == 0 {
		fmt.Print("There are no allocated strings")
	} else {
		fmt.Println("Biggest string: ", stringArray[biggestStringIndex], " (index ", biggestStringIndex+1, ")")
	}
}
