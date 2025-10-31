package ies

import "rrc/utils"

// PUCCH-ConfigDedicated-v1130-n1PUCCH-AN-CS-v1130-setup ::= SEQUENCE
type PucchConfigdedicatedV1130N1pucchAnCsV1130Setup struct {
	N1pucchAnCsListp1R11 []utils.INTEGER `lb:2,ub:4`
}
