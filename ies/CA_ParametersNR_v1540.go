package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

type CA_ParametersNR_v1540 struct {
	SimultaneousSRS_AssocCSI_RS_AllCC         *int64                                                           `lb:5,ub:32,optional`
	Csi_RS_IM_ReceptionForFeedbackPerBandComb *CA_ParametersNR_v1540_csi_RS_IM_ReceptionForFeedbackPerBandComb `lb:1,ub:64,optional`
	SimultaneousCSI_ReportsAllCC              *int64                                                           `lb:5,ub:32,optional`
	DualPA_Architecture                       *CA_ParametersNR_v1540_dualPA_Architecture                       `optional`
}

func (ie *CA_ParametersNR_v1540) Encode(w *aper.AperWriter) error {
	var err error
	preambleBits := []bool{ie.SimultaneousSRS_AssocCSI_RS_AllCC != nil, ie.Csi_RS_IM_ReceptionForFeedbackPerBandComb != nil, ie.SimultaneousCSI_ReportsAllCC != nil, ie.DualPA_Architecture != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.SimultaneousSRS_AssocCSI_RS_AllCC != nil {
		if err = w.WriteInteger(*ie.SimultaneousSRS_AssocCSI_RS_AllCC, &aper.Constraint{Lb: 5, Ub: 32}, false); err != nil {
			return utils.WrapError("Encode SimultaneousSRS_AssocCSI_RS_AllCC", err)
		}
	}
	if ie.Csi_RS_IM_ReceptionForFeedbackPerBandComb != nil {
		if err = ie.Csi_RS_IM_ReceptionForFeedbackPerBandComb.Encode(w); err != nil {
			return utils.WrapError("Encode Csi_RS_IM_ReceptionForFeedbackPerBandComb", err)
		}
	}
	if ie.SimultaneousCSI_ReportsAllCC != nil {
		if err = w.WriteInteger(*ie.SimultaneousCSI_ReportsAllCC, &aper.Constraint{Lb: 5, Ub: 32}, false); err != nil {
			return utils.WrapError("Encode SimultaneousCSI_ReportsAllCC", err)
		}
	}
	if ie.DualPA_Architecture != nil {
		if err = ie.DualPA_Architecture.Encode(w); err != nil {
			return utils.WrapError("Encode DualPA_Architecture", err)
		}
	}
	return nil
}

func (ie *CA_ParametersNR_v1540) Decode(r *aper.AperReader) error {
	var err error
	var SimultaneousSRS_AssocCSI_RS_AllCCPresent bool
	if SimultaneousSRS_AssocCSI_RS_AllCCPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var Csi_RS_IM_ReceptionForFeedbackPerBandCombPresent bool
	if Csi_RS_IM_ReceptionForFeedbackPerBandCombPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var SimultaneousCSI_ReportsAllCCPresent bool
	if SimultaneousCSI_ReportsAllCCPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var DualPA_ArchitecturePresent bool
	if DualPA_ArchitecturePresent, err = r.ReadBool(); err != nil {
		return err
	}
	if SimultaneousSRS_AssocCSI_RS_AllCCPresent {
		var tmp_int_SimultaneousSRS_AssocCSI_RS_AllCC int64
		if tmp_int_SimultaneousSRS_AssocCSI_RS_AllCC, err = r.ReadInteger(&aper.Constraint{Lb: 5, Ub: 32}, false); err != nil {
			return utils.WrapError("Decode SimultaneousSRS_AssocCSI_RS_AllCC", err)
		}
		ie.SimultaneousSRS_AssocCSI_RS_AllCC = &tmp_int_SimultaneousSRS_AssocCSI_RS_AllCC
	}
	if Csi_RS_IM_ReceptionForFeedbackPerBandCombPresent {
		ie.Csi_RS_IM_ReceptionForFeedbackPerBandComb = new(CA_ParametersNR_v1540_csi_RS_IM_ReceptionForFeedbackPerBandComb)
		if err = ie.Csi_RS_IM_ReceptionForFeedbackPerBandComb.Decode(r); err != nil {
			return utils.WrapError("Decode Csi_RS_IM_ReceptionForFeedbackPerBandComb", err)
		}
	}
	if SimultaneousCSI_ReportsAllCCPresent {
		var tmp_int_SimultaneousCSI_ReportsAllCC int64
		if tmp_int_SimultaneousCSI_ReportsAllCC, err = r.ReadInteger(&aper.Constraint{Lb: 5, Ub: 32}, false); err != nil {
			return utils.WrapError("Decode SimultaneousCSI_ReportsAllCC", err)
		}
		ie.SimultaneousCSI_ReportsAllCC = &tmp_int_SimultaneousCSI_ReportsAllCC
	}
	if DualPA_ArchitecturePresent {
		ie.DualPA_Architecture = new(CA_ParametersNR_v1540_dualPA_Architecture)
		if err = ie.DualPA_Architecture.Decode(r); err != nil {
			return utils.WrapError("Decode DualPA_Architecture", err)
		}
	}
	return nil
}
