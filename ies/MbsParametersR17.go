package ies

import "rrc/utils"

// MBS-Parameters-r17 ::= SEQUENCE
type MbsParametersR17 struct {
	MaxmrbAddR17 *utils.INTEGER `lb:0,ub:16`
}
