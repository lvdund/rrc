package ies

import "rrc/utils"

// UplinkPowerControlDedicated-deltaMCS-Enabled ::= ENUMERATED
type UplinkpowercontroldedicatedDeltamcsEnabled struct {
	Value utils.ENUMERATED
}

const (
	UplinkpowercontroldedicatedDeltamcsEnabledEnumeratedNothing = iota
	UplinkpowercontroldedicatedDeltamcsEnabledEnumeratedEn0
	UplinkpowercontroldedicatedDeltamcsEnabledEnumeratedEn1
)
