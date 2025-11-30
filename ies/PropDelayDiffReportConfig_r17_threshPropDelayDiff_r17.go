package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

const (
	PropDelayDiffReportConfig_r17_threshPropDelayDiff_r17_Enum_ms0dot5 aper.Enumerated = 0
	PropDelayDiffReportConfig_r17_threshPropDelayDiff_r17_Enum_ms1     aper.Enumerated = 1
	PropDelayDiffReportConfig_r17_threshPropDelayDiff_r17_Enum_ms2     aper.Enumerated = 2
	PropDelayDiffReportConfig_r17_threshPropDelayDiff_r17_Enum_ms3     aper.Enumerated = 3
	PropDelayDiffReportConfig_r17_threshPropDelayDiff_r17_Enum_ms4     aper.Enumerated = 4
	PropDelayDiffReportConfig_r17_threshPropDelayDiff_r17_Enum_ms5     aper.Enumerated = 5
	PropDelayDiffReportConfig_r17_threshPropDelayDiff_r17_Enum_ms6     aper.Enumerated = 6
	PropDelayDiffReportConfig_r17_threshPropDelayDiff_r17_Enum_ms7     aper.Enumerated = 7
	PropDelayDiffReportConfig_r17_threshPropDelayDiff_r17_Enum_ms8     aper.Enumerated = 8
	PropDelayDiffReportConfig_r17_threshPropDelayDiff_r17_Enum_ms9     aper.Enumerated = 9
	PropDelayDiffReportConfig_r17_threshPropDelayDiff_r17_Enum_ms10    aper.Enumerated = 10
	PropDelayDiffReportConfig_r17_threshPropDelayDiff_r17_Enum_spare5  aper.Enumerated = 11
	PropDelayDiffReportConfig_r17_threshPropDelayDiff_r17_Enum_spare4  aper.Enumerated = 12
	PropDelayDiffReportConfig_r17_threshPropDelayDiff_r17_Enum_spare3  aper.Enumerated = 13
	PropDelayDiffReportConfig_r17_threshPropDelayDiff_r17_Enum_spare2  aper.Enumerated = 14
	PropDelayDiffReportConfig_r17_threshPropDelayDiff_r17_Enum_spare1  aper.Enumerated = 15
)

type PropDelayDiffReportConfig_r17_threshPropDelayDiff_r17 struct {
	Value aper.Enumerated `lb:0,ub:15,madatory`
}

func (ie *PropDelayDiffReportConfig_r17_threshPropDelayDiff_r17) Encode(w *aper.AperWriter) error {
	var err error
	if err = w.WriteEnumerate(uint64(ie.Value), aper.Constraint{Lb: 0, Ub: 15}, false); err != nil {
		return utils.WrapError("Encode PropDelayDiffReportConfig_r17_threshPropDelayDiff_r17", err)
	}
	return nil
}

func (ie *PropDelayDiffReportConfig_r17_threshPropDelayDiff_r17) Decode(r *aper.AperReader) error {
	var err error
	var v uint64
	if v, err = r.ReadEnumerate(aper.Constraint{Lb: 0, Ub: 15}, false); err != nil {
		return utils.WrapError("Decode PropDelayDiffReportConfig_r17_threshPropDelayDiff_r17", err)
	}
	ie.Value = aper.Enumerated(v)
	return nil
}
