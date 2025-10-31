package ies

// RRCConnectionReconfiguration-v1510-IEs-nr-Config-r15 ::= CHOICE
const (
	RrcconnectionreconfigurationV1510IesNrConfigR15ChoiceNothing = iota
	RrcconnectionreconfigurationV1510IesNrConfigR15ChoiceRelease
	RrcconnectionreconfigurationV1510IesNrConfigR15ChoiceSetup
)

type RrcconnectionreconfigurationV1510IesNrConfigR15 struct {
	Choice  uint64
	Release *struct{}
	Setup   *RrcconnectionreconfigurationV1510IesNrConfigR15Setup
}
