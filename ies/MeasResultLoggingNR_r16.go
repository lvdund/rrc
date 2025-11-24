package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type MeasResultLoggingNR_r16 struct {
	PhysCellId_r16      PhysCellId          `madatory`
	ResultsSSB_Cell_r16 MeasQuantityResults `madatory`
	NumberOfGoodSSB_r16 *int64              `lb:1,ub:maxNrofSSBs_r16,optional`
}

func (ie *MeasResultLoggingNR_r16) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.NumberOfGoodSSB_r16 != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if err = ie.PhysCellId_r16.Encode(w); err != nil {
		return utils.WrapError("Encode PhysCellId_r16", err)
	}
	if err = ie.ResultsSSB_Cell_r16.Encode(w); err != nil {
		return utils.WrapError("Encode ResultsSSB_Cell_r16", err)
	}
	if ie.NumberOfGoodSSB_r16 != nil {
		if err = w.WriteInteger(*ie.NumberOfGoodSSB_r16, &uper.Constraint{Lb: 1, Ub: maxNrofSSBs_r16}, false); err != nil {
			return utils.WrapError("Encode NumberOfGoodSSB_r16", err)
		}
	}
	return nil
}

func (ie *MeasResultLoggingNR_r16) Decode(r *uper.UperReader) error {
	var err error
	var NumberOfGoodSSB_r16Present bool
	if NumberOfGoodSSB_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	if err = ie.PhysCellId_r16.Decode(r); err != nil {
		return utils.WrapError("Decode PhysCellId_r16", err)
	}
	if err = ie.ResultsSSB_Cell_r16.Decode(r); err != nil {
		return utils.WrapError("Decode ResultsSSB_Cell_r16", err)
	}
	if NumberOfGoodSSB_r16Present {
		var tmp_int_NumberOfGoodSSB_r16 int64
		if tmp_int_NumberOfGoodSSB_r16, err = r.ReadInteger(&uper.Constraint{Lb: 1, Ub: maxNrofSSBs_r16}, false); err != nil {
			return utils.WrapError("Decode NumberOfGoodSSB_r16", err)
		}
		ie.NumberOfGoodSSB_r16 = &tmp_int_NumberOfGoodSSB_r16
	}
	return nil
}
