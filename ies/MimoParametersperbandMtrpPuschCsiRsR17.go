package ies

import "rrc/utils"

// MIMO-ParametersPerBand-mTRP-PUSCH-CSI-RS-r17 ::= SEQUENCE
type MimoParametersperbandMtrpPuschCsiRsR17 struct {
	MaxnumperiodicsrsR17         utils.INTEGER `lb:0,ub:8`
	MaxnumaperiodicsrsR17        utils.INTEGER `lb:0,ub:8`
	MaxnumspSrsR17               utils.INTEGER `lb:0,ub:8`
	NumsrsResourceperccR17       utils.INTEGER `lb:0,ub:16`
	NumsrsResourcenoncodebookR17 utils.INTEGER `lb:0,ub:2`
}
