package ies

import "rrc/utils"

// SupportedOperatorDic-r15 ::= SEQUENCE
type SupportedoperatordicR15 struct {
	VersionofdictionaryR15 utils.INTEGER
	AssociatedplmnIdR15    PlmnIdentity
}
