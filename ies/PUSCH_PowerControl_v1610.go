package ies

import (
	"bytes"

	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

type PUSCH_PowerControl_v1610 struct {
	PathlossReferenceRSToAddModListSizeExt_v1610  []PUSCH_PathlossReferenceRS_r16             `lb:1,ub:maxNrofPUSCH_PathlossReferenceRSsDiff_r16,optional`
	PathlossReferenceRSToReleaseListSizeExt_v1610 []PUSCH_PathlossReferenceRS_Id_v1610        `lb:1,ub:maxNrofPUSCH_PathlossReferenceRSsDiff_r16,optional`
	P0_PUSCH_SetList_r16                          []P0_PUSCH_Set_r16                          `lb:1,ub:maxNrofSRI_PUSCH_Mappings,optional`
	Olpc_ParameterSet                             *PUSCH_PowerControl_v1610_olpc_ParameterSet `lb:1,ub:2,optional`
	Sri_PUSCH_MappingToAddModList2_r17            []SRI_PUSCH_PowerControl                    `lb:1,ub:maxNrofSRI_PUSCH_Mappings,optional,ext-1`
	Sri_PUSCH_MappingToReleaseList2_r17           []SRI_PUSCH_PowerControlId                  `lb:1,ub:maxNrofSRI_PUSCH_Mappings,optional,ext-1`
	P0_PUSCH_SetList2_r17                         []P0_PUSCH_Set_r16                          `lb:1,ub:maxNrofSRI_PUSCH_Mappings,optional,ext-1`
	Dummy                                         []DummyPathlossReferenceRS_v1710            `lb:1,ub:maxNrofPUSCH_PathlossReferenceRSs_r16,optional,ext-1`
}

func (ie *PUSCH_PowerControl_v1610) Encode(w *aper.AperWriter) error {
	var err error
	hasExtensions := len(ie.Sri_PUSCH_MappingToAddModList2_r17) > 0 || len(ie.Sri_PUSCH_MappingToReleaseList2_r17) > 0 || len(ie.P0_PUSCH_SetList2_r17) > 0 || len(ie.Dummy) > 0
	preambleBits := []bool{hasExtensions, len(ie.PathlossReferenceRSToAddModListSizeExt_v1610) > 0, len(ie.PathlossReferenceRSToReleaseListSizeExt_v1610) > 0, len(ie.P0_PUSCH_SetList_r16) > 0, ie.Olpc_ParameterSet != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if len(ie.PathlossReferenceRSToAddModListSizeExt_v1610) > 0 {
		tmp_PathlossReferenceRSToAddModListSizeExt_v1610 := utils.NewSequence[*PUSCH_PathlossReferenceRS_r16]([]*PUSCH_PathlossReferenceRS_r16{}, aper.Constraint{Lb: 1, Ub: maxNrofPUSCH_PathlossReferenceRSsDiff_r16}, false)
		for _, i := range ie.PathlossReferenceRSToAddModListSizeExt_v1610 {
			tmp_PathlossReferenceRSToAddModListSizeExt_v1610.Value = append(tmp_PathlossReferenceRSToAddModListSizeExt_v1610.Value, &i)
		}
		if err = tmp_PathlossReferenceRSToAddModListSizeExt_v1610.Encode(w); err != nil {
			return utils.WrapError("Encode PathlossReferenceRSToAddModListSizeExt_v1610", err)
		}
	}
	if len(ie.PathlossReferenceRSToReleaseListSizeExt_v1610) > 0 {
		tmp_PathlossReferenceRSToReleaseListSizeExt_v1610 := utils.NewSequence[*PUSCH_PathlossReferenceRS_Id_v1610]([]*PUSCH_PathlossReferenceRS_Id_v1610{}, aper.Constraint{Lb: 1, Ub: maxNrofPUSCH_PathlossReferenceRSsDiff_r16}, false)
		for _, i := range ie.PathlossReferenceRSToReleaseListSizeExt_v1610 {
			tmp_PathlossReferenceRSToReleaseListSizeExt_v1610.Value = append(tmp_PathlossReferenceRSToReleaseListSizeExt_v1610.Value, &i)
		}
		if err = tmp_PathlossReferenceRSToReleaseListSizeExt_v1610.Encode(w); err != nil {
			return utils.WrapError("Encode PathlossReferenceRSToReleaseListSizeExt_v1610", err)
		}
	}
	if len(ie.P0_PUSCH_SetList_r16) > 0 {
		tmp_P0_PUSCH_SetList_r16 := utils.NewSequence[*P0_PUSCH_Set_r16]([]*P0_PUSCH_Set_r16{}, aper.Constraint{Lb: 1, Ub: maxNrofSRI_PUSCH_Mappings}, false)
		for _, i := range ie.P0_PUSCH_SetList_r16 {
			tmp_P0_PUSCH_SetList_r16.Value = append(tmp_P0_PUSCH_SetList_r16.Value, &i)
		}
		if err = tmp_P0_PUSCH_SetList_r16.Encode(w); err != nil {
			return utils.WrapError("Encode P0_PUSCH_SetList_r16", err)
		}
	}
	if ie.Olpc_ParameterSet != nil {
		if err = ie.Olpc_ParameterSet.Encode(w); err != nil {
			return utils.WrapError("Encode Olpc_ParameterSet", err)
		}
	}
	if hasExtensions {
		// Extension bitmap: 1 bits for 1 extension groups
		extBitmap := []bool{len(ie.Sri_PUSCH_MappingToAddModList2_r17) > 0 || len(ie.Sri_PUSCH_MappingToReleaseList2_r17) > 0 || len(ie.P0_PUSCH_SetList2_r17) > 0 || len(ie.Dummy) > 0}
		if err := w.WriteExtBitMap(extBitmap); err != nil {
			return utils.WrapError("WriteExtBitMap PUSCH_PowerControl_v1610", err)
		}

		// encode extension group 1
		if extBitmap[0] {
			extBuf := new(bytes.Buffer)
			extWriter := aper.NewWriter(extBuf)

			// Write preamble bits for optional fields in extension group 1
			optionals_ext_1 := []bool{len(ie.Sri_PUSCH_MappingToAddModList2_r17) > 0, len(ie.Sri_PUSCH_MappingToReleaseList2_r17) > 0, len(ie.P0_PUSCH_SetList2_r17) > 0, len(ie.Dummy) > 0}
			for _, bit := range optionals_ext_1 {
				if err := extWriter.WriteBool(bit); err != nil {
					return err
				}
			}

			// encode Sri_PUSCH_MappingToAddModList2_r17 optional
			if len(ie.Sri_PUSCH_MappingToAddModList2_r17) > 0 {
				tmp_Sri_PUSCH_MappingToAddModList2_r17 := utils.NewSequence[*SRI_PUSCH_PowerControl]([]*SRI_PUSCH_PowerControl{}, aper.Constraint{Lb: 1, Ub: maxNrofSRI_PUSCH_Mappings}, false)
				for _, i := range ie.Sri_PUSCH_MappingToAddModList2_r17 {
					tmp_Sri_PUSCH_MappingToAddModList2_r17.Value = append(tmp_Sri_PUSCH_MappingToAddModList2_r17.Value, &i)
				}
				if err = tmp_Sri_PUSCH_MappingToAddModList2_r17.Encode(extWriter); err != nil {
					return utils.WrapError("Encode Sri_PUSCH_MappingToAddModList2_r17", err)
				}
			}
			// encode Sri_PUSCH_MappingToReleaseList2_r17 optional
			if len(ie.Sri_PUSCH_MappingToReleaseList2_r17) > 0 {
				tmp_Sri_PUSCH_MappingToReleaseList2_r17 := utils.NewSequence[*SRI_PUSCH_PowerControlId]([]*SRI_PUSCH_PowerControlId{}, aper.Constraint{Lb: 1, Ub: maxNrofSRI_PUSCH_Mappings}, false)
				for _, i := range ie.Sri_PUSCH_MappingToReleaseList2_r17 {
					tmp_Sri_PUSCH_MappingToReleaseList2_r17.Value = append(tmp_Sri_PUSCH_MappingToReleaseList2_r17.Value, &i)
				}
				if err = tmp_Sri_PUSCH_MappingToReleaseList2_r17.Encode(extWriter); err != nil {
					return utils.WrapError("Encode Sri_PUSCH_MappingToReleaseList2_r17", err)
				}
			}
			// encode P0_PUSCH_SetList2_r17 optional
			if len(ie.P0_PUSCH_SetList2_r17) > 0 {
				tmp_P0_PUSCH_SetList2_r17 := utils.NewSequence[*P0_PUSCH_Set_r16]([]*P0_PUSCH_Set_r16{}, aper.Constraint{Lb: 1, Ub: maxNrofSRI_PUSCH_Mappings}, false)
				for _, i := range ie.P0_PUSCH_SetList2_r17 {
					tmp_P0_PUSCH_SetList2_r17.Value = append(tmp_P0_PUSCH_SetList2_r17.Value, &i)
				}
				if err = tmp_P0_PUSCH_SetList2_r17.Encode(extWriter); err != nil {
					return utils.WrapError("Encode P0_PUSCH_SetList2_r17", err)
				}
			}
			// encode Dummy optional
			if len(ie.Dummy) > 0 {
				tmp_Dummy := utils.NewSequence[*DummyPathlossReferenceRS_v1710]([]*DummyPathlossReferenceRS_v1710{}, aper.Constraint{Lb: 1, Ub: maxNrofPUSCH_PathlossReferenceRSs_r16}, false)
				for _, i := range ie.Dummy {
					tmp_Dummy.Value = append(tmp_Dummy.Value, &i)
				}
				if err = tmp_Dummy.Encode(extWriter); err != nil {
					return utils.WrapError("Encode Dummy", err)
				}
			}

			if err := extWriter.Close(); err != nil {
				return err
			}

			if err := w.WriteOpenType(extBuf.Bytes()); err != nil {
				return err
			}
		}
	}
	return nil
}

func (ie *PUSCH_PowerControl_v1610) Decode(r *aper.AperReader) error {
	var err error
	var extensionBit bool
	if extensionBit, err = r.ReadBool(); err != nil {
		return err
	}
	var PathlossReferenceRSToAddModListSizeExt_v1610Present bool
	if PathlossReferenceRSToAddModListSizeExt_v1610Present, err = r.ReadBool(); err != nil {
		return err
	}
	var PathlossReferenceRSToReleaseListSizeExt_v1610Present bool
	if PathlossReferenceRSToReleaseListSizeExt_v1610Present, err = r.ReadBool(); err != nil {
		return err
	}
	var P0_PUSCH_SetList_r16Present bool
	if P0_PUSCH_SetList_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var Olpc_ParameterSetPresent bool
	if Olpc_ParameterSetPresent, err = r.ReadBool(); err != nil {
		return err
	}
	if PathlossReferenceRSToAddModListSizeExt_v1610Present {
		tmp_PathlossReferenceRSToAddModListSizeExt_v1610 := utils.NewSequence[*PUSCH_PathlossReferenceRS_r16]([]*PUSCH_PathlossReferenceRS_r16{}, aper.Constraint{Lb: 1, Ub: maxNrofPUSCH_PathlossReferenceRSsDiff_r16}, false)
		fn_PathlossReferenceRSToAddModListSizeExt_v1610 := func() *PUSCH_PathlossReferenceRS_r16 {
			return new(PUSCH_PathlossReferenceRS_r16)
		}
		if err = tmp_PathlossReferenceRSToAddModListSizeExt_v1610.Decode(r, fn_PathlossReferenceRSToAddModListSizeExt_v1610); err != nil {
			return utils.WrapError("Decode PathlossReferenceRSToAddModListSizeExt_v1610", err)
		}
		ie.PathlossReferenceRSToAddModListSizeExt_v1610 = []PUSCH_PathlossReferenceRS_r16{}
		for _, i := range tmp_PathlossReferenceRSToAddModListSizeExt_v1610.Value {
			ie.PathlossReferenceRSToAddModListSizeExt_v1610 = append(ie.PathlossReferenceRSToAddModListSizeExt_v1610, *i)
		}
	}
	if PathlossReferenceRSToReleaseListSizeExt_v1610Present {
		tmp_PathlossReferenceRSToReleaseListSizeExt_v1610 := utils.NewSequence[*PUSCH_PathlossReferenceRS_Id_v1610]([]*PUSCH_PathlossReferenceRS_Id_v1610{}, aper.Constraint{Lb: 1, Ub: maxNrofPUSCH_PathlossReferenceRSsDiff_r16}, false)
		fn_PathlossReferenceRSToReleaseListSizeExt_v1610 := func() *PUSCH_PathlossReferenceRS_Id_v1610 {
			return new(PUSCH_PathlossReferenceRS_Id_v1610)
		}
		if err = tmp_PathlossReferenceRSToReleaseListSizeExt_v1610.Decode(r, fn_PathlossReferenceRSToReleaseListSizeExt_v1610); err != nil {
			return utils.WrapError("Decode PathlossReferenceRSToReleaseListSizeExt_v1610", err)
		}
		ie.PathlossReferenceRSToReleaseListSizeExt_v1610 = []PUSCH_PathlossReferenceRS_Id_v1610{}
		for _, i := range tmp_PathlossReferenceRSToReleaseListSizeExt_v1610.Value {
			ie.PathlossReferenceRSToReleaseListSizeExt_v1610 = append(ie.PathlossReferenceRSToReleaseListSizeExt_v1610, *i)
		}
	}
	if P0_PUSCH_SetList_r16Present {
		tmp_P0_PUSCH_SetList_r16 := utils.NewSequence[*P0_PUSCH_Set_r16]([]*P0_PUSCH_Set_r16{}, aper.Constraint{Lb: 1, Ub: maxNrofSRI_PUSCH_Mappings}, false)
		fn_P0_PUSCH_SetList_r16 := func() *P0_PUSCH_Set_r16 {
			return new(P0_PUSCH_Set_r16)
		}
		if err = tmp_P0_PUSCH_SetList_r16.Decode(r, fn_P0_PUSCH_SetList_r16); err != nil {
			return utils.WrapError("Decode P0_PUSCH_SetList_r16", err)
		}
		ie.P0_PUSCH_SetList_r16 = []P0_PUSCH_Set_r16{}
		for _, i := range tmp_P0_PUSCH_SetList_r16.Value {
			ie.P0_PUSCH_SetList_r16 = append(ie.P0_PUSCH_SetList_r16, *i)
		}
	}
	if Olpc_ParameterSetPresent {
		ie.Olpc_ParameterSet = new(PUSCH_PowerControl_v1610_olpc_ParameterSet)
		if err = ie.Olpc_ParameterSet.Decode(r); err != nil {
			return utils.WrapError("Decode Olpc_ParameterSet", err)
		}
	}

	if extensionBit {
		// Read extension bitmap: 1 bits for 1 extension groups
		extBitmap, err := r.ReadExtBitMap()
		if err != nil {
			return err
		}

		// decode extension group 1
		if len(extBitmap) > 0 && extBitmap[0] {
			extBytes, err := r.ReadOpenType()
			if err != nil {
				return err
			}

			extReader := aper.NewReader(bytes.NewReader(extBytes))

			Sri_PUSCH_MappingToAddModList2_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			Sri_PUSCH_MappingToReleaseList2_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			P0_PUSCH_SetList2_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			DummyPresent, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			// decode Sri_PUSCH_MappingToAddModList2_r17 optional
			if Sri_PUSCH_MappingToAddModList2_r17Present {
				tmp_Sri_PUSCH_MappingToAddModList2_r17 := utils.NewSequence[*SRI_PUSCH_PowerControl]([]*SRI_PUSCH_PowerControl{}, aper.Constraint{Lb: 1, Ub: maxNrofSRI_PUSCH_Mappings}, false)
				fn_Sri_PUSCH_MappingToAddModList2_r17 := func() *SRI_PUSCH_PowerControl {
					return new(SRI_PUSCH_PowerControl)
				}
				if err = tmp_Sri_PUSCH_MappingToAddModList2_r17.Decode(extReader, fn_Sri_PUSCH_MappingToAddModList2_r17); err != nil {
					return utils.WrapError("Decode Sri_PUSCH_MappingToAddModList2_r17", err)
				}
				ie.Sri_PUSCH_MappingToAddModList2_r17 = []SRI_PUSCH_PowerControl{}
				for _, i := range tmp_Sri_PUSCH_MappingToAddModList2_r17.Value {
					ie.Sri_PUSCH_MappingToAddModList2_r17 = append(ie.Sri_PUSCH_MappingToAddModList2_r17, *i)
				}
			}
			// decode Sri_PUSCH_MappingToReleaseList2_r17 optional
			if Sri_PUSCH_MappingToReleaseList2_r17Present {
				tmp_Sri_PUSCH_MappingToReleaseList2_r17 := utils.NewSequence[*SRI_PUSCH_PowerControlId]([]*SRI_PUSCH_PowerControlId{}, aper.Constraint{Lb: 1, Ub: maxNrofSRI_PUSCH_Mappings}, false)
				fn_Sri_PUSCH_MappingToReleaseList2_r17 := func() *SRI_PUSCH_PowerControlId {
					return new(SRI_PUSCH_PowerControlId)
				}
				if err = tmp_Sri_PUSCH_MappingToReleaseList2_r17.Decode(extReader, fn_Sri_PUSCH_MappingToReleaseList2_r17); err != nil {
					return utils.WrapError("Decode Sri_PUSCH_MappingToReleaseList2_r17", err)
				}
				ie.Sri_PUSCH_MappingToReleaseList2_r17 = []SRI_PUSCH_PowerControlId{}
				for _, i := range tmp_Sri_PUSCH_MappingToReleaseList2_r17.Value {
					ie.Sri_PUSCH_MappingToReleaseList2_r17 = append(ie.Sri_PUSCH_MappingToReleaseList2_r17, *i)
				}
			}
			// decode P0_PUSCH_SetList2_r17 optional
			if P0_PUSCH_SetList2_r17Present {
				tmp_P0_PUSCH_SetList2_r17 := utils.NewSequence[*P0_PUSCH_Set_r16]([]*P0_PUSCH_Set_r16{}, aper.Constraint{Lb: 1, Ub: maxNrofSRI_PUSCH_Mappings}, false)
				fn_P0_PUSCH_SetList2_r17 := func() *P0_PUSCH_Set_r16 {
					return new(P0_PUSCH_Set_r16)
				}
				if err = tmp_P0_PUSCH_SetList2_r17.Decode(extReader, fn_P0_PUSCH_SetList2_r17); err != nil {
					return utils.WrapError("Decode P0_PUSCH_SetList2_r17", err)
				}
				ie.P0_PUSCH_SetList2_r17 = []P0_PUSCH_Set_r16{}
				for _, i := range tmp_P0_PUSCH_SetList2_r17.Value {
					ie.P0_PUSCH_SetList2_r17 = append(ie.P0_PUSCH_SetList2_r17, *i)
				}
			}
			// decode Dummy optional
			if DummyPresent {
				tmp_Dummy := utils.NewSequence[*DummyPathlossReferenceRS_v1710]([]*DummyPathlossReferenceRS_v1710{}, aper.Constraint{Lb: 1, Ub: maxNrofPUSCH_PathlossReferenceRSs_r16}, false)
				fn_Dummy := func() *DummyPathlossReferenceRS_v1710 {
					return new(DummyPathlossReferenceRS_v1710)
				}
				if err = tmp_Dummy.Decode(extReader, fn_Dummy); err != nil {
					return utils.WrapError("Decode Dummy", err)
				}
				ie.Dummy = []DummyPathlossReferenceRS_v1710{}
				for _, i := range tmp_Dummy.Value {
					ie.Dummy = append(ie.Dummy, *i)
				}
			}
		}
	}
	return nil
}
