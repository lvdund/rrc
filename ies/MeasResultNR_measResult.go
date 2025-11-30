package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

type MeasResultNR_measResult struct {
	CellResults    *MeasResultNR_measResult_cellResults    `optional`
	RsIndexResults *MeasResultNR_measResult_rsIndexResults `optional`
}

func (ie *MeasResultNR_measResult) Encode(w *aper.AperWriter) error {
	var err error
	preambleBits := []bool{ie.CellResults != nil, ie.RsIndexResults != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.CellResults != nil {
		if err = ie.CellResults.Encode(w); err != nil {
			return utils.WrapError("Encode CellResults", err)
		}
	}
	if ie.RsIndexResults != nil {
		if err = ie.RsIndexResults.Encode(w); err != nil {
			return utils.WrapError("Encode RsIndexResults", err)
		}
	}
	return nil
}

func (ie *MeasResultNR_measResult) Decode(r *aper.AperReader) error {
	var err error
	var CellResultsPresent bool
	if CellResultsPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var RsIndexResultsPresent bool
	if RsIndexResultsPresent, err = r.ReadBool(); err != nil {
		return err
	}
	if CellResultsPresent {
		ie.CellResults = new(MeasResultNR_measResult_cellResults)
		if err = ie.CellResults.Decode(r); err != nil {
			return utils.WrapError("Decode CellResults", err)
		}
	}
	if RsIndexResultsPresent {
		ie.RsIndexResults = new(MeasResultNR_measResult_rsIndexResults)
		if err = ie.RsIndexResults.Decode(r); err != nil {
			return utils.WrapError("Decode RsIndexResults", err)
		}
	}
	return nil
}
