package main

import (
	"fmt"
	"log"
	"math/rand"
	"os"
	"strconv"
	"time"
)

type Pool string

const lowercasePool Pool = "abcdefghijklmnopqrstuvwxyz"
const uppercasePool Pool = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
const numberPool Pool = "1234567890"
const runePool Pool = lowercasePool + uppercasePool + numberPool

type Password string

func main() {
	rand.Seed(time.Now().UnixNano())

	pwLength := obtainPasswordLengthFromArgs()
	var pw Password

	for !pw.hasElementsFromAllPools() {
		pw = generatePassword(pwLength)
	}
	
	fmt.Println(pw)
}

func generatePassword(pwLength int) Password {
	var bytes []byte
	for i := 0; i < pwLength; i++ {
		i := rand.Intn(len(runePool))
		bytes = append(bytes, runePool[i])
	}
	return Password(bytes)
}

// true if pw has elements from every pool
func (pw Password) hasElementsFromAllPools() bool {
	return (pw.hasElementFrom(numberPool) &&
		pw.hasElementFrom(uppercasePool) &&
		pw.hasElementFrom(lowercasePool))
}

// returns true if pool contains given byte
func (pool Pool) containsByte(b byte) bool {
	for i := range pool {
		if pool[i] == b {
			return true
		}
	}
	return false
}

// true if contains at least one element from specified pool
func (pw Password) hasElementFrom(pool Pool) bool {
	for i := range pw {
		if pool.containsByte(pw[i]) {
			return true
		}
	}
	return false
}

// default password length is 20, unless first arg specifies an alternative length
func obtainPasswordLengthFromArgs() int {
	args := os.Args[1:]

	if len(args) == 0 {
		return 20
	} else {
		pwLen, err := strconv.ParseInt(args[0], 10, 32)
		if err != nil {
			log.Fatal("First argument must be an integer.")
		}
		return int(pwLen)
	}
}
