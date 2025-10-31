package ies

// ReestabNCellInfoList ::= SEQUENCE (SIZE (1..maxCellPrep)) OF ReestabNCellInfo
type ReestabNCellInfoList struct {
	Value []ReestabNCellInfo `lb:1,ub:maxCellPrep`
}
