package ies

// SCCH-MessageType-c1 ::= CHOICE
const (
	ScchMessagetypeC1ChoiceNothing = iota
	ScchMessagetypeC1ChoiceMeasurementreportsidelink
	ScchMessagetypeC1ChoiceRrcreconfigurationsidelink
	ScchMessagetypeC1ChoiceRrcreconfigurationcompletesidelink
	ScchMessagetypeC1ChoiceRrcreconfigurationfailuresidelink
	ScchMessagetypeC1ChoiceUecapabilityenquirysidelink
	ScchMessagetypeC1ChoiceUecapabilityinformationsidelink
	ScchMessagetypeC1ChoiceUumessagetransfersidelinkR17
	ScchMessagetypeC1ChoiceRemoteueinformationsidelinkR17
)

type ScchMessagetypeC1 struct {
	Choice                             uint64
	Measurementreportsidelink          *Measurementreportsidelink
	Rrcreconfigurationsidelink         *Rrcreconfigurationsidelink
	Rrcreconfigurationcompletesidelink *Rrcreconfigurationcompletesidelink
	Rrcreconfigurationfailuresidelink  *Rrcreconfigurationfailuresidelink
	Uecapabilityenquirysidelink        *Uecapabilityenquirysidelink
	Uecapabilityinformationsidelink    *Uecapabilityinformationsidelink
	UumessagetransfersidelinkR17       *UumessagetransfersidelinkR17
	RemoteueinformationsidelinkR17     *RemoteueinformationsidelinkR17
}
