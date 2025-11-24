package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type MeasResultNR_measResult_cellResults struct {
	ResultsSSB_Cell    *MeasQuantityResults `optional`
	ResultsCSI_RS_Cell *MeasQuantityResults `optional`
}

func (ie *MeasResultNR_measResult_cellResults) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.ResultsSSB_Cell != nil, ie.ResultsCSI_RS_Cell != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.ResultsSSB_Cell != nil {
		if err = ie.ResultsSSB_Cell.Encode(w); err != nil {
			return utils.WrapError("Encode ResultsSSB_Cell", err)
		}
	}
	if ie.ResultsCSI_RS_Cell != nil {
		if err = ie.ResultsCSI_RS_Cell.Encode(w); err != nil {
			return utils.WrapError("Encode ResultsCSI_RS_Cell", err)
		}
	}
	return nil
}

func (ie *MeasResultNR_measResult_cellResults) Decode(r *uper.UperReader) error {
	var err error
	var ResultsSSB_CellPresent bool
	if ResultsSSB_CellPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var ResultsCSI_RS_CellPresent bool
	if ResultsCSI_RS_CellPresent, err = r.ReadBool(); err != nil {
		return err
	}
	if ResultsSSB_CellPresent {
		ie.ResultsSSB_Cell = new(MeasQuantityResults)
		if err = ie.ResultsSSB_Cell.Decode(r); err != nil {
			return utils.WrapError("Decode ResultsSSB_Cell", err)
		}
	}
	if ResultsCSI_RS_CellPresent {
		ie.ResultsCSI_RS_Cell = new(MeasQuantityResults)
		if err = ie.ResultsCSI_RS_Cell.Decode(r); err != nil {
			return utils.WrapError("Decode ResultsCSI_RS_Cell", err)
		}
	}
	return nil
}
