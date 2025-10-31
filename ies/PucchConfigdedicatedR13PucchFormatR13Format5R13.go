package ies

// PUCCH-ConfigDedicated-r13-pucch-Format-r13-format5-r13 ::= SEQUENCE
type PucchConfigdedicatedR13PucchFormatR13Format5R13 struct {
	Format5ResourceconfigurationR13         []Format5ResourceR13 `lb:4,ub:4`
	Format5MulticsiResourceconfigurationR13 *Format5ResourceR13
}
