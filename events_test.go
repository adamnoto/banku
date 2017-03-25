package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCreateCoupon(t *testing.T) {
	expectedResult := CreateEvent{
		CouponName: "Sushi Bonanza",
	}

	event := newCouponEvent("Sushi Bonanza")

	assert.Equal(t, expectedResult, event)
}

func TestInvalidateCoupon(t *testing.T) {
	id := "123"

	event := newInvalidateEvent(id)

	assert.Equal(t, id, *event.CouponId)
}

func TestAcceptCoupon(t *testing.T) {
	id := "123"
	transId := "T028MN"

	event := newAcceptEvent(id, transId)

	assert.Equal(t, id, *event.CouponId)
	assert.Equal(t, transId, event.TransID)
}
