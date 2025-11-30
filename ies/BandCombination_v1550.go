package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

type BandCombination_v1550 struct {
	Ca_ParametersNR_v1550 CA_ParametersNR_v1550 `madatory`
}

func (ie *BandCombination_v1550) Encode(w *aper.AperWriter) error {
	var err error
	if err = ie.Ca_ParametersNR_v1550.Encode(w); err != nil {
		return utils.WrapError("Encode Ca_ParametersNR_v1550", err)
	}
	return nil
}

func (ie *BandCombination_v1550) Decode(r *aper.AperReader) error {
	var err error
	if err = ie.Ca_ParametersNR_v1550.Decode(r); err != nil {
		return utils.WrapError("Decode Ca_ParametersNR_v1550", err)
	}
	return nil
}
