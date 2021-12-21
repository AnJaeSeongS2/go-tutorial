package basics

import (
	cryptoRand "crypto/rand"
	"fmt"
	"log"
	"math/big"
	"math/rand"
	"time"
)

func TourPackage() {
	randN := time.Now().UnixNano()
	rand.Seed(randN)
	fmt.Println("My favorite first number is ", rand.Intn(10))

	// crypto/rand use CSPRNG. and on linux, it use /dev/urandom (heavy job)
	cryptoRandN, err := cryptoRand.Int(cryptoRand.Reader, big.NewInt(10))
	if err != nil {
		log.Panic(err)
	}
	fmt.Println("My favorite second number is ", cryptoRandN)
}
