package ies

import "rrc/utils"

// SIB12-IEs-r16-sl-L3U2N-RelayDiscovery-r17 ::= ENUMERATED
type Sib12IesR16SlL3u2nRelaydiscoveryR17 struct {
	Value utils.ENUMERATED
}

const (
	Sib12IesR16SlL3u2nRelaydiscoveryR17EnumeratedNothing = iota
	Sib12IesR16SlL3u2nRelaydiscoveryR17EnumeratedEnabled
)
