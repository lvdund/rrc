package ies

// SPUCCH-Config-r15 ::= CHOICE
const (
	SpucchConfigR15ChoiceNothing = iota
	SpucchConfigR15ChoiceRelease
	SpucchConfigR15ChoiceSetup
)

type SpucchConfigR15 struct {
	Choice  uint64
	Release *struct{}
	Setup   *SpucchConfigR15Setup
}
