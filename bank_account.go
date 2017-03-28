package main

import (
	"errors"
	"strconv"
)

type BankAccount struct {
	Id      string
	Name    string
	Balance int
}

func FetchAccount(id string) (*BankAccount, error) {
	cmd := Redis.HGetAll(id)
	if err := cmd.Err(); err != nil {
		return nil, err
	}

	data := cmd.Val()
	if len(data) == 0 {
		return nil, nil
	} else {
		return ToAccount(data)
	}
}

func updateAccount(id string, data map[string]interface{}) (*BankAccount, error) {
	cmd := Redis.HMSet(id, data)

	if err := cmd.Err(); err != nil {
		return nil, err
	} else {
		return FetchAccount(id)
	}
}

func ToAccount(m map[string]string) (*BankAccount, error) {
	balance, err := strconv.Atoi(m["Balance"])
	if err != nil {
		return nil, err
	}

	if _, ok := m["Id"]; !ok {
		return nil, errors.New("Missing account ID")
	}

	return &BankAccount{
		Id:      m["Id"],
		Name:    m["Name"],
		Balance: balance,
	}, nil
}
