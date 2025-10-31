package ies

import "rrc/utils"

// SL-ConfigDedicatedNR-r16 ::= SEQUENCE
// Extensible
type SlConfigdedicatednrR16 struct {
	SlPhyMacRlcConfigR16             *SlPhyMacRlcConfigR16
	SlRadiobearertoreleaselistR16    *[]SlrbUuConfigindexR16   `lb:1,ub:maxNrofSLRBR16`
	SlRadiobearertoaddmodlistR16     *[]SlRadiobearerconfigR16 `lb:1,ub:maxNrofSLRBR16`
	SlMeasconfiginfotoreleaselistR16 *[]SlDestinationindexR16  `lb:1,ub:maxNrofSLDestR16`
	SlMeasconfiginfotoaddmodlistR16  *[]SlMeasconfiginfoR16    `lb:1,ub:maxNrofSLDestR16`
	T400R16                          *SlConfigdedicatednrR16T400R16
	SlPhyMacRlcConfigV1700           *utils.Setuprelease[SlPhyMacRlcConfigV1700] `ext`
	SlDiscconfigR17                  *utils.Setuprelease[SlDiscconfigR17]        `ext`
}
