package ies

// DLInformationTransfer-r8-IEs-dedicatedInfoType ::= CHOICE
const (
	DlinformationtransferR8IesDedicatedinfotypeChoiceNothing = iota
	DlinformationtransferR8IesDedicatedinfotypeChoiceDedicatedinfonas
	DlinformationtransferR8IesDedicatedinfotypeChoiceDedicatedinfocdma20001xrtt
	DlinformationtransferR8IesDedicatedinfotypeChoiceDedicatedinfocdma2000Hrpd
)

type DlinformationtransferR8IesDedicatedinfotype struct {
	Choice                     uint64
	Dedicatedinfonas           *Dedicatedinfonas
	Dedicatedinfocdma20001xrtt *Dedicatedinfocdma2000
	Dedicatedinfocdma2000Hrpd  *Dedicatedinfocdma2000
}
