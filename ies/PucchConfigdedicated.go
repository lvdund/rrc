package ies

// PUCCH-ConfigDedicated ::= SEQUENCE
type PucchConfigdedicated struct {
	Acknackrepetition      PucchConfigdedicatedAcknackrepetition
	TddAcknackfeedbackmode *PucchConfigdedicatedTddAcknackfeedbackmode
}
