package ies

import (
	"bytes"

	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type ControlResourceSet struct {
	ControlResourceSetId          ControlResourceSetId                           `madatory`
	FrequencyDomainResources      uper.BitString                                 `lb:45,ub:45,madatory`
	Duration                      int64                                          `lb:1,ub:maxCoReSetDuration,madatory`
	Cce_REG_MappingType           *ControlResourceSet_cce_REG_MappingType        `lb:0,ub:maxNrofPhysicalResourceBlocks_1,optional`
	PrecoderGranularity           ControlResourceSet_precoderGranularity         `madatory`
	Tci_StatesPDCCH_ToAddList     []TCI_StateId                                  `lb:1,ub:maxNrofTCI_StatesPDCCH,optional`
	Tci_StatesPDCCH_ToReleaseList []TCI_StateId                                  `lb:1,ub:maxNrofTCI_StatesPDCCH,optional`
	Tci_PresentInDCI              *ControlResourceSet_tci_PresentInDCI           `optional`
	Pdcch_DMRS_ScramblingID       *int64                                         `lb:0,ub:65535,optional`
	Rb_Offset_r16                 *int64                                         `lb:0,ub:5,optional,ext-1`
	Tci_PresentDCI_1_2_r16        *int64                                         `lb:1,ub:3,optional,ext-1`
	CoresetPoolIndex_r16          *int64                                         `lb:0,ub:1,optional,ext-1`
	ControlResourceSetId_v1610    *ControlResourceSetId_v1610                    `optional,ext-1`
	FollowUnifiedTCI_State_r17    *ControlResourceSet_followUnifiedTCI_State_r17 `optional,ext-2`
}

func (ie *ControlResourceSet) Encode(w *uper.UperWriter) error {
	var err error
	hasExtensions := ie.Rb_Offset_r16 != nil || ie.Tci_PresentDCI_1_2_r16 != nil || ie.CoresetPoolIndex_r16 != nil || ie.ControlResourceSetId_v1610 != nil || ie.FollowUnifiedTCI_State_r17 != nil
	preambleBits := []bool{hasExtensions, ie.Cce_REG_MappingType != nil, len(ie.Tci_StatesPDCCH_ToAddList) > 0, len(ie.Tci_StatesPDCCH_ToReleaseList) > 0, ie.Tci_PresentInDCI != nil, ie.Pdcch_DMRS_ScramblingID != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if err = ie.ControlResourceSetId.Encode(w); err != nil {
		return utils.WrapError("Encode ControlResourceSetId", err)
	}
	if err = w.WriteBitString(ie.FrequencyDomainResources.Bytes, uint(ie.FrequencyDomainResources.NumBits), &uper.Constraint{Lb: 45, Ub: 45}, false); err != nil {
		return utils.WrapError("WriteBitString FrequencyDomainResources", err)
	}
	if err = w.WriteInteger(ie.Duration, &uper.Constraint{Lb: 1, Ub: maxCoReSetDuration}, false); err != nil {
		return utils.WrapError("WriteInteger Duration", err)
	}
	if ie.Cce_REG_MappingType != nil {
		if err = ie.Cce_REG_MappingType.Encode(w); err != nil {
			return utils.WrapError("Encode Cce_REG_MappingType", err)
		}
	}
	if err = ie.PrecoderGranularity.Encode(w); err != nil {
		return utils.WrapError("Encode PrecoderGranularity", err)
	}
	if len(ie.Tci_StatesPDCCH_ToAddList) > 0 {
		tmp_Tci_StatesPDCCH_ToAddList := utils.NewSequence[*TCI_StateId]([]*TCI_StateId{}, uper.Constraint{Lb: 1, Ub: maxNrofTCI_StatesPDCCH}, false)
		for _, i := range ie.Tci_StatesPDCCH_ToAddList {
			tmp_Tci_StatesPDCCH_ToAddList.Value = append(tmp_Tci_StatesPDCCH_ToAddList.Value, &i)
		}
		if err = tmp_Tci_StatesPDCCH_ToAddList.Encode(w); err != nil {
			return utils.WrapError("Encode Tci_StatesPDCCH_ToAddList", err)
		}
	}
	if len(ie.Tci_StatesPDCCH_ToReleaseList) > 0 {
		tmp_Tci_StatesPDCCH_ToReleaseList := utils.NewSequence[*TCI_StateId]([]*TCI_StateId{}, uper.Constraint{Lb: 1, Ub: maxNrofTCI_StatesPDCCH}, false)
		for _, i := range ie.Tci_StatesPDCCH_ToReleaseList {
			tmp_Tci_StatesPDCCH_ToReleaseList.Value = append(tmp_Tci_StatesPDCCH_ToReleaseList.Value, &i)
		}
		if err = tmp_Tci_StatesPDCCH_ToReleaseList.Encode(w); err != nil {
			return utils.WrapError("Encode Tci_StatesPDCCH_ToReleaseList", err)
		}
	}
	if ie.Tci_PresentInDCI != nil {
		if err = ie.Tci_PresentInDCI.Encode(w); err != nil {
			return utils.WrapError("Encode Tci_PresentInDCI", err)
		}
	}
	if ie.Pdcch_DMRS_ScramblingID != nil {
		if err = w.WriteInteger(*ie.Pdcch_DMRS_ScramblingID, &uper.Constraint{Lb: 0, Ub: 65535}, false); err != nil {
			return utils.WrapError("Encode Pdcch_DMRS_ScramblingID", err)
		}
	}
	if hasExtensions {
		// Extension bitmap: 2 bits for 2 extension groups
		extBitmap := []bool{ie.Rb_Offset_r16 != nil || ie.Tci_PresentDCI_1_2_r16 != nil || ie.CoresetPoolIndex_r16 != nil || ie.ControlResourceSetId_v1610 != nil, ie.FollowUnifiedTCI_State_r17 != nil}
		if err := w.WriteExtBitMap(extBitmap); err != nil {
			return utils.WrapError("WriteExtBitMap ControlResourceSet", err)
		}

		// encode extension group 1
		if extBitmap[0] {
			extBuf := new(bytes.Buffer)
			extWriter := uper.NewWriter(extBuf)

			// Write preamble bits for optional fields in extension group 1
			optionals_ext_1 := []bool{ie.Rb_Offset_r16 != nil, ie.Tci_PresentDCI_1_2_r16 != nil, ie.CoresetPoolIndex_r16 != nil, ie.ControlResourceSetId_v1610 != nil}
			for _, bit := range optionals_ext_1 {
				if err := extWriter.WriteBool(bit); err != nil {
					return err
				}
			}

			// encode Rb_Offset_r16 optional
			if ie.Rb_Offset_r16 != nil {
				if err = extWriter.WriteInteger(*ie.Rb_Offset_r16, &uper.Constraint{Lb: 0, Ub: 5}, false); err != nil {
					return utils.WrapError("Encode Rb_Offset_r16", err)
				}
			}
			// encode Tci_PresentDCI_1_2_r16 optional
			if ie.Tci_PresentDCI_1_2_r16 != nil {
				if err = extWriter.WriteInteger(*ie.Tci_PresentDCI_1_2_r16, &uper.Constraint{Lb: 1, Ub: 3}, false); err != nil {
					return utils.WrapError("Encode Tci_PresentDCI_1_2_r16", err)
				}
			}
			// encode CoresetPoolIndex_r16 optional
			if ie.CoresetPoolIndex_r16 != nil {
				if err = extWriter.WriteInteger(*ie.CoresetPoolIndex_r16, &uper.Constraint{Lb: 0, Ub: 1}, false); err != nil {
					return utils.WrapError("Encode CoresetPoolIndex_r16", err)
				}
			}
			// encode ControlResourceSetId_v1610 optional
			if ie.ControlResourceSetId_v1610 != nil {
				if err = ie.ControlResourceSetId_v1610.Encode(extWriter); err != nil {
					return utils.WrapError("Encode ControlResourceSetId_v1610", err)
				}
			}

			if err := extWriter.Close(); err != nil {
				return err
			}

			if err := w.WriteOpenType(extBuf.Bytes()); err != nil {
				return err
			}
		}

		// encode extension group 2
		if extBitmap[1] {
			extBuf := new(bytes.Buffer)
			extWriter := uper.NewWriter(extBuf)

			// Write preamble bits for optional fields in extension group 2
			optionals_ext_2 := []bool{ie.FollowUnifiedTCI_State_r17 != nil}
			for _, bit := range optionals_ext_2 {
				if err := extWriter.WriteBool(bit); err != nil {
					return err
				}
			}

			// encode FollowUnifiedTCI_State_r17 optional
			if ie.FollowUnifiedTCI_State_r17 != nil {
				if err = ie.FollowUnifiedTCI_State_r17.Encode(extWriter); err != nil {
					return utils.WrapError("Encode FollowUnifiedTCI_State_r17", err)
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

func (ie *ControlResourceSet) Decode(r *uper.UperReader) error {
	var err error
	var extensionBit bool
	if extensionBit, err = r.ReadBool(); err != nil {
		return err
	}
	var Cce_REG_MappingTypePresent bool
	if Cce_REG_MappingTypePresent, err = r.ReadBool(); err != nil {
		return err
	}
	var Tci_StatesPDCCH_ToAddListPresent bool
	if Tci_StatesPDCCH_ToAddListPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var Tci_StatesPDCCH_ToReleaseListPresent bool
	if Tci_StatesPDCCH_ToReleaseListPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var Tci_PresentInDCIPresent bool
	if Tci_PresentInDCIPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var Pdcch_DMRS_ScramblingIDPresent bool
	if Pdcch_DMRS_ScramblingIDPresent, err = r.ReadBool(); err != nil {
		return err
	}
	if err = ie.ControlResourceSetId.Decode(r); err != nil {
		return utils.WrapError("Decode ControlResourceSetId", err)
	}
	var tmp_bs_FrequencyDomainResources []byte
	var n_FrequencyDomainResources uint
	if tmp_bs_FrequencyDomainResources, n_FrequencyDomainResources, err = r.ReadBitString(&uper.Constraint{Lb: 45, Ub: 45}, false); err != nil {
		return utils.WrapError("ReadBitString FrequencyDomainResources", err)
	}
	ie.FrequencyDomainResources = uper.BitString{
		Bytes:   tmp_bs_FrequencyDomainResources,
		NumBits: uint64(n_FrequencyDomainResources),
	}
	var tmp_int_Duration int64
	if tmp_int_Duration, err = r.ReadInteger(&uper.Constraint{Lb: 1, Ub: maxCoReSetDuration}, false); err != nil {
		return utils.WrapError("ReadInteger Duration", err)
	}
	ie.Duration = tmp_int_Duration
	if Cce_REG_MappingTypePresent {
		ie.Cce_REG_MappingType = new(ControlResourceSet_cce_REG_MappingType)
		if err = ie.Cce_REG_MappingType.Decode(r); err != nil {
			return utils.WrapError("Decode Cce_REG_MappingType", err)
		}
	}
	if err = ie.PrecoderGranularity.Decode(r); err != nil {
		return utils.WrapError("Decode PrecoderGranularity", err)
	}
	if Tci_StatesPDCCH_ToAddListPresent {
		tmp_Tci_StatesPDCCH_ToAddList := utils.NewSequence[*TCI_StateId]([]*TCI_StateId{}, uper.Constraint{Lb: 1, Ub: maxNrofTCI_StatesPDCCH}, false)
		fn_Tci_StatesPDCCH_ToAddList := func() *TCI_StateId {
			return new(TCI_StateId)
		}
		if err = tmp_Tci_StatesPDCCH_ToAddList.Decode(r, fn_Tci_StatesPDCCH_ToAddList); err != nil {
			return utils.WrapError("Decode Tci_StatesPDCCH_ToAddList", err)
		}
		ie.Tci_StatesPDCCH_ToAddList = []TCI_StateId{}
		for _, i := range tmp_Tci_StatesPDCCH_ToAddList.Value {
			ie.Tci_StatesPDCCH_ToAddList = append(ie.Tci_StatesPDCCH_ToAddList, *i)
		}
	}
	if Tci_StatesPDCCH_ToReleaseListPresent {
		tmp_Tci_StatesPDCCH_ToReleaseList := utils.NewSequence[*TCI_StateId]([]*TCI_StateId{}, uper.Constraint{Lb: 1, Ub: maxNrofTCI_StatesPDCCH}, false)
		fn_Tci_StatesPDCCH_ToReleaseList := func() *TCI_StateId {
			return new(TCI_StateId)
		}
		if err = tmp_Tci_StatesPDCCH_ToReleaseList.Decode(r, fn_Tci_StatesPDCCH_ToReleaseList); err != nil {
			return utils.WrapError("Decode Tci_StatesPDCCH_ToReleaseList", err)
		}
		ie.Tci_StatesPDCCH_ToReleaseList = []TCI_StateId{}
		for _, i := range tmp_Tci_StatesPDCCH_ToReleaseList.Value {
			ie.Tci_StatesPDCCH_ToReleaseList = append(ie.Tci_StatesPDCCH_ToReleaseList, *i)
		}
	}
	if Tci_PresentInDCIPresent {
		ie.Tci_PresentInDCI = new(ControlResourceSet_tci_PresentInDCI)
		if err = ie.Tci_PresentInDCI.Decode(r); err != nil {
			return utils.WrapError("Decode Tci_PresentInDCI", err)
		}
	}
	if Pdcch_DMRS_ScramblingIDPresent {
		var tmp_int_Pdcch_DMRS_ScramblingID int64
		if tmp_int_Pdcch_DMRS_ScramblingID, err = r.ReadInteger(&uper.Constraint{Lb: 0, Ub: 65535}, false); err != nil {
			return utils.WrapError("Decode Pdcch_DMRS_ScramblingID", err)
		}
		ie.Pdcch_DMRS_ScramblingID = &tmp_int_Pdcch_DMRS_ScramblingID
	}

	if extensionBit {
		// Read extension bitmap: 2 bits for 2 extension groups
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

			Rb_Offset_r16Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			Tci_PresentDCI_1_2_r16Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			CoresetPoolIndex_r16Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			ControlResourceSetId_v1610Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			// decode Rb_Offset_r16 optional
			if Rb_Offset_r16Present {
				var tmp_int_Rb_Offset_r16 int64
				if tmp_int_Rb_Offset_r16, err = extReader.ReadInteger(&uper.Constraint{Lb: 0, Ub: 5}, false); err != nil {
					return utils.WrapError("Decode Rb_Offset_r16", err)
				}
				ie.Rb_Offset_r16 = &tmp_int_Rb_Offset_r16
			}
			// decode Tci_PresentDCI_1_2_r16 optional
			if Tci_PresentDCI_1_2_r16Present {
				var tmp_int_Tci_PresentDCI_1_2_r16 int64
				if tmp_int_Tci_PresentDCI_1_2_r16, err = extReader.ReadInteger(&uper.Constraint{Lb: 1, Ub: 3}, false); err != nil {
					return utils.WrapError("Decode Tci_PresentDCI_1_2_r16", err)
				}
				ie.Tci_PresentDCI_1_2_r16 = &tmp_int_Tci_PresentDCI_1_2_r16
			}
			// decode CoresetPoolIndex_r16 optional
			if CoresetPoolIndex_r16Present {
				var tmp_int_CoresetPoolIndex_r16 int64
				if tmp_int_CoresetPoolIndex_r16, err = extReader.ReadInteger(&uper.Constraint{Lb: 0, Ub: 1}, false); err != nil {
					return utils.WrapError("Decode CoresetPoolIndex_r16", err)
				}
				ie.CoresetPoolIndex_r16 = &tmp_int_CoresetPoolIndex_r16
			}
			// decode ControlResourceSetId_v1610 optional
			if ControlResourceSetId_v1610Present {
				ie.ControlResourceSetId_v1610 = new(ControlResourceSetId_v1610)
				if err = ie.ControlResourceSetId_v1610.Decode(extReader); err != nil {
					return utils.WrapError("Decode ControlResourceSetId_v1610", err)
				}
			}
		}
		// decode extension group 2
		if len(extBitmap) > 1 && extBitmap[1] {
			extBytes, err := r.ReadOpenType()
			if err != nil {
				return err
			}

			extReader := uper.NewReader(bytes.NewReader(extBytes))

			FollowUnifiedTCI_State_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			// decode FollowUnifiedTCI_State_r17 optional
			if FollowUnifiedTCI_State_r17Present {
				ie.FollowUnifiedTCI_State_r17 = new(ControlResourceSet_followUnifiedTCI_State_r17)
				if err = ie.FollowUnifiedTCI_State_r17.Decode(extReader); err != nil {
					return utils.WrapError("Decode FollowUnifiedTCI_State_r17", err)
				}
			}
		}
	}
	return nil
}
