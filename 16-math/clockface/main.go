package main

import (
	"github.com/schmiddim/learn-go-with-tests/16-math"
	"os"
	"time"
)

func main() {
	t := time.Now()
	clockface.SVGWriter(os.Stdout, t)
}
