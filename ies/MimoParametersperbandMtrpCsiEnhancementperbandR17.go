package ies

import "rrc/utils"

// MIMO-ParametersPerBand-mTRP-CSI-EnhancementPerBand-r17 ::= SEQUENCE
type MimoParametersperbandMtrpCsiEnhancementperbandR17 struct {
	MaxnumnzpCsiRsR17          utils.INTEGER `lb:0,ub:8`
	CsiReportModeR17           MimoParametersperbandMtrpCsiEnhancementperbandR17CsiReportModeR17
	SupportedcomboacrossccsR17 []CsiMultitrpSupportedcombinationsR17 `lb:1,ub:16`
	CodebookmodencjtR17        MimoParametersperbandMtrpCsiEnhancementperbandR17CodebookmodencjtR17
}
