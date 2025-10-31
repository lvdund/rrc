package ies

import "rrc/utils"

// RRCSetupComplete-v1700-IEs ::= SEQUENCE
type RrcsetupcompleteV1700 struct {
	OnboardingrequestR17 *utils.BOOLEAN
	Noncriticalextension *RrcsetupcompleteV1700IesNoncriticalextension
}
