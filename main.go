package main

import (
	"fmt"
	"github.com/fatih/color"
	"io/ioutil"
	"os"
)

const STATE_FILE = "/.local/share/khron-krato/state"
const TIMES_FILE = "/.local/share/khron-krato/times"

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

func (s State) Set() {
	home, err := os.UserHomeDir()
	if err != nil {
		fmt.Println(err)
	}

	if err := os.WriteFile(home+STATE_FILE, []byte(s.String()), 0666); err != nil {
		fmt.Println(err)
	}
}

func Get() State {
	home, err := os.UserHomeDir()
	if err != nil {
		fmt.Println(err)
	}

	content, err := ioutil.ReadFile(home + STATE_FILE)
	if err != nil {
		fmt.Println(err)
	}

	if string(content) == "working" {
		return Working
	} else {
		return Off
	}
}

func main() {
	args := parse_args()
	var state State

	state = Get()

	if args.work {
		state = Working
		state.Set()
	} else if args.off {
		state = Off
		state.Set()
	} else if args.toggle {
		if state == Working {
			state = Off
			state.Set()
		} else {
			state = Working
			state.Set()
		}
	} else if args.check {
		state.Print()
	} else {
		fmt.Println("Welcome to khron-krato")
		fmt.Println("Use --help for more info")
	}

}
