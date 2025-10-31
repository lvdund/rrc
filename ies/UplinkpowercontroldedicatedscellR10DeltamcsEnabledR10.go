package ies

import "rrc/utils"

// UplinkPowerControlDedicatedSCell-r10-deltaMCS-Enabled-r10 ::= ENUMERATED
type UplinkpowercontroldedicatedscellR10DeltamcsEnabledR10 struct {
	Value utils.ENUMERATED
}

const (
	UplinkpowercontroldedicatedscellR10DeltamcsEnabledR10EnumeratedNothing = iota
	UplinkpowercontroldedicatedscellR10DeltamcsEnabledR10EnumeratedEn0
	UplinkpowercontroldedicatedscellR10DeltamcsEnabledR10EnumeratedEn1
)
