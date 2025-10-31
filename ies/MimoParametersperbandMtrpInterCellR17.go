package ies

import "rrc/utils"

// MIMO-ParametersPerBand-mTRP-inter-Cell-r17 ::= SEQUENCE
type MimoParametersperbandMtrpInterCellR17 struct {
	MaxnumadditionalpciCase1R17 utils.INTEGER `lb:0,ub:7`
	MaxnumadditionalpciCase2R17 utils.INTEGER `lb:0,ub:7`
}
