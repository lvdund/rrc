package ies

// PUSCH-ConfigDedicated-v1530-ce-PUSCH-SubPRB-Config-r15 ::= CHOICE
const (
	PuschConfigdedicatedV1530CePuschSubprbConfigR15ChoiceNothing = iota
	PuschConfigdedicatedV1530CePuschSubprbConfigR15ChoiceRelease
	PuschConfigdedicatedV1530CePuschSubprbConfigR15ChoiceSetup
)

type PuschConfigdedicatedV1530CePuschSubprbConfigR15 struct {
	Choice  uint64
	Release *struct{}
	Setup   *PuschConfigdedicatedV1530CePuschSubprbConfigR15Setup
}
