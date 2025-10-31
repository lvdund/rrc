package ies

// SL-MeasConfigCommon-r16 ::= SEQUENCE
// Extensible
type SlMeasconfigcommonR16 struct {
	SlMeasobjectlistcommonR16   *SlMeasobjectlistR16
	SlReportconfiglistcommonR16 *SlReportconfiglistR16
	SlMeasidlistcommonR16       *SlMeasidlistR16
	SlQuantityconfigcommonR16   *SlQuantityconfigR16
}
