package main

import (
  "bufio"
  "encoding/csv"
  "os"
  "io"
)

type CsvParser struct {
  Data [][]string
  File *os.File
}

func (c CsvParser) init(filename string) CsvParser {
  f, _ := os.Open(filename)
  c.File = f
  return c
}

func (c CsvParser) parse() CsvParser {
  var data [][]string
  r := csv.NewReader(bufio.NewReader(c.File))
  for {
    record, err := r.Read()

    if err == io.EOF {
      break
    }
    data = append(data, record)
  }
  c.Data = data
  return c
}
