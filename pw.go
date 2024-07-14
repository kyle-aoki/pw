package main

import (
	"crypto/rand"
	"flag"
	"fmt"
	"math/big"
)

var (
	length = flag.Int("l", 32, "length of password")
)

const pool = "ABCDEFGHIJKLMNOPQRSTUVWXYZ" + "0123456789"

func main() {
	flag.Parse()
	var str []byte
	for i := 0; i < *length; i++ {
		ch := pool[randInt(len(pool))]
		str = append(str, ch)
	}
	pw := formatNetPassword(string(str))
	fmt.Print(pw)
}

func formatNetPassword(str string) string {
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
