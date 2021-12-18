package core

import (
	"log"

)

type client struct {
}

func NewClient() *client {
	return &client{	}
}

func (p *client) GetRate(account string) (int32, error) {
	log.Printf("GetRate")
	
	log.Printf("--------------------------------------")
	log.Printf("- GetRate Calling another Service !!!!")
	log.Printf("--------------------------------------")

	return 1, nil
}