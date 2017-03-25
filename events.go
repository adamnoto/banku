package main

/* each event expressed as its own struct */

type Event struct {
	CouponId *string `json:"id,omitempty"`
}

type CreateEvent struct {
	Event
	CouponName string `json:"couponName"`
}

type InvalidateEvent struct {
	Event
}

type AcceptEvent struct {
	Event
	TransID string `json:"transId"`
}

/* helper to create events */

func newCouponEvent(name string) CreateEvent {
	return CreateEvent{
		CouponName: name,
	}
}

func newInvalidateEvent(id string) InvalidateEvent {
	return InvalidateEvent{
		Event{CouponId: &id},
	}
}

func newAcceptEvent(id string, transactionId string) AcceptEvent {
	event := new(AcceptEvent)
	event.CouponId = &id
	event.TransID = transactionId
	return *event
}
