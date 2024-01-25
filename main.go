package main

import "main/internal/mess"

func main() {
  
	opt1 := mess.Options{
		Type:   mess.I,
		Prefix: mess.P_MESS,
    Separator: mess.TypeSeparator(" : "),
	}

	opt2 := mess.Options{
		Type:   mess.E,
		Prefix: mess.P_MESS,
	}

	opt3 := mess.Options{
		Type:   mess.S,
		Prefix: mess.P_TIME,
	}

	data := make([]string, 3)

	data[0] = "Hello"
	data[1] = "World"

	mp := make(map[string]int)

	mp["Hello 1"] = 1
	mp["World 2"] = 2
	mp["World 3"] = 3

	mess.Print(opt1, data, mp, "123213", 234, 3, true, "\n")
	mess.Println(opt2, data, mp, "123213", 234, 3, true)
	mess.Printf(opt3, "UPS: %#v,\n%#v", data, mp)
}
