package ies

// DL-DCCH-MessageType-c1 ::= CHOICE
const (
	DlDcchMessagetypeC1ChoiceNothing = iota
	DlDcchMessagetypeC1ChoiceCsfbparametersresponsecdma2000
	DlDcchMessagetypeC1ChoiceDlinformationtransfer
	DlDcchMessagetypeC1ChoiceHandoverfromeutrapreparationrequest
	DlDcchMessagetypeC1ChoiceMobilityfromeutracommand
	DlDcchMessagetypeC1ChoiceRrcconnectionreconfiguration
	DlDcchMessagetypeC1ChoiceRrcconnectionrelease
	DlDcchMessagetypeC1ChoiceSecuritymodecommand
	DlDcchMessagetypeC1ChoiceUecapabilityenquiry
	DlDcchMessagetypeC1ChoiceCountercheck
	DlDcchMessagetypeC1ChoiceUeinformationrequestR9
	DlDcchMessagetypeC1ChoiceLoggedmeasurementconfigurationR10
	DlDcchMessagetypeC1ChoiceRnreconfigurationR10
	DlDcchMessagetypeC1ChoiceRrcconnectionresumeR13
	DlDcchMessagetypeC1ChoiceDldedicatedmessagesegmentR16
	DlDcchMessagetypeC1ChoiceSpare2
	DlDcchMessagetypeC1ChoiceSpare1
)

type DlDcchMessagetypeC1 struct {
	Choice                              uint64
	Csfbparametersresponsecdma2000      *Csfbparametersresponsecdma2000
	Dlinformationtransfer               *Dlinformationtransfer
	Handoverfromeutrapreparationrequest *Handoverfromeutrapreparationrequest
	Mobilityfromeutracommand            *Mobilityfromeutracommand
	Rrcconnectionreconfiguration        *Rrcconnectionreconfiguration
	Rrcconnectionrelease                *Rrcconnectionrelease
	Securitymodecommand                 *Securitymodecommand
	Uecapabilityenquiry                 *Uecapabilityenquiry
	Countercheck                        *Countercheck
	UeinformationrequestR9              *UeinformationrequestR9
	LoggedmeasurementconfigurationR10   *LoggedmeasurementconfigurationR10
	RnreconfigurationR10                *RnreconfigurationR10
	RrcconnectionresumeR13              *RrcconnectionresumeR13
	DldedicatedmessagesegmentR16        *DldedicatedmessagesegmentR16
	Spare2                              *struct{}
	Spare1                              *struct{}
}
