package ies

import "rrc/utils"

// PhyLayerParameters-v1530-stti-SPT-Capabilities-r15-srs-DCI7-TriggeringFS2-r15 ::= ENUMERATED
type PhylayerparametersV1530SttiSptCapabilitiesR15SrsDci7Triggeringfs2R15 struct {
	Value utils.ENUMERATED
}

const (
	PhylayerparametersV1530SttiSptCapabilitiesR15SrsDci7Triggeringfs2R15EnumeratedNothing = iota
	PhylayerparametersV1530SttiSptCapabilitiesR15SrsDci7Triggeringfs2R15EnumeratedSupported
)
