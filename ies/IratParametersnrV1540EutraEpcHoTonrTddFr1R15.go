package ies

import "rrc/utils"

// IRAT-ParametersNR-v1540-eutra-EPC-HO-ToNR-TDD-FR1-r15 ::= ENUMERATED
type IratParametersnrV1540EutraEpcHoTonrTddFr1R15 struct {
	Value utils.ENUMERATED
}

const (
	IratParametersnrV1540EutraEpcHoTonrTddFr1R15EnumeratedNothing = iota
	IratParametersnrV1540EutraEpcHoTonrTddFr1R15EnumeratedSupported
)
