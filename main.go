package main

import (
	"flag"
	"fmt"

	"github.com/costa365/pwdgen"
)

func parseFlags() pwdgen.Config {
	var cfg pwdgen.Config

	uppers := flag.Int("u", 4, "Number of upper case characters")
	lowers := flag.Int("l", 4, "Number of lower case characters")
	digits := flag.Int("d", 4, "Number of digit [0-9] characters")
	specials := flag.Int("s", 4, "Number of special characters")

	flag.Parse()

	cfg.UpperAlphas = *uppers
	cfg.LowerAlphas = *lowers
	cfg.Digits = *digits
	cfg.Specials = *specials

	return cfg
}

func main() {
	pwdgen.PwdConfig = parseFlags()

	fmt.Println(pwdgen.Password())
}
