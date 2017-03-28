package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

/* let the program behave as a producer when executed */
func mainProducer() {
	var err error
	reader := bufio.NewReader(os.Stdin)
	kafka := newKafkaSyncProducer()

	for {
		fmt.Print("-> ")
		text, _ := reader.ReadString('\n')
		text = strings.Replace(text, "\n", "", -1)
		args := strings.Split(text, "###")
		cmd := args[0]

		switch cmd {
		case "create":
			if len(args) == 2 {
				accName := args[1]
				event := NewCreateAccountEvent(accName)
				sendMsg(kafka, event)
			} else {
				fmt.Println("Only specify create###Account Name")
			}
		case "deposit":
			if len(args) == 3 {
				accId := args[1]
				if amount, err := strconv.Atoi(args[2]); err == nil {
					event := NewDepositEvent(accId, amount)
					sendMsg(kafka, event)
				}
			} else {
				fmt.Println("Only specify deposit###Account ID###amount")
			}
		case "withdraw":
			if len(args) == 3 {
				accId := args[1]
				if amount, err := strconv.Atoi(args[2]); err == nil {
					event := NewWithdrawEvent(accId, amount)
					sendMsg(kafka, event)
				}
			} else {
				fmt.Println("Only specify withdraw###Account ID###amount")
			}
		case "transfer":
			if len(args) == 4 {
				sourceId := args[1]
				targetId := args[2]
				if amount, err := strconv.Atoi(args[3]); err == nil {
					event := NewTransferEvent(sourceId, targetId, amount)
					sendMsg(kafka, event)
				}
			} else {
				fmt.Println("Only specify transfer###Source ID###Target ID####amount")
			}
		default:
			fmt.Printf("Unknown command %s, only: create, deposit, withdraw, transfer\n", cmd)
		}

		if err != nil {
			fmt.Printf("Error: %s\n", err)
			err = nil
		}
	}
}
