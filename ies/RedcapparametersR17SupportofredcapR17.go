package ies

import "rrc/utils"

// RedCapParameters-r17-supportOfRedCap-r17 ::= ENUMERATED
type RedcapparametersR17SupportofredcapR17 struct {
	Value utils.ENUMERATED
}

const (
	RedcapparametersR17SupportofredcapR17EnumeratedNothing = iota
	RedcapparametersR17SupportofredcapR17EnumeratedSupported
)
