package ies

// CG-ConfigInfo-v1700-IEs ::= SEQUENCE
type CgConfiginfoV1700 struct {
	CandidatecelllistcpcR17                  *CandidatecelllistcpcR17
	TwophrmodemcgR17                         *CgConfiginfoV1700IesTwophrmodemcgR17
	LowmobilityevaluationconnectedinpcellR17 *CgConfiginfoV1700IesLowmobilityevaluationconnectedinpcellR17
	Noncriticalextension                     *CgConfiginfoV1730
}
