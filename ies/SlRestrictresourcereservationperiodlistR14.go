package ies

// SL-RestrictResourceReservationPeriodList-r14 ::= SEQUENCE OF SL-RestrictResourceReservationPeriod-r14
// SIZE (1..maxReservationPeriod-r14)
type SlRestrictresourcereservationperiodlistR14 struct {
	Value []SlRestrictresourcereservationperiodR14 `lb:1,ub:maxReservationPeriodR14`
}
