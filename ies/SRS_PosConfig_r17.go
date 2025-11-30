package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

type SRS_PosConfig_r17 struct {
	Srs_PosResourceSetToReleaseList_r17 []SRS_PosResourceSetId_r16 `lb:1,ub:maxNrofSRS_PosResourceSets_r16,optional`
	Srs_PosResourceSetToAddModList_r17  []SRS_PosResourceSet_r16   `lb:1,ub:maxNrofSRS_PosResourceSets_r16,optional`
	Srs_PosResourceToReleaseList_r17    []SRS_PosResourceId_r16    `lb:1,ub:maxNrofSRS_PosResources_r16,optional`
	Srs_PosResourceToAddModList_r17     []SRS_PosResource_r16      `lb:1,ub:maxNrofSRS_PosResources_r16,optional`
}

func (ie *SRS_PosConfig_r17) Encode(w *aper.AperWriter) error {
	var err error
	preambleBits := []bool{len(ie.Srs_PosResourceSetToReleaseList_r17) > 0, len(ie.Srs_PosResourceSetToAddModList_r17) > 0, len(ie.Srs_PosResourceToReleaseList_r17) > 0, len(ie.Srs_PosResourceToAddModList_r17) > 0}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if len(ie.Srs_PosResourceSetToReleaseList_r17) > 0 {
		tmp_Srs_PosResourceSetToReleaseList_r17 := utils.NewSequence[*SRS_PosResourceSetId_r16]([]*SRS_PosResourceSetId_r16{}, aper.Constraint{Lb: 1, Ub: maxNrofSRS_PosResourceSets_r16}, false)
		for _, i := range ie.Srs_PosResourceSetToReleaseList_r17 {
			tmp_Srs_PosResourceSetToReleaseList_r17.Value = append(tmp_Srs_PosResourceSetToReleaseList_r17.Value, &i)
		}
		if err = tmp_Srs_PosResourceSetToReleaseList_r17.Encode(w); err != nil {
			return utils.WrapError("Encode Srs_PosResourceSetToReleaseList_r17", err)
		}
	}
	if len(ie.Srs_PosResourceSetToAddModList_r17) > 0 {
		tmp_Srs_PosResourceSetToAddModList_r17 := utils.NewSequence[*SRS_PosResourceSet_r16]([]*SRS_PosResourceSet_r16{}, aper.Constraint{Lb: 1, Ub: maxNrofSRS_PosResourceSets_r16}, false)
		for _, i := range ie.Srs_PosResourceSetToAddModList_r17 {
			tmp_Srs_PosResourceSetToAddModList_r17.Value = append(tmp_Srs_PosResourceSetToAddModList_r17.Value, &i)
		}
		if err = tmp_Srs_PosResourceSetToAddModList_r17.Encode(w); err != nil {
			return utils.WrapError("Encode Srs_PosResourceSetToAddModList_r17", err)
		}
	}
	if len(ie.Srs_PosResourceToReleaseList_r17) > 0 {
		tmp_Srs_PosResourceToReleaseList_r17 := utils.NewSequence[*SRS_PosResourceId_r16]([]*SRS_PosResourceId_r16{}, aper.Constraint{Lb: 1, Ub: maxNrofSRS_PosResources_r16}, false)
		for _, i := range ie.Srs_PosResourceToReleaseList_r17 {
			tmp_Srs_PosResourceToReleaseList_r17.Value = append(tmp_Srs_PosResourceToReleaseList_r17.Value, &i)
		}
		if err = tmp_Srs_PosResourceToReleaseList_r17.Encode(w); err != nil {
			return utils.WrapError("Encode Srs_PosResourceToReleaseList_r17", err)
		}
	}
	if len(ie.Srs_PosResourceToAddModList_r17) > 0 {
		tmp_Srs_PosResourceToAddModList_r17 := utils.NewSequence[*SRS_PosResource_r16]([]*SRS_PosResource_r16{}, aper.Constraint{Lb: 1, Ub: maxNrofSRS_PosResources_r16}, false)
		for _, i := range ie.Srs_PosResourceToAddModList_r17 {
			tmp_Srs_PosResourceToAddModList_r17.Value = append(tmp_Srs_PosResourceToAddModList_r17.Value, &i)
		}
		if err = tmp_Srs_PosResourceToAddModList_r17.Encode(w); err != nil {
			return utils.WrapError("Encode Srs_PosResourceToAddModList_r17", err)
		}
	}
	return nil
}

