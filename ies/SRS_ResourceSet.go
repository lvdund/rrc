package ies

import (
	"bytes"

	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type SRS_ResourceSet struct {
	Srs_ResourceSetId                SRS_ResourceSetId                                 `madatory`
	Srs_ResourceIdList               []SRS_ResourceId                                  `lb:1,ub:maxNrofSRS_ResourcesPerSet,optional`
	ResourceType                     []int64                                           `lb:1,ub:maxNrofSRS_TriggerStates_2,e_lb:1,e_ub:maxNrofSRS_TriggerStates_1,optional`
	Usage                            SRS_ResourceSet_usage                             `madatory,ext`
	Alpha                            *Alpha                                            `optional,ext`
	P0                               *int64                                            `lb:-202,ub:24,optional,ext`
	PathlossReferenceRS              *PathlossReferenceRS_Config                       `optional,ext`
	Srs_PowerControlAdjustmentStates *SRS_ResourceSet_srs_PowerControlAdjustmentStates `optional,ext`
	PathlossReferenceRSList_r16      *PathlossReferenceRSList_r16                      `optional,ext-2,setuprelease`
	UsagePDC_r17                     *SRS_ResourceSet_usagePDC_r17                     `optional,ext-3`
	AvailableSlotOffsetList_r17      []AvailableSlotOffset_r17                         `lb:1,ub:4,optional,ext-3`
	FollowUnifiedTCI_StateSRS_r17    *SRS_ResourceSet_followUnifiedTCI_StateSRS_r17    `optional,ext-3`
}

func (ie *SRS_ResourceSet) Encode(w *uper.UperWriter) error {
	var err error
	hasExtensions := ie.PathlossReferenceRSList_r16 != nil || ie.UsagePDC_r17 != nil || len(ie.AvailableSlotOffsetList_r17) > 0 || ie.FollowUnifiedTCI_StateSRS_r17 != nil
	preambleBits := []bool{hasExtensions, len(ie.Srs_ResourceIdList) > 0, len(ie.ResourceType) > 0}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if err = ie.Srs_ResourceSetId.Encode(w); err != nil {
		return utils.WrapError("Encode Srs_ResourceSetId", err)
	}
	if len(ie.Srs_ResourceIdList) > 0 {
		tmp_Srs_ResourceIdList := utils.NewSequence[*SRS_ResourceId]([]*SRS_ResourceId{}, uper.Constraint{Lb: 1, Ub: maxNrofSRS_ResourcesPerSet}, false)
		for _, i := range ie.Srs_ResourceIdList {
			tmp_Srs_ResourceIdList.Value = append(tmp_Srs_ResourceIdList.Value, &i)
		}
		if err = tmp_Srs_ResourceIdList.Encode(w); err != nil {
			return utils.WrapError("Encode Srs_ResourceIdList", err)
		}
	}
	if len(ie.ResourceType) > 0 {
		tmp_ResourceType := utils.NewSequence[*utils.INTEGER]([]*utils.INTEGER{}, uper.Constraint{Lb: 1, Ub: maxNrofSRS_TriggerStates_2}, false)
		for _, i := range ie.ResourceType {
			tmp_ie := utils.NewINTEGER(int64(i), uper.Constraint{Lb: 1, Ub: maxNrofSRS_TriggerStates_1}, false)
			tmp_ResourceType.Value = append(tmp_ResourceType.Value, &tmp_ie)
		}
		if err = tmp_ResourceType.Encode(w); err != nil {
			return utils.WrapError("Encode ResourceType", err)
		}
	}
	if hasExtensions {
		// Extension bitmap: 2 bits for 2 extension groups
		extBitmap := []bool{ie.PathlossReferenceRSList_r16 != nil, ie.UsagePDC_r17 != nil || len(ie.AvailableSlotOffsetList_r17) > 0 || ie.FollowUnifiedTCI_StateSRS_r17 != nil}
		if err := w.WriteExtBitMap(extBitmap); err != nil {
			return utils.WrapError("WriteExtBitMap SRS_ResourceSet", err)
		}

		// encode extension group 2
		if extBitmap[1] {
			extBuf := new(bytes.Buffer)
			extWriter := uper.NewWriter(extBuf)

			// Write preamble bits for optional fields in extension group 2
			optionals_ext_2 := []bool{ie.PathlossReferenceRSList_r16 != nil}
			for _, bit := range optionals_ext_2 {
				if err := extWriter.WriteBool(bit); err != nil {
					return err
				}
			}

			// encode PathlossReferenceRSList_r16 optional
			if ie.PathlossReferenceRSList_r16 != nil {
				tmp_PathlossReferenceRSList_r16 := utils.SetupRelease[*PathlossReferenceRSList_r16]{
					Setup: ie.PathlossReferenceRSList_r16,
				}
				if err = tmp_PathlossReferenceRSList_r16.Encode(extWriter); err != nil {
					return utils.WrapError("Encode PathlossReferenceRSList_r16", err)
				}
			}

			if err := extWriter.Close(); err != nil {
				return err
			}

			if err := w.WriteOpenType(extBuf.Bytes()); err != nil {
				return err
			}
		}

		// encode extension group 3
		if extBitmap[2] {
			extBuf := new(bytes.Buffer)
			extWriter := uper.NewWriter(extBuf)

			// Write preamble bits for optional fields in extension group 3
			optionals_ext_3 := []bool{ie.UsagePDC_r17 != nil, len(ie.AvailableSlotOffsetList_r17) > 0, ie.FollowUnifiedTCI_StateSRS_r17 != nil}
			for _, bit := range optionals_ext_3 {
				if err := extWriter.WriteBool(bit); err != nil {
					return err
				}
			}

			// encode UsagePDC_r17 optional
			if ie.UsagePDC_r17 != nil {
				if err = ie.UsagePDC_r17.Encode(extWriter); err != nil {
					return utils.WrapError("Encode UsagePDC_r17", err)
				}
			}
			// encode AvailableSlotOffsetList_r17 optional
			if len(ie.AvailableSlotOffsetList_r17) > 0 {
				tmp_AvailableSlotOffsetList_r17 := utils.NewSequence[*AvailableSlotOffset_r17]([]*AvailableSlotOffset_r17{}, uper.Constraint{Lb: 1, Ub: 4}, false)
				for _, i := range ie.AvailableSlotOffsetList_r17 {
					tmp_AvailableSlotOffsetList_r17.Value = append(tmp_AvailableSlotOffsetList_r17.Value, &i)
				}
				if err = tmp_AvailableSlotOffsetList_r17.Encode(extWriter); err != nil {
					return utils.WrapError("Encode AvailableSlotOffsetList_r17", err)
				}
			}
			// encode FollowUnifiedTCI_StateSRS_r17 optional
			if ie.FollowUnifiedTCI_StateSRS_r17 != nil {
				if err = ie.FollowUnifiedTCI_StateSRS_r17.Encode(extWriter); err != nil {
					return utils.WrapError("Encode FollowUnifiedTCI_StateSRS_r17", err)
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

func (ie *SRS_ResourceSet) Decode(r *uper.UperReader) error {
	var err error
	var extensionBit bool
	if extensionBit, err = r.ReadBool(); err != nil {
		return err
	}
	var Srs_ResourceIdListPresent bool
	if Srs_ResourceIdListPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var ResourceTypePresent bool
	if ResourceTypePresent, err = r.ReadBool(); err != nil {
		return err
	}
	if err = ie.Srs_ResourceSetId.Decode(r); err != nil {
		return utils.WrapError("Decode Srs_ResourceSetId", err)
	}
	if Srs_ResourceIdListPresent {
		tmp_Srs_ResourceIdList := utils.NewSequence[*SRS_ResourceId]([]*SRS_ResourceId{}, uper.Constraint{Lb: 1, Ub: maxNrofSRS_ResourcesPerSet}, false)
		fn_Srs_ResourceIdList := func() *SRS_ResourceId {
			return new(SRS_ResourceId)
		}
		if err = tmp_Srs_ResourceIdList.Decode(r, fn_Srs_ResourceIdList); err != nil {
			return utils.WrapError("Decode Srs_ResourceIdList", err)
		}
		ie.Srs_ResourceIdList = []SRS_ResourceId{}
		for _, i := range tmp_Srs_ResourceIdList.Value {
			ie.Srs_ResourceIdList = append(ie.Srs_ResourceIdList, *i)
		}
	}
	if ResourceTypePresent {
		tmp_ResourceType := utils.NewSequence[*utils.INTEGER]([]*utils.INTEGER{}, uper.Constraint{Lb: 1, Ub: maxNrofSRS_TriggerStates_2}, false)
		fn_ResourceType := func() *utils.INTEGER {
			ie := utils.NewINTEGER(0, uper.Constraint{Lb: 1, Ub: maxNrofSRS_TriggerStates_1}, false)
			return &ie
		}
		if err = tmp_ResourceType.Decode(r, fn_ResourceType); err != nil {
			return utils.WrapError("Decode ResourceType", err)
		}
		ie.ResourceType = []int64{}
		for _, i := range tmp_ResourceType.Value {
			ie.ResourceType = append(ie.ResourceType, int64(i.Value))
		}
	}

	if extensionBit {
		// Read extension bitmap: 2 bits for 2 extension groups
		extBitmap, err := r.ReadExtBitMap()
		if err != nil {
			return err
		}

		// decode extension group 2
		if len(extBitmap) > 1 && extBitmap[1] {
			extBytes, err := r.ReadOpenType()
			if err != nil {
				return err
			}

			extReader := uper.NewReader(bytes.NewReader(extBytes))

			PathlossReferenceRSList_r16Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			// decode PathlossReferenceRSList_r16 optional
			if PathlossReferenceRSList_r16Present {
				tmp_PathlossReferenceRSList_r16 := utils.SetupRelease[*PathlossReferenceRSList_r16]{}
				if err = tmp_PathlossReferenceRSList_r16.Decode(extReader); err != nil {
					return utils.WrapError("Decode PathlossReferenceRSList_r16", err)
				}
				ie.PathlossReferenceRSList_r16 = tmp_PathlossReferenceRSList_r16.Setup
			}
		}
		// decode extension group 3
		if len(extBitmap) > 2 && extBitmap[2] {
			extBytes, err := r.ReadOpenType()
			if err != nil {
				return err
			}

			extReader := uper.NewReader(bytes.NewReader(extBytes))

			UsagePDC_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			AvailableSlotOffsetList_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			FollowUnifiedTCI_StateSRS_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			// decode UsagePDC_r17 optional
			if UsagePDC_r17Present {
				ie.UsagePDC_r17 = new(SRS_ResourceSet_usagePDC_r17)
				if err = ie.UsagePDC_r17.Decode(extReader); err != nil {
					return utils.WrapError("Decode UsagePDC_r17", err)
				}
			}
			// decode AvailableSlotOffsetList_r17 optional
			if AvailableSlotOffsetList_r17Present {
				tmp_AvailableSlotOffsetList_r17 := utils.NewSequence[*AvailableSlotOffset_r17]([]*AvailableSlotOffset_r17{}, uper.Constraint{Lb: 1, Ub: 4}, false)
				fn_AvailableSlotOffsetList_r17 := func() *AvailableSlotOffset_r17 {
					return new(AvailableSlotOffset_r17)
				}
				if err = tmp_AvailableSlotOffsetList_r17.Decode(extReader, fn_AvailableSlotOffsetList_r17); err != nil {
					return utils.WrapError("Decode AvailableSlotOffsetList_r17", err)
				}
				ie.AvailableSlotOffsetList_r17 = []AvailableSlotOffset_r17{}
				for _, i := range tmp_AvailableSlotOffsetList_r17.Value {
					ie.AvailableSlotOffsetList_r17 = append(ie.AvailableSlotOffsetList_r17, *i)
				}
			}
			// decode FollowUnifiedTCI_StateSRS_r17 optional
			if FollowUnifiedTCI_StateSRS_r17Present {
				ie.FollowUnifiedTCI_StateSRS_r17 = new(SRS_ResourceSet_followUnifiedTCI_StateSRS_r17)
				if err = ie.FollowUnifiedTCI_StateSRS_r17.Decode(extReader); err != nil {
					return utils.WrapError("Decode FollowUnifiedTCI_StateSRS_r17", err)
				}
			}
		}
	}
	return nil
}
