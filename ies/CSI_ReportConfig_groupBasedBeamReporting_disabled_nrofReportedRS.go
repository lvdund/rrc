package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

const (
	CSI_ReportConfig_groupBasedBeamReporting_disabled_nrofReportedRS_Enum_n1 aper.Enumerated = 0
	CSI_ReportConfig_groupBasedBeamReporting_disabled_nrofReportedRS_Enum_n2 aper.Enumerated = 1
	CSI_ReportConfig_groupBasedBeamReporting_disabled_nrofReportedRS_Enum_n3 aper.Enumerated = 2
	CSI_ReportConfig_groupBasedBeamReporting_disabled_nrofReportedRS_Enum_n4 aper.Enumerated = 3
)

type CSI_ReportConfig_groupBasedBeamReporting_disabled_nrofReportedRS struct {
	Value aper.Enumerated `lb:0,ub:3,madatory`
}

func (ie *CSI_ReportConfig_groupBasedBeamReporting_disabled_nrofReportedRS) Encode(w *aper.AperWriter) error {
	var err error
	if err = w.WriteEnumerate(uint64(ie.Value), aper.Constraint{Lb: 0, Ub: 3}, false); err != nil {
		return utils.WrapError("Encode CSI_ReportConfig_groupBasedBeamReporting_disabled_nrofReportedRS", err)
	}
	return nil
}

func (ie *CSI_ReportConfig_groupBasedBeamReporting_disabled_nrofReportedRS) Decode(r *aper.AperReader) error {
	var err error
	var v uint64
	if v, err = r.ReadEnumerate(aper.Constraint{Lb: 0, Ub: 3}, false); err != nil {
		return utils.WrapError("Decode CSI_ReportConfig_groupBasedBeamReporting_disabled_nrofReportedRS", err)
	}
	ie.Value = aper.Enumerated(v)
	return nil
}
