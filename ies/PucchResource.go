package ies

// PUCCH-Resource ::= SEQUENCE
type PucchResource struct {
	PucchResourceid           PucchResourceid
	Startingprb               PrbId
	Intraslotfrequencyhopping *PucchResourceIntraslotfrequencyhopping
	Secondhopprb              *PrbId
	Format                    PucchResourceFormat
}
