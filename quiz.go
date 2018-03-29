package main

import (
  "fmt"
  "time"
  "os"
)

type Quiz struct {
  Content [][]string
}

func main() {
  var answers []string
  csvData := new(CsvParser).init("problems.csv").parse()
  quiz := new(Quiz).init(csvData.Data)
  quiz.startTimer()
  for i := range quiz.Content {
    fmt.Println(quiz.Content[i][0])
    var input string
    fmt.Scanln(&input)
    answers = append(answers, input)
  }
  grade := quiz.grade(quiz.score(answers))
  fmt.Printf("Your score is: %v%%", grade)
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

func (q Quiz) score(answers []string) int {
  correctCount := 0
  for i, v := range answers {
    if q.check(i, v) == true {
      correctCount = correctCount + 1
    }
  }
  return correctCount
}

func (q Quiz) grade(correctAnswers int) int {
  total := len(q.Content)
  percent := (float64(correctAnswers) / float64(total)) * 100
  return int(percent)
}

func (q Quiz) startTimer() {
  timer := time.NewTimer(time.Second * 30)
  go func() {
    <- timer.C
    fmt.Println("Time is up! Try again next time")
    os.Exit(0)
  }()
}
