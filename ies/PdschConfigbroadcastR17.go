package ies

// PDSCH-ConfigBroadcast-r17 ::= SEQUENCE
type PdschConfigbroadcastR17 struct {
	PdschconfiglistR17               []PdschConfigptmR17 `lb:1,ub:maxNrofPDSCHConfigptmR17`
	PdschTimedomainallocationlistR17 *PdschTimedomainresourceallocationlistR16
	RatematchpatterntoaddmodlistR17  *[]Ratematchpattern `lb:1,ub:maxNrofRateMatchPatterns`
	LteCrsTomatcharoundR17           *RatematchpatternlteCrs
	McsTableR17                      *PdschConfigbroadcastR17McsTableR17
	XoverheadR17                     *PdschConfigbroadcastR17XoverheadR17
}
