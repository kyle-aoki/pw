package main

import (
	"crypto/rand"
	"fmt"
	"math/big"
	"os"
	"strconv"
)

const pool = "ABCDEFGHIJKLMNOPQRSTUVWXYZ" + "0123456789"

func main() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println(r)
		}
	}()
	length := 30
	if len(os.Args) > 1 {
		i64, err := strconv.ParseInt(os.Args[1], 10, 64)
		check(err)
		length = int(i64)
	}
	var str []byte
	for i := 0; i < length; i++ {
		ch := pool[randInt(len(pool))]
		str = append(str, ch)
	}
	pw := formatPassword(string(str))
	fmt.Print(pw)
}

func formatPassword(str string) string {
	return fmt.Sprintf("%sx!", str)
}

func randInt(mx int) int64 {
	n, err := rand.Int(rand.Reader, big.NewInt(int64(mx)))
	check(err)
	return n.Int64()
}

func check(err error) {
	if err != nil {
		panic(err)
	}
}
