package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

type CSI_ReportConfig_semiPersistentOnPUSCH_v1530 struct {
	ReportSlotConfig_v1530 CSI_ReportConfig_semiPersistentOnPUSCH_v1530_reportSlotConfig_v1530 `madatory`
}

func (ie *CSI_ReportConfig_semiPersistentOnPUSCH_v1530) Encode(w *aper.AperWriter) error {
	var err error
	if err = ie.ReportSlotConfig_v1530.Encode(w); err != nil {
		return utils.WrapError("Encode ReportSlotConfig_v1530", err)
	}
	return nil
}

func (ie *CSI_ReportConfig_semiPersistentOnPUSCH_v1530) Decode(r *aper.AperReader) error {
	var err error
	if err = ie.ReportSlotConfig_v1530.Decode(r); err != nil {
		return utils.WrapError("Decode ReportSlotConfig_v1530", err)
	}
	return nil
}
