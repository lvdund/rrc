package ies

// SL-DiscTxResource-r13 ::= CHOICE
const (
	SlDisctxresourceR13ChoiceNothing = iota
	SlDisctxresourceR13ChoiceRelease
	SlDisctxresourceR13ChoiceSetup
)

type SlDisctxresourceR13 struct {
	Choice  uint64
	Release *struct{}
	Setup   *SlDisctxresourceR13Setup
}
