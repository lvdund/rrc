package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type CFRA_CSIRS_Resource struct {
	Csi_RS           CSI_RS_Index `madatory`
	Ra_OccasionList  []int64      `lb:1,ub:maxRA_OccasionsPerCSIRS,e_lb:0,e_ub:maxRA_Occasions_1,madatory`
	Ra_PreambleIndex int64        `lb:0,ub:63,madatory`
}

func (ie *CFRA_CSIRS_Resource) Encode(w *uper.UperWriter) error {
	var err error
	if err = ie.Csi_RS.Encode(w); err != nil {
		return utils.WrapError("Encode Csi_RS", err)
	}
	tmp_Ra_OccasionList := utils.NewSequence[*utils.INTEGER]([]*utils.INTEGER{}, uper.Constraint{Lb: 1, Ub: maxRA_OccasionsPerCSIRS}, false)
	for _, i := range ie.Ra_OccasionList {
		tmp_ie := utils.NewINTEGER(int64(i), uper.Constraint{Lb: 0, Ub: maxRA_Occasions_1}, false)
		tmp_Ra_OccasionList.Value = append(tmp_Ra_OccasionList.Value, &tmp_ie)
	}
	if err = tmp_Ra_OccasionList.Encode(w); err != nil {
		return utils.WrapError("Encode Ra_OccasionList", err)
	}
	if err = w.WriteInteger(ie.Ra_PreambleIndex, &uper.Constraint{Lb: 0, Ub: 63}, false); err != nil {
		return utils.WrapError("WriteInteger Ra_PreambleIndex", err)
	}
	return nil
}

func (ie *CFRA_CSIRS_Resource) Decode(r *uper.UperReader) error {
	var err error
	if err = ie.Csi_RS.Decode(r); err != nil {
		return utils.WrapError("Decode Csi_RS", err)
	}
	tmp_Ra_OccasionList := utils.NewSequence[*utils.INTEGER]([]*utils.INTEGER{}, uper.Constraint{Lb: 1, Ub: maxRA_OccasionsPerCSIRS}, false)
	fn_Ra_OccasionList := func() *utils.INTEGER {
		ie := utils.NewINTEGER(0, uper.Constraint{Lb: 0, Ub: maxRA_Occasions_1}, false)
		return &ie
	}
	if err = tmp_Ra_OccasionList.Decode(r, fn_Ra_OccasionList); err != nil {
		return utils.WrapError("Decode Ra_OccasionList", err)
	}
	ie.Ra_OccasionList = []int64{}
	for _, i := range tmp_Ra_OccasionList.Value {
		ie.Ra_OccasionList = append(ie.Ra_OccasionList, int64(i.Value))
	}
	var tmp_int_Ra_PreambleIndex int64
	if tmp_int_Ra_PreambleIndex, err = r.ReadInteger(&uper.Constraint{Lb: 0, Ub: 63}, false); err != nil {
		return utils.WrapError("ReadInteger Ra_PreambleIndex", err)
	}
	ie.Ra_PreambleIndex = tmp_int_Ra_PreambleIndex
	return nil
}
