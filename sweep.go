package main

import (
	"log"
)

type Sweep struct{}

func (s *Sweep) Help() string {
	return "sweep"
}

func (s *Sweep) Run(args []string) int {
	log.Println("sweep!")
	return 0
}

func (s *Sweep) Synopsis() string {
	return "Delete unread items"
}
