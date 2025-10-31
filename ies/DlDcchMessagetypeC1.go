package ies

// DL-DCCH-MessageType-c1 ::= CHOICE
const (
	DlDcchMessagetypeC1ChoiceNothing = iota
	DlDcchMessagetypeC1ChoiceRrcreconfiguration
	DlDcchMessagetypeC1ChoiceRrcresume
	DlDcchMessagetypeC1ChoiceRrcrelease
	DlDcchMessagetypeC1ChoiceRrcreestablishment
	DlDcchMessagetypeC1ChoiceSecuritymodecommand
	DlDcchMessagetypeC1ChoiceDlinformationtransfer
	DlDcchMessagetypeC1ChoiceUecapabilityenquiry
	DlDcchMessagetypeC1ChoiceCountercheck
	DlDcchMessagetypeC1ChoiceMobilityfromnrcommand
	DlDcchMessagetypeC1ChoiceDldedicatedmessagesegmentR16
	DlDcchMessagetypeC1ChoiceUeinformationrequestR16
	DlDcchMessagetypeC1ChoiceDlinformationtransfermrdcR16
	DlDcchMessagetypeC1ChoiceLoggedmeasurementconfigurationR16
	DlDcchMessagetypeC1ChoiceSpare3
	DlDcchMessagetypeC1ChoiceSpare2
	DlDcchMessagetypeC1ChoiceSpare1
)

type DlDcchMessagetypeC1 struct {
	Choice                            uint64
	Rrcreconfiguration                *Rrcreconfiguration
	Rrcresume                         *Rrcresume
	Rrcrelease                        *Rrcrelease
	Rrcreestablishment                *Rrcreestablishment
	Securitymodecommand               *Securitymodecommand
	Dlinformationtransfer             *Dlinformationtransfer
	Uecapabilityenquiry               *Uecapabilityenquiry
	Countercheck                      *Countercheck
	Mobilityfromnrcommand             *Mobilityfromnrcommand
	DldedicatedmessagesegmentR16      *DldedicatedmessagesegmentR16
	UeinformationrequestR16           *UeinformationrequestR16
	DlinformationtransfermrdcR16      *DlinformationtransfermrdcR16
	LoggedmeasurementconfigurationR16 *LoggedmeasurementconfigurationR16
	Spare3                            *struct{}
	Spare2                            *struct{}
	Spare1                            *struct{}
}
