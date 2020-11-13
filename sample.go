package sample

import (
  "fmt"
)

type Sample struct {
  name string
}

func (s *Sample) Print() {
  fmt.Println(s.name)
}

func New(n string) *Sample {
  return Sample{name: n}
}


