package main

import (
	"dkstar88/mathgen/cmd"
	"github.com/shiena/ansicolor"
	log "github.com/sirupsen/logrus"
	"os"
)

func init() {
	log.SetFormatter(&log.TextFormatter{ForceColors: true})
	log.SetOutput(ansicolor.NewAnsiColorWriter(os.Stdout))
	// Only log the warning severity or above.
	log.SetLevel(log.InfoLevel)
}

func main() {
	cmd.Execute()
}
