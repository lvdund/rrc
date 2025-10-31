package ies

import "rrc/utils"

// SharedSpectrumChAccessParamsPerBand-r16-extCP-rangeCG-PUSCH-r16 ::= ENUMERATED
type SharedspectrumchaccessparamsperbandR16ExtcpRangecgPuschR16 struct {
	Value utils.ENUMERATED
}

const (
	SharedspectrumchaccessparamsperbandR16ExtcpRangecgPuschR16EnumeratedNothing = iota
	SharedspectrumchaccessparamsperbandR16ExtcpRangecgPuschR16EnumeratedSupported
)
