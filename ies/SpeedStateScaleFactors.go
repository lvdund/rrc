package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type SpeedStateScaleFactors struct {
	Sf_Medium SpeedStateScaleFactors_sf_Medium `madatory`
	Sf_High   SpeedStateScaleFactors_sf_High   `madatory`
}

func (ie *SpeedStateScaleFactors) Encode(w *uper.UperWriter) error {
	var err error
	if err = ie.Sf_Medium.Encode(w); err != nil {
		return utils.WrapError("Encode Sf_Medium", err)
	}
	if err = ie.Sf_High.Encode(w); err != nil {
		return utils.WrapError("Encode Sf_High", err)
	}
	return nil
}

func (ie *SpeedStateScaleFactors) Decode(r *uper.UperReader) error {
	var err error
	if err = ie.Sf_Medium.Decode(r); err != nil {
		return utils.WrapError("Decode Sf_Medium", err)
	}
	if err = ie.Sf_High.Decode(r); err != nil {
		return utils.WrapError("Decode Sf_High", err)
	}
	return nil
}
