package ies

import "rrc/utils"

// FeatureSetDownlinkPerCC-channelBW-90mhz ::= ENUMERATED
type FeaturesetdownlinkperccChannelbw90mhz struct {
	Value utils.ENUMERATED
}

const (
	FeaturesetdownlinkperccChannelbw90mhzEnumeratedNothing = iota
	FeaturesetdownlinkperccChannelbw90mhzEnumeratedSupported
)
