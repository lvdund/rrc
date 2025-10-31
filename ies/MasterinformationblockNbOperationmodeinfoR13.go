package ies

// MasterInformationBlock-NB-operationModeInfo-r13 ::= CHOICE
const (
	MasterinformationblockNbOperationmodeinfoR13ChoiceNothing = iota
	MasterinformationblockNbOperationmodeinfoR13ChoiceInbandSamepciR13
	MasterinformationblockNbOperationmodeinfoR13ChoiceInbandDifferentpciR13
	MasterinformationblockNbOperationmodeinfoR13ChoiceGuardbandR13
	MasterinformationblockNbOperationmodeinfoR13ChoiceStandaloneR13
)

type MasterinformationblockNbOperationmodeinfoR13 struct {
	Choice                uint64
	InbandSamepciR13      *InbandSamepciNbR13
	InbandDifferentpciR13 *InbandDifferentpciNbR13
	GuardbandR13          *GuardbandNbR13
	StandaloneR13         *StandaloneNbR13
}
