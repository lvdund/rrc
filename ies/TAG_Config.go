package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type TAG_Config struct {
	Tag_ToReleaseList []TAG_Id `lb:1,ub:maxNrofTAGs,optional`
	Tag_ToAddModList  []TAG    `lb:1,ub:maxNrofTAGs,optional`
}

func (ie *TAG_Config) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{len(ie.Tag_ToReleaseList) > 0, len(ie.Tag_ToAddModList) > 0}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if len(ie.Tag_ToReleaseList) > 0 {
		tmp_Tag_ToReleaseList := utils.NewSequence[*TAG_Id]([]*TAG_Id{}, uper.Constraint{Lb: 1, Ub: maxNrofTAGs}, false)
		for _, i := range ie.Tag_ToReleaseList {
			tmp_Tag_ToReleaseList.Value = append(tmp_Tag_ToReleaseList.Value, &i)
		}
		if err = tmp_Tag_ToReleaseList.Encode(w); err != nil {
			return utils.WrapError("Encode Tag_ToReleaseList", err)
		}
	}
	if len(ie.Tag_ToAddModList) > 0 {
		tmp_Tag_ToAddModList := utils.NewSequence[*TAG]([]*TAG{}, uper.Constraint{Lb: 1, Ub: maxNrofTAGs}, false)
		for _, i := range ie.Tag_ToAddModList {
			tmp_Tag_ToAddModList.Value = append(tmp_Tag_ToAddModList.Value, &i)
		}
		if err = tmp_Tag_ToAddModList.Encode(w); err != nil {
			return utils.WrapError("Encode Tag_ToAddModList", err)
		}
	}
	return nil
}

func (ie *TAG_Config) Decode(r *uper.UperReader) error {
	var err error
	var Tag_ToReleaseListPresent bool
	if Tag_ToReleaseListPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var Tag_ToAddModListPresent bool
	if Tag_ToAddModListPresent, err = r.ReadBool(); err != nil {
		return err
	}
	if Tag_ToReleaseListPresent {
		tmp_Tag_ToReleaseList := utils.NewSequence[*TAG_Id]([]*TAG_Id{}, uper.Constraint{Lb: 1, Ub: maxNrofTAGs}, false)
		fn_Tag_ToReleaseList := func() *TAG_Id {
			return new(TAG_Id)
		}
		if err = tmp_Tag_ToReleaseList.Decode(r, fn_Tag_ToReleaseList); err != nil {
			return utils.WrapError("Decode Tag_ToReleaseList", err)
		}
		ie.Tag_ToReleaseList = []TAG_Id{}
		for _, i := range tmp_Tag_ToReleaseList.Value {
			ie.Tag_ToReleaseList = append(ie.Tag_ToReleaseList, *i)
		}
	}
	if Tag_ToAddModListPresent {
		tmp_Tag_ToAddModList := utils.NewSequence[*TAG]([]*TAG{}, uper.Constraint{Lb: 1, Ub: maxNrofTAGs}, false)
		fn_Tag_ToAddModList := func() *TAG {
			return new(TAG)
		}
		if err = tmp_Tag_ToAddModList.Decode(r, fn_Tag_ToAddModList); err != nil {
			return utils.WrapError("Decode Tag_ToAddModList", err)
		}
		ie.Tag_ToAddModList = []TAG{}
		for _, i := range tmp_Tag_ToAddModList.Value {
			ie.Tag_ToAddModList = append(ie.Tag_ToAddModList, *i)
		}
	}
	return nil
}
