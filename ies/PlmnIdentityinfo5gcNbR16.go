package ies

// PLMN-IdentityInfo-5GC-NB-r16 ::= SEQUENCE
type PlmnIdentityinfo5gcNbR16 struct {
	PlmnIdentity5gcR16            PlmnIdentityinfo5gcNbR16PlmnIdentity5gcR16
	CellreservedforoperatoruseR16 PlmnIdentityinfo5gcNbR16CellreservedforoperatoruseR16
	NgUDatatransferR16            *bool
	UpCiot5gsOptimisationR16      *bool
}
