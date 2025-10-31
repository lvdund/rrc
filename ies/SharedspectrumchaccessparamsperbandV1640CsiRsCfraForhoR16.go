package ies

import "rrc/utils"

// SharedSpectrumChAccessParamsPerBand-v1640-csi-RS-CFRA-ForHO-r16 ::= ENUMERATED
type SharedspectrumchaccessparamsperbandV1640CsiRsCfraForhoR16 struct {
	Value utils.ENUMERATED
}

const (
	SharedspectrumchaccessparamsperbandV1640CsiRsCfraForhoR16EnumeratedNothing = iota
	SharedspectrumchaccessparamsperbandV1640CsiRsCfraForhoR16EnumeratedSupported
)
