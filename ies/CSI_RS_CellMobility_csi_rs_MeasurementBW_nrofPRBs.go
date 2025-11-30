package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

const (
	CSI_RS_CellMobility_csi_rs_MeasurementBW_nrofPRBs_Enum_size24  aper.Enumerated = 0
	CSI_RS_CellMobility_csi_rs_MeasurementBW_nrofPRBs_Enum_size48  aper.Enumerated = 1
	CSI_RS_CellMobility_csi_rs_MeasurementBW_nrofPRBs_Enum_size96  aper.Enumerated = 2
	CSI_RS_CellMobility_csi_rs_MeasurementBW_nrofPRBs_Enum_size192 aper.Enumerated = 3
	CSI_RS_CellMobility_csi_rs_MeasurementBW_nrofPRBs_Enum_size264 aper.Enumerated = 4
)

type CSI_RS_CellMobility_csi_rs_MeasurementBW_nrofPRBs struct {
	Value aper.Enumerated `lb:0,ub:4,madatory`
}

func (ie *CSI_RS_CellMobility_csi_rs_MeasurementBW_nrofPRBs) Encode(w *aper.AperWriter) error {
	var err error
	if err = w.WriteEnumerate(uint64(ie.Value), aper.Constraint{Lb: 0, Ub: 4}, false); err != nil {
		return utils.WrapError("Encode CSI_RS_CellMobility_csi_rs_MeasurementBW_nrofPRBs", err)
	}
	return nil
}

func (ie *CSI_RS_CellMobility_csi_rs_MeasurementBW_nrofPRBs) Decode(r *aper.AperReader) error {
	var err error
	var v uint64
	if v, err = r.ReadEnumerate(aper.Constraint{Lb: 0, Ub: 4}, false); err != nil {
		return utils.WrapError("Decode CSI_RS_CellMobility_csi_rs_MeasurementBW_nrofPRBs", err)
	}
	ie.Value = aper.Enumerated(v)
	return nil
}
