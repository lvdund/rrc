package ies

import "rrc/utils"

// CellGlobalIdNR-r16 ::= SEQUENCE
type CellglobalidnrR16 struct {
	PlmnIdentityR16     PlmnIdentity
	CellidentityR16     CellidentitynrR15
	TrackingareacodeR16 *TrackingareacodenrR15
}
