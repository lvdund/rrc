package ies

// AreaConfig-r16 ::= CHOICE
const (
	AreaconfigR16ChoiceNothing = iota
	AreaconfigR16ChoiceCellglobalidlistR16
	AreaconfigR16ChoiceTrackingareacodelistR16
	AreaconfigR16ChoiceTrackingareaidentitylistR16
)

type AreaconfigR16 struct {
	Choice                      uint64
	CellglobalidlistR16         *CellglobalidlistR16
	TrackingareacodelistR16     *TrackingareacodelistR16
	TrackingareaidentitylistR16 *TrackingareaidentitylistR16
}
