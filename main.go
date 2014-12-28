package main

import (
	"iangreenleaf/noted/notes"
)

func main() {
  m := notes.NewServer()
	m.Run()
}
