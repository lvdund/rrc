package ies

// CSI-RS-ConfigBeamformed-v1430 ::= SEQUENCE
type CsiRsConfigbeamformedV1430 struct {
	CsiRsConfignzpAplistR14        *[]CsiRsConfignzpR11 `lb:1,ub:8`
	NzpResourceconfigoriginalV1430 *CsiRsConfigNzpV1430
	CsiRsNzpActivationR14          *CsiRsConfignzpActivationR14
}
