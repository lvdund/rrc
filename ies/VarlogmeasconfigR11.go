package ies

// VarLogMeasConfig-r11 ::= SEQUENCE
type VarlogmeasconfigR11 struct {
	AreaconfigurationR10   *AreaconfigurationR10
	AreaconfigurationV1130 *AreaconfigurationV1130
	LoggingdurationR10     LoggingdurationR10
	LoggingintervalR10     LoggingintervalR10
}
