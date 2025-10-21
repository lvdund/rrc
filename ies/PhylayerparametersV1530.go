package ies

import "rrc/utils"

// PhyLayerParameters-v1530 ::= SEQUENCE
type PhylayerparametersV1530 struct {
	SttiSptCapabilitiesR15        *interface{}
	CeCapabilitiesR15             *interface{}
	ShortcqiForscellactivationR15 *utils.ENUMERATED
	MimoCbsrAdvancedcsiR15        *utils.ENUMERATED
	CrsIntfmitigR15               *utils.ENUMERATED
	UlPowercontrolenhancementsR15 *utils.ENUMERATED
	UrllcCapabilitiesR15          *interface{}
	AltmcsTableR15                *utils.ENUMERATED
}
