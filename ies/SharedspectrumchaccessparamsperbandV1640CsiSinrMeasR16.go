package ies

import "rrc/utils"

// SharedSpectrumChAccessParamsPerBand-v1640-csi-SINR-Meas-r16 ::= ENUMERATED
type SharedspectrumchaccessparamsperbandV1640CsiSinrMeasR16 struct {
	Value utils.ENUMERATED
}

const (
	SharedspectrumchaccessparamsperbandV1640CsiSinrMeasR16EnumeratedNothing = iota
	SharedspectrumchaccessparamsperbandV1640CsiSinrMeasR16EnumeratedSupported
)
