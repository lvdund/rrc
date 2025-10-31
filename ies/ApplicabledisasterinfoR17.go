package ies

// ApplicableDisasterInfo-r17 ::= CHOICE
const (
	ApplicabledisasterinfoR17ChoiceNothing = iota
	ApplicabledisasterinfoR17ChoiceNodisasterroamingR17
	ApplicabledisasterinfoR17ChoiceDisasterrelatedindicationR17
	ApplicabledisasterinfoR17ChoiceCommonplmnsR17
	ApplicabledisasterinfoR17ChoiceDedicatedplmnsR17
)

type ApplicabledisasterinfoR17 struct {
	Choice                       uint64
	NodisasterroamingR17         *struct{}
	DisasterrelatedindicationR17 *struct{}
	CommonplmnsR17               *struct{}
	DedicatedplmnsR17            *[]PlmnIdentity `lb:1,ub:maxPLMN`
}
