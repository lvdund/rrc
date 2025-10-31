package ies

import "rrc/utils"

// RadioBearerConfig ::= SEQUENCE
// Extensible
type Radiobearerconfig struct {
	SrbToaddmodlist     *SrbToaddmodlist
	Srb3Torelease       *utils.BOOLEAN
	DrbToaddmodlist     *DrbToaddmodlist
	DrbToreleaselist    *DrbToreleaselist
	Securityconfig      *Securityconfig
	MrbToaddmodlistR17  *MrbToaddmodlistR17  `ext`
	MrbToreleaselistR17 *MrbToreleaselistR17 `ext`
	Srb4ToaddmodR17     *SrbToaddmod         `ext`
	Srb4ToreleaseR17    *utils.BOOLEAN       `ext`
}
