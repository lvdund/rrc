package ies

import "rrc/utils"

// NRDC-Parameters-v1570-sfn-SyncNRDC ::= ENUMERATED
type NrdcParametersV1570SfnSyncnrdc struct {
	Value utils.ENUMERATED
}

const (
	NrdcParametersV1570SfnSyncnrdcEnumeratedNothing = iota
	NrdcParametersV1570SfnSyncnrdcEnumeratedSupported
)
