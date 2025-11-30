package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

type CA_ParametersNR_v1540_csi_RS_IM_ReceptionForFeedbackPerBandComb struct {
	MaxNumberSimultaneousNZP_CSI_RS_ActBWP_AllCC        *int64 `lb:1,ub:64,optional`
	TotalNumberPortsSimultaneousNZP_CSI_RS_ActBWP_AllCC *int64 `lb:2,ub:256,optional`
}

func (ie *CA_ParametersNR_v1540_csi_RS_IM_ReceptionForFeedbackPerBandComb) Encode(w *aper.AperWriter) error {
	var err error
	preambleBits := []bool{ie.MaxNumberSimultaneousNZP_CSI_RS_ActBWP_AllCC != nil, ie.TotalNumberPortsSimultaneousNZP_CSI_RS_ActBWP_AllCC != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.MaxNumberSimultaneousNZP_CSI_RS_ActBWP_AllCC != nil {
		if err = w.WriteInteger(*ie.MaxNumberSimultaneousNZP_CSI_RS_ActBWP_AllCC, &aper.Constraint{Lb: 1, Ub: 64}, false); err != nil {
			return utils.WrapError("Encode MaxNumberSimultaneousNZP_CSI_RS_ActBWP_AllCC", err)
		}
	}
	if ie.TotalNumberPortsSimultaneousNZP_CSI_RS_ActBWP_AllCC != nil {
		if err = w.WriteInteger(*ie.TotalNumberPortsSimultaneousNZP_CSI_RS_ActBWP_AllCC, &aper.Constraint{Lb: 2, Ub: 256}, false); err != nil {
			return utils.WrapError("Encode TotalNumberPortsSimultaneousNZP_CSI_RS_ActBWP_AllCC", err)
		}
	}
	return nil
}

func (ie *CA_ParametersNR_v1540_csi_RS_IM_ReceptionForFeedbackPerBandComb) Decode(r *aper.AperReader) error {
	var err error
	var MaxNumberSimultaneousNZP_CSI_RS_ActBWP_AllCCPresent bool
	if MaxNumberSimultaneousNZP_CSI_RS_ActBWP_AllCCPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var TotalNumberPortsSimultaneousNZP_CSI_RS_ActBWP_AllCCPresent bool
	if TotalNumberPortsSimultaneousNZP_CSI_RS_ActBWP_AllCCPresent, err = r.ReadBool(); err != nil {
		return err
	}
	if MaxNumberSimultaneousNZP_CSI_RS_ActBWP_AllCCPresent {
		var tmp_int_MaxNumberSimultaneousNZP_CSI_RS_ActBWP_AllCC int64
		if tmp_int_MaxNumberSimultaneousNZP_CSI_RS_ActBWP_AllCC, err = r.ReadInteger(&aper.Constraint{Lb: 1, Ub: 64}, false); err != nil {
			return utils.WrapError("Decode MaxNumberSimultaneousNZP_CSI_RS_ActBWP_AllCC", err)
		}
		ie.MaxNumberSimultaneousNZP_CSI_RS_ActBWP_AllCC = &tmp_int_MaxNumberSimultaneousNZP_CSI_RS_ActBWP_AllCC
	}
	if TotalNumberPortsSimultaneousNZP_CSI_RS_ActBWP_AllCCPresent {
		var tmp_int_TotalNumberPortsSimultaneousNZP_CSI_RS_ActBWP_AllCC int64
		if tmp_int_TotalNumberPortsSimultaneousNZP_CSI_RS_ActBWP_AllCC, err = r.ReadInteger(&aper.Constraint{Lb: 2, Ub: 256}, false); err != nil {
			return utils.WrapError("Decode TotalNumberPortsSimultaneousNZP_CSI_RS_ActBWP_AllCC", err)
		}
		ie.TotalNumberPortsSimultaneousNZP_CSI_RS_ActBWP_AllCC = &tmp_int_TotalNumberPortsSimultaneousNZP_CSI_RS_ActBWP_AllCC
	}
	return nil
}
