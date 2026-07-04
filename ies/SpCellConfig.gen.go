// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: SpCellConfig (line 5663).

var spCellConfigConstraints = per.SequenceConstraints{
	Extensible: true,
	RootComponents: []per.ComponentInfo{
		{Name: "servCellIndex", Optional: true},
		{Name: "reconfigurationWithSync", Optional: true},
		{Name: "rlf-TimersAndConstants", Optional: true},
		{Name: "rlmInSyncOutOfSyncThreshold", Optional: true},
		{Name: "spCellConfigDedicated", Optional: true},
	},
	ExtComponents: []per.ComponentInfo{
		{Name: "extension-addition-0", Optional: true},
	},
}

var spCellConfigRlfTimersAndConstantsConstraints = per.ChoiceConstraints{
	Extensible: false,
	RootAlternatives: []per.AlternativeInfo{
		{Name: "release"},
		{Name: "setup"},
	},
}

const (
	SpCellConfig_Rlf_TimersAndConstants_Release = 0
	SpCellConfig_Rlf_TimersAndConstants_Setup   = 1
)

const (
	SpCellConfig_RlmInSyncOutOfSyncThreshold_N1 = 0
)

var spCellConfigRlmInSyncOutOfSyncThresholdConstraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{SpCellConfig_RlmInSyncOutOfSyncThreshold_N1},
}

var spCellConfigExtDeactivatedSCGConfigR17Constraints = per.ChoiceConstraints{
	Extensible: false,
	RootAlternatives: []per.AlternativeInfo{
		{Name: "release"},
		{Name: "setup"},
	},
}

const (
	SpCellConfig_Ext_DeactivatedSCG_Config_r17_Release = 0
	SpCellConfig_Ext_DeactivatedSCG_Config_r17_Setup   = 1
)

const (
	SpCellConfig_Ext_LowMobilityEvaluationConnected_r17_S_SearchDeltaP_Connected_r17_DB3    = 0
	SpCellConfig_Ext_LowMobilityEvaluationConnected_r17_S_SearchDeltaP_Connected_r17_DB6    = 1
	SpCellConfig_Ext_LowMobilityEvaluationConnected_r17_S_SearchDeltaP_Connected_r17_DB9    = 2
	SpCellConfig_Ext_LowMobilityEvaluationConnected_r17_S_SearchDeltaP_Connected_r17_DB12   = 3
	SpCellConfig_Ext_LowMobilityEvaluationConnected_r17_S_SearchDeltaP_Connected_r17_DB15   = 4
	SpCellConfig_Ext_LowMobilityEvaluationConnected_r17_S_SearchDeltaP_Connected_r17_Spare3 = 5
	SpCellConfig_Ext_LowMobilityEvaluationConnected_r17_S_SearchDeltaP_Connected_r17_Spare2 = 6
	SpCellConfig_Ext_LowMobilityEvaluationConnected_r17_S_SearchDeltaP_Connected_r17_Spare1 = 7
)

var spCellConfigExtLowMobilityEvaluationConnectedR17SSearchDeltaPConnectedR17Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{SpCellConfig_Ext_LowMobilityEvaluationConnected_r17_S_SearchDeltaP_Connected_r17_DB3, SpCellConfig_Ext_LowMobilityEvaluationConnected_r17_S_SearchDeltaP_Connected_r17_DB6, SpCellConfig_Ext_LowMobilityEvaluationConnected_r17_S_SearchDeltaP_Connected_r17_DB9, SpCellConfig_Ext_LowMobilityEvaluationConnected_r17_S_SearchDeltaP_Connected_r17_DB12, SpCellConfig_Ext_LowMobilityEvaluationConnected_r17_S_SearchDeltaP_Connected_r17_DB15, SpCellConfig_Ext_LowMobilityEvaluationConnected_r17_S_SearchDeltaP_Connected_r17_Spare3, SpCellConfig_Ext_LowMobilityEvaluationConnected_r17_S_SearchDeltaP_Connected_r17_Spare2, SpCellConfig_Ext_LowMobilityEvaluationConnected_r17_S_SearchDeltaP_Connected_r17_Spare1},
}

