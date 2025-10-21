package ies

import "rrc/utils"

// PhyLayerParameters-v1430 ::= SEQUENCE
type PhylayerparametersV1430 struct {
	CePuschNbMaxtbsR14                *utils.ENUMERATED
	CePdschPuschMaxbandwidthR14       *utils.ENUMERATED
	CeHarqAckbundlingR14              *utils.ENUMERATED
	CePdschTenprocessesR14            *utils.ENUMERATED
	CeRetuningsymbolsR14              *utils.ENUMERATED
	CePdschPuschEnhancementR14        *utils.ENUMERATED
	CeSchedulingenhancementR14        *utils.ENUMERATED
	CeSrsEnhancementR14               *utils.ENUMERATED
	CePucchEnhancementR14             *utils.ENUMERATED
	CeClosedlooptxantennaselectionR14 *utils.ENUMERATED
	TddSpecialsubframeR14             *utils.ENUMERATED
	TddTtiBundlingR14                 *utils.ENUMERATED
	DmrsLessupptsR14                  *utils.ENUMERATED
	MimoUeParametersV1430             *MimoUeParametersV1430
	AlternativetbsIndexR14            *utils.ENUMERATED
	FembmsUnicastParametersR14        *FembmsUnicastParametersR14
}
