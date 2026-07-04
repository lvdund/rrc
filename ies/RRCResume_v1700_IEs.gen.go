// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: RRCResume-v1700-IEs (line 1564).

var rRCResumeV1700IEsConstraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "sl-ConfigDedicatedNR-r17", Optional: true},
		{Name: "sl-L2RemoteUE-Config-r17", Optional: true},
		{Name: "needForGapNCSG-ConfigNR-r17", Optional: true},
		{Name: "needForGapNCSG-ConfigEUTRA-r17", Optional: true},
		{Name: "scg-State-r17", Optional: true},
		{Name: "appLayerMeasConfig-r17", Optional: true},
		{Name: "nonCriticalExtension", Optional: true},
	},
}

var rRCResume_v1700_IEsSlConfigDedicatedNRR17Constraints = per.ChoiceConstraints{
	Extensible: false,
	RootAlternatives: []per.AlternativeInfo{
		{Name: "release"},
		{Name: "setup"},
	},
}

const (
	RRCResume_v1700_IEs_Sl_ConfigDedicatedNR_r17_Release = 0
	RRCResume_v1700_IEs_Sl_ConfigDedicatedNR_r17_Setup   = 1
)

var rRCResume_v1700_IEsSlL2RemoteUEConfigR17Constraints = per.ChoiceConstraints{
	Extensible: false,
	RootAlternatives: []per.AlternativeInfo{
		{Name: "release"},
		{Name: "setup"},
	},
}

const (
	RRCResume_v1700_IEs_Sl_L2RemoteUE_Config_r17_Release = 0
	RRCResume_v1700_IEs_Sl_L2RemoteUE_Config_r17_Setup   = 1
)

var rRCResume_v1700_IEsNeedForGapNCSGConfigNRR17Constraints = per.ChoiceConstraints{
	Extensible: false,
	RootAlternatives: []per.AlternativeInfo{
		{Name: "release"},
		{Name: "setup"},
	},
}

const (
	RRCResume_v1700_IEs_NeedForGapNCSG_ConfigNR_r17_Release = 0
	RRCResume_v1700_IEs_NeedForGapNCSG_ConfigNR_r17_Setup   = 1
)

var rRCResume_v1700_IEsNeedForGapNCSGConfigEUTRAR17Constraints = per.ChoiceConstraints{
	Extensible: false,
	RootAlternatives: []per.AlternativeInfo{
		{Name: "release"},
		{Name: "setup"},
	},
}

const (
	RRCResume_v1700_IEs_NeedForGapNCSG_ConfigEUTRA_r17_Release = 0
	RRCResume_v1700_IEs_NeedForGapNCSG_ConfigEUTRA_r17_Setup   = 1
)

const (
	RRCResume_v1700_IEs_Scg_State_r17_Deactivated = 0
)

var rRCResumeV1700IEsScgStateR17Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{RRCResume_v1700_IEs_Scg_State_r17_Deactivated},
}

type RRCResume_v1700_IEs struct {
	Sl_ConfigDedicatedNR_r17 *struct {
		Choice  int
		Release *struct{}
		Setup   *SL_ConfigDedicatedNR_r16
	}
	Sl_L2RemoteUE_Config_r17 *struct {
		Choice  int
		Release *struct{}
		Setup   *SL_L2RemoteUE_Config_r17
	}
	NeedForGapNCSG_ConfigNR_r17 *struct {
		Choice  int
		Release *struct{}
		Setup   *NeedForGapNCSG_ConfigNR_r17
	}
	NeedForGapNCSG_ConfigEUTRA_r17 *struct {
		Choice  int
		Release *struct{}
		Setup   *NeedForGapNCSG_ConfigEUTRA_r17
	}
	Scg_State_r17          *int64
	AppLayerMeasConfig_r17 *AppLayerMeasConfig_r17
	NonCriticalExtension   *RRCResume_v1800_IEs
}

