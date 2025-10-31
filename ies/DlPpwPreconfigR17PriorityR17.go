package ies

import "rrc/utils"

// DL-PPW-PreConfig-r17-priority-r17 ::= ENUMERATED
type DlPpwPreconfigR17PriorityR17 struct {
	Value utils.ENUMERATED
}

const (
	DlPpwPreconfigR17PriorityR17EnumeratedNothing = iota
	DlPpwPreconfigR17PriorityR17EnumeratedSt1
	DlPpwPreconfigR17PriorityR17EnumeratedSt2
	DlPpwPreconfigR17PriorityR17EnumeratedSt3
)
