package ies

import "rrc/utils"

// PLMN-IdentityInfo-v1610 ::= SEQUENCE
type PlmnIdentityinfoV1610 struct {
	CpCiot5gsOptimisationR16 *utils.ENUMERATED
	UpCiot5gsOptimisationR16 *utils.ENUMERATED
	IabSupportR16            *utils.ENUMERATED
}
