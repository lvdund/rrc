package ies

import "rrc/utils"

// OtherConfig-v1540 ::= SEQUENCE
// Extensible
type OtherconfigV1540 struct {
	Overheatingassistanceconfig *utils.Setuprelease[Overheatingassistanceconfig]
}
