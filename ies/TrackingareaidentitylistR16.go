package ies

// TrackingAreaIdentityList-r16 ::= SEQUENCE OF TrackingAreaIdentity-r16
// SIZE (1..8)
type TrackingareaidentitylistR16 struct {
	Value []TrackingareaidentityR16 `lb:1,ub:8`
}
