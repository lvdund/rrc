package ies

// SL-MeasConfig-r16 ::= SEQUENCE
// Extensible
type SlMeasconfigR16 struct {
	SlMeasobjecttoremovelistR16   *SlMeasobjecttoremovelistR16
	SlMeasobjecttoaddmodlistR16   *SlMeasobjectlistR16
	SlReportconfigtoremovelistR16 *SlReportconfigtoremovelistR16
	SlReportconfigtoaddmodlistR16 *SlReportconfiglistR16
	SlMeasidtoremovelistR16       *SlMeasidtoremovelistR16
	SlMeasidtoaddmodlistR16       *SlMeasidlistR16
	SlQuantityconfigR16           *SlQuantityconfigR16
}
