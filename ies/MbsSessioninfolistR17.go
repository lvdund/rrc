package ies

// MBS-SessionInfoList-r17 ::= SEQUENCE OF MBS-SessionInfo-r17
// SIZE (1..maxNrofMBS-Session-r17)
type MbsSessioninfolistR17 struct {
	Value []MbsSessioninfoR17 `lb:1,ub:maxNrofMBSSessionR17`
}
