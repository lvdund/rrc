package ies

import "rrc/utils"

// IRAT-ParametersNR-v1540-eutra-5GC-HO-ToNR-TDD-FR1-r15 ::= ENUMERATED
type IratParametersnrV1540Eutra5gcHoTonrTddFr1R15 struct {
	Value utils.ENUMERATED
}

const (
	IratParametersnrV1540Eutra5gcHoTonrTddFr1R15EnumeratedNothing = iota
	IratParametersnrV1540Eutra5gcHoTonrTddFr1R15EnumeratedSupported
)
