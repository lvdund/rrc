package ies

// MeasGapConfig ::= CHOICE
// Extensible
const (
	MeasgapconfigChoiceNothing = iota
	MeasgapconfigChoiceRelease
	MeasgapconfigChoiceSetup
)

type Measgapconfig struct {
	Choice  uint64
	Release *struct{}
	Setup   *MeasgapconfigSetup
}
