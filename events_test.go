package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCreateCoupon(t *testing.T) {
	expectedResult := Event{
		Name: "CreateCoupon",
		Data: &map[string]string{
			"name": "Sushi Bonanza",
		},
	}

	eventPayload := createCuponEvent("Sushi Bonanza")

	assert.Equal(t, expectedResult, eventPayload)
}

func TestInvalidateCoupon(t *testing.T) {
	id := "123"
	expectedResult := Event{
		Name: "InvalidateCoupon",
		Id:   &id,
	}

	eventPayload := invalidateCouponEvent("123")

	assert.Equal(t, expectedResult, eventPayload)
}

func TestAcceptCoupon(t *testing.T) {
	id := "123"
	expectedResult := Event{
		Name: "AcceptCoupon",
		Id:   &id,
		Data: &map[string]string{
			"transactionId": "T1234",
		},
	}

	eventPayload := acceptCouponEvent(id, "T1234")

	assert.Equal(t, expectedResult, eventPayload)
}
