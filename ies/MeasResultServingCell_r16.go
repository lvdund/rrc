package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

type MeasResultServingCell_r16 struct {
	ResultsSSB_Cell MeasQuantityResults                   `madatory`
	ResultsSSB      *MeasResultServingCell_r16_resultsSSB `lb:1,ub:maxNrofSSBs_r16,optional`
}

func (ie *MeasResultServingCell_r16) Encode(w *aper.AperWriter) error {
	var err error
	preambleBits := []bool{ie.ResultsSSB != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if err = ie.ResultsSSB_Cell.Encode(w); err != nil {
		return utils.WrapError("Encode ResultsSSB_Cell", err)
	}
	if ie.ResultsSSB != nil {
		if err = ie.ResultsSSB.Encode(w); err != nil {
			return utils.WrapError("Encode ResultsSSB", err)
		}
	}
	return nil
}

func (ie *MeasResultServingCell_r16) Decode(r *aper.AperReader) error {
	var err error
	var ResultsSSBPresent bool
	if ResultsSSBPresent, err = r.ReadBool(); err != nil {
		return err
	}
	if err = ie.ResultsSSB_Cell.Decode(r); err != nil {
		return utils.WrapError("Decode ResultsSSB_Cell", err)
	}
	if ResultsSSBPresent {
		ie.ResultsSSB = new(MeasResultServingCell_r16_resultsSSB)
		if err = ie.ResultsSSB.Decode(r); err != nil {
			return utils.WrapError("Decode ResultsSSB", err)
		}
	}
	return nil
}
