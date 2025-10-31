package ies

import "rrc/utils"

// SharedSpectrumChAccessParamsPerBand-v1630-dl-ReceptionIntraCellGuardband-r16 ::= ENUMERATED
type SharedspectrumchaccessparamsperbandV1630DlReceptionintracellguardbandR16 struct {
	Value utils.ENUMERATED
}

const (
	SharedspectrumchaccessparamsperbandV1630DlReceptionintracellguardbandR16EnumeratedNothing = iota
	SharedspectrumchaccessparamsperbandV1630DlReceptionintracellguardbandR16EnumeratedSupported
)
