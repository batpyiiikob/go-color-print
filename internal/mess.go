package mess

import (
	"fmt"
	"time"
)

type TypeMessage string

const (
	E     TypeMessage = "\033[31m" // error -> text by red
	S     TypeMessage = "\033[32m" // success -> text by green
	W     TypeMessage = "\033[33m" // warning -> text by yellow
	I     TypeMessage = "\033[34m" // information -> text by blue
	P     TypeMessage = "\033[35m" // progress -> text by purple
	C     TypeMessage = "\033[36m" // comment -> text by cyan
	Reset TypeMessage = "\033[0m"  // default color
)

type TypeSeparator string

type TypePrefix int

const (
	P_MESS TypePrefix = iota // by TypeMessage
	P_TIME                   // print current time
)

type Options struct {
	Prefix    TypePrefix
	Separator TypeSeparator
	Type      TypeMessage
}

func Print(opt Options, data ...interface{}) {
	printPrefix(opt)

	var i int = 0
	for _, v := range data {
		fmt.Print(opt.Type, v, Reset)
		i++
		if i != len(data) {
			if opt.Separator != "" {
				fmt.Print(opt.Type, opt.Separator, Reset)
			} else {
				fmt.Print(" ")
			}
		}
	}
}

func Println(opt Options, data ...interface{}) {
	printPrefix(opt)
	for _, v := range data {
		fmt.Print(opt.Type, v, Reset, "\n")
	}
}

func Printf(opt Options, str string, data ...interface{}) {
	printPrefix(opt)
	fmt.Printf(string(opt.Type)+str+string(Reset), data...)
}

func printPrefix(opt Options) {
	if opt.Prefix == P_MESS {
		switch opt.Type {
		case E:
			fmt.Print(opt.Type, "::ERROR:: ", Reset)
		case S:
			fmt.Print(opt.Type, "::SUCCESS:: ", Reset)
		case W:
			fmt.Print(opt.Type, "::WARNING:: ", Reset)
		case I:
			fmt.Print(opt.Type, "::INFO:: ", Reset)
		case P:
			fmt.Print(opt.Type, "::PROGRESS:: ", Reset)
		case C:
			fmt.Print(opt.Type, "::COMMENT:: ", Reset)
		}
	}

	if opt.Prefix == P_TIME {
		fmt.Print(opt.Type, time.Now().Format("2006-01-02 15:04:05 "), Reset)
	}
}
