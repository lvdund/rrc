package ies

import "rrc/utils"

// SharedSpectrumChAccessParamsPerBand-v1710-ul-Semi-StaticChAccessDependentConfig-r17 ::= ENUMERATED
type SharedspectrumchaccessparamsperbandV1710UlSemiStaticchaccessdependentconfigR17 struct {
	Value utils.ENUMERATED
}

const (
	SharedspectrumchaccessparamsperbandV1710UlSemiStaticchaccessdependentconfigR17EnumeratedNothing = iota
	SharedspectrumchaccessparamsperbandV1710UlSemiStaticchaccessdependentconfigR17EnumeratedSupported
)
