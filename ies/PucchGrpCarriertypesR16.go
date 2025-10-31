package ies

// PUCCH-Grp-CarrierTypes-r16 ::= SEQUENCE
type PucchGrpCarriertypesR16 struct {
	Fr1NonsharedtddR16 *PucchGrpCarriertypesR16Fr1NonsharedtddR16
	Fr1SharedtddR16    *PucchGrpCarriertypesR16Fr1SharedtddR16
	Fr1NonsharedfddR16 *PucchGrpCarriertypesR16Fr1NonsharedfddR16
	Fr2R16             *PucchGrpCarriertypesR16Fr2R16
}
