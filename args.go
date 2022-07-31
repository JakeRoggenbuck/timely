package main

import (
	"flag"
)

type Args struct {
	work   bool
	off    bool
	toggle bool
	check  bool
	inline bool
}

func parse_args() Args {
	work := flag.Bool("work", false, "Set to work")
	off := flag.Bool("off", false, "Set to off")
	toggle := flag.Bool("toggle", false, "Set to toggle")
	check := flag.Bool("check", false, "Set to check")
	inline := flag.Bool("inline", false, "No newline after print")

	flag.Parse()

	return Args{*work, *off, *toggle, *check, *inline}
}
