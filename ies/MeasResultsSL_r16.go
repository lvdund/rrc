package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

type MeasResultsSL_r16 struct {
	MeasResultsListSL_r16 MeasResultsSL_r16_measResultsListSL_r16 `madatory`
}

func (ie *MeasResultsSL_r16) Encode(w *aper.AperWriter) error {
	var err error
	if err = ie.MeasResultsListSL_r16.Encode(w); err != nil {
		return utils.WrapError("Encode MeasResultsListSL_r16", err)
	}
	return nil
}

func (ie *MeasResultsSL_r16) Decode(r *aper.AperReader) error {
	var err error
	if err = ie.MeasResultsListSL_r16.Decode(r); err != nil {
		return utils.WrapError("Decode MeasResultsListSL_r16", err)
	}
	return nil
}
