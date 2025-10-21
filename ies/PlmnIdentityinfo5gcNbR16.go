package ies

import "rrc/utils"

// PLMN-IdentityInfo-5GC-NB-r16 ::= SEQUENCE
type PlmnIdentityinfo5gcNbR16 struct {
	PlmnIdentity5gcR16            interface{}
	CellreservedforoperatoruseR16 utils.ENUMERATED
	NgUDatatransferR16            *utils.ENUMERATED
	UpCiot5gsOptimisationR16      *utils.ENUMERATED
}
