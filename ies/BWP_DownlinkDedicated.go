package ies

import (
	"bytes"

	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type BWP_DownlinkDedicated struct {
	Pdcch_Config                          *PDCCH_Config                        `optional,setuprelease`
	Pdsch_Config                          *PDSCH_Config                        `optional,setuprelease`
	Sps_Config                            *SPS_Config                          `optional,setuprelease`
	RadioLinkMonitoringConfig             *RadioLinkMonitoringConfig           `optional,setuprelease`
	Sps_ConfigToAddModList_r16            *SPS_ConfigToAddModList_r16          `optional,ext-1`
	Sps_ConfigToReleaseList_r16           *SPS_ConfigToReleaseList_r16         `optional,ext-1`
	Sps_ConfigDeactivationStateList_r16   *SPS_ConfigDeactivationStateList_r16 `optional,ext-1`
	BeamFailureRecoverySCellConfig_r16    *BeamFailureRecoveryRSConfig_r16     `optional,ext-1,setuprelease`
	Sl_PDCCH_Config_r16                   *PDCCH_Config                        `optional,ext-1,setuprelease`
	Sl_V2X_PDCCH_Config_r16               *PDCCH_Config                        `optional,ext-1,setuprelease`
	PreConfGapStatus_r17                  *uper.BitString                      `lb:maxNrofGapId_r17,ub:maxNrofGapId_r17,optional,ext-2`
	BeamFailureRecoverySpCellConfig_r17   *BeamFailureRecoveryRSConfig_r16     `optional,ext-2,setuprelease`
	Harq_FeedbackEnablingforSPSactive_r17 *bool                                `optional,ext-2`
	Cfr_ConfigMulticast_r17               *CFR_ConfigMulticast_r17             `optional,ext-2,setuprelease`
	Dl_PPW_PreConfigToAddModList_r17      *DL_PPW_PreConfigToAddModList_r17    `optional,ext-2`
	Dl_PPW_PreConfigToReleaseList_r17     *DL_PPW_PreConfigToReleaseList_r17   `optional,ext-2`
	NonCellDefiningSSB_r17                *NonCellDefiningSSB_r17              `optional,ext-2`
	ServingCellMO_r17                     *MeasObjectId                        `optional,ext-2`
}

func (ie *BWP_DownlinkDedicated) Encode(w *uper.UperWriter) error {
	var err error
	hasExtensions := ie.Sps_ConfigToAddModList_r16 != nil || ie.Sps_ConfigToReleaseList_r16 != nil || ie.Sps_ConfigDeactivationStateList_r16 != nil || ie.BeamFailureRecoverySCellConfig_r16 != nil || ie.Sl_PDCCH_Config_r16 != nil || ie.Sl_V2X_PDCCH_Config_r16 != nil || ie.PreConfGapStatus_r17 != nil || ie.BeamFailureRecoverySpCellConfig_r17 != nil || ie.Harq_FeedbackEnablingforSPSactive_r17 != nil || ie.Cfr_ConfigMulticast_r17 != nil || ie.Dl_PPW_PreConfigToAddModList_r17 != nil || ie.Dl_PPW_PreConfigToReleaseList_r17 != nil || ie.NonCellDefiningSSB_r17 != nil || ie.ServingCellMO_r17 != nil
	preambleBits := []bool{hasExtensions, ie.Pdcch_Config != nil, ie.Pdsch_Config != nil, ie.Sps_Config != nil, ie.RadioLinkMonitoringConfig != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.Pdcch_Config != nil {
		tmp_Pdcch_Config := utils.SetupRelease[*PDCCH_Config]{
			Setup: ie.Pdcch_Config,
		}
		if err = tmp_Pdcch_Config.Encode(w); err != nil {
			return utils.WrapError("Encode Pdcch_Config", err)
		}
	}
	if ie.Pdsch_Config != nil {
		tmp_Pdsch_Config := utils.SetupRelease[*PDSCH_Config]{
			Setup: ie.Pdsch_Config,
		}
		if err = tmp_Pdsch_Config.Encode(w); err != nil {
			return utils.WrapError("Encode Pdsch_Config", err)
		}
	}
	if ie.Sps_Config != nil {
		tmp_Sps_Config := utils.SetupRelease[*SPS_Config]{
			Setup: ie.Sps_Config,
		}
		if err = tmp_Sps_Config.Encode(w); err != nil {
			return utils.WrapError("Encode Sps_Config", err)
		}
	}
	if ie.RadioLinkMonitoringConfig != nil {
		tmp_RadioLinkMonitoringConfig := utils.SetupRelease[*RadioLinkMonitoringConfig]{
			Setup: ie.RadioLinkMonitoringConfig,
		}
		if err = tmp_RadioLinkMonitoringConfig.Encode(w); err != nil {
			return utils.WrapError("Encode RadioLinkMonitoringConfig", err)
		}
	}
	if hasExtensions {
		// Extension bitmap: 2 bits for 2 extension groups
		extBitmap := []bool{ie.Sps_ConfigToAddModList_r16 != nil || ie.Sps_ConfigToReleaseList_r16 != nil || ie.Sps_ConfigDeactivationStateList_r16 != nil || ie.BeamFailureRecoverySCellConfig_r16 != nil || ie.Sl_PDCCH_Config_r16 != nil || ie.Sl_V2X_PDCCH_Config_r16 != nil, ie.PreConfGapStatus_r17 != nil || ie.BeamFailureRecoverySpCellConfig_r17 != nil || ie.Harq_FeedbackEnablingforSPSactive_r17 != nil || ie.Cfr_ConfigMulticast_r17 != nil || ie.Dl_PPW_PreConfigToAddModList_r17 != nil || ie.Dl_PPW_PreConfigToReleaseList_r17 != nil || ie.NonCellDefiningSSB_r17 != nil || ie.ServingCellMO_r17 != nil}
		if err := w.WriteExtBitMap(extBitmap); err != nil {
			return utils.WrapError("WriteExtBitMap BWP_DownlinkDedicated", err)
		}

		// encode extension group 1
		if extBitmap[0] {
			extBuf := new(bytes.Buffer)
			extWriter := uper.NewWriter(extBuf)

			// Write preamble bits for optional fields in extension group 1
			optionals_ext_1 := []bool{ie.Sps_ConfigToAddModList_r16 != nil, ie.Sps_ConfigToReleaseList_r16 != nil, ie.Sps_ConfigDeactivationStateList_r16 != nil, ie.BeamFailureRecoverySCellConfig_r16 != nil, ie.Sl_PDCCH_Config_r16 != nil, ie.Sl_V2X_PDCCH_Config_r16 != nil}
			for _, bit := range optionals_ext_1 {
				if err := extWriter.WriteBool(bit); err != nil {
					return err
				}
			}

			// encode Sps_ConfigToAddModList_r16 optional
			if ie.Sps_ConfigToAddModList_r16 != nil {
				if err = ie.Sps_ConfigToAddModList_r16.Encode(extWriter); err != nil {
					return utils.WrapError("Encode Sps_ConfigToAddModList_r16", err)
				}
			}
			// encode Sps_ConfigToReleaseList_r16 optional
			if ie.Sps_ConfigToReleaseList_r16 != nil {
				if err = ie.Sps_ConfigToReleaseList_r16.Encode(extWriter); err != nil {
					return utils.WrapError("Encode Sps_ConfigToReleaseList_r16", err)
				}
			}
			// encode Sps_ConfigDeactivationStateList_r16 optional
			if ie.Sps_ConfigDeactivationStateList_r16 != nil {
				if err = ie.Sps_ConfigDeactivationStateList_r16.Encode(extWriter); err != nil {
					return utils.WrapError("Encode Sps_ConfigDeactivationStateList_r16", err)
				}
			}
			// encode BeamFailureRecoverySCellConfig_r16 optional
			if ie.BeamFailureRecoverySCellConfig_r16 != nil {
				tmp_BeamFailureRecoverySCellConfig_r16 := utils.SetupRelease[*BeamFailureRecoveryRSConfig_r16]{
					Setup: ie.BeamFailureRecoverySCellConfig_r16,
				}
				if err = tmp_BeamFailureRecoverySCellConfig_r16.Encode(extWriter); err != nil {
					return utils.WrapError("Encode BeamFailureRecoverySCellConfig_r16", err)
				}
			}
			// encode Sl_PDCCH_Config_r16 optional
			if ie.Sl_PDCCH_Config_r16 != nil {
				tmp_Sl_PDCCH_Config_r16 := utils.SetupRelease[*PDCCH_Config]{
					Setup: ie.Sl_PDCCH_Config_r16,
				}
				if err = tmp_Sl_PDCCH_Config_r16.Encode(extWriter); err != nil {
					return utils.WrapError("Encode Sl_PDCCH_Config_r16", err)
				}
			}
			// encode Sl_V2X_PDCCH_Config_r16 optional
			if ie.Sl_V2X_PDCCH_Config_r16 != nil {
				tmp_Sl_V2X_PDCCH_Config_r16 := utils.SetupRelease[*PDCCH_Config]{
					Setup: ie.Sl_V2X_PDCCH_Config_r16,
				}
				if err = tmp_Sl_V2X_PDCCH_Config_r16.Encode(extWriter); err != nil {
					return utils.WrapError("Encode Sl_V2X_PDCCH_Config_r16", err)
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
			optionals_ext_2 := []bool{ie.PreConfGapStatus_r17 != nil, ie.BeamFailureRecoverySpCellConfig_r17 != nil, ie.Harq_FeedbackEnablingforSPSactive_r17 != nil, ie.Cfr_ConfigMulticast_r17 != nil, ie.Dl_PPW_PreConfigToAddModList_r17 != nil, ie.Dl_PPW_PreConfigToReleaseList_r17 != nil, ie.NonCellDefiningSSB_r17 != nil, ie.ServingCellMO_r17 != nil}
			for _, bit := range optionals_ext_2 {
				if err := extWriter.WriteBool(bit); err != nil {
					return err
				}
			}

			// encode PreConfGapStatus_r17 optional
			if ie.PreConfGapStatus_r17 != nil {
				if err = extWriter.WriteBitString(ie.PreConfGapStatus_r17.Bytes, uint(ie.PreConfGapStatus_r17.NumBits), &uper.Constraint{Lb: maxNrofGapId_r17, Ub: maxNrofGapId_r17}, false); err != nil {
					return utils.WrapError("Encode PreConfGapStatus_r17", err)
				}
			}
			// encode BeamFailureRecoverySpCellConfig_r17 optional
			if ie.BeamFailureRecoverySpCellConfig_r17 != nil {
				tmp_BeamFailureRecoverySpCellConfig_r17 := utils.SetupRelease[*BeamFailureRecoveryRSConfig_r16]{
					Setup: ie.BeamFailureRecoverySpCellConfig_r17,
				}
				if err = tmp_BeamFailureRecoverySpCellConfig_r17.Encode(extWriter); err != nil {
					return utils.WrapError("Encode BeamFailureRecoverySpCellConfig_r17", err)
				}
			}
			// encode Harq_FeedbackEnablingforSPSactive_r17 optional
			if ie.Harq_FeedbackEnablingforSPSactive_r17 != nil {
				if err = extWriter.WriteBoolean(*ie.Harq_FeedbackEnablingforSPSactive_r17); err != nil {
					return utils.WrapError("Encode Harq_FeedbackEnablingforSPSactive_r17", err)
				}
			}
			// encode Cfr_ConfigMulticast_r17 optional
			if ie.Cfr_ConfigMulticast_r17 != nil {
				tmp_Cfr_ConfigMulticast_r17 := utils.SetupRelease[*CFR_ConfigMulticast_r17]{
					Setup: ie.Cfr_ConfigMulticast_r17,
				}
				if err = tmp_Cfr_ConfigMulticast_r17.Encode(extWriter); err != nil {
					return utils.WrapError("Encode Cfr_ConfigMulticast_r17", err)
				}
			}
			// encode Dl_PPW_PreConfigToAddModList_r17 optional
			if ie.Dl_PPW_PreConfigToAddModList_r17 != nil {
				if err = ie.Dl_PPW_PreConfigToAddModList_r17.Encode(extWriter); err != nil {
					return utils.WrapError("Encode Dl_PPW_PreConfigToAddModList_r17", err)
				}
			}
			// encode Dl_PPW_PreConfigToReleaseList_r17 optional
			if ie.Dl_PPW_PreConfigToReleaseList_r17 != nil {
				if err = ie.Dl_PPW_PreConfigToReleaseList_r17.Encode(extWriter); err != nil {
					return utils.WrapError("Encode Dl_PPW_PreConfigToReleaseList_r17", err)
				}
			}
			// encode NonCellDefiningSSB_r17 optional
			if ie.NonCellDefiningSSB_r17 != nil {
				if err = ie.NonCellDefiningSSB_r17.Encode(extWriter); err != nil {
					return utils.WrapError("Encode NonCellDefiningSSB_r17", err)
				}
			}
			// encode ServingCellMO_r17 optional
			if ie.ServingCellMO_r17 != nil {
				if err = ie.ServingCellMO_r17.Encode(extWriter); err != nil {
					return utils.WrapError("Encode ServingCellMO_r17", err)
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

func (ie *BWP_DownlinkDedicated) Decode(r *uper.UperReader) error {
	var err error
	var extensionBit bool
	if extensionBit, err = r.ReadBool(); err != nil {
		return err
	}
	var Pdcch_ConfigPresent bool
	if Pdcch_ConfigPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var Pdsch_ConfigPresent bool
	if Pdsch_ConfigPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var Sps_ConfigPresent bool
	if Sps_ConfigPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var RadioLinkMonitoringConfigPresent bool
	if RadioLinkMonitoringConfigPresent, err = r.ReadBool(); err != nil {
		return err
	}
	if Pdcch_ConfigPresent {
		tmp_Pdcch_Config := utils.SetupRelease[*PDCCH_Config]{}
		if err = tmp_Pdcch_Config.Decode(r); err != nil {
			return utils.WrapError("Decode Pdcch_Config", err)
		}
		ie.Pdcch_Config = tmp_Pdcch_Config.Setup
	}
	if Pdsch_ConfigPresent {
		tmp_Pdsch_Config := utils.SetupRelease[*PDSCH_Config]{}
		if err = tmp_Pdsch_Config.Decode(r); err != nil {
			return utils.WrapError("Decode Pdsch_Config", err)
		}
		ie.Pdsch_Config = tmp_Pdsch_Config.Setup
	}
	if Sps_ConfigPresent {
		tmp_Sps_Config := utils.SetupRelease[*SPS_Config]{}
		if err = tmp_Sps_Config.Decode(r); err != nil {
			return utils.WrapError("Decode Sps_Config", err)
		}
		ie.Sps_Config = tmp_Sps_Config.Setup
	}
	if RadioLinkMonitoringConfigPresent {
		tmp_RadioLinkMonitoringConfig := utils.SetupRelease[*RadioLinkMonitoringConfig]{}
		if err = tmp_RadioLinkMonitoringConfig.Decode(r); err != nil {
			return utils.WrapError("Decode RadioLinkMonitoringConfig", err)
		}
		ie.RadioLinkMonitoringConfig = tmp_RadioLinkMonitoringConfig.Setup
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

			Sps_ConfigToAddModList_r16Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			Sps_ConfigToReleaseList_r16Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			Sps_ConfigDeactivationStateList_r16Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			BeamFailureRecoverySCellConfig_r16Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			Sl_PDCCH_Config_r16Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			Sl_V2X_PDCCH_Config_r16Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			// decode Sps_ConfigToAddModList_r16 optional
			if Sps_ConfigToAddModList_r16Present {
				ie.Sps_ConfigToAddModList_r16 = new(SPS_ConfigToAddModList_r16)
				if err = ie.Sps_ConfigToAddModList_r16.Decode(extReader); err != nil {
					return utils.WrapError("Decode Sps_ConfigToAddModList_r16", err)
				}
			}
			// decode Sps_ConfigToReleaseList_r16 optional
			if Sps_ConfigToReleaseList_r16Present {
				ie.Sps_ConfigToReleaseList_r16 = new(SPS_ConfigToReleaseList_r16)
				if err = ie.Sps_ConfigToReleaseList_r16.Decode(extReader); err != nil {
					return utils.WrapError("Decode Sps_ConfigToReleaseList_r16", err)
				}
			}
			// decode Sps_ConfigDeactivationStateList_r16 optional
			if Sps_ConfigDeactivationStateList_r16Present {
				ie.Sps_ConfigDeactivationStateList_r16 = new(SPS_ConfigDeactivationStateList_r16)
				if err = ie.Sps_ConfigDeactivationStateList_r16.Decode(extReader); err != nil {
					return utils.WrapError("Decode Sps_ConfigDeactivationStateList_r16", err)
				}
			}
			// decode BeamFailureRecoverySCellConfig_r16 optional
			if BeamFailureRecoverySCellConfig_r16Present {
				tmp_BeamFailureRecoverySCellConfig_r16 := utils.SetupRelease[*BeamFailureRecoveryRSConfig_r16]{}
				if err = tmp_BeamFailureRecoverySCellConfig_r16.Decode(extReader); err != nil {
					return utils.WrapError("Decode BeamFailureRecoverySCellConfig_r16", err)
				}
				ie.BeamFailureRecoverySCellConfig_r16 = tmp_BeamFailureRecoverySCellConfig_r16.Setup
			}
			// decode Sl_PDCCH_Config_r16 optional
			if Sl_PDCCH_Config_r16Present {
				tmp_Sl_PDCCH_Config_r16 := utils.SetupRelease[*PDCCH_Config]{}
				if err = tmp_Sl_PDCCH_Config_r16.Decode(extReader); err != nil {
					return utils.WrapError("Decode Sl_PDCCH_Config_r16", err)
				}
				ie.Sl_PDCCH_Config_r16 = tmp_Sl_PDCCH_Config_r16.Setup
			}
			// decode Sl_V2X_PDCCH_Config_r16 optional
			if Sl_V2X_PDCCH_Config_r16Present {
				tmp_Sl_V2X_PDCCH_Config_r16 := utils.SetupRelease[*PDCCH_Config]{}
				if err = tmp_Sl_V2X_PDCCH_Config_r16.Decode(extReader); err != nil {
					return utils.WrapError("Decode Sl_V2X_PDCCH_Config_r16", err)
				}
				ie.Sl_V2X_PDCCH_Config_r16 = tmp_Sl_V2X_PDCCH_Config_r16.Setup
			}
		}
		// decode extension group 2
		if len(extBitmap) > 1 && extBitmap[1] {
			extBytes, err := r.ReadOpenType()
			if err != nil {
				return err
			}

			extReader := uper.NewReader(bytes.NewReader(extBytes))

			PreConfGapStatus_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			BeamFailureRecoverySpCellConfig_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			Harq_FeedbackEnablingforSPSactive_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			Cfr_ConfigMulticast_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			Dl_PPW_PreConfigToAddModList_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			Dl_PPW_PreConfigToReleaseList_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			NonCellDefiningSSB_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			ServingCellMO_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			// decode PreConfGapStatus_r17 optional
			if PreConfGapStatus_r17Present {
				var tmp_bs_PreConfGapStatus_r17 []byte
				var n_PreConfGapStatus_r17 uint
				if tmp_bs_PreConfGapStatus_r17, n_PreConfGapStatus_r17, err = extReader.ReadBitString(&uper.Constraint{Lb: maxNrofGapId_r17, Ub: maxNrofGapId_r17}, false); err != nil {
					return utils.WrapError("Decode PreConfGapStatus_r17", err)
				}
				tmp_bitstring := uper.BitString{
					Bytes:   tmp_bs_PreConfGapStatus_r17,
					NumBits: uint64(n_PreConfGapStatus_r17),
				}
				ie.PreConfGapStatus_r17 = &tmp_bitstring
			}
			// decode BeamFailureRecoverySpCellConfig_r17 optional
			if BeamFailureRecoverySpCellConfig_r17Present {
				tmp_BeamFailureRecoverySpCellConfig_r17 := utils.SetupRelease[*BeamFailureRecoveryRSConfig_r16]{}
				if err = tmp_BeamFailureRecoverySpCellConfig_r17.Decode(extReader); err != nil {
					return utils.WrapError("Decode BeamFailureRecoverySpCellConfig_r17", err)
				}
				ie.BeamFailureRecoverySpCellConfig_r17 = tmp_BeamFailureRecoverySpCellConfig_r17.Setup
			}
			// decode Harq_FeedbackEnablingforSPSactive_r17 optional
			if Harq_FeedbackEnablingforSPSactive_r17Present {
				var tmp_bool_Harq_FeedbackEnablingforSPSactive_r17 bool
				if tmp_bool_Harq_FeedbackEnablingforSPSactive_r17, err = extReader.ReadBoolean(); err != nil {
					return utils.WrapError("Decode Harq_FeedbackEnablingforSPSactive_r17", err)
				}
				ie.Harq_FeedbackEnablingforSPSactive_r17 = &tmp_bool_Harq_FeedbackEnablingforSPSactive_r17
			}
			// decode Cfr_ConfigMulticast_r17 optional
			if Cfr_ConfigMulticast_r17Present {
				tmp_Cfr_ConfigMulticast_r17 := utils.SetupRelease[*CFR_ConfigMulticast_r17]{}
				if err = tmp_Cfr_ConfigMulticast_r17.Decode(extReader); err != nil {
					return utils.WrapError("Decode Cfr_ConfigMulticast_r17", err)
				}
				ie.Cfr_ConfigMulticast_r17 = tmp_Cfr_ConfigMulticast_r17.Setup
			}
			// decode Dl_PPW_PreConfigToAddModList_r17 optional
			if Dl_PPW_PreConfigToAddModList_r17Present {
				ie.Dl_PPW_PreConfigToAddModList_r17 = new(DL_PPW_PreConfigToAddModList_r17)
				if err = ie.Dl_PPW_PreConfigToAddModList_r17.Decode(extReader); err != nil {
					return utils.WrapError("Decode Dl_PPW_PreConfigToAddModList_r17", err)
				}
			}
			// decode Dl_PPW_PreConfigToReleaseList_r17 optional
			if Dl_PPW_PreConfigToReleaseList_r17Present {
				ie.Dl_PPW_PreConfigToReleaseList_r17 = new(DL_PPW_PreConfigToReleaseList_r17)
				if err = ie.Dl_PPW_PreConfigToReleaseList_r17.Decode(extReader); err != nil {
					return utils.WrapError("Decode Dl_PPW_PreConfigToReleaseList_r17", err)
				}
			}
			// decode NonCellDefiningSSB_r17 optional
			if NonCellDefiningSSB_r17Present {
				ie.NonCellDefiningSSB_r17 = new(NonCellDefiningSSB_r17)
				if err = ie.NonCellDefiningSSB_r17.Decode(extReader); err != nil {
					return utils.WrapError("Decode NonCellDefiningSSB_r17", err)
				}
			}
			// decode ServingCellMO_r17 optional
			if ServingCellMO_r17Present {
				ie.ServingCellMO_r17 = new(MeasObjectId)
				if err = ie.ServingCellMO_r17.Decode(extReader); err != nil {
					return utils.WrapError("Decode ServingCellMO_r17", err)
				}
			}
		}
	}
	return nil
}
