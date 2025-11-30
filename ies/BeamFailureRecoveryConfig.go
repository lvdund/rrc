package ies

import (
	"bytes"

	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

type BeamFailureRecoveryConfig struct {
	RootSequenceIndex_BFR        *int64                                              `lb:0,ub:137,optional`
	Rach_ConfigBFR               *RACH_ConfigGeneric                                 `optional`
	Rsrp_ThresholdSSB            *RSRP_Range                                         `optional`
	CandidateBeamRSList          []PRACH_ResourceDedicatedBFR                        `lb:1,ub:maxNrofCandidateBeams,optional`
	Ssb_perRACH_Occasion         *BeamFailureRecoveryConfig_ssb_perRACH_Occasion     `optional`
	Ra_ssb_OccasionMaskIndex     *int64                                              `lb:0,ub:15,optional`
	RecoverySearchSpaceId        *SearchSpaceId                                      `optional`
	Ra_Prioritization            *RA_Prioritization                                  `optional`
	BeamFailureRecoveryTimer     *BeamFailureRecoveryConfig_beamFailureRecoveryTimer `optional`
	Msg1_SubcarrierSpacing       *SubcarrierSpacing                                  `optional,ext-1`
	Ra_PrioritizationTwoStep_r16 *RA_Prioritization                                  `optional,ext-2`
	CandidateBeamRSListExt_v1610 *CandidateBeamRSListExt_r16                         `optional,ext-2,setuprelease`
	SpCell_BFR_CBRA_r16          *BeamFailureRecoveryConfig_spCell_BFR_CBRA_r16      `optional,ext-3`
}

func (ie *BeamFailureRecoveryConfig) Encode(w *aper.AperWriter) error {
	var err error
	hasExtensions := ie.Msg1_SubcarrierSpacing != nil || ie.Ra_PrioritizationTwoStep_r16 != nil || ie.CandidateBeamRSListExt_v1610 != nil || ie.SpCell_BFR_CBRA_r16 != nil
	preambleBits := []bool{hasExtensions, ie.RootSequenceIndex_BFR != nil, ie.Rach_ConfigBFR != nil, ie.Rsrp_ThresholdSSB != nil, len(ie.CandidateBeamRSList) > 0, ie.Ssb_perRACH_Occasion != nil, ie.Ra_ssb_OccasionMaskIndex != nil, ie.RecoverySearchSpaceId != nil, ie.Ra_Prioritization != nil, ie.BeamFailureRecoveryTimer != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.RootSequenceIndex_BFR != nil {
		if err = w.WriteInteger(*ie.RootSequenceIndex_BFR, &aper.Constraint{Lb: 0, Ub: 137}, false); err != nil {
			return utils.WrapError("Encode RootSequenceIndex_BFR", err)
		}
	}
	if ie.Rach_ConfigBFR != nil {
		if err = ie.Rach_ConfigBFR.Encode(w); err != nil {
			return utils.WrapError("Encode Rach_ConfigBFR", err)
		}
	}
	if ie.Rsrp_ThresholdSSB != nil {
		if err = ie.Rsrp_ThresholdSSB.Encode(w); err != nil {
			return utils.WrapError("Encode Rsrp_ThresholdSSB", err)
		}
	}
	if len(ie.CandidateBeamRSList) > 0 {
		tmp_CandidateBeamRSList := utils.NewSequence[*PRACH_ResourceDedicatedBFR]([]*PRACH_ResourceDedicatedBFR{}, aper.Constraint{Lb: 1, Ub: maxNrofCandidateBeams}, false)
		for _, i := range ie.CandidateBeamRSList {
			tmp_CandidateBeamRSList.Value = append(tmp_CandidateBeamRSList.Value, &i)
		}
		if err = tmp_CandidateBeamRSList.Encode(w); err != nil {
			return utils.WrapError("Encode CandidateBeamRSList", err)
		}
	}
	if ie.Ssb_perRACH_Occasion != nil {
		if err = ie.Ssb_perRACH_Occasion.Encode(w); err != nil {
			return utils.WrapError("Encode Ssb_perRACH_Occasion", err)
		}
	}
	if ie.Ra_ssb_OccasionMaskIndex != nil {
		if err = w.WriteInteger(*ie.Ra_ssb_OccasionMaskIndex, &aper.Constraint{Lb: 0, Ub: 15}, false); err != nil {
			return utils.WrapError("Encode Ra_ssb_OccasionMaskIndex", err)
		}
	}
	if ie.RecoverySearchSpaceId != nil {
		if err = ie.RecoverySearchSpaceId.Encode(w); err != nil {
			return utils.WrapError("Encode RecoverySearchSpaceId", err)
		}
	}
	if ie.Ra_Prioritization != nil {
		if err = ie.Ra_Prioritization.Encode(w); err != nil {
			return utils.WrapError("Encode Ra_Prioritization", err)
		}
	}
	if ie.BeamFailureRecoveryTimer != nil {
		if err = ie.BeamFailureRecoveryTimer.Encode(w); err != nil {
			return utils.WrapError("Encode BeamFailureRecoveryTimer", err)
		}
	}
	if hasExtensions {
		// Extension bitmap: 3 bits for 3 extension groups
		extBitmap := []bool{ie.Msg1_SubcarrierSpacing != nil, ie.Ra_PrioritizationTwoStep_r16 != nil || ie.CandidateBeamRSListExt_v1610 != nil, ie.SpCell_BFR_CBRA_r16 != nil}
		if err := w.WriteExtBitMap(extBitmap); err != nil {
			return utils.WrapError("WriteExtBitMap BeamFailureRecoveryConfig", err)
		}

		// encode extension group 1
		if extBitmap[0] {
			extBuf := new(bytes.Buffer)
			extWriter := aper.NewWriter(extBuf)

			// Write preamble bits for optional fields in extension group 1
			optionals_ext_1 := []bool{ie.Msg1_SubcarrierSpacing != nil}
			for _, bit := range optionals_ext_1 {
				if err := extWriter.WriteBool(bit); err != nil {
					return err
				}
			}

			// encode Msg1_SubcarrierSpacing optional
			if ie.Msg1_SubcarrierSpacing != nil {
				if err = ie.Msg1_SubcarrierSpacing.Encode(extWriter); err != nil {
					return utils.WrapError("Encode Msg1_SubcarrierSpacing", err)
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
			optionals_ext_2 := []bool{ie.Ra_PrioritizationTwoStep_r16 != nil, ie.CandidateBeamRSListExt_v1610 != nil}
			for _, bit := range optionals_ext_2 {
				if err := extWriter.WriteBool(bit); err != nil {
					return err
				}
			}

			// encode Ra_PrioritizationTwoStep_r16 optional
			if ie.Ra_PrioritizationTwoStep_r16 != nil {
				if err = ie.Ra_PrioritizationTwoStep_r16.Encode(extWriter); err != nil {
					return utils.WrapError("Encode Ra_PrioritizationTwoStep_r16", err)
				}
			}
			// encode CandidateBeamRSListExt_v1610 optional
			if ie.CandidateBeamRSListExt_v1610 != nil {
				tmp_CandidateBeamRSListExt_v1610 := utils.SetupRelease[*CandidateBeamRSListExt_r16]{
					Setup: ie.CandidateBeamRSListExt_v1610,
				}
				if err = tmp_CandidateBeamRSListExt_v1610.Encode(extWriter); err != nil {
					return utils.WrapError("Encode CandidateBeamRSListExt_v1610", err)
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
			extWriter := aper.NewWriter(extBuf)

			// Write preamble bits for optional fields in extension group 3
			optionals_ext_3 := []bool{ie.SpCell_BFR_CBRA_r16 != nil}
			for _, bit := range optionals_ext_3 {
				if err := extWriter.WriteBool(bit); err != nil {
					return err
				}
			}

			// encode SpCell_BFR_CBRA_r16 optional
			if ie.SpCell_BFR_CBRA_r16 != nil {
				if err = ie.SpCell_BFR_CBRA_r16.Encode(extWriter); err != nil {
					return utils.WrapError("Encode SpCell_BFR_CBRA_r16", err)
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

func (ie *BeamFailureRecoveryConfig) Decode(r *aper.AperReader) error {
	var err error
	var extensionBit bool
	if extensionBit, err = r.ReadBool(); err != nil {
		return err
	}
	var RootSequenceIndex_BFRPresent bool
	if RootSequenceIndex_BFRPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var Rach_ConfigBFRPresent bool
	if Rach_ConfigBFRPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var Rsrp_ThresholdSSBPresent bool
	if Rsrp_ThresholdSSBPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var CandidateBeamRSListPresent bool
	if CandidateBeamRSListPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var Ssb_perRACH_OccasionPresent bool
	if Ssb_perRACH_OccasionPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var Ra_ssb_OccasionMaskIndexPresent bool
	if Ra_ssb_OccasionMaskIndexPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var RecoverySearchSpaceIdPresent bool
	if RecoverySearchSpaceIdPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var Ra_PrioritizationPresent bool
	if Ra_PrioritizationPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var BeamFailureRecoveryTimerPresent bool
	if BeamFailureRecoveryTimerPresent, err = r.ReadBool(); err != nil {
		return err
	}
	if RootSequenceIndex_BFRPresent {
		var tmp_int_RootSequenceIndex_BFR int64
		if tmp_int_RootSequenceIndex_BFR, err = r.ReadInteger(&aper.Constraint{Lb: 0, Ub: 137}, false); err != nil {
			return utils.WrapError("Decode RootSequenceIndex_BFR", err)
		}
		ie.RootSequenceIndex_BFR = &tmp_int_RootSequenceIndex_BFR
	}
	if Rach_ConfigBFRPresent {
		ie.Rach_ConfigBFR = new(RACH_ConfigGeneric)
		if err = ie.Rach_ConfigBFR.Decode(r); err != nil {
			return utils.WrapError("Decode Rach_ConfigBFR", err)
		}
	}
	if Rsrp_ThresholdSSBPresent {
		ie.Rsrp_ThresholdSSB = new(RSRP_Range)
		if err = ie.Rsrp_ThresholdSSB.Decode(r); err != nil {
			return utils.WrapError("Decode Rsrp_ThresholdSSB", err)
		}
	}
	if CandidateBeamRSListPresent {
		tmp_CandidateBeamRSList := utils.NewSequence[*PRACH_ResourceDedicatedBFR]([]*PRACH_ResourceDedicatedBFR{}, aper.Constraint{Lb: 1, Ub: maxNrofCandidateBeams}, false)
		fn_CandidateBeamRSList := func() *PRACH_ResourceDedicatedBFR {
			return new(PRACH_ResourceDedicatedBFR)
		}
		if err = tmp_CandidateBeamRSList.Decode(r, fn_CandidateBeamRSList); err != nil {
			return utils.WrapError("Decode CandidateBeamRSList", err)
		}
		ie.CandidateBeamRSList = []PRACH_ResourceDedicatedBFR{}
		for _, i := range tmp_CandidateBeamRSList.Value {
			ie.CandidateBeamRSList = append(ie.CandidateBeamRSList, *i)
		}
	}
	if Ssb_perRACH_OccasionPresent {
		ie.Ssb_perRACH_Occasion = new(BeamFailureRecoveryConfig_ssb_perRACH_Occasion)
		if err = ie.Ssb_perRACH_Occasion.Decode(r); err != nil {
			return utils.WrapError("Decode Ssb_perRACH_Occasion", err)
		}
	}
	if Ra_ssb_OccasionMaskIndexPresent {
		var tmp_int_Ra_ssb_OccasionMaskIndex int64
		if tmp_int_Ra_ssb_OccasionMaskIndex, err = r.ReadInteger(&aper.Constraint{Lb: 0, Ub: 15}, false); err != nil {
			return utils.WrapError("Decode Ra_ssb_OccasionMaskIndex", err)
		}
		ie.Ra_ssb_OccasionMaskIndex = &tmp_int_Ra_ssb_OccasionMaskIndex
	}
	if RecoverySearchSpaceIdPresent {
		ie.RecoverySearchSpaceId = new(SearchSpaceId)
		if err = ie.RecoverySearchSpaceId.Decode(r); err != nil {
			return utils.WrapError("Decode RecoverySearchSpaceId", err)
		}
	}
	if Ra_PrioritizationPresent {
		ie.Ra_Prioritization = new(RA_Prioritization)
		if err = ie.Ra_Prioritization.Decode(r); err != nil {
			return utils.WrapError("Decode Ra_Prioritization", err)
		}
	}
	if BeamFailureRecoveryTimerPresent {
		ie.BeamFailureRecoveryTimer = new(BeamFailureRecoveryConfig_beamFailureRecoveryTimer)
		if err = ie.BeamFailureRecoveryTimer.Decode(r); err != nil {
			return utils.WrapError("Decode BeamFailureRecoveryTimer", err)
		}
	}

	if extensionBit {
		// Read extension bitmap: 3 bits for 3 extension groups
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

			Msg1_SubcarrierSpacingPresent, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			// decode Msg1_SubcarrierSpacing optional
			if Msg1_SubcarrierSpacingPresent {
				ie.Msg1_SubcarrierSpacing = new(SubcarrierSpacing)
				if err = ie.Msg1_SubcarrierSpacing.Decode(extReader); err != nil {
					return utils.WrapError("Decode Msg1_SubcarrierSpacing", err)
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

			Ra_PrioritizationTwoStep_r16Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			CandidateBeamRSListExt_v1610Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			// decode Ra_PrioritizationTwoStep_r16 optional
			if Ra_PrioritizationTwoStep_r16Present {
				ie.Ra_PrioritizationTwoStep_r16 = new(RA_Prioritization)
				if err = ie.Ra_PrioritizationTwoStep_r16.Decode(extReader); err != nil {
					return utils.WrapError("Decode Ra_PrioritizationTwoStep_r16", err)
				}
			}
			// decode CandidateBeamRSListExt_v1610 optional
			if CandidateBeamRSListExt_v1610Present {
				tmp_CandidateBeamRSListExt_v1610 := utils.SetupRelease[*CandidateBeamRSListExt_r16]{}
				if err = tmp_CandidateBeamRSListExt_v1610.Decode(extReader); err != nil {
					return utils.WrapError("Decode CandidateBeamRSListExt_v1610", err)
				}
				ie.CandidateBeamRSListExt_v1610 = tmp_CandidateBeamRSListExt_v1610.Setup
			}
		}
		// decode extension group 3
		if len(extBitmap) > 2 && extBitmap[2] {
			extBytes, err := r.ReadOpenType()
			if err != nil {
				return err
			}

			extReader := aper.NewReader(bytes.NewReader(extBytes))

			SpCell_BFR_CBRA_r16Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			// decode SpCell_BFR_CBRA_r16 optional
			if SpCell_BFR_CBRA_r16Present {
				ie.SpCell_BFR_CBRA_r16 = new(BeamFailureRecoveryConfig_spCell_BFR_CBRA_r16)
				if err = ie.SpCell_BFR_CBRA_r16.Decode(extReader); err != nil {
					return utils.WrapError("Decode SpCell_BFR_CBRA_r16", err)
				}
			}
		}
	}
	return nil
}
