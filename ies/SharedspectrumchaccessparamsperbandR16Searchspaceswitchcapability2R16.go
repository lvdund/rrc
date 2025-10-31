package ies

import "rrc/utils"

// SharedSpectrumChAccessParamsPerBand-r16-searchSpaceSwitchCapability2-r16 ::= ENUMERATED
type SharedspectrumchaccessparamsperbandR16Searchspaceswitchcapability2R16 struct {
	Value utils.ENUMERATED
}

const (
	SharedspectrumchaccessparamsperbandR16Searchspaceswitchcapability2R16EnumeratedNothing = iota
	SharedspectrumchaccessparamsperbandR16Searchspaceswitchcapability2R16EnumeratedSupported
)
