package ies

// RRM-Config ::= SEQUENCE
// Extensible
type RrmConfig struct {
	UeInactivetime               *RrmConfigUeInactivetime
	Candidatecellinfolist        *Measresultlist2nr
	CandidatecellinfolistsnEutra *MeasresultservfreqlisteutraScg `ext`
}
