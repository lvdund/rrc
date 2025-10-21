package ies

import "rrc/utils"

// SCellToAddModList-r10 ::= SEQUENCE OF SCellToAddMod-r10
// SIZE (1..maxSCell-r10)
type ScelltoaddmodlistR10 struct {
	Value []ScelltoaddmodR10
}
