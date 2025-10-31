package ies

// DL-DCCH-MessageType-NB-c1 ::= CHOICE
const (
	DlDcchMessagetypeNbC1ChoiceNothing = iota
	DlDcchMessagetypeNbC1ChoiceDlinformationtransferR13
	DlDcchMessagetypeNbC1ChoiceRrcconnectionreconfigurationR13
	DlDcchMessagetypeNbC1ChoiceRrcconnectionreleaseR13
	DlDcchMessagetypeNbC1ChoiceSecuritymodecommandR13
	DlDcchMessagetypeNbC1ChoiceUecapabilityenquiryR13
	DlDcchMessagetypeNbC1ChoiceRrcconnectionresumeR13
	DlDcchMessagetypeNbC1ChoiceUeinformationrequestR16
	DlDcchMessagetypeNbC1ChoiceSpare1
)

type DlDcchMessagetypeNbC1 struct {
	Choice                          uint64
	DlinformationtransferR13        *DlinformationtransferNb
	RrcconnectionreconfigurationR13 *RrcconnectionreconfigurationNb
	RrcconnectionreleaseR13         *RrcconnectionreleaseNb
	SecuritymodecommandR13          *Securitymodecommand
	UecapabilityenquiryR13          *UecapabilityenquiryNb
	RrcconnectionresumeR13          *RrcconnectionresumeNb
	UeinformationrequestR16         *UeinformationrequestNbR16
	Spare1                          *struct{}
}
