package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

type BWP_UplinkDedicated_ul_TCI_StateList_r17_explicitlist struct {
	Ul_TCI_ToAddModList_r17  []TCI_UL_State_r17    `lb:1,ub:maxUL_TCI_r17,optional`
	Ul_TCI_ToReleaseList_r17 []TCI_UL_State_Id_r17 `lb:1,ub:maxUL_TCI_r17,optional`
}

func (ie *BWP_UplinkDedicated_ul_TCI_StateList_r17_explicitlist) Encode(w *aper.AperWriter) error {
	var err error
	preambleBits := []bool{len(ie.Ul_TCI_ToAddModList_r17) > 0, len(ie.Ul_TCI_ToReleaseList_r17) > 0}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if len(ie.Ul_TCI_ToAddModList_r17) > 0 {
		tmp_Ul_TCI_ToAddModList_r17 := utils.NewSequence[*TCI_UL_State_r17]([]*TCI_UL_State_r17{}, aper.Constraint{Lb: 1, Ub: maxUL_TCI_r17}, false)
		for _, i := range ie.Ul_TCI_ToAddModList_r17 {
			tmp_Ul_TCI_ToAddModList_r17.Value = append(tmp_Ul_TCI_ToAddModList_r17.Value, &i)
		}
		if err = tmp_Ul_TCI_ToAddModList_r17.Encode(w); err != nil {
			return utils.WrapError("Encode Ul_TCI_ToAddModList_r17", err)
		}
	}
	if len(ie.Ul_TCI_ToReleaseList_r17) > 0 {
		tmp_Ul_TCI_ToReleaseList_r17 := utils.NewSequence[*TCI_UL_State_Id_r17]([]*TCI_UL_State_Id_r17{}, aper.Constraint{Lb: 1, Ub: maxUL_TCI_r17}, false)
		for _, i := range ie.Ul_TCI_ToReleaseList_r17 {
			tmp_Ul_TCI_ToReleaseList_r17.Value = append(tmp_Ul_TCI_ToReleaseList_r17.Value, &i)
		}
		if err = tmp_Ul_TCI_ToReleaseList_r17.Encode(w); err != nil {
			return utils.WrapError("Encode Ul_TCI_ToReleaseList_r17", err)
		}
	}
	return nil
}

func (ie *BWP_UplinkDedicated_ul_TCI_StateList_r17_explicitlist) Decode(r *aper.AperReader) error {
	var err error
	var Ul_TCI_ToAddModList_r17Present bool
	if Ul_TCI_ToAddModList_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var Ul_TCI_ToReleaseList_r17Present bool
	if Ul_TCI_ToReleaseList_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	if Ul_TCI_ToAddModList_r17Present {
		tmp_Ul_TCI_ToAddModList_r17 := utils.NewSequence[*TCI_UL_State_r17]([]*TCI_UL_State_r17{}, aper.Constraint{Lb: 1, Ub: maxUL_TCI_r17}, false)
		fn_Ul_TCI_ToAddModList_r17 := func() *TCI_UL_State_r17 {
			return new(TCI_UL_State_r17)
		}
		if err = tmp_Ul_TCI_ToAddModList_r17.Decode(r, fn_Ul_TCI_ToAddModList_r17); err != nil {
			return utils.WrapError("Decode Ul_TCI_ToAddModList_r17", err)
		}
		ie.Ul_TCI_ToAddModList_r17 = []TCI_UL_State_r17{}
		for _, i := range tmp_Ul_TCI_ToAddModList_r17.Value {
			ie.Ul_TCI_ToAddModList_r17 = append(ie.Ul_TCI_ToAddModList_r17, *i)
		}
	}
	if Ul_TCI_ToReleaseList_r17Present {
		tmp_Ul_TCI_ToReleaseList_r17 := utils.NewSequence[*TCI_UL_State_Id_r17]([]*TCI_UL_State_Id_r17{}, aper.Constraint{Lb: 1, Ub: maxUL_TCI_r17}, false)
		fn_Ul_TCI_ToReleaseList_r17 := func() *TCI_UL_State_Id_r17 {
			return new(TCI_UL_State_Id_r17)
		}
		if err = tmp_Ul_TCI_ToReleaseList_r17.Decode(r, fn_Ul_TCI_ToReleaseList_r17); err != nil {
			return utils.WrapError("Decode Ul_TCI_ToReleaseList_r17", err)
		}
		ie.Ul_TCI_ToReleaseList_r17 = []TCI_UL_State_Id_r17{}
		for _, i := range tmp_Ul_TCI_ToReleaseList_r17.Value {
			ie.Ul_TCI_ToReleaseList_r17 = append(ie.Ul_TCI_ToReleaseList_r17, *i)
		}
	}
	return nil
}
