package ies

import (
	"bytes"

	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type SRS_Config struct {
	Srs_ResourceSetToReleaseList            []SRS_ResourceSetId          `lb:1,ub:maxNrofSRS_ResourceSets,optional`
	Srs_ResourceSetToAddModList             []SRS_ResourceSet            `lb:1,ub:maxNrofSRS_ResourceSets,optional`
	Srs_ResourceToReleaseList               []SRS_ResourceId             `lb:1,ub:maxNrofSRS_Resources,optional`
	Srs_ResourceToAddModList                []SRS_Resource               `lb:1,ub:maxNrofSRS_Resources,optional`
	Tpc_Accumulation                        *SRS_Config_tpc_Accumulation `optional`
	Srs_RequestDCI_1_2_r16                  *int64                       `lb:1,ub:2,optional,ext-1`
	Srs_RequestDCI_0_2_r16                  *int64                       `lb:1,ub:2,optional,ext-1`
	Srs_ResourceSetToAddModListDCI_0_2_r16  []SRS_ResourceSet            `lb:1,ub:maxNrofSRS_ResourceSets,optional,ext-1`
	Srs_ResourceSetToReleaseListDCI_0_2_r16 []SRS_ResourceSetId          `lb:1,ub:maxNrofSRS_ResourceSets,optional,ext-1`
	Srs_PosResourceSetToReleaseList_r16     []SRS_PosResourceSetId_r16   `lb:1,ub:maxNrofSRS_PosResourceSets_r16,optional,ext-1`
	Srs_PosResourceSetToAddModList_r16      []SRS_PosResourceSet_r16     `lb:1,ub:maxNrofSRS_PosResourceSets_r16,optional,ext-1`
	Srs_PosResourceToReleaseList_r16        []SRS_PosResourceId_r16      `lb:1,ub:maxNrofSRS_PosResources_r16,optional,ext-1`
	Srs_PosResourceToAddModList_r16         []SRS_PosResource_r16        `lb:1,ub:maxNrofSRS_PosResources_r16,optional,ext-1`
}

func (ie *SRS_Config) Encode(w *uper.UperWriter) error {
	var err error
	hasExtensions := ie.Srs_RequestDCI_1_2_r16 != nil || ie.Srs_RequestDCI_0_2_r16 != nil || len(ie.Srs_ResourceSetToAddModListDCI_0_2_r16) > 0 || len(ie.Srs_ResourceSetToReleaseListDCI_0_2_r16) > 0 || len(ie.Srs_PosResourceSetToReleaseList_r16) > 0 || len(ie.Srs_PosResourceSetToAddModList_r16) > 0 || len(ie.Srs_PosResourceToReleaseList_r16) > 0 || len(ie.Srs_PosResourceToAddModList_r16) > 0
	preambleBits := []bool{hasExtensions, len(ie.Srs_ResourceSetToReleaseList) > 0, len(ie.Srs_ResourceSetToAddModList) > 0, len(ie.Srs_ResourceToReleaseList) > 0, len(ie.Srs_ResourceToAddModList) > 0, ie.Tpc_Accumulation != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if len(ie.Srs_ResourceSetToReleaseList) > 0 {
		tmp_Srs_ResourceSetToReleaseList := utils.NewSequence[*SRS_ResourceSetId]([]*SRS_ResourceSetId{}, uper.Constraint{Lb: 1, Ub: maxNrofSRS_ResourceSets}, false)
		for _, i := range ie.Srs_ResourceSetToReleaseList {
			tmp_Srs_ResourceSetToReleaseList.Value = append(tmp_Srs_ResourceSetToReleaseList.Value, &i)
		}
		if err = tmp_Srs_ResourceSetToReleaseList.Encode(w); err != nil {
			return utils.WrapError("Encode Srs_ResourceSetToReleaseList", err)
		}
	}
	if len(ie.Srs_ResourceSetToAddModList) > 0 {
		tmp_Srs_ResourceSetToAddModList := utils.NewSequence[*SRS_ResourceSet]([]*SRS_ResourceSet{}, uper.Constraint{Lb: 1, Ub: maxNrofSRS_ResourceSets}, false)
		for _, i := range ie.Srs_ResourceSetToAddModList {
			tmp_Srs_ResourceSetToAddModList.Value = append(tmp_Srs_ResourceSetToAddModList.Value, &i)
		}
		if err = tmp_Srs_ResourceSetToAddModList.Encode(w); err != nil {
			return utils.WrapError("Encode Srs_ResourceSetToAddModList", err)
		}
	}
	if len(ie.Srs_ResourceToReleaseList) > 0 {
		tmp_Srs_ResourceToReleaseList := utils.NewSequence[*SRS_ResourceId]([]*SRS_ResourceId{}, uper.Constraint{Lb: 1, Ub: maxNrofSRS_Resources}, false)
		for _, i := range ie.Srs_ResourceToReleaseList {
			tmp_Srs_ResourceToReleaseList.Value = append(tmp_Srs_ResourceToReleaseList.Value, &i)
		}
		if err = tmp_Srs_ResourceToReleaseList.Encode(w); err != nil {
			return utils.WrapError("Encode Srs_ResourceToReleaseList", err)
		}
	}
	if len(ie.Srs_ResourceToAddModList) > 0 {
		tmp_Srs_ResourceToAddModList := utils.NewSequence[*SRS_Resource]([]*SRS_Resource{}, uper.Constraint{Lb: 1, Ub: maxNrofSRS_Resources}, false)
		for _, i := range ie.Srs_ResourceToAddModList {
			tmp_Srs_ResourceToAddModList.Value = append(tmp_Srs_ResourceToAddModList.Value, &i)
		}
		if err = tmp_Srs_ResourceToAddModList.Encode(w); err != nil {
			return utils.WrapError("Encode Srs_ResourceToAddModList", err)
		}
	}
	if ie.Tpc_Accumulation != nil {
		if err = ie.Tpc_Accumulation.Encode(w); err != nil {
			return utils.WrapError("Encode Tpc_Accumulation", err)
		}
	}
	if hasExtensions {
		// Extension bitmap: 1 bits for 1 extension groups
		extBitmap := []bool{ie.Srs_RequestDCI_1_2_r16 != nil || ie.Srs_RequestDCI_0_2_r16 != nil || len(ie.Srs_ResourceSetToAddModListDCI_0_2_r16) > 0 || len(ie.Srs_ResourceSetToReleaseListDCI_0_2_r16) > 0 || len(ie.Srs_PosResourceSetToReleaseList_r16) > 0 || len(ie.Srs_PosResourceSetToAddModList_r16) > 0 || len(ie.Srs_PosResourceToReleaseList_r16) > 0 || len(ie.Srs_PosResourceToAddModList_r16) > 0}
		if err := w.WriteExtBitMap(extBitmap); err != nil {
			return utils.WrapError("WriteExtBitMap SRS_Config", err)
		}

		// encode extension group 1
		if extBitmap[0] {
			extBuf := new(bytes.Buffer)
			extWriter := uper.NewWriter(extBuf)

			// Write preamble bits for optional fields in extension group 1
			optionals_ext_1 := []bool{ie.Srs_RequestDCI_1_2_r16 != nil, ie.Srs_RequestDCI_0_2_r16 != nil, len(ie.Srs_ResourceSetToAddModListDCI_0_2_r16) > 0, len(ie.Srs_ResourceSetToReleaseListDCI_0_2_r16) > 0, len(ie.Srs_PosResourceSetToReleaseList_r16) > 0, len(ie.Srs_PosResourceSetToAddModList_r16) > 0, len(ie.Srs_PosResourceToReleaseList_r16) > 0, len(ie.Srs_PosResourceToAddModList_r16) > 0}
			for _, bit := range optionals_ext_1 {
				if err := extWriter.WriteBool(bit); err != nil {
					return err
				}
			}

			// encode Srs_RequestDCI_1_2_r16 optional
			if ie.Srs_RequestDCI_1_2_r16 != nil {
				if err = extWriter.WriteInteger(*ie.Srs_RequestDCI_1_2_r16, &uper.Constraint{Lb: 1, Ub: 2}, false); err != nil {
					return utils.WrapError("Encode Srs_RequestDCI_1_2_r16", err)
				}
			}
			// encode Srs_RequestDCI_0_2_r16 optional
			if ie.Srs_RequestDCI_0_2_r16 != nil {
				if err = extWriter.WriteInteger(*ie.Srs_RequestDCI_0_2_r16, &uper.Constraint{Lb: 1, Ub: 2}, false); err != nil {
					return utils.WrapError("Encode Srs_RequestDCI_0_2_r16", err)
				}
			}
			// encode Srs_ResourceSetToAddModListDCI_0_2_r16 optional
			if len(ie.Srs_ResourceSetToAddModListDCI_0_2_r16) > 0 {
				tmp_Srs_ResourceSetToAddModListDCI_0_2_r16 := utils.NewSequence[*SRS_ResourceSet]([]*SRS_ResourceSet{}, uper.Constraint{Lb: 1, Ub: maxNrofSRS_ResourceSets}, false)
				for _, i := range ie.Srs_ResourceSetToAddModListDCI_0_2_r16 {
					tmp_Srs_ResourceSetToAddModListDCI_0_2_r16.Value = append(tmp_Srs_ResourceSetToAddModListDCI_0_2_r16.Value, &i)
				}
				if err = tmp_Srs_ResourceSetToAddModListDCI_0_2_r16.Encode(extWriter); err != nil {
					return utils.WrapError("Encode Srs_ResourceSetToAddModListDCI_0_2_r16", err)
				}
			}
			// encode Srs_ResourceSetToReleaseListDCI_0_2_r16 optional
			if len(ie.Srs_ResourceSetToReleaseListDCI_0_2_r16) > 0 {
				tmp_Srs_ResourceSetToReleaseListDCI_0_2_r16 := utils.NewSequence[*SRS_ResourceSetId]([]*SRS_ResourceSetId{}, uper.Constraint{Lb: 1, Ub: maxNrofSRS_ResourceSets}, false)
				for _, i := range ie.Srs_ResourceSetToReleaseListDCI_0_2_r16 {
					tmp_Srs_ResourceSetToReleaseListDCI_0_2_r16.Value = append(tmp_Srs_ResourceSetToReleaseListDCI_0_2_r16.Value, &i)
				}
				if err = tmp_Srs_ResourceSetToReleaseListDCI_0_2_r16.Encode(extWriter); err != nil {
					return utils.WrapError("Encode Srs_ResourceSetToReleaseListDCI_0_2_r16", err)
				}
			}
			// encode Srs_PosResourceSetToReleaseList_r16 optional
			if len(ie.Srs_PosResourceSetToReleaseList_r16) > 0 {
				tmp_Srs_PosResourceSetToReleaseList_r16 := utils.NewSequence[*SRS_PosResourceSetId_r16]([]*SRS_PosResourceSetId_r16{}, uper.Constraint{Lb: 1, Ub: maxNrofSRS_PosResourceSets_r16}, false)
				for _, i := range ie.Srs_PosResourceSetToReleaseList_r16 {
					tmp_Srs_PosResourceSetToReleaseList_r16.Value = append(tmp_Srs_PosResourceSetToReleaseList_r16.Value, &i)
				}
				if err = tmp_Srs_PosResourceSetToReleaseList_r16.Encode(extWriter); err != nil {
					return utils.WrapError("Encode Srs_PosResourceSetToReleaseList_r16", err)
				}
			}
			// encode Srs_PosResourceSetToAddModList_r16 optional
			if len(ie.Srs_PosResourceSetToAddModList_r16) > 0 {
				tmp_Srs_PosResourceSetToAddModList_r16 := utils.NewSequence[*SRS_PosResourceSet_r16]([]*SRS_PosResourceSet_r16{}, uper.Constraint{Lb: 1, Ub: maxNrofSRS_PosResourceSets_r16}, false)
				for _, i := range ie.Srs_PosResourceSetToAddModList_r16 {
					tmp_Srs_PosResourceSetToAddModList_r16.Value = append(tmp_Srs_PosResourceSetToAddModList_r16.Value, &i)
				}
				if err = tmp_Srs_PosResourceSetToAddModList_r16.Encode(extWriter); err != nil {
					return utils.WrapError("Encode Srs_PosResourceSetToAddModList_r16", err)
				}
			}
			// encode Srs_PosResourceToReleaseList_r16 optional
			if len(ie.Srs_PosResourceToReleaseList_r16) > 0 {
				tmp_Srs_PosResourceToReleaseList_r16 := utils.NewSequence[*SRS_PosResourceId_r16]([]*SRS_PosResourceId_r16{}, uper.Constraint{Lb: 1, Ub: maxNrofSRS_PosResources_r16}, false)
				for _, i := range ie.Srs_PosResourceToReleaseList_r16 {
					tmp_Srs_PosResourceToReleaseList_r16.Value = append(tmp_Srs_PosResourceToReleaseList_r16.Value, &i)
				}
				if err = tmp_Srs_PosResourceToReleaseList_r16.Encode(extWriter); err != nil {
					return utils.WrapError("Encode Srs_PosResourceToReleaseList_r16", err)
				}
			}
			// encode Srs_PosResourceToAddModList_r16 optional
			if len(ie.Srs_PosResourceToAddModList_r16) > 0 {
				tmp_Srs_PosResourceToAddModList_r16 := utils.NewSequence[*SRS_PosResource_r16]([]*SRS_PosResource_r16{}, uper.Constraint{Lb: 1, Ub: maxNrofSRS_PosResources_r16}, false)
				for _, i := range ie.Srs_PosResourceToAddModList_r16 {
					tmp_Srs_PosResourceToAddModList_r16.Value = append(tmp_Srs_PosResourceToAddModList_r16.Value, &i)
				}
				if err = tmp_Srs_PosResourceToAddModList_r16.Encode(extWriter); err != nil {
					return utils.WrapError("Encode Srs_PosResourceToAddModList_r16", err)
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

func (ie *SRS_Config) Decode(r *uper.UperReader) error {
	var err error
	var extensionBit bool
	if extensionBit, err = r.ReadBool(); err != nil {
		return err
	}
	var Srs_ResourceSetToReleaseListPresent bool
	if Srs_ResourceSetToReleaseListPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var Srs_ResourceSetToAddModListPresent bool
	if Srs_ResourceSetToAddModListPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var Srs_ResourceToReleaseListPresent bool
	if Srs_ResourceToReleaseListPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var Srs_ResourceToAddModListPresent bool
	if Srs_ResourceToAddModListPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var Tpc_AccumulationPresent bool
	if Tpc_AccumulationPresent, err = r.ReadBool(); err != nil {
		return err
	}
	if Srs_ResourceSetToReleaseListPresent {
		tmp_Srs_ResourceSetToReleaseList := utils.NewSequence[*SRS_ResourceSetId]([]*SRS_ResourceSetId{}, uper.Constraint{Lb: 1, Ub: maxNrofSRS_ResourceSets}, false)
		fn_Srs_ResourceSetToReleaseList := func() *SRS_ResourceSetId {
			return new(SRS_ResourceSetId)
		}
		if err = tmp_Srs_ResourceSetToReleaseList.Decode(r, fn_Srs_ResourceSetToReleaseList); err != nil {
			return utils.WrapError("Decode Srs_ResourceSetToReleaseList", err)
		}
		ie.Srs_ResourceSetToReleaseList = []SRS_ResourceSetId{}
		for _, i := range tmp_Srs_ResourceSetToReleaseList.Value {
			ie.Srs_ResourceSetToReleaseList = append(ie.Srs_ResourceSetToReleaseList, *i)
		}
	}
	if Srs_ResourceSetToAddModListPresent {
		tmp_Srs_ResourceSetToAddModList := utils.NewSequence[*SRS_ResourceSet]([]*SRS_ResourceSet{}, uper.Constraint{Lb: 1, Ub: maxNrofSRS_ResourceSets}, false)
		fn_Srs_ResourceSetToAddModList := func() *SRS_ResourceSet {
			return new(SRS_ResourceSet)
		}
		if err = tmp_Srs_ResourceSetToAddModList.Decode(r, fn_Srs_ResourceSetToAddModList); err != nil {
			return utils.WrapError("Decode Srs_ResourceSetToAddModList", err)
		}
		ie.Srs_ResourceSetToAddModList = []SRS_ResourceSet{}
		for _, i := range tmp_Srs_ResourceSetToAddModList.Value {
			ie.Srs_ResourceSetToAddModList = append(ie.Srs_ResourceSetToAddModList, *i)
		}
	}
	if Srs_ResourceToReleaseListPresent {
		tmp_Srs_ResourceToReleaseList := utils.NewSequence[*SRS_ResourceId]([]*SRS_ResourceId{}, uper.Constraint{Lb: 1, Ub: maxNrofSRS_Resources}, false)
		fn_Srs_ResourceToReleaseList := func() *SRS_ResourceId {
			return new(SRS_ResourceId)
		}
		if err = tmp_Srs_ResourceToReleaseList.Decode(r, fn_Srs_ResourceToReleaseList); err != nil {
			return utils.WrapError("Decode Srs_ResourceToReleaseList", err)
		}
		ie.Srs_ResourceToReleaseList = []SRS_ResourceId{}
		for _, i := range tmp_Srs_ResourceToReleaseList.Value {
			ie.Srs_ResourceToReleaseList = append(ie.Srs_ResourceToReleaseList, *i)
		}
	}
	if Srs_ResourceToAddModListPresent {
		tmp_Srs_ResourceToAddModList := utils.NewSequence[*SRS_Resource]([]*SRS_Resource{}, uper.Constraint{Lb: 1, Ub: maxNrofSRS_Resources}, false)
		fn_Srs_ResourceToAddModList := func() *SRS_Resource {
			return new(SRS_Resource)
		}
		if err = tmp_Srs_ResourceToAddModList.Decode(r, fn_Srs_ResourceToAddModList); err != nil {
			return utils.WrapError("Decode Srs_ResourceToAddModList", err)
		}
		ie.Srs_ResourceToAddModList = []SRS_Resource{}
		for _, i := range tmp_Srs_ResourceToAddModList.Value {
			ie.Srs_ResourceToAddModList = append(ie.Srs_ResourceToAddModList, *i)
		}
	}
	if Tpc_AccumulationPresent {
		ie.Tpc_Accumulation = new(SRS_Config_tpc_Accumulation)
		if err = ie.Tpc_Accumulation.Decode(r); err != nil {
			return utils.WrapError("Decode Tpc_Accumulation", err)
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

			extReader := uper.NewReader(bytes.NewReader(extBytes))

			Srs_RequestDCI_1_2_r16Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			Srs_RequestDCI_0_2_r16Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			Srs_ResourceSetToAddModListDCI_0_2_r16Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			Srs_ResourceSetToReleaseListDCI_0_2_r16Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			Srs_PosResourceSetToReleaseList_r16Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			Srs_PosResourceSetToAddModList_r16Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			Srs_PosResourceToReleaseList_r16Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			Srs_PosResourceToAddModList_r16Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			// decode Srs_RequestDCI_1_2_r16 optional
			if Srs_RequestDCI_1_2_r16Present {
				var tmp_int_Srs_RequestDCI_1_2_r16 int64
				if tmp_int_Srs_RequestDCI_1_2_r16, err = extReader.ReadInteger(&uper.Constraint{Lb: 1, Ub: 2}, false); err != nil {
					return utils.WrapError("Decode Srs_RequestDCI_1_2_r16", err)
				}
				ie.Srs_RequestDCI_1_2_r16 = &tmp_int_Srs_RequestDCI_1_2_r16
			}
			// decode Srs_RequestDCI_0_2_r16 optional
			if Srs_RequestDCI_0_2_r16Present {
				var tmp_int_Srs_RequestDCI_0_2_r16 int64
				if tmp_int_Srs_RequestDCI_0_2_r16, err = extReader.ReadInteger(&uper.Constraint{Lb: 1, Ub: 2}, false); err != nil {
					return utils.WrapError("Decode Srs_RequestDCI_0_2_r16", err)
				}
				ie.Srs_RequestDCI_0_2_r16 = &tmp_int_Srs_RequestDCI_0_2_r16
			}
			// decode Srs_ResourceSetToAddModListDCI_0_2_r16 optional
			if Srs_ResourceSetToAddModListDCI_0_2_r16Present {
				tmp_Srs_ResourceSetToAddModListDCI_0_2_r16 := utils.NewSequence[*SRS_ResourceSet]([]*SRS_ResourceSet{}, uper.Constraint{Lb: 1, Ub: maxNrofSRS_ResourceSets}, false)
				fn_Srs_ResourceSetToAddModListDCI_0_2_r16 := func() *SRS_ResourceSet {
					return new(SRS_ResourceSet)
				}
				if err = tmp_Srs_ResourceSetToAddModListDCI_0_2_r16.Decode(extReader, fn_Srs_ResourceSetToAddModListDCI_0_2_r16); err != nil {
					return utils.WrapError("Decode Srs_ResourceSetToAddModListDCI_0_2_r16", err)
				}
				ie.Srs_ResourceSetToAddModListDCI_0_2_r16 = []SRS_ResourceSet{}
				for _, i := range tmp_Srs_ResourceSetToAddModListDCI_0_2_r16.Value {
					ie.Srs_ResourceSetToAddModListDCI_0_2_r16 = append(ie.Srs_ResourceSetToAddModListDCI_0_2_r16, *i)
				}
			}
			// decode Srs_ResourceSetToReleaseListDCI_0_2_r16 optional
			if Srs_ResourceSetToReleaseListDCI_0_2_r16Present {
				tmp_Srs_ResourceSetToReleaseListDCI_0_2_r16 := utils.NewSequence[*SRS_ResourceSetId]([]*SRS_ResourceSetId{}, uper.Constraint{Lb: 1, Ub: maxNrofSRS_ResourceSets}, false)
				fn_Srs_ResourceSetToReleaseListDCI_0_2_r16 := func() *SRS_ResourceSetId {
					return new(SRS_ResourceSetId)
				}
				if err = tmp_Srs_ResourceSetToReleaseListDCI_0_2_r16.Decode(extReader, fn_Srs_ResourceSetToReleaseListDCI_0_2_r16); err != nil {
					return utils.WrapError("Decode Srs_ResourceSetToReleaseListDCI_0_2_r16", err)
				}
				ie.Srs_ResourceSetToReleaseListDCI_0_2_r16 = []SRS_ResourceSetId{}
				for _, i := range tmp_Srs_ResourceSetToReleaseListDCI_0_2_r16.Value {
					ie.Srs_ResourceSetToReleaseListDCI_0_2_r16 = append(ie.Srs_ResourceSetToReleaseListDCI_0_2_r16, *i)
				}
			}
			// decode Srs_PosResourceSetToReleaseList_r16 optional
			if Srs_PosResourceSetToReleaseList_r16Present {
				tmp_Srs_PosResourceSetToReleaseList_r16 := utils.NewSequence[*SRS_PosResourceSetId_r16]([]*SRS_PosResourceSetId_r16{}, uper.Constraint{Lb: 1, Ub: maxNrofSRS_PosResourceSets_r16}, false)
				fn_Srs_PosResourceSetToReleaseList_r16 := func() *SRS_PosResourceSetId_r16 {
					return new(SRS_PosResourceSetId_r16)
				}
				if err = tmp_Srs_PosResourceSetToReleaseList_r16.Decode(extReader, fn_Srs_PosResourceSetToReleaseList_r16); err != nil {
					return utils.WrapError("Decode Srs_PosResourceSetToReleaseList_r16", err)
				}
				ie.Srs_PosResourceSetToReleaseList_r16 = []SRS_PosResourceSetId_r16{}
				for _, i := range tmp_Srs_PosResourceSetToReleaseList_r16.Value {
					ie.Srs_PosResourceSetToReleaseList_r16 = append(ie.Srs_PosResourceSetToReleaseList_r16, *i)
				}
			}
			// decode Srs_PosResourceSetToAddModList_r16 optional
			if Srs_PosResourceSetToAddModList_r16Present {
				tmp_Srs_PosResourceSetToAddModList_r16 := utils.NewSequence[*SRS_PosResourceSet_r16]([]*SRS_PosResourceSet_r16{}, uper.Constraint{Lb: 1, Ub: maxNrofSRS_PosResourceSets_r16}, false)
				fn_Srs_PosResourceSetToAddModList_r16 := func() *SRS_PosResourceSet_r16 {
					return new(SRS_PosResourceSet_r16)
				}
				if err = tmp_Srs_PosResourceSetToAddModList_r16.Decode(extReader, fn_Srs_PosResourceSetToAddModList_r16); err != nil {
					return utils.WrapError("Decode Srs_PosResourceSetToAddModList_r16", err)
				}
				ie.Srs_PosResourceSetToAddModList_r16 = []SRS_PosResourceSet_r16{}
				for _, i := range tmp_Srs_PosResourceSetToAddModList_r16.Value {
					ie.Srs_PosResourceSetToAddModList_r16 = append(ie.Srs_PosResourceSetToAddModList_r16, *i)
				}
			}
			// decode Srs_PosResourceToReleaseList_r16 optional
			if Srs_PosResourceToReleaseList_r16Present {
				tmp_Srs_PosResourceToReleaseList_r16 := utils.NewSequence[*SRS_PosResourceId_r16]([]*SRS_PosResourceId_r16{}, uper.Constraint{Lb: 1, Ub: maxNrofSRS_PosResources_r16}, false)
				fn_Srs_PosResourceToReleaseList_r16 := func() *SRS_PosResourceId_r16 {
					return new(SRS_PosResourceId_r16)
				}
				if err = tmp_Srs_PosResourceToReleaseList_r16.Decode(extReader, fn_Srs_PosResourceToReleaseList_r16); err != nil {
					return utils.WrapError("Decode Srs_PosResourceToReleaseList_r16", err)
				}
				ie.Srs_PosResourceToReleaseList_r16 = []SRS_PosResourceId_r16{}
				for _, i := range tmp_Srs_PosResourceToReleaseList_r16.Value {
					ie.Srs_PosResourceToReleaseList_r16 = append(ie.Srs_PosResourceToReleaseList_r16, *i)
				}
			}
			// decode Srs_PosResourceToAddModList_r16 optional
			if Srs_PosResourceToAddModList_r16Present {
				tmp_Srs_PosResourceToAddModList_r16 := utils.NewSequence[*SRS_PosResource_r16]([]*SRS_PosResource_r16{}, uper.Constraint{Lb: 1, Ub: maxNrofSRS_PosResources_r16}, false)
				fn_Srs_PosResourceToAddModList_r16 := func() *SRS_PosResource_r16 {
					return new(SRS_PosResource_r16)
				}
				if err = tmp_Srs_PosResourceToAddModList_r16.Decode(extReader, fn_Srs_PosResourceToAddModList_r16); err != nil {
					return utils.WrapError("Decode Srs_PosResourceToAddModList_r16", err)
				}
				ie.Srs_PosResourceToAddModList_r16 = []SRS_PosResource_r16{}
				for _, i := range tmp_Srs_PosResourceToAddModList_r16.Value {
					ie.Srs_PosResourceToAddModList_r16 = append(ie.Srs_PosResourceToAddModList_r16, *i)
				}
			}
		}
	}
	return nil
}
