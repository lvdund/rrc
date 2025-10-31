package ies

import "rrc/utils"

// SharedSpectrumChAccessParamsPerBand-r16-ul-ChannelBW-SCell-10mhz-r16 ::= ENUMERATED
type SharedspectrumchaccessparamsperbandR16UlChannelbwScell10mhzR16 struct {
	Value utils.ENUMERATED
}

const (
	SharedspectrumchaccessparamsperbandR16UlChannelbwScell10mhzR16EnumeratedNothing = iota
	SharedspectrumchaccessparamsperbandR16UlChannelbwScell10mhzR16EnumeratedSupported
)
