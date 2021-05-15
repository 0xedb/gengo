package main

import (
	"fmt"
	"os"
	"os/user"

	"github.com/0xedb/gengo/repl"
)

const BANNER = `
▄████ ▓█████  ███▄    █   ▄████  ▒█████  
██▒ ▀█▒▓█   ▀  ██ ▀█   █  ██▒ ▀█▒▒██▒  ██▒
▒██░▄▄▄░▒███   ▓██  ▀█ ██▒▒██░▄▄▄░▒██░  ██▒
░▓█  ██▓▒▓█  ▄ ▓██▒  ▐▌██▒░▓█  ██▓▒██   ██░
░▒▓███▀▒░▒████▒▒██░   ▓██░░▒▓███▀▒░ ████▓▒░
░▒   ▒ ░░ ▒░ ░░ ▒░   ▒ ▒  ░▒   ▒ ░ ▒░▒░▒░ 
 ░   ░  ░ ░  ░░ ░░   ░ ▒░  ░   ░   ░ ▒ ▒░ 
░ ░   ░    ░      ░   ░ ░ ░ ░   ░ ░ ░ ░ ▒  
	 ░    ░  ░         ░       ░     ░ ░  
										  
`

func main() {
	user, err := user.Current()
	if err != nil {
		panic(err)
	}
	fmt.Printf("Hello %s! Welcome to Gengo programming language!\n",
		user.Username)
	fmt.Printf("%s\n", BANNER)
	repl.Start(os.Stdin, os.Stdout)
}
