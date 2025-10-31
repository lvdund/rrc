package ies

// HandoverFromEUTRAPreparationRequest-v1020-IEs ::= SEQUENCE
type HandoverfromeutrapreparationrequestV1020 struct {
	DualrxtxredirectindicatorR10    *bool
	Redirectcarriercdma20001xrttR10 *Carrierfreqcdma2000
	Noncriticalextension            *HandoverfromeutrapreparationrequestV1020IesNoncriticalextension
}
