package ies

// MBMS-Parameters-v1470 ::= SEQUENCE
type MbmsParametersV1470 struct {
	MbmsMaxbwR14               MbmsParametersV1470MbmsMaxbwR14
	MbmsScalingfactor1dot25R14 *MbmsParametersV1470MbmsScalingfactor1dot25R14
	MbmsScalingfactor7dot5R14  *MbmsParametersV1470MbmsScalingfactor7dot5R14
}
