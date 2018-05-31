package models

import (
	"time"
	"fmt"
	"labix.org/v2/mgo/bson"
)

type Account struct {
	AccountId string `json:"accountId" bson:"accountId"`
}

type Quote struct {
	Id string `json:"id" bson:"id"`
	SalesPersonUserId string `json:"salesPersonUserId" bson:"salesPersonUserId"`
	Opportunity string `json:"opportunity" bson:"opportunity"`
}

type AutomationEvent struct {
	EventId string `json:"eventId" bson:"eventId"`
}

type InventoryState struct { /* Inventory State Model */
	State string `json:"state" bson:"state"`
	UserId string `json:"userId" bson:"userId"`
	Comment string `json:"comment" bson:"comment"`
	AutomationEvent AutomationEvent `json:"automationEvent" bson:"automationEvent"`
	Device bson.ObjectId `bson:"device"`
	Account Account `json:"account" bson:"account"`
	Quote Quote `json:"quote" bson:"quote"`
	TimeStamp time.Time `json:"timestamp" bson:"timestamp"`
}

func (i *InventoryState) SetState(state string) error {
	states := []string{"AVAILABLE", "MAINTENANCE", "SUSPENDED", "DECOMMISSIONED", "DELETED"}
	for _, s := range states {
		if state == s {
			i.State = s
			return nil
		}
	}
	return ValidationError{fmt.Sprintf("inventory state should be one of %v", states), ErrorUnprocessable}
}

func NewInventoryState(state string, userId string, comment string, device bson.ObjectId,
	automationEvent string, accountId string, quoteId string,
	salesperson string, opportunity string) (error, InventoryState) {

	account := Account{AccountId: accountId}
	quote := Quote{Id: quoteId, SalesPersonUserId: salesperson, Opportunity: opportunity}
	automation := AutomationEvent{EventId: automationEvent}
	
	inventoryState := InventoryState{
		UserId: userId,
		Comment: comment,
		Device: device,
		Account: account,
		Quote: quote,
		AutomationEvent: automation}
	if err := inventoryState.SetState(state); err != nil {
		return err, inventoryState
	}
	
	return nil, inventoryState
}

func (i *InventoryState) Save() error {
	i.TimeStamp = time.Now()
	return nil
}
