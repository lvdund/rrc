package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

type BandParametersSidelinkEUTRA_NR_r16_nr struct {
	BandParametersSidelinkNR_r16 BandParametersSidelink_r16 `madatory`
}

func (ie *BandParametersSidelinkEUTRA_NR_r16_nr) Encode(w *aper.AperWriter) error {
	var err error
	if err = ie.BandParametersSidelinkNR_r16.Encode(w); err != nil {
		return utils.WrapError("Encode BandParametersSidelinkNR_r16", err)
	}
	return nil
}

func (ie *BandParametersSidelinkEUTRA_NR_r16_nr) Decode(r *aper.AperReader) error {
	var err error
	if err = ie.BandParametersSidelinkNR_r16.Decode(r); err != nil {
		return utils.WrapError("Decode BandParametersSidelinkNR_r16", err)
	}
	return nil
}
