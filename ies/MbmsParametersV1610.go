package ies

// MBMS-Parameters-v1610 ::= SEQUENCE
type MbmsParametersV1610 struct {
	MbmsScalingfactor2dot5R16    *MbmsParametersV1610MbmsScalingfactor2dot5R16
	MbmsScalingfactor0dot37R16   *MbmsParametersV1610MbmsScalingfactor0dot37R16
	MbmsSupportedbandinfolistR16 []MbmsSupportedbandinfoR16 `lb:1,ub:maxBands`
}
