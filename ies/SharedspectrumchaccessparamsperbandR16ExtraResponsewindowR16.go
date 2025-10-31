package ies

import "rrc/utils"

// SharedSpectrumChAccessParamsPerBand-r16-extRA-ResponseWindow-r16 ::= ENUMERATED
type SharedspectrumchaccessparamsperbandR16ExtraResponsewindowR16 struct {
	Value utils.ENUMERATED
}

const (
	SharedspectrumchaccessparamsperbandR16ExtraResponsewindowR16EnumeratedNothing = iota
	SharedspectrumchaccessparamsperbandR16ExtraResponsewindowR16EnumeratedSupported
)
