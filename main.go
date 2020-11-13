package main

import (
  "github.com/StephanieSunshine/go-uber-fxtest/sample"
  "github.com/davecgh/go-spew/spew"
)

func main() {
  s := sample.New("Steph")
  spew.Dump(s)
}
