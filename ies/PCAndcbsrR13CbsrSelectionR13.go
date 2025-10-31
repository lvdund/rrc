package ies

// P-C-AndCBSR-r13-cbsr-Selection-r13 ::= CHOICE
const (
	PCAndcbsrR13CbsrSelectionR13ChoiceNothing = iota
	PCAndcbsrR13CbsrSelectionR13ChoiceNonprecodedR13
	PCAndcbsrR13CbsrSelectionR13ChoiceBeamformedk1aR13
	PCAndcbsrR13CbsrSelectionR13ChoiceBeamformedknR13
)

type PCAndcbsrR13CbsrSelectionR13 struct {
	Choice           uint64
	NonprecodedR13   *PCAndcbsrR13CbsrSelectionR13NonprecodedR13
	Beamformedk1aR13 *PCAndcbsrR13CbsrSelectionR13Beamformedk1aR13
	BeamformedknR13  *PCAndcbsrR13CbsrSelectionR13BeamformedknR13
}
