package ies

// PerRAInfoList-v1660 ::= SEQUENCE OF PerRACSI-RSInfo-v1660
// SIZE (1..200)
type PerrainfolistV1660 struct {
	Value []PerracsiRsinfoV1660 `lb:1,ub:200`
}
