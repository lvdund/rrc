package ies

import "rrc/utils"

// CSI-RS-ConfigBeamformed-r13-channelMeasRestriction-r13 ::= ENUMERATED
type CsiRsConfigbeamformedR13ChannelmeasrestrictionR13 struct {
	Value utils.ENUMERATED
}

const (
	CsiRsConfigbeamformedR13ChannelmeasrestrictionR13EnumeratedNothing = iota
	CsiRsConfigbeamformedR13ChannelmeasrestrictionR13EnumeratedOn
)
