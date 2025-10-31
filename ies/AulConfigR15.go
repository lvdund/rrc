package ies

// AUL-Config-r15 ::= CHOICE
const (
	AulConfigR15ChoiceNothing = iota
	AulConfigR15ChoiceRelease
	AulConfigR15ChoiceSetup
)

type AulConfigR15 struct {
	Choice  uint64
	Release *struct{}
	Setup   *AulConfigR15Setup
}
