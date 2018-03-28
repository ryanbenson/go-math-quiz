package main

import (
  "fmt"
)

type Quiz struct {
  Content [][]string
}

func main() {
  fmt.Println("Let's do this")
}

func (q Quiz) init(content [][]string) Quiz {
  q.Content = content
  return q
}

func (q Quiz) check(i int, val string) bool {
  valid := false
  if q.Content[i][1] == val {
    valid = true
  }
  return valid
}
