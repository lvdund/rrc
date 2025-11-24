package ies

import (
	"bytes"

	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type SpCellConfig struct {
	ServCellIndex                      *ServCellIndex                                   `optional`
	ReconfigurationWithSync            *ReconfigurationWithSync                         `optional`
	Rlf_TimersAndConstants             *RLF_TimersAndConstants                          `optional,setuprelease`
	RlmInSyncOutOfSyncThreshold        *SpCellConfig_rlmInSyncOutOfSyncThreshold        `optional`
	SpCellConfigDedicated              *ServingCellConfig                               `optional`
	LowMobilityEvaluationConnected_r17 *SpCellConfig_lowMobilityEvaluationConnected_r17 `optional,ext-1`
	GoodServingCellEvaluationRLM_r17   *GoodServingCellEvaluation_r17                   `optional,ext-1`
	GoodServingCellEvaluationBFD_r17   *GoodServingCellEvaluation_r17                   `optional,ext-1`
	DeactivatedSCG_Config_r17          *DeactivatedSCG_Config_r17                       `optional,ext-1,setuprelease`
}

func (ie *SpCellConfig) Encode(w *uper.UperWriter) error {
	var err error
	hasExtensions := ie.LowMobilityEvaluationConnected_r17 != nil || ie.GoodServingCellEvaluationRLM_r17 != nil || ie.GoodServingCellEvaluationBFD_r17 != nil || ie.DeactivatedSCG_Config_r17 != nil
	preambleBits := []bool{hasExtensions, ie.ServCellIndex != nil, ie.ReconfigurationWithSync != nil, ie.Rlf_TimersAndConstants != nil, ie.RlmInSyncOutOfSyncThreshold != nil, ie.SpCellConfigDedicated != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.ServCellIndex != nil {
		if err = ie.ServCellIndex.Encode(w); err != nil {
			return utils.WrapError("Encode ServCellIndex", err)
		}
	}
	if ie.ReconfigurationWithSync != nil {
		if err = ie.ReconfigurationWithSync.Encode(w); err != nil {
			return utils.WrapError("Encode ReconfigurationWithSync", err)
		}
	}
	if ie.Rlf_TimersAndConstants != nil {
		tmp_Rlf_TimersAndConstants := utils.SetupRelease[*RLF_TimersAndConstants]{
			Setup: ie.Rlf_TimersAndConstants,
		}
		if err = tmp_Rlf_TimersAndConstants.Encode(w); err != nil {
			return utils.WrapError("Encode Rlf_TimersAndConstants", err)
		}
	}
	if ie.RlmInSyncOutOfSyncThreshold != nil {
		if err = ie.RlmInSyncOutOfSyncThreshold.Encode(w); err != nil {
			return utils.WrapError("Encode RlmInSyncOutOfSyncThreshold", err)
		}
	}
	if ie.SpCellConfigDedicated != nil {
		if err = ie.SpCellConfigDedicated.Encode(w); err != nil {
			return utils.WrapError("Encode SpCellConfigDedicated", err)
		}
	}
	if hasExtensions {
		// Extension bitmap: 1 bits for 1 extension groups
		extBitmap := []bool{ie.LowMobilityEvaluationConnected_r17 != nil || ie.GoodServingCellEvaluationRLM_r17 != nil || ie.GoodServingCellEvaluationBFD_r17 != nil || ie.DeactivatedSCG_Config_r17 != nil}
		if err := w.WriteExtBitMap(extBitmap); err != nil {
			return utils.WrapError("WriteExtBitMap SpCellConfig", err)
		}

		// encode extension group 1
		if extBitmap[0] {
			extBuf := new(bytes.Buffer)
			extWriter := uper.NewWriter(extBuf)

			// Write preamble bits for optional fields in extension group 1
			optionals_ext_1 := []bool{ie.LowMobilityEvaluationConnected_r17 != nil, ie.GoodServingCellEvaluationRLM_r17 != nil, ie.GoodServingCellEvaluationBFD_r17 != nil, ie.DeactivatedSCG_Config_r17 != nil}
			for _, bit := range optionals_ext_1 {
				if err := extWriter.WriteBool(bit); err != nil {
					return err
				}
			}

			// encode LowMobilityEvaluationConnected_r17 optional
			if ie.LowMobilityEvaluationConnected_r17 != nil {
				if err = ie.LowMobilityEvaluationConnected_r17.Encode(extWriter); err != nil {
					return utils.WrapError("Encode LowMobilityEvaluationConnected_r17", err)
				}
			}
			// encode GoodServingCellEvaluationRLM_r17 optional
			if ie.GoodServingCellEvaluationRLM_r17 != nil {
				if err = ie.GoodServingCellEvaluationRLM_r17.Encode(extWriter); err != nil {
					return utils.WrapError("Encode GoodServingCellEvaluationRLM_r17", err)
				}
			}
			// encode GoodServingCellEvaluationBFD_r17 optional
			if ie.GoodServingCellEvaluationBFD_r17 != nil {
				if err = ie.GoodServingCellEvaluationBFD_r17.Encode(extWriter); err != nil {
					return utils.WrapError("Encode GoodServingCellEvaluationBFD_r17", err)
				}
			}
			// encode DeactivatedSCG_Config_r17 optional
			if ie.DeactivatedSCG_Config_r17 != nil {
				tmp_DeactivatedSCG_Config_r17 := utils.SetupRelease[*DeactivatedSCG_Config_r17]{
					Setup: ie.DeactivatedSCG_Config_r17,
				}
				if err = tmp_DeactivatedSCG_Config_r17.Encode(extWriter); err != nil {
					return utils.WrapError("Encode DeactivatedSCG_Config_r17", err)
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

func (ie *SpCellConfig) Decode(r *uper.UperReader) error {
	var err error
	var extensionBit bool
	if extensionBit, err = r.ReadBool(); err != nil {
		return err
	}
	var ServCellIndexPresent bool
	if ServCellIndexPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var ReconfigurationWithSyncPresent bool
	if ReconfigurationWithSyncPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var Rlf_TimersAndConstantsPresent bool
	if Rlf_TimersAndConstantsPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var RlmInSyncOutOfSyncThresholdPresent bool
	if RlmInSyncOutOfSyncThresholdPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var SpCellConfigDedicatedPresent bool
	if SpCellConfigDedicatedPresent, err = r.ReadBool(); err != nil {
		return err
	}
	if ServCellIndexPresent {
		ie.ServCellIndex = new(ServCellIndex)
		if err = ie.ServCellIndex.Decode(r); err != nil {
			return utils.WrapError("Decode ServCellIndex", err)
		}
	}
	if ReconfigurationWithSyncPresent {
		ie.ReconfigurationWithSync = new(ReconfigurationWithSync)
		if err = ie.ReconfigurationWithSync.Decode(r); err != nil {
			return utils.WrapError("Decode ReconfigurationWithSync", err)
		}
	}
	if Rlf_TimersAndConstantsPresent {
		tmp_Rlf_TimersAndConstants := utils.SetupRelease[*RLF_TimersAndConstants]{}
		if err = tmp_Rlf_TimersAndConstants.Decode(r); err != nil {
			return utils.WrapError("Decode Rlf_TimersAndConstants", err)
		}
		ie.Rlf_TimersAndConstants = tmp_Rlf_TimersAndConstants.Setup
	}
	if RlmInSyncOutOfSyncThresholdPresent {
		ie.RlmInSyncOutOfSyncThreshold = new(SpCellConfig_rlmInSyncOutOfSyncThreshold)
		if err = ie.RlmInSyncOutOfSyncThreshold.Decode(r); err != nil {
			return utils.WrapError("Decode RlmInSyncOutOfSyncThreshold", err)
		}
	}
	if SpCellConfigDedicatedPresent {
		ie.SpCellConfigDedicated = new(ServingCellConfig)
		if err = ie.SpCellConfigDedicated.Decode(r); err != nil {
			return utils.WrapError("Decode SpCellConfigDedicated", err)
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

			LowMobilityEvaluationConnected_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			GoodServingCellEvaluationRLM_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			GoodServingCellEvaluationBFD_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			DeactivatedSCG_Config_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			// decode LowMobilityEvaluationConnected_r17 optional
			if LowMobilityEvaluationConnected_r17Present {
				ie.LowMobilityEvaluationConnected_r17 = new(SpCellConfig_lowMobilityEvaluationConnected_r17)
				if err = ie.LowMobilityEvaluationConnected_r17.Decode(extReader); err != nil {
					return utils.WrapError("Decode LowMobilityEvaluationConnected_r17", err)
				}
			}
			// decode GoodServingCellEvaluationRLM_r17 optional
			if GoodServingCellEvaluationRLM_r17Present {
				ie.GoodServingCellEvaluationRLM_r17 = new(GoodServingCellEvaluation_r17)
				if err = ie.GoodServingCellEvaluationRLM_r17.Decode(extReader); err != nil {
					return utils.WrapError("Decode GoodServingCellEvaluationRLM_r17", err)
				}
			}
			// decode GoodServingCellEvaluationBFD_r17 optional
			if GoodServingCellEvaluationBFD_r17Present {
				ie.GoodServingCellEvaluationBFD_r17 = new(GoodServingCellEvaluation_r17)
				if err = ie.GoodServingCellEvaluationBFD_r17.Decode(extReader); err != nil {
					return utils.WrapError("Decode GoodServingCellEvaluationBFD_r17", err)
				}
			}
			// decode DeactivatedSCG_Config_r17 optional
			if DeactivatedSCG_Config_r17Present {
				tmp_DeactivatedSCG_Config_r17 := utils.SetupRelease[*DeactivatedSCG_Config_r17]{}
				if err = tmp_DeactivatedSCG_Config_r17.Decode(extReader); err != nil {
					return utils.WrapError("Decode DeactivatedSCG_Config_r17", err)
				}
				ie.DeactivatedSCG_Config_r17 = tmp_DeactivatedSCG_Config_r17.Setup
			}
		}
	}
	return nil
}
