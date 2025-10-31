package ies

// ULInformationTransfer-r8-IEs-dedicatedInfoType ::= CHOICE
const (
	UlinformationtransferR8IesDedicatedinfotypeChoiceNothing = iota
	UlinformationtransferR8IesDedicatedinfotypeChoiceDedicatedinfonas
	UlinformationtransferR8IesDedicatedinfotypeChoiceDedicatedinfocdma20001xrtt
	UlinformationtransferR8IesDedicatedinfotypeChoiceDedicatedinfocdma2000Hrpd
)

type UlinformationtransferR8IesDedicatedinfotype struct {
	Choice                     uint64
	Dedicatedinfonas           *Dedicatedinfonas
	Dedicatedinfocdma20001xrtt *Dedicatedinfocdma2000
	Dedicatedinfocdma2000Hrpd  *Dedicatedinfocdma2000
}
