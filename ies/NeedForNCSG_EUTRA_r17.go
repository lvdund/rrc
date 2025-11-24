package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type NeedForNCSG_EUTRA_r17 struct {
	BandEUTRA_r17     FreqBandIndicatorEUTRA                  `madatory`
	GapIndication_r17 NeedForNCSG_EUTRA_r17_gapIndication_r17 `madatory`
}

func (ie *NeedForNCSG_EUTRA_r17) Encode(w *uper.UperWriter) error {
	var err error
	if err = ie.BandEUTRA_r17.Encode(w); err != nil {
		return utils.WrapError("Encode BandEUTRA_r17", err)
	}
	if err = ie.GapIndication_r17.Encode(w); err != nil {
		return utils.WrapError("Encode GapIndication_r17", err)
	}
	return nil
}

func (ie *NeedForNCSG_EUTRA_r17) Decode(r *uper.UperReader) error {
	var err error
	if err = ie.BandEUTRA_r17.Decode(r); err != nil {
		return utils.WrapError("Decode BandEUTRA_r17", err)
	}
	if err = ie.GapIndication_r17.Decode(r); err != nil {
		return utils.WrapError("Decode GapIndication_r17", err)
	}
	return nil
}
