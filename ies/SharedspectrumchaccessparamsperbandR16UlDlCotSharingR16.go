package ies

import "rrc/utils"

// SharedSpectrumChAccessParamsPerBand-r16-ul-DL-COT-Sharing-r16 ::= ENUMERATED
type SharedspectrumchaccessparamsperbandR16UlDlCotSharingR16 struct {
	Value utils.ENUMERATED
}

const (
	SharedspectrumchaccessparamsperbandR16UlDlCotSharingR16EnumeratedNothing = iota
	SharedspectrumchaccessparamsperbandR16UlDlCotSharingR16EnumeratedSupported
)
