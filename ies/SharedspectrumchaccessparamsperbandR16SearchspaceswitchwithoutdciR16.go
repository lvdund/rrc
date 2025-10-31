package ies

import "rrc/utils"

// SharedSpectrumChAccessParamsPerBand-r16-searchSpaceSwitchWithoutDCI-r16 ::= ENUMERATED
type SharedspectrumchaccessparamsperbandR16SearchspaceswitchwithoutdciR16 struct {
	Value utils.ENUMERATED
}

const (
	SharedspectrumchaccessparamsperbandR16SearchspaceswitchwithoutdciR16EnumeratedNothing = iota
	SharedspectrumchaccessparamsperbandR16SearchspaceswitchwithoutdciR16EnumeratedSupported
)
