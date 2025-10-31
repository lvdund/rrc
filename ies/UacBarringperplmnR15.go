package ies

import "rrc/utils"

// UAC-BarringPerPLMN-r15 ::= SEQUENCE
type UacBarringperplmnR15 struct {
	PlmnIdentityindexR15    utils.INTEGER                                `lb:0,ub:maxPLMNR11`
	UacAcBarringlisttypeR15 *UacBarringperplmnR15UacAcBarringlisttypeR15 `lb:maxAccessCat1R15,ub:maxAccessCat1R15`
}
