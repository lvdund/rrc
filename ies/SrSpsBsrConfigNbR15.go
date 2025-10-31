package ies

// SR-SPS-BSR-Config-NB-r15 ::= CHOICE
const (
	SrSpsBsrConfigNbR15ChoiceNothing = iota
	SrSpsBsrConfigNbR15ChoiceRelease
	SrSpsBsrConfigNbR15ChoiceSetup
)

type SrSpsBsrConfigNbR15 struct {
	Choice  uint64
	Release *struct{}
	Setup   *SrSpsBsrConfigNbR15Setup
}
