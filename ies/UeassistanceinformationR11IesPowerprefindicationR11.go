package ies

import "rrc/utils"

// UEAssistanceInformation-r11-IEs-powerPrefIndication-r11 ::= ENUMERATED
type UeassistanceinformationR11IesPowerprefindicationR11 struct {
	Value utils.ENUMERATED
}

const (
	UeassistanceinformationR11IesPowerprefindicationR11EnumeratedNothing = iota
	UeassistanceinformationR11IesPowerprefindicationR11EnumeratedNormal
	UeassistanceinformationR11IesPowerprefindicationR11EnumeratedLowpowerconsumption
)
