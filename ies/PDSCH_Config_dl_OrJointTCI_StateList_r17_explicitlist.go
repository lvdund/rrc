package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

type PDSCH_Config_dl_OrJointTCI_StateList_r17_explicitlist struct {
	Dl_OrJointTCI_StateToAddModList_r17  []TCI_State   `lb:1,ub:maxNrofTCI_States,optional`
	Dl_OrJointTCI_StateToReleaseList_r17 []TCI_StateId `lb:1,ub:maxNrofTCI_States,optional`
}

func (ie *PDSCH_Config_dl_OrJointTCI_StateList_r17_explicitlist) Encode(w *aper.AperWriter) error {
	var err error
	preambleBits := []bool{len(ie.Dl_OrJointTCI_StateToAddModList_r17) > 0, len(ie.Dl_OrJointTCI_StateToReleaseList_r17) > 0}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if len(ie.Dl_OrJointTCI_StateToAddModList_r17) > 0 {
		tmp_Dl_OrJointTCI_StateToAddModList_r17 := utils.NewSequence[*TCI_State]([]*TCI_State{}, aper.Constraint{Lb: 1, Ub: maxNrofTCI_States}, false)
		for _, i := range ie.Dl_OrJointTCI_StateToAddModList_r17 {
			tmp_Dl_OrJointTCI_StateToAddModList_r17.Value = append(tmp_Dl_OrJointTCI_StateToAddModList_r17.Value, &i)
		}
		if err = tmp_Dl_OrJointTCI_StateToAddModList_r17.Encode(w); err != nil {
			return utils.WrapError("Encode Dl_OrJointTCI_StateToAddModList_r17", err)
		}
	}
	if len(ie.Dl_OrJointTCI_StateToReleaseList_r17) > 0 {
		tmp_Dl_OrJointTCI_StateToReleaseList_r17 := utils.NewSequence[*TCI_StateId]([]*TCI_StateId{}, aper.Constraint{Lb: 1, Ub: maxNrofTCI_States}, false)
		for _, i := range ie.Dl_OrJointTCI_StateToReleaseList_r17 {
			tmp_Dl_OrJointTCI_StateToReleaseList_r17.Value = append(tmp_Dl_OrJointTCI_StateToReleaseList_r17.Value, &i)
		}
		if err = tmp_Dl_OrJointTCI_StateToReleaseList_r17.Encode(w); err != nil {
			return utils.WrapError("Encode Dl_OrJointTCI_StateToReleaseList_r17", err)
		}
	}
	return nil
}

func (ie *PDSCH_Config_dl_OrJointTCI_StateList_r17_explicitlist) Decode(r *aper.AperReader) error {
	var err error
	var Dl_OrJointTCI_StateToAddModList_r17Present bool
	if Dl_OrJointTCI_StateToAddModList_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var Dl_OrJointTCI_StateToReleaseList_r17Present bool
	if Dl_OrJointTCI_StateToReleaseList_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	if Dl_OrJointTCI_StateToAddModList_r17Present {
		tmp_Dl_OrJointTCI_StateToAddModList_r17 := utils.NewSequence[*TCI_State]([]*TCI_State{}, aper.Constraint{Lb: 1, Ub: maxNrofTCI_States}, false)
		fn_Dl_OrJointTCI_StateToAddModList_r17 := func() *TCI_State {
			return new(TCI_State)
		}
		if err = tmp_Dl_OrJointTCI_StateToAddModList_r17.Decode(r, fn_Dl_OrJointTCI_StateToAddModList_r17); err != nil {
			return utils.WrapError("Decode Dl_OrJointTCI_StateToAddModList_r17", err)
		}
		ie.Dl_OrJointTCI_StateToAddModList_r17 = []TCI_State{}
		for _, i := range tmp_Dl_OrJointTCI_StateToAddModList_r17.Value {
			ie.Dl_OrJointTCI_StateToAddModList_r17 = append(ie.Dl_OrJointTCI_StateToAddModList_r17, *i)
		}
	}
	if Dl_OrJointTCI_StateToReleaseList_r17Present {
		tmp_Dl_OrJointTCI_StateToReleaseList_r17 := utils.NewSequence[*TCI_StateId]([]*TCI_StateId{}, aper.Constraint{Lb: 1, Ub: maxNrofTCI_States}, false)
		fn_Dl_OrJointTCI_StateToReleaseList_r17 := func() *TCI_StateId {
			return new(TCI_StateId)
		}
		if err = tmp_Dl_OrJointTCI_StateToReleaseList_r17.Decode(r, fn_Dl_OrJointTCI_StateToReleaseList_r17); err != nil {
			return utils.WrapError("Decode Dl_OrJointTCI_StateToReleaseList_r17", err)
		}
		ie.Dl_OrJointTCI_StateToReleaseList_r17 = []TCI_StateId{}
		for _, i := range tmp_Dl_OrJointTCI_StateToReleaseList_r17.Value {
			ie.Dl_OrJointTCI_StateToReleaseList_r17 = append(ie.Dl_OrJointTCI_StateToReleaseList_r17, *i)
		}
	}
	return nil
}
