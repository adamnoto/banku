package main_test

import (
	. "banku"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Event processor", func() {
	var (
		name = "John Smith"
	)

	Describe("NewCreateAccountEvent", func() {
		It("can create an account", func() {
			acc, err := NewCreateAccountEvent(name).Process()

			Expect(err).To(BeNil())
			Expect(acc.Name).To(Equal(name))
			Expect(acc.Balance).To(Equal(0))
		})
	})

	Describe("NewDepositEvent", func() {
		It("can deposit to an account", func() {
			acc, _ := NewCreateAccountEvent(name).Process()

			acc, _ = NewDepositEvent(acc.Id, 20).Process()
			acc, err := NewDepositEvent(acc.Id, 25).Process()

			Expect(err).To(BeNil())
			Expect(acc.Balance).To(Equal(45))
		})
	})

	Describe("NewWithdrawEvent", func() {
		var acc *BankAccount

		BeforeEach(func() {
			acc, _ = NewCreateAccountEvent(name).Process()
			acc, _ = NewDepositEvent(acc.Id, 20).Process()
		})

		It("can withdraw money from an account with sufficient balance", func() {
			acc, err := NewWithdrawEvent(acc.Id, 30).Process()

			Expect(err).NotTo(BeNil())
			Expect(acc).To(BeNil())
		})

		It("can withdraw money from an account with sufficient balance", func() {
			acc, _ = NewWithdrawEvent(acc.Id, 5).Process()
			acc, _ = NewWithdrawEvent(acc.Id, 2).Process()

			Expect(acc.Balance).To(Equal(13))
		})
	})

	Describe("NewTransferEvent", func() {
		var accAdam, accNatan *BankAccount

		BeforeEach(func() {
			accAdam, _ = NewCreateAccountEvent("Adam Smith").Process()
			accNatan, _ = NewCreateAccountEvent("Natan Newstaff").Process()
			accAdam, _ = NewDepositEvent(accAdam.Id, 50).Process()
			accNatan, _ = NewDepositEvent(accNatan.Id, 100).Process()
		})

		It("can transfer when amount is sufficient", func() {
			NewTransferEvent(accNatan.Id, accAdam.Id, 30).Process()

			accAdam, _ = FetchAccount(accAdam.Id)
			accNatan, _ = FetchAccount(accNatan.Id)
			Expect(accAdam.Balance).To(Equal(80))
			Expect(accNatan.Balance).To(Equal(70))
		})

		It("cannot transfer when insufficient amount", func() {
			_, err := NewTransferEvent(accNatan.Id, accAdam.Id, 150).Process()

			accAdam, _ = FetchAccount(accAdam.Id)
			accNatan, _ = FetchAccount(accNatan.Id)
			Expect(accAdam.Balance).To(Equal(50))
			Expect(accNatan.Balance).To(Equal(100))
			Expect(err).NotTo(BeNil())
		})
	})
})
