package ies

import "rrc/utils"

// SIB12-IEs-r16-sl-NonRelayDiscovery-r17 ::= ENUMERATED
type Sib12IesR16SlNonrelaydiscoveryR17 struct {
	Value utils.ENUMERATED
}

const (
	Sib12IesR16SlNonrelaydiscoveryR17EnumeratedNothing = iota
	Sib12IesR16SlNonrelaydiscoveryR17EnumeratedEnabled
)
