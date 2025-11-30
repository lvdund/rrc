package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

type CA_ParametersEUTRA_v1560 struct {
	Fd_MIMO_TotalWeightedLayers *int64 `lb:2,ub:128,optional`
}

func (ie *CA_ParametersEUTRA_v1560) Encode(w *aper.AperWriter) error {
	var err error
	preambleBits := []bool{ie.Fd_MIMO_TotalWeightedLayers != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.Fd_MIMO_TotalWeightedLayers != nil {
		if err = w.WriteInteger(*ie.Fd_MIMO_TotalWeightedLayers, &aper.Constraint{Lb: 2, Ub: 128}, false); err != nil {
			return utils.WrapError("Encode Fd_MIMO_TotalWeightedLayers", err)
		}
	}
	return nil
}

func (ie *CA_ParametersEUTRA_v1560) Decode(r *aper.AperReader) error {
	var err error
	var Fd_MIMO_TotalWeightedLayersPresent bool
	if Fd_MIMO_TotalWeightedLayersPresent, err = r.ReadBool(); err != nil {
		return err
	}
	if Fd_MIMO_TotalWeightedLayersPresent {
		var tmp_int_Fd_MIMO_TotalWeightedLayers int64
		if tmp_int_Fd_MIMO_TotalWeightedLayers, err = r.ReadInteger(&aper.Constraint{Lb: 2, Ub: 128}, false); err != nil {
			return utils.WrapError("Decode Fd_MIMO_TotalWeightedLayers", err)
		}
		ie.Fd_MIMO_TotalWeightedLayers = &tmp_int_Fd_MIMO_TotalWeightedLayers
	}
	return nil
}
