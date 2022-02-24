package main

import (
	"flag"

	"github.com/tinygoprogs/pwgen"
)

var exclude = flag.String("exclude", "", "exclude all these character from generation")
var include = flag.String("include", "", "include all these character in generation; Note: too many includes will produce low entropy password!")
var length = flag.Int("len", 32, "length of generated password")

func main() {
	flag.Parse()
	c := &pwgen.GenPasswordConfig{
		MustExclude: *exclude,
		MustInclude: *include,
		Len:         *length,
	}
	println(pwgen.GenPassword(c))
}