const (
	SpCellConfig_Ext_LowMobilityEvaluationConnected_r17_T_SearchDeltaP_Connected_r17_S5     = 0
	SpCellConfig_Ext_LowMobilityEvaluationConnected_r17_T_SearchDeltaP_Connected_r17_S10    = 1
	SpCellConfig_Ext_LowMobilityEvaluationConnected_r17_T_SearchDeltaP_Connected_r17_S20    = 2
	SpCellConfig_Ext_LowMobilityEvaluationConnected_r17_T_SearchDeltaP_Connected_r17_S30    = 3
	SpCellConfig_Ext_LowMobilityEvaluationConnected_r17_T_SearchDeltaP_Connected_r17_S60    = 4
	SpCellConfig_Ext_LowMobilityEvaluationConnected_r17_T_SearchDeltaP_Connected_r17_S120   = 5
	SpCellConfig_Ext_LowMobilityEvaluationConnected_r17_T_SearchDeltaP_Connected_r17_S180   = 6
	SpCellConfig_Ext_LowMobilityEvaluationConnected_r17_T_SearchDeltaP_Connected_r17_S240   = 7
	SpCellConfig_Ext_LowMobilityEvaluationConnected_r17_T_SearchDeltaP_Connected_r17_S300   = 8
	SpCellConfig_Ext_LowMobilityEvaluationConnected_r17_T_SearchDeltaP_Connected_r17_Spare7 = 9
	SpCellConfig_Ext_LowMobilityEvaluationConnected_r17_T_SearchDeltaP_Connected_r17_Spare6 = 10
	SpCellConfig_Ext_LowMobilityEvaluationConnected_r17_T_SearchDeltaP_Connected_r17_Spare5 = 11
	SpCellConfig_Ext_LowMobilityEvaluationConnected_r17_T_SearchDeltaP_Connected_r17_Spare4 = 12
	SpCellConfig_Ext_LowMobilityEvaluationConnected_r17_T_SearchDeltaP_Connected_r17_Spare3 = 13
	SpCellConfig_Ext_LowMobilityEvaluationConnected_r17_T_SearchDeltaP_Connected_r17_Spare2 = 14
	SpCellConfig_Ext_LowMobilityEvaluationConnected_r17_T_SearchDeltaP_Connected_r17_Spare1 = 15
)

var spCellConfigExtLowMobilityEvaluationConnectedR17TSearchDeltaPConnectedR17Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{SpCellConfig_Ext_LowMobilityEvaluationConnected_r17_T_SearchDeltaP_Connected_r17_S5, SpCellConfig_Ext_LowMobilityEvaluationConnected_r17_T_SearchDeltaP_Connected_r17_S10, SpCellConfig_Ext_LowMobilityEvaluationConnected_r17_T_SearchDeltaP_Connected_r17_S20, SpCellConfig_Ext_LowMobilityEvaluationConnected_r17_T_SearchDeltaP_Connected_r17_S30, SpCellConfig_Ext_LowMobilityEvaluationConnected_r17_T_SearchDeltaP_Connected_r17_S60, SpCellConfig_Ext_LowMobilityEvaluationConnected_r17_T_SearchDeltaP_Connected_r17_S120, SpCellConfig_Ext_LowMobilityEvaluationConnected_r17_T_SearchDeltaP_Connected_r17_S180, SpCellConfig_Ext_LowMobilityEvaluationConnected_r17_T_SearchDeltaP_Connected_r17_S240, SpCellConfig_Ext_LowMobilityEvaluationConnected_r17_T_SearchDeltaP_Connected_r17_S300, SpCellConfig_Ext_LowMobilityEvaluationConnected_r17_T_SearchDeltaP_Connected_r17_Spare7, SpCellConfig_Ext_LowMobilityEvaluationConnected_r17_T_SearchDeltaP_Connected_r17_Spare6, SpCellConfig_Ext_LowMobilityEvaluationConnected_r17_T_SearchDeltaP_Connected_r17_Spare5, SpCellConfig_Ext_LowMobilityEvaluationConnected_r17_T_SearchDeltaP_Connected_r17_Spare4, SpCellConfig_Ext_LowMobilityEvaluationConnected_r17_T_SearchDeltaP_Connected_r17_Spare3, SpCellConfig_Ext_LowMobilityEvaluationConnected_r17_T_SearchDeltaP_Connected_r17_Spare2, SpCellConfig_Ext_LowMobilityEvaluationConnected_r17_T_SearchDeltaP_Connected_r17_Spare1},
}

