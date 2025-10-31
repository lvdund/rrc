package ies

import "rrc/utils"

// SRS-CarrierSwitching-srs-SwitchFromCarrier ::= ENUMERATED
type SrsCarrierswitchingSrsSwitchfromcarrier struct {
	Value utils.ENUMERATED
}

const (
	SrsCarrierswitchingSrsSwitchfromcarrierEnumeratedNothing = iota
	SrsCarrierswitchingSrsSwitchfromcarrierEnumeratedSul
	SrsCarrierswitchingSrsSwitchfromcarrierEnumeratedNul
)