func (ie *RRCResume_v1700_IEs) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(rRCResumeV1700IEsConstraints)

	// 1. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.Sl_ConfigDedicatedNR_r17 != nil, ie.Sl_L2RemoteUE_Config_r17 != nil, ie.NeedForGapNCSG_ConfigNR_r17 != nil, ie.NeedForGapNCSG_ConfigEUTRA_r17 != nil, ie.Scg_State_r17 != nil, ie.AppLayerMeasConfig_r17 != nil, ie.NonCriticalExtension != nil}); err != nil {
		return err
	}

	// 2. sl-ConfigDedicatedNR-r17: choice
	{
		if ie.Sl_ConfigDedicatedNR_r17 != nil {
			choiceEnc := e.NewChoiceEncoder(rRCResume_v1700_IEsSlConfigDedicatedNRR17Constraints)
			if err := choiceEnc.EncodeChoice(int64((*ie.Sl_ConfigDedicatedNR_r17).Choice), false, nil); err != nil {
				return err
			}
			switch (*ie.Sl_ConfigDedicatedNR_r17).Choice {
			case RRCResume_v1700_IEs_Sl_ConfigDedicatedNR_r17_Release:
				if err := e.EncodeNull(); err != nil {
					return err
				}
			case RRCResume_v1700_IEs_Sl_ConfigDedicatedNR_r17_Setup:
				if err := (*ie.Sl_ConfigDedicatedNR_r17).Setup.Encode(e); err != nil {
					return err
				}
			default:
				return &per.ConstraintViolationError{TypeName: "choice", Value: int64((*ie.Sl_ConfigDedicatedNR_r17).Choice), Constraint: "index out of range"}
			}
		}
	}

	// 3. sl-L2RemoteUE-Config-r17: choice
	{
		if ie.Sl_L2RemoteUE_Config_r17 != nil {
			choiceEnc := e.NewChoiceEncoder(rRCResume_v1700_IEsSlL2RemoteUEConfigR17Constraints)
			if err := choiceEnc.EncodeChoice(int64((*ie.Sl_L2RemoteUE_Config_r17).Choice), false, nil); err != nil {
				return err
			}
			switch (*ie.Sl_L2RemoteUE_Config_r17).Choice {
			case RRCResume_v1700_IEs_Sl_L2RemoteUE_Config_r17_Release:
				if err := e.EncodeNull(); err != nil {
					return err
				}
			case RRCResume_v1700_IEs_Sl_L2RemoteUE_Config_r17_Setup:
				if err := (*ie.Sl_L2RemoteUE_Config_r17).Setup.Encode(e); err != nil {
					return err
				}
			default:
				return &per.ConstraintViolationError{TypeName: "choice", Value: int64((*ie.Sl_L2RemoteUE_Config_r17).Choice), Constraint: "index out of range"}
			}
		}
	}

	// 4. needForGapNCSG-ConfigNR-r17: choice
	{
		if ie.NeedForGapNCSG_ConfigNR_r17 != nil {
			choiceEnc := e.NewChoiceEncoder(rRCResume_v1700_IEsNeedForGapNCSGConfigNRR17Constraints)
			if err := choiceEnc.EncodeChoice(int64((*ie.NeedForGapNCSG_ConfigNR_r17).Choice), false, nil); err != nil {
				return err
			}
			switch (*ie.NeedForGapNCSG_ConfigNR_r17).Choice {
			case RRCResume_v1700_IEs_NeedForGapNCSG_ConfigNR_r17_Release:
				if err := e.EncodeNull(); err != nil {
					return err
				}
			case RRCResume_v1700_IEs_NeedForGapNCSG_ConfigNR_r17_Setup:
				if err := (*ie.NeedForGapNCSG_ConfigNR_r17).Setup.Encode(e); err != nil {
					return err
				}
			default:
				return &per.ConstraintViolationError{TypeName: "choice", Value: int64((*ie.NeedForGapNCSG_ConfigNR_r17).Choice), Constraint: "index out of range"}
			}
		}
	}

	// 5. needForGapNCSG-ConfigEUTRA-r17: choice
	{
		if ie.NeedForGapNCSG_ConfigEUTRA_r17 != nil {
			choiceEnc := e.NewChoiceEncoder(rRCResume_v1700_IEsNeedForGapNCSGConfigEUTRAR17Constraints)
			if err := choiceEnc.EncodeChoice(int64((*ie.NeedForGapNCSG_ConfigEUTRA_r17).Choice), false, nil); err != nil {
				return err
			}
			switch (*ie.NeedForGapNCSG_ConfigEUTRA_r17).Choice {
			case RRCResume_v1700_IEs_NeedForGapNCSG_ConfigEUTRA_r17_Release:
				if err := e.EncodeNull(); err != nil {
					return err
				}
			case RRCResume_v1700_IEs_NeedForGapNCSG_ConfigEUTRA_r17_Setup:
				if err := (*ie.NeedForGapNCSG_ConfigEUTRA_r17).Setup.Encode(e); err != nil {
					return err
				}
			default:
				return &per.ConstraintViolationError{TypeName: "choice", Value: int64((*ie.NeedForGapNCSG_ConfigEUTRA_r17).Choice), Constraint: "index out of range"}
			}
		}
	}

	// 6. scg-State-r17: enumerated
	{
		if ie.Scg_State_r17 != nil {
			if err := e.EncodeEnumerated(*ie.Scg_State_r17, rRCResumeV1700IEsScgStateR17Constraints); err != nil {
				return err
			}
		}
	}

	// 7. appLayerMeasConfig-r17: ref
	{
		if ie.AppLayerMeasConfig_r17 != nil {
			if err := ie.AppLayerMeasConfig_r17.Encode(e); err != nil {
				return err
			}
		}
	}

	// 8. nonCriticalExtension: ref
	{
		if ie.NonCriticalExtension != nil {
			if err := ie.NonCriticalExtension.Encode(e); err != nil {
				return err
			}
		}
	}

	return nil
}

