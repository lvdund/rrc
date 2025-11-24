package ies

import (
	"bytes"

	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type RACH_ConfigCommonTwoStepRA_r16 struct {
	Rach_ConfigGenericTwoStepRA_r16                    RACH_ConfigGenericTwoStepRA_r16                                                    `madatory`
	MsgA_TotalNumberOfRA_Preambles_r16                 *int64                                                                             `lb:1,ub:63,optional`
	MsgA_SSB_PerRACH_OccasionAndCB_PreamblesPerSSB_r16 *RACH_ConfigCommonTwoStepRA_r16_msgA_SSB_PerRACH_OccasionAndCB_PreamblesPerSSB_r16 `lb:1,ub:16,optional`
	MsgA_CB_PreamblesPerSSB_PerSharedRO_r16            *int64                                                                             `lb:1,ub:60,optional`
	MsgA_SSB_SharedRO_MaskIndex_r16                    *int64                                                                             `lb:1,ub:15,optional`
	GroupB_ConfiguredTwoStepRA_r16                     *GroupB_ConfiguredTwoStepRA_r16                                                    `optional`
	MsgA_PRACH_RootSequenceIndex_r16                   *RACH_ConfigCommonTwoStepRA_r16_msgA_PRACH_RootSequenceIndex_r16                   `lb:0,ub:837,optional`
	MsgA_TransMax_r16                                  *RACH_ConfigCommonTwoStepRA_r16_msgA_TransMax_r16                                  `optional`
	MsgA_RSRP_Threshold_r16                            *RSRP_Range                                                                        `optional`
	MsgA_RSRP_ThresholdSSB_r16                         *RSRP_Range                                                                        `optional`
	MsgA_SubcarrierSpacing_r16                         *SubcarrierSpacing                                                                 `optional`
	MsgA_RestrictedSetConfig_r16                       *RACH_ConfigCommonTwoStepRA_r16_msgA_RestrictedSetConfig_r16                       `optional`
	Ra_PrioritizationForAccessIdentityTwoStep_r16      *RACH_ConfigCommonTwoStepRA_r16_ra_PrioritizationForAccessIdentityTwoStep_r16      `lb:2,ub:2,optional`
	Ra_ContentionResolutionTimer_r16                   *RACH_ConfigCommonTwoStepRA_r16_ra_ContentionResolutionTimer_r16                   `optional`
	Ra_PrioritizationForSlicingTwoStep_r17             *RA_PrioritizationForSlicing_r17                                                   `optional,ext-1`
	FeatureCombinationPreamblesList_r17                []FeatureCombinationPreambles_r17                                                  `lb:1,ub:maxFeatureCombPreamblesPerRACHResource_r17,optional,ext-1`
}

func (ie *RACH_ConfigCommonTwoStepRA_r16) Encode(w *uper.UperWriter) error {
	var err error
	hasExtensions := ie.Ra_PrioritizationForSlicingTwoStep_r17 != nil || len(ie.FeatureCombinationPreamblesList_r17) > 0
	preambleBits := []bool{hasExtensions, ie.MsgA_TotalNumberOfRA_Preambles_r16 != nil, ie.MsgA_SSB_PerRACH_OccasionAndCB_PreamblesPerSSB_r16 != nil, ie.MsgA_CB_PreamblesPerSSB_PerSharedRO_r16 != nil, ie.MsgA_SSB_SharedRO_MaskIndex_r16 != nil, ie.GroupB_ConfiguredTwoStepRA_r16 != nil, ie.MsgA_PRACH_RootSequenceIndex_r16 != nil, ie.MsgA_TransMax_r16 != nil, ie.MsgA_RSRP_Threshold_r16 != nil, ie.MsgA_RSRP_ThresholdSSB_r16 != nil, ie.MsgA_SubcarrierSpacing_r16 != nil, ie.MsgA_RestrictedSetConfig_r16 != nil, ie.Ra_PrioritizationForAccessIdentityTwoStep_r16 != nil, ie.Ra_ContentionResolutionTimer_r16 != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if err = ie.Rach_ConfigGenericTwoStepRA_r16.Encode(w); err != nil {
		return utils.WrapError("Encode Rach_ConfigGenericTwoStepRA_r16", err)
	}
	if ie.MsgA_TotalNumberOfRA_Preambles_r16 != nil {
		if err = w.WriteInteger(*ie.MsgA_TotalNumberOfRA_Preambles_r16, &uper.Constraint{Lb: 1, Ub: 63}, false); err != nil {
			return utils.WrapError("Encode MsgA_TotalNumberOfRA_Preambles_r16", err)
		}
	}
	if ie.MsgA_SSB_PerRACH_OccasionAndCB_PreamblesPerSSB_r16 != nil {
		if err = ie.MsgA_SSB_PerRACH_OccasionAndCB_PreamblesPerSSB_r16.Encode(w); err != nil {
			return utils.WrapError("Encode MsgA_SSB_PerRACH_OccasionAndCB_PreamblesPerSSB_r16", err)
		}
	}
	if ie.MsgA_CB_PreamblesPerSSB_PerSharedRO_r16 != nil {
		if err = w.WriteInteger(*ie.MsgA_CB_PreamblesPerSSB_PerSharedRO_r16, &uper.Constraint{Lb: 1, Ub: 60}, false); err != nil {
			return utils.WrapError("Encode MsgA_CB_PreamblesPerSSB_PerSharedRO_r16", err)
		}
	}
	if ie.MsgA_SSB_SharedRO_MaskIndex_r16 != nil {
		if err = w.WriteInteger(*ie.MsgA_SSB_SharedRO_MaskIndex_r16, &uper.Constraint{Lb: 1, Ub: 15}, false); err != nil {
			return utils.WrapError("Encode MsgA_SSB_SharedRO_MaskIndex_r16", err)
		}
	}
	if ie.GroupB_ConfiguredTwoStepRA_r16 != nil {
		if err = ie.GroupB_ConfiguredTwoStepRA_r16.Encode(w); err != nil {
			return utils.WrapError("Encode GroupB_ConfiguredTwoStepRA_r16", err)
		}
	}
	if ie.MsgA_PRACH_RootSequenceIndex_r16 != nil {
		if err = ie.MsgA_PRACH_RootSequenceIndex_r16.Encode(w); err != nil {
			return utils.WrapError("Encode MsgA_PRACH_RootSequenceIndex_r16", err)
		}
	}
	if ie.MsgA_TransMax_r16 != nil {
		if err = ie.MsgA_TransMax_r16.Encode(w); err != nil {
			return utils.WrapError("Encode MsgA_TransMax_r16", err)
		}
	}
	if ie.MsgA_RSRP_Threshold_r16 != nil {
		if err = ie.MsgA_RSRP_Threshold_r16.Encode(w); err != nil {
			return utils.WrapError("Encode MsgA_RSRP_Threshold_r16", err)
		}
	}
	if ie.MsgA_RSRP_ThresholdSSB_r16 != nil {
		if err = ie.MsgA_RSRP_ThresholdSSB_r16.Encode(w); err != nil {
			return utils.WrapError("Encode MsgA_RSRP_ThresholdSSB_r16", err)
		}
	}
	if ie.MsgA_SubcarrierSpacing_r16 != nil {
		if err = ie.MsgA_SubcarrierSpacing_r16.Encode(w); err != nil {
			return utils.WrapError("Encode MsgA_SubcarrierSpacing_r16", err)
		}
	}
	if ie.MsgA_RestrictedSetConfig_r16 != nil {
		if err = ie.MsgA_RestrictedSetConfig_r16.Encode(w); err != nil {
			return utils.WrapError("Encode MsgA_RestrictedSetConfig_r16", err)
		}
	}
	if ie.Ra_PrioritizationForAccessIdentityTwoStep_r16 != nil {
		if err = ie.Ra_PrioritizationForAccessIdentityTwoStep_r16.Encode(w); err != nil {
			return utils.WrapError("Encode Ra_PrioritizationForAccessIdentityTwoStep_r16", err)
		}
	}
	if ie.Ra_ContentionResolutionTimer_r16 != nil {
		if err = ie.Ra_ContentionResolutionTimer_r16.Encode(w); err != nil {
			return utils.WrapError("Encode Ra_ContentionResolutionTimer_r16", err)
		}
	}
	if hasExtensions {
		// Extension bitmap: 1 bits for 1 extension groups
		extBitmap := []bool{ie.Ra_PrioritizationForSlicingTwoStep_r17 != nil || len(ie.FeatureCombinationPreamblesList_r17) > 0}
		if err := w.WriteExtBitMap(extBitmap); err != nil {
			return utils.WrapError("WriteExtBitMap RACH_ConfigCommonTwoStepRA_r16", err)
		}

		// encode extension group 1
		if extBitmap[0] {
			extBuf := new(bytes.Buffer)
			extWriter := uper.NewWriter(extBuf)

			// Write preamble bits for optional fields in extension group 1
			optionals_ext_1 := []bool{ie.Ra_PrioritizationForSlicingTwoStep_r17 != nil, len(ie.FeatureCombinationPreamblesList_r17) > 0}
			for _, bit := range optionals_ext_1 {
				if err := extWriter.WriteBool(bit); err != nil {
					return err
				}
			}

			// encode Ra_PrioritizationForSlicingTwoStep_r17 optional
			if ie.Ra_PrioritizationForSlicingTwoStep_r17 != nil {
				if err = ie.Ra_PrioritizationForSlicingTwoStep_r17.Encode(extWriter); err != nil {
					return utils.WrapError("Encode Ra_PrioritizationForSlicingTwoStep_r17", err)
				}
			}
			// encode FeatureCombinationPreamblesList_r17 optional
			if len(ie.FeatureCombinationPreamblesList_r17) > 0 {
				tmp_FeatureCombinationPreamblesList_r17 := utils.NewSequence[*FeatureCombinationPreambles_r17]([]*FeatureCombinationPreambles_r17{}, uper.Constraint{Lb: 1, Ub: maxFeatureCombPreamblesPerRACHResource_r17}, false)
				for _, i := range ie.FeatureCombinationPreamblesList_r17 {
					tmp_FeatureCombinationPreamblesList_r17.Value = append(tmp_FeatureCombinationPreamblesList_r17.Value, &i)
				}
				if err = tmp_FeatureCombinationPreamblesList_r17.Encode(extWriter); err != nil {
					return utils.WrapError("Encode FeatureCombinationPreamblesList_r17", err)
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

func (ie *RACH_ConfigCommonTwoStepRA_r16) Decode(r *uper.UperReader) error {
	var err error
	var extensionBit bool
	if extensionBit, err = r.ReadBool(); err != nil {
		return err
	}
	var MsgA_TotalNumberOfRA_Preambles_r16Present bool
	if MsgA_TotalNumberOfRA_Preambles_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var MsgA_SSB_PerRACH_OccasionAndCB_PreamblesPerSSB_r16Present bool
	if MsgA_SSB_PerRACH_OccasionAndCB_PreamblesPerSSB_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var MsgA_CB_PreamblesPerSSB_PerSharedRO_r16Present bool
	if MsgA_CB_PreamblesPerSSB_PerSharedRO_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var MsgA_SSB_SharedRO_MaskIndex_r16Present bool
	if MsgA_SSB_SharedRO_MaskIndex_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var GroupB_ConfiguredTwoStepRA_r16Present bool
	if GroupB_ConfiguredTwoStepRA_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var MsgA_PRACH_RootSequenceIndex_r16Present bool
	if MsgA_PRACH_RootSequenceIndex_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var MsgA_TransMax_r16Present bool
	if MsgA_TransMax_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var MsgA_RSRP_Threshold_r16Present bool
	if MsgA_RSRP_Threshold_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var MsgA_RSRP_ThresholdSSB_r16Present bool
	if MsgA_RSRP_ThresholdSSB_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var MsgA_SubcarrierSpacing_r16Present bool
	if MsgA_SubcarrierSpacing_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var MsgA_RestrictedSetConfig_r16Present bool
	if MsgA_RestrictedSetConfig_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var Ra_PrioritizationForAccessIdentityTwoStep_r16Present bool
	if Ra_PrioritizationForAccessIdentityTwoStep_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var Ra_ContentionResolutionTimer_r16Present bool
	if Ra_ContentionResolutionTimer_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	if err = ie.Rach_ConfigGenericTwoStepRA_r16.Decode(r); err != nil {
		return utils.WrapError("Decode Rach_ConfigGenericTwoStepRA_r16", err)
	}
	if MsgA_TotalNumberOfRA_Preambles_r16Present {
		var tmp_int_MsgA_TotalNumberOfRA_Preambles_r16 int64
		if tmp_int_MsgA_TotalNumberOfRA_Preambles_r16, err = r.ReadInteger(&uper.Constraint{Lb: 1, Ub: 63}, false); err != nil {
			return utils.WrapError("Decode MsgA_TotalNumberOfRA_Preambles_r16", err)
		}
		ie.MsgA_TotalNumberOfRA_Preambles_r16 = &tmp_int_MsgA_TotalNumberOfRA_Preambles_r16
	}
	if MsgA_SSB_PerRACH_OccasionAndCB_PreamblesPerSSB_r16Present {
		ie.MsgA_SSB_PerRACH_OccasionAndCB_PreamblesPerSSB_r16 = new(RACH_ConfigCommonTwoStepRA_r16_msgA_SSB_PerRACH_OccasionAndCB_PreamblesPerSSB_r16)
		if err = ie.MsgA_SSB_PerRACH_OccasionAndCB_PreamblesPerSSB_r16.Decode(r); err != nil {
			return utils.WrapError("Decode MsgA_SSB_PerRACH_OccasionAndCB_PreamblesPerSSB_r16", err)
		}
	}
	if MsgA_CB_PreamblesPerSSB_PerSharedRO_r16Present {
		var tmp_int_MsgA_CB_PreamblesPerSSB_PerSharedRO_r16 int64
		if tmp_int_MsgA_CB_PreamblesPerSSB_PerSharedRO_r16, err = r.ReadInteger(&uper.Constraint{Lb: 1, Ub: 60}, false); err != nil {
			return utils.WrapError("Decode MsgA_CB_PreamblesPerSSB_PerSharedRO_r16", err)
		}
		ie.MsgA_CB_PreamblesPerSSB_PerSharedRO_r16 = &tmp_int_MsgA_CB_PreamblesPerSSB_PerSharedRO_r16
	}
	if MsgA_SSB_SharedRO_MaskIndex_r16Present {
		var tmp_int_MsgA_SSB_SharedRO_MaskIndex_r16 int64
		if tmp_int_MsgA_SSB_SharedRO_MaskIndex_r16, err = r.ReadInteger(&uper.Constraint{Lb: 1, Ub: 15}, false); err != nil {
			return utils.WrapError("Decode MsgA_SSB_SharedRO_MaskIndex_r16", err)
		}
		ie.MsgA_SSB_SharedRO_MaskIndex_r16 = &tmp_int_MsgA_SSB_SharedRO_MaskIndex_r16
	}
	if GroupB_ConfiguredTwoStepRA_r16Present {
		ie.GroupB_ConfiguredTwoStepRA_r16 = new(GroupB_ConfiguredTwoStepRA_r16)
		if err = ie.GroupB_ConfiguredTwoStepRA_r16.Decode(r); err != nil {
			return utils.WrapError("Decode GroupB_ConfiguredTwoStepRA_r16", err)
		}
	}
	if MsgA_PRACH_RootSequenceIndex_r16Present {
		ie.MsgA_PRACH_RootSequenceIndex_r16 = new(RACH_ConfigCommonTwoStepRA_r16_msgA_PRACH_RootSequenceIndex_r16)
		if err = ie.MsgA_PRACH_RootSequenceIndex_r16.Decode(r); err != nil {
			return utils.WrapError("Decode MsgA_PRACH_RootSequenceIndex_r16", err)
		}
	}
	if MsgA_TransMax_r16Present {
		ie.MsgA_TransMax_r16 = new(RACH_ConfigCommonTwoStepRA_r16_msgA_TransMax_r16)
		if err = ie.MsgA_TransMax_r16.Decode(r); err != nil {
			return utils.WrapError("Decode MsgA_TransMax_r16", err)
		}
	}
	if MsgA_RSRP_Threshold_r16Present {
		ie.MsgA_RSRP_Threshold_r16 = new(RSRP_Range)
		if err = ie.MsgA_RSRP_Threshold_r16.Decode(r); err != nil {
			return utils.WrapError("Decode MsgA_RSRP_Threshold_r16", err)
		}
	}
	if MsgA_RSRP_ThresholdSSB_r16Present {
		ie.MsgA_RSRP_ThresholdSSB_r16 = new(RSRP_Range)
		if err = ie.MsgA_RSRP_ThresholdSSB_r16.Decode(r); err != nil {
			return utils.WrapError("Decode MsgA_RSRP_ThresholdSSB_r16", err)
		}
	}
	if MsgA_SubcarrierSpacing_r16Present {
		ie.MsgA_SubcarrierSpacing_r16 = new(SubcarrierSpacing)
		if err = ie.MsgA_SubcarrierSpacing_r16.Decode(r); err != nil {
			return utils.WrapError("Decode MsgA_SubcarrierSpacing_r16", err)
		}
	}
	if MsgA_RestrictedSetConfig_r16Present {
		ie.MsgA_RestrictedSetConfig_r16 = new(RACH_ConfigCommonTwoStepRA_r16_msgA_RestrictedSetConfig_r16)
		if err = ie.MsgA_RestrictedSetConfig_r16.Decode(r); err != nil {
			return utils.WrapError("Decode MsgA_RestrictedSetConfig_r16", err)
		}
	}
	if Ra_PrioritizationForAccessIdentityTwoStep_r16Present {
		ie.Ra_PrioritizationForAccessIdentityTwoStep_r16 = new(RACH_ConfigCommonTwoStepRA_r16_ra_PrioritizationForAccessIdentityTwoStep_r16)
		if err = ie.Ra_PrioritizationForAccessIdentityTwoStep_r16.Decode(r); err != nil {
			return utils.WrapError("Decode Ra_PrioritizationForAccessIdentityTwoStep_r16", err)
		}
	}
	if Ra_ContentionResolutionTimer_r16Present {
		ie.Ra_ContentionResolutionTimer_r16 = new(RACH_ConfigCommonTwoStepRA_r16_ra_ContentionResolutionTimer_r16)
		if err = ie.Ra_ContentionResolutionTimer_r16.Decode(r); err != nil {
			return utils.WrapError("Decode Ra_ContentionResolutionTimer_r16", err)
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

			Ra_PrioritizationForSlicingTwoStep_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			FeatureCombinationPreamblesList_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			// decode Ra_PrioritizationForSlicingTwoStep_r17 optional
			if Ra_PrioritizationForSlicingTwoStep_r17Present {
				ie.Ra_PrioritizationForSlicingTwoStep_r17 = new(RA_PrioritizationForSlicing_r17)
				if err = ie.Ra_PrioritizationForSlicingTwoStep_r17.Decode(extReader); err != nil {
					return utils.WrapError("Decode Ra_PrioritizationForSlicingTwoStep_r17", err)
				}
			}
			// decode FeatureCombinationPreamblesList_r17 optional
			if FeatureCombinationPreamblesList_r17Present {
				tmp_FeatureCombinationPreamblesList_r17 := utils.NewSequence[*FeatureCombinationPreambles_r17]([]*FeatureCombinationPreambles_r17{}, uper.Constraint{Lb: 1, Ub: maxFeatureCombPreamblesPerRACHResource_r17}, false)
				fn_FeatureCombinationPreamblesList_r17 := func() *FeatureCombinationPreambles_r17 {
					return new(FeatureCombinationPreambles_r17)
				}
				if err = tmp_FeatureCombinationPreamblesList_r17.Decode(extReader, fn_FeatureCombinationPreamblesList_r17); err != nil {
					return utils.WrapError("Decode FeatureCombinationPreamblesList_r17", err)
				}
				ie.FeatureCombinationPreamblesList_r17 = []FeatureCombinationPreambles_r17{}
				for _, i := range tmp_FeatureCombinationPreamblesList_r17.Value {
					ie.FeatureCombinationPreamblesList_r17 = append(ie.FeatureCombinationPreamblesList_r17, *i)
				}
			}
		}
	}
	return nil
}
