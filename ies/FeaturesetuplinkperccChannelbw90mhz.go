package ies

import "rrc/utils"

// FeatureSetUplinkPerCC-channelBW-90mhz ::= ENUMERATED
type FeaturesetuplinkperccChannelbw90mhz struct {
	Value utils.ENUMERATED
}

const (
	FeaturesetuplinkperccChannelbw90mhzEnumeratedNothing = iota
	FeaturesetuplinkperccChannelbw90mhzEnumeratedSupported
)
