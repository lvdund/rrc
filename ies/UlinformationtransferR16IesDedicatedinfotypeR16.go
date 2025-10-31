package ies

// ULInformationTransfer-r16-IEs-dedicatedInfoType-r16 ::= CHOICE
const (
	UlinformationtransferR16IesDedicatedinfotypeR16ChoiceNothing = iota
	UlinformationtransferR16IesDedicatedinfotypeR16ChoiceDedicatedinfonasR16
	UlinformationtransferR16IesDedicatedinfotypeR16ChoiceDedicatedinfocdma20001xrttR16
	UlinformationtransferR16IesDedicatedinfotypeR16ChoiceDedicatedinfocdma2000HrpdR16
)

type UlinformationtransferR16IesDedicatedinfotypeR16 struct {
	Choice                        uint64
	DedicatedinfonasR16           *Dedicatedinfonas
	Dedicatedinfocdma20001xrttR16 *Dedicatedinfocdma2000
	Dedicatedinfocdma2000HrpdR16  *Dedicatedinfocdma2000
}
