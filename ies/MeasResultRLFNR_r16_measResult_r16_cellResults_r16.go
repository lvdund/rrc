package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type MeasResultRLFNR_r16_measResult_r16_cellResults_r16 struct {
	ResultsSSB_Cell_r16    *MeasQuantityResults `optional`
	ResultsCSI_RS_Cell_r16 *MeasQuantityResults `optional`
}

func (ie *MeasResultRLFNR_r16_measResult_r16_cellResults_r16) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.ResultsSSB_Cell_r16 != nil, ie.ResultsCSI_RS_Cell_r16 != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.ResultsSSB_Cell_r16 != nil {
		if err = ie.ResultsSSB_Cell_r16.Encode(w); err != nil {
			return utils.WrapError("Encode ResultsSSB_Cell_r16", err)
		}
	}
	if ie.ResultsCSI_RS_Cell_r16 != nil {
		if err = ie.ResultsCSI_RS_Cell_r16.Encode(w); err != nil {
			return utils.WrapError("Encode ResultsCSI_RS_Cell_r16", err)
		}
	}
	return nil
}

func (ie *MeasResultRLFNR_r16_measResult_r16_cellResults_r16) Decode(r *uper.UperReader) error {
	var err error
	var ResultsSSB_Cell_r16Present bool
	if ResultsSSB_Cell_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var ResultsCSI_RS_Cell_r16Present bool
	if ResultsCSI_RS_Cell_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	if ResultsSSB_Cell_r16Present {
		ie.ResultsSSB_Cell_r16 = new(MeasQuantityResults)
		if err = ie.ResultsSSB_Cell_r16.Decode(r); err != nil {
			return utils.WrapError("Decode ResultsSSB_Cell_r16", err)
		}
	}
	if ResultsCSI_RS_Cell_r16Present {
		ie.ResultsCSI_RS_Cell_r16 = new(MeasQuantityResults)
		if err = ie.ResultsCSI_RS_Cell_r16.Decode(r); err != nil {
			return utils.WrapError("Decode ResultsCSI_RS_Cell_r16", err)
		}
	}
	return nil
}
