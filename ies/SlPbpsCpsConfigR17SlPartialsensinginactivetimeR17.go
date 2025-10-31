package ies

import "rrc/utils"

// SL-PBPS-CPS-Config-r17-sl-PartialSensingInactiveTime-r17 ::= ENUMERATED
type SlPbpsCpsConfigR17SlPartialsensinginactivetimeR17 struct {
	Value utils.ENUMERATED
}

const (
	SlPbpsCpsConfigR17SlPartialsensinginactivetimeR17EnumeratedNothing = iota
	SlPbpsCpsConfigR17SlPartialsensinginactivetimeR17EnumeratedEnabled
	SlPbpsCpsConfigR17SlPartialsensinginactivetimeR17EnumeratedDisabled
)
