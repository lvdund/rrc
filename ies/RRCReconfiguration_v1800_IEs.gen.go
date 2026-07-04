// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: RRCReconfiguration-v1800-IEs (line 1033).

var rRCReconfigurationV1800IEsConstraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "needForInterruptionConfigNR-r18", Optional: true},
		{Name: "aerial-Config-r18", Optional: true},
		{Name: "sl-IndirectPathAddChange-r18", Optional: true},
		{Name: "n3c-IndirectPathAddChange-r18", Optional: true},
		{Name: "n3c-IndirectPathConfigRelay-r18", Optional: true},
		{Name: "otherConfig-v1800", Optional: true},
		{Name: "srs-PosResourceSetAggBW-CombinationList-r18", Optional: true},
		{Name: "ltm-Config-r18", Optional: true},
		{Name: "nonCriticalExtension", Optional: true},
	},
}

const (
	RRCReconfiguration_v1800_IEs_NeedForInterruptionConfigNR_r18_Disabled = 0
	RRCReconfiguration_v1800_IEs_NeedForInterruptionConfigNR_r18_Enabled  = 1
)

var rRCReconfigurationV1800IEsNeedForInterruptionConfigNRR18Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{RRCReconfiguration_v1800_IEs_NeedForInterruptionConfigNR_r18_Disabled, RRCReconfiguration_v1800_IEs_NeedForInterruptionConfigNR_r18_Enabled},
}

var rRCReconfiguration_v1800_IEsAerialConfigR18Constraints = per.ChoiceConstraints{
	Extensible: false,
	RootAlternatives: []per.AlternativeInfo{
		{Name: "release"},
		{Name: "setup"},
	},
}

const (
	RRCReconfiguration_v1800_IEs_Aerial_Config_r18_Release = 0
	RRCReconfiguration_v1800_IEs_Aerial_Config_r18_Setup   = 1
)

var rRCReconfiguration_v1800_IEsSlIndirectPathAddChangeR18Constraints = per.ChoiceConstraints{
	Extensible: false,
	RootAlternatives: []per.AlternativeInfo{
		{Name: "release"},
		{Name: "setup"},
	},
}

const (
	RRCReconfiguration_v1800_IEs_Sl_IndirectPathAddChange_r18_Release = 0
	RRCReconfiguration_v1800_IEs_Sl_IndirectPathAddChange_r18_Setup   = 1
)

var rRCReconfiguration_v1800_IEsN3cIndirectPathAddChangeR18Constraints = per.ChoiceConstraints{
	Extensible: false,
	RootAlternatives: []per.AlternativeInfo{
		{Name: "release"},
		{Name: "setup"},
	},
}

const (
	RRCReconfiguration_v1800_IEs_N3c_IndirectPathAddChange_r18_Release = 0
	RRCReconfiguration_v1800_IEs_N3c_IndirectPathAddChange_r18_Setup   = 1
)

var rRCReconfiguration_v1800_IEsN3cIndirectPathConfigRelayR18Constraints = per.ChoiceConstraints{
	Extensible: false,
	RootAlternatives: []per.AlternativeInfo{
		{Name: "release"},
		{Name: "setup"},
	},
}

const (
	RRCReconfiguration_v1800_IEs_N3c_IndirectPathConfigRelay_r18_Release = 0
	RRCReconfiguration_v1800_IEs_N3c_IndirectPathConfigRelay_r18_Setup   = 1
)

var rRCReconfiguration_v1800_IEsSrsPosResourceSetAggBWCombinationListR18Constraints = per.ChoiceConstraints{
	Extensible: false,
	RootAlternatives: []per.AlternativeInfo{
		{Name: "release"},
		{Name: "setup"},
	},
}

const (
	RRCReconfiguration_v1800_IEs_Srs_PosResourceSetAggBW_CombinationList_r18_Release = 0
	RRCReconfiguration_v1800_IEs_Srs_PosResourceSetAggBW_CombinationList_r18_Setup   = 1
)

