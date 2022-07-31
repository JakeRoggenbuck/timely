package main

import (
	"fmt"
	"github.com/fatih/color"
	"os"
)

const KHRON_STATE_WORKING = "KHRON_STATE_WORKING"

type State int

const (
	Working State = iota + 1
	Off
)

func (s State) String() string {
	states := [...]string{"working", "off"}
	if s < Working || s > Off {
		return fmt.Sprintf("State(%d)", int(s))
	}
	return states[s-1]
}

func (s State) Char() string {
	states := [...]string{"▶", "▷"}
	if s < Working || s > Off {
		return fmt.Sprintf("State(%d)", int(s))
	}
	return states[s-1]
}

func (s State) Print() {
	var c *color.Color
	if s == Working {
		c = color.New(color.FgRed)
	} else {
		c = color.New(color.FgBlue)
	}

	c.Println(s.Char() + " " + s.String())
}

func main() {
	current := os.Getenv(KHRON_STATE_WORKING)
	var state State

	if current == "true" {
		state = Working
	} else { 
		state = Off
	}

	state.Print()
}