func (ie *RRCResume_v1700_IEs) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(rRCResumeV1700IEsConstraints)

	// 1. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 2. sl-ConfigDedicatedNR-r17: choice
	{
		if seq.IsComponentPresent(0) {
			ie.Sl_ConfigDedicatedNR_r17 = &struct {
				Choice  int
				Release *struct{}
				Setup   *SL_ConfigDedicatedNR_r16
			}{}
			choiceDec := d.NewChoiceDecoder(rRCResume_v1700_IEsSlConfigDedicatedNRR17Constraints)
			index, isExt, _, err := choiceDec.DecodeChoice()
			if err != nil {
				return err
			}
			(*ie.Sl_ConfigDedicatedNR_r17).Choice = int(index)
			if isExt {
				return &per.DecodeError{TypeName: "choice", Reason: "extension alternative not supported"}
			}
			switch index {
			case RRCResume_v1700_IEs_Sl_ConfigDedicatedNR_r17_Release:
				(*ie.Sl_ConfigDedicatedNR_r17).Release = &struct{}{}
				if err := d.DecodeNull(); err != nil {
					return err
				}
			case RRCResume_v1700_IEs_Sl_ConfigDedicatedNR_r17_Setup:
				(*ie.Sl_ConfigDedicatedNR_r17).Setup = new(SL_ConfigDedicatedNR_r16)
				if err := (*ie.Sl_ConfigDedicatedNR_r17).Setup.Decode(d); err != nil {
					return err
				}
			}
		}
	}

	// 3. sl-L2RemoteUE-Config-r17: choice
	{
		if seq.IsComponentPresent(1) {
			ie.Sl_L2RemoteUE_Config_r17 = &struct {
				Choice  int
				Release *struct{}
				Setup   *SL_L2RemoteUE_Config_r17
			}{}
			choiceDec := d.NewChoiceDecoder(rRCResume_v1700_IEsSlL2RemoteUEConfigR17Constraints)
			index, isExt, _, err := choiceDec.DecodeChoice()
			if err != nil {
				return err
			}
			(*ie.Sl_L2RemoteUE_Config_r17).Choice = int(index)
			if isExt {
				return &per.DecodeError{TypeName: "choice", Reason: "extension alternative not supported"}
			}
			switch index {
			case RRCResume_v1700_IEs_Sl_L2RemoteUE_Config_r17_Release:
				(*ie.Sl_L2RemoteUE_Config_r17).Release = &struct{}{}
				if err := d.DecodeNull(); err != nil {
					return err
				}
			case RRCResume_v1700_IEs_Sl_L2RemoteUE_Config_r17_Setup:
				(*ie.Sl_L2RemoteUE_Config_r17).Setup = new(SL_L2RemoteUE_Config_r17)
				if err := (*ie.Sl_L2RemoteUE_Config_r17).Setup.Decode(d); err != nil {
					return err
				}
			}
		}
	}

	// 4. needForGapNCSG-ConfigNR-r17: choice
	{
		if seq.IsComponentPresent(2) {
			ie.NeedForGapNCSG_ConfigNR_r17 = &struct {
				Choice  int
				Release *struct{}
				Setup   *NeedForGapNCSG_ConfigNR_r17
			}{}
			choiceDec := d.NewChoiceDecoder(rRCResume_v1700_IEsNeedForGapNCSGConfigNRR17Constraints)
			index, isExt, _, err := choiceDec.DecodeChoice()
			if err != nil {
				return err
			}
			(*ie.NeedForGapNCSG_ConfigNR_r17).Choice = int(index)
			if isExt {
				return &per.DecodeError{TypeName: "choice", Reason: "extension alternative not supported"}
			}
			switch index {
			case RRCResume_v1700_IEs_NeedForGapNCSG_ConfigNR_r17_Release:
				(*ie.NeedForGapNCSG_ConfigNR_r17).Release = &struct{}{}
				if err := d.DecodeNull(); err != nil {
					return err
				}
			case RRCResume_v1700_IEs_NeedForGapNCSG_ConfigNR_r17_Setup:
				(*ie.NeedForGapNCSG_ConfigNR_r17).Setup = new(NeedForGapNCSG_ConfigNR_r17)
				if err := (*ie.NeedForGapNCSG_ConfigNR_r17).Setup.Decode(d); err != nil {
					return err
				}
			}
		}
	}

	// 5. needForGapNCSG-ConfigEUTRA-r17: choice
	{
		if seq.IsComponentPresent(3) {
			ie.NeedForGapNCSG_ConfigEUTRA_r17 = &struct {
				Choice  int
				Release *struct{}
				Setup   *NeedForGapNCSG_ConfigEUTRA_r17
			}{}
			choiceDec := d.NewChoiceDecoder(rRCResume_v1700_IEsNeedForGapNCSGConfigEUTRAR17Constraints)
			index, isExt, _, err := choiceDec.DecodeChoice()
			if err != nil {
				return err
			}
			(*ie.NeedForGapNCSG_ConfigEUTRA_r17).Choice = int(index)
			if isExt {
				return &per.DecodeError{TypeName: "choice", Reason: "extension alternative not supported"}
			}
			switch index {
			case RRCResume_v1700_IEs_NeedForGapNCSG_ConfigEUTRA_r17_Release:
				(*ie.NeedForGapNCSG_ConfigEUTRA_r17).Release = &struct{}{}
				if err := d.DecodeNull(); err != nil {
					return err
				}
			case RRCResume_v1700_IEs_NeedForGapNCSG_ConfigEUTRA_r17_Setup:
				(*ie.NeedForGapNCSG_ConfigEUTRA_r17).Setup = new(NeedForGapNCSG_ConfigEUTRA_r17)
				if err := (*ie.NeedForGapNCSG_ConfigEUTRA_r17).Setup.Decode(d); err != nil {
					return err
				}
			}
		}
	}

	// 6. scg-State-r17: enumerated
	{
		if seq.IsComponentPresent(4) {
			idx, err := d.DecodeEnumerated(rRCResumeV1700IEsScgStateR17Constraints)
			if err != nil {
				return err
			}
			ie.Scg_State_r17 = &idx
		}
	}

	// 7. appLayerMeasConfig-r17: ref
	{
		if seq.IsComponentPresent(5) {
			ie.AppLayerMeasConfig_r17 = new(AppLayerMeasConfig_r17)
			if err := ie.AppLayerMeasConfig_r17.Decode(d); err != nil {
				return err
			}
		}
	}

	// 8. nonCriticalExtension: ref
	{
		if seq.IsComponentPresent(6) {
			ie.NonCriticalExtension = new(RRCResume_v1800_IEs)
			if err := ie.NonCriticalExtension.Decode(d); err != nil {
				return err
			}
		}
	}

	return nil
}
