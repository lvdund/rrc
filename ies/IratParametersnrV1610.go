package ies

import "rrc/utils"

// IRAT-ParametersNR-v1610 ::= SEQUENCE
type IratParametersnrV1610 struct {
	NrHoToenDcR16             *utils.ENUMERATED
	CeEutra5gcHoTonrFddFr1R16 *utils.ENUMERATED
	CeEutra5gcHoTonrTddFr1R16 *utils.ENUMERATED
	CeEutra5gcHoTonrFddFr2R16 *utils.ENUMERATED
	CeEutra5gcHoTonrTddFr2R16 *utils.ENUMERATED
}
