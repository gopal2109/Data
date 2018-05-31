package schemas

import (
	"rpis/api/models"
)

type InventoryState struct { /* Inventory State Model */
	State string `json:"state"`
	Comment string `json:"comment"`
	AutomationEvent models.AutomationEvent `json:"automationEvent,omitempty"`
	Account models.Account `json:"account,omitempty"`
	Quote models.Quote `json:"quote,omitempty"`
}
