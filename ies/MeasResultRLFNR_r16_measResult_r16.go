package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type MeasResultRLFNR_r16_measResult_r16 struct {
	CellResults_r16    *MeasResultRLFNR_r16_measResult_r16_cellResults_r16    `optional`
	RsIndexResults_r16 *MeasResultRLFNR_r16_measResult_r16_rsIndexResults_r16 `lb:64,ub:64,optional`
}

func (ie *MeasResultRLFNR_r16_measResult_r16) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.CellResults_r16 != nil, ie.RsIndexResults_r16 != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.CellResults_r16 != nil {
		if err = ie.CellResults_r16.Encode(w); err != nil {
			return utils.WrapError("Encode CellResults_r16", err)
		}
	}
	if ie.RsIndexResults_r16 != nil {
		if err = ie.RsIndexResults_r16.Encode(w); err != nil {
			return utils.WrapError("Encode RsIndexResults_r16", err)
		}
	}
	return nil
}

func (ie *MeasResultRLFNR_r16_measResult_r16) Decode(r *uper.UperReader) error {
	var err error
	var CellResults_r16Present bool
	if CellResults_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var RsIndexResults_r16Present bool
	if RsIndexResults_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	if CellResults_r16Present {
		ie.CellResults_r16 = new(MeasResultRLFNR_r16_measResult_r16_cellResults_r16)
		if err = ie.CellResults_r16.Decode(r); err != nil {
			return utils.WrapError("Decode CellResults_r16", err)
		}
	}
	if RsIndexResults_r16Present {
		ie.RsIndexResults_r16 = new(MeasResultRLFNR_r16_measResult_r16_rsIndexResults_r16)
		if err = ie.RsIndexResults_r16.Decode(r); err != nil {
			return utils.WrapError("Decode RsIndexResults_r16", err)
		}
	}
	return nil
}
