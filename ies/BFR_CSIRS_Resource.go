package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

type BFR_CSIRS_Resource struct {
	Csi_RS           NZP_CSI_RS_ResourceId `madatory`
	Ra_OccasionList  []int64               `lb:1,ub:maxRA_OccasionsPerCSIRS,e_lb:0,e_ub:maxRA_Occasions_1,optional`
	Ra_PreambleIndex *int64                `lb:0,ub:63,optional`
}

func (ie *BFR_CSIRS_Resource) Encode(w *aper.AperWriter) error {
	var err error
	preambleBits := []bool{len(ie.Ra_OccasionList) > 0, ie.Ra_PreambleIndex != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if err = ie.Csi_RS.Encode(w); err != nil {
		return utils.WrapError("Encode Csi_RS", err)
	}
	if len(ie.Ra_OccasionList) > 0 {
		tmp_Ra_OccasionList := utils.NewSequence[*utils.INTEGER]([]*utils.INTEGER{}, aper.Constraint{Lb: 1, Ub: maxRA_OccasionsPerCSIRS}, false)
		for _, i := range ie.Ra_OccasionList {
			tmp_ie := utils.NewINTEGER(int64(i), aper.Constraint{Lb: 0, Ub: maxRA_Occasions_1}, false)
			tmp_Ra_OccasionList.Value = append(tmp_Ra_OccasionList.Value, &tmp_ie)
		}
		if err = tmp_Ra_OccasionList.Encode(w); err != nil {
			return utils.WrapError("Encode Ra_OccasionList", err)
		}
	}
	if ie.Ra_PreambleIndex != nil {
		if err = w.WriteInteger(*ie.Ra_PreambleIndex, &aper.Constraint{Lb: 0, Ub: 63}, false); err != nil {
			return utils.WrapError("Encode Ra_PreambleIndex", err)
		}
	}
	return nil
}

func (ie *BFR_CSIRS_Resource) Decode(r *aper.AperReader) error {
	var err error
	var Ra_OccasionListPresent bool
	if Ra_OccasionListPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var Ra_PreambleIndexPresent bool
	if Ra_PreambleIndexPresent, err = r.ReadBool(); err != nil {
		return err
	}
	if err = ie.Csi_RS.Decode(r); err != nil {
		return utils.WrapError("Decode Csi_RS", err)
	}
	if Ra_OccasionListPresent {
		tmp_Ra_OccasionList := utils.NewSequence[*utils.INTEGER]([]*utils.INTEGER{}, aper.Constraint{Lb: 1, Ub: maxRA_OccasionsPerCSIRS}, false)
		fn_Ra_OccasionList := func() *utils.INTEGER {
			ie := utils.NewINTEGER(0, aper.Constraint{Lb: 0, Ub: maxRA_Occasions_1}, false)
			return &ie
		}
		if err = tmp_Ra_OccasionList.Decode(r, fn_Ra_OccasionList); err != nil {
			return utils.WrapError("Decode Ra_OccasionList", err)
		}
		ie.Ra_OccasionList = []int64{}
		for _, i := range tmp_Ra_OccasionList.Value {
			ie.Ra_OccasionList = append(ie.Ra_OccasionList, int64(i.Value))
		}
	}
	if Ra_PreambleIndexPresent {
		var tmp_int_Ra_PreambleIndex int64
		if tmp_int_Ra_PreambleIndex, err = r.ReadInteger(&aper.Constraint{Lb: 0, Ub: 63}, false); err != nil {
			return utils.WrapError("Decode Ra_PreambleIndex", err)
		}
		ie.Ra_PreambleIndex = &tmp_int_Ra_PreambleIndex
	}
	return nil
}
