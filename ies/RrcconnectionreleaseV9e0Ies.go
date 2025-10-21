package ies

import "rrc/utils"

// RRCConnectionRelease-v9e0-IEs ::= SEQUENCE
type RrcconnectionreleaseV9e0Ies struct {
	RedirectedcarrierinfoV9e0       *RedirectedcarrierinfoV9e0
	IdlemodemobilitycontrolinfoV9e0 *IdlemodemobilitycontrolinfoV9e0
	Noncriticalextension            *interface{}
}
