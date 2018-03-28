package main

import (
  "testing"
  . "github.com/franela/goblin"
  . "github.com/onsi/gomega"
)

func TestQuiz(t *testing.T) {
  g := Goblin(t)

  //special hook for gomega
  RegisterFailHandler(func(m string, _ ...int) { g.Fail(m) })

  g.Describe("Quiz", func() {

    g.It("should attach the list of questions and answers", func() {
      var questionAnswers = [][]string{{"5+5", "10"},{"7+3", "10"},{"1+1", "2"},{"8+3", "11"},{"1+2", "3"},{"8+6", "14"},{"3+1", "4"},{"1+4", "5"},{"5+1", "6"},{"2+3", "5"},{"3+3", "6"},{"2+4", "6"},{"5+2", "7"}}
      q := new(Quiz).init(questionAnswers)

      Ω(q.Content).Should(Equal(questionAnswers))
    })

    g.It("should tell you if the answer is right or wrong based on question index", func() {
      var questionAnswers = [][]string{{"5+5", "10"},{"7+3", "10"}}
      q := new(Quiz).init(questionAnswers)
      valid := q.check(0, "10")
      invalid := q.check(1, "42")

      Ω(valid).Should(BeTrue())
      Ω(invalid).Should(BeFalse())
    })

  })
}
