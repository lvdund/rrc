package ies

import "rrc/utils"

// IRAT-ParametersNR-v1540-eutra-5GC-HO-ToNR-TDD-FR2-r15 ::= ENUMERATED
type IratParametersnrV1540Eutra5gcHoTonrTddFr2R15 struct {
	Value utils.ENUMERATED
}

const (
	IratParametersnrV1540Eutra5gcHoTonrTddFr2R15EnumeratedNothing = iota
	IratParametersnrV1540Eutra5gcHoTonrTddFr2R15EnumeratedSupported
)
