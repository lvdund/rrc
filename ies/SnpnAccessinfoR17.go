package ies

import "rrc/utils"

// SNPN-AccessInfo-r17 ::= SEQUENCE
type SnpnAccessinfoR17 struct {
	ExtchSupportedR17             *utils.BOOLEAN
	ExtchWithoutconfigallowedR17  *utils.BOOLEAN
	OnboardingenabledR17          *utils.BOOLEAN
	ImsemergencysupportforsnpnR17 *utils.BOOLEAN
}
