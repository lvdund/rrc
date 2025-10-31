package ies

// DedicatedSIBRequest-r16-IEs-onDemandSIB-RequestList-r16 ::= SEQUENCE
type DedicatedsibrequestR16IesOndemandsibRequestlistR16 struct {
	RequestedsibListR16    *[]SibReqinfoR16    `lb:1,ub:maxOnDemandSIBR16`
	RequestedpossibListR16 *[]PossibReqinfoR16 `lb:1,ub:maxOnDemandPosSIBR16`
}
