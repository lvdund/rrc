package ies

import "rrc/utils"

// SharedSpectrumChAccessParamsPerBand-r16-configuredUL-Tx-r16 ::= ENUMERATED
type SharedspectrumchaccessparamsperbandR16ConfiguredulTxR16 struct {
	Value utils.ENUMERATED
}

const (
	SharedspectrumchaccessparamsperbandR16ConfiguredulTxR16EnumeratedNothing = iota
	SharedspectrumchaccessparamsperbandR16ConfiguredulTxR16EnumeratedSupported
)
