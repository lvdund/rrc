package ies

import "rrc/utils"

// IRAT-ParametersNR-v1540 ::= SEQUENCE
type IratParametersnrV1540 struct {
	Eutra5gcHoTonrFddFr1R15  *utils.ENUMERATED
	Eutra5gcHoTonrTddFr1R15  *utils.ENUMERATED
	Eutra5gcHoTonrFddFr2R15  *utils.ENUMERATED
	Eutra5gcHoTonrTddFr2R15  *utils.ENUMERATED
	EutraEpcHoTonrFddFr1R15  *utils.ENUMERATED
	EutraEpcHoTonrTddFr1R15  *utils.ENUMERATED
	EutraEpcHoTonrFddFr2R15  *utils.ENUMERATED
	EutraEpcHoTonrTddFr2R15  *utils.ENUMERATED
	ImsVoiceovernrFr1R15     *utils.ENUMERATED
	ImsVoiceovernrFr2R15     *utils.ENUMERATED
	SaNrR15                  *utils.ENUMERATED
	SupportedbandlistnrSaR15 *SupportedbandlistnrR15
}
