package ies

import (
	"bytes"

	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type BWP_UplinkDedicated struct {
	Pucch_Config                                        *PUCCH_Config                                        `optional,setuprelease`
	Pusch_Config                                        *PUSCH_Config                                        `optional,setuprelease`
	ConfiguredGrantConfig                               *ConfiguredGrantConfig                               `optional,setuprelease`
	Srs_Config                                          *SRS_Config                                          `optional,setuprelease`
	BeamFailureRecoveryConfig                           *BeamFailureRecoveryConfig                           `optional,setuprelease`
	Sl_PUCCH_Config_r16                                 *PUCCH_Config                                        `optional,ext-1,setuprelease`
	Cp_ExtensionC2_r16                                  *int64                                               `lb:1,ub:28,optional,ext-1`
	Cp_ExtensionC3_r16                                  *int64                                               `lb:1,ub:28,optional,ext-1`
	UseInterlacePUCCH_PUSCH_r16                         *BWP_UplinkDedicated_useInterlacePUCCH_PUSCH_r16     `optional,ext-1`
	Pucch_ConfigurationList_r16                         *PUCCH_ConfigurationList_r16                         `optional,ext-1,setuprelease`
	Lbt_FailureRecoveryConfig_r16                       *LBT_FailureRecoveryConfig_r16                       `optional,ext-1,setuprelease`
	ConfiguredGrantConfigToAddModList_r16               *ConfiguredGrantConfigToAddModList_r16               `optional,ext-1`
	ConfiguredGrantConfigToReleaseList_r16              *ConfiguredGrantConfigToReleaseList_r16              `optional,ext-1`
	ConfiguredGrantConfigType2DeactivationStateList_r16 *ConfiguredGrantConfigType2DeactivationStateList_r16 `optional,ext-1`
	Ul_TCI_StateList_r17                                *BWP_UplinkDedicated_ul_TCI_StateList_r17            `lb:1,ub:maxUL_TCI_r17,optional,ext-2`
	Ul_powerControl_r17                                 *Uplink_powerControlId_r17                           `optional,ext-2`
	Pucch_ConfigurationListMulticast1_r17               *PUCCH_ConfigurationList_r16                         `optional,ext-2,setuprelease`
	Pucch_ConfigurationListMulticast2_r17               *PUCCH_ConfigurationList_r16                         `optional,ext-2,setuprelease`
	Pucch_ConfigMulticast1_r17                          *PUCCH_Config                                        `optional,ext-3,setuprelease`
	Pucch_ConfigMulticast2_r17                          *PUCCH_Config                                        `optional,ext-3,setuprelease`
	PathlossReferenceRSToAddModList_r17                 []PathlossReferenceRS_r17                            `lb:1,ub:maxNrofPathlossReferenceRSs_r17,optional,ext-4`
	PathlossReferenceRSToReleaseList_r17                []PathlossReferenceRS_Id_r17                         `lb:1,ub:maxNrofPathlossReferenceRSs_r17,optional,ext-4`
}

func (ie *BWP_UplinkDedicated) Encode(w *uper.UperWriter) error {
	var err error
	hasExtensions := ie.Sl_PUCCH_Config_r16 != nil || ie.Cp_ExtensionC2_r16 != nil || ie.Cp_ExtensionC3_r16 != nil || ie.UseInterlacePUCCH_PUSCH_r16 != nil || ie.Pucch_ConfigurationList_r16 != nil || ie.Lbt_FailureRecoveryConfig_r16 != nil || ie.ConfiguredGrantConfigToAddModList_r16 != nil || ie.ConfiguredGrantConfigToReleaseList_r16 != nil || ie.ConfiguredGrantConfigType2DeactivationStateList_r16 != nil || ie.Ul_TCI_StateList_r17 != nil || ie.Ul_powerControl_r17 != nil || ie.Pucch_ConfigurationListMulticast1_r17 != nil || ie.Pucch_ConfigurationListMulticast2_r17 != nil || ie.Pucch_ConfigMulticast1_r17 != nil || ie.Pucch_ConfigMulticast2_r17 != nil || len(ie.PathlossReferenceRSToAddModList_r17) > 0 || len(ie.PathlossReferenceRSToReleaseList_r17) > 0
	preambleBits := []bool{hasExtensions, ie.Pucch_Config != nil, ie.Pusch_Config != nil, ie.ConfiguredGrantConfig != nil, ie.Srs_Config != nil, ie.BeamFailureRecoveryConfig != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.Pucch_Config != nil {
		tmp_Pucch_Config := utils.SetupRelease[*PUCCH_Config]{
			Setup: ie.Pucch_Config,
		}
		if err = tmp_Pucch_Config.Encode(w); err != nil {
			return utils.WrapError("Encode Pucch_Config", err)
		}
	}
	if ie.Pusch_Config != nil {
		tmp_Pusch_Config := utils.SetupRelease[*PUSCH_Config]{
			Setup: ie.Pusch_Config,
		}
		if err = tmp_Pusch_Config.Encode(w); err != nil {
			return utils.WrapError("Encode Pusch_Config", err)
		}
	}
	if ie.ConfiguredGrantConfig != nil {
		tmp_ConfiguredGrantConfig := utils.SetupRelease[*ConfiguredGrantConfig]{
			Setup: ie.ConfiguredGrantConfig,
		}
		if err = tmp_ConfiguredGrantConfig.Encode(w); err != nil {
			return utils.WrapError("Encode ConfiguredGrantConfig", err)
		}
	}
	if ie.Srs_Config != nil {
		tmp_Srs_Config := utils.SetupRelease[*SRS_Config]{
			Setup: ie.Srs_Config,
		}
		if err = tmp_Srs_Config.Encode(w); err != nil {
			return utils.WrapError("Encode Srs_Config", err)
		}
	}
	if ie.BeamFailureRecoveryConfig != nil {
		tmp_BeamFailureRecoveryConfig := utils.SetupRelease[*BeamFailureRecoveryConfig]{
			Setup: ie.BeamFailureRecoveryConfig,
		}
		if err = tmp_BeamFailureRecoveryConfig.Encode(w); err != nil {
			return utils.WrapError("Encode BeamFailureRecoveryConfig", err)
		}
	}
	if hasExtensions {
		// Extension bitmap: 4 bits for 4 extension groups
		extBitmap := []bool{ie.Sl_PUCCH_Config_r16 != nil || ie.Cp_ExtensionC2_r16 != nil || ie.Cp_ExtensionC3_r16 != nil || ie.UseInterlacePUCCH_PUSCH_r16 != nil || ie.Pucch_ConfigurationList_r16 != nil || ie.Lbt_FailureRecoveryConfig_r16 != nil || ie.ConfiguredGrantConfigToAddModList_r16 != nil || ie.ConfiguredGrantConfigToReleaseList_r16 != nil || ie.ConfiguredGrantConfigType2DeactivationStateList_r16 != nil, ie.Ul_TCI_StateList_r17 != nil || ie.Ul_powerControl_r17 != nil || ie.Pucch_ConfigurationListMulticast1_r17 != nil || ie.Pucch_ConfigurationListMulticast2_r17 != nil, ie.Pucch_ConfigMulticast1_r17 != nil || ie.Pucch_ConfigMulticast2_r17 != nil, len(ie.PathlossReferenceRSToAddModList_r17) > 0 || len(ie.PathlossReferenceRSToReleaseList_r17) > 0}
		if err := w.WriteExtBitMap(extBitmap); err != nil {
			return utils.WrapError("WriteExtBitMap BWP_UplinkDedicated", err)
		}

		// encode extension group 1
		if extBitmap[0] {
			extBuf := new(bytes.Buffer)
			extWriter := uper.NewWriter(extBuf)

			// Write preamble bits for optional fields in extension group 1
			optionals_ext_1 := []bool{ie.Sl_PUCCH_Config_r16 != nil, ie.Cp_ExtensionC2_r16 != nil, ie.Cp_ExtensionC3_r16 != nil, ie.UseInterlacePUCCH_PUSCH_r16 != nil, ie.Pucch_ConfigurationList_r16 != nil, ie.Lbt_FailureRecoveryConfig_r16 != nil, ie.ConfiguredGrantConfigToAddModList_r16 != nil, ie.ConfiguredGrantConfigToReleaseList_r16 != nil, ie.ConfiguredGrantConfigType2DeactivationStateList_r16 != nil}
			for _, bit := range optionals_ext_1 {
				if err := extWriter.WriteBool(bit); err != nil {
					return err
				}
			}

			// encode Sl_PUCCH_Config_r16 optional
			if ie.Sl_PUCCH_Config_r16 != nil {
				tmp_Sl_PUCCH_Config_r16 := utils.SetupRelease[*PUCCH_Config]{
					Setup: ie.Sl_PUCCH_Config_r16,
				}
				if err = tmp_Sl_PUCCH_Config_r16.Encode(extWriter); err != nil {
					return utils.WrapError("Encode Sl_PUCCH_Config_r16", err)
				}
			}
			// encode Cp_ExtensionC2_r16 optional
			if ie.Cp_ExtensionC2_r16 != nil {
				if err = extWriter.WriteInteger(*ie.Cp_ExtensionC2_r16, &uper.Constraint{Lb: 1, Ub: 28}, false); err != nil {
					return utils.WrapError("Encode Cp_ExtensionC2_r16", err)
				}
			}
			// encode Cp_ExtensionC3_r16 optional
			if ie.Cp_ExtensionC3_r16 != nil {
				if err = extWriter.WriteInteger(*ie.Cp_ExtensionC3_r16, &uper.Constraint{Lb: 1, Ub: 28}, false); err != nil {
					return utils.WrapError("Encode Cp_ExtensionC3_r16", err)
				}
			}
			// encode UseInterlacePUCCH_PUSCH_r16 optional
			if ie.UseInterlacePUCCH_PUSCH_r16 != nil {
				if err = ie.UseInterlacePUCCH_PUSCH_r16.Encode(extWriter); err != nil {
					return utils.WrapError("Encode UseInterlacePUCCH_PUSCH_r16", err)
				}
			}
			// encode Pucch_ConfigurationList_r16 optional
			if ie.Pucch_ConfigurationList_r16 != nil {
				tmp_Pucch_ConfigurationList_r16 := utils.SetupRelease[*PUCCH_ConfigurationList_r16]{
					Setup: ie.Pucch_ConfigurationList_r16,
				}
				if err = tmp_Pucch_ConfigurationList_r16.Encode(extWriter); err != nil {
					return utils.WrapError("Encode Pucch_ConfigurationList_r16", err)
				}
			}
			// encode Lbt_FailureRecoveryConfig_r16 optional
			if ie.Lbt_FailureRecoveryConfig_r16 != nil {
				tmp_Lbt_FailureRecoveryConfig_r16 := utils.SetupRelease[*LBT_FailureRecoveryConfig_r16]{
					Setup: ie.Lbt_FailureRecoveryConfig_r16,
				}
				if err = tmp_Lbt_FailureRecoveryConfig_r16.Encode(extWriter); err != nil {
					return utils.WrapError("Encode Lbt_FailureRecoveryConfig_r16", err)
				}
			}
			// encode ConfiguredGrantConfigToAddModList_r16 optional
			if ie.ConfiguredGrantConfigToAddModList_r16 != nil {
				if err = ie.ConfiguredGrantConfigToAddModList_r16.Encode(extWriter); err != nil {
					return utils.WrapError("Encode ConfiguredGrantConfigToAddModList_r16", err)
				}
			}
			// encode ConfiguredGrantConfigToReleaseList_r16 optional
			if ie.ConfiguredGrantConfigToReleaseList_r16 != nil {
				if err = ie.ConfiguredGrantConfigToReleaseList_r16.Encode(extWriter); err != nil {
					return utils.WrapError("Encode ConfiguredGrantConfigToReleaseList_r16", err)
				}
			}
			// encode ConfiguredGrantConfigType2DeactivationStateList_r16 optional
			if ie.ConfiguredGrantConfigType2DeactivationStateList_r16 != nil {
				if err = ie.ConfiguredGrantConfigType2DeactivationStateList_r16.Encode(extWriter); err != nil {
					return utils.WrapError("Encode ConfiguredGrantConfigType2DeactivationStateList_r16", err)
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
			optionals_ext_2 := []bool{ie.Ul_TCI_StateList_r17 != nil, ie.Ul_powerControl_r17 != nil, ie.Pucch_ConfigurationListMulticast1_r17 != nil, ie.Pucch_ConfigurationListMulticast2_r17 != nil}
			for _, bit := range optionals_ext_2 {
				if err := extWriter.WriteBool(bit); err != nil {
					return err
				}
			}

			// encode Ul_TCI_StateList_r17 optional
			if ie.Ul_TCI_StateList_r17 != nil {
				if err = ie.Ul_TCI_StateList_r17.Encode(extWriter); err != nil {
					return utils.WrapError("Encode Ul_TCI_StateList_r17", err)
				}
			}
			// encode Ul_powerControl_r17 optional
			if ie.Ul_powerControl_r17 != nil {
				if err = ie.Ul_powerControl_r17.Encode(extWriter); err != nil {
					return utils.WrapError("Encode Ul_powerControl_r17", err)
				}
			}
			// encode Pucch_ConfigurationListMulticast1_r17 optional
			if ie.Pucch_ConfigurationListMulticast1_r17 != nil {
				tmp_Pucch_ConfigurationListMulticast1_r17 := utils.SetupRelease[*PUCCH_ConfigurationList_r16]{
					Setup: ie.Pucch_ConfigurationListMulticast1_r17,
				}
				if err = tmp_Pucch_ConfigurationListMulticast1_r17.Encode(extWriter); err != nil {
					return utils.WrapError("Encode Pucch_ConfigurationListMulticast1_r17", err)
				}
			}
			// encode Pucch_ConfigurationListMulticast2_r17 optional
			if ie.Pucch_ConfigurationListMulticast2_r17 != nil {
				tmp_Pucch_ConfigurationListMulticast2_r17 := utils.SetupRelease[*PUCCH_ConfigurationList_r16]{
					Setup: ie.Pucch_ConfigurationListMulticast2_r17,
				}
				if err = tmp_Pucch_ConfigurationListMulticast2_r17.Encode(extWriter); err != nil {
					return utils.WrapError("Encode Pucch_ConfigurationListMulticast2_r17", err)
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
			optionals_ext_3 := []bool{ie.Pucch_ConfigMulticast1_r17 != nil, ie.Pucch_ConfigMulticast2_r17 != nil}
			for _, bit := range optionals_ext_3 {
				if err := extWriter.WriteBool(bit); err != nil {
					return err
				}
			}

			// encode Pucch_ConfigMulticast1_r17 optional
			if ie.Pucch_ConfigMulticast1_r17 != nil {
				tmp_Pucch_ConfigMulticast1_r17 := utils.SetupRelease[*PUCCH_Config]{
					Setup: ie.Pucch_ConfigMulticast1_r17,
				}
				if err = tmp_Pucch_ConfigMulticast1_r17.Encode(extWriter); err != nil {
					return utils.WrapError("Encode Pucch_ConfigMulticast1_r17", err)
				}
			}
			// encode Pucch_ConfigMulticast2_r17 optional
			if ie.Pucch_ConfigMulticast2_r17 != nil {
				tmp_Pucch_ConfigMulticast2_r17 := utils.SetupRelease[*PUCCH_Config]{
					Setup: ie.Pucch_ConfigMulticast2_r17,
				}
				if err = tmp_Pucch_ConfigMulticast2_r17.Encode(extWriter); err != nil {
					return utils.WrapError("Encode Pucch_ConfigMulticast2_r17", err)
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
			optionals_ext_4 := []bool{len(ie.PathlossReferenceRSToAddModList_r17) > 0, len(ie.PathlossReferenceRSToReleaseList_r17) > 0}
			for _, bit := range optionals_ext_4 {
				if err := extWriter.WriteBool(bit); err != nil {
					return err
				}
			}

			// encode PathlossReferenceRSToAddModList_r17 optional
			if len(ie.PathlossReferenceRSToAddModList_r17) > 0 {
				tmp_PathlossReferenceRSToAddModList_r17 := utils.NewSequence[*PathlossReferenceRS_r17]([]*PathlossReferenceRS_r17{}, uper.Constraint{Lb: 1, Ub: maxNrofPathlossReferenceRSs_r17}, false)
				for _, i := range ie.PathlossReferenceRSToAddModList_r17 {
					tmp_PathlossReferenceRSToAddModList_r17.Value = append(tmp_PathlossReferenceRSToAddModList_r17.Value, &i)
				}
				if err = tmp_PathlossReferenceRSToAddModList_r17.Encode(extWriter); err != nil {
					return utils.WrapError("Encode PathlossReferenceRSToAddModList_r17", err)
				}
			}
			// encode PathlossReferenceRSToReleaseList_r17 optional
			if len(ie.PathlossReferenceRSToReleaseList_r17) > 0 {
				tmp_PathlossReferenceRSToReleaseList_r17 := utils.NewSequence[*PathlossReferenceRS_Id_r17]([]*PathlossReferenceRS_Id_r17{}, uper.Constraint{Lb: 1, Ub: maxNrofPathlossReferenceRSs_r17}, false)
				for _, i := range ie.PathlossReferenceRSToReleaseList_r17 {
					tmp_PathlossReferenceRSToReleaseList_r17.Value = append(tmp_PathlossReferenceRSToReleaseList_r17.Value, &i)
				}
				if err = tmp_PathlossReferenceRSToReleaseList_r17.Encode(extWriter); err != nil {
					return utils.WrapError("Encode PathlossReferenceRSToReleaseList_r17", err)
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

func (ie *BWP_UplinkDedicated) Decode(r *uper.UperReader) error {
	var err error
	var extensionBit bool
	if extensionBit, err = r.ReadBool(); err != nil {
		return err
	}
	var Pucch_ConfigPresent bool
	if Pucch_ConfigPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var Pusch_ConfigPresent bool
	if Pusch_ConfigPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var ConfiguredGrantConfigPresent bool
	if ConfiguredGrantConfigPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var Srs_ConfigPresent bool
	if Srs_ConfigPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var BeamFailureRecoveryConfigPresent bool
	if BeamFailureRecoveryConfigPresent, err = r.ReadBool(); err != nil {
		return err
	}
	if Pucch_ConfigPresent {
		tmp_Pucch_Config := utils.SetupRelease[*PUCCH_Config]{}
		if err = tmp_Pucch_Config.Decode(r); err != nil {
			return utils.WrapError("Decode Pucch_Config", err)
		}
		ie.Pucch_Config = tmp_Pucch_Config.Setup
	}
	if Pusch_ConfigPresent {
		tmp_Pusch_Config := utils.SetupRelease[*PUSCH_Config]{}
		if err = tmp_Pusch_Config.Decode(r); err != nil {
			return utils.WrapError("Decode Pusch_Config", err)
		}
		ie.Pusch_Config = tmp_Pusch_Config.Setup
	}
	if ConfiguredGrantConfigPresent {
		tmp_ConfiguredGrantConfig := utils.SetupRelease[*ConfiguredGrantConfig]{}
		if err = tmp_ConfiguredGrantConfig.Decode(r); err != nil {
			return utils.WrapError("Decode ConfiguredGrantConfig", err)
		}
		ie.ConfiguredGrantConfig = tmp_ConfiguredGrantConfig.Setup
	}
	if Srs_ConfigPresent {
		tmp_Srs_Config := utils.SetupRelease[*SRS_Config]{}
		if err = tmp_Srs_Config.Decode(r); err != nil {
			return utils.WrapError("Decode Srs_Config", err)
		}
		ie.Srs_Config = tmp_Srs_Config.Setup
	}
	if BeamFailureRecoveryConfigPresent {
		tmp_BeamFailureRecoveryConfig := utils.SetupRelease[*BeamFailureRecoveryConfig]{}
		if err = tmp_BeamFailureRecoveryConfig.Decode(r); err != nil {
			return utils.WrapError("Decode BeamFailureRecoveryConfig", err)
		}
		ie.BeamFailureRecoveryConfig = tmp_BeamFailureRecoveryConfig.Setup
	}

	if extensionBit {
		// Read extension bitmap: 4 bits for 4 extension groups
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

			Sl_PUCCH_Config_r16Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			Cp_ExtensionC2_r16Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			Cp_ExtensionC3_r16Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			UseInterlacePUCCH_PUSCH_r16Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			Pucch_ConfigurationList_r16Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			Lbt_FailureRecoveryConfig_r16Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			ConfiguredGrantConfigToAddModList_r16Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			ConfiguredGrantConfigToReleaseList_r16Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			ConfiguredGrantConfigType2DeactivationStateList_r16Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			// decode Sl_PUCCH_Config_r16 optional
			if Sl_PUCCH_Config_r16Present {
				tmp_Sl_PUCCH_Config_r16 := utils.SetupRelease[*PUCCH_Config]{}
				if err = tmp_Sl_PUCCH_Config_r16.Decode(extReader); err != nil {
					return utils.WrapError("Decode Sl_PUCCH_Config_r16", err)
				}
				ie.Sl_PUCCH_Config_r16 = tmp_Sl_PUCCH_Config_r16.Setup
			}
			// decode Cp_ExtensionC2_r16 optional
			if Cp_ExtensionC2_r16Present {
				var tmp_int_Cp_ExtensionC2_r16 int64
				if tmp_int_Cp_ExtensionC2_r16, err = extReader.ReadInteger(&uper.Constraint{Lb: 1, Ub: 28}, false); err != nil {
					return utils.WrapError("Decode Cp_ExtensionC2_r16", err)
				}
				ie.Cp_ExtensionC2_r16 = &tmp_int_Cp_ExtensionC2_r16
			}
			// decode Cp_ExtensionC3_r16 optional
			if Cp_ExtensionC3_r16Present {
				var tmp_int_Cp_ExtensionC3_r16 int64
				if tmp_int_Cp_ExtensionC3_r16, err = extReader.ReadInteger(&uper.Constraint{Lb: 1, Ub: 28}, false); err != nil {
					return utils.WrapError("Decode Cp_ExtensionC3_r16", err)
				}
				ie.Cp_ExtensionC3_r16 = &tmp_int_Cp_ExtensionC3_r16
			}
			// decode UseInterlacePUCCH_PUSCH_r16 optional
			if UseInterlacePUCCH_PUSCH_r16Present {
				ie.UseInterlacePUCCH_PUSCH_r16 = new(BWP_UplinkDedicated_useInterlacePUCCH_PUSCH_r16)
				if err = ie.UseInterlacePUCCH_PUSCH_r16.Decode(extReader); err != nil {
					return utils.WrapError("Decode UseInterlacePUCCH_PUSCH_r16", err)
				}
			}
			// decode Pucch_ConfigurationList_r16 optional
			if Pucch_ConfigurationList_r16Present {
				tmp_Pucch_ConfigurationList_r16 := utils.SetupRelease[*PUCCH_ConfigurationList_r16]{}
				if err = tmp_Pucch_ConfigurationList_r16.Decode(extReader); err != nil {
					return utils.WrapError("Decode Pucch_ConfigurationList_r16", err)
				}
				ie.Pucch_ConfigurationList_r16 = tmp_Pucch_ConfigurationList_r16.Setup
			}
			// decode Lbt_FailureRecoveryConfig_r16 optional
			if Lbt_FailureRecoveryConfig_r16Present {
				tmp_Lbt_FailureRecoveryConfig_r16 := utils.SetupRelease[*LBT_FailureRecoveryConfig_r16]{}
				if err = tmp_Lbt_FailureRecoveryConfig_r16.Decode(extReader); err != nil {
					return utils.WrapError("Decode Lbt_FailureRecoveryConfig_r16", err)
				}
				ie.Lbt_FailureRecoveryConfig_r16 = tmp_Lbt_FailureRecoveryConfig_r16.Setup
			}
			// decode ConfiguredGrantConfigToAddModList_r16 optional
			if ConfiguredGrantConfigToAddModList_r16Present {
				ie.ConfiguredGrantConfigToAddModList_r16 = new(ConfiguredGrantConfigToAddModList_r16)
				if err = ie.ConfiguredGrantConfigToAddModList_r16.Decode(extReader); err != nil {
					return utils.WrapError("Decode ConfiguredGrantConfigToAddModList_r16", err)
				}
			}
			// decode ConfiguredGrantConfigToReleaseList_r16 optional
			if ConfiguredGrantConfigToReleaseList_r16Present {
				ie.ConfiguredGrantConfigToReleaseList_r16 = new(ConfiguredGrantConfigToReleaseList_r16)
				if err = ie.ConfiguredGrantConfigToReleaseList_r16.Decode(extReader); err != nil {
					return utils.WrapError("Decode ConfiguredGrantConfigToReleaseList_r16", err)
				}
			}
			// decode ConfiguredGrantConfigType2DeactivationStateList_r16 optional
			if ConfiguredGrantConfigType2DeactivationStateList_r16Present {
				ie.ConfiguredGrantConfigType2DeactivationStateList_r16 = new(ConfiguredGrantConfigType2DeactivationStateList_r16)
				if err = ie.ConfiguredGrantConfigType2DeactivationStateList_r16.Decode(extReader); err != nil {
					return utils.WrapError("Decode ConfiguredGrantConfigType2DeactivationStateList_r16", err)
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

			Ul_TCI_StateList_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			Ul_powerControl_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			Pucch_ConfigurationListMulticast1_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			Pucch_ConfigurationListMulticast2_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			// decode Ul_TCI_StateList_r17 optional
			if Ul_TCI_StateList_r17Present {
				ie.Ul_TCI_StateList_r17 = new(BWP_UplinkDedicated_ul_TCI_StateList_r17)
				if err = ie.Ul_TCI_StateList_r17.Decode(extReader); err != nil {
					return utils.WrapError("Decode Ul_TCI_StateList_r17", err)
				}
			}
			// decode Ul_powerControl_r17 optional
			if Ul_powerControl_r17Present {
				ie.Ul_powerControl_r17 = new(Uplink_powerControlId_r17)
				if err = ie.Ul_powerControl_r17.Decode(extReader); err != nil {
					return utils.WrapError("Decode Ul_powerControl_r17", err)
				}
			}
			// decode Pucch_ConfigurationListMulticast1_r17 optional
			if Pucch_ConfigurationListMulticast1_r17Present {
				tmp_Pucch_ConfigurationListMulticast1_r17 := utils.SetupRelease[*PUCCH_ConfigurationList_r16]{}
				if err = tmp_Pucch_ConfigurationListMulticast1_r17.Decode(extReader); err != nil {
					return utils.WrapError("Decode Pucch_ConfigurationListMulticast1_r17", err)
				}
				ie.Pucch_ConfigurationListMulticast1_r17 = tmp_Pucch_ConfigurationListMulticast1_r17.Setup
			}
			// decode Pucch_ConfigurationListMulticast2_r17 optional
			if Pucch_ConfigurationListMulticast2_r17Present {
				tmp_Pucch_ConfigurationListMulticast2_r17 := utils.SetupRelease[*PUCCH_ConfigurationList_r16]{}
				if err = tmp_Pucch_ConfigurationListMulticast2_r17.Decode(extReader); err != nil {
					return utils.WrapError("Decode Pucch_ConfigurationListMulticast2_r17", err)
				}
				ie.Pucch_ConfigurationListMulticast2_r17 = tmp_Pucch_ConfigurationListMulticast2_r17.Setup
			}
		}
		// decode extension group 3
		if len(extBitmap) > 2 && extBitmap[2] {
			extBytes, err := r.ReadOpenType()
			if err != nil {
				return err
			}

			extReader := uper.NewReader(bytes.NewReader(extBytes))

			Pucch_ConfigMulticast1_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			Pucch_ConfigMulticast2_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			// decode Pucch_ConfigMulticast1_r17 optional
			if Pucch_ConfigMulticast1_r17Present {
				tmp_Pucch_ConfigMulticast1_r17 := utils.SetupRelease[*PUCCH_Config]{}
				if err = tmp_Pucch_ConfigMulticast1_r17.Decode(extReader); err != nil {
					return utils.WrapError("Decode Pucch_ConfigMulticast1_r17", err)
				}
				ie.Pucch_ConfigMulticast1_r17 = tmp_Pucch_ConfigMulticast1_r17.Setup
			}
			// decode Pucch_ConfigMulticast2_r17 optional
			if Pucch_ConfigMulticast2_r17Present {
				tmp_Pucch_ConfigMulticast2_r17 := utils.SetupRelease[*PUCCH_Config]{}
				if err = tmp_Pucch_ConfigMulticast2_r17.Decode(extReader); err != nil {
					return utils.WrapError("Decode Pucch_ConfigMulticast2_r17", err)
				}
				ie.Pucch_ConfigMulticast2_r17 = tmp_Pucch_ConfigMulticast2_r17.Setup
			}
		}
		// decode extension group 4
		if len(extBitmap) > 3 && extBitmap[3] {
			extBytes, err := r.ReadOpenType()
			if err != nil {
				return err
			}

			extReader := uper.NewReader(bytes.NewReader(extBytes))

			PathlossReferenceRSToAddModList_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			PathlossReferenceRSToReleaseList_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			// decode PathlossReferenceRSToAddModList_r17 optional
			if PathlossReferenceRSToAddModList_r17Present {
				tmp_PathlossReferenceRSToAddModList_r17 := utils.NewSequence[*PathlossReferenceRS_r17]([]*PathlossReferenceRS_r17{}, uper.Constraint{Lb: 1, Ub: maxNrofPathlossReferenceRSs_r17}, false)
				fn_PathlossReferenceRSToAddModList_r17 := func() *PathlossReferenceRS_r17 {
					return new(PathlossReferenceRS_r17)
				}
				if err = tmp_PathlossReferenceRSToAddModList_r17.Decode(extReader, fn_PathlossReferenceRSToAddModList_r17); err != nil {
					return utils.WrapError("Decode PathlossReferenceRSToAddModList_r17", err)
				}
				ie.PathlossReferenceRSToAddModList_r17 = []PathlossReferenceRS_r17{}
				for _, i := range tmp_PathlossReferenceRSToAddModList_r17.Value {
					ie.PathlossReferenceRSToAddModList_r17 = append(ie.PathlossReferenceRSToAddModList_r17, *i)
				}
			}
			// decode PathlossReferenceRSToReleaseList_r17 optional
			if PathlossReferenceRSToReleaseList_r17Present {
				tmp_PathlossReferenceRSToReleaseList_r17 := utils.NewSequence[*PathlossReferenceRS_Id_r17]([]*PathlossReferenceRS_Id_r17{}, uper.Constraint{Lb: 1, Ub: maxNrofPathlossReferenceRSs_r17}, false)
				fn_PathlossReferenceRSToReleaseList_r17 := func() *PathlossReferenceRS_Id_r17 {
					return new(PathlossReferenceRS_Id_r17)
				}
				if err = tmp_PathlossReferenceRSToReleaseList_r17.Decode(extReader, fn_PathlossReferenceRSToReleaseList_r17); err != nil {
					return utils.WrapError("Decode PathlossReferenceRSToReleaseList_r17", err)
				}
				ie.PathlossReferenceRSToReleaseList_r17 = []PathlossReferenceRS_Id_r17{}
				for _, i := range tmp_PathlossReferenceRSToReleaseList_r17.Value {
					ie.PathlossReferenceRSToReleaseList_r17 = append(ie.PathlossReferenceRSToReleaseList_r17, *i)
				}
			}
		}
	}
	return nil
}
