package main

import ()

func main() {
	conf := config{
		commandRegistry: getCommands(),
	}
	startRepl(&conf)
}
