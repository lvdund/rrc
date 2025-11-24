package ies

import (
	"bytes"

	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type MAC_CellGroupConfig struct {
	Drx_Config                              *DRX_Config                                             `optional,setuprelease`
	SchedulingRequestConfig                 *SchedulingRequestConfig                                `optional`
	Bsr_Config                              *BSR_Config                                             `optional`
	Tag_Config                              *TAG_Config                                             `optional`
	Phr_Config                              *PHR_Config                                             `optional,setuprelease`
	SkipUplinkTxDynamic                     bool                                                    `madatory`
	Csi_Mask                                *bool                                                   `optional,ext-1`
	DataInactivityTimer                     *DataInactivityTimer                                    `optional,ext-1,setuprelease`
	UsePreBSR_r16                           *MAC_CellGroupConfig_usePreBSR_r16                      `optional,ext-2`
	SchedulingRequestID_LBT_SCell_r16       *SchedulingRequestId                                    `optional,ext-2`
	Lch_BasedPrioritization_r16             *MAC_CellGroupConfig_lch_BasedPrioritization_r16        `optional,ext-2`
	SchedulingRequestID_BFR_SCell_r16       *SchedulingRequestId                                    `optional,ext-2`
	Drx_ConfigSecondaryGroup_r16            *DRX_ConfigSecondaryGroup_r16                           `optional,ext-2,setuprelease`
	EnhancedSkipUplinkTxDynamic_r16         *MAC_CellGroupConfig_enhancedSkipUplinkTxDynamic_r16    `optional,ext-3`
	EnhancedSkipUplinkTxConfigured_r16      *MAC_CellGroupConfig_enhancedSkipUplinkTxConfigured_r16 `optional,ext-3`
	IntraCG_Prioritization_r17              *MAC_CellGroupConfig_intraCG_Prioritization_r17         `optional,ext-4`
	Drx_ConfigSL_r17                        *DRX_ConfigSL_r17                                       `optional,ext-4,setuprelease`
	Drx_ConfigExt_v1700                     *DRX_ConfigExt_v1700                                    `optional,ext-4,setuprelease`
	SchedulingRequestID_BFR_r17             *SchedulingRequestId                                    `optional,ext-4`
	SchedulingRequestID_BFR2_r17            *SchedulingRequestId                                    `optional,ext-4`
	SchedulingRequestConfig_v1700           *SchedulingRequestConfig_v1700                          `optional,ext-4`
	Tar_Config_r17                          *TAR_Config_r17                                         `optional,ext-4,setuprelease`
	G_RNTI_ConfigToAddModList_r17           []MBS_RNTI_SpecificConfig_r17                           `lb:1,ub:maxG_RNTI_r17,optional,ext-4`
	G_RNTI_ConfigToReleaseList_r17          []MBS_RNTI_SpecificConfigId_r17                         `lb:1,ub:maxG_RNTI_r17,optional,ext-4`
	G_CS_RNTI_ConfigToAddModList_r17        []MBS_RNTI_SpecificConfig_r17                           `lb:1,ub:maxG_CS_RNTI_r17,optional,ext-4`
	G_CS_RNTI_ConfigToReleaseList_r17       []MBS_RNTI_SpecificConfigId_r17                         `lb:1,ub:maxG_CS_RNTI_r17,optional,ext-4`
	AllowCSI_SRS_Tx_MulticastDRX_Active_r17 *bool                                                   `optional,ext-4`
	SchedulingRequestID_PosMG_Request_r17   *SchedulingRequestId                                    `optional,ext-5`
	Drx_LastTransmissionUL_r17              *MAC_CellGroupConfig_drx_LastTransmissionUL_r17         `optional,ext-5`
}

func (ie *MAC_CellGroupConfig) Encode(w *uper.UperWriter) error {
	var err error
	hasExtensions := ie.Csi_Mask != nil || ie.DataInactivityTimer != nil || ie.UsePreBSR_r16 != nil || ie.SchedulingRequestID_LBT_SCell_r16 != nil || ie.Lch_BasedPrioritization_r16 != nil || ie.SchedulingRequestID_BFR_SCell_r16 != nil || ie.Drx_ConfigSecondaryGroup_r16 != nil || ie.EnhancedSkipUplinkTxDynamic_r16 != nil || ie.EnhancedSkipUplinkTxConfigured_r16 != nil || ie.IntraCG_Prioritization_r17 != nil || ie.Drx_ConfigSL_r17 != nil || ie.Drx_ConfigExt_v1700 != nil || ie.SchedulingRequestID_BFR_r17 != nil || ie.SchedulingRequestID_BFR2_r17 != nil || ie.SchedulingRequestConfig_v1700 != nil || ie.Tar_Config_r17 != nil || len(ie.G_RNTI_ConfigToAddModList_r17) > 0 || len(ie.G_RNTI_ConfigToReleaseList_r17) > 0 || len(ie.G_CS_RNTI_ConfigToAddModList_r17) > 0 || len(ie.G_CS_RNTI_ConfigToReleaseList_r17) > 0 || ie.AllowCSI_SRS_Tx_MulticastDRX_Active_r17 != nil || ie.SchedulingRequestID_PosMG_Request_r17 != nil || ie.Drx_LastTransmissionUL_r17 != nil
	preambleBits := []bool{hasExtensions, ie.Drx_Config != nil, ie.SchedulingRequestConfig != nil, ie.Bsr_Config != nil, ie.Tag_Config != nil, ie.Phr_Config != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.Drx_Config != nil {
		tmp_Drx_Config := utils.SetupRelease[*DRX_Config]{
			Setup: ie.Drx_Config,
		}
		if err = tmp_Drx_Config.Encode(w); err != nil {
			return utils.WrapError("Encode Drx_Config", err)
		}
	}
	if ie.SchedulingRequestConfig != nil {
		if err = ie.SchedulingRequestConfig.Encode(w); err != nil {
			return utils.WrapError("Encode SchedulingRequestConfig", err)
		}
	}
	if ie.Bsr_Config != nil {
		if err = ie.Bsr_Config.Encode(w); err != nil {
			return utils.WrapError("Encode Bsr_Config", err)
		}
	}
	if ie.Tag_Config != nil {
		if err = ie.Tag_Config.Encode(w); err != nil {
			return utils.WrapError("Encode Tag_Config", err)
		}
	}
	if ie.Phr_Config != nil {
		tmp_Phr_Config := utils.SetupRelease[*PHR_Config]{
			Setup: ie.Phr_Config,
		}
		if err = tmp_Phr_Config.Encode(w); err != nil {
			return utils.WrapError("Encode Phr_Config", err)
		}
	}
	if err = w.WriteBoolean(ie.SkipUplinkTxDynamic); err != nil {
		return utils.WrapError("WriteBoolean SkipUplinkTxDynamic", err)
	}
	if hasExtensions {
		// Extension bitmap: 5 bits for 5 extension groups
		extBitmap := []bool{ie.Csi_Mask != nil || ie.DataInactivityTimer != nil, ie.UsePreBSR_r16 != nil || ie.SchedulingRequestID_LBT_SCell_r16 != nil || ie.Lch_BasedPrioritization_r16 != nil || ie.SchedulingRequestID_BFR_SCell_r16 != nil || ie.Drx_ConfigSecondaryGroup_r16 != nil, ie.EnhancedSkipUplinkTxDynamic_r16 != nil || ie.EnhancedSkipUplinkTxConfigured_r16 != nil, ie.IntraCG_Prioritization_r17 != nil || ie.Drx_ConfigSL_r17 != nil || ie.Drx_ConfigExt_v1700 != nil || ie.SchedulingRequestID_BFR_r17 != nil || ie.SchedulingRequestID_BFR2_r17 != nil || ie.SchedulingRequestConfig_v1700 != nil || ie.Tar_Config_r17 != nil || len(ie.G_RNTI_ConfigToAddModList_r17) > 0 || len(ie.G_RNTI_ConfigToReleaseList_r17) > 0 || len(ie.G_CS_RNTI_ConfigToAddModList_r17) > 0 || len(ie.G_CS_RNTI_ConfigToReleaseList_r17) > 0 || ie.AllowCSI_SRS_Tx_MulticastDRX_Active_r17 != nil, ie.SchedulingRequestID_PosMG_Request_r17 != nil || ie.Drx_LastTransmissionUL_r17 != nil}
		if err := w.WriteExtBitMap(extBitmap); err != nil {
			return utils.WrapError("WriteExtBitMap MAC_CellGroupConfig", err)
		}

		// encode extension group 1
		if extBitmap[0] {
			extBuf := new(bytes.Buffer)
			extWriter := uper.NewWriter(extBuf)

			// Write preamble bits for optional fields in extension group 1
			optionals_ext_1 := []bool{ie.Csi_Mask != nil, ie.DataInactivityTimer != nil}
			for _, bit := range optionals_ext_1 {
				if err := extWriter.WriteBool(bit); err != nil {
					return err
				}
			}

			// encode Csi_Mask optional
			if ie.Csi_Mask != nil {
				if err = extWriter.WriteBoolean(*ie.Csi_Mask); err != nil {
					return utils.WrapError("Encode Csi_Mask", err)
				}
			}
			// encode DataInactivityTimer optional
			if ie.DataInactivityTimer != nil {
				tmp_DataInactivityTimer := utils.SetupRelease[*DataInactivityTimer]{
					Setup: ie.DataInactivityTimer,
				}
				if err = tmp_DataInactivityTimer.Encode(extWriter); err != nil {
					return utils.WrapError("Encode DataInactivityTimer", err)
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
			optionals_ext_2 := []bool{ie.UsePreBSR_r16 != nil, ie.SchedulingRequestID_LBT_SCell_r16 != nil, ie.Lch_BasedPrioritization_r16 != nil, ie.SchedulingRequestID_BFR_SCell_r16 != nil, ie.Drx_ConfigSecondaryGroup_r16 != nil}
			for _, bit := range optionals_ext_2 {
				if err := extWriter.WriteBool(bit); err != nil {
					return err
				}
			}

			// encode UsePreBSR_r16 optional
			if ie.UsePreBSR_r16 != nil {
				if err = ie.UsePreBSR_r16.Encode(extWriter); err != nil {
					return utils.WrapError("Encode UsePreBSR_r16", err)
				}
			}
			// encode SchedulingRequestID_LBT_SCell_r16 optional
			if ie.SchedulingRequestID_LBT_SCell_r16 != nil {
				if err = ie.SchedulingRequestID_LBT_SCell_r16.Encode(extWriter); err != nil {
					return utils.WrapError("Encode SchedulingRequestID_LBT_SCell_r16", err)
				}
			}
			// encode Lch_BasedPrioritization_r16 optional
			if ie.Lch_BasedPrioritization_r16 != nil {
				if err = ie.Lch_BasedPrioritization_r16.Encode(extWriter); err != nil {
					return utils.WrapError("Encode Lch_BasedPrioritization_r16", err)
				}
			}
			// encode SchedulingRequestID_BFR_SCell_r16 optional
			if ie.SchedulingRequestID_BFR_SCell_r16 != nil {
				if err = ie.SchedulingRequestID_BFR_SCell_r16.Encode(extWriter); err != nil {
					return utils.WrapError("Encode SchedulingRequestID_BFR_SCell_r16", err)
				}
			}
			// encode Drx_ConfigSecondaryGroup_r16 optional
			if ie.Drx_ConfigSecondaryGroup_r16 != nil {
				tmp_Drx_ConfigSecondaryGroup_r16 := utils.SetupRelease[*DRX_ConfigSecondaryGroup_r16]{
					Setup: ie.Drx_ConfigSecondaryGroup_r16,
				}
				if err = tmp_Drx_ConfigSecondaryGroup_r16.Encode(extWriter); err != nil {
					return utils.WrapError("Encode Drx_ConfigSecondaryGroup_r16", err)
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
			optionals_ext_3 := []bool{ie.EnhancedSkipUplinkTxDynamic_r16 != nil, ie.EnhancedSkipUplinkTxConfigured_r16 != nil}
			for _, bit := range optionals_ext_3 {
				if err := extWriter.WriteBool(bit); err != nil {
					return err
				}
			}

			// encode EnhancedSkipUplinkTxDynamic_r16 optional
			if ie.EnhancedSkipUplinkTxDynamic_r16 != nil {
				if err = ie.EnhancedSkipUplinkTxDynamic_r16.Encode(extWriter); err != nil {
					return utils.WrapError("Encode EnhancedSkipUplinkTxDynamic_r16", err)
				}
			}
			// encode EnhancedSkipUplinkTxConfigured_r16 optional
			if ie.EnhancedSkipUplinkTxConfigured_r16 != nil {
				if err = ie.EnhancedSkipUplinkTxConfigured_r16.Encode(extWriter); err != nil {
					return utils.WrapError("Encode EnhancedSkipUplinkTxConfigured_r16", err)
				}
			}

			if err := extWriter.Close(); err != nil {
				return err
			}

			if err := w.WriteOpenType(extBuf.Bytes()); err != nil {
				return err
			}
		}

		// encode extension group 4
		if extBitmap[3] {
			extBuf := new(bytes.Buffer)
			extWriter := uper.NewWriter(extBuf)

			// Write preamble bits for optional fields in extension group 4
			optionals_ext_4 := []bool{ie.IntraCG_Prioritization_r17 != nil, ie.Drx_ConfigSL_r17 != nil, ie.Drx_ConfigExt_v1700 != nil, ie.SchedulingRequestID_BFR_r17 != nil, ie.SchedulingRequestID_BFR2_r17 != nil, ie.SchedulingRequestConfig_v1700 != nil, ie.Tar_Config_r17 != nil, len(ie.G_RNTI_ConfigToAddModList_r17) > 0, len(ie.G_RNTI_ConfigToReleaseList_r17) > 0, len(ie.G_CS_RNTI_ConfigToAddModList_r17) > 0, len(ie.G_CS_RNTI_ConfigToReleaseList_r17) > 0, ie.AllowCSI_SRS_Tx_MulticastDRX_Active_r17 != nil}
			for _, bit := range optionals_ext_4 {
				if err := extWriter.WriteBool(bit); err != nil {
					return err
				}
			}

			// encode IntraCG_Prioritization_r17 optional
			if ie.IntraCG_Prioritization_r17 != nil {
				if err = ie.IntraCG_Prioritization_r17.Encode(extWriter); err != nil {
					return utils.WrapError("Encode IntraCG_Prioritization_r17", err)
				}
			}
			// encode Drx_ConfigSL_r17 optional
			if ie.Drx_ConfigSL_r17 != nil {
				tmp_Drx_ConfigSL_r17 := utils.SetupRelease[*DRX_ConfigSL_r17]{
					Setup: ie.Drx_ConfigSL_r17,
				}
				if err = tmp_Drx_ConfigSL_r17.Encode(extWriter); err != nil {
					return utils.WrapError("Encode Drx_ConfigSL_r17", err)
				}
			}
			// encode Drx_ConfigExt_v1700 optional
			if ie.Drx_ConfigExt_v1700 != nil {
				tmp_Drx_ConfigExt_v1700 := utils.SetupRelease[*DRX_ConfigExt_v1700]{
					Setup: ie.Drx_ConfigExt_v1700,
				}
				if err = tmp_Drx_ConfigExt_v1700.Encode(extWriter); err != nil {
					return utils.WrapError("Encode Drx_ConfigExt_v1700", err)
				}
			}
			// encode SchedulingRequestID_BFR_r17 optional
			if ie.SchedulingRequestID_BFR_r17 != nil {
				if err = ie.SchedulingRequestID_BFR_r17.Encode(extWriter); err != nil {
					return utils.WrapError("Encode SchedulingRequestID_BFR_r17", err)
				}
			}
			// encode SchedulingRequestID_BFR2_r17 optional
			if ie.SchedulingRequestID_BFR2_r17 != nil {
				if err = ie.SchedulingRequestID_BFR2_r17.Encode(extWriter); err != nil {
					return utils.WrapError("Encode SchedulingRequestID_BFR2_r17", err)
				}
			}
			// encode SchedulingRequestConfig_v1700 optional
			if ie.SchedulingRequestConfig_v1700 != nil {
				if err = ie.SchedulingRequestConfig_v1700.Encode(extWriter); err != nil {
					return utils.WrapError("Encode SchedulingRequestConfig_v1700", err)
				}
			}
			// encode Tar_Config_r17 optional
			if ie.Tar_Config_r17 != nil {
				tmp_Tar_Config_r17 := utils.SetupRelease[*TAR_Config_r17]{
					Setup: ie.Tar_Config_r17,
				}
				if err = tmp_Tar_Config_r17.Encode(extWriter); err != nil {
					return utils.WrapError("Encode Tar_Config_r17", err)
				}
			}
			// encode G_RNTI_ConfigToAddModList_r17 optional
			if len(ie.G_RNTI_ConfigToAddModList_r17) > 0 {
				tmp_G_RNTI_ConfigToAddModList_r17 := utils.NewSequence[*MBS_RNTI_SpecificConfig_r17]([]*MBS_RNTI_SpecificConfig_r17{}, uper.Constraint{Lb: 1, Ub: maxG_RNTI_r17}, false)
				for _, i := range ie.G_RNTI_ConfigToAddModList_r17 {
					tmp_G_RNTI_ConfigToAddModList_r17.Value = append(tmp_G_RNTI_ConfigToAddModList_r17.Value, &i)
				}
				if err = tmp_G_RNTI_ConfigToAddModList_r17.Encode(extWriter); err != nil {
					return utils.WrapError("Encode G_RNTI_ConfigToAddModList_r17", err)
				}
			}
			// encode G_RNTI_ConfigToReleaseList_r17 optional
			if len(ie.G_RNTI_ConfigToReleaseList_r17) > 0 {
				tmp_G_RNTI_ConfigToReleaseList_r17 := utils.NewSequence[*MBS_RNTI_SpecificConfigId_r17]([]*MBS_RNTI_SpecificConfigId_r17{}, uper.Constraint{Lb: 1, Ub: maxG_RNTI_r17}, false)
				for _, i := range ie.G_RNTI_ConfigToReleaseList_r17 {
					tmp_G_RNTI_ConfigToReleaseList_r17.Value = append(tmp_G_RNTI_ConfigToReleaseList_r17.Value, &i)
				}
				if err = tmp_G_RNTI_ConfigToReleaseList_r17.Encode(extWriter); err != nil {
					return utils.WrapError("Encode G_RNTI_ConfigToReleaseList_r17", err)
				}
			}
			// encode G_CS_RNTI_ConfigToAddModList_r17 optional
			if len(ie.G_CS_RNTI_ConfigToAddModList_r17) > 0 {
				tmp_G_CS_RNTI_ConfigToAddModList_r17 := utils.NewSequence[*MBS_RNTI_SpecificConfig_r17]([]*MBS_RNTI_SpecificConfig_r17{}, uper.Constraint{Lb: 1, Ub: maxG_CS_RNTI_r17}, false)
				for _, i := range ie.G_CS_RNTI_ConfigToAddModList_r17 {
					tmp_G_CS_RNTI_ConfigToAddModList_r17.Value = append(tmp_G_CS_RNTI_ConfigToAddModList_r17.Value, &i)
				}
				if err = tmp_G_CS_RNTI_ConfigToAddModList_r17.Encode(extWriter); err != nil {
					return utils.WrapError("Encode G_CS_RNTI_ConfigToAddModList_r17", err)
				}
			}
			// encode G_CS_RNTI_ConfigToReleaseList_r17 optional
			if len(ie.G_CS_RNTI_ConfigToReleaseList_r17) > 0 {
				tmp_G_CS_RNTI_ConfigToReleaseList_r17 := utils.NewSequence[*MBS_RNTI_SpecificConfigId_r17]([]*MBS_RNTI_SpecificConfigId_r17{}, uper.Constraint{Lb: 1, Ub: maxG_CS_RNTI_r17}, false)
				for _, i := range ie.G_CS_RNTI_ConfigToReleaseList_r17 {
					tmp_G_CS_RNTI_ConfigToReleaseList_r17.Value = append(tmp_G_CS_RNTI_ConfigToReleaseList_r17.Value, &i)
				}
				if err = tmp_G_CS_RNTI_ConfigToReleaseList_r17.Encode(extWriter); err != nil {
					return utils.WrapError("Encode G_CS_RNTI_ConfigToReleaseList_r17", err)
				}
			}
			// encode AllowCSI_SRS_Tx_MulticastDRX_Active_r17 optional
			if ie.AllowCSI_SRS_Tx_MulticastDRX_Active_r17 != nil {
				if err = extWriter.WriteBoolean(*ie.AllowCSI_SRS_Tx_MulticastDRX_Active_r17); err != nil {
					return utils.WrapError("Encode AllowCSI_SRS_Tx_MulticastDRX_Active_r17", err)
				}
			}

			if err := extWriter.Close(); err != nil {
				return err
			}

			if err := w.WriteOpenType(extBuf.Bytes()); err != nil {
				return err
			}
		}

		// encode extension group 5
		if extBitmap[4] {
			extBuf := new(bytes.Buffer)
			extWriter := uper.NewWriter(extBuf)

			// Write preamble bits for optional fields in extension group 5
			optionals_ext_5 := []bool{ie.SchedulingRequestID_PosMG_Request_r17 != nil, ie.Drx_LastTransmissionUL_r17 != nil}
			for _, bit := range optionals_ext_5 {
				if err := extWriter.WriteBool(bit); err != nil {
					return err
				}
			}

			// encode SchedulingRequestID_PosMG_Request_r17 optional
			if ie.SchedulingRequestID_PosMG_Request_r17 != nil {
				if err = ie.SchedulingRequestID_PosMG_Request_r17.Encode(extWriter); err != nil {
					return utils.WrapError("Encode SchedulingRequestID_PosMG_Request_r17", err)
				}
			}
			// encode Drx_LastTransmissionUL_r17 optional
			if ie.Drx_LastTransmissionUL_r17 != nil {
				if err = ie.Drx_LastTransmissionUL_r17.Encode(extWriter); err != nil {
					return utils.WrapError("Encode Drx_LastTransmissionUL_r17", err)
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

func (ie *MAC_CellGroupConfig) Decode(r *uper.UperReader) error {
	var err error
	var extensionBit bool
	if extensionBit, err = r.ReadBool(); err != nil {
		return err
	}
	var Drx_ConfigPresent bool
	if Drx_ConfigPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var SchedulingRequestConfigPresent bool
	if SchedulingRequestConfigPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var Bsr_ConfigPresent bool
	if Bsr_ConfigPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var Tag_ConfigPresent bool
	if Tag_ConfigPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var Phr_ConfigPresent bool
	if Phr_ConfigPresent, err = r.ReadBool(); err != nil {
		return err
	}
	if Drx_ConfigPresent {
		tmp_Drx_Config := utils.SetupRelease[*DRX_Config]{}
		if err = tmp_Drx_Config.Decode(r); err != nil {
			return utils.WrapError("Decode Drx_Config", err)
		}
		ie.Drx_Config = tmp_Drx_Config.Setup
	}
	if SchedulingRequestConfigPresent {
		ie.SchedulingRequestConfig = new(SchedulingRequestConfig)
		if err = ie.SchedulingRequestConfig.Decode(r); err != nil {
			return utils.WrapError("Decode SchedulingRequestConfig", err)
		}
	}
	if Bsr_ConfigPresent {
		ie.Bsr_Config = new(BSR_Config)
		if err = ie.Bsr_Config.Decode(r); err != nil {
			return utils.WrapError("Decode Bsr_Config", err)
		}
	}
	if Tag_ConfigPresent {
		ie.Tag_Config = new(TAG_Config)
		if err = ie.Tag_Config.Decode(r); err != nil {
			return utils.WrapError("Decode Tag_Config", err)
		}
	}
	if Phr_ConfigPresent {
		tmp_Phr_Config := utils.SetupRelease[*PHR_Config]{}
		if err = tmp_Phr_Config.Decode(r); err != nil {
			return utils.WrapError("Decode Phr_Config", err)
		}
		ie.Phr_Config = tmp_Phr_Config.Setup
	}
	var tmp_bool_SkipUplinkTxDynamic bool
	if tmp_bool_SkipUplinkTxDynamic, err = r.ReadBoolean(); err != nil {
		return utils.WrapError("ReadBoolean SkipUplinkTxDynamic", err)
	}
	ie.SkipUplinkTxDynamic = tmp_bool_SkipUplinkTxDynamic

	if extensionBit {
		// Read extension bitmap: 5 bits for 5 extension groups
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

			Csi_MaskPresent, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			DataInactivityTimerPresent, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			// decode Csi_Mask optional
			if Csi_MaskPresent {
				var tmp_bool_Csi_Mask bool
				if tmp_bool_Csi_Mask, err = extReader.ReadBoolean(); err != nil {
					return utils.WrapError("Decode Csi_Mask", err)
				}
				ie.Csi_Mask = &tmp_bool_Csi_Mask
			}
			// decode DataInactivityTimer optional
			if DataInactivityTimerPresent {
				tmp_DataInactivityTimer := utils.SetupRelease[*DataInactivityTimer]{}
				if err = tmp_DataInactivityTimer.Decode(extReader); err != nil {
					return utils.WrapError("Decode DataInactivityTimer", err)
				}
				ie.DataInactivityTimer = tmp_DataInactivityTimer.Setup
			}
		}
		// decode extension group 2
		if len(extBitmap) > 1 && extBitmap[1] {
			extBytes, err := r.ReadOpenType()
			if err != nil {
				return err
			}

			extReader := uper.NewReader(bytes.NewReader(extBytes))

			UsePreBSR_r16Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			SchedulingRequestID_LBT_SCell_r16Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			Lch_BasedPrioritization_r16Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			SchedulingRequestID_BFR_SCell_r16Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			Drx_ConfigSecondaryGroup_r16Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			// decode UsePreBSR_r16 optional
			if UsePreBSR_r16Present {
				ie.UsePreBSR_r16 = new(MAC_CellGroupConfig_usePreBSR_r16)
				if err = ie.UsePreBSR_r16.Decode(extReader); err != nil {
					return utils.WrapError("Decode UsePreBSR_r16", err)
				}
			}
			// decode SchedulingRequestID_LBT_SCell_r16 optional
			if SchedulingRequestID_LBT_SCell_r16Present {
				ie.SchedulingRequestID_LBT_SCell_r16 = new(SchedulingRequestId)
				if err = ie.SchedulingRequestID_LBT_SCell_r16.Decode(extReader); err != nil {
					return utils.WrapError("Decode SchedulingRequestID_LBT_SCell_r16", err)
				}
			}
			// decode Lch_BasedPrioritization_r16 optional
			if Lch_BasedPrioritization_r16Present {
				ie.Lch_BasedPrioritization_r16 = new(MAC_CellGroupConfig_lch_BasedPrioritization_r16)
				if err = ie.Lch_BasedPrioritization_r16.Decode(extReader); err != nil {
					return utils.WrapError("Decode Lch_BasedPrioritization_r16", err)
				}
			}
			// decode SchedulingRequestID_BFR_SCell_r16 optional
			if SchedulingRequestID_BFR_SCell_r16Present {
				ie.SchedulingRequestID_BFR_SCell_r16 = new(SchedulingRequestId)
				if err = ie.SchedulingRequestID_BFR_SCell_r16.Decode(extReader); err != nil {
					return utils.WrapError("Decode SchedulingRequestID_BFR_SCell_r16", err)
				}
			}
			// decode Drx_ConfigSecondaryGroup_r16 optional
			if Drx_ConfigSecondaryGroup_r16Present {
				tmp_Drx_ConfigSecondaryGroup_r16 := utils.SetupRelease[*DRX_ConfigSecondaryGroup_r16]{}
				if err = tmp_Drx_ConfigSecondaryGroup_r16.Decode(extReader); err != nil {
					return utils.WrapError("Decode Drx_ConfigSecondaryGroup_r16", err)
				}
				ie.Drx_ConfigSecondaryGroup_r16 = tmp_Drx_ConfigSecondaryGroup_r16.Setup
			}
		}
		// decode extension group 3
		if len(extBitmap) > 2 && extBitmap[2] {
			extBytes, err := r.ReadOpenType()
			if err != nil {
				return err
			}

			extReader := uper.NewReader(bytes.NewReader(extBytes))

			EnhancedSkipUplinkTxDynamic_r16Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			EnhancedSkipUplinkTxConfigured_r16Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			// decode EnhancedSkipUplinkTxDynamic_r16 optional
			if EnhancedSkipUplinkTxDynamic_r16Present {
				ie.EnhancedSkipUplinkTxDynamic_r16 = new(MAC_CellGroupConfig_enhancedSkipUplinkTxDynamic_r16)
				if err = ie.EnhancedSkipUplinkTxDynamic_r16.Decode(extReader); err != nil {
					return utils.WrapError("Decode EnhancedSkipUplinkTxDynamic_r16", err)
				}
			}
			// decode EnhancedSkipUplinkTxConfigured_r16 optional
			if EnhancedSkipUplinkTxConfigured_r16Present {
				ie.EnhancedSkipUplinkTxConfigured_r16 = new(MAC_CellGroupConfig_enhancedSkipUplinkTxConfigured_r16)
				if err = ie.EnhancedSkipUplinkTxConfigured_r16.Decode(extReader); err != nil {
					return utils.WrapError("Decode EnhancedSkipUplinkTxConfigured_r16", err)
				}
			}
		}
		// decode extension group 4
		if len(extBitmap) > 3 && extBitmap[3] {
			extBytes, err := r.ReadOpenType()
			if err != nil {
				return err
			}

			extReader := uper.NewReader(bytes.NewReader(extBytes))

			IntraCG_Prioritization_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			Drx_ConfigSL_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			Drx_ConfigExt_v1700Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			SchedulingRequestID_BFR_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			SchedulingRequestID_BFR2_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			SchedulingRequestConfig_v1700Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			Tar_Config_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			G_RNTI_ConfigToAddModList_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			G_RNTI_ConfigToReleaseList_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			G_CS_RNTI_ConfigToAddModList_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			G_CS_RNTI_ConfigToReleaseList_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			AllowCSI_SRS_Tx_MulticastDRX_Active_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			// decode IntraCG_Prioritization_r17 optional
			if IntraCG_Prioritization_r17Present {
				ie.IntraCG_Prioritization_r17 = new(MAC_CellGroupConfig_intraCG_Prioritization_r17)
				if err = ie.IntraCG_Prioritization_r17.Decode(extReader); err != nil {
					return utils.WrapError("Decode IntraCG_Prioritization_r17", err)
				}
			}
			// decode Drx_ConfigSL_r17 optional
			if Drx_ConfigSL_r17Present {
				tmp_Drx_ConfigSL_r17 := utils.SetupRelease[*DRX_ConfigSL_r17]{}
				if err = tmp_Drx_ConfigSL_r17.Decode(extReader); err != nil {
					return utils.WrapError("Decode Drx_ConfigSL_r17", err)
				}
				ie.Drx_ConfigSL_r17 = tmp_Drx_ConfigSL_r17.Setup
			}
			// decode Drx_ConfigExt_v1700 optional
			if Drx_ConfigExt_v1700Present {
				tmp_Drx_ConfigExt_v1700 := utils.SetupRelease[*DRX_ConfigExt_v1700]{}
				if err = tmp_Drx_ConfigExt_v1700.Decode(extReader); err != nil {
					return utils.WrapError("Decode Drx_ConfigExt_v1700", err)
				}
				ie.Drx_ConfigExt_v1700 = tmp_Drx_ConfigExt_v1700.Setup
			}
			// decode SchedulingRequestID_BFR_r17 optional
			if SchedulingRequestID_BFR_r17Present {
				ie.SchedulingRequestID_BFR_r17 = new(SchedulingRequestId)
				if err = ie.SchedulingRequestID_BFR_r17.Decode(extReader); err != nil {
					return utils.WrapError("Decode SchedulingRequestID_BFR_r17", err)
				}
			}
			// decode SchedulingRequestID_BFR2_r17 optional
			if SchedulingRequestID_BFR2_r17Present {
				ie.SchedulingRequestID_BFR2_r17 = new(SchedulingRequestId)
				if err = ie.SchedulingRequestID_BFR2_r17.Decode(extReader); err != nil {
					return utils.WrapError("Decode SchedulingRequestID_BFR2_r17", err)
				}
			}
			// decode SchedulingRequestConfig_v1700 optional
			if SchedulingRequestConfig_v1700Present {
				ie.SchedulingRequestConfig_v1700 = new(SchedulingRequestConfig_v1700)
				if err = ie.SchedulingRequestConfig_v1700.Decode(extReader); err != nil {
					return utils.WrapError("Decode SchedulingRequestConfig_v1700", err)
				}
			}
			// decode Tar_Config_r17 optional
			if Tar_Config_r17Present {
				tmp_Tar_Config_r17 := utils.SetupRelease[*TAR_Config_r17]{}
				if err = tmp_Tar_Config_r17.Decode(extReader); err != nil {
					return utils.WrapError("Decode Tar_Config_r17", err)
				}
				ie.Tar_Config_r17 = tmp_Tar_Config_r17.Setup
			}
			// decode G_RNTI_ConfigToAddModList_r17 optional
			if G_RNTI_ConfigToAddModList_r17Present {
				tmp_G_RNTI_ConfigToAddModList_r17 := utils.NewSequence[*MBS_RNTI_SpecificConfig_r17]([]*MBS_RNTI_SpecificConfig_r17{}, uper.Constraint{Lb: 1, Ub: maxG_RNTI_r17}, false)
				fn_G_RNTI_ConfigToAddModList_r17 := func() *MBS_RNTI_SpecificConfig_r17 {
					return new(MBS_RNTI_SpecificConfig_r17)
				}
				if err = tmp_G_RNTI_ConfigToAddModList_r17.Decode(extReader, fn_G_RNTI_ConfigToAddModList_r17); err != nil {
					return utils.WrapError("Decode G_RNTI_ConfigToAddModList_r17", err)
				}
				ie.G_RNTI_ConfigToAddModList_r17 = []MBS_RNTI_SpecificConfig_r17{}
				for _, i := range tmp_G_RNTI_ConfigToAddModList_r17.Value {
					ie.G_RNTI_ConfigToAddModList_r17 = append(ie.G_RNTI_ConfigToAddModList_r17, *i)
				}
			}
			// decode G_RNTI_ConfigToReleaseList_r17 optional
			if G_RNTI_ConfigToReleaseList_r17Present {
				tmp_G_RNTI_ConfigToReleaseList_r17 := utils.NewSequence[*MBS_RNTI_SpecificConfigId_r17]([]*MBS_RNTI_SpecificConfigId_r17{}, uper.Constraint{Lb: 1, Ub: maxG_RNTI_r17}, false)
				fn_G_RNTI_ConfigToReleaseList_r17 := func() *MBS_RNTI_SpecificConfigId_r17 {
					return new(MBS_RNTI_SpecificConfigId_r17)
				}
				if err = tmp_G_RNTI_ConfigToReleaseList_r17.Decode(extReader, fn_G_RNTI_ConfigToReleaseList_r17); err != nil {
					return utils.WrapError("Decode G_RNTI_ConfigToReleaseList_r17", err)
				}
				ie.G_RNTI_ConfigToReleaseList_r17 = []MBS_RNTI_SpecificConfigId_r17{}
				for _, i := range tmp_G_RNTI_ConfigToReleaseList_r17.Value {
					ie.G_RNTI_ConfigToReleaseList_r17 = append(ie.G_RNTI_ConfigToReleaseList_r17, *i)
				}
			}
			// decode G_CS_RNTI_ConfigToAddModList_r17 optional
			if G_CS_RNTI_ConfigToAddModList_r17Present {
				tmp_G_CS_RNTI_ConfigToAddModList_r17 := utils.NewSequence[*MBS_RNTI_SpecificConfig_r17]([]*MBS_RNTI_SpecificConfig_r17{}, uper.Constraint{Lb: 1, Ub: maxG_CS_RNTI_r17}, false)
				fn_G_CS_RNTI_ConfigToAddModList_r17 := func() *MBS_RNTI_SpecificConfig_r17 {
					return new(MBS_RNTI_SpecificConfig_r17)
				}
				if err = tmp_G_CS_RNTI_ConfigToAddModList_r17.Decode(extReader, fn_G_CS_RNTI_ConfigToAddModList_r17); err != nil {
					return utils.WrapError("Decode G_CS_RNTI_ConfigToAddModList_r17", err)
				}
				ie.G_CS_RNTI_ConfigToAddModList_r17 = []MBS_RNTI_SpecificConfig_r17{}
				for _, i := range tmp_G_CS_RNTI_ConfigToAddModList_r17.Value {
					ie.G_CS_RNTI_ConfigToAddModList_r17 = append(ie.G_CS_RNTI_ConfigToAddModList_r17, *i)
				}
			}
			// decode G_CS_RNTI_ConfigToReleaseList_r17 optional
			if G_CS_RNTI_ConfigToReleaseList_r17Present {
				tmp_G_CS_RNTI_ConfigToReleaseList_r17 := utils.NewSequence[*MBS_RNTI_SpecificConfigId_r17]([]*MBS_RNTI_SpecificConfigId_r17{}, uper.Constraint{Lb: 1, Ub: maxG_CS_RNTI_r17}, false)
				fn_G_CS_RNTI_ConfigToReleaseList_r17 := func() *MBS_RNTI_SpecificConfigId_r17 {
					return new(MBS_RNTI_SpecificConfigId_r17)
				}
				if err = tmp_G_CS_RNTI_ConfigToReleaseList_r17.Decode(extReader, fn_G_CS_RNTI_ConfigToReleaseList_r17); err != nil {
					return utils.WrapError("Decode G_CS_RNTI_ConfigToReleaseList_r17", err)
				}
				ie.G_CS_RNTI_ConfigToReleaseList_r17 = []MBS_RNTI_SpecificConfigId_r17{}
				for _, i := range tmp_G_CS_RNTI_ConfigToReleaseList_r17.Value {
					ie.G_CS_RNTI_ConfigToReleaseList_r17 = append(ie.G_CS_RNTI_ConfigToReleaseList_r17, *i)
				}
			}
			// decode AllowCSI_SRS_Tx_MulticastDRX_Active_r17 optional
			if AllowCSI_SRS_Tx_MulticastDRX_Active_r17Present {
				var tmp_bool_AllowCSI_SRS_Tx_MulticastDRX_Active_r17 bool
				if tmp_bool_AllowCSI_SRS_Tx_MulticastDRX_Active_r17, err = extReader.ReadBoolean(); err != nil {
					return utils.WrapError("Decode AllowCSI_SRS_Tx_MulticastDRX_Active_r17", err)
				}
				ie.AllowCSI_SRS_Tx_MulticastDRX_Active_r17 = &tmp_bool_AllowCSI_SRS_Tx_MulticastDRX_Active_r17
			}
		}
		// decode extension group 5
		if len(extBitmap) > 4 && extBitmap[4] {
			extBytes, err := r.ReadOpenType()
			if err != nil {
				return err
			}

			extReader := uper.NewReader(bytes.NewReader(extBytes))

			SchedulingRequestID_PosMG_Request_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			Drx_LastTransmissionUL_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			// decode SchedulingRequestID_PosMG_Request_r17 optional
			if SchedulingRequestID_PosMG_Request_r17Present {
				ie.SchedulingRequestID_PosMG_Request_r17 = new(SchedulingRequestId)
				if err = ie.SchedulingRequestID_PosMG_Request_r17.Decode(extReader); err != nil {
					return utils.WrapError("Decode SchedulingRequestID_PosMG_Request_r17", err)
				}
			}
			// decode Drx_LastTransmissionUL_r17 optional
			if Drx_LastTransmissionUL_r17Present {
				ie.Drx_LastTransmissionUL_r17 = new(MAC_CellGroupConfig_drx_LastTransmissionUL_r17)
				if err = ie.Drx_LastTransmissionUL_r17.Decode(extReader); err != nil {
					return utils.WrapError("Decode Drx_LastTransmissionUL_r17", err)
				}
			}
		}
	}
	return nil
}
