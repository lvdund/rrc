package ies

import "rrc/utils"

// MRDC-Parameters-v1580-dynamicPowerSharingNEDC ::= ENUMERATED
type MrdcParametersV1580Dynamicpowersharingnedc struct {
	Value utils.ENUMERATED
}

const (
	MrdcParametersV1580DynamicpowersharingnedcEnumeratedNothing = iota
	MrdcParametersV1580DynamicpowersharingnedcEnumeratedSupported
)
