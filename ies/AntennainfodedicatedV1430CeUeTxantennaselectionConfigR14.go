package ies

import "rrc/utils"

// AntennaInfoDedicated-v1430-ce-UE-TxAntennaSelection-config-r14 ::= ENUMERATED
type AntennainfodedicatedV1430CeUeTxantennaselectionConfigR14 struct {
	Value utils.ENUMERATED
}

const (
	AntennainfodedicatedV1430CeUeTxantennaselectionConfigR14EnumeratedNothing = iota
	AntennainfodedicatedV1430CeUeTxantennaselectionConfigR14EnumeratedOn
)
