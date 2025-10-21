package ies

import "rrc/utils"

// BT-NameList-r15 ::= SEQUENCE OF BT-Name-r15
// SIZE (1..maxBT-Name-r15)
type BtNamelistR15 struct {
	Value utils.Sequence[BtNameR15]
}
