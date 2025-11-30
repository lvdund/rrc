package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

const (
	CSI_ReportConfig_semiPersistentOnPUSCH_v1530_reportSlotConfig_v1530_Enum_sl4  aper.Enumerated = 0
	CSI_ReportConfig_semiPersistentOnPUSCH_v1530_reportSlotConfig_v1530_Enum_sl8  aper.Enumerated = 1
	CSI_ReportConfig_semiPersistentOnPUSCH_v1530_reportSlotConfig_v1530_Enum_sl16 aper.Enumerated = 2
)

type CSI_ReportConfig_semiPersistentOnPUSCH_v1530_reportSlotConfig_v1530 struct {
	Value aper.Enumerated `lb:0,ub:2,madatory`
}

func (ie *CSI_ReportConfig_semiPersistentOnPUSCH_v1530_reportSlotConfig_v1530) Encode(w *aper.AperWriter) error {
	var err error
	if err = w.WriteEnumerate(uint64(ie.Value), aper.Constraint{Lb: 0, Ub: 2}, false); err != nil {
		return utils.WrapError("Encode CSI_ReportConfig_semiPersistentOnPUSCH_v1530_reportSlotConfig_v1530", err)
	}
	return nil
}

func (ie *CSI_ReportConfig_semiPersistentOnPUSCH_v1530_reportSlotConfig_v1530) Decode(r *aper.AperReader) error {
	var err error
	var v uint64
	if v, err = r.ReadEnumerate(aper.Constraint{Lb: 0, Ub: 2}, false); err != nil {
		return utils.WrapError("Decode CSI_ReportConfig_semiPersistentOnPUSCH_v1530_reportSlotConfig_v1530", err)
	}
	ie.Value = aper.Enumerated(v)
	return nil
}
