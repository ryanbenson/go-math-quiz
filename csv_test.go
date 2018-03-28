package main

import (
  "testing"
  . "github.com/franela/goblin"
  . "github.com/onsi/gomega"
)

func TestConfig(t *testing.T) {
  g := Goblin(t)

  //special hook for gomega
  RegisterFailHandler(func(m string, _ ...int) { g.Fail(m) })

  g.Describe("CSV", func() {

    g.It("should parse the csv and get the data as an array", func() {
      // c := new(csv).init().parse()

      Ω("test").Should(Equal("list of records"))
    })

  })
}
