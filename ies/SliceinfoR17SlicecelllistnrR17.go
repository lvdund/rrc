package ies

// SliceInfo-r17-sliceCellListNR-r17 ::= CHOICE
const (
	SliceinfoR17SlicecelllistnrR17ChoiceNothing = iota
	SliceinfoR17SlicecelllistnrR17ChoiceSliceallowedcelllistnrR17
	SliceinfoR17SlicecelllistnrR17ChoiceSliceexcludedcelllistnrR17
)

type SliceinfoR17SlicecelllistnrR17 struct {
	Choice                     uint64
	SliceallowedcelllistnrR17  *SlicecelllistnrR17
	SliceexcludedcelllistnrR17 *SlicecelllistnrR17
}
