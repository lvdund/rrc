package ies

import "rrc/utils"

// SCellToAddModExt-v1430 ::= SEQUENCE
// Extensible
type ScelltoaddmodextV1430 struct {
	SrsSwitchfromservcellindexR14 *utils.INTEGER `lb:0,ub:31`
}
