package ies

import "rrc/utils"

// IRAT-ParametersNR-v1610-ce-EUTRA-5GC-HO-ToNR-TDD-FR2-r16 ::= ENUMERATED
type IratParametersnrV1610CeEutra5gcHoTonrTddFr2R16 struct {
	Value utils.ENUMERATED
}

const (
	IratParametersnrV1610CeEutra5gcHoTonrTddFr2R16EnumeratedNothing = iota
	IratParametersnrV1610CeEutra5gcHoTonrTddFr2R16EnumeratedSupported
)
