package ies

// PLMN-IdentityList-v1610 ::= SEQUENCE OF PLMN-IdentityInfo-v1610
// SIZE (1..maxPLMN-r11)
type PlmnIdentitylistV1610 struct {
	Value []PlmnIdentityinfoV1610 `lb:1,ub:maxPLMNR11`
}
