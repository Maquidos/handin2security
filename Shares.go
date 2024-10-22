package main

import (
	"math/rand"
)

type Shares struct {
	Share1, Share2, Share3 int
}

func generateShares(secret int, R int) Shares {
	share1 := rand.Intn(R)
	share2 := rand.Intn(R)
	share3 := (secret - share1 - share2 + R+1) % (R+1) // Ensures secret = share1 + share2 + share3
	return Shares{Share1: share1, Share2: share2, Share3: share3}
}

