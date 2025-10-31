package ies

import "rrc/utils"

// SIB-Type-v12j0 ::= utils.ENUMERATED // Extensible
type SibTypeV12j0 struct {
	Value utils.ENUMERATED
}

const (
	SibTypeV12j0EnumeratedNothing = iota
	SibTypeV12j0EnumeratedSibtype19_V1250
	SibTypeV12j0EnumeratedSibtype20_V1310
	SibTypeV12j0EnumeratedSibtype21_V1430
	SibTypeV12j0EnumeratedSibtype24_V1530
	SibTypeV12j0EnumeratedSibtype25_V1530
	SibTypeV12j0EnumeratedSibtype26_V1530
	SibTypeV12j0EnumeratedSibtype26a_V1610
	SibTypeV12j0EnumeratedSibtype27_V1610
	SibTypeV12j0EnumeratedSibtype28_V1610
	SibTypeV12j0EnumeratedSibtype29_V1610
	SibTypeV12j0EnumeratedSpare6
	SibTypeV12j0EnumeratedSpare5
	SibTypeV12j0EnumeratedSpare4
	SibTypeV12j0EnumeratedSpare3
	SibTypeV12j0EnumeratedSpare2
	SibTypeV12j0EnumeratedSpare1
)
