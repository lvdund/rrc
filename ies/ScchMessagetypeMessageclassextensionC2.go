package ies

// SCCH-MessageType-messageClassExtension-c2 ::= CHOICE
const (
	ScchMessagetypeMessageclassextensionC2ChoiceNothing = iota
	ScchMessagetypeMessageclassextensionC2ChoiceNotificationmessagesidelinkR17
	ScchMessagetypeMessageclassextensionC2ChoiceUeassistanceinformationsidelinkR17
	ScchMessagetypeMessageclassextensionC2ChoiceSpare6
	ScchMessagetypeMessageclassextensionC2ChoiceSpare5
	ScchMessagetypeMessageclassextensionC2ChoiceSpare4
	ScchMessagetypeMessageclassextensionC2ChoiceSpare3
	ScchMessagetypeMessageclassextensionC2ChoiceSpare2
	ScchMessagetypeMessageclassextensionC2ChoiceSpare1
)

type ScchMessagetypeMessageclassextensionC2 struct {
	Choice                             uint64
	NotificationmessagesidelinkR17     *NotificationmessagesidelinkR17
	UeassistanceinformationsidelinkR17 *UeassistanceinformationsidelinkR17
	Spare6                             *struct{}
	Spare5                             *struct{}
	Spare4                             *struct{}
	Spare3                             *struct{}
	Spare2                             *struct{}
	Spare1                             *struct{}
}