func (ie *SRS_PosConfig_r17) Decode(r *aper.AperReader) error {
	var err error
	var Srs_PosResourceSetToReleaseList_r17Present bool
	if Srs_PosResourceSetToReleaseList_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var Srs_PosResourceSetToAddModList_r17Present bool
	if Srs_PosResourceSetToAddModList_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var Srs_PosResourceToReleaseList_r17Present bool
	if Srs_PosResourceToReleaseList_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var Srs_PosResourceToAddModList_r17Present bool
	if Srs_PosResourceToAddModList_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	if Srs_PosResourceSetToReleaseList_r17Present {
		tmp_Srs_PosResourceSetToReleaseList_r17 := utils.NewSequence[*SRS_PosResourceSetId_r16]([]*SRS_PosResourceSetId_r16{}, aper.Constraint{Lb: 1, Ub: maxNrofSRS_PosResourceSets_r16}, false)
		fn_Srs_PosResourceSetToReleaseList_r17 := func() *SRS_PosResourceSetId_r16 {
			return new(SRS_PosResourceSetId_r16)
		}
		if err = tmp_Srs_PosResourceSetToReleaseList_r17.Decode(r, fn_Srs_PosResourceSetToReleaseList_r17); err != nil {
			return utils.WrapError("Decode Srs_PosResourceSetToReleaseList_r17", err)
		}
		ie.Srs_PosResourceSetToReleaseList_r17 = []SRS_PosResourceSetId_r16{}
		for _, i := range tmp_Srs_PosResourceSetToReleaseList_r17.Value {
			ie.Srs_PosResourceSetToReleaseList_r17 = append(ie.Srs_PosResourceSetToReleaseList_r17, *i)
		}
	}
	if Srs_PosResourceSetToAddModList_r17Present {
		tmp_Srs_PosResourceSetToAddModList_r17 := utils.NewSequence[*SRS_PosResourceSet_r16]([]*SRS_PosResourceSet_r16{}, aper.Constraint{Lb: 1, Ub: maxNrofSRS_PosResourceSets_r16}, false)
		fn_Srs_PosResourceSetToAddModList_r17 := func() *SRS_PosResourceSet_r16 {
			return new(SRS_PosResourceSet_r16)
		}
		if err = tmp_Srs_PosResourceSetToAddModList_r17.Decode(r, fn_Srs_PosResourceSetToAddModList_r17); err != nil {
			return utils.WrapError("Decode Srs_PosResourceSetToAddModList_r17", err)
		}
		ie.Srs_PosResourceSetToAddModList_r17 = []SRS_PosResourceSet_r16{}
		for _, i := range tmp_Srs_PosResourceSetToAddModList_r17.Value {
			ie.Srs_PosResourceSetToAddModList_r17 = append(ie.Srs_PosResourceSetToAddModList_r17, *i)
		}
	}
	if Srs_PosResourceToReleaseList_r17Present {
		tmp_Srs_PosResourceToReleaseList_r17 := utils.NewSequence[*SRS_PosResourceId_r16]([]*SRS_PosResourceId_r16{}, aper.Constraint{Lb: 1, Ub: maxNrofSRS_PosResources_r16}, false)
		fn_Srs_PosResourceToReleaseList_r17 := func() *SRS_PosResourceId_r16 {
			return new(SRS_PosResourceId_r16)
		}
		if err = tmp_Srs_PosResourceToReleaseList_r17.Decode(r, fn_Srs_PosResourceToReleaseList_r17); err != nil {
			return utils.WrapError("Decode Srs_PosResourceToReleaseList_r17", err)
		}
		ie.Srs_PosResourceToReleaseList_r17 = []SRS_PosResourceId_r16{}
		for _, i := range tmp_Srs_PosResourceToReleaseList_r17.Value {
			ie.Srs_PosResourceToReleaseList_r17 = append(ie.Srs_PosResourceToReleaseList_r17, *i)
		}
	}
	if Srs_PosResourceToAddModList_r17Present {
		tmp_Srs_PosResourceToAddModList_r17 := utils.NewSequence[*SRS_PosResource_r16]([]*SRS_PosResource_r16{}, aper.Constraint{Lb: 1, Ub: maxNrofSRS_PosResources_r16}, false)
		fn_Srs_PosResourceToAddModList_r17 := func() *SRS_PosResource_r16 {
			return new(SRS_PosResource_r16)
		}
		if err = tmp_Srs_PosResourceToAddModList_r17.Decode(r, fn_Srs_PosResourceToAddModList_r17); err != nil {
			return utils.WrapError("Decode Srs_PosResourceToAddModList_r17", err)
		}
		ie.Srs_PosResourceToAddModList_r17 = []SRS_PosResource_r16{}
		for _, i := range tmp_Srs_PosResourceToAddModList_r17.Value {
			ie.Srs_PosResourceToAddModList_r17 = append(ie.Srs_PosResourceToAddModList_r17, *i)
		}
	}
	return nil
}
