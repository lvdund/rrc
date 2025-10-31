package ies

// TrackingAreaCodeList-v1130 ::= SEQUENCE
type TrackingareacodelistV1130 struct {
	PlmnIdentityPertacListR11 []PlmnIdentity `lb:1,ub:8`
}