type SpCellConfig struct {
	ServCellIndex           *ServCellIndex
	ReconfigurationWithSync *ReconfigurationWithSync
	Rlf_TimersAndConstants  *struct {
		Choice  int
		Release *struct{}
		Setup   *RLF_TimersAndConstants
	}
	RlmInSyncOutOfSyncThreshold        *int64
	SpCellConfigDedicated              *ServingCellConfig
	LowMobilityEvaluationConnected_r17 *struct {
		S_SearchDeltaP_Connected_r17 int64
		T_SearchDeltaP_Connected_r17 int64
	}
	GoodServingCellEvaluationRLM_r17 *GoodServingCellEvaluation_r17
	GoodServingCellEvaluationBFD_r17 *GoodServingCellEvaluation_r17
	DeactivatedSCG_Config_r17        *struct {
		Choice  int
		Release *struct{}
		Setup   *DeactivatedSCG_Config_r17
	}
}

func (ie *SpCellConfig) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(spCellConfigConstraints)

	// Extension-group presence (one var per [[...]] group).
	hasExtGroup0 := ie.LowMobilityEvaluationConnected_r17 != nil || ie.GoodServingCellEvaluationRLM_r17 != nil || ie.GoodServingCellEvaluationBFD_r17 != nil || ie.DeactivatedSCG_Config_r17 != nil
	hasExtensions := hasExtGroup0

	// 1. SEQUENCE extension bit.
	if err := seq.EncodeExtensionBit(hasExtensions); err != nil {
		return err
	}

	// 2. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.ServCellIndex != nil, ie.ReconfigurationWithSync != nil, ie.Rlf_TimersAndConstants != nil, ie.RlmInSyncOutOfSyncThreshold != nil, ie.SpCellConfigDedicated != nil}); err != nil {
		return err
	}

	// 3. servCellIndex: ref
	{
		if ie.ServCellIndex != nil {
			if err := ie.ServCellIndex.Encode(e); err != nil {
				return err
			}
		}
	}

	// 4. reconfigurationWithSync: ref
	{
		if ie.ReconfigurationWithSync != nil {
			if err := ie.ReconfigurationWithSync.Encode(e); err != nil {
				return err
			}
		}
	}

	// 5. rlf-TimersAndConstants: choice
	{
		if ie.Rlf_TimersAndConstants != nil {
			choiceEnc := e.NewChoiceEncoder(spCellConfigRlfTimersAndConstantsConstraints)
			if err := choiceEnc.EncodeChoice(int64((*ie.Rlf_TimersAndConstants).Choice), false, nil); err != nil {
				return err
			}
			switch (*ie.Rlf_TimersAndConstants).Choice {
			case SpCellConfig_Rlf_TimersAndConstants_Release:
				if err := e.EncodeNull(); err != nil {
					return err
				}
			case SpCellConfig_Rlf_TimersAndConstants_Setup:
				if err := (*ie.Rlf_TimersAndConstants).Setup.Encode(e); err != nil {
					return err
				}
			default:
				return &per.ConstraintViolationError{TypeName: "choice", Value: int64((*ie.Rlf_TimersAndConstants).Choice), Constraint: "index out of range"}
			}
		}
	}

	// 6. rlmInSyncOutOfSyncThreshold: enumerated
	{
		if ie.RlmInSyncOutOfSyncThreshold != nil {
			if err := e.EncodeEnumerated(*ie.RlmInSyncOutOfSyncThreshold, spCellConfigRlmInSyncOutOfSyncThresholdConstraints); err != nil {
				return err
			}
		}
	}

	// 7. spCellConfigDedicated: ref
	{
		if ie.SpCellConfigDedicated != nil {
			if err := ie.SpCellConfigDedicated.Encode(e); err != nil {
				return err
			}
		}
	}

	if hasExtensions {
		extPresence := []bool{hasExtGroup0}
		var extValues [][]byte

		// Extension group 0:
		if hasExtGroup0 {
			ex := per.NewEncoder(e.Variant())
			groupSeq := ex.NewSequenceEncoder(per.SequenceConstraints{
				Extensible: false,
				RootComponents: []per.ComponentInfo{
					{Name: "lowMobilityEvaluationConnected-r17", Optional: true},
					{Name: "goodServingCellEvaluationRLM-r17", Optional: true},
					{Name: "goodServingCellEvaluationBFD-r17", Optional: true},
					{Name: "deactivatedSCG-Config-r17", Optional: true},
				},
			})
			if err := groupSeq.EncodePreamble([]bool{ie.LowMobilityEvaluationConnected_r17 != nil, ie.GoodServingCellEvaluationRLM_r17 != nil, ie.GoodServingCellEvaluationBFD_r17 != nil, ie.DeactivatedSCG_Config_r17 != nil}); err != nil {
				return err
			}

			if ie.LowMobilityEvaluationConnected_r17 != nil {
				c := ie.LowMobilityEvaluationConnected_r17
				if err := ex.EncodeEnumerated(c.S_SearchDeltaP_Connected_r17, spCellConfigExtLowMobilityEvaluationConnectedR17SSearchDeltaPConnectedR17Constraints); err != nil {
					return err
				}
				if err := ex.EncodeEnumerated(c.T_SearchDeltaP_Connected_r17, spCellConfigExtLowMobilityEvaluationConnectedR17TSearchDeltaPConnectedR17Constraints); err != nil {
					return err
				}
			}

			if ie.GoodServingCellEvaluationRLM_r17 != nil {
				if err := ie.GoodServingCellEvaluationRLM_r17.Encode(ex); err != nil {
					return err
				}
			}

			if ie.GoodServingCellEvaluationBFD_r17 != nil {
				if err := ie.GoodServingCellEvaluationBFD_r17.Encode(ex); err != nil {
					return err
				}
			}

			if ie.DeactivatedSCG_Config_r17 != nil {
				choiceEnc := ex.NewChoiceEncoder(spCellConfigExtDeactivatedSCGConfigR17Constraints)
				if err := choiceEnc.EncodeChoice(int64((*ie.DeactivatedSCG_Config_r17).Choice), false, nil); err != nil {
					return err
				}
				switch (*ie.DeactivatedSCG_Config_r17).Choice {
				case SpCellConfig_Ext_DeactivatedSCG_Config_r17_Release:
					if err := ex.EncodeNull(); err != nil {
						return err
					}
				case SpCellConfig_Ext_DeactivatedSCG_Config_r17_Setup:
					if err := (*ie.DeactivatedSCG_Config_r17).Setup.Encode(ex); err != nil {
						return err
					}
				}
			}
			extValues = append(extValues, ex.Bytes())
		}

		if err := seq.EncodeExtensionAdditions(extPresence, extValues); err != nil {
			return err
		}
	}

	return nil
}

