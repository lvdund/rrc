package ies

// MasterInformationBlock-TDD-NB-r15-operationModeInfo-r15 ::= CHOICE
const (
	MasterinformationblockTddNbR15OperationmodeinfoR15ChoiceNothing = iota
	MasterinformationblockTddNbR15OperationmodeinfoR15ChoiceInbandSamepciR15
	MasterinformationblockTddNbR15OperationmodeinfoR15ChoiceInbandDifferentpciR15
	MasterinformationblockTddNbR15OperationmodeinfoR15ChoiceGuardbandR15
	MasterinformationblockTddNbR15OperationmodeinfoR15ChoiceStandaloneR15
)

type MasterinformationblockTddNbR15OperationmodeinfoR15 struct {
	Choice                uint64
	InbandSamepciR15      *InbandSamepciTddNbR15
	InbandDifferentpciR15 *InbandDifferentpciTddNbR15
	GuardbandR15          *GuardbandtddNbR15
	StandaloneR15         *StandalonetddNbR15
}
