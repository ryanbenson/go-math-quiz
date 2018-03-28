package main

import (
  "testing"
  . "github.com/franela/goblin"
  . "github.com/onsi/gomega"
)

func TestCsv(t *testing.T) {
  g := Goblin(t)

  //special hook for gomega
  RegisterFailHandler(func(m string, _ ...int) { g.Fail(m) })

  g.Describe("CSV", func() {

    g.It("should parse the csv and get the data as an array", func() {
      filename := "problems.csv"
      c := new(CsvParser).init(filename).parse()

      var answers = [][]string{{"5+5", "10"},{"7+3", "10"},{"1+1", "2"},{"8+3", "11"},{"1+2", "3"},{"8+6", "14"},{"3+1", "4"},{"1+4", "5"},{"5+1", "6"},{"2+3", "5"},{"3+3", "6"},{"2+4", "6"},{"5+2", "7"}}
      Ω(c.Data).Should(Equal(answers))
    })

  })
}