func (ie *SpCellConfig) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(spCellConfigConstraints)

	// 1. SEQUENCE extension bit.
	if err := seq.DecodeExtensionBit(); err != nil {
		return err
	}

	// 2. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 3. servCellIndex: ref
	{
		if seq.IsComponentPresent(0) {
			ie.ServCellIndex = new(ServCellIndex)
			if err := ie.ServCellIndex.Decode(d); err != nil {
				return err
			}
		}
	}

	// 4. reconfigurationWithSync: ref
	{
		if seq.IsComponentPresent(1) {
			ie.ReconfigurationWithSync = new(ReconfigurationWithSync)
			if err := ie.ReconfigurationWithSync.Decode(d); err != nil {
				return err
			}
		}
	}

	// 5. rlf-TimersAndConstants: choice
	{
		if seq.IsComponentPresent(2) {
			ie.Rlf_TimersAndConstants = &struct {
				Choice  int
				Release *struct{}
				Setup   *RLF_TimersAndConstants
			}{}
			choiceDec := d.NewChoiceDecoder(spCellConfigRlfTimersAndConstantsConstraints)
			index, isExt, _, err := choiceDec.DecodeChoice()
			if err != nil {
				return err
			}
			(*ie.Rlf_TimersAndConstants).Choice = int(index)
			if isExt {
				return &per.DecodeError{TypeName: "choice", Reason: "extension alternative not supported"}
			}
			switch index {
			case SpCellConfig_Rlf_TimersAndConstants_Release:
				(*ie.Rlf_TimersAndConstants).Release = &struct{}{}
				if err := d.DecodeNull(); err != nil {
					return err
				}
			case SpCellConfig_Rlf_TimersAndConstants_Setup:
				(*ie.Rlf_TimersAndConstants).Setup = new(RLF_TimersAndConstants)
				if err := (*ie.Rlf_TimersAndConstants).Setup.Decode(d); err != nil {
					return err
				}
			}
		}
	}

	// 6. rlmInSyncOutOfSyncThreshold: enumerated
	{
		if seq.IsComponentPresent(3) {
			idx, err := d.DecodeEnumerated(spCellConfigRlmInSyncOutOfSyncThresholdConstraints)
			if err != nil {
				return err
			}
			ie.RlmInSyncOutOfSyncThreshold = &idx
		}
	}

	// 7. spCellConfigDedicated: ref
	{
		if seq.IsComponentPresent(4) {
			ie.SpCellConfigDedicated = new(ServingCellConfig)
			if err := ie.SpCellConfigDedicated.Decode(d); err != nil {
				return err
			}
		}
	}

	// Extension additions as open types.
	extValues, err := seq.DecodeExtensionAdditions()
	if err != nil {
		return err
	}
	extValueIndex := 0

	// Extension group 0:
	if seq.IsExtensionPresent(0) {
		dx := per.NewDecoder(extValues[extValueIndex], d.Variant())
		extValueIndex++
		groupSeq := dx.NewSequenceDecoder(per.SequenceConstraints{
			Extensible: false,
			RootComponents: []per.ComponentInfo{
				{Name: "lowMobilityEvaluationConnected-r17", Optional: true},
				{Name: "goodServingCellEvaluationRLM-r17", Optional: true},
				{Name: "goodServingCellEvaluationBFD-r17", Optional: true},
				{Name: "deactivatedSCG-Config-r17", Optional: true},
			},
		})
		if err := groupSeq.DecodePreamble(); err != nil {
			return err
		}

		if groupSeq.IsComponentPresent(0) {
			ie.LowMobilityEvaluationConnected_r17 = &struct {
				S_SearchDeltaP_Connected_r17 int64
				T_SearchDeltaP_Connected_r17 int64
			}{}
			c := ie.LowMobilityEvaluationConnected_r17
			{
				v, err := dx.DecodeEnumerated(spCellConfigExtLowMobilityEvaluationConnectedR17SSearchDeltaPConnectedR17Constraints)
				if err != nil {
					return err
				}
				c.S_SearchDeltaP_Connected_r17 = v
			}
			{
				v, err := dx.DecodeEnumerated(spCellConfigExtLowMobilityEvaluationConnectedR17TSearchDeltaPConnectedR17Constraints)
				if err != nil {
					return err
				}
				c.T_SearchDeltaP_Connected_r17 = v
			}
		}

		if groupSeq.IsComponentPresent(1) {
			ie.GoodServingCellEvaluationRLM_r17 = new(GoodServingCellEvaluation_r17)
			if err := ie.GoodServingCellEvaluationRLM_r17.Decode(dx); err != nil {
				return err
			}
		}

		if groupSeq.IsComponentPresent(2) {
			ie.GoodServingCellEvaluationBFD_r17 = new(GoodServingCellEvaluation_r17)
			if err := ie.GoodServingCellEvaluationBFD_r17.Decode(dx); err != nil {
				return err
			}
		}

		if groupSeq.IsComponentPresent(3) {
			ie.DeactivatedSCG_Config_r17 = &struct {
				Choice  int
				Release *struct{}
				Setup   *DeactivatedSCG_Config_r17
			}{}
			choiceDec := dx.NewChoiceDecoder(spCellConfigExtDeactivatedSCGConfigR17Constraints)
			index, _, _, err := choiceDec.DecodeChoice()
			if err != nil {
				return err
			}
			(*ie.DeactivatedSCG_Config_r17).Choice = int(index)
			switch index {
			case SpCellConfig_Ext_DeactivatedSCG_Config_r17_Release:
				(*ie.DeactivatedSCG_Config_r17).Release = &struct{}{}
				if err := dx.DecodeNull(); err != nil {
					return err
				}
			case SpCellConfig_Ext_DeactivatedSCG_Config_r17_Setup:
				(*ie.DeactivatedSCG_Config_r17).Setup = new(DeactivatedSCG_Config_r17)
				if err := (*ie.DeactivatedSCG_Config_r17).Setup.Decode(dx); err != nil {
					return err
				}
			}
		}
	}

	return nil
}
