package ies

// SL-RequestedSIB-List-r17 ::= SEQUENCE OF SL-SIB-ReqInfo-r17
// SIZE (maxSIB-MessagePlus1-r17)
type SlRequestedsibListR17 struct {
	Value []SlSibReqinfoR17 `lb:maxSIBMessageplus1R17,ub:maxSIBMessageplus1R17`
}
