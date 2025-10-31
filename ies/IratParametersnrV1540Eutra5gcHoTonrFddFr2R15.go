package ies

import "rrc/utils"

// IRAT-ParametersNR-v1540-eutra-5GC-HO-ToNR-FDD-FR2-r15 ::= ENUMERATED
type IratParametersnrV1540Eutra5gcHoTonrFddFr2R15 struct {
	Value utils.ENUMERATED
}

const (
	IratParametersnrV1540Eutra5gcHoTonrFddFr2R15EnumeratedNothing = iota
	IratParametersnrV1540Eutra5gcHoTonrFddFr2R15EnumeratedSupported
)
