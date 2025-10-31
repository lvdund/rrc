package ies

// CG-Config-v1700-IEs ::= SEQUENCE
type CgConfigV1700 struct {
	CandidatecellinfolistcpcR17 *CandidatecellinfolistcpcR17
	TwophrmodescgR17            *CgConfigV1700IesTwophrmodescgR17
	Noncriticalextension        *CgConfigV1730
}
