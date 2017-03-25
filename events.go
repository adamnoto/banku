package main

type Event struct {
	name string
	id   *string `omitempty`
	data *map[string]string
}

func createCuponEvent(name string) Event {
	return Event{
		name: "CreateCoupon",
		data: &map[string]string{
			name: name,
		},
	}
}

func invalidateCouponEvent(id string) Event {
	return Event{
		name: "InvalidateCoupon",
		id:   id,
	}
}

func acceptCouponEvent(id string, transactionId string) Event {
	return Event{
		name: "AcceptCoupon",
		id:   id,
		data: &map[string]string{
			transactionId: transactionId,
		},
	}
}
