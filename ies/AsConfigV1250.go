package ies

import "rrc/utils"

// AS-Config-v1250 ::= SEQUENCE
type AsConfigV1250 struct {
	SourcewlanOffloadconfigR12 *WlanOffloadconfigR12
	SourceslCommconfigR12      *SlCommconfigR12
	SourceslDiscconfigR12      *SlDiscconfigR12
}
