package ies

import "rrc/utils"

// SearchSpaceSwitchConfig-r17 ::= SEQUENCE
type SearchspaceswitchconfigR17 struct {
	SearchspaceswitchtimerR17 *ScsSpecificdurationR17
	SearchspaceswitchdelayR17 *utils.INTEGER `lb:0,ub:52`
}
