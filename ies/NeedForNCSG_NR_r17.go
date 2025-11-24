package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type NeedForNCSG_NR_r17 struct {
	BandNR_r17        FreqBandIndicatorNR                  `madatory`
	GapIndication_r17 NeedForNCSG_NR_r17_gapIndication_r17 `madatory`
}

func (ie *NeedForNCSG_NR_r17) Encode(w *uper.UperWriter) error {
	var err error
	if err = ie.BandNR_r17.Encode(w); err != nil {
		return utils.WrapError("Encode BandNR_r17", err)
	}
	if err = ie.GapIndication_r17.Encode(w); err != nil {
		return utils.WrapError("Encode GapIndication_r17", err)
	}
	return nil
}

func (ie *NeedForNCSG_NR_r17) Decode(r *uper.UperReader) error {
	var err error
	if err = ie.BandNR_r17.Decode(r); err != nil {
		return utils.WrapError("Decode BandNR_r17", err)
	}
	if err = ie.GapIndication_r17.Decode(r); err != nil {
		return utils.WrapError("Decode GapIndication_r17", err)
	}
	return nil
}
