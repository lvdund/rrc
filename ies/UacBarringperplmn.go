package ies

import "rrc/utils"

// UAC-BarringPerPLMN ::= SEQUENCE
type UacBarringperplmn struct {
	PlmnIdentityindex    utils.INTEGER                          `lb:0,ub:maxPLMN`
	UacAcbarringlisttype *UacBarringperplmnUacAcbarringlisttype `lb:maxAccessCat1,ub:maxAccessCat1`
}
