package main

import (
	"fmt"
	"strings"
)

type martian struct{}

func (m martian) talk() string {
	return "nack nack"
}

type laser int

func (l laser) talk() string {
	return strings.Repeat("pew ", int(l))
}

type talker interface {
	talk() string
}

func shout(t talker) {
	louder := strings.ToUpper(t.talk())
	fmt.Println(louder)
}

func main() {
	shout(martian{}) // NACK NACK
	shout(laser(2))  // PEW PEW

	type crater struct{}
	// shout(crater{})
	// crater does not implement talker (missing talk method)
	// craterはtalkerを実装しません(talkメソッドが欠けています)
}
