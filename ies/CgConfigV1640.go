package ies

// CG-Config-v1640-IEs ::= SEQUENCE
type CgConfigV1640 struct {
	ServcellinfolistscgNrR16    *ServcellinfolistscgNrR16
	ServcellinfolistscgEutraR16 *ServcellinfolistscgEutraR16
	Noncriticalextension        *CgConfigV1700
}
