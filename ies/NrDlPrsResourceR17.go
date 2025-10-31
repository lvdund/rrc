package ies

import "rrc/utils"

// NR-DL-PRS-Resource-r17 ::= SEQUENCE
// Extensible
type NrDlPrsResourceR17 struct {
	NrDlPrsResourceidR17         NrDlPrsResourceidR17
	DlPrsSequenceidR17           utils.INTEGER `lb:0,ub:4095`
	DlPrsCombsizenAndreoffsetR17 NrDlPrsResourceR17DlPrsCombsizenAndreoffsetR17
	DlPrsResourceslotoffsetR17   utils.INTEGER `lb:0,ub:maxNrofPRSResourceoffsetvalue1R17`
	DlPrsResourcesymboloffsetR17 utils.INTEGER `lb:0,ub:12`
	DlPrsQclInfoR17              *DlPrsQclInfoR17
}
