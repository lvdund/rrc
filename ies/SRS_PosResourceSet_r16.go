package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type SRS_PosResourceSet_r16 struct {
	Srs_PosResourceSetId_r16    SRS_PosResourceSetId_r16                            `madatory`
	Srs_PosResourceIdList_r16   []SRS_PosResourceId_r16                             `lb:1,ub:maxNrofSRS_ResourcesPerSet,optional`
	ResourceType_r16            ResourceType_r16_SRS_PosResourceSet_r16             `madatory`
	Alpha_r16                   *Alpha                                              `optional`
	P0_r16                      *int64                                              `lb:-202,ub:24,optional`
	PathlossReferenceRS_Pos_r16 *SRS_PosResourceSet_r16_pathlossReferenceRS_Pos_r16 `optional`
}

func (ie *SRS_PosResourceSet_r16) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{len(ie.Srs_PosResourceIdList_r16) > 0, ie.Alpha_r16 != nil, ie.P0_r16 != nil, ie.PathlossReferenceRS_Pos_r16 != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if err = ie.Srs_PosResourceSetId_r16.Encode(w); err != nil {
		return utils.WrapError("Encode Srs_PosResourceSetId_r16", err)
	}
	if len(ie.Srs_PosResourceIdList_r16) > 0 {
		tmp_Srs_PosResourceIdList_r16 := utils.NewSequence[*SRS_PosResourceId_r16]([]*SRS_PosResourceId_r16{}, uper.Constraint{Lb: 1, Ub: maxNrofSRS_ResourcesPerSet}, false)
		for _, i := range ie.Srs_PosResourceIdList_r16 {
			tmp_Srs_PosResourceIdList_r16.Value = append(tmp_Srs_PosResourceIdList_r16.Value, &i)
		}
		if err = tmp_Srs_PosResourceIdList_r16.Encode(w); err != nil {
			return utils.WrapError("Encode Srs_PosResourceIdList_r16", err)
		}
	}
	if err = ie.ResourceType_r16.Encode(w); err != nil {
		return utils.WrapError("Encode ResourceType_r16", err)
	}
	if ie.Alpha_r16 != nil {
		if err = ie.Alpha_r16.Encode(w); err != nil {
			return utils.WrapError("Encode Alpha_r16", err)
		}
	}
	if ie.P0_r16 != nil {
		if err = w.WriteInteger(*ie.P0_r16, &uper.Constraint{Lb: -202, Ub: 24}, false); err != nil {
			return utils.WrapError("Encode P0_r16", err)
		}
	}
	if ie.PathlossReferenceRS_Pos_r16 != nil {
		if err = ie.PathlossReferenceRS_Pos_r16.Encode(w); err != nil {
			return utils.WrapError("Encode PathlossReferenceRS_Pos_r16", err)
		}
	}
	return nil
}

func (ie *SRS_PosResourceSet_r16) Decode(r *uper.UperReader) error {
	var err error
	var Srs_PosResourceIdList_r16Present bool
	if Srs_PosResourceIdList_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var Alpha_r16Present bool
	if Alpha_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var P0_r16Present bool
	if P0_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var PathlossReferenceRS_Pos_r16Present bool
	if PathlossReferenceRS_Pos_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	if err = ie.Srs_PosResourceSetId_r16.Decode(r); err != nil {
		return utils.WrapError("Decode Srs_PosResourceSetId_r16", err)
	}
	if Srs_PosResourceIdList_r16Present {
		tmp_Srs_PosResourceIdList_r16 := utils.NewSequence[*SRS_PosResourceId_r16]([]*SRS_PosResourceId_r16{}, uper.Constraint{Lb: 1, Ub: maxNrofSRS_ResourcesPerSet}, false)
		fn_Srs_PosResourceIdList_r16 := func() *SRS_PosResourceId_r16 {
			return new(SRS_PosResourceId_r16)
		}
		if err = tmp_Srs_PosResourceIdList_r16.Decode(r, fn_Srs_PosResourceIdList_r16); err != nil {
			return utils.WrapError("Decode Srs_PosResourceIdList_r16", err)
		}
		ie.Srs_PosResourceIdList_r16 = []SRS_PosResourceId_r16{}
		for _, i := range tmp_Srs_PosResourceIdList_r16.Value {
			ie.Srs_PosResourceIdList_r16 = append(ie.Srs_PosResourceIdList_r16, *i)
		}
	}
	if err = ie.ResourceType_r16.Decode(r); err != nil {
		return utils.WrapError("Decode ResourceType_r16", err)
	}
	if Alpha_r16Present {
		ie.Alpha_r16 = new(Alpha)
		if err = ie.Alpha_r16.Decode(r); err != nil {
			return utils.WrapError("Decode Alpha_r16", err)
		}
	}
	if P0_r16Present {
		var tmp_int_P0_r16 int64
		if tmp_int_P0_r16, err = r.ReadInteger(&uper.Constraint{Lb: -202, Ub: 24}, false); err != nil {
			return utils.WrapError("Decode P0_r16", err)
		}
		ie.P0_r16 = &tmp_int_P0_r16
	}
	if PathlossReferenceRS_Pos_r16Present {
		ie.PathlossReferenceRS_Pos_r16 = new(SRS_PosResourceSet_r16_pathlossReferenceRS_Pos_r16)
		if err = ie.PathlossReferenceRS_Pos_r16.Decode(r); err != nil {
			return utils.WrapError("Decode PathlossReferenceRS_Pos_r16", err)
		}
	}
	return nil
}
