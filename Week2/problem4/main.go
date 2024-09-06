package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Person struct {
	Name      string
	Job       string
	BirthYear int
}

func main() {
	var persons []Person
	file, err := os.Open("person.txt")

	if err != nil {
		fmt.Println("Error opening file:", err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		fields := strings.Split(line, "|")

		birthYear, err := strconv.Atoi(fields[2])

		if err != nil {
			fmt.Println("Error converting birth year:", err)
			continue
		}

		person := Person{
			Name:      fields[0],
			Job:       fields[1],
			BirthYear: birthYear,
		}

		persons = append(persons, person)
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
	}

	fmt.Println(persons)
}