var rRCReconfiguration_v1800_IEsLtmConfigR18Constraints = per.ChoiceConstraints{
	Extensible: false,
	RootAlternatives: []per.AlternativeInfo{
		{Name: "release"},
		{Name: "setup"},
	},
}

const (
	RRCReconfiguration_v1800_IEs_Ltm_Config_r18_Release = 0
	RRCReconfiguration_v1800_IEs_Ltm_Config_r18_Setup   = 1
)

type RRCReconfiguration_v1800_IEs struct {
	NeedForInterruptionConfigNR_r18 *int64
	Aerial_Config_r18               *struct {
		Choice  int
		Release *struct{}
		Setup   *Aerial_Config_r18
	}
	Sl_IndirectPathAddChange_r18 *struct {
		Choice  int
		Release *struct{}
		Setup   *SL_IndirectPathAddChange_r18
	}
	N3c_IndirectPathAddChange_r18 *struct {
		Choice  int
		Release *struct{}
		Setup   *N3C_IndirectPathAddChange_r18
	}
	N3c_IndirectPathConfigRelay_r18 *struct {
		Choice  int
		Release *struct{}
		Setup   *N3C_IndirectPathConfigRelay_r18
	}
	OtherConfig_v1800                           *OtherConfig_v1800
	Srs_PosResourceSetAggBW_CombinationList_r18 *struct {
		Choice  int
		Release *struct{}
		Setup   *SRS_PosResourceSetAggBW_CombinationList_r18
	}
	Ltm_Config_r18 *struct {
		Choice  int
		Release *struct{}
		Setup   *LTM_Config_r18
	}
	NonCriticalExtension *RRCReconfiguration_v1830_IEs
}

