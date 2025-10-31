package ies

import "rrc/utils"

// IRAT-ParametersNR-v1610-ce-EUTRA-5GC-HO-ToNR-FDD-FR1-r16 ::= ENUMERATED
type IratParametersnrV1610CeEutra5gcHoTonrFddFr1R16 struct {
	Value utils.ENUMERATED
}

const (
	IratParametersnrV1610CeEutra5gcHoTonrFddFr1R16EnumeratedNothing = iota
	IratParametersnrV1610CeEutra5gcHoTonrFddFr1R16EnumeratedSupported
)
