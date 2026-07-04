// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/rrc/common"
)

// Source: UplinkConfig (line 14728).

var uplinkConfigConstraints = per.SequenceConstraints{
	Extensible: true,
	RootComponents: []per.ComponentInfo{
		{Name: "initialUplinkBWP", Optional: true},
		{Name: "uplinkBWP-ToReleaseList", Optional: true},
		{Name: "uplinkBWP-ToAddModList", Optional: true},
		{Name: "firstActiveUplinkBWP-Id", Optional: true},
		{Name: "pusch-ServingCellConfig", Optional: true},
		{Name: "carrierSwitching", Optional: true},
	},
	ExtComponents: []per.ComponentInfo{
		{Name: "extension-addition-0", Optional: true},
		{Name: "extension-addition-1", Optional: true},
		{Name: "extension-addition-2", Optional: true},
	},
}

var uplinkConfigUplinkBWPToReleaseListConstraints = per.SizeRange(1, common.MaxNrofBWPs)

var uplinkConfigUplinkBWPToAddModListConstraints = per.SizeRange(1, common.MaxNrofBWPs)

var uplinkConfigPuschServingCellConfigConstraints = per.ChoiceConstraints{
	Extensible: false,
	RootAlternatives: []per.AlternativeInfo{
		{Name: "release"},
		{Name: "setup"},
	},
}

const (
	UplinkConfig_Pusch_ServingCellConfig_Release = 0
	UplinkConfig_Pusch_ServingCellConfig_Setup   = 1
)

var uplinkConfigCarrierSwitchingConstraints = per.ChoiceConstraints{
	Extensible: false,
	RootAlternatives: []per.AlternativeInfo{
		{Name: "release"},
		{Name: "setup"},
	},
}

const (
	UplinkConfig_CarrierSwitching_Release = 0
	UplinkConfig_CarrierSwitching_Setup   = 1
)

var uplinkConfigExtUplinkChannelBWPerSCSListConstraints = per.SizeRange(1, common.MaxSCSs)

const (
	UplinkConfig_Ext_EnablePL_RS_UpdateForPUSCH_SRS_r16_Enabled = 0
)

var uplinkConfigExtEnablePLRSUpdateForPUSCHSRSR16Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{UplinkConfig_Ext_EnablePL_RS_UpdateForPUSCH_SRS_r16_Enabled},
}

const (
	UplinkConfig_Ext_EnableDefaultBeamPL_ForPUSCH0_0_r16_Enabled = 0
)

var uplinkConfigExtEnableDefaultBeamPLForPUSCH00R16Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{UplinkConfig_Ext_EnableDefaultBeamPL_ForPUSCH0_0_r16_Enabled},
}

const (
	UplinkConfig_Ext_EnableDefaultBeamPL_ForPUCCH_r16_Enabled = 0
)

var uplinkConfigExtEnableDefaultBeamPLForPUCCHR16Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{UplinkConfig_Ext_EnableDefaultBeamPL_ForPUCCH_r16_Enabled},
}

const (
	UplinkConfig_Ext_EnableDefaultBeamPL_ForSRS_r16_Enabled = 0
)

var uplinkConfigExtEnableDefaultBeamPLForSRSR16Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{UplinkConfig_Ext_EnableDefaultBeamPL_ForSRS_r16_Enabled},
}

var uplinkConfigExtUplinkTxSwitchingR16Constraints = per.ChoiceConstraints{
	Extensible: false,
	RootAlternatives: []per.AlternativeInfo{
		{Name: "release"},
		{Name: "setup"},
	},
}

const (
	UplinkConfig_Ext_UplinkTxSwitching_r16_Release = 0
	UplinkConfig_Ext_UplinkTxSwitching_r16_Setup   = 1
)

const (
	UplinkConfig_Ext_Mpr_PowerBoost_FR2_r16_True = 0
)

var uplinkConfigExtMprPowerBoostFR2R16Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{UplinkConfig_Ext_Mpr_PowerBoost_FR2_r16_True},
}

var uplinkConfigExtSrsPosTxHoppingR18Constraints = per.ChoiceConstraints{
	Extensible: false,
	RootAlternatives: []per.AlternativeInfo{
		{Name: "release"},
		{Name: "setup"},
	},
}

const (
	UplinkConfig_Ext_Srs_PosTx_Hopping_r18_Release = 0
	UplinkConfig_Ext_Srs_PosTx_Hopping_r18_Setup   = 1
)

