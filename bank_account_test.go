package main_test

/* don't commit this file yet, let this file be used to test
   that when you push something and open PR, Semaphore will run
	 a test */

import (
	. "banku"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("BankAccount", func() {
	var (
		id   = "817212-2939283923-232323-2323"
		name = "John Smith"
	)

	Describe("ToAccount", func() {
		It("cannot convert wrongly formatted number", func() {
			data := map[string]string{
				"Id":      id,
				"Name":    name,
				"Balance": "ABC",
			}

			acc, err := ToAccount(data)

			Expect(acc).To(BeNil())
			Expect(err).NotTo(BeNil())
		})

		It("can return a BankAccount object given correct attributes", func() {
			data := map[string]string{
				"Id":      id,
				"Name":    name,
				"Balance": "10000",
			}

			acc, err := ToAccount(data)

			Expect(acc).NotTo(Equal(nil))
			Expect(acc.Id).To(Equal(id))
			Expect(acc.Name).To(Equal(name))
			Expect(acc.Balance).To(Equal(10000))
			Expect(err).To(BeNil())
		})
	})
})
