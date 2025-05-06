package main

import (
	"fmt"
	"math/rand"
	"os"
)

func main() {
	dir := "sampleDir-" + fmt.Sprint(rand.ExpFloat64())
	err := os.Mkdir(dir, 0777)
	os.Chmod(dir, 0777)
	if err != nil {
		panic(err)
		return
	}
}
