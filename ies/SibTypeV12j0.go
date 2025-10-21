package ies

import "rrc/utils"

// SIB-Type-v12j0 ::= ENUMERATED
// Extensible
type SibTypeV12j0 struct {
	Value utils.ENUMERATED
}

const (
	SibTypeV12j0Sibtype19V1250  = 0
	SibTypeV12j0Sibtype20V1310  = 1
	SibTypeV12j0Sibtype21V1430  = 2
	SibTypeV12j0Sibtype24V1530  = 3
	SibTypeV12j0Sibtype25V1530  = 4
	SibTypeV12j0Sibtype26V1530  = 5
	SibTypeV12j0Sibtype26aV1610 = 6
	SibTypeV12j0Sibtype27V1610  = 7
	SibTypeV12j0Sibtype28V1610  = 8
	SibTypeV12j0Sibtype29V1610  = 9
	SibTypeV12j0Spare6          = 10
	SibTypeV12j0Spare5          = 11
	SibTypeV12j0Spare4          = 12
	SibTypeV12j0Spare3          = 13
	SibTypeV12j0Spare2          = 14
	SibTypeV12j0Spare1          = 15
)
