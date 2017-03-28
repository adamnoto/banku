package main

import (
	"errors"
)

func (e CreateEvent) Process() (*BankAccount, error) {
	return updateAccount(e.AccId, map[string]interface{}{
		"Id":      e.AccId,
		"Name":    e.AccName,
		"Balance": "0",
	})
}

func (e DepositEvent) Process() (*BankAccount, error) {
	if acc, err := FetchAccount(e.AccId); err != nil {
		return nil, err
	} else {
		newBalance := acc.Balance + e.Amount
		return updateAccount(e.AccId, map[string]interface{}{
			"Balance": newBalance,
		})
	}
}

func (e WithdrawEvent) Process() (*BankAccount, error) {
	if acc, err := FetchAccount(e.AccId); err != nil {
		return nil, err
	} else {
		if acc.Balance >= e.Amount {
			newBalance := acc.Balance - e.Amount
			return updateAccount(e.AccId, map[string]interface{}{
				"Balance": newBalance,
			})
		} else {
			return nil, errors.New("Insufficient amount")
		}
	}
}

func (e TransferEvent) Process() (*BankAccount, error) {
	if acc, err := FetchAccount(e.AccId); err != nil {
		return nil, err
	} else {
		if destAcc, err := FetchAccount(e.TargetId); err != nil {
			return nil, err
		} else {
			if acc.Balance >= e.Amount {
				acc.Balance -= e.Amount
				destAcc.Balance += e.Amount
				if _, err := updateAccount(destAcc.Id, map[string]interface{}{
					"Balance": destAcc.Balance,
				}); err != nil {
					return nil, err
				} else {
					return updateAccount(acc.Id, map[string]interface{}{
						"Balance": acc.Balance,
					})
				}
			} else {
				return nil, errors.New("Insufficient amount")
			}
		}
	}
}
