package ies

// ReportUplinkTxDirectCurrentMoreCarrier-r17 ::= SEQUENCE OF IntraBandCC-CombinationReqList-r17
// SIZE (1.. maxSimultaneousBands)
type ReportuplinktxdirectcurrentmorecarrierR17 struct {
	Value []IntrabandccCombinationreqlistR17 `lb:1,ub:maxSimultaneousBands`
}
