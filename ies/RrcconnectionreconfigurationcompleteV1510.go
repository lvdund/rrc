package ies

import "rrc/utils"

// RRCConnectionReconfigurationComplete-v1510-IEs ::= SEQUENCE
type RrcconnectionreconfigurationcompleteV1510 struct {
	ScgConfigresponsenrR15 *utils.OCTETSTRING
	Noncriticalextension   *RrcconnectionreconfigurationcompleteV1530
}
