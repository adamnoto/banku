package main

type Event struct {
	Name string             `json:"name"`
	Id   *string            `json:"id,omitempty"`
	Data *map[string]string `json:"data"`
}

func createCuponEvent(name string) Event {
	return Event{
		Name: "CreateCoupon",
		Data: &map[string]string{
			"name": name,
		},
	}
}

func invalidateCouponEvent(id string) Event {
	return Event{
		Name: "InvalidateCoupon",
		Id:   &id,
	}
}

func acceptCouponEvent(id string, transactionId string) Event {
	return Event{
		Name: "AcceptCoupon",
		Id:   &id,
		Data: &map[string]string{
			"transactionId": transactionId,
		},
	}
}
