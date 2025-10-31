package ies

import "rrc/utils"

// SearchSpaceSwitchConfig-r16 ::= SEQUENCE
type SearchspaceswitchconfigR16 struct {
	CellgroupsforswitchlistR16 *[]CellgroupforswitchR16 `lb:1,ub:4`
	SearchspaceswitchdelayR16  *utils.INTEGER           `lb:0,ub:52`
}
