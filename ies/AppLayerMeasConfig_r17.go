package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type AppLayerMeasConfig_r17 struct {
	MeasConfigAppLayerToAddModList_r17  []MeasConfigAppLayer_r17                   `lb:1,ub:maxNrofAppLayerMeas_r17,optional`
	MeasConfigAppLayerToReleaseList_r17 []MeasConfigAppLayerId_r17                 `lb:1,ub:maxNrofAppLayerMeas_r17,optional`
	Rrc_SegAllowed_r17                  *AppLayerMeasConfig_r17_rrc_SegAllowed_r17 `optional`
}

func (ie *AppLayerMeasConfig_r17) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{len(ie.MeasConfigAppLayerToAddModList_r17) > 0, len(ie.MeasConfigAppLayerToReleaseList_r17) > 0, ie.Rrc_SegAllowed_r17 != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if len(ie.MeasConfigAppLayerToAddModList_r17) > 0 {
		tmp_MeasConfigAppLayerToAddModList_r17 := utils.NewSequence[*MeasConfigAppLayer_r17]([]*MeasConfigAppLayer_r17{}, uper.Constraint{Lb: 1, Ub: maxNrofAppLayerMeas_r17}, false)
		for _, i := range ie.MeasConfigAppLayerToAddModList_r17 {
			tmp_MeasConfigAppLayerToAddModList_r17.Value = append(tmp_MeasConfigAppLayerToAddModList_r17.Value, &i)
		}
		if err = tmp_MeasConfigAppLayerToAddModList_r17.Encode(w); err != nil {
			return utils.WrapError("Encode MeasConfigAppLayerToAddModList_r17", err)
		}
	}
	if len(ie.MeasConfigAppLayerToReleaseList_r17) > 0 {
		tmp_MeasConfigAppLayerToReleaseList_r17 := utils.NewSequence[*MeasConfigAppLayerId_r17]([]*MeasConfigAppLayerId_r17{}, uper.Constraint{Lb: 1, Ub: maxNrofAppLayerMeas_r17}, false)
		for _, i := range ie.MeasConfigAppLayerToReleaseList_r17 {
			tmp_MeasConfigAppLayerToReleaseList_r17.Value = append(tmp_MeasConfigAppLayerToReleaseList_r17.Value, &i)
		}
		if err = tmp_MeasConfigAppLayerToReleaseList_r17.Encode(w); err != nil {
			return utils.WrapError("Encode MeasConfigAppLayerToReleaseList_r17", err)
		}
	}
	if ie.Rrc_SegAllowed_r17 != nil {
		if err = ie.Rrc_SegAllowed_r17.Encode(w); err != nil {
			return utils.WrapError("Encode Rrc_SegAllowed_r17", err)
		}
	}
	return nil
}

func (ie *AppLayerMeasConfig_r17) Decode(r *uper.UperReader) error {
	var err error
	var MeasConfigAppLayerToAddModList_r17Present bool
	if MeasConfigAppLayerToAddModList_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var MeasConfigAppLayerToReleaseList_r17Present bool
	if MeasConfigAppLayerToReleaseList_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var Rrc_SegAllowed_r17Present bool
	if Rrc_SegAllowed_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	if MeasConfigAppLayerToAddModList_r17Present {
		tmp_MeasConfigAppLayerToAddModList_r17 := utils.NewSequence[*MeasConfigAppLayer_r17]([]*MeasConfigAppLayer_r17{}, uper.Constraint{Lb: 1, Ub: maxNrofAppLayerMeas_r17}, false)
		fn_MeasConfigAppLayerToAddModList_r17 := func() *MeasConfigAppLayer_r17 {
			return new(MeasConfigAppLayer_r17)
		}
		if err = tmp_MeasConfigAppLayerToAddModList_r17.Decode(r, fn_MeasConfigAppLayerToAddModList_r17); err != nil {
			return utils.WrapError("Decode MeasConfigAppLayerToAddModList_r17", err)
		}
		ie.MeasConfigAppLayerToAddModList_r17 = []MeasConfigAppLayer_r17{}
		for _, i := range tmp_MeasConfigAppLayerToAddModList_r17.Value {
			ie.MeasConfigAppLayerToAddModList_r17 = append(ie.MeasConfigAppLayerToAddModList_r17, *i)
		}
	}
	if MeasConfigAppLayerToReleaseList_r17Present {
		tmp_MeasConfigAppLayerToReleaseList_r17 := utils.NewSequence[*MeasConfigAppLayerId_r17]([]*MeasConfigAppLayerId_r17{}, uper.Constraint{Lb: 1, Ub: maxNrofAppLayerMeas_r17}, false)
		fn_MeasConfigAppLayerToReleaseList_r17 := func() *MeasConfigAppLayerId_r17 {
			return new(MeasConfigAppLayerId_r17)
		}
		if err = tmp_MeasConfigAppLayerToReleaseList_r17.Decode(r, fn_MeasConfigAppLayerToReleaseList_r17); err != nil {
			return utils.WrapError("Decode MeasConfigAppLayerToReleaseList_r17", err)
		}
		ie.MeasConfigAppLayerToReleaseList_r17 = []MeasConfigAppLayerId_r17{}
		for _, i := range tmp_MeasConfigAppLayerToReleaseList_r17.Value {
			ie.MeasConfigAppLayerToReleaseList_r17 = append(ie.MeasConfigAppLayerToReleaseList_r17, *i)
		}
	}
	if Rrc_SegAllowed_r17Present {
		ie.Rrc_SegAllowed_r17 = new(AppLayerMeasConfig_r17_rrc_SegAllowed_r17)
		if err = ie.Rrc_SegAllowed_r17.Decode(r); err != nil {
			return utils.WrapError("Decode Rrc_SegAllowed_r17", err)
		}
	}
	return nil
}
