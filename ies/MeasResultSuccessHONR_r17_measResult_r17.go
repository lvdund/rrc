package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type MeasResultSuccessHONR_r17_measResult_r17 struct {
	CellResults_r17    *MeasResultSuccessHONR_r17_measResult_r17_cellResults_r17    `optional`
	RsIndexResults_r17 *MeasResultSuccessHONR_r17_measResult_r17_rsIndexResults_r17 `optional`
}

func (ie *MeasResultSuccessHONR_r17_measResult_r17) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.CellResults_r17 != nil, ie.RsIndexResults_r17 != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.CellResults_r17 != nil {
		if err = ie.CellResults_r17.Encode(w); err != nil {
			return utils.WrapError("Encode CellResults_r17", err)
		}
	}
	if ie.RsIndexResults_r17 != nil {
		if err = ie.RsIndexResults_r17.Encode(w); err != nil {
			return utils.WrapError("Encode RsIndexResults_r17", err)
		}
	}
	return nil
}

func (ie *MeasResultSuccessHONR_r17_measResult_r17) Decode(r *uper.UperReader) error {
	var err error
	var CellResults_r17Present bool
	if CellResults_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var RsIndexResults_r17Present bool
	if RsIndexResults_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	if CellResults_r17Present {
		ie.CellResults_r17 = new(MeasResultSuccessHONR_r17_measResult_r17_cellResults_r17)
		if err = ie.CellResults_r17.Decode(r); err != nil {
			return utils.WrapError("Decode CellResults_r17", err)
		}
	}
	if RsIndexResults_r17Present {
		ie.RsIndexResults_r17 = new(MeasResultSuccessHONR_r17_measResult_r17_rsIndexResults_r17)
		if err = ie.RsIndexResults_r17.Decode(r); err != nil {
			return utils.WrapError("Decode RsIndexResults_r17", err)
		}
	}
	return nil
}
