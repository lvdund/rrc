package ies

import "rrc/utils"

// ResourceReservationConfig-NB-r16 ::= SEQUENCE
// Extensible
type ResourcereservationconfigNbR16 struct {
	PeriodicityR16         ResourcereservationconfigNbR16PeriodicityR16
	StartpositionR16       utils.INTEGER `lb:0,ub:15`
	ResourcereservationR16 *ResourcereservationconfigNbR16ResourcereservationR16
}
