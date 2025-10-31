package ies

import "rrc/utils"

// UDT-RestrictingPerPLMN-r13 ::= SEQUENCE
type UdtRestrictingperplmnR13 struct {
	PlmnIdentityindexR13 utils.INTEGER `lb:0,ub:maxPLMNR11`
	UdtRestrictingR13    *UdtRestrictingR13
}
