package ies

import "rrc/utils"

// SharedSpectrumChAccessParamsPerBand-v1640-csi-RSRP-AndRSRQ-MeasWithSSB-r16 ::= ENUMERATED
type SharedspectrumchaccessparamsperbandV1640CsiRsrpAndrsrqMeaswithssbR16 struct {
	Value utils.ENUMERATED
}

const (
	SharedspectrumchaccessparamsperbandV1640CsiRsrpAndrsrqMeaswithssbR16EnumeratedNothing = iota
	SharedspectrumchaccessparamsperbandV1640CsiRsrpAndrsrqMeaswithssbR16EnumeratedSupported
)
