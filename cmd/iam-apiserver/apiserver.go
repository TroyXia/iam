package main

import (
	"github.com/TroyXia/iam/internal/apiserver"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UTC().UnixNano())

	apiserver.NewApp("iam-apiserver").Run()
}
