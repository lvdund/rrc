package ies

import "rrc/utils"

// MIMO-ParametersPerBand-mTRP-GroupBasedL1-RSRP-r17 ::= SEQUENCE
type MimoParametersperbandMtrpGroupbasedl1RsrpR17 struct {
	MaxnumbeamgroupsR17   utils.INTEGER `lb:0,ub:4`
	MaxnumrsWithinslotR17 MimoParametersperbandMtrpGroupbasedl1RsrpR17MaxnumrsWithinslotR17
	MaxnumrsAcrossslotR17 MimoParametersperbandMtrpGroupbasedl1RsrpR17MaxnumrsAcrossslotR17
}
