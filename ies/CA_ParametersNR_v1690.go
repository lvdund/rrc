package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type CA_ParametersNR_v1690 struct {
	Csi_ReportingCrossPUCCH_Grp_r16 *CA_ParametersNR_v1690_csi_ReportingCrossPUCCH_Grp_r16 `lb:1,ub:maxCarrierTypePairList_r16,optional`
}

func (ie *CA_ParametersNR_v1690) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.Csi_ReportingCrossPUCCH_Grp_r16 != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.Csi_ReportingCrossPUCCH_Grp_r16 != nil {
		if err = ie.Csi_ReportingCrossPUCCH_Grp_r16.Encode(w); err != nil {
			return utils.WrapError("Encode Csi_ReportingCrossPUCCH_Grp_r16", err)
		}
	}
	return nil
}

func (ie *CA_ParametersNR_v1690) Decode(r *uper.UperReader) error {
	var err error
	var Csi_ReportingCrossPUCCH_Grp_r16Present bool
	if Csi_ReportingCrossPUCCH_Grp_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	if Csi_ReportingCrossPUCCH_Grp_r16Present {
		ie.Csi_ReportingCrossPUCCH_Grp_r16 = new(CA_ParametersNR_v1690_csi_ReportingCrossPUCCH_Grp_r16)
		if err = ie.Csi_ReportingCrossPUCCH_Grp_r16.Decode(r); err != nil {
			return utils.WrapError("Decode Csi_ReportingCrossPUCCH_Grp_r16", err)
		}
	}
	return nil
}
