package ies

// CellReselectionPriorities ::= SEQUENCE
// Extensible
type Cellreselectionpriorities struct {
	Freqprioritylisteutra               *Freqprioritylisteutra
	Freqprioritylistnr                  *Freqprioritylistnr
	T320                                *CellreselectionprioritiesT320
	FreqprioritylistdedicatedslicingR17 *FreqprioritylistdedicatedslicingR17 `ext`
}
