package ies

import "rrc/utils"

// SharedSpectrumChAccessParamsPerBand-v1640-csi-RSRP-AndRSRQ-MeasWithoutSSB-r16 ::= ENUMERATED
type SharedspectrumchaccessparamsperbandV1640CsiRsrpAndrsrqMeaswithoutssbR16 struct {
	Value utils.ENUMERATED
}

const (
	SharedspectrumchaccessparamsperbandV1640CsiRsrpAndrsrqMeaswithoutssbR16EnumeratedNothing = iota
	SharedspectrumchaccessparamsperbandV1640CsiRsrpAndrsrqMeaswithoutssbR16EnumeratedSupported
)
