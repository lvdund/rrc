package ies

import "rrc/utils"

// N1PUCCH-AN-InfoList-r13 ::= SEQUENCE OF utils.INTEGER // SIZE (1..maxCE-Level-r13)
type N1pucchAnInfolistR13 struct {
	Value []utils.INTEGER `lb:1,ub:maxCELevelR13`
}
