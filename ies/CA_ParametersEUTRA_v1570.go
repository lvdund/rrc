package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

type CA_ParametersEUTRA_v1570 struct {
	Dl_1024QAM_TotalWeightedLayers *int64 `lb:0,ub:10,optional`
}

func (ie *CA_ParametersEUTRA_v1570) Encode(w *aper.AperWriter) error {
	var err error
	preambleBits := []bool{ie.Dl_1024QAM_TotalWeightedLayers != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.Dl_1024QAM_TotalWeightedLayers != nil {
		if err = w.WriteInteger(*ie.Dl_1024QAM_TotalWeightedLayers, &aper.Constraint{Lb: 0, Ub: 10}, false); err != nil {
			return utils.WrapError("Encode Dl_1024QAM_TotalWeightedLayers", err)
		}
	}
	return nil
}

func (ie *CA_ParametersEUTRA_v1570) Decode(r *aper.AperReader) error {
	var err error
	var Dl_1024QAM_TotalWeightedLayersPresent bool
	if Dl_1024QAM_TotalWeightedLayersPresent, err = r.ReadBool(); err != nil {
		return err
	}
	if Dl_1024QAM_TotalWeightedLayersPresent {
		var tmp_int_Dl_1024QAM_TotalWeightedLayers int64
		if tmp_int_Dl_1024QAM_TotalWeightedLayers, err = r.ReadInteger(&aper.Constraint{Lb: 0, Ub: 10}, false); err != nil {
			return utils.WrapError("Decode Dl_1024QAM_TotalWeightedLayers", err)
		}
		ie.Dl_1024QAM_TotalWeightedLayers = &tmp_int_Dl_1024QAM_TotalWeightedLayers
	}
	return nil
}
