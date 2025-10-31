package ies

import "rrc/utils"

// PUCCH-Group-Config-r17-fr1-FR1-NonSharedTDD-r17 ::= ENUMERATED
type PucchGroupConfigR17Fr1Fr1NonsharedtddR17 struct {
	Value utils.ENUMERATED
}

const (
	PucchGroupConfigR17Fr1Fr1NonsharedtddR17EnumeratedNothing = iota
	PucchGroupConfigR17Fr1Fr1NonsharedtddR17EnumeratedSupported
)
