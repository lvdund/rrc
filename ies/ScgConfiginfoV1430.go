package ies

// SCG-ConfigInfo-v1430-IEs ::= SEQUENCE
type ScgConfiginfoV1430 struct {
	MakebeforebreakscgReqR14 *bool
	MeasgapconfigperccList   *MeasgapconfigperccListR14
	Noncriticalextension     *ScgConfiginfoV1530
}
