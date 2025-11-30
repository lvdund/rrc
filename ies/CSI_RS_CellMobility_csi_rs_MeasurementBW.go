package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

type CSI_RS_CellMobility_csi_rs_MeasurementBW struct {
	NrofPRBs CSI_RS_CellMobility_csi_rs_MeasurementBW_nrofPRBs `madatory`
	StartPRB int64                                             `lb:0,ub:2169,madatory`
}

func (ie *CSI_RS_CellMobility_csi_rs_MeasurementBW) Encode(w *aper.AperWriter) error {
	var err error
	if err = ie.NrofPRBs.Encode(w); err != nil {
		return utils.WrapError("Encode NrofPRBs", err)
	}
	if err = w.WriteInteger(ie.StartPRB, &aper.Constraint{Lb: 0, Ub: 2169}, false); err != nil {
		return utils.WrapError("WriteInteger StartPRB", err)
	}
	return nil
}

func (ie *CSI_RS_CellMobility_csi_rs_MeasurementBW) Decode(r *aper.AperReader) error {
	var err error
	if err = ie.NrofPRBs.Decode(r); err != nil {
		return utils.WrapError("Decode NrofPRBs", err)
	}
	var tmp_int_StartPRB int64
	if tmp_int_StartPRB, err = r.ReadInteger(&aper.Constraint{Lb: 0, Ub: 2169}, false); err != nil {
		return utils.WrapError("ReadInteger StartPRB", err)
	}
	ie.StartPRB = tmp_int_StartPRB
	return nil
}
