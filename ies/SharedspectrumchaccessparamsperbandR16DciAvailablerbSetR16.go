package ies

import "rrc/utils"

// SharedSpectrumChAccessParamsPerBand-r16-dci-AvailableRB-Set-r16 ::= ENUMERATED
type SharedspectrumchaccessparamsperbandR16DciAvailablerbSetR16 struct {
	Value utils.ENUMERATED
}

const (
	SharedspectrumchaccessparamsperbandR16DciAvailablerbSetR16EnumeratedNothing = iota
	SharedspectrumchaccessparamsperbandR16DciAvailablerbSetR16EnumeratedSupported
)
