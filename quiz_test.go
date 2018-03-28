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

    g.It("should tell you how many answers you got right", func() {
      var questionAnswers = [][]string{{"5+5", "10"},{"7+3", "10"},{"1+1", "2"},{"8+3", "11"},{"1+2", "3"}}
      var answers = []string{"10","10","2","5","3"}
      validAnswers := 4
      q := new(Quiz).init(questionAnswers)
      correctAnswers := q.score(answers)
      Ω(correctAnswers).Should(Equal(validAnswers))
    })

    g.It("should give you a grade based on answers correct vs total questions", func() {
      var questionAnswers = [][]string{{"5+5", "10"},{"7+3", "10"},{"1+1", "2"},{"8+3", "11"},{"1+2", "3"}}
      validAnswers := 4
      q := new(Quiz).init(questionAnswers)
      grade := q.grade(validAnswers)
      Ω(grade).Should(Equal(80))
    })

  })
}
