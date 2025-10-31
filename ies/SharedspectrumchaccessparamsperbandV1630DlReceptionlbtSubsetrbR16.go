package ies

import "rrc/utils"

// SharedSpectrumChAccessParamsPerBand-v1630-dl-ReceptionLBT-subsetRB-r16 ::= ENUMERATED
type SharedspectrumchaccessparamsperbandV1630DlReceptionlbtSubsetrbR16 struct {
	Value utils.ENUMERATED
}

const (
	SharedspectrumchaccessparamsperbandV1630DlReceptionlbtSubsetrbR16EnumeratedNothing = iota
	SharedspectrumchaccessparamsperbandV1630DlReceptionlbtSubsetrbR16EnumeratedSupported
)
