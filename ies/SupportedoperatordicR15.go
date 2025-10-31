package ies

import "rrc/utils"

// SupportedOperatorDic-r15 ::= SEQUENCE
type SupportedoperatordicR15 struct {
	VersionofdictionaryR15 utils.INTEGER `lb:0,ub:15`
	AssociatedplmnIdR15    PlmnIdentity
}
