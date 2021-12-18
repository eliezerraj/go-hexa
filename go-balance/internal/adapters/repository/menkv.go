package repository

import (
	"log"
	"encoding/json"

	"github.com/go-hexa/go-balance/pkg"
	"github.com/go-hexa/go-balance/internal/core"
)

type memkv struct {
	kv map[string][]byte
}

func NewMemKV() *memkv {
	return &memkv{kv: map[string][]byte{}}
}

func (repo *memkv) ListBalance() ([]core.Balance, error) {
	log.Printf("ListBalance")

	log.Printf("####################################")
	log.Printf("- ListBalance DataBase Access !!!!")
	log.Printf("####################################")

	var res []core.Balance
	for _, value := range repo.kv {
		balance := core.Balance{}
		err := json.Unmarshal(value, &balance)
		if err != nil {
			log.Print(err)
		}
		res = append(res, balance)
	}
	return res ,nil
}

func (repo *memkv) GetBalance(account string) (core.Balance, error) {
	log.Printf("GetBalance")

	log.Printf("####################################")
	log.Printf("- GetBalance DataBase Access !!!!")
	log.Printf("####################################")

	if value, ok := repo.kv[account]; ok {
		balance := core.Balance{}
		err := json.Unmarshal(value, &balance)
		if err != nil {
			return core.Balance{}, err
		}
		return balance, nil
	}
	return core.Balance{}, pkg.NotFound
}

func (repo *memkv) AddBalance(balance core.Balance) error {
	log.Printf("repo-AddBalance")

	log.Printf("####################################")
	log.Printf("- GetBalance DataBase Access !!!!")
	log.Printf("####################################")

	bytes, err := json.Marshal(balance)
	if err != nil {
		return err
	}
	repo.kv[balance.Id] = bytes
	return nil
}