package ies

import "rrc/utils"

// CA-ParametersEUTRA-v1570 ::= SEQUENCE
type CaParameterseutraV1570 struct {
	Dl1024qamTotalweightedlayers *utils.INTEGER `lb:0,ub:10`
}
