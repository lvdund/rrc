package ies

// SL-V2X-ConfigDedicated-r14-commTxResources-r14 ::= CHOICE
const (
	SlV2xConfigdedicatedR14CommtxresourcesR14ChoiceNothing = iota
	SlV2xConfigdedicatedR14CommtxresourcesR14ChoiceRelease
	SlV2xConfigdedicatedR14CommtxresourcesR14ChoiceSetup
)

type SlV2xConfigdedicatedR14CommtxresourcesR14 struct {
	Choice  uint64
	Release *struct{}
	Setup   *SlV2xConfigdedicatedR14CommtxresourcesR14Setup
}
