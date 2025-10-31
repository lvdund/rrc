package ies

// NeighCellsCRS-Info-r11 ::= CHOICE
const (
	NeighcellscrsInfoR11ChoiceNothing = iota
	NeighcellscrsInfoR11ChoiceRelease
	NeighcellscrsInfoR11ChoiceSetup
)

type NeighcellscrsInfoR11 struct {
	Choice  uint64
	Release *struct{}
	Setup   *CrsAssistanceinfolistR11
}
