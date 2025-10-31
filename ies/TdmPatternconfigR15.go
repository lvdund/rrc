package ies

// TDM-PatternConfig-r15 ::= CHOICE
const (
	TdmPatternconfigR15ChoiceNothing = iota
	TdmPatternconfigR15ChoiceRelease
	TdmPatternconfigR15ChoiceSetup
)

type TdmPatternconfigR15 struct {
	Choice  uint64
	Release *struct{}
	Setup   *TdmPatternconfigR15Setup
}
