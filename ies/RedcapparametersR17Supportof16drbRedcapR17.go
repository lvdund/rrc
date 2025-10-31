package ies

import "rrc/utils"

// RedCapParameters-r17-supportOf16DRB-RedCap-r17 ::= ENUMERATED
type RedcapparametersR17Supportof16drbRedcapR17 struct {
	Value utils.ENUMERATED
}

const (
	RedcapparametersR17Supportof16drbRedcapR17EnumeratedNothing = iota
	RedcapparametersR17Supportof16drbRedcapR17EnumeratedSupported
)
