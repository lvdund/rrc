package ies

// CrossCarrierSchedulingConfig-schedulingCellInfo ::= CHOICE
const (
	CrosscarrierschedulingconfigSchedulingcellinfoChoiceNothing = iota
	CrosscarrierschedulingconfigSchedulingcellinfoChoiceOwn
	CrosscarrierschedulingconfigSchedulingcellinfoChoiceOther
)

type CrosscarrierschedulingconfigSchedulingcellinfo struct {
	Choice uint64
	Own    *CrosscarrierschedulingconfigSchedulingcellinfoOwn
	Other  *CrosscarrierschedulingconfigSchedulingcellinfoOther
}
