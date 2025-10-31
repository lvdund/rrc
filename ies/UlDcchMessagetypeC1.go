package ies

// UL-DCCH-MessageType-c1 ::= CHOICE
const (
	UlDcchMessagetypeC1ChoiceNothing = iota
	UlDcchMessagetypeC1ChoiceMeasurementreport
	UlDcchMessagetypeC1ChoiceRrcreconfigurationcomplete
	UlDcchMessagetypeC1ChoiceRrcsetupcomplete
	UlDcchMessagetypeC1ChoiceRrcreestablishmentcomplete
	UlDcchMessagetypeC1ChoiceRrcresumecomplete
	UlDcchMessagetypeC1ChoiceSecuritymodecomplete
	UlDcchMessagetypeC1ChoiceSecuritymodefailure
	UlDcchMessagetypeC1ChoiceUlinformationtransfer
	UlDcchMessagetypeC1ChoiceLocationmeasurementindication
	UlDcchMessagetypeC1ChoiceUecapabilityinformation
	UlDcchMessagetypeC1ChoiceCountercheckresponse
	UlDcchMessagetypeC1ChoiceUeassistanceinformation
	UlDcchMessagetypeC1ChoiceFailureinformation
	UlDcchMessagetypeC1ChoiceUlinformationtransfermrdc
	UlDcchMessagetypeC1ChoiceScgfailureinformation
	UlDcchMessagetypeC1ChoiceScgfailureinformationeutra
)

type UlDcchMessagetypeC1 struct {
	Choice                        uint64
	Measurementreport             *Measurementreport
	Rrcreconfigurationcomplete    *Rrcreconfigurationcomplete
	Rrcsetupcomplete              *Rrcsetupcomplete
	Rrcreestablishmentcomplete    *Rrcreestablishmentcomplete
	Rrcresumecomplete             *Rrcresumecomplete
	Securitymodecomplete          *Securitymodecomplete
	Securitymodefailure           *Securitymodefailure
	Ulinformationtransfer         *Ulinformationtransfer
	Locationmeasurementindication *Locationmeasurementindication
	Uecapabilityinformation       *Uecapabilityinformation
	Countercheckresponse          *Countercheckresponse
	Ueassistanceinformation       *Ueassistanceinformation
	Failureinformation            *Failureinformation
	Ulinformationtransfermrdc     *Ulinformationtransfermrdc
	Scgfailureinformation         *Scgfailureinformation
	Scgfailureinformationeutra    *Scgfailureinformationeutra
}
