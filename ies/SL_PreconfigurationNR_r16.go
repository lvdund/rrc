package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

type SL_PreconfigurationNR_r16 struct {
	SidelinkPreconfigNR_r16 SidelinkPreconfigNR_r16 `madatory`
}

func (ie *SL_PreconfigurationNR_r16) Encode(w *aper.AperWriter) error {
	var err error
	if err = ie.SidelinkPreconfigNR_r16.Encode(w); err != nil {
		return utils.WrapError("Encode SidelinkPreconfigNR_r16", err)
	}
	return nil
}

func (ie *SL_PreconfigurationNR_r16) Decode(r *aper.AperReader) error {
	var err error
	if err = ie.SidelinkPreconfigNR_r16.Decode(r); err != nil {
		return utils.WrapError("Decode SidelinkPreconfigNR_r16", err)
	}
	return nil
}
