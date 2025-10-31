package ies

// PUSCH-ConfigDedicated-r13-uciOnPUSCH ::= CHOICE
const (
	PuschConfigdedicatedR13UcionpuschChoiceNothing = iota
	PuschConfigdedicatedR13UcionpuschChoiceRelease
	PuschConfigdedicatedR13UcionpuschChoiceSetup
)

type PuschConfigdedicatedR13Ucionpusch struct {
	Choice  uint64
	Release *struct{}
	Setup   *PuschConfigdedicatedR13UcionpuschSetup
}
