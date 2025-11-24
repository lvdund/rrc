package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type SchedulingRequestConfig struct {
	SchedulingRequestToAddModList  []SchedulingRequestToAddMod `lb:1,ub:maxNrofSR_ConfigPerCellGroup,optional`
	SchedulingRequestToReleaseList []SchedulingRequestId       `lb:1,ub:maxNrofSR_ConfigPerCellGroup,optional`
}

func (ie *SchedulingRequestConfig) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{len(ie.SchedulingRequestToAddModList) > 0, len(ie.SchedulingRequestToReleaseList) > 0}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if len(ie.SchedulingRequestToAddModList) > 0 {
		tmp_SchedulingRequestToAddModList := utils.NewSequence[*SchedulingRequestToAddMod]([]*SchedulingRequestToAddMod{}, uper.Constraint{Lb: 1, Ub: maxNrofSR_ConfigPerCellGroup}, false)
		for _, i := range ie.SchedulingRequestToAddModList {
			tmp_SchedulingRequestToAddModList.Value = append(tmp_SchedulingRequestToAddModList.Value, &i)
		}
		if err = tmp_SchedulingRequestToAddModList.Encode(w); err != nil {
			return utils.WrapError("Encode SchedulingRequestToAddModList", err)
		}
	}
	if len(ie.SchedulingRequestToReleaseList) > 0 {
		tmp_SchedulingRequestToReleaseList := utils.NewSequence[*SchedulingRequestId]([]*SchedulingRequestId{}, uper.Constraint{Lb: 1, Ub: maxNrofSR_ConfigPerCellGroup}, false)
		for _, i := range ie.SchedulingRequestToReleaseList {
			tmp_SchedulingRequestToReleaseList.Value = append(tmp_SchedulingRequestToReleaseList.Value, &i)
		}
		if err = tmp_SchedulingRequestToReleaseList.Encode(w); err != nil {
			return utils.WrapError("Encode SchedulingRequestToReleaseList", err)
		}
	}
	return nil
}

func (ie *SchedulingRequestConfig) Decode(r *uper.UperReader) error {
	var err error
	var SchedulingRequestToAddModListPresent bool
	if SchedulingRequestToAddModListPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var SchedulingRequestToReleaseListPresent bool
	if SchedulingRequestToReleaseListPresent, err = r.ReadBool(); err != nil {
		return err
	}
	if SchedulingRequestToAddModListPresent {
		tmp_SchedulingRequestToAddModList := utils.NewSequence[*SchedulingRequestToAddMod]([]*SchedulingRequestToAddMod{}, uper.Constraint{Lb: 1, Ub: maxNrofSR_ConfigPerCellGroup}, false)
		fn_SchedulingRequestToAddModList := func() *SchedulingRequestToAddMod {
			return new(SchedulingRequestToAddMod)
		}
		if err = tmp_SchedulingRequestToAddModList.Decode(r, fn_SchedulingRequestToAddModList); err != nil {
			return utils.WrapError("Decode SchedulingRequestToAddModList", err)
		}
		ie.SchedulingRequestToAddModList = []SchedulingRequestToAddMod{}
		for _, i := range tmp_SchedulingRequestToAddModList.Value {
			ie.SchedulingRequestToAddModList = append(ie.SchedulingRequestToAddModList, *i)
		}
	}
	if SchedulingRequestToReleaseListPresent {
		tmp_SchedulingRequestToReleaseList := utils.NewSequence[*SchedulingRequestId]([]*SchedulingRequestId{}, uper.Constraint{Lb: 1, Ub: maxNrofSR_ConfigPerCellGroup}, false)
		fn_SchedulingRequestToReleaseList := func() *SchedulingRequestId {
			return new(SchedulingRequestId)
		}
		if err = tmp_SchedulingRequestToReleaseList.Decode(r, fn_SchedulingRequestToReleaseList); err != nil {
			return utils.WrapError("Decode SchedulingRequestToReleaseList", err)
		}
		ie.SchedulingRequestToReleaseList = []SchedulingRequestId{}
		for _, i := range tmp_SchedulingRequestToReleaseList.Value {
			ie.SchedulingRequestToReleaseList = append(ie.SchedulingRequestToReleaseList, *i)
		}
	}
	return nil
}
