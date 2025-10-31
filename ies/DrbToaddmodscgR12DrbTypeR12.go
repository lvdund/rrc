package ies

// DRB-ToAddModSCG-r12-drb-Type-r12 ::= CHOICE
const (
	DrbToaddmodscgR12DrbTypeR12ChoiceNothing = iota
	DrbToaddmodscgR12DrbTypeR12ChoiceSplitR12
	DrbToaddmodscgR12DrbTypeR12ChoiceScgR12
)

type DrbToaddmodscgR12DrbTypeR12 struct {
	Choice   uint64
	SplitR12 *struct{}
	ScgR12   *DrbToaddmodscgR12DrbTypeR12ScgR12
}
