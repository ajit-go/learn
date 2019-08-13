package main

import (
	"bufio"
	"fmt"
	"log"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

func getPositiveRandomNumber(max int) int {
	seconds := time.Now().Unix()
	rand.Seed(seconds)
	return rand.Intn(max) + 1
}

func main() {
	rnd := getPositiveRandomNumber(100)
	//fmt.Printf("Random number: %d", rnd)
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("----Max retries 10----")
	i := 0
	for ; i < 10; i++ {
		fmt.Println("\nGuess a number between 1 and 100...!: ")
		response, err := reader.ReadString('\n')
		response = strings.TrimSpace(response)
		//fmt.Printf("-----response: %s\n", response)
		if err != nil {
			log.Fatal(err)
		}

		intResponse, err := strconv.Atoi(response)
		//fmt.Printf("intResponse: %d\n", intResponse)

		if err == nil {
			if intResponse > rnd {
				fmt.Println("Oops, Your guess was HIGH.")
			} else if intResponse < rnd {
				fmt.Println("Oops, Your guess was LOW.")
			} else {
				fmt.Println("Good job! you guess it.")
			}
		} else {
			fmt.Printf("Oops, you entered an invalid number." + err.Error())
		}
	}
	if i > 9 {
		fmt.Println("Better luck next time..!")
	}
}
