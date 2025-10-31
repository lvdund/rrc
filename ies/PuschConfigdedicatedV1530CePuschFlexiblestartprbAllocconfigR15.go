package ies

// PUSCH-ConfigDedicated-v1530-ce-PUSCH-FlexibleStartPRB-AllocConfig-r15 ::= CHOICE
const (
	PuschConfigdedicatedV1530CePuschFlexiblestartprbAllocconfigR15ChoiceNothing = iota
	PuschConfigdedicatedV1530CePuschFlexiblestartprbAllocconfigR15ChoiceRelease
	PuschConfigdedicatedV1530CePuschFlexiblestartprbAllocconfigR15ChoiceSetup
)

type PuschConfigdedicatedV1530CePuschFlexiblestartprbAllocconfigR15 struct {
	Choice  uint64
	Release *struct{}
	Setup   *PuschConfigdedicatedV1530CePuschFlexiblestartprbAllocconfigR15Setup
}
