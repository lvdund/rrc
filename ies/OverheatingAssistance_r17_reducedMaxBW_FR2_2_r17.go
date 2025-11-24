package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type OverheatingAssistance_r17_reducedMaxBW_FR2_2_r17 struct {
	ReducedBW_FR2_2_DL_r17 ReducedAggregatedBandwidth_r17 `madatory`
	ReducedBW_FR2_2_UL_r17 ReducedAggregatedBandwidth_r17 `madatory`
}

func (ie *OverheatingAssistance_r17_reducedMaxBW_FR2_2_r17) Encode(w *uper.UperWriter) error {
	var err error
	if err = ie.ReducedBW_FR2_2_DL_r17.Encode(w); err != nil {
		return utils.WrapError("Encode ReducedBW_FR2_2_DL_r17", err)
	}
	if err = ie.ReducedBW_FR2_2_UL_r17.Encode(w); err != nil {
		return utils.WrapError("Encode ReducedBW_FR2_2_UL_r17", err)
	}
	return nil
}

func (ie *OverheatingAssistance_r17_reducedMaxBW_FR2_2_r17) Decode(r *uper.UperReader) error {
	var err error
	if err = ie.ReducedBW_FR2_2_DL_r17.Decode(r); err != nil {
		return utils.WrapError("Decode ReducedBW_FR2_2_DL_r17", err)
	}
	if err = ie.ReducedBW_FR2_2_UL_r17.Decode(r); err != nil {
		return utils.WrapError("Decode ReducedBW_FR2_2_UL_r17", err)
	}
	return nil
}
