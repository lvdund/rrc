package ies

// UL-DCCH-MessageType-c1 ::= CHOICE
const (
	UlDcchMessagetypeC1ChoiceNothing = iota
	UlDcchMessagetypeC1ChoiceCsfbparametersrequestcdma2000
	UlDcchMessagetypeC1ChoiceMeasurementreport
	UlDcchMessagetypeC1ChoiceRrcconnectionreconfigurationcomplete
	UlDcchMessagetypeC1ChoiceRrcconnectionreestablishmentcomplete
	UlDcchMessagetypeC1ChoiceRrcconnectionsetupcomplete
	UlDcchMessagetypeC1ChoiceSecuritymodecomplete
	UlDcchMessagetypeC1ChoiceSecuritymodefailure
	UlDcchMessagetypeC1ChoiceUecapabilityinformation
	UlDcchMessagetypeC1ChoiceUlhandoverpreparationtransfer
	UlDcchMessagetypeC1ChoiceUlinformationtransfer
	UlDcchMessagetypeC1ChoiceCountercheckresponse
	UlDcchMessagetypeC1ChoiceUeinformationresponseR9
	UlDcchMessagetypeC1ChoiceProximityindicationR9
	UlDcchMessagetypeC1ChoiceRnreconfigurationcompleteR10
	UlDcchMessagetypeC1ChoiceMbmscountingresponseR10
	UlDcchMessagetypeC1ChoiceInterfreqrstdmeasurementindicationR10
)

type UlDcchMessagetypeC1 struct {
	Choice                                uint64
	Csfbparametersrequestcdma2000         *Csfbparametersrequestcdma2000
	Measurementreport                     *Measurementreport
	Rrcconnectionreconfigurationcomplete  *Rrcconnectionreconfigurationcomplete
	Rrcconnectionreestablishmentcomplete  *Rrcconnectionreestablishmentcomplete
	Rrcconnectionsetupcomplete            *Rrcconnectionsetupcomplete
	Securitymodecomplete                  *Securitymodecomplete
	Securitymodefailure                   *Securitymodefailure
	Uecapabilityinformation               *Uecapabilityinformation
	Ulhandoverpreparationtransfer         *Ulhandoverpreparationtransfer
	Ulinformationtransfer                 *Ulinformationtransfer
	Countercheckresponse                  *Countercheckresponse
	UeinformationresponseR9               *UeinformationresponseR9
	ProximityindicationR9                 *ProximityindicationR9
	RnreconfigurationcompleteR10          *RnreconfigurationcompleteR10
	MbmscountingresponseR10               *MbmscountingresponseR10
	InterfreqrstdmeasurementindicationR10 *InterfreqrstdmeasurementindicationR10
}
