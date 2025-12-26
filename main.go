package main

func main() {
	cfg := &config{
		Next:     nil,
		Previous: nil,
	}

	startRepl(cfg)
}