const (
	UplinkConfig_Ext_EnablePL_RS_UpdateForType1CG_PUSCH_r18_Enabled = 0
)

var uplinkConfigExtEnablePLRSUpdateForType1CGPUSCHR18Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{UplinkConfig_Ext_EnablePL_RS_UpdateForType1CG_PUSCH_r18_Enabled},
}

type UplinkConfig struct {
	InitialUplinkBWP        *BWP_UplinkDedicated
	UplinkBWP_ToReleaseList []BWP_Id
	UplinkBWP_ToAddModList  []BWP_Uplink
	FirstActiveUplinkBWP_Id *BWP_Id
	Pusch_ServingCellConfig *struct {
		Choice  int
		Release *struct{}
		Setup   *PUSCH_ServingCellConfig
	}
	CarrierSwitching *struct {
		Choice  int
		Release *struct{}
		Setup   *SRS_CarrierSwitching
	}
	PowerBoostPi2BPSK                   *bool
	UplinkChannelBW_PerSCS_List         []SCS_SpecificCarrier
	EnablePL_RS_UpdateForPUSCH_SRS_r16  *int64
	EnableDefaultBeamPL_ForPUSCH0_0_r16 *int64
	EnableDefaultBeamPL_ForPUCCH_r16    *int64
	EnableDefaultBeamPL_ForSRS_r16      *int64
	UplinkTxSwitching_r16               *struct {
		Choice  int
		Release *struct{}
		Setup   *UplinkTxSwitching_r16
	}
	Mpr_PowerBoost_FR2_r16 *int64
	Srs_PosTx_Hopping_r18  *struct {
		Choice  int
		Release *struct{}
		Setup   *SRS_PosTx_Hopping_r18
	}
	EnablePL_RS_UpdateForType1CG_PUSCH_r18 *int64
	PowerBoostPi2BPSK_r18                  *bool
	PowerBoostQPSK_r18                     *bool
}

