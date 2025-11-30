package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

type PDSCH_HARQ_ACK_EnhType3_r17 struct {
	Pdsch_HARQ_ACK_EnhType3Index_r17 PDSCH_HARQ_ACK_EnhType3Index_r17                            `madatory`
	Applicable_r17                   []int64                                                     `lb:1,ub:maxNrofServingCells,e_lb:0,e_ub:1,madatory`
	Pdsch_HARQ_ACK_EnhType3NDI_r17   *PDSCH_HARQ_ACK_EnhType3_r17_pdsch_HARQ_ACK_EnhType3NDI_r17 `optional`
	Pdsch_HARQ_ACK_EnhType3CBG_r17   *PDSCH_HARQ_ACK_EnhType3_r17_pdsch_HARQ_ACK_EnhType3CBG_r17 `optional`
}

func (ie *PDSCH_HARQ_ACK_EnhType3_r17) Encode(w *aper.AperWriter) error {
	var err error
	preambleBits := []bool{ie.Pdsch_HARQ_ACK_EnhType3NDI_r17 != nil, ie.Pdsch_HARQ_ACK_EnhType3CBG_r17 != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if err = ie.Pdsch_HARQ_ACK_EnhType3Index_r17.Encode(w); err != nil {
		return utils.WrapError("Encode Pdsch_HARQ_ACK_EnhType3Index_r17", err)
	}
	tmp_Applicable_r17 := utils.NewSequence[*utils.INTEGER]([]*utils.INTEGER{}, aper.Constraint{Lb: 1, Ub: maxNrofServingCells}, false)
	for _, i := range ie.Applicable_r17 {
		tmp_ie := utils.NewINTEGER(int64(i), aper.Constraint{Lb: 0, Ub: 1}, false)
		tmp_Applicable_r17.Value = append(tmp_Applicable_r17.Value, &tmp_ie)
	}
	if err = tmp_Applicable_r17.Encode(w); err != nil {
		return utils.WrapError("Encode Applicable_r17", err)
	}
	if ie.Pdsch_HARQ_ACK_EnhType3NDI_r17 != nil {
		if err = ie.Pdsch_HARQ_ACK_EnhType3NDI_r17.Encode(w); err != nil {
			return utils.WrapError("Encode Pdsch_HARQ_ACK_EnhType3NDI_r17", err)
		}
	}
	if ie.Pdsch_HARQ_ACK_EnhType3CBG_r17 != nil {
		if err = ie.Pdsch_HARQ_ACK_EnhType3CBG_r17.Encode(w); err != nil {
			return utils.WrapError("Encode Pdsch_HARQ_ACK_EnhType3CBG_r17", err)
		}
	}
	return nil
}

func (ie *PDSCH_HARQ_ACK_EnhType3_r17) Decode(r *aper.AperReader) error {
	var err error
	var Pdsch_HARQ_ACK_EnhType3NDI_r17Present bool
	if Pdsch_HARQ_ACK_EnhType3NDI_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var Pdsch_HARQ_ACK_EnhType3CBG_r17Present bool
	if Pdsch_HARQ_ACK_EnhType3CBG_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	if err = ie.Pdsch_HARQ_ACK_EnhType3Index_r17.Decode(r); err != nil {
		return utils.WrapError("Decode Pdsch_HARQ_ACK_EnhType3Index_r17", err)
	}
	tmp_Applicable_r17 := utils.NewSequence[*utils.INTEGER]([]*utils.INTEGER{}, aper.Constraint{Lb: 1, Ub: maxNrofServingCells}, false)
	fn_Applicable_r17 := func() *utils.INTEGER {
		ie := utils.NewINTEGER(0, aper.Constraint{Lb: 0, Ub: 1}, false)
		return &ie
	}
	if err = tmp_Applicable_r17.Decode(r, fn_Applicable_r17); err != nil {
		return utils.WrapError("Decode Applicable_r17", err)
	}
	ie.Applicable_r17 = []int64{}
	for _, i := range tmp_Applicable_r17.Value {
		ie.Applicable_r17 = append(ie.Applicable_r17, int64(i.Value))
	}
	if Pdsch_HARQ_ACK_EnhType3NDI_r17Present {
		ie.Pdsch_HARQ_ACK_EnhType3NDI_r17 = new(PDSCH_HARQ_ACK_EnhType3_r17_pdsch_HARQ_ACK_EnhType3NDI_r17)
		if err = ie.Pdsch_HARQ_ACK_EnhType3NDI_r17.Decode(r); err != nil {
			return utils.WrapError("Decode Pdsch_HARQ_ACK_EnhType3NDI_r17", err)
		}
	}
	if Pdsch_HARQ_ACK_EnhType3CBG_r17Present {
		ie.Pdsch_HARQ_ACK_EnhType3CBG_r17 = new(PDSCH_HARQ_ACK_EnhType3_r17_pdsch_HARQ_ACK_EnhType3CBG_r17)
		if err = ie.Pdsch_HARQ_ACK_EnhType3CBG_r17.Decode(r); err != nil {
			return utils.WrapError("Decode Pdsch_HARQ_ACK_EnhType3CBG_r17", err)
		}
	}
	return nil
}
