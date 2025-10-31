package ies

// SL-DiscConfig-r12-discTxResources-r12 ::= CHOICE
const (
	SlDiscconfigR12DisctxresourcesR12ChoiceNothing = iota
	SlDiscconfigR12DisctxresourcesR12ChoiceRelease
	SlDiscconfigR12DisctxresourcesR12ChoiceSetup
)

type SlDiscconfigR12DisctxresourcesR12 struct {
	Choice  uint64
	Release *struct{}
	Setup   *SlDiscconfigR12DisctxresourcesR12Setup
}