func (ie *UplinkConfig) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(uplinkConfigConstraints)

	// Extension-group presence (one var per [[...]] group).
	hasExtGroup0 := ie.PowerBoostPi2BPSK != nil || ie.UplinkChannelBW_PerSCS_List != nil
	hasExtGroup1 := ie.EnablePL_RS_UpdateForPUSCH_SRS_r16 != nil || ie.EnableDefaultBeamPL_ForPUSCH0_0_r16 != nil || ie.EnableDefaultBeamPL_ForPUCCH_r16 != nil || ie.EnableDefaultBeamPL_ForSRS_r16 != nil || ie.UplinkTxSwitching_r16 != nil || ie.Mpr_PowerBoost_FR2_r16 != nil
	hasExtGroup2 := ie.Srs_PosTx_Hopping_r18 != nil || ie.EnablePL_RS_UpdateForType1CG_PUSCH_r18 != nil || ie.PowerBoostPi2BPSK_r18 != nil || ie.PowerBoostQPSK_r18 != nil
	hasExtensions := hasExtGroup0 || hasExtGroup1 || hasExtGroup2

	// 1. SEQUENCE extension bit.
	if err := seq.EncodeExtensionBit(hasExtensions); err != nil {
		return err
	}

	// 2. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.InitialUplinkBWP != nil, ie.UplinkBWP_ToReleaseList != nil, ie.UplinkBWP_ToAddModList != nil, ie.FirstActiveUplinkBWP_Id != nil, ie.Pusch_ServingCellConfig != nil, ie.CarrierSwitching != nil}); err != nil {
		return err
	}

	// 3. initialUplinkBWP: ref
	{
		if ie.InitialUplinkBWP != nil {
			if err := ie.InitialUplinkBWP.Encode(e); err != nil {
				return err
			}
		}
	}

	// 4. uplinkBWP-ToReleaseList: sequence-of
	{
		if ie.UplinkBWP_ToReleaseList != nil {
			seqOf := e.NewSequenceOfEncoder(uplinkConfigUplinkBWPToReleaseListConstraints)
			if err := seqOf.EncodeLength(int64(len(ie.UplinkBWP_ToReleaseList))); err != nil {
				return err
			}
			for i := range ie.UplinkBWP_ToReleaseList {
				if err := ie.UplinkBWP_ToReleaseList[i].Encode(e); err != nil {
					return err
				}
			}
		}
	}

	// 5. uplinkBWP-ToAddModList: sequence-of
	{
		if ie.UplinkBWP_ToAddModList != nil {
			seqOf := e.NewSequenceOfEncoder(uplinkConfigUplinkBWPToAddModListConstraints)
			if err := seqOf.EncodeLength(int64(len(ie.UplinkBWP_ToAddModList))); err != nil {
				return err
			}
			for i := range ie.UplinkBWP_ToAddModList {
				if err := ie.UplinkBWP_ToAddModList[i].Encode(e); err != nil {
					return err
				}
			}
		}
	}

	// 6. firstActiveUplinkBWP-Id: ref
	{
		if ie.FirstActiveUplinkBWP_Id != nil {
			if err := ie.FirstActiveUplinkBWP_Id.Encode(e); err != nil {
				return err
			}
		}
	}

	// 7. pusch-ServingCellConfig: choice
	{
		if ie.Pusch_ServingCellConfig != nil {
			choiceEnc := e.NewChoiceEncoder(uplinkConfigPuschServingCellConfigConstraints)
			if err := choiceEnc.EncodeChoice(int64((*ie.Pusch_ServingCellConfig).Choice), false, nil); err != nil {
				return err
			}
			switch (*ie.Pusch_ServingCellConfig).Choice {
			case UplinkConfig_Pusch_ServingCellConfig_Release:
				if err := e.EncodeNull(); err != nil {
					return err
				}
			case UplinkConfig_Pusch_ServingCellConfig_Setup:
				if err := (*ie.Pusch_ServingCellConfig).Setup.Encode(e); err != nil {
					return err
				}
			default:
				return &per.ConstraintViolationError{TypeName: "choice", Value: int64((*ie.Pusch_ServingCellConfig).Choice), Constraint: "index out of range"}
			}
		}
	}

	// 8. carrierSwitching: choice
	{
		if ie.CarrierSwitching != nil {
			choiceEnc := e.NewChoiceEncoder(uplinkConfigCarrierSwitchingConstraints)
			if err := choiceEnc.EncodeChoice(int64((*ie.CarrierSwitching).Choice), false, nil); err != nil {
				return err
			}
			switch (*ie.CarrierSwitching).Choice {
			case UplinkConfig_CarrierSwitching_Release:
				if err := e.EncodeNull(); err != nil {
					return err
				}
			case UplinkConfig_CarrierSwitching_Setup:
				if err := (*ie.CarrierSwitching).Setup.Encode(e); err != nil {
					return err
				}
			default:
				return &per.ConstraintViolationError{TypeName: "choice", Value: int64((*ie.CarrierSwitching).Choice), Constraint: "index out of range"}
			}
		}
	}

	if hasExtensions {
		extPresence := []bool{hasExtGroup0, hasExtGroup1, hasExtGroup2}
		var extValues [][]byte

		// Extension group 0:
		if hasExtGroup0 {
			ex := per.NewEncoder(e.Variant())
			groupSeq := ex.NewSequenceEncoder(per.SequenceConstraints{
				Extensible: false,
				RootComponents: []per.ComponentInfo{
					{Name: "powerBoostPi2BPSK", Optional: true},
					{Name: "uplinkChannelBW-PerSCS-List", Optional: true},
				},
			})
			if err := groupSeq.EncodePreamble([]bool{ie.PowerBoostPi2BPSK != nil, ie.UplinkChannelBW_PerSCS_List != nil}); err != nil {
				return err
			}

			if ie.PowerBoostPi2BPSK != nil {
				if err := ex.EncodeBoolean(*ie.PowerBoostPi2BPSK); err != nil {
					return err
				}
			}

			if ie.UplinkChannelBW_PerSCS_List != nil {
				seqOf := ex.NewSequenceOfEncoder(uplinkConfigExtUplinkChannelBWPerSCSListConstraints)
				if err := seqOf.EncodeLength(int64(len(ie.UplinkChannelBW_PerSCS_List))); err != nil {
					return err
				}
				for i := range ie.UplinkChannelBW_PerSCS_List {
					if err := ie.UplinkChannelBW_PerSCS_List[i].Encode(ex); err != nil {
						return err
					}
				}
			}
			extValues = append(extValues, ex.Bytes())
		}

		// Extension group 1:
		if hasExtGroup1 {
			ex := per.NewEncoder(e.Variant())
			groupSeq := ex.NewSequenceEncoder(per.SequenceConstraints{
				Extensible: false,
				RootComponents: []per.ComponentInfo{
					{Name: "enablePL-RS-UpdateForPUSCH-SRS-r16", Optional: true},
					{Name: "enableDefaultBeamPL-ForPUSCH0-0-r16", Optional: true},
					{Name: "enableDefaultBeamPL-ForPUCCH-r16", Optional: true},
					{Name: "enableDefaultBeamPL-ForSRS-r16", Optional: true},
					{Name: "uplinkTxSwitching-r16", Optional: true},
					{Name: "mpr-PowerBoost-FR2-r16", Optional: true},
				},
			})
			if err := groupSeq.EncodePreamble([]bool{ie.EnablePL_RS_UpdateForPUSCH_SRS_r16 != nil, ie.EnableDefaultBeamPL_ForPUSCH0_0_r16 != nil, ie.EnableDefaultBeamPL_ForPUCCH_r16 != nil, ie.EnableDefaultBeamPL_ForSRS_r16 != nil, ie.UplinkTxSwitching_r16 != nil, ie.Mpr_PowerBoost_FR2_r16 != nil}); err != nil {
				return err
			}

			if ie.EnablePL_RS_UpdateForPUSCH_SRS_r16 != nil {
				if err := ex.EncodeEnumerated(*ie.EnablePL_RS_UpdateForPUSCH_SRS_r16, uplinkConfigExtEnablePLRSUpdateForPUSCHSRSR16Constraints); err != nil {
					return err
				}
			}

			if ie.EnableDefaultBeamPL_ForPUSCH0_0_r16 != nil {
				if err := ex.EncodeEnumerated(*ie.EnableDefaultBeamPL_ForPUSCH0_0_r16, uplinkConfigExtEnableDefaultBeamPLForPUSCH00R16Constraints); err != nil {
					return err
				}
			}

			if ie.EnableDefaultBeamPL_ForPUCCH_r16 != nil {
				if err := ex.EncodeEnumerated(*ie.EnableDefaultBeamPL_ForPUCCH_r16, uplinkConfigExtEnableDefaultBeamPLForPUCCHR16Constraints); err != nil {
					return err
				}
			}

			if ie.EnableDefaultBeamPL_ForSRS_r16 != nil {
				if err := ex.EncodeEnumerated(*ie.EnableDefaultBeamPL_ForSRS_r16, uplinkConfigExtEnableDefaultBeamPLForSRSR16Constraints); err != nil {
					return err
				}
			}

			if ie.UplinkTxSwitching_r16 != nil {
				choiceEnc := ex.NewChoiceEncoder(uplinkConfigExtUplinkTxSwitchingR16Constraints)
				if err := choiceEnc.EncodeChoice(int64((*ie.UplinkTxSwitching_r16).Choice), false, nil); err != nil {
					return err
				}
				switch (*ie.UplinkTxSwitching_r16).Choice {
				case UplinkConfig_Ext_UplinkTxSwitching_r16_Release:
					if err := ex.EncodeNull(); err != nil {
						return err
					}
				case UplinkConfig_Ext_UplinkTxSwitching_r16_Setup:
					if err := (*ie.UplinkTxSwitching_r16).Setup.Encode(ex); err != nil {
						return err
					}
				}
			}

			if ie.Mpr_PowerBoost_FR2_r16 != nil {
				if err := ex.EncodeEnumerated(*ie.Mpr_PowerBoost_FR2_r16, uplinkConfigExtMprPowerBoostFR2R16Constraints); err != nil {
					return err
				}
			}
			extValues = append(extValues, ex.Bytes())
		}

		// Extension group 2:
		if hasExtGroup2 {
			ex := per.NewEncoder(e.Variant())
			groupSeq := ex.NewSequenceEncoder(per.SequenceConstraints{
				Extensible: false,
				RootComponents: []per.ComponentInfo{
					{Name: "srs-PosTx-Hopping-r18", Optional: true},
					{Name: "enablePL-RS-UpdateForType1CG-PUSCH-r18", Optional: true},
					{Name: "powerBoostPi2BPSK-r18", Optional: true},
					{Name: "powerBoostQPSK-r18", Optional: true},
				},
			})
			if err := groupSeq.EncodePreamble([]bool{ie.Srs_PosTx_Hopping_r18 != nil, ie.EnablePL_RS_UpdateForType1CG_PUSCH_r18 != nil, ie.PowerBoostPi2BPSK_r18 != nil, ie.PowerBoostQPSK_r18 != nil}); err != nil {
				return err
			}

			if ie.Srs_PosTx_Hopping_r18 != nil {
				choiceEnc := ex.NewChoiceEncoder(uplinkConfigExtSrsPosTxHoppingR18Constraints)
				if err := choiceEnc.EncodeChoice(int64((*ie.Srs_PosTx_Hopping_r18).Choice), false, nil); err != nil {
					return err
				}
				switch (*ie.Srs_PosTx_Hopping_r18).Choice {
				case UplinkConfig_Ext_Srs_PosTx_Hopping_r18_Release:
					if err := ex.EncodeNull(); err != nil {
						return err
					}
				case UplinkConfig_Ext_Srs_PosTx_Hopping_r18_Setup:
					if err := (*ie.Srs_PosTx_Hopping_r18).Setup.Encode(ex); err != nil {
						return err
					}
				}
			}

			if ie.EnablePL_RS_UpdateForType1CG_PUSCH_r18 != nil {
				if err := ex.EncodeEnumerated(*ie.EnablePL_RS_UpdateForType1CG_PUSCH_r18, uplinkConfigExtEnablePLRSUpdateForType1CGPUSCHR18Constraints); err != nil {
					return err
				}
			}

			if ie.PowerBoostPi2BPSK_r18 != nil {
				if err := ex.EncodeBoolean(*ie.PowerBoostPi2BPSK_r18); err != nil {
					return err
				}
			}

			if ie.PowerBoostQPSK_r18 != nil {
				if err := ex.EncodeBoolean(*ie.PowerBoostQPSK_r18); err != nil {
					return err
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

func (ie *UplinkConfig) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(uplinkConfigConstraints)

	// 1. SEQUENCE extension bit.
	if err := seq.DecodeExtensionBit(); err != nil {
		return err
	}

	// 2. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 3. initialUplinkBWP: ref
	{
		if seq.IsComponentPresent(0) {
			ie.InitialUplinkBWP = new(BWP_UplinkDedicated)
			if err := ie.InitialUplinkBWP.Decode(d); err != nil {
				return err
			}
		}
	}

	// 4. uplinkBWP-ToReleaseList: sequence-of
	{
		if seq.IsComponentPresent(1) {
			seqOf := d.NewSequenceOfDecoder(uplinkConfigUplinkBWPToReleaseListConstraints)
			n, err := seqOf.DecodeLength()
			if err != nil {
				return err
			}
			ie.UplinkBWP_ToReleaseList = make([]BWP_Id, n)
			for i := int64(0); i < n; i++ {
				if err := ie.UplinkBWP_ToReleaseList[i].Decode(d); err != nil {
					return err
				}
			}
		}
	}

	// 5. uplinkBWP-ToAddModList: sequence-of
	{
		if seq.IsComponentPresent(2) {
			seqOf := d.NewSequenceOfDecoder(uplinkConfigUplinkBWPToAddModListConstraints)
			n, err := seqOf.DecodeLength()
			if err != nil {
				return err
			}
			ie.UplinkBWP_ToAddModList = make([]BWP_Uplink, n)
			for i := int64(0); i < n; i++ {
				if err := ie.UplinkBWP_ToAddModList[i].Decode(d); err != nil {
					return err
				}
			}
		}
	}

	// 6. firstActiveUplinkBWP-Id: ref
	{
		if seq.IsComponentPresent(3) {
			ie.FirstActiveUplinkBWP_Id = new(BWP_Id)
			if err := ie.FirstActiveUplinkBWP_Id.Decode(d); err != nil {
				return err
			}
		}
	}

	// 7. pusch-ServingCellConfig: choice
	{
		if seq.IsComponentPresent(4) {
			ie.Pusch_ServingCellConfig = &struct {
				Choice  int
				Release *struct{}
				Setup   *PUSCH_ServingCellConfig
			}{}
			choiceDec := d.NewChoiceDecoder(uplinkConfigPuschServingCellConfigConstraints)
			index, isExt, _, err := choiceDec.DecodeChoice()
			if err != nil {
				return err
			}
			(*ie.Pusch_ServingCellConfig).Choice = int(index)
			if isExt {
				return &per.DecodeError{TypeName: "choice", Reason: "extension alternative not supported"}
			}
			switch index {
			case UplinkConfig_Pusch_ServingCellConfig_Release:
				(*ie.Pusch_ServingCellConfig).Release = &struct{}{}
				if err := d.DecodeNull(); err != nil {
					return err
				}
			case UplinkConfig_Pusch_ServingCellConfig_Setup:
				(*ie.Pusch_ServingCellConfig).Setup = new(PUSCH_ServingCellConfig)
				if err := (*ie.Pusch_ServingCellConfig).Setup.Decode(d); err != nil {
					return err
				}
			}
		}
	}

	// 8. carrierSwitching: choice
	{
		if seq.IsComponentPresent(5) {
			ie.CarrierSwitching = &struct {
				Choice  int
				Release *struct{}
				Setup   *SRS_CarrierSwitching
			}{}
			choiceDec := d.NewChoiceDecoder(uplinkConfigCarrierSwitchingConstraints)
			index, isExt, _, err := choiceDec.DecodeChoice()
			if err != nil {
				return err
			}
			(*ie.CarrierSwitching).Choice = int(index)
			if isExt {
				return &per.DecodeError{TypeName: "choice", Reason: "extension alternative not supported"}
			}
			switch index {
			case UplinkConfig_CarrierSwitching_Release:
				(*ie.CarrierSwitching).Release = &struct{}{}
				if err := d.DecodeNull(); err != nil {
					return err
				}
			case UplinkConfig_CarrierSwitching_Setup:
				(*ie.CarrierSwitching).Setup = new(SRS_CarrierSwitching)
				if err := (*ie.CarrierSwitching).Setup.Decode(d); err != nil {
					return err
				}
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
				{Name: "powerBoostPi2BPSK", Optional: true},
				{Name: "uplinkChannelBW-PerSCS-List", Optional: true},
			},
		})
		if err := groupSeq.DecodePreamble(); err != nil {
			return err
		}

		if groupSeq.IsComponentPresent(0) {
			v, err := dx.DecodeBoolean()
			if err != nil {
				return err
			}
			ie.PowerBoostPi2BPSK = &v
		}

		if groupSeq.IsComponentPresent(1) {
			seqOf := dx.NewSequenceOfDecoder(uplinkConfigExtUplinkChannelBWPerSCSListConstraints)
			n, err := seqOf.DecodeLength()
			if err != nil {
				return err
			}
			ie.UplinkChannelBW_PerSCS_List = make([]SCS_SpecificCarrier, n)
			for i := int64(0); i < n; i++ {
				if err := ie.UplinkChannelBW_PerSCS_List[i].Decode(dx); err != nil {
					return err
				}
			}
		}
	}

	// Extension group 1:
	if seq.IsExtensionPresent(1) {
		dx := per.NewDecoder(extValues[extValueIndex], d.Variant())
		extValueIndex++
		groupSeq := dx.NewSequenceDecoder(per.SequenceConstraints{
			Extensible: false,
			RootComponents: []per.ComponentInfo{
				{Name: "enablePL-RS-UpdateForPUSCH-SRS-r16", Optional: true},
				{Name: "enableDefaultBeamPL-ForPUSCH0-0-r16", Optional: true},
				{Name: "enableDefaultBeamPL-ForPUCCH-r16", Optional: true},
				{Name: "enableDefaultBeamPL-ForSRS-r16", Optional: true},
				{Name: "uplinkTxSwitching-r16", Optional: true},
				{Name: "mpr-PowerBoost-FR2-r16", Optional: true},
			},
		})
		if err := groupSeq.DecodePreamble(); err != nil {
			return err
		}

		if groupSeq.IsComponentPresent(0) {
			v, err := dx.DecodeEnumerated(uplinkConfigExtEnablePLRSUpdateForPUSCHSRSR16Constraints)
			if err != nil {
				return err
			}
			ie.EnablePL_RS_UpdateForPUSCH_SRS_r16 = &v
		}

		if groupSeq.IsComponentPresent(1) {
			v, err := dx.DecodeEnumerated(uplinkConfigExtEnableDefaultBeamPLForPUSCH00R16Constraints)
			if err != nil {
				return err
			}
			ie.EnableDefaultBeamPL_ForPUSCH0_0_r16 = &v
		}

		if groupSeq.IsComponentPresent(2) {
			v, err := dx.DecodeEnumerated(uplinkConfigExtEnableDefaultBeamPLForPUCCHR16Constraints)
			if err != nil {
				return err
			}
			ie.EnableDefaultBeamPL_ForPUCCH_r16 = &v
		}

		if groupSeq.IsComponentPresent(3) {
			v, err := dx.DecodeEnumerated(uplinkConfigExtEnableDefaultBeamPLForSRSR16Constraints)
			if err != nil {
				return err
			}
			ie.EnableDefaultBeamPL_ForSRS_r16 = &v
		}

		if groupSeq.IsComponentPresent(4) {
			ie.UplinkTxSwitching_r16 = &struct {
				Choice  int
				Release *struct{}
				Setup   *UplinkTxSwitching_r16
			}{}
			choiceDec := dx.NewChoiceDecoder(uplinkConfigExtUplinkTxSwitchingR16Constraints)
			index, _, _, err := choiceDec.DecodeChoice()
			if err != nil {
				return err
			}
			(*ie.UplinkTxSwitching_r16).Choice = int(index)
			switch index {
			case UplinkConfig_Ext_UplinkTxSwitching_r16_Release:
				(*ie.UplinkTxSwitching_r16).Release = &struct{}{}
				if err := dx.DecodeNull(); err != nil {
					return err
				}
			case UplinkConfig_Ext_UplinkTxSwitching_r16_Setup:
				(*ie.UplinkTxSwitching_r16).Setup = new(UplinkTxSwitching_r16)
				if err := (*ie.UplinkTxSwitching_r16).Setup.Decode(dx); err != nil {
					return err
				}
			}
		}

		if groupSeq.IsComponentPresent(5) {
			v, err := dx.DecodeEnumerated(uplinkConfigExtMprPowerBoostFR2R16Constraints)
			if err != nil {
				return err
			}
			ie.Mpr_PowerBoost_FR2_r16 = &v
		}
	}

	// Extension group 2:
	if seq.IsExtensionPresent(2) {
		dx := per.NewDecoder(extValues[extValueIndex], d.Variant())
		extValueIndex++
		groupSeq := dx.NewSequenceDecoder(per.SequenceConstraints{
			Extensible: false,
			RootComponents: []per.ComponentInfo{
				{Name: "srs-PosTx-Hopping-r18", Optional: true},
				{Name: "enablePL-RS-UpdateForType1CG-PUSCH-r18", Optional: true},
				{Name: "powerBoostPi2BPSK-r18", Optional: true},
				{Name: "powerBoostQPSK-r18", Optional: true},
			},
		})
		if err := groupSeq.DecodePreamble(); err != nil {
			return err
		}

		if groupSeq.IsComponentPresent(0) {
			ie.Srs_PosTx_Hopping_r18 = &struct {
				Choice  int
				Release *struct{}
				Setup   *SRS_PosTx_Hopping_r18
			}{}
			choiceDec := dx.NewChoiceDecoder(uplinkConfigExtSrsPosTxHoppingR18Constraints)
			index, _, _, err := choiceDec.DecodeChoice()
			if err != nil {
				return err
			}
			(*ie.Srs_PosTx_Hopping_r18).Choice = int(index)
			switch index {
			case UplinkConfig_Ext_Srs_PosTx_Hopping_r18_Release:
				(*ie.Srs_PosTx_Hopping_r18).Release = &struct{}{}
				if err := dx.DecodeNull(); err != nil {
					return err
				}
			case UplinkConfig_Ext_Srs_PosTx_Hopping_r18_Setup:
				(*ie.Srs_PosTx_Hopping_r18).Setup = new(SRS_PosTx_Hopping_r18)
				if err := (*ie.Srs_PosTx_Hopping_r18).Setup.Decode(dx); err != nil {
					return err
				}
			}
		}

		if groupSeq.IsComponentPresent(1) {
			v, err := dx.DecodeEnumerated(uplinkConfigExtEnablePLRSUpdateForType1CGPUSCHR18Constraints)
			if err != nil {
				return err
			}
			ie.EnablePL_RS_UpdateForType1CG_PUSCH_r18 = &v
		}

		if groupSeq.IsComponentPresent(2) {
			v, err := dx.DecodeBoolean()
			if err != nil {
				return err
			}
			ie.PowerBoostPi2BPSK_r18 = &v
		}

		if groupSeq.IsComponentPresent(3) {
			v, err := dx.DecodeBoolean()
			if err != nil {
				return err
			}
			ie.PowerBoostQPSK_r18 = &v
		}
	}

	return nil
}
