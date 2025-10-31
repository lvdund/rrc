package ies

import "rrc/utils"

// IRAT-ParametersNR-v1540-eutra-EPC-HO-ToNR-FDD-FR1-r15 ::= ENUMERATED
type IratParametersnrV1540EutraEpcHoTonrFddFr1R15 struct {
	Value utils.ENUMERATED
}

const (
	IratParametersnrV1540EutraEpcHoTonrFddFr1R15EnumeratedNothing = iota
	IratParametersnrV1540EutraEpcHoTonrFddFr1R15EnumeratedSupported
)
