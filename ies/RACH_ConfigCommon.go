package ies

import (
	"bytes"

	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

type RACH_ConfigCommon struct {
	Rach_ConfigGeneric                        RACH_ConfigGeneric                                           `madatory`
	TotalNumberOfRA_Preambles                 *int64                                                       `lb:1,ub:63,optional`
	Ssb_perRACH_OccasionAndCB_PreamblesPerSSB *RACH_ConfigCommon_ssb_perRACH_OccasionAndCB_PreamblesPerSSB `lb:1,ub:16,optional`
	GroupBconfigured                          *RACH_ConfigCommon_groupBconfigured                          `lb:1,ub:64,optional`
	Ra_ContentionResolutionTimer              RACH_ConfigCommon_ra_ContentionResolutionTimer               `madatory`
	Rsrp_ThresholdSSB                         *RSRP_Range                                                  `optional`
	Rsrp_ThresholdSSB_SUL                     *RSRP_Range                                                  `optional`
	Prach_RootSequenceIndex                   RACH_ConfigCommon_prach_RootSequenceIndex                    `lb:0,ub:837,madatory`
	Msg1_SubcarrierSpacing                    *SubcarrierSpacing                                           `optional`
	RestrictedSetConfig                       RACH_ConfigCommon_restrictedSetConfig                        `madatory`
	Msg3_transformPrecoder                    *RACH_ConfigCommon_msg3_transformPrecoder                    `optional`
	Ra_PrioritizationForAccessIdentity_r16    *RACH_ConfigCommon_ra_PrioritizationForAccessIdentity_r16    `lb:2,ub:2,optional,ext-1`
	Prach_RootSequenceIndex_r16               *RACH_ConfigCommon_prach_RootSequenceIndex_r16               `lb:0,ub:569,optional,ext-1`
	Ra_PrioritizationForSlicing_r17           *RA_PrioritizationForSlicing_r17                             `optional,ext-2`
	FeatureCombinationPreamblesList_r17       []FeatureCombinationPreambles_r17                            `lb:1,ub:maxFeatureCombPreamblesPerRACHResource_r17,optional,ext-2`
}

func (ie *RACH_ConfigCommon) Encode(w *aper.AperWriter) error {
	var err error
	hasExtensions := ie.Ra_PrioritizationForAccessIdentity_r16 != nil || ie.Prach_RootSequenceIndex_r16 != nil || ie.Ra_PrioritizationForSlicing_r17 != nil || len(ie.FeatureCombinationPreamblesList_r17) > 0
	preambleBits := []bool{hasExtensions, ie.TotalNumberOfRA_Preambles != nil, ie.Ssb_perRACH_OccasionAndCB_PreamblesPerSSB != nil, ie.GroupBconfigured != nil, ie.Rsrp_ThresholdSSB != nil, ie.Rsrp_ThresholdSSB_SUL != nil, ie.Msg1_SubcarrierSpacing != nil, ie.Msg3_transformPrecoder != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if err = ie.Rach_ConfigGeneric.Encode(w); err != nil {
		return utils.WrapError("Encode Rach_ConfigGeneric", err)
	}
	if ie.TotalNumberOfRA_Preambles != nil {
		if err = w.WriteInteger(*ie.TotalNumberOfRA_Preambles, &aper.Constraint{Lb: 1, Ub: 63}, false); err != nil {
			return utils.WrapError("Encode TotalNumberOfRA_Preambles", err)
		}
	}
	if ie.Ssb_perRACH_OccasionAndCB_PreamblesPerSSB != nil {
		if err = ie.Ssb_perRACH_OccasionAndCB_PreamblesPerSSB.Encode(w); err != nil {
			return utils.WrapError("Encode Ssb_perRACH_OccasionAndCB_PreamblesPerSSB", err)
		}
	}
	if ie.GroupBconfigured != nil {
		if err = ie.GroupBconfigured.Encode(w); err != nil {
			return utils.WrapError("Encode GroupBconfigured", err)
		}
	}
	if err = ie.Ra_ContentionResolutionTimer.Encode(w); err != nil {
		return utils.WrapError("Encode Ra_ContentionResolutionTimer", err)
	}
	if ie.Rsrp_ThresholdSSB != nil {
		if err = ie.Rsrp_ThresholdSSB.Encode(w); err != nil {
			return utils.WrapError("Encode Rsrp_ThresholdSSB", err)
		}
	}
	if ie.Rsrp_ThresholdSSB_SUL != nil {
		if err = ie.Rsrp_ThresholdSSB_SUL.Encode(w); err != nil {
			return utils.WrapError("Encode Rsrp_ThresholdSSB_SUL", err)
		}
	}
	if err = ie.Prach_RootSequenceIndex.Encode(w); err != nil {
		return utils.WrapError("Encode Prach_RootSequenceIndex", err)
	}
	if ie.Msg1_SubcarrierSpacing != nil {
		if err = ie.Msg1_SubcarrierSpacing.Encode(w); err != nil {
			return utils.WrapError("Encode Msg1_SubcarrierSpacing", err)
		}
	}
	if err = ie.RestrictedSetConfig.Encode(w); err != nil {
		return utils.WrapError("Encode RestrictedSetConfig", err)
	}
	if ie.Msg3_transformPrecoder != nil {
		if err = ie.Msg3_transformPrecoder.Encode(w); err != nil {
			return utils.WrapError("Encode Msg3_transformPrecoder", err)
		}
	}
	if hasExtensions {
		// Extension bitmap: 2 bits for 2 extension groups
		extBitmap := []bool{ie.Ra_PrioritizationForAccessIdentity_r16 != nil || ie.Prach_RootSequenceIndex_r16 != nil, ie.Ra_PrioritizationForSlicing_r17 != nil || len(ie.FeatureCombinationPreamblesList_r17) > 0}
		if err := w.WriteExtBitMap(extBitmap); err != nil {
			return utils.WrapError("WriteExtBitMap RACH_ConfigCommon", err)
		}

		// encode extension group 1
		if extBitmap[0] {
			extBuf := new(bytes.Buffer)
			extWriter := aper.NewWriter(extBuf)

			// Write preamble bits for optional fields in extension group 1
			optionals_ext_1 := []bool{ie.Ra_PrioritizationForAccessIdentity_r16 != nil, ie.Prach_RootSequenceIndex_r16 != nil}
			for _, bit := range optionals_ext_1 {
				if err := extWriter.WriteBool(bit); err != nil {
					return err
				}
			}

			// encode Ra_PrioritizationForAccessIdentity_r16 optional
			if ie.Ra_PrioritizationForAccessIdentity_r16 != nil {
				if err = ie.Ra_PrioritizationForAccessIdentity_r16.Encode(extWriter); err != nil {
					return utils.WrapError("Encode Ra_PrioritizationForAccessIdentity_r16", err)
				}
			}
			// encode Prach_RootSequenceIndex_r16 optional
			if ie.Prach_RootSequenceIndex_r16 != nil {
				if err = ie.Prach_RootSequenceIndex_r16.Encode(extWriter); err != nil {
					return utils.WrapError("Encode Prach_RootSequenceIndex_r16", err)
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
			extWriter := aper.NewWriter(extBuf)

			// Write preamble bits for optional fields in extension group 2
			optionals_ext_2 := []bool{ie.Ra_PrioritizationForSlicing_r17 != nil, len(ie.FeatureCombinationPreamblesList_r17) > 0}
			for _, bit := range optionals_ext_2 {
				if err := extWriter.WriteBool(bit); err != nil {
					return err
				}
			}

			// encode Ra_PrioritizationForSlicing_r17 optional
			if ie.Ra_PrioritizationForSlicing_r17 != nil {
				if err = ie.Ra_PrioritizationForSlicing_r17.Encode(extWriter); err != nil {
					return utils.WrapError("Encode Ra_PrioritizationForSlicing_r17", err)
				}
			}
			// encode FeatureCombinationPreamblesList_r17 optional
			if len(ie.FeatureCombinationPreamblesList_r17) > 0 {
				tmp_FeatureCombinationPreamblesList_r17 := utils.NewSequence[*FeatureCombinationPreambles_r17]([]*FeatureCombinationPreambles_r17{}, aper.Constraint{Lb: 1, Ub: maxFeatureCombPreamblesPerRACHResource_r17}, false)
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

func (ie *RACH_ConfigCommon) Decode(r *aper.AperReader) error {
	var err error
	var extensionBit bool
	if extensionBit, err = r.ReadBool(); err != nil {
		return err
	}
	var TotalNumberOfRA_PreamblesPresent bool
	if TotalNumberOfRA_PreamblesPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var Ssb_perRACH_OccasionAndCB_PreamblesPerSSBPresent bool
	if Ssb_perRACH_OccasionAndCB_PreamblesPerSSBPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var GroupBconfiguredPresent bool
	if GroupBconfiguredPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var Rsrp_ThresholdSSBPresent bool
	if Rsrp_ThresholdSSBPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var Rsrp_ThresholdSSB_SULPresent bool
	if Rsrp_ThresholdSSB_SULPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var Msg1_SubcarrierSpacingPresent bool
	if Msg1_SubcarrierSpacingPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var Msg3_transformPrecoderPresent bool
	if Msg3_transformPrecoderPresent, err = r.ReadBool(); err != nil {
		return err
	}
	if err = ie.Rach_ConfigGeneric.Decode(r); err != nil {
		return utils.WrapError("Decode Rach_ConfigGeneric", err)
	}
	if TotalNumberOfRA_PreamblesPresent {
		var tmp_int_TotalNumberOfRA_Preambles int64
		if tmp_int_TotalNumberOfRA_Preambles, err = r.ReadInteger(&aper.Constraint{Lb: 1, Ub: 63}, false); err != nil {
			return utils.WrapError("Decode TotalNumberOfRA_Preambles", err)
		}
		ie.TotalNumberOfRA_Preambles = &tmp_int_TotalNumberOfRA_Preambles
	}
	if Ssb_perRACH_OccasionAndCB_PreamblesPerSSBPresent {
		ie.Ssb_perRACH_OccasionAndCB_PreamblesPerSSB = new(RACH_ConfigCommon_ssb_perRACH_OccasionAndCB_PreamblesPerSSB)
		if err = ie.Ssb_perRACH_OccasionAndCB_PreamblesPerSSB.Decode(r); err != nil {
			return utils.WrapError("Decode Ssb_perRACH_OccasionAndCB_PreamblesPerSSB", err)
		}
	}
	if GroupBconfiguredPresent {
		ie.GroupBconfigured = new(RACH_ConfigCommon_groupBconfigured)
		if err = ie.GroupBconfigured.Decode(r); err != nil {
			return utils.WrapError("Decode GroupBconfigured", err)
		}
	}
	if err = ie.Ra_ContentionResolutionTimer.Decode(r); err != nil {
		return utils.WrapError("Decode Ra_ContentionResolutionTimer", err)
	}
	if Rsrp_ThresholdSSBPresent {
		ie.Rsrp_ThresholdSSB = new(RSRP_Range)
		if err = ie.Rsrp_ThresholdSSB.Decode(r); err != nil {
			return utils.WrapError("Decode Rsrp_ThresholdSSB", err)
		}
	}
	if Rsrp_ThresholdSSB_SULPresent {
		ie.Rsrp_ThresholdSSB_SUL = new(RSRP_Range)
		if err = ie.Rsrp_ThresholdSSB_SUL.Decode(r); err != nil {
			return utils.WrapError("Decode Rsrp_ThresholdSSB_SUL", err)
		}
	}
	if err = ie.Prach_RootSequenceIndex.Decode(r); err != nil {
		return utils.WrapError("Decode Prach_RootSequenceIndex", err)
	}
	if Msg1_SubcarrierSpacingPresent {
		ie.Msg1_SubcarrierSpacing = new(SubcarrierSpacing)
		if err = ie.Msg1_SubcarrierSpacing.Decode(r); err != nil {
			return utils.WrapError("Decode Msg1_SubcarrierSpacing", err)
		}
	}
	if err = ie.RestrictedSetConfig.Decode(r); err != nil {
		return utils.WrapError("Decode RestrictedSetConfig", err)
	}
	if Msg3_transformPrecoderPresent {
		ie.Msg3_transformPrecoder = new(RACH_ConfigCommon_msg3_transformPrecoder)
		if err = ie.Msg3_transformPrecoder.Decode(r); err != nil {
			return utils.WrapError("Decode Msg3_transformPrecoder", err)
		}
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

			extReader := aper.NewReader(bytes.NewReader(extBytes))

			Ra_PrioritizationForAccessIdentity_r16Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			Prach_RootSequenceIndex_r16Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			// decode Ra_PrioritizationForAccessIdentity_r16 optional
			if Ra_PrioritizationForAccessIdentity_r16Present {
				ie.Ra_PrioritizationForAccessIdentity_r16 = new(RACH_ConfigCommon_ra_PrioritizationForAccessIdentity_r16)
				if err = ie.Ra_PrioritizationForAccessIdentity_r16.Decode(extReader); err != nil {
					return utils.WrapError("Decode Ra_PrioritizationForAccessIdentity_r16", err)
				}
			}
			// decode Prach_RootSequenceIndex_r16 optional
			if Prach_RootSequenceIndex_r16Present {
				ie.Prach_RootSequenceIndex_r16 = new(RACH_ConfigCommon_prach_RootSequenceIndex_r16)
				if err = ie.Prach_RootSequenceIndex_r16.Decode(extReader); err != nil {
					return utils.WrapError("Decode Prach_RootSequenceIndex_r16", err)
				}
			}
		}
		// decode extension group 2
		if len(extBitmap) > 1 && extBitmap[1] {
			extBytes, err := r.ReadOpenType()
			if err != nil {
				return err
			}

			extReader := aper.NewReader(bytes.NewReader(extBytes))

			Ra_PrioritizationForSlicing_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			FeatureCombinationPreamblesList_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			// decode Ra_PrioritizationForSlicing_r17 optional
			if Ra_PrioritizationForSlicing_r17Present {
				ie.Ra_PrioritizationForSlicing_r17 = new(RA_PrioritizationForSlicing_r17)
				if err = ie.Ra_PrioritizationForSlicing_r17.Decode(extReader); err != nil {
					return utils.WrapError("Decode Ra_PrioritizationForSlicing_r17", err)
				}
			}
			// decode FeatureCombinationPreamblesList_r17 optional
			if FeatureCombinationPreamblesList_r17Present {
				tmp_FeatureCombinationPreamblesList_r17 := utils.NewSequence[*FeatureCombinationPreambles_r17]([]*FeatureCombinationPreambles_r17{}, aper.Constraint{Lb: 1, Ub: maxFeatureCombPreamblesPerRACHResource_r17}, false)
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
