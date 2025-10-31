package ies

// IdleModeMobilityControlInfo ::= SEQUENCE
// Extensible
type Idlemodemobilitycontrolinfo struct {
	Freqprioritylisteutra      *Freqprioritylisteutra
	Freqprioritylistgeran      *Freqsprioritylistgeran
	FreqprioritylistutraFdd    *FreqprioritylistutraFdd
	FreqprioritylistutraTdd    *FreqprioritylistutraTdd
	Bandclassprioritylisthrpd  *Bandclassprioritylisthrpd
	Bandclassprioritylist1xrtt *Bandclassprioritylist1xrtt
	T320                       *IdlemodemobilitycontrolinfoT320
}
