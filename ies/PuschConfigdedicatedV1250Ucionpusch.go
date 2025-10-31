package ies

// PUSCH-ConfigDedicated-v1250-uciOnPUSCH ::= CHOICE
const (
	PuschConfigdedicatedV1250UcionpuschChoiceNothing = iota
	PuschConfigdedicatedV1250UcionpuschChoiceRelease
	PuschConfigdedicatedV1250UcionpuschChoiceSetup
)

type PuschConfigdedicatedV1250Ucionpusch struct {
	Choice  uint64
	Release *struct{}
	Setup   *PuschConfigdedicatedV1250UcionpuschSetup
}