func (ie *RRCReconfiguration_v1800_IEs) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(rRCReconfigurationV1800IEsConstraints)

	// 1. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.NeedForInterruptionConfigNR_r18 != nil, ie.Aerial_Config_r18 != nil, ie.Sl_IndirectPathAddChange_r18 != nil, ie.N3c_IndirectPathAddChange_r18 != nil, ie.N3c_IndirectPathConfigRelay_r18 != nil, ie.OtherConfig_v1800 != nil, ie.Srs_PosResourceSetAggBW_CombinationList_r18 != nil, ie.Ltm_Config_r18 != nil, ie.NonCriticalExtension != nil}); err != nil {
		return err
	}

	// 2. needForInterruptionConfigNR-r18: enumerated
	{
		if ie.NeedForInterruptionConfigNR_r18 != nil {
			if err := e.EncodeEnumerated(*ie.NeedForInterruptionConfigNR_r18, rRCReconfigurationV1800IEsNeedForInterruptionConfigNRR18Constraints); err != nil {
				return err
			}
		}
	}

	// 3. aerial-Config-r18: choice
	{
		if ie.Aerial_Config_r18 != nil {
			choiceEnc := e.NewChoiceEncoder(rRCReconfiguration_v1800_IEsAerialConfigR18Constraints)
			if err := choiceEnc.EncodeChoice(int64((*ie.Aerial_Config_r18).Choice), false, nil); err != nil {
				return err
			}
			switch (*ie.Aerial_Config_r18).Choice {
			case RRCReconfiguration_v1800_IEs_Aerial_Config_r18_Release:
				if err := e.EncodeNull(); err != nil {
					return err
				}
			case RRCReconfiguration_v1800_IEs_Aerial_Config_r18_Setup:
				if err := (*ie.Aerial_Config_r18).Setup.Encode(e); err != nil {
					return err
				}
			default:
				return &per.ConstraintViolationError{TypeName: "choice", Value: int64((*ie.Aerial_Config_r18).Choice), Constraint: "index out of range"}
			}
		}
	}

	// 4. sl-IndirectPathAddChange-r18: choice
	{
		if ie.Sl_IndirectPathAddChange_r18 != nil {
			choiceEnc := e.NewChoiceEncoder(rRCReconfiguration_v1800_IEsSlIndirectPathAddChangeR18Constraints)
			if err := choiceEnc.EncodeChoice(int64((*ie.Sl_IndirectPathAddChange_r18).Choice), false, nil); err != nil {
				return err
			}
			switch (*ie.Sl_IndirectPathAddChange_r18).Choice {
			case RRCReconfiguration_v1800_IEs_Sl_IndirectPathAddChange_r18_Release:
				if err := e.EncodeNull(); err != nil {
					return err
				}
			case RRCReconfiguration_v1800_IEs_Sl_IndirectPathAddChange_r18_Setup:
				if err := (*ie.Sl_IndirectPathAddChange_r18).Setup.Encode(e); err != nil {
					return err
				}
			default:
				return &per.ConstraintViolationError{TypeName: "choice", Value: int64((*ie.Sl_IndirectPathAddChange_r18).Choice), Constraint: "index out of range"}
			}
		}
	}

	// 5. n3c-IndirectPathAddChange-r18: choice
	{
		if ie.N3c_IndirectPathAddChange_r18 != nil {
			choiceEnc := e.NewChoiceEncoder(rRCReconfiguration_v1800_IEsN3cIndirectPathAddChangeR18Constraints)
			if err := choiceEnc.EncodeChoice(int64((*ie.N3c_IndirectPathAddChange_r18).Choice), false, nil); err != nil {
				return err
			}
			switch (*ie.N3c_IndirectPathAddChange_r18).Choice {
			case RRCReconfiguration_v1800_IEs_N3c_IndirectPathAddChange_r18_Release:
				if err := e.EncodeNull(); err != nil {
					return err
				}
			case RRCReconfiguration_v1800_IEs_N3c_IndirectPathAddChange_r18_Setup:
				if err := (*ie.N3c_IndirectPathAddChange_r18).Setup.Encode(e); err != nil {
					return err
				}
			default:
				return &per.ConstraintViolationError{TypeName: "choice", Value: int64((*ie.N3c_IndirectPathAddChange_r18).Choice), Constraint: "index out of range"}
			}
		}
	}

	// 6. n3c-IndirectPathConfigRelay-r18: choice
	{
		if ie.N3c_IndirectPathConfigRelay_r18 != nil {
			choiceEnc := e.NewChoiceEncoder(rRCReconfiguration_v1800_IEsN3cIndirectPathConfigRelayR18Constraints)
			if err := choiceEnc.EncodeChoice(int64((*ie.N3c_IndirectPathConfigRelay_r18).Choice), false, nil); err != nil {
				return err
			}
			switch (*ie.N3c_IndirectPathConfigRelay_r18).Choice {
			case RRCReconfiguration_v1800_IEs_N3c_IndirectPathConfigRelay_r18_Release:
				if err := e.EncodeNull(); err != nil {
					return err
				}
			case RRCReconfiguration_v1800_IEs_N3c_IndirectPathConfigRelay_r18_Setup:
				if err := (*ie.N3c_IndirectPathConfigRelay_r18).Setup.Encode(e); err != nil {
					return err
				}
			default:
				return &per.ConstraintViolationError{TypeName: "choice", Value: int64((*ie.N3c_IndirectPathConfigRelay_r18).Choice), Constraint: "index out of range"}
			}
		}
	}

	// 7. otherConfig-v1800: ref
	{
		if ie.OtherConfig_v1800 != nil {
			if err := ie.OtherConfig_v1800.Encode(e); err != nil {
				return err
			}
		}
	}

	// 8. srs-PosResourceSetAggBW-CombinationList-r18: choice
	{
		if ie.Srs_PosResourceSetAggBW_CombinationList_r18 != nil {
			choiceEnc := e.NewChoiceEncoder(rRCReconfiguration_v1800_IEsSrsPosResourceSetAggBWCombinationListR18Constraints)
			if err := choiceEnc.EncodeChoice(int64((*ie.Srs_PosResourceSetAggBW_CombinationList_r18).Choice), false, nil); err != nil {
				return err
			}
			switch (*ie.Srs_PosResourceSetAggBW_CombinationList_r18).Choice {
			case RRCReconfiguration_v1800_IEs_Srs_PosResourceSetAggBW_CombinationList_r18_Release:
				if err := e.EncodeNull(); err != nil {
					return err
				}
			case RRCReconfiguration_v1800_IEs_Srs_PosResourceSetAggBW_CombinationList_r18_Setup:
				if err := (*ie.Srs_PosResourceSetAggBW_CombinationList_r18).Setup.Encode(e); err != nil {
					return err
				}
			default:
				return &per.ConstraintViolationError{TypeName: "choice", Value: int64((*ie.Srs_PosResourceSetAggBW_CombinationList_r18).Choice), Constraint: "index out of range"}
			}
		}
	}

	// 9. ltm-Config-r18: choice
	{
		if ie.Ltm_Config_r18 != nil {
			choiceEnc := e.NewChoiceEncoder(rRCReconfiguration_v1800_IEsLtmConfigR18Constraints)
			if err := choiceEnc.EncodeChoice(int64((*ie.Ltm_Config_r18).Choice), false, nil); err != nil {
				return err
			}
			switch (*ie.Ltm_Config_r18).Choice {
			case RRCReconfiguration_v1800_IEs_Ltm_Config_r18_Release:
				if err := e.EncodeNull(); err != nil {
					return err
				}
			case RRCReconfiguration_v1800_IEs_Ltm_Config_r18_Setup:
				if err := (*ie.Ltm_Config_r18).Setup.Encode(e); err != nil {
					return err
				}
			default:
				return &per.ConstraintViolationError{TypeName: "choice", Value: int64((*ie.Ltm_Config_r18).Choice), Constraint: "index out of range"}
			}
		}
	}

	// 10. nonCriticalExtension: ref
	{
		if ie.NonCriticalExtension != nil {
			if err := ie.NonCriticalExtension.Encode(e); err != nil {
				return err
			}
		}
	}

	return nil
}

