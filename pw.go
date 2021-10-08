package main

import (
	"fmt"
	"log"
	"math/rand"
	"os"
	"strconv"
	"time"
)

const lower_case_letters = "abcdefghijklmnopqrstuvwxyz"
const upper_case_letters = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
const numbers = "1234567890"
const symbols = "*&@"
const char_pool = lower_case_letters + upper_case_letters + numbers + symbols

func main() {
	rand.Seed(time.Now().UnixNano())
	pwLength := obtainLengthFromArgs()
	validatePwLength(pwLength)

	for i := 0; i < 3000; i++ {
		pw := generatePassword(pwLength)
		if containsAll(pw) {
			fmt.Print(pw)
			os.Exit(0)
		}
	}
	log.Fatal("Failed to generate password.")
}

func containsAll(pw string) bool {
	return (contains(symbols, pw) &&
		contains(numbers, pw) &&
		contains(upper_case_letters, pw) &&
		contains(lower_case_letters, pw))
}

func isMember(pool string, char rune) bool {
	for _, r := range pool {
		if r == char {
			return true
		}
	}
	return false
}

func contains(pool string, pw string) bool {
	for _, char := range pw {
		if isMember(pool, char) {
			return true
		}
	}
	return false
}

func generatePassword(pwLength int) string {
	pw := ""
	for i := 0; i < pwLength; i++ {
		randomInteger := rand.Intn(len(char_pool))
		pw += string(char_pool[randomInteger])
	}
	return pw
}

func obtainLengthFromArgs() int {
	args := os.Args[1:]
	
	if len(args) == 0 {
		return 20
	} else {
		pwl, err := strconv.ParseInt(args[0], 10, 32)
		if err != nil {
			log.Fatal("First argument must be an integer.")
		}

		return int(pwl)
	}
}

func validatePwLength(pwLength int) {
	if pwLength < 4 {
		log.Fatal("Password too short. Minumum 4 characters long.")
	}
}
