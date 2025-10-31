package ies

// PUCCH-ConfigDedicated-r13-pucch-Format-r13-format4-r13 ::= SEQUENCE
type PucchConfigdedicatedR13PucchFormatR13Format4R13 struct {
	Format4ResourceconfigurationR13         []Format4ResourceR13  `lb:4,ub:4`
	Format4MulticsiResourceconfigurationR13 *[]Format4ResourceR13 `lb:1,ub:2`
}
