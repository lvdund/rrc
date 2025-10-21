package ies

import "rrc/utils"

// SCellToAddModListSCG-r12 ::= SEQUENCE OF Cell-ToAddMod-r12
// SIZE (1..maxSCell-r10)
type ScelltoaddmodlistscgR12 struct {
	Value utils.Sequence[CellToaddmodR12]
}
