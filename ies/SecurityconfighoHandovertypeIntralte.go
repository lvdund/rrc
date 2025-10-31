package ies

import "rrc/utils"

// SecurityConfigHO-handoverType-intraLTE ::= SEQUENCE
type SecurityconfighoHandovertypeIntralte struct {
	Securityalgorithmconfig *Securityalgorithmconfig
	Keychangeindicator      utils.BOOLEAN
	Nexthopchainingcount    Nexthopchainingcount
}
