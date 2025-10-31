package ies

// RegisteredAMF ::= SEQUENCE
type Registeredamf struct {
	PlmnIdentity  *PlmnIdentity
	AmfIdentifier AmfIdentifier
}
