package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

type CSI_ReportConfig_groupBasedBeamReporting_v1710 struct {
	NrofReportedGroups_r17 CSI_ReportConfig_groupBasedBeamReporting_v1710_nrofReportedGroups_r17 `madatory`
}

func (ie *CSI_ReportConfig_groupBasedBeamReporting_v1710) Encode(w *aper.AperWriter) error {
	var err error
	if err = ie.NrofReportedGroups_r17.Encode(w); err != nil {
		return utils.WrapError("Encode NrofReportedGroups_r17", err)
	}
	return nil
}

func (ie *CSI_ReportConfig_groupBasedBeamReporting_v1710) Decode(r *aper.AperReader) error {
	var err error
	if err = ie.NrofReportedGroups_r17.Decode(r); err != nil {
		return utils.WrapError("Decode NrofReportedGroups_r17", err)
	}
	return nil
}
