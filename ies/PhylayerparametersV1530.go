package ies

// PhyLayerParameters-v1530 ::= SEQUENCE
type PhylayerparametersV1530 struct {
	SttiSptCapabilitiesR15        *PhylayerparametersV1530SttiSptCapabilitiesR15
	CeCapabilitiesR15             *PhylayerparametersV1530CeCapabilitiesR15
	ShortcqiForscellactivationR15 *PhylayerparametersV1530ShortcqiForscellactivationR15
	MimoCbsrAdvancedcsiR15        *PhylayerparametersV1530MimoCbsrAdvancedcsiR15
	CrsIntfmitigR15               *PhylayerparametersV1530CrsIntfmitigR15
	UlPowercontrolenhancementsR15 *PhylayerparametersV1530UlPowercontrolenhancementsR15
	UrllcCapabilitiesR15          *PhylayerparametersV1530UrllcCapabilitiesR15
	AltmcsTableR15                *PhylayerparametersV1530AltmcsTableR15
}
