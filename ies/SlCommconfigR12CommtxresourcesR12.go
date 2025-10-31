package ies

// SL-CommConfig-r12-commTxResources-r12 ::= CHOICE
const (
	SlCommconfigR12CommtxresourcesR12ChoiceNothing = iota
	SlCommconfigR12CommtxresourcesR12ChoiceRelease
	SlCommconfigR12CommtxresourcesR12ChoiceSetup
)

type SlCommconfigR12CommtxresourcesR12 struct {
	Choice  uint64
	Release *struct{}
	Setup   *SlCommconfigR12CommtxresourcesR12Setup
}
