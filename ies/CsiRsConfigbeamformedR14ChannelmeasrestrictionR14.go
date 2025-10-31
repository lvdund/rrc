package ies

import "rrc/utils"

// CSI-RS-ConfigBeamformed-r14-channelMeasRestriction-r14 ::= ENUMERATED
type CsiRsConfigbeamformedR14ChannelmeasrestrictionR14 struct {
	Value utils.ENUMERATED
}

const (
	CsiRsConfigbeamformedR14ChannelmeasrestrictionR14EnumeratedNothing = iota
	CsiRsConfigbeamformedR14ChannelmeasrestrictionR14EnumeratedOn
)
