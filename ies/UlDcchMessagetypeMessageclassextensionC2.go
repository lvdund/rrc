package ies

// UL-DCCH-MessageType-messageClassExtension-c2 ::= CHOICE
const (
	UlDcchMessagetypeMessageclassextensionC2ChoiceNothing = iota
	UlDcchMessagetypeMessageclassextensionC2ChoiceUldedicatedmessagesegmentR16
	UlDcchMessagetypeMessageclassextensionC2ChoiceDedicatedsibrequestR16
	UlDcchMessagetypeMessageclassextensionC2ChoiceMcgfailureinformationR16
	UlDcchMessagetypeMessageclassextensionC2ChoiceUeinformationresponseR16
	UlDcchMessagetypeMessageclassextensionC2ChoiceSidelinkueinformationnrR16
	UlDcchMessagetypeMessageclassextensionC2ChoiceUlinformationtransferiratR16
	UlDcchMessagetypeMessageclassextensionC2ChoiceIabotherinformationR16
	UlDcchMessagetypeMessageclassextensionC2ChoiceMbsinterestindicationR17
	UlDcchMessagetypeMessageclassextensionC2ChoiceUepositioningassistanceinfoR17
	UlDcchMessagetypeMessageclassextensionC2ChoiceMeasurementreportapplayerR17
	UlDcchMessagetypeMessageclassextensionC2ChoiceSpare6
	UlDcchMessagetypeMessageclassextensionC2ChoiceSpare5
	UlDcchMessagetypeMessageclassextensionC2ChoiceSpare4
	UlDcchMessagetypeMessageclassextensionC2ChoiceSpare3
	UlDcchMessagetypeMessageclassextensionC2ChoiceSpare2
	UlDcchMessagetypeMessageclassextensionC2ChoiceSpare1
)

type UlDcchMessagetypeMessageclassextensionC2 struct {
	Choice                         uint64
	UldedicatedmessagesegmentR16   *UldedicatedmessagesegmentR16
	DedicatedsibrequestR16         *DedicatedsibrequestR16
	McgfailureinformationR16       *McgfailureinformationR16
	UeinformationresponseR16       *UeinformationresponseR16
	SidelinkueinformationnrR16     *SidelinkueinformationnrR16
	UlinformationtransferiratR16   *UlinformationtransferiratR16
	IabotherinformationR16         *IabotherinformationR16
	MbsinterestindicationR17       *MbsinterestindicationR17
	UepositioningassistanceinfoR17 *UepositioningassistanceinfoR17
	MeasurementreportapplayerR17   *MeasurementreportapplayerR17
	Spare6                         *struct{}
	Spare5                         *struct{}
	Spare4                         *struct{}
	Spare3                         *struct{}
	Spare2                         *struct{}
	Spare1                         *struct{}
}
