package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

type MeasResultFailedCell_r16_measResult_r16 struct {
	CellResults_r16    MeasResultFailedCell_r16_measResult_r16_cellResults_r16    `madatory`
	RsIndexResults_r16 MeasResultFailedCell_r16_measResult_r16_rsIndexResults_r16 `madatory`
}

func (ie *MeasResultFailedCell_r16_measResult_r16) Encode(w *aper.AperWriter) error {
	var err error
	if err = ie.CellResults_r16.Encode(w); err != nil {
		return utils.WrapError("Encode CellResults_r16", err)
	}
	if err = ie.RsIndexResults_r16.Encode(w); err != nil {
		return utils.WrapError("Encode RsIndexResults_r16", err)
	}
	return nil
}

func (ie *MeasResultFailedCell_r16_measResult_r16) Decode(r *aper.AperReader) error {
	var err error
	if err = ie.CellResults_r16.Decode(r); err != nil {
		return utils.WrapError("Decode CellResults_r16", err)
	}
	if err = ie.RsIndexResults_r16.Decode(r); err != nil {
		return utils.WrapError("Decode RsIndexResults_r16", err)
	}
	return nil
}
