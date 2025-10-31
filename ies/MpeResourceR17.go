package ies

// MPE-Resource-r17 ::= SEQUENCE
type MpeResourceR17 struct {
	MpeResourceidR17      MpeResourceidR17
	CellR17               *Servcellindex
	AdditionalpciR17      *AdditionalpciindexR17
	MpeReferencesignalR17 MpeResourceR17MpeReferencesignalR17
}
