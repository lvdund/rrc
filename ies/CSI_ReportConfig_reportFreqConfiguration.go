package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type CSI_ReportConfig_reportFreqConfiguration struct {
	Cqi_FormatIndicator *CSI_ReportConfig_reportFreqConfiguration_cqi_FormatIndicator `optional`
	Pmi_FormatIndicator *CSI_ReportConfig_reportFreqConfiguration_pmi_FormatIndicator `optional`
	Csi_ReportingBand   *CSI_ReportConfig_reportFreqConfiguration_csi_ReportingBand   `lb:3,ub:3,optional`
}

func (ie *CSI_ReportConfig_reportFreqConfiguration) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.Cqi_FormatIndicator != nil, ie.Pmi_FormatIndicator != nil, ie.Csi_ReportingBand != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.Cqi_FormatIndicator != nil {
		if err = ie.Cqi_FormatIndicator.Encode(w); err != nil {
			return utils.WrapError("Encode Cqi_FormatIndicator", err)
		}
	}
	if ie.Pmi_FormatIndicator != nil {
		if err = ie.Pmi_FormatIndicator.Encode(w); err != nil {
			return utils.WrapError("Encode Pmi_FormatIndicator", err)
		}
	}
	if ie.Csi_ReportingBand != nil {
		if err = ie.Csi_ReportingBand.Encode(w); err != nil {
			return utils.WrapError("Encode Csi_ReportingBand", err)
		}
	}
	return nil
}

func (ie *CSI_ReportConfig_reportFreqConfiguration) Decode(r *uper.UperReader) error {
	var err error
	var Cqi_FormatIndicatorPresent bool
	if Cqi_FormatIndicatorPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var Pmi_FormatIndicatorPresent bool
	if Pmi_FormatIndicatorPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var Csi_ReportingBandPresent bool
	if Csi_ReportingBandPresent, err = r.ReadBool(); err != nil {
		return err
	}
	if Cqi_FormatIndicatorPresent {
		ie.Cqi_FormatIndicator = new(CSI_ReportConfig_reportFreqConfiguration_cqi_FormatIndicator)
		if err = ie.Cqi_FormatIndicator.Decode(r); err != nil {
			return utils.WrapError("Decode Cqi_FormatIndicator", err)
		}
	}
	if Pmi_FormatIndicatorPresent {
		ie.Pmi_FormatIndicator = new(CSI_ReportConfig_reportFreqConfiguration_pmi_FormatIndicator)
		if err = ie.Pmi_FormatIndicator.Decode(r); err != nil {
			return utils.WrapError("Decode Pmi_FormatIndicator", err)
		}
	}
	if Csi_ReportingBandPresent {
		ie.Csi_ReportingBand = new(CSI_ReportConfig_reportFreqConfiguration_csi_ReportingBand)
		if err = ie.Csi_ReportingBand.Decode(r); err != nil {
			return utils.WrapError("Decode Csi_ReportingBand", err)
		}
	}
	return nil
}
