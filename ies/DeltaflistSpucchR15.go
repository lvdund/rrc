package ies

// DeltaFList-SPUCCH-r15 ::= CHOICE
// Extensible
const (
	DeltaflistSpucchR15ChoiceNothing = iota
	DeltaflistSpucchR15ChoiceRelease
	DeltaflistSpucchR15ChoiceSetup
)

type DeltaflistSpucchR15 struct {
	Choice  uint64
	Release *struct{}
	Setup   *DeltaflistSpucchR15Setup
}
