package ies

// TAG-Config ::= SEQUENCE
type TagConfig struct {
	TagToreleaselist *[]TagId `lb:1,ub:maxNrofTAGs`
	TagToaddmodlist  *[]Tag   `lb:1,ub:maxNrofTAGs`
}
