package ies

import "rrc/utils"

// SharedSpectrumChAccessParamsPerBand-r16-mux-CG-UCI-HARQ-ACK-r16 ::= ENUMERATED
type SharedspectrumchaccessparamsperbandR16MuxCgUciHarqAckR16 struct {
	Value utils.ENUMERATED
}

const (
	SharedspectrumchaccessparamsperbandR16MuxCgUciHarqAckR16EnumeratedNothing = iota
	SharedspectrumchaccessparamsperbandR16MuxCgUciHarqAckR16EnumeratedSupported
)
