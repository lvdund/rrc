package ies

// SL-V2X-ConfigDedicated-r14 ::= SEQUENCE
// Extensible
type SlV2xConfigdedicatedR14 struct {
	CommtxresourcesR14          *SlV2xConfigdedicatedR14CommtxresourcesR14
	V2xInterfreqinfolistR14     *SlInterfreqinfolistv2xR14
	ThresslTxprioritizationR14  *SlPriorityR13
	TypetxsyncR14               *SlTypetxsyncR14
	CbrDedicatedtxconfiglistR14 *SlCbrCommontxconfiglistR14
}
