package ies

// AreaConfiguration-r10 ::= CHOICE
const (
	AreaconfigurationR10ChoiceNothing = iota
	AreaconfigurationR10ChoiceCellglobalidlistR10
	AreaconfigurationR10ChoiceTrackingareacodelistR10
)

type AreaconfigurationR10 struct {
	Choice                  uint64
	CellglobalidlistR10     *CellglobalidlistR10
	TrackingareacodelistR10 *TrackingareacodelistR10
}
