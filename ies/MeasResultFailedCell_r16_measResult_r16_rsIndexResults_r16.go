package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

type MeasResultFailedCell_r16_measResult_r16_rsIndexResults_r16 struct {
	ResultsSSB_Indexes_r16 ResultsPerSSB_IndexList `madatory`
}

func (ie *MeasResultFailedCell_r16_measResult_r16_rsIndexResults_r16) Encode(w *aper.AperWriter) error {
	var err error
	if err = ie.ResultsSSB_Indexes_r16.Encode(w); err != nil {
		return utils.WrapError("Encode ResultsSSB_Indexes_r16", err)
	}
	return nil
}

func (ie *MeasResultFailedCell_r16_measResult_r16_rsIndexResults_r16) Decode(r *aper.AperReader) error {
	var err error
	if err = ie.ResultsSSB_Indexes_r16.Decode(r); err != nil {
		return utils.WrapError("Decode ResultsSSB_Indexes_r16", err)
	}
	return nil
}
