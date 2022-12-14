package main

import (
	"fmt"
	"github.com/fatih/color"
	"io/ioutil"
	"os"
	"time"
	"log"
)

const STATE_FILE = "/.local/share/timely/state"
const TIMES_FILE = "/.local/share/timely/times"

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

func (s State) Println(no_color bool, no_word bool) {
	var c *color.Color
	if s == Working {
		c = color.New(color.FgRed)
	} else {
		c = color.New(color.FgBlue)
	}

	if no_color {
		c.DisableColor()
	} else {
		c.EnableColor()
	}

	if no_word {
		c.Println(s.Char())
	} else {
		c.Println(s.Char() + " " + s.String())
	}
}

func (s State) Print(no_color bool, no_word bool) {
	var c *color.Color
	if s == Working {
		c = color.New(color.FgRed)
	} else {
		c = color.New(color.FgBlue)
	}

	if no_color {
		c.DisableColor()
	} else {
		c.EnableColor()
	}

	if no_word  {
		c.Print(s.Char())
	} else {
		c.Print(s.Char() + " " + s.String())
	}
}

func (s State) Time() string {
	t := time.Now()
	return "Started " + s.String() + " at " + t.Format("01-02-2006 15:04:05 Monday") + "\n"
}

func (s State) Set() {
	home, err := os.UserHomeDir()
	if err != nil {
		fmt.Println(err)
	}

	if err := os.WriteFile(home+STATE_FILE, []byte(s.String()), 0666); err != nil {
		log.Fatal(err)
	}


	f, err := os.OpenFile(home+TIMES_FILE, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600)
	if err != nil {
    	log.Fatal(err)
	}

	defer f.Close()

	if _, err = f.WriteString(s.Time()); err != nil {
		log.Fatal(err)
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
		if state != Working {
			state = Working
			state.Set()
		}
	} else if args.off {
		if state != Off {
			state = Off
			state.Set()
		}
	} else if args.toggle {
		if state == Working {
			state = Off
			state.Set()
		} else {
			state = Working
			state.Set()
		}
	} else if args.check {
		if args.inline {
			state.Print(args.no_color, args.no_word)
		} else {
			state.Println(args.no_color, args.no_word)
		}
	} else {
		fmt.Println("Welcome to timely")
		fmt.Println("Use --help for more info")
	}
}
