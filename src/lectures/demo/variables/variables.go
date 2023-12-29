package main

import "fmt"

func main() {
	var myName = "Aleisha"
	fmt.Println("my name is", myName)

	var name string = "Bob"
	fmt.Println("name = ", name)

	userName := "admin"
	fmt.Println("username =", userName)

	var sum int
	fmt.Println("The sum is", sum)

	part1, other := 1, 5
	fmt.Println("part 1 is", part1, "other is", other)
	fmt.Printf("part 1 is %d other is %d \n", part1, other)

	part2, other := 2, 0
	fmt.Printf("part 2 is %d other is now %d \n", part2, other) // Printf allows for formatting symbols but does not add spacing or new line at end so new line must be included in print statement

	sum = part1 + part2
	fmt.Println("sum =", sum)

	var (
		lessonName = "Variables"
		lessonType = "Demo"
	)

	fmt.Print("lessonName = ", lessonName) // Print allows for formatting symbols but does not add spacing or new line at end
	fmt.Println("lessonType =", lessonType)

	word1, word2, _ := "hello", "world", "!"
	fmt.Println(word1, word2)
}
