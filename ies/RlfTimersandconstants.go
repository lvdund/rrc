package ies

// RLF-TimersAndConstants ::= SEQUENCE
// Extensible
type RlfTimersandconstants struct {
	T310 RlfTimersandconstantsT310
	N310 RlfTimersandconstantsN310
	N311 RlfTimersandconstantsN311
	T311 RlfTimersandconstantsT311 `ext`
}
