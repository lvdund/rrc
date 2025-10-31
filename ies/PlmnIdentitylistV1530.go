package ies

// PLMN-IdentityList-v1530 ::= SEQUENCE OF PLMN-IdentityInfo-v1530
// SIZE (1..maxPLMN-r11)
type PlmnIdentitylistV1530 struct {
	Value []PlmnIdentityinfoV1530 `lb:1,ub:maxPLMNR11`
}