func (ie *RRCReconfiguration_v1800_IEs) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(rRCReconfigurationV1800IEsConstraints)

	// 1. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 2. needForInterruptionConfigNR-r18: enumerated
	{
		if seq.IsComponentPresent(0) {
			idx, err := d.DecodeEnumerated(rRCReconfigurationV1800IEsNeedForInterruptionConfigNRR18Constraints)
			if err != nil {
				return err
			}
			ie.NeedForInterruptionConfigNR_r18 = &idx
		}
	}

	// 3. aerial-Config-r18: choice
	{
		if seq.IsComponentPresent(1) {
			ie.Aerial_Config_r18 = &struct {
				Choice  int
				Release *struct{}
				Setup   *Aerial_Config_r18
			}{}
			choiceDec := d.NewChoiceDecoder(rRCReconfiguration_v1800_IEsAerialConfigR18Constraints)
			index, isExt, _, err := choiceDec.DecodeChoice()
			if err != nil {
				return err
			}
			(*ie.Aerial_Config_r18).Choice = int(index)
			if isExt {
				return &per.DecodeError{TypeName: "choice", Reason: "extension alternative not supported"}
			}
			switch index {
			case RRCReconfiguration_v1800_IEs_Aerial_Config_r18_Release:
				(*ie.Aerial_Config_r18).Release = &struct{}{}
				if err := d.DecodeNull(); err != nil {
					return err
				}
			case RRCReconfiguration_v1800_IEs_Aerial_Config_r18_Setup:
				(*ie.Aerial_Config_r18).Setup = new(Aerial_Config_r18)
				if err := (*ie.Aerial_Config_r18).Setup.Decode(d); err != nil {
					return err
				}
			}
		}
	}

	// 4. sl-IndirectPathAddChange-r18: choice
	{
		if seq.IsComponentPresent(2) {
			ie.Sl_IndirectPathAddChange_r18 = &struct {
				Choice  int
				Release *struct{}
				Setup   *SL_IndirectPathAddChange_r18
			}{}
			choiceDec := d.NewChoiceDecoder(rRCReconfiguration_v1800_IEsSlIndirectPathAddChangeR18Constraints)
			index, isExt, _, err := choiceDec.DecodeChoice()
			if err != nil {
				return err
			}
			(*ie.Sl_IndirectPathAddChange_r18).Choice = int(index)
			if isExt {
				return &per.DecodeError{TypeName: "choice", Reason: "extension alternative not supported"}
			}
			switch index {
			case RRCReconfiguration_v1800_IEs_Sl_IndirectPathAddChange_r18_Release:
				(*ie.Sl_IndirectPathAddChange_r18).Release = &struct{}{}
				if err := d.DecodeNull(); err != nil {
					return err
				}
			case RRCReconfiguration_v1800_IEs_Sl_IndirectPathAddChange_r18_Setup:
				(*ie.Sl_IndirectPathAddChange_r18).Setup = new(SL_IndirectPathAddChange_r18)
				if err := (*ie.Sl_IndirectPathAddChange_r18).Setup.Decode(d); err != nil {
					return err
				}
			}
		}
	}

	// 5. n3c-IndirectPathAddChange-r18: choice
	{
		if seq.IsComponentPresent(3) {
			ie.N3c_IndirectPathAddChange_r18 = &struct {
				Choice  int
				Release *struct{}
				Setup   *N3C_IndirectPathAddChange_r18
			}{}
			choiceDec := d.NewChoiceDecoder(rRCReconfiguration_v1800_IEsN3cIndirectPathAddChangeR18Constraints)
			index, isExt, _, err := choiceDec.DecodeChoice()
			if err != nil {
				return err
			}
			(*ie.N3c_IndirectPathAddChange_r18).Choice = int(index)
			if isExt {
				return &per.DecodeError{TypeName: "choice", Reason: "extension alternative not supported"}
			}
			switch index {
			case RRCReconfiguration_v1800_IEs_N3c_IndirectPathAddChange_r18_Release:
				(*ie.N3c_IndirectPathAddChange_r18).Release = &struct{}{}
				if err := d.DecodeNull(); err != nil {
					return err
				}
			case RRCReconfiguration_v1800_IEs_N3c_IndirectPathAddChange_r18_Setup:
				(*ie.N3c_IndirectPathAddChange_r18).Setup = new(N3C_IndirectPathAddChange_r18)
				if err := (*ie.N3c_IndirectPathAddChange_r18).Setup.Decode(d); err != nil {
					return err
				}
			}
		}
	}

	// 6. n3c-IndirectPathConfigRelay-r18: choice
	{
		if seq.IsComponentPresent(4) {
			ie.N3c_IndirectPathConfigRelay_r18 = &struct {
				Choice  int
				Release *struct{}
				Setup   *N3C_IndirectPathConfigRelay_r18
			}{}
			choiceDec := d.NewChoiceDecoder(rRCReconfiguration_v1800_IEsN3cIndirectPathConfigRelayR18Constraints)
			index, isExt, _, err := choiceDec.DecodeChoice()
			if err != nil {
				return err
			}
			(*ie.N3c_IndirectPathConfigRelay_r18).Choice = int(index)
			if isExt {
				return &per.DecodeError{TypeName: "choice", Reason: "extension alternative not supported"}
			}
			switch index {
			case RRCReconfiguration_v1800_IEs_N3c_IndirectPathConfigRelay_r18_Release:
				(*ie.N3c_IndirectPathConfigRelay_r18).Release = &struct{}{}
				if err := d.DecodeNull(); err != nil {
					return err
				}
			case RRCReconfiguration_v1800_IEs_N3c_IndirectPathConfigRelay_r18_Setup:
				(*ie.N3c_IndirectPathConfigRelay_r18).Setup = new(N3C_IndirectPathConfigRelay_r18)
				if err := (*ie.N3c_IndirectPathConfigRelay_r18).Setup.Decode(d); err != nil {
					return err
				}
			}
		}
	}

	// 7. otherConfig-v1800: ref
	{
		if seq.IsComponentPresent(5) {
			ie.OtherConfig_v1800 = new(OtherConfig_v1800)
			if err := ie.OtherConfig_v1800.Decode(d); err != nil {
				return err
			}
		}
	}

	// 8. srs-PosResourceSetAggBW-CombinationList-r18: choice
	{
		if seq.IsComponentPresent(6) {
			ie.Srs_PosResourceSetAggBW_CombinationList_r18 = &struct {
				Choice  int
				Release *struct{}
				Setup   *SRS_PosResourceSetAggBW_CombinationList_r18
			}{}
			choiceDec := d.NewChoiceDecoder(rRCReconfiguration_v1800_IEsSrsPosResourceSetAggBWCombinationListR18Constraints)
			index, isExt, _, err := choiceDec.DecodeChoice()
			if err != nil {
				return err
			}
			(*ie.Srs_PosResourceSetAggBW_CombinationList_r18).Choice = int(index)
			if isExt {
				return &per.DecodeError{TypeName: "choice", Reason: "extension alternative not supported"}
			}
			switch index {
			case RRCReconfiguration_v1800_IEs_Srs_PosResourceSetAggBW_CombinationList_r18_Release:
				(*ie.Srs_PosResourceSetAggBW_CombinationList_r18).Release = &struct{}{}
				if err := d.DecodeNull(); err != nil {
					return err
				}
			case RRCReconfiguration_v1800_IEs_Srs_PosResourceSetAggBW_CombinationList_r18_Setup:
				(*ie.Srs_PosResourceSetAggBW_CombinationList_r18).Setup = new(SRS_PosResourceSetAggBW_CombinationList_r18)
				if err := (*ie.Srs_PosResourceSetAggBW_CombinationList_r18).Setup.Decode(d); err != nil {
					return err
				}
			}
		}
	}

	// 9. ltm-Config-r18: choice
	{
		if seq.IsComponentPresent(7) {
			ie.Ltm_Config_r18 = &struct {
				Choice  int
				Release *struct{}
				Setup   *LTM_Config_r18
			}{}
			choiceDec := d.NewChoiceDecoder(rRCReconfiguration_v1800_IEsLtmConfigR18Constraints)
			index, isExt, _, err := choiceDec.DecodeChoice()
			if err != nil {
				return err
			}
			(*ie.Ltm_Config_r18).Choice = int(index)
			if isExt {
				return &per.DecodeError{TypeName: "choice", Reason: "extension alternative not supported"}
			}
			switch index {
			case RRCReconfiguration_v1800_IEs_Ltm_Config_r18_Release:
				(*ie.Ltm_Config_r18).Release = &struct{}{}
				if err := d.DecodeNull(); err != nil {
					return err
				}
			case RRCReconfiguration_v1800_IEs_Ltm_Config_r18_Setup:
				(*ie.Ltm_Config_r18).Setup = new(LTM_Config_r18)
				if err := (*ie.Ltm_Config_r18).Setup.Decode(d); err != nil {
					return err
				}
			}
		}
	}

	// 10. nonCriticalExtension: ref
	{
		if seq.IsComponentPresent(8) {
			ie.NonCriticalExtension = new(RRCReconfiguration_v1830_IEs)
			if err := ie.NonCriticalExtension.Decode(d); err != nil {
				return err
			}
		}
	}

	return nil
}
