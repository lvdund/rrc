package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

type CSI_ReportConfig_groupBasedBeamReporting_disabled struct {
	NrofReportedRS *CSI_ReportConfig_groupBasedBeamReporting_disabled_nrofReportedRS `optional`
}

func (ie *CSI_ReportConfig_groupBasedBeamReporting_disabled) Encode(w *aper.AperWriter) error {
	var err error
	preambleBits := []bool{ie.NrofReportedRS != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.NrofReportedRS != nil {
		if err = ie.NrofReportedRS.Encode(w); err != nil {
			return utils.WrapError("Encode NrofReportedRS", err)
		}
	}
	return nil
}

func (ie *CSI_ReportConfig_groupBasedBeamReporting_disabled) Decode(r *aper.AperReader) error {
	var err error
	var NrofReportedRSPresent bool
	if NrofReportedRSPresent, err = r.ReadBool(); err != nil {
		return err
	}
	if NrofReportedRSPresent {
		ie.NrofReportedRS = new(CSI_ReportConfig_groupBasedBeamReporting_disabled_nrofReportedRS)
		if err = ie.NrofReportedRS.Decode(r); err != nil {
			return utils.WrapError("Decode NrofReportedRS", err)
		}
	}
	return nil
}
