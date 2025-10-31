package ies

import "rrc/utils"

// CA-ParametersNR-v1610-crossCarrierA-CSI-trigDiffSCS-r16 ::= ENUMERATED
type CaParametersnrV1610CrosscarrieraCsiTrigdiffscsR16 struct {
	Value utils.ENUMERATED
}

const (
	CaParametersnrV1610CrosscarrieraCsiTrigdiffscsR16EnumeratedNothing = iota
	CaParametersnrV1610CrosscarrieraCsiTrigdiffscsR16EnumeratedHighera_Csi_Scs
	CaParametersnrV1610CrosscarrieraCsiTrigdiffscsR16EnumeratedLowera_Csi_Scs
	CaParametersnrV1610CrosscarrieraCsiTrigdiffscsR16EnumeratedBoth
)
