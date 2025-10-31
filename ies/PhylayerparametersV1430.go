package ies

// PhyLayerParameters-v1430 ::= SEQUENCE
type PhylayerparametersV1430 struct {
	CePuschNbMaxtbsR14                *PhylayerparametersV1430CePuschNbMaxtbsR14
	CePdschPuschMaxbandwidthR14       *PhylayerparametersV1430CePdschPuschMaxbandwidthR14
	CeHarqAckbundlingR14              *PhylayerparametersV1430CeHarqAckbundlingR14
	CePdschTenprocessesR14            *PhylayerparametersV1430CePdschTenprocessesR14
	CeRetuningsymbolsR14              *PhylayerparametersV1430CeRetuningsymbolsR14
	CePdschPuschEnhancementR14        *PhylayerparametersV1430CePdschPuschEnhancementR14
	CeSchedulingenhancementR14        *PhylayerparametersV1430CeSchedulingenhancementR14
	CeSrsEnhancementR14               *PhylayerparametersV1430CeSrsEnhancementR14
	CePucchEnhancementR14             *PhylayerparametersV1430CePucchEnhancementR14
	CeClosedlooptxantennaselectionR14 *PhylayerparametersV1430CeClosedlooptxantennaselectionR14
	TddSpecialsubframeR14             *PhylayerparametersV1430TddSpecialsubframeR14
	TddTtiBundlingR14                 *PhylayerparametersV1430TddTtiBundlingR14
	DmrsLessupptsR14                  *PhylayerparametersV1430DmrsLessupptsR14
	MimoUeParametersV1430             *MimoUeParametersV1430
	AlternativetbsIndexR14            *PhylayerparametersV1430AlternativetbsIndexR14
	FembmsUnicastParametersR14        *FembmsUnicastParametersR14
}
