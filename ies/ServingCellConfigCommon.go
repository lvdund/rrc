package ies

import (
	"bytes"

	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type ServingCellConfigCommon struct {
	PhysCellId                     *PhysCellId                                             `optional`
	DownlinkConfigCommon           *DownlinkConfigCommon                                   `optional`
	UplinkConfigCommon             *UplinkConfigCommon                                     `optional`
	SupplementaryUplinkConfig      *UplinkConfigCommon                                     `optional`
	N_TimingAdvanceOffset          *ServingCellConfigCommon_n_TimingAdvanceOffset          `optional`
	Ssb_PositionsInBurst           *ServingCellConfigCommon_ssb_PositionsInBurst           `lb:4,ub:4,optional`
	Ssb_periodicityServingCell     *ServingCellConfigCommon_ssb_periodicityServingCell     `optional`
	Dmrs_TypeA_Position            ServingCellConfigCommon_dmrs_TypeA_Position             `madatory`
	Lte_CRS_ToMatchAround          *RateMatchPatternLTE_CRS                                `optional,setuprelease`
	RateMatchPatternToAddModList   []RateMatchPattern                                      `lb:1,ub:maxNrofRateMatchPatterns,optional`
	RateMatchPatternToReleaseList  []RateMatchPatternId                                    `lb:1,ub:maxNrofRateMatchPatterns,optional`
	SsbSubcarrierSpacing           *SubcarrierSpacing                                      `optional`
	Tdd_UL_DL_ConfigurationCommon  *TDD_UL_DL_ConfigCommon                                 `optional`
	Ss_PBCH_BlockPower             int64                                                   `lb:-60,ub:50,madatory`
	ChannelAccessMode_r16          *ServingCellConfigCommon_channelAccessMode_r16          `optional,ext-1`
	DiscoveryBurstWindowLength_r16 *ServingCellConfigCommon_discoveryBurstWindowLength_r16 `optional,ext-1`
	Ssb_PositionQCL_r16            *SSB_PositionQCL_Relation_r16                           `optional,ext-1`
	HighSpeedConfig_r16            *HighSpeedConfig_r16                                    `optional,ext-1`
	HighSpeedConfig_v1700          *HighSpeedConfig_v1700                                  `optional,ext-2`
	ChannelAccessMode2_r17         *ServingCellConfigCommon_channelAccessMode2_r17         `optional,ext-2`
	DiscoveryBurstWindowLength_r17 *ServingCellConfigCommon_discoveryBurstWindowLength_r17 `optional,ext-2`
	Ssb_PositionQCL_r17            *SSB_PositionQCL_Relation_r17                           `optional,ext-2`
	HighSpeedConfigFR2_r17         *HighSpeedConfigFR2_r17                                 `optional,ext-2`
	UplinkConfigCommon_v1700       *UplinkConfigCommon_v1700                               `optional,ext-2`
	Ntn_Config_r17                 *NTN_Config_r17                                         `optional,ext-2`
	FeaturePriorities_r17          *FeaturePriorities_r17                                  `optional,ext-3`
}

func (ie *ServingCellConfigCommon) Encode(w *uper.UperWriter) error {
	var err error
	hasExtensions := ie.ChannelAccessMode_r16 != nil || ie.DiscoveryBurstWindowLength_r16 != nil || ie.Ssb_PositionQCL_r16 != nil || ie.HighSpeedConfig_r16 != nil || ie.HighSpeedConfig_v1700 != nil || ie.ChannelAccessMode2_r17 != nil || ie.DiscoveryBurstWindowLength_r17 != nil || ie.Ssb_PositionQCL_r17 != nil || ie.HighSpeedConfigFR2_r17 != nil || ie.UplinkConfigCommon_v1700 != nil || ie.Ntn_Config_r17 != nil || ie.FeaturePriorities_r17 != nil
	preambleBits := []bool{hasExtensions, ie.PhysCellId != nil, ie.DownlinkConfigCommon != nil, ie.UplinkConfigCommon != nil, ie.SupplementaryUplinkConfig != nil, ie.N_TimingAdvanceOffset != nil, ie.Ssb_PositionsInBurst != nil, ie.Ssb_periodicityServingCell != nil, ie.Lte_CRS_ToMatchAround != nil, len(ie.RateMatchPatternToAddModList) > 0, len(ie.RateMatchPatternToReleaseList) > 0, ie.SsbSubcarrierSpacing != nil, ie.Tdd_UL_DL_ConfigurationCommon != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.PhysCellId != nil {
		if err = ie.PhysCellId.Encode(w); err != nil {
			return utils.WrapError("Encode PhysCellId", err)
		}
	}
	if ie.DownlinkConfigCommon != nil {
		if err = ie.DownlinkConfigCommon.Encode(w); err != nil {
			return utils.WrapError("Encode DownlinkConfigCommon", err)
		}
	}
	if ie.UplinkConfigCommon != nil {
		if err = ie.UplinkConfigCommon.Encode(w); err != nil {
			return utils.WrapError("Encode UplinkConfigCommon", err)
		}
	}
	if ie.SupplementaryUplinkConfig != nil {
		if err = ie.SupplementaryUplinkConfig.Encode(w); err != nil {
			return utils.WrapError("Encode SupplementaryUplinkConfig", err)
		}
	}
	if ie.N_TimingAdvanceOffset != nil {
		if err = ie.N_TimingAdvanceOffset.Encode(w); err != nil {
			return utils.WrapError("Encode N_TimingAdvanceOffset", err)
		}
	}
	if ie.Ssb_PositionsInBurst != nil {
		if err = ie.Ssb_PositionsInBurst.Encode(w); err != nil {
			return utils.WrapError("Encode Ssb_PositionsInBurst", err)
		}
	}
	if ie.Ssb_periodicityServingCell != nil {
		if err = ie.Ssb_periodicityServingCell.Encode(w); err != nil {
			return utils.WrapError("Encode Ssb_periodicityServingCell", err)
		}
	}
	if err = ie.Dmrs_TypeA_Position.Encode(w); err != nil {
		return utils.WrapError("Encode Dmrs_TypeA_Position", err)
	}
	if ie.Lte_CRS_ToMatchAround != nil {
		tmp_Lte_CRS_ToMatchAround := utils.SetupRelease[*RateMatchPatternLTE_CRS]{
			Setup: ie.Lte_CRS_ToMatchAround,
		}
		if err = tmp_Lte_CRS_ToMatchAround.Encode(w); err != nil {
			return utils.WrapError("Encode Lte_CRS_ToMatchAround", err)
		}
	}
	if len(ie.RateMatchPatternToAddModList) > 0 {
		tmp_RateMatchPatternToAddModList := utils.NewSequence[*RateMatchPattern]([]*RateMatchPattern{}, uper.Constraint{Lb: 1, Ub: maxNrofRateMatchPatterns}, false)
		for _, i := range ie.RateMatchPatternToAddModList {
			tmp_RateMatchPatternToAddModList.Value = append(tmp_RateMatchPatternToAddModList.Value, &i)
		}
		if err = tmp_RateMatchPatternToAddModList.Encode(w); err != nil {
			return utils.WrapError("Encode RateMatchPatternToAddModList", err)
		}
	}
	if len(ie.RateMatchPatternToReleaseList) > 0 {
		tmp_RateMatchPatternToReleaseList := utils.NewSequence[*RateMatchPatternId]([]*RateMatchPatternId{}, uper.Constraint{Lb: 1, Ub: maxNrofRateMatchPatterns}, false)
		for _, i := range ie.RateMatchPatternToReleaseList {
			tmp_RateMatchPatternToReleaseList.Value = append(tmp_RateMatchPatternToReleaseList.Value, &i)
		}
		if err = tmp_RateMatchPatternToReleaseList.Encode(w); err != nil {
			return utils.WrapError("Encode RateMatchPatternToReleaseList", err)
		}
	}
	if ie.SsbSubcarrierSpacing != nil {
		if err = ie.SsbSubcarrierSpacing.Encode(w); err != nil {
			return utils.WrapError("Encode SsbSubcarrierSpacing", err)
		}
	}
	if ie.Tdd_UL_DL_ConfigurationCommon != nil {
		if err = ie.Tdd_UL_DL_ConfigurationCommon.Encode(w); err != nil {
			return utils.WrapError("Encode Tdd_UL_DL_ConfigurationCommon", err)
		}
	}
	if err = w.WriteInteger(ie.Ss_PBCH_BlockPower, &uper.Constraint{Lb: -60, Ub: 50}, false); err != nil {
		return utils.WrapError("WriteInteger Ss_PBCH_BlockPower", err)
	}
	if hasExtensions {
		// Extension bitmap: 3 bits for 3 extension groups
		extBitmap := []bool{ie.ChannelAccessMode_r16 != nil || ie.DiscoveryBurstWindowLength_r16 != nil || ie.Ssb_PositionQCL_r16 != nil || ie.HighSpeedConfig_r16 != nil, ie.HighSpeedConfig_v1700 != nil || ie.ChannelAccessMode2_r17 != nil || ie.DiscoveryBurstWindowLength_r17 != nil || ie.Ssb_PositionQCL_r17 != nil || ie.HighSpeedConfigFR2_r17 != nil || ie.UplinkConfigCommon_v1700 != nil || ie.Ntn_Config_r17 != nil, ie.FeaturePriorities_r17 != nil}
		if err := w.WriteExtBitMap(extBitmap); err != nil {
			return utils.WrapError("WriteExtBitMap ServingCellConfigCommon", err)
		}

		// encode extension group 1
		if extBitmap[0] {
			extBuf := new(bytes.Buffer)
			extWriter := uper.NewWriter(extBuf)

			// Write preamble bits for optional fields in extension group 1
			optionals_ext_1 := []bool{ie.ChannelAccessMode_r16 != nil, ie.DiscoveryBurstWindowLength_r16 != nil, ie.Ssb_PositionQCL_r16 != nil, ie.HighSpeedConfig_r16 != nil}
			for _, bit := range optionals_ext_1 {
				if err := extWriter.WriteBool(bit); err != nil {
					return err
				}
			}

			// encode ChannelAccessMode_r16 optional
			if ie.ChannelAccessMode_r16 != nil {
				if err = ie.ChannelAccessMode_r16.Encode(extWriter); err != nil {
					return utils.WrapError("Encode ChannelAccessMode_r16", err)
				}
			}
			// encode DiscoveryBurstWindowLength_r16 optional
			if ie.DiscoveryBurstWindowLength_r16 != nil {
				if err = ie.DiscoveryBurstWindowLength_r16.Encode(extWriter); err != nil {
					return utils.WrapError("Encode DiscoveryBurstWindowLength_r16", err)
				}
			}
			// encode Ssb_PositionQCL_r16 optional
			if ie.Ssb_PositionQCL_r16 != nil {
				if err = ie.Ssb_PositionQCL_r16.Encode(extWriter); err != nil {
					return utils.WrapError("Encode Ssb_PositionQCL_r16", err)
				}
			}
			// encode HighSpeedConfig_r16 optional
			if ie.HighSpeedConfig_r16 != nil {
				if err = ie.HighSpeedConfig_r16.Encode(extWriter); err != nil {
					return utils.WrapError("Encode HighSpeedConfig_r16", err)
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
			optionals_ext_2 := []bool{ie.HighSpeedConfig_v1700 != nil, ie.ChannelAccessMode2_r17 != nil, ie.DiscoveryBurstWindowLength_r17 != nil, ie.Ssb_PositionQCL_r17 != nil, ie.HighSpeedConfigFR2_r17 != nil, ie.UplinkConfigCommon_v1700 != nil, ie.Ntn_Config_r17 != nil}
			for _, bit := range optionals_ext_2 {
				if err := extWriter.WriteBool(bit); err != nil {
					return err
				}
			}

			// encode HighSpeedConfig_v1700 optional
			if ie.HighSpeedConfig_v1700 != nil {
				if err = ie.HighSpeedConfig_v1700.Encode(extWriter); err != nil {
					return utils.WrapError("Encode HighSpeedConfig_v1700", err)
				}
			}
			// encode ChannelAccessMode2_r17 optional
			if ie.ChannelAccessMode2_r17 != nil {
				if err = ie.ChannelAccessMode2_r17.Encode(extWriter); err != nil {
					return utils.WrapError("Encode ChannelAccessMode2_r17", err)
				}
			}
			// encode DiscoveryBurstWindowLength_r17 optional
			if ie.DiscoveryBurstWindowLength_r17 != nil {
				if err = ie.DiscoveryBurstWindowLength_r17.Encode(extWriter); err != nil {
					return utils.WrapError("Encode DiscoveryBurstWindowLength_r17", err)
				}
			}
			// encode Ssb_PositionQCL_r17 optional
			if ie.Ssb_PositionQCL_r17 != nil {
				if err = ie.Ssb_PositionQCL_r17.Encode(extWriter); err != nil {
					return utils.WrapError("Encode Ssb_PositionQCL_r17", err)
				}
			}
			// encode HighSpeedConfigFR2_r17 optional
			if ie.HighSpeedConfigFR2_r17 != nil {
				if err = ie.HighSpeedConfigFR2_r17.Encode(extWriter); err != nil {
					return utils.WrapError("Encode HighSpeedConfigFR2_r17", err)
				}
			}
			// encode UplinkConfigCommon_v1700 optional
			if ie.UplinkConfigCommon_v1700 != nil {
				if err = ie.UplinkConfigCommon_v1700.Encode(extWriter); err != nil {
					return utils.WrapError("Encode UplinkConfigCommon_v1700", err)
				}
			}
			// encode Ntn_Config_r17 optional
			if ie.Ntn_Config_r17 != nil {
				if err = ie.Ntn_Config_r17.Encode(extWriter); err != nil {
					return utils.WrapError("Encode Ntn_Config_r17", err)
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
			optionals_ext_3 := []bool{ie.FeaturePriorities_r17 != nil}
			for _, bit := range optionals_ext_3 {
				if err := extWriter.WriteBool(bit); err != nil {
					return err
				}
			}

			// encode FeaturePriorities_r17 optional
			if ie.FeaturePriorities_r17 != nil {
				if err = ie.FeaturePriorities_r17.Encode(extWriter); err != nil {
					return utils.WrapError("Encode FeaturePriorities_r17", err)
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

func (ie *ServingCellConfigCommon) Decode(r *uper.UperReader) error {
	var err error
	var extensionBit bool
	if extensionBit, err = r.ReadBool(); err != nil {
		return err
	}
	var PhysCellIdPresent bool
	if PhysCellIdPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var DownlinkConfigCommonPresent bool
	if DownlinkConfigCommonPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var UplinkConfigCommonPresent bool
	if UplinkConfigCommonPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var SupplementaryUplinkConfigPresent bool
	if SupplementaryUplinkConfigPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var N_TimingAdvanceOffsetPresent bool
	if N_TimingAdvanceOffsetPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var Ssb_PositionsInBurstPresent bool
	if Ssb_PositionsInBurstPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var Ssb_periodicityServingCellPresent bool
	if Ssb_periodicityServingCellPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var Lte_CRS_ToMatchAroundPresent bool
	if Lte_CRS_ToMatchAroundPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var RateMatchPatternToAddModListPresent bool
	if RateMatchPatternToAddModListPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var RateMatchPatternToReleaseListPresent bool
	if RateMatchPatternToReleaseListPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var SsbSubcarrierSpacingPresent bool
	if SsbSubcarrierSpacingPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var Tdd_UL_DL_ConfigurationCommonPresent bool
	if Tdd_UL_DL_ConfigurationCommonPresent, err = r.ReadBool(); err != nil {
		return err
	}
	if PhysCellIdPresent {
		ie.PhysCellId = new(PhysCellId)
		if err = ie.PhysCellId.Decode(r); err != nil {
			return utils.WrapError("Decode PhysCellId", err)
		}
	}
	if DownlinkConfigCommonPresent {
		ie.DownlinkConfigCommon = new(DownlinkConfigCommon)
		if err = ie.DownlinkConfigCommon.Decode(r); err != nil {
			return utils.WrapError("Decode DownlinkConfigCommon", err)
		}
	}
	if UplinkConfigCommonPresent {
		ie.UplinkConfigCommon = new(UplinkConfigCommon)
		if err = ie.UplinkConfigCommon.Decode(r); err != nil {
			return utils.WrapError("Decode UplinkConfigCommon", err)
		}
	}
	if SupplementaryUplinkConfigPresent {
		ie.SupplementaryUplinkConfig = new(UplinkConfigCommon)
		if err = ie.SupplementaryUplinkConfig.Decode(r); err != nil {
			return utils.WrapError("Decode SupplementaryUplinkConfig", err)
		}
	}
	if N_TimingAdvanceOffsetPresent {
		ie.N_TimingAdvanceOffset = new(ServingCellConfigCommon_n_TimingAdvanceOffset)
		if err = ie.N_TimingAdvanceOffset.Decode(r); err != nil {
			return utils.WrapError("Decode N_TimingAdvanceOffset", err)
		}
	}
	if Ssb_PositionsInBurstPresent {
		ie.Ssb_PositionsInBurst = new(ServingCellConfigCommon_ssb_PositionsInBurst)
		if err = ie.Ssb_PositionsInBurst.Decode(r); err != nil {
			return utils.WrapError("Decode Ssb_PositionsInBurst", err)
		}
	}
	if Ssb_periodicityServingCellPresent {
		ie.Ssb_periodicityServingCell = new(ServingCellConfigCommon_ssb_periodicityServingCell)
		if err = ie.Ssb_periodicityServingCell.Decode(r); err != nil {
			return utils.WrapError("Decode Ssb_periodicityServingCell", err)
		}
	}
	if err = ie.Dmrs_TypeA_Position.Decode(r); err != nil {
		return utils.WrapError("Decode Dmrs_TypeA_Position", err)
	}
	if Lte_CRS_ToMatchAroundPresent {
		tmp_Lte_CRS_ToMatchAround := utils.SetupRelease[*RateMatchPatternLTE_CRS]{}
		if err = tmp_Lte_CRS_ToMatchAround.Decode(r); err != nil {
			return utils.WrapError("Decode Lte_CRS_ToMatchAround", err)
		}
		ie.Lte_CRS_ToMatchAround = tmp_Lte_CRS_ToMatchAround.Setup
	}
	if RateMatchPatternToAddModListPresent {
		tmp_RateMatchPatternToAddModList := utils.NewSequence[*RateMatchPattern]([]*RateMatchPattern{}, uper.Constraint{Lb: 1, Ub: maxNrofRateMatchPatterns}, false)
		fn_RateMatchPatternToAddModList := func() *RateMatchPattern {
			return new(RateMatchPattern)
		}
		if err = tmp_RateMatchPatternToAddModList.Decode(r, fn_RateMatchPatternToAddModList); err != nil {
			return utils.WrapError("Decode RateMatchPatternToAddModList", err)
		}
		ie.RateMatchPatternToAddModList = []RateMatchPattern{}
		for _, i := range tmp_RateMatchPatternToAddModList.Value {
			ie.RateMatchPatternToAddModList = append(ie.RateMatchPatternToAddModList, *i)
		}
	}
	if RateMatchPatternToReleaseListPresent {
		tmp_RateMatchPatternToReleaseList := utils.NewSequence[*RateMatchPatternId]([]*RateMatchPatternId{}, uper.Constraint{Lb: 1, Ub: maxNrofRateMatchPatterns}, false)
		fn_RateMatchPatternToReleaseList := func() *RateMatchPatternId {
			return new(RateMatchPatternId)
		}
		if err = tmp_RateMatchPatternToReleaseList.Decode(r, fn_RateMatchPatternToReleaseList); err != nil {
			return utils.WrapError("Decode RateMatchPatternToReleaseList", err)
		}
		ie.RateMatchPatternToReleaseList = []RateMatchPatternId{}
		for _, i := range tmp_RateMatchPatternToReleaseList.Value {
			ie.RateMatchPatternToReleaseList = append(ie.RateMatchPatternToReleaseList, *i)
		}
	}
	if SsbSubcarrierSpacingPresent {
		ie.SsbSubcarrierSpacing = new(SubcarrierSpacing)
		if err = ie.SsbSubcarrierSpacing.Decode(r); err != nil {
			return utils.WrapError("Decode SsbSubcarrierSpacing", err)
		}
	}
	if Tdd_UL_DL_ConfigurationCommonPresent {
		ie.Tdd_UL_DL_ConfigurationCommon = new(TDD_UL_DL_ConfigCommon)
		if err = ie.Tdd_UL_DL_ConfigurationCommon.Decode(r); err != nil {
			return utils.WrapError("Decode Tdd_UL_DL_ConfigurationCommon", err)
		}
	}
	var tmp_int_Ss_PBCH_BlockPower int64
	if tmp_int_Ss_PBCH_BlockPower, err = r.ReadInteger(&uper.Constraint{Lb: -60, Ub: 50}, false); err != nil {
		return utils.WrapError("ReadInteger Ss_PBCH_BlockPower", err)
	}
	ie.Ss_PBCH_BlockPower = tmp_int_Ss_PBCH_BlockPower

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

			extReader := uper.NewReader(bytes.NewReader(extBytes))

			ChannelAccessMode_r16Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			DiscoveryBurstWindowLength_r16Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			Ssb_PositionQCL_r16Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			HighSpeedConfig_r16Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			// decode ChannelAccessMode_r16 optional
			if ChannelAccessMode_r16Present {
				ie.ChannelAccessMode_r16 = new(ServingCellConfigCommon_channelAccessMode_r16)
				if err = ie.ChannelAccessMode_r16.Decode(extReader); err != nil {
					return utils.WrapError("Decode ChannelAccessMode_r16", err)
				}
			}
			// decode DiscoveryBurstWindowLength_r16 optional
			if DiscoveryBurstWindowLength_r16Present {
				ie.DiscoveryBurstWindowLength_r16 = new(ServingCellConfigCommon_discoveryBurstWindowLength_r16)
				if err = ie.DiscoveryBurstWindowLength_r16.Decode(extReader); err != nil {
					return utils.WrapError("Decode DiscoveryBurstWindowLength_r16", err)
				}
			}
			// decode Ssb_PositionQCL_r16 optional
			if Ssb_PositionQCL_r16Present {
				ie.Ssb_PositionQCL_r16 = new(SSB_PositionQCL_Relation_r16)
				if err = ie.Ssb_PositionQCL_r16.Decode(extReader); err != nil {
					return utils.WrapError("Decode Ssb_PositionQCL_r16", err)
				}
			}
			// decode HighSpeedConfig_r16 optional
			if HighSpeedConfig_r16Present {
				ie.HighSpeedConfig_r16 = new(HighSpeedConfig_r16)
				if err = ie.HighSpeedConfig_r16.Decode(extReader); err != nil {
					return utils.WrapError("Decode HighSpeedConfig_r16", err)
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

			HighSpeedConfig_v1700Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			ChannelAccessMode2_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			DiscoveryBurstWindowLength_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			Ssb_PositionQCL_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			HighSpeedConfigFR2_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			UplinkConfigCommon_v1700Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			Ntn_Config_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			// decode HighSpeedConfig_v1700 optional
			if HighSpeedConfig_v1700Present {
				ie.HighSpeedConfig_v1700 = new(HighSpeedConfig_v1700)
				if err = ie.HighSpeedConfig_v1700.Decode(extReader); err != nil {
					return utils.WrapError("Decode HighSpeedConfig_v1700", err)
				}
			}
			// decode ChannelAccessMode2_r17 optional
			if ChannelAccessMode2_r17Present {
				ie.ChannelAccessMode2_r17 = new(ServingCellConfigCommon_channelAccessMode2_r17)
				if err = ie.ChannelAccessMode2_r17.Decode(extReader); err != nil {
					return utils.WrapError("Decode ChannelAccessMode2_r17", err)
				}
			}
			// decode DiscoveryBurstWindowLength_r17 optional
			if DiscoveryBurstWindowLength_r17Present {
				ie.DiscoveryBurstWindowLength_r17 = new(ServingCellConfigCommon_discoveryBurstWindowLength_r17)
				if err = ie.DiscoveryBurstWindowLength_r17.Decode(extReader); err != nil {
					return utils.WrapError("Decode DiscoveryBurstWindowLength_r17", err)
				}
			}
			// decode Ssb_PositionQCL_r17 optional
			if Ssb_PositionQCL_r17Present {
				ie.Ssb_PositionQCL_r17 = new(SSB_PositionQCL_Relation_r17)
				if err = ie.Ssb_PositionQCL_r17.Decode(extReader); err != nil {
					return utils.WrapError("Decode Ssb_PositionQCL_r17", err)
				}
			}
			// decode HighSpeedConfigFR2_r17 optional
			if HighSpeedConfigFR2_r17Present {
				ie.HighSpeedConfigFR2_r17 = new(HighSpeedConfigFR2_r17)
				if err = ie.HighSpeedConfigFR2_r17.Decode(extReader); err != nil {
					return utils.WrapError("Decode HighSpeedConfigFR2_r17", err)
				}
			}
			// decode UplinkConfigCommon_v1700 optional
			if UplinkConfigCommon_v1700Present {
				ie.UplinkConfigCommon_v1700 = new(UplinkConfigCommon_v1700)
				if err = ie.UplinkConfigCommon_v1700.Decode(extReader); err != nil {
					return utils.WrapError("Decode UplinkConfigCommon_v1700", err)
				}
			}
			// decode Ntn_Config_r17 optional
			if Ntn_Config_r17Present {
				ie.Ntn_Config_r17 = new(NTN_Config_r17)
				if err = ie.Ntn_Config_r17.Decode(extReader); err != nil {
					return utils.WrapError("Decode Ntn_Config_r17", err)
				}
			}
		}
		// decode extension group 3
		if len(extBitmap) > 2 && extBitmap[2] {
			extBytes, err := r.ReadOpenType()
			if err != nil {
				return err
			}

			extReader := uper.NewReader(bytes.NewReader(extBytes))

			FeaturePriorities_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			// decode FeaturePriorities_r17 optional
			if FeaturePriorities_r17Present {
				ie.FeaturePriorities_r17 = new(FeaturePriorities_r17)
				if err = ie.FeaturePriorities_r17.Decode(extReader); err != nil {
					return utils.WrapError("Decode FeaturePriorities_r17", err)
				}
			}
		}
	}
	return nil
}
