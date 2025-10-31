package ies

// PosSI-SchedulingInfo-r16 ::= SEQUENCE
// Extensible
type PossiSchedulinginfoR16 struct {
	PosschedulinginfolistR16    []PosschedulinginfoR16 `lb:1,ub:maxSIMessage`
	PossiRequestconfigR16       *SiRequestconfig
	PossiRequestconfigsulR16    *SiRequestconfig
	PossiRequestconfigredcapR17 *SiRequestconfig `ext`
}
