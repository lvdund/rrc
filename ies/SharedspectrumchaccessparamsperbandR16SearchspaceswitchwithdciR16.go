package ies

import "rrc/utils"

// SharedSpectrumChAccessParamsPerBand-r16-searchSpaceSwitchWithDCI-r16 ::= ENUMERATED
type SharedspectrumchaccessparamsperbandR16SearchspaceswitchwithdciR16 struct {
	Value utils.ENUMERATED
}

const (
	SharedspectrumchaccessparamsperbandR16SearchspaceswitchwithdciR16EnumeratedNothing = iota
	SharedspectrumchaccessparamsperbandR16SearchspaceswitchwithdciR16EnumeratedSupported
)
