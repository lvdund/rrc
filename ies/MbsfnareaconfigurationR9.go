package ies

// MBSFNAreaConfiguration-r9 ::= SEQUENCE
type MbsfnareaconfigurationR9 struct {
	CommonsfAllocR9       CommonsfAllocpatternlistR9
	CommonsfAllocperiodR9 MbsfnareaconfigurationR9CommonsfAllocperiodR9
	PmchInfolistR9        PmchInfolistR9
	Noncriticalextension  *MbsfnareaconfigurationV930
}
