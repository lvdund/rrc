package ies

// MBSFNAreaConfiguration-v1610-IEs ::= SEQUENCE
type MbsfnareaconfigurationV1610 struct {
	CommonsfAllocV1610   *CommonsfAllocpatternlistV1610
	Noncriticalextension *MbsfnareaconfigurationV1610IesNoncriticalextension
}
