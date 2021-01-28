package main

import (
	"flag"

	"github.com/borosr/flutter-screenshot/src/executor"
	log "github.com/sirupsen/logrus"
)

func main() {
	v := flag.Bool("verbose", false, "") // TODO
	flag.Parse()

	if *v {
		log.SetLevel(log.DebugLevel)
	}

	if err := executor.Run(); err != nil {
		log.Fatal(err)
	}
}
