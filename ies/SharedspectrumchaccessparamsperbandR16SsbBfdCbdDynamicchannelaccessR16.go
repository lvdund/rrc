package ies

import "rrc/utils"

// SharedSpectrumChAccessParamsPerBand-r16-ssb-BFD-CBD-dynamicChannelAccess-r16 ::= ENUMERATED
type SharedspectrumchaccessparamsperbandR16SsbBfdCbdDynamicchannelaccessR16 struct {
	Value utils.ENUMERATED
}

const (
	SharedspectrumchaccessparamsperbandR16SsbBfdCbdDynamicchannelaccessR16EnumeratedNothing = iota
	SharedspectrumchaccessparamsperbandR16SsbBfdCbdDynamicchannelaccessR16EnumeratedSupported
)
