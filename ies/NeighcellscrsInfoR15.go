package ies

// NeighCellsCRS-Info-r15 ::= CHOICE
const (
	NeighcellscrsInfoR15ChoiceNothing = iota
	NeighcellscrsInfoR15ChoiceRelease
	NeighcellscrsInfoR15ChoiceSetup
)

type NeighcellscrsInfoR15 struct {
	Choice  uint64
	Release *struct{}
	Setup   *CrsAssistanceinfolistR15
}
