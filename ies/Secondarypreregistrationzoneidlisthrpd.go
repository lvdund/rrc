package ies

// SecondaryPreRegistrationZoneIdListHRPD ::= SEQUENCE OF PreRegistrationZoneIdHRPD
// SIZE (1..2)
type Secondarypreregistrationzoneidlisthrpd struct {
	Value []Preregistrationzoneidhrpd `lb:1,ub:2`
}
