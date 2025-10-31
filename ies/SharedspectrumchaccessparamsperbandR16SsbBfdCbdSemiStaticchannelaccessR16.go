package ies

import "rrc/utils"

// SharedSpectrumChAccessParamsPerBand-r16-ssb-BFD-CBD-semi-staticChannelAccess-r16 ::= ENUMERATED
type SharedspectrumchaccessparamsperbandR16SsbBfdCbdSemiStaticchannelaccessR16 struct {
	Value utils.ENUMERATED
}

const (
	SharedspectrumchaccessparamsperbandR16SsbBfdCbdSemiStaticchannelaccessR16EnumeratedNothing = iota
	SharedspectrumchaccessparamsperbandR16SsbBfdCbdSemiStaticchannelaccessR16EnumeratedSupported
)
