package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type MUSIM_GapConfig_r17 struct {
	Musim_GapToReleaseList_r17 []MUSIM_GapId_r17  `lb:1,ub:3,optional`
	Musim_GapToAddModList_r17  []MUSIM_Gap_r17    `lb:1,ub:3,optional`
	Musim_AperiodicGap_r17     *MUSIM_GapInfo_r17 `optional`
}

func (ie *MUSIM_GapConfig_r17) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{len(ie.Musim_GapToReleaseList_r17) > 0, len(ie.Musim_GapToAddModList_r17) > 0, ie.Musim_AperiodicGap_r17 != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if len(ie.Musim_GapToReleaseList_r17) > 0 {
		tmp_Musim_GapToReleaseList_r17 := utils.NewSequence[*MUSIM_GapId_r17]([]*MUSIM_GapId_r17{}, uper.Constraint{Lb: 1, Ub: 3}, false)
		for _, i := range ie.Musim_GapToReleaseList_r17 {
			tmp_Musim_GapToReleaseList_r17.Value = append(tmp_Musim_GapToReleaseList_r17.Value, &i)
		}
		if err = tmp_Musim_GapToReleaseList_r17.Encode(w); err != nil {
			return utils.WrapError("Encode Musim_GapToReleaseList_r17", err)
		}
	}
	if len(ie.Musim_GapToAddModList_r17) > 0 {
		tmp_Musim_GapToAddModList_r17 := utils.NewSequence[*MUSIM_Gap_r17]([]*MUSIM_Gap_r17{}, uper.Constraint{Lb: 1, Ub: 3}, false)
		for _, i := range ie.Musim_GapToAddModList_r17 {
			tmp_Musim_GapToAddModList_r17.Value = append(tmp_Musim_GapToAddModList_r17.Value, &i)
		}
		if err = tmp_Musim_GapToAddModList_r17.Encode(w); err != nil {
			return utils.WrapError("Encode Musim_GapToAddModList_r17", err)
		}
	}
	if ie.Musim_AperiodicGap_r17 != nil {
		if err = ie.Musim_AperiodicGap_r17.Encode(w); err != nil {
			return utils.WrapError("Encode Musim_AperiodicGap_r17", err)
		}
	}
	return nil
}

func (ie *MUSIM_GapConfig_r17) Decode(r *uper.UperReader) error {
	var err error
	var Musim_GapToReleaseList_r17Present bool
	if Musim_GapToReleaseList_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var Musim_GapToAddModList_r17Present bool
	if Musim_GapToAddModList_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var Musim_AperiodicGap_r17Present bool
	if Musim_AperiodicGap_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	if Musim_GapToReleaseList_r17Present {
		tmp_Musim_GapToReleaseList_r17 := utils.NewSequence[*MUSIM_GapId_r17]([]*MUSIM_GapId_r17{}, uper.Constraint{Lb: 1, Ub: 3}, false)
		fn_Musim_GapToReleaseList_r17 := func() *MUSIM_GapId_r17 {
			return new(MUSIM_GapId_r17)
		}
		if err = tmp_Musim_GapToReleaseList_r17.Decode(r, fn_Musim_GapToReleaseList_r17); err != nil {
			return utils.WrapError("Decode Musim_GapToReleaseList_r17", err)
		}
		ie.Musim_GapToReleaseList_r17 = []MUSIM_GapId_r17{}
		for _, i := range tmp_Musim_GapToReleaseList_r17.Value {
			ie.Musim_GapToReleaseList_r17 = append(ie.Musim_GapToReleaseList_r17, *i)
		}
	}
	if Musim_GapToAddModList_r17Present {
		tmp_Musim_GapToAddModList_r17 := utils.NewSequence[*MUSIM_Gap_r17]([]*MUSIM_Gap_r17{}, uper.Constraint{Lb: 1, Ub: 3}, false)
		fn_Musim_GapToAddModList_r17 := func() *MUSIM_Gap_r17 {
			return new(MUSIM_Gap_r17)
		}
		if err = tmp_Musim_GapToAddModList_r17.Decode(r, fn_Musim_GapToAddModList_r17); err != nil {
			return utils.WrapError("Decode Musim_GapToAddModList_r17", err)
		}
		ie.Musim_GapToAddModList_r17 = []MUSIM_Gap_r17{}
		for _, i := range tmp_Musim_GapToAddModList_r17.Value {
			ie.Musim_GapToAddModList_r17 = append(ie.Musim_GapToAddModList_r17, *i)
		}
	}
	if Musim_AperiodicGap_r17Present {
		ie.Musim_AperiodicGap_r17 = new(MUSIM_GapInfo_r17)
		if err = ie.Musim_AperiodicGap_r17.Decode(r); err != nil {
			return utils.WrapError("Decode Musim_AperiodicGap_r17", err)
		}
	}
	return nil
}
