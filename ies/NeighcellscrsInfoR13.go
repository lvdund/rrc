package ies

// NeighCellsCRS-Info-r13 ::= CHOICE
const (
	NeighcellscrsInfoR13ChoiceNothing = iota
	NeighcellscrsInfoR13ChoiceRelease
	NeighcellscrsInfoR13ChoiceSetup
)

type NeighcellscrsInfoR13 struct {
	Choice  uint64
	Release *struct{}
	Setup   *CrsAssistanceinfolistR13
}
