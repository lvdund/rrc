package ies

import "rrc/utils"

// PUCCH-Group-Config-r17-fr2-FR2-NonSharedTDD-r17 ::= ENUMERATED
type PucchGroupConfigR17Fr2Fr2NonsharedtddR17 struct {
	Value utils.ENUMERATED
}

const (
	PucchGroupConfigR17Fr2Fr2NonsharedtddR17EnumeratedNothing = iota
	PucchGroupConfigR17Fr2Fr2NonsharedtddR17EnumeratedSupported
)
