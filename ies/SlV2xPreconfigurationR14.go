package ies

import "rrc/utils"

// SL-V2X-Preconfiguration-r14 ::= SEQUENCE
// Extensible
type SlV2xPreconfigurationR14 struct {
	V2xPreconfigfreqlistR14  SlV2xPreconfigfreqlistR14
	AnchorcarrierfreqlistR14 *SlAnchorcarrierfreqlistV2xR14
	CbrPreconfiglistR14      *SlCbrPreconfigtxconfiglistR14
}
