package ies

// SPDCCH-Config-r15 ::= CHOICE
const (
	SpdcchConfigR15ChoiceNothing = iota
	SpdcchConfigR15ChoiceRelease
	SpdcchConfigR15ChoiceSetup
)

type SpdcchConfigR15 struct {
	Choice  uint64
	Release *struct{}
	Setup   *SpdcchConfigR15Setup
}
