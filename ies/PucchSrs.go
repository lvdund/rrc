package ies

// PUCCH-SRS ::= SEQUENCE
type PucchSrs struct {
	Resource  SrsResourceid
	Uplinkbwp BwpId
}
