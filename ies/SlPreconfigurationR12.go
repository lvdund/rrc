package ies

// SL-Preconfiguration-r12 ::= SEQUENCE
// Extensible
type SlPreconfigurationR12 struct {
	PreconfiggeneralR12 SlPreconfiggeneralR12
	PreconfigsyncR12    SlPreconfigsyncR12
	PreconfigcommR12    SlPreconfigcommpoollist4R12
}
