package main

import (
	"crypto/rand"
	"fmt"
	"math/big"
	"os"
	"strconv"
)

const letters = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
const numbers = "0123456789"
const pool = letters + numbers

func main() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println(r)
		}
	}()
	var length = 32
	if len(os.Args) > 1 {
		i64, err := strconv.ParseInt(os.Args[1], 10, 64)
		if err != nil {
			panic(err)
		}
		length = int(i64)
	}
	var s []byte
	for i := 0; i < length; i++ {
		s = append(s, pool[randInt(len(pool))])
	}
	fmt.Print(string(s))
}

func randInt(mx int) int64 {
	n, _ := rand.Int(rand.Reader, big.NewInt(int64(mx)))
	return n.Int64()
}
