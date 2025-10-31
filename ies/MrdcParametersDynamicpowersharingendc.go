package ies

import "rrc/utils"

// MRDC-Parameters-dynamicPowerSharingENDC ::= ENUMERATED
type MrdcParametersDynamicpowersharingendc struct {
	Value utils.ENUMERATED
}

const (
	MrdcParametersDynamicpowersharingendcEnumeratedNothing = iota
	MrdcParametersDynamicpowersharingendcEnumeratedSupported
)
