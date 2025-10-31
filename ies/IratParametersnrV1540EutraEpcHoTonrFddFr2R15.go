package ies

import "rrc/utils"

// IRAT-ParametersNR-v1540-eutra-EPC-HO-ToNR-FDD-FR2-r15 ::= ENUMERATED
type IratParametersnrV1540EutraEpcHoTonrFddFr2R15 struct {
	Value utils.ENUMERATED
}

const (
	IratParametersnrV1540EutraEpcHoTonrFddFr2R15EnumeratedNothing = iota
	IratParametersnrV1540EutraEpcHoTonrFddFr2R15EnumeratedSupported
)
