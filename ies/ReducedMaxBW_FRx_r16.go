package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

type ReducedMaxBW_FRx_r16 struct {
	ReducedBW_DL_r16 ReducedAggregatedBandwidth `madatory`
	ReducedBW_UL_r16 ReducedAggregatedBandwidth `madatory`
}

func (ie *ReducedMaxBW_FRx_r16) Encode(w *aper.AperWriter) error {
	var err error
	if err = ie.ReducedBW_DL_r16.Encode(w); err != nil {
		return utils.WrapError("Encode ReducedBW_DL_r16", err)
	}
	if err = ie.ReducedBW_UL_r16.Encode(w); err != nil {
		return utils.WrapError("Encode ReducedBW_UL_r16", err)
	}
	return nil
}

func (ie *ReducedMaxBW_FRx_r16) Decode(r *aper.AperReader) error {
	var err error
	if err = ie.ReducedBW_DL_r16.Decode(r); err != nil {
		return utils.WrapError("Decode ReducedBW_DL_r16", err)
	}
	if err = ie.ReducedBW_UL_r16.Decode(r); err != nil {
		return utils.WrapError("Decode ReducedBW_UL_r16", err)
	}
	return nil
}
