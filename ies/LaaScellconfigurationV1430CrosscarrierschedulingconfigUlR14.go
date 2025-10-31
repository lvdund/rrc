package ies

// LAA-SCellConfiguration-v1430-crossCarrierSchedulingConfig-UL-r14 ::= CHOICE
const (
	LaaScellconfigurationV1430CrosscarrierschedulingconfigUlR14ChoiceNothing = iota
	LaaScellconfigurationV1430CrosscarrierschedulingconfigUlR14ChoiceRelease
	LaaScellconfigurationV1430CrosscarrierschedulingconfigUlR14ChoiceSetup
)

type LaaScellconfigurationV1430CrosscarrierschedulingconfigUlR14 struct {
	Choice  uint64
	Release *struct{}
	Setup   *LaaScellconfigurationV1430CrosscarrierschedulingconfigUlR14Setup
}
