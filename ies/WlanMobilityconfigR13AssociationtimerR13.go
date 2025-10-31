package ies

import "rrc/utils"

// WLAN-MobilityConfig-r13-associationTimer-r13 ::= ENUMERATED
type WlanMobilityconfigR13AssociationtimerR13 struct {
	Value utils.ENUMERATED
}

const (
	WlanMobilityconfigR13AssociationtimerR13EnumeratedNothing = iota
	WlanMobilityconfigR13AssociationtimerR13EnumeratedS10
	WlanMobilityconfigR13AssociationtimerR13EnumeratedS30
	WlanMobilityconfigR13AssociationtimerR13EnumeratedS60
	WlanMobilityconfigR13AssociationtimerR13EnumeratedS120
	WlanMobilityconfigR13AssociationtimerR13EnumeratedS240
)
