package ies

// RACH-CE-LevelInfo-r13 ::= SEQUENCE
// Extensible
type RachCeLevelinfoR13 struct {
	PreamblemappinginfoR13          RachCeLevelinfoR13PreamblemappinginfoR13
	RaResponsewindowsizeR13         RachCeLevelinfoR13RaResponsewindowsizeR13
	MacContentionresolutiontimerR13 RachCeLevelinfoR13MacContentionresolutiontimerR13
	RarHoppingconfigR13             RachCeLevelinfoR13RarHoppingconfigR13
}
