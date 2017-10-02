package main

import (
	"math/rand"
	"time"
)

type Genome struct {
	score    int64
	sequence string
}

func main() {
	rand.Seed(time.Now().Unix())

}
