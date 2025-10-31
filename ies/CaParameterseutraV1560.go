package ies

import "rrc/utils"

// CA-ParametersEUTRA-v1560 ::= SEQUENCE
type CaParameterseutraV1560 struct {
	FdMimoTotalweightedlayers *utils.INTEGER `lb:0,ub:128`
}
