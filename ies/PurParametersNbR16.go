package ies

// PUR-Parameters-NB-r16 ::= SEQUENCE
type PurParametersNbR16 struct {
	PurCpEpcR16           *PurParametersNbR16PurCpEpcR16
	PurCp5gcR16           *PurParametersNbR16PurCp5gcR16
	PurUpEpcR16           *PurParametersNbR16PurUpEpcR16
	PurUp5gcR16           *PurParametersNbR16PurUp5gcR16
	PurNrsrpValidationR16 *PurParametersNbR16PurNrsrpValidationR16
	PurCpL1ackR16         *PurParametersNbR16PurCpL1ackR16
}
