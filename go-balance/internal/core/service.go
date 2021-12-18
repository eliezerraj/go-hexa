package core

import (
	"log"
	"context"
)

type service struct {
	repo 	BalanceRepositoryPort
	cliente BalanceClientPort
	ctx 	context.Context
}

func NewService(repo BalanceRepositoryPort, clie BalanceClientPort) *service {
	return &service{
		cliente: clie,
		ctx: context.Background(),
	}
}

func (p *service) AddBalance(balance Balance) error {
	log.Printf("AddBalance")
	
	log.Printf("--------------------------------------")
	log.Printf("- AddBalance Doing Business Rules !!!!")
	log.Printf("--------------------------------------")

	err := p.repo.AddBalance(balance)
	if err != nil {
		return err
	}
	return nil
}

func (p *service) GetBalance(account string) (Balance, error) {
	log.Printf("GetBalance")
	
	log.Printf("--------------------------------------")
	log.Printf("- GetBalance Doing Business Rules !!!!")
	log.Printf("--------------------------------------")

	res, err := p.repo.GetBalance(account)
	if err != nil {
		return Balance{}, err
	}
	return res, nil
}

func (p *service) ListBalance() ([]Balance, error) {
	log.Printf("ListBalance")
	
	log.Printf("--------------------------------------")
	log.Printf("- ListBalance Doing Business Rules !!!!")
	log.Printf("--------------------------------------")

	res, err := p.repo.ListBalance()
	if err != nil {
		return nil, err
	}

	for _, value := range res {
		rate, err := p.cliente.GetRate(value.Account)
		if err != nil {
			return nil, err
		}
		log.Printf("rate %v %v \n", value.Account , rate)
	}

	return res, nil
}
