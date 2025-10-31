package ies

import "rrc/utils"

// SharedSpectrumChAccessParamsPerBand-r16-dci-ChOccupancyDuration-r16 ::= ENUMERATED
type SharedspectrumchaccessparamsperbandR16DciChoccupancydurationR16 struct {
	Value utils.ENUMERATED
}

const (
	SharedspectrumchaccessparamsperbandR16DciChoccupancydurationR16EnumeratedNothing = iota
	SharedspectrumchaccessparamsperbandR16DciChoccupancydurationR16EnumeratedSupported
)
