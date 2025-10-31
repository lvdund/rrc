package ies

// PDCP-Parameters-v1610 ::= SEQUENCE
type PdcpParametersV1610 struct {
	PdcpVersionchangewithouthoR16 *PdcpParametersV1610PdcpVersionchangewithouthoR16
	EhcR16                        *PdcpParametersV1610EhcR16
	ContinueehcContextR16         *PdcpParametersV1610ContinueehcContextR16
	MaxnumberehcContextsR16       *PdcpParametersV1610MaxnumberehcContextsR16
	JointehcRohcConfigR16         *PdcpParametersV1610JointehcRohcConfigR16
}
