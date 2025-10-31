package ies

import "rrc/utils"

// SharedSpectrumChAccessParamsPerBand-r16-occ-PRB-PF2-PF3-r16 ::= ENUMERATED
type SharedspectrumchaccessparamsperbandR16OccPrbPf2Pf3R16 struct {
	Value utils.ENUMERATED
}

const (
	SharedspectrumchaccessparamsperbandR16OccPrbPf2Pf3R16EnumeratedNothing = iota
	SharedspectrumchaccessparamsperbandR16OccPrbPf2Pf3R16EnumeratedSupported
)
