package ies

import "rrc/utils"

// SL-PBPS-CPS-Config-r17-sl-Additional-PBPS-Occasion-r17 ::= ENUMERATED
type SlPbpsCpsConfigR17SlAdditionalPbpsOccasionR17 struct {
	Value utils.ENUMERATED
}

const (
	SlPbpsCpsConfigR17SlAdditionalPbpsOccasionR17EnumeratedNothing = iota
	SlPbpsCpsConfigR17SlAdditionalPbpsOccasionR17EnumeratedMonitored
)
