package ies

// RLC-BearerConfig-r15 ::= CHOICE
const (
	RlcBearerconfigR15ChoiceNothing = iota
	RlcBearerconfigR15ChoiceRelease
	RlcBearerconfigR15ChoiceSetup
)

type RlcBearerconfigR15 struct {
	Choice  uint64
	Release *struct{}
	Setup   *RlcBearerconfigR15Setup
}
