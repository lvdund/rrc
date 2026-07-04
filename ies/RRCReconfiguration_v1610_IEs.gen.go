// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: RRCReconfiguration-v1610-IEs (line 1002).

var rRCReconfigurationV1610IEsConstraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "otherConfig-v1610", Optional: true},
		{Name: "bap-Config-r16", Optional: true},
		{Name: "iab-IP-AddressConfigurationList-r16", Optional: true},
		{Name: "conditionalReconfiguration-r16", Optional: true},
		{Name: "daps-SourceRelease-r16", Optional: true},
		{Name: "t316-r16", Optional: true},
		{Name: "needForGapsConfigNR-r16", Optional: true},
		{Name: "onDemandSIB-Request-r16", Optional: true},
		{Name: "dedicatedPosSysInfoDelivery-r16", Optional: true},
		{Name: "sl-ConfigDedicatedNR-r16", Optional: true},
		{Name: "sl-ConfigDedicatedEUTRA-Info-r16", Optional: true},
		{Name: "targetCellSMTC-SCG-r16", Optional: true},
		{Name: "nonCriticalExtension", Optional: true},
	},
}

var rRCReconfiguration_v1610_IEsBapConfigR16Constraints = per.ChoiceConstraints{
	Extensible: false,
	RootAlternatives: []per.AlternativeInfo{
		{Name: "release"},
		{Name: "setup"},
	},
}

const (
	RRCReconfiguration_v1610_IEs_Bap_Config_r16_Release = 0
	RRCReconfiguration_v1610_IEs_Bap_Config_r16_Setup   = 1
)

const (
	RRCReconfiguration_v1610_IEs_Daps_SourceRelease_r16_True = 0
)

var rRCReconfigurationV1610IEsDapsSourceReleaseR16Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{RRCReconfiguration_v1610_IEs_Daps_SourceRelease_r16_True},
}

var rRCReconfiguration_v1610_IEsT316R16Constraints = per.ChoiceConstraints{
	Extensible: false,
	RootAlternatives: []per.AlternativeInfo{
		{Name: "release"},
		{Name: "setup"},
	},
}

const (
	RRCReconfiguration_v1610_IEs_T316_r16_Release = 0
	RRCReconfiguration_v1610_IEs_T316_r16_Setup   = 1
)

var rRCReconfiguration_v1610_IEsNeedForGapsConfigNRR16Constraints = per.ChoiceConstraints{
	Extensible: false,
	RootAlternatives: []per.AlternativeInfo{
		{Name: "release"},
		{Name: "setup"},
	},
}

const (
	RRCReconfiguration_v1610_IEs_NeedForGapsConfigNR_r16_Release = 0
	RRCReconfiguration_v1610_IEs_NeedForGapsConfigNR_r16_Setup   = 1
)

var rRCReconfiguration_v1610_IEsOnDemandSIBRequestR16Constraints = per.ChoiceConstraints{
	Extensible: false,
	RootAlternatives: []per.AlternativeInfo{
		{Name: "release"},
		{Name: "setup"},
	},
}

const (
	RRCReconfiguration_v1610_IEs_OnDemandSIB_Request_r16_Release = 0
	RRCReconfiguration_v1610_IEs_OnDemandSIB_Request_r16_Setup   = 1
)

var rRCReconfigurationV1610IEsDedicatedPosSysInfoDeliveryR16Constraints = per.SizeConstraints{}

var rRCReconfiguration_v1610_IEsSlConfigDedicatedNRR16Constraints = per.ChoiceConstraints{
	Extensible: false,
	RootAlternatives: []per.AlternativeInfo{
		{Name: "release"},
		{Name: "setup"},
	},
}

const (
	RRCReconfiguration_v1610_IEs_Sl_ConfigDedicatedNR_r16_Release = 0
	RRCReconfiguration_v1610_IEs_Sl_ConfigDedicatedNR_r16_Setup   = 1
)

var rRCReconfiguration_v1610_IEsSlConfigDedicatedEUTRAInfoR16Constraints = per.ChoiceConstraints{
	Extensible: false,
	RootAlternatives: []per.AlternativeInfo{
		{Name: "release"},
		{Name: "setup"},
	},
}

const (
	RRCReconfiguration_v1610_IEs_Sl_ConfigDedicatedEUTRA_Info_r16_Release = 0
	RRCReconfiguration_v1610_IEs_Sl_ConfigDedicatedEUTRA_Info_r16_Setup   = 1
)

type RRCReconfiguration_v1610_IEs struct {
	OtherConfig_v1610 *OtherConfig_v1610
	Bap_Config_r16    *struct {
		Choice  int
		Release *struct{}
		Setup   *BAP_Config_r16
	}
	Iab_IP_AddressConfigurationList_r16 *IAB_IP_AddressConfigurationList_r16
	ConditionalReconfiguration_r16      *ConditionalReconfiguration_r16
	Daps_SourceRelease_r16              *int64
	T316_r16                            *struct {
		Choice  int
		Release *struct{}
		Setup   *T316_r16
	}
	NeedForGapsConfigNR_r16 *struct {
		Choice  int
		Release *struct{}
		Setup   *NeedForGapsConfigNR_r16
	}
	OnDemandSIB_Request_r16 *struct {
		Choice  int
		Release *struct{}
		Setup   *OnDemandSIB_Request_r16
	}
	DedicatedPosSysInfoDelivery_r16 []byte
	Sl_ConfigDedicatedNR_r16        *struct {
		Choice  int
		Release *struct{}
		Setup   *SL_ConfigDedicatedNR_r16
	}
	Sl_ConfigDedicatedEUTRA_Info_r16 *struct {
		Choice  int
		Release *struct{}
		Setup   *SL_ConfigDedicatedEUTRA_Info_r16
	}
	TargetCellSMTC_SCG_r16 *SSB_MTC
	NonCriticalExtension   *RRCReconfiguration_v1700_IEs
}

func (ie *RRCReconfiguration_v1610_IEs) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(rRCReconfigurationV1610IEsConstraints)

	// 1. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.OtherConfig_v1610 != nil, ie.Bap_Config_r16 != nil, ie.Iab_IP_AddressConfigurationList_r16 != nil, ie.ConditionalReconfiguration_r16 != nil, ie.Daps_SourceRelease_r16 != nil, ie.T316_r16 != nil, ie.NeedForGapsConfigNR_r16 != nil, ie.OnDemandSIB_Request_r16 != nil, ie.DedicatedPosSysInfoDelivery_r16 != nil, ie.Sl_ConfigDedicatedNR_r16 != nil, ie.Sl_ConfigDedicatedEUTRA_Info_r16 != nil, ie.TargetCellSMTC_SCG_r16 != nil, ie.NonCriticalExtension != nil}); err != nil {
		return err
	}

	// 2. otherConfig-v1610: ref
	{
		if ie.OtherConfig_v1610 != nil {
			if err := ie.OtherConfig_v1610.Encode(e); err != nil {
				return err
			}
		}
	}

	// 3. bap-Config-r16: choice
	{
		if ie.Bap_Config_r16 != nil {
			choiceEnc := e.NewChoiceEncoder(rRCReconfiguration_v1610_IEsBapConfigR16Constraints)
			if err := choiceEnc.EncodeChoice(int64((*ie.Bap_Config_r16).Choice), false, nil); err != nil {
				return err
			}
			switch (*ie.Bap_Config_r16).Choice {
			case RRCReconfiguration_v1610_IEs_Bap_Config_r16_Release:
				if err := e.EncodeNull(); err != nil {
					return err
				}
			case RRCReconfiguration_v1610_IEs_Bap_Config_r16_Setup:
				if err := (*ie.Bap_Config_r16).Setup.Encode(e); err != nil {
					return err
				}
			default:
				return &per.ConstraintViolationError{TypeName: "choice", Value: int64((*ie.Bap_Config_r16).Choice), Constraint: "index out of range"}
			}
		}
	}

	// 4. iab-IP-AddressConfigurationList-r16: ref
	{
		if ie.Iab_IP_AddressConfigurationList_r16 != nil {
			if err := ie.Iab_IP_AddressConfigurationList_r16.Encode(e); err != nil {
				return err
			}
		}
	}

	// 5. conditionalReconfiguration-r16: ref
	{
		if ie.ConditionalReconfiguration_r16 != nil {
			if err := ie.ConditionalReconfiguration_r16.Encode(e); err != nil {
				return err
			}
		}
	}

	// 6. daps-SourceRelease-r16: enumerated
	{
		if ie.Daps_SourceRelease_r16 != nil {
			if err := e.EncodeEnumerated(*ie.Daps_SourceRelease_r16, rRCReconfigurationV1610IEsDapsSourceReleaseR16Constraints); err != nil {
				return err
			}
		}
	}

	// 7. t316-r16: choice
	{
		if ie.T316_r16 != nil {
			choiceEnc := e.NewChoiceEncoder(rRCReconfiguration_v1610_IEsT316R16Constraints)
			if err := choiceEnc.EncodeChoice(int64((*ie.T316_r16).Choice), false, nil); err != nil {
				return err
			}
			switch (*ie.T316_r16).Choice {
			case RRCReconfiguration_v1610_IEs_T316_r16_Release:
				if err := e.EncodeNull(); err != nil {
					return err
				}
			case RRCReconfiguration_v1610_IEs_T316_r16_Setup:
				if err := (*ie.T316_r16).Setup.Encode(e); err != nil {
					return err
				}
			default:
				return &per.ConstraintViolationError{TypeName: "choice", Value: int64((*ie.T316_r16).Choice), Constraint: "index out of range"}
			}
		}
	}

	// 8. needForGapsConfigNR-r16: choice
	{
		if ie.NeedForGapsConfigNR_r16 != nil {
			choiceEnc := e.NewChoiceEncoder(rRCReconfiguration_v1610_IEsNeedForGapsConfigNRR16Constraints)
			if err := choiceEnc.EncodeChoice(int64((*ie.NeedForGapsConfigNR_r16).Choice), false, nil); err != nil {
				return err
			}
			switch (*ie.NeedForGapsConfigNR_r16).Choice {
			case RRCReconfiguration_v1610_IEs_NeedForGapsConfigNR_r16_Release:
				if err := e.EncodeNull(); err != nil {
					return err
				}
			case RRCReconfiguration_v1610_IEs_NeedForGapsConfigNR_r16_Setup:
				if err := (*ie.NeedForGapsConfigNR_r16).Setup.Encode(e); err != nil {
					return err
				}
			default:
				return &per.ConstraintViolationError{TypeName: "choice", Value: int64((*ie.NeedForGapsConfigNR_r16).Choice), Constraint: "index out of range"}
			}
		}
	}

	// 9. onDemandSIB-Request-r16: choice
	{
		if ie.OnDemandSIB_Request_r16 != nil {
			choiceEnc := e.NewChoiceEncoder(rRCReconfiguration_v1610_IEsOnDemandSIBRequestR16Constraints)
			if err := choiceEnc.EncodeChoice(int64((*ie.OnDemandSIB_Request_r16).Choice), false, nil); err != nil {
				return err
			}
			switch (*ie.OnDemandSIB_Request_r16).Choice {
			case RRCReconfiguration_v1610_IEs_OnDemandSIB_Request_r16_Release:
				if err := e.EncodeNull(); err != nil {
					return err
				}
			case RRCReconfiguration_v1610_IEs_OnDemandSIB_Request_r16_Setup:
				if err := (*ie.OnDemandSIB_Request_r16).Setup.Encode(e); err != nil {
					return err
				}
			default:
				return &per.ConstraintViolationError{TypeName: "choice", Value: int64((*ie.OnDemandSIB_Request_r16).Choice), Constraint: "index out of range"}
			}
		}
	}

	// 10. dedicatedPosSysInfoDelivery-r16: octet-string
	{
		if ie.DedicatedPosSysInfoDelivery_r16 != nil {
			if err := e.EncodeOctetString(ie.DedicatedPosSysInfoDelivery_r16, rRCReconfigurationV1610IEsDedicatedPosSysInfoDeliveryR16Constraints); err != nil {
				return err
			}
		}
	}

	// 11. sl-ConfigDedicatedNR-r16: choice
	{
		if ie.Sl_ConfigDedicatedNR_r16 != nil {
			choiceEnc := e.NewChoiceEncoder(rRCReconfiguration_v1610_IEsSlConfigDedicatedNRR16Constraints)
			if err := choiceEnc.EncodeChoice(int64((*ie.Sl_ConfigDedicatedNR_r16).Choice), false, nil); err != nil {
				return err
			}
			switch (*ie.Sl_ConfigDedicatedNR_r16).Choice {
			case RRCReconfiguration_v1610_IEs_Sl_ConfigDedicatedNR_r16_Release:
				if err := e.EncodeNull(); err != nil {
					return err
				}
			case RRCReconfiguration_v1610_IEs_Sl_ConfigDedicatedNR_r16_Setup:
				if err := (*ie.Sl_ConfigDedicatedNR_r16).Setup.Encode(e); err != nil {
					return err
				}
			default:
				return &per.ConstraintViolationError{TypeName: "choice", Value: int64((*ie.Sl_ConfigDedicatedNR_r16).Choice), Constraint: "index out of range"}
			}
		}
	}

	// 12. sl-ConfigDedicatedEUTRA-Info-r16: choice
	{
		if ie.Sl_ConfigDedicatedEUTRA_Info_r16 != nil {
			choiceEnc := e.NewChoiceEncoder(rRCReconfiguration_v1610_IEsSlConfigDedicatedEUTRAInfoR16Constraints)
			if err := choiceEnc.EncodeChoice(int64((*ie.Sl_ConfigDedicatedEUTRA_Info_r16).Choice), false, nil); err != nil {
				return err
			}
			switch (*ie.Sl_ConfigDedicatedEUTRA_Info_r16).Choice {
			case RRCReconfiguration_v1610_IEs_Sl_ConfigDedicatedEUTRA_Info_r16_Release:
				if err := e.EncodeNull(); err != nil {
					return err
				}
			case RRCReconfiguration_v1610_IEs_Sl_ConfigDedicatedEUTRA_Info_r16_Setup:
				if err := (*ie.Sl_ConfigDedicatedEUTRA_Info_r16).Setup.Encode(e); err != nil {
					return err
				}
			default:
				return &per.ConstraintViolationError{TypeName: "choice", Value: int64((*ie.Sl_ConfigDedicatedEUTRA_Info_r16).Choice), Constraint: "index out of range"}
			}
		}
	}

	// 13. targetCellSMTC-SCG-r16: ref
	{
		if ie.TargetCellSMTC_SCG_r16 != nil {
			if err := ie.TargetCellSMTC_SCG_r16.Encode(e); err != nil {
				return err
			}
		}
	}

	// 14. nonCriticalExtension: ref
	{
		if ie.NonCriticalExtension != nil {
			if err := ie.NonCriticalExtension.Encode(e); err != nil {
				return err
			}
		}
	}

	return nil
}

func (ie *RRCReconfiguration_v1610_IEs) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(rRCReconfigurationV1610IEsConstraints)

	// 1. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 2. otherConfig-v1610: ref
	{
		if seq.IsComponentPresent(0) {
			ie.OtherConfig_v1610 = new(OtherConfig_v1610)
			if err := ie.OtherConfig_v1610.Decode(d); err != nil {
				return err
			}
		}
	}

	// 3. bap-Config-r16: choice
	{
		if seq.IsComponentPresent(1) {
			ie.Bap_Config_r16 = &struct {
				Choice  int
				Release *struct{}
				Setup   *BAP_Config_r16
			}{}
			choiceDec := d.NewChoiceDecoder(rRCReconfiguration_v1610_IEsBapConfigR16Constraints)
			index, isExt, _, err := choiceDec.DecodeChoice()
			if err != nil {
				return err
			}
			(*ie.Bap_Config_r16).Choice = int(index)
			if isExt {
				return &per.DecodeError{TypeName: "choice", Reason: "extension alternative not supported"}
			}
			switch index {
			case RRCReconfiguration_v1610_IEs_Bap_Config_r16_Release:
				(*ie.Bap_Config_r16).Release = &struct{}{}
				if err := d.DecodeNull(); err != nil {
					return err
				}
			case RRCReconfiguration_v1610_IEs_Bap_Config_r16_Setup:
				(*ie.Bap_Config_r16).Setup = new(BAP_Config_r16)
				if err := (*ie.Bap_Config_r16).Setup.Decode(d); err != nil {
					return err
				}
			}
		}
	}

	// 4. iab-IP-AddressConfigurationList-r16: ref
	{
		if seq.IsComponentPresent(2) {
			ie.Iab_IP_AddressConfigurationList_r16 = new(IAB_IP_AddressConfigurationList_r16)
			if err := ie.Iab_IP_AddressConfigurationList_r16.Decode(d); err != nil {
				return err
			}
		}
	}

	// 5. conditionalReconfiguration-r16: ref
	{
		if seq.IsComponentPresent(3) {
			ie.ConditionalReconfiguration_r16 = new(ConditionalReconfiguration_r16)
			if err := ie.ConditionalReconfiguration_r16.Decode(d); err != nil {
				return err
			}
		}
	}

	// 6. daps-SourceRelease-r16: enumerated
	{
		if seq.IsComponentPresent(4) {
			idx, err := d.DecodeEnumerated(rRCReconfigurationV1610IEsDapsSourceReleaseR16Constraints)
			if err != nil {
				return err
			}
			ie.Daps_SourceRelease_r16 = &idx
		}
	}

	// 7. t316-r16: choice
	{
		if seq.IsComponentPresent(5) {
			ie.T316_r16 = &struct {
				Choice  int
				Release *struct{}
				Setup   *T316_r16
			}{}
			choiceDec := d.NewChoiceDecoder(rRCReconfiguration_v1610_IEsT316R16Constraints)
			index, isExt, _, err := choiceDec.DecodeChoice()
			if err != nil {
				return err
			}
			(*ie.T316_r16).Choice = int(index)
			if isExt {
				return &per.DecodeError{TypeName: "choice", Reason: "extension alternative not supported"}
			}
			switch index {
			case RRCReconfiguration_v1610_IEs_T316_r16_Release:
				(*ie.T316_r16).Release = &struct{}{}
				if err := d.DecodeNull(); err != nil {
					return err
				}
			case RRCReconfiguration_v1610_IEs_T316_r16_Setup:
				(*ie.T316_r16).Setup = new(T316_r16)
				if err := (*ie.T316_r16).Setup.Decode(d); err != nil {
					return err
				}
			}
		}
	}

	// 8. needForGapsConfigNR-r16: choice
	{
		if seq.IsComponentPresent(6) {
			ie.NeedForGapsConfigNR_r16 = &struct {
				Choice  int
				Release *struct{}
				Setup   *NeedForGapsConfigNR_r16
			}{}
			choiceDec := d.NewChoiceDecoder(rRCReconfiguration_v1610_IEsNeedForGapsConfigNRR16Constraints)
			index, isExt, _, err := choiceDec.DecodeChoice()
			if err != nil {
				return err
			}
			(*ie.NeedForGapsConfigNR_r16).Choice = int(index)
			if isExt {
				return &per.DecodeError{TypeName: "choice", Reason: "extension alternative not supported"}
			}
			switch index {
			case RRCReconfiguration_v1610_IEs_NeedForGapsConfigNR_r16_Release:
				(*ie.NeedForGapsConfigNR_r16).Release = &struct{}{}
				if err := d.DecodeNull(); err != nil {
					return err
				}
			case RRCReconfiguration_v1610_IEs_NeedForGapsConfigNR_r16_Setup:
				(*ie.NeedForGapsConfigNR_r16).Setup = new(NeedForGapsConfigNR_r16)
				if err := (*ie.NeedForGapsConfigNR_r16).Setup.Decode(d); err != nil {
					return err
				}
			}
		}
	}

	// 9. onDemandSIB-Request-r16: choice
	{
		if seq.IsComponentPresent(7) {
			ie.OnDemandSIB_Request_r16 = &struct {
				Choice  int
				Release *struct{}
				Setup   *OnDemandSIB_Request_r16
			}{}
			choiceDec := d.NewChoiceDecoder(rRCReconfiguration_v1610_IEsOnDemandSIBRequestR16Constraints)
			index, isExt, _, err := choiceDec.DecodeChoice()
			if err != nil {
				return err
			}
			(*ie.OnDemandSIB_Request_r16).Choice = int(index)
			if isExt {
				return &per.DecodeError{TypeName: "choice", Reason: "extension alternative not supported"}
			}
			switch index {
			case RRCReconfiguration_v1610_IEs_OnDemandSIB_Request_r16_Release:
				(*ie.OnDemandSIB_Request_r16).Release = &struct{}{}
				if err := d.DecodeNull(); err != nil {
					return err
				}
			case RRCReconfiguration_v1610_IEs_OnDemandSIB_Request_r16_Setup:
				(*ie.OnDemandSIB_Request_r16).Setup = new(OnDemandSIB_Request_r16)
				if err := (*ie.OnDemandSIB_Request_r16).Setup.Decode(d); err != nil {
					return err
				}
			}
		}
	}

	// 10. dedicatedPosSysInfoDelivery-r16: octet-string
	{
		if seq.IsComponentPresent(8) {
			v, err := d.DecodeOctetString(rRCReconfigurationV1610IEsDedicatedPosSysInfoDeliveryR16Constraints)
			if err != nil {
				return err
			}
			ie.DedicatedPosSysInfoDelivery_r16 = v
		}
	}

	// 11. sl-ConfigDedicatedNR-r16: choice
	{
		if seq.IsComponentPresent(9) {
			ie.Sl_ConfigDedicatedNR_r16 = &struct {
				Choice  int
				Release *struct{}
				Setup   *SL_ConfigDedicatedNR_r16
			}{}
			choiceDec := d.NewChoiceDecoder(rRCReconfiguration_v1610_IEsSlConfigDedicatedNRR16Constraints)
			index, isExt, _, err := choiceDec.DecodeChoice()
			if err != nil {
				return err
			}
			(*ie.Sl_ConfigDedicatedNR_r16).Choice = int(index)
			if isExt {
				return &per.DecodeError{TypeName: "choice", Reason: "extension alternative not supported"}
			}
			switch index {
			case RRCReconfiguration_v1610_IEs_Sl_ConfigDedicatedNR_r16_Release:
				(*ie.Sl_ConfigDedicatedNR_r16).Release = &struct{}{}
				if err := d.DecodeNull(); err != nil {
					return err
				}
			case RRCReconfiguration_v1610_IEs_Sl_ConfigDedicatedNR_r16_Setup:
				(*ie.Sl_ConfigDedicatedNR_r16).Setup = new(SL_ConfigDedicatedNR_r16)
				if err := (*ie.Sl_ConfigDedicatedNR_r16).Setup.Decode(d); err != nil {
					return err
				}
			}
		}
	}

	// 12. sl-ConfigDedicatedEUTRA-Info-r16: choice
	{
		if seq.IsComponentPresent(10) {
			ie.Sl_ConfigDedicatedEUTRA_Info_r16 = &struct {
				Choice  int
				Release *struct{}
				Setup   *SL_ConfigDedicatedEUTRA_Info_r16
			}{}
			choiceDec := d.NewChoiceDecoder(rRCReconfiguration_v1610_IEsSlConfigDedicatedEUTRAInfoR16Constraints)
			index, isExt, _, err := choiceDec.DecodeChoice()
			if err != nil {
				return err
			}
			(*ie.Sl_ConfigDedicatedEUTRA_Info_r16).Choice = int(index)
			if isExt {
				return &per.DecodeError{TypeName: "choice", Reason: "extension alternative not supported"}
			}
			switch index {
			case RRCReconfiguration_v1610_IEs_Sl_ConfigDedicatedEUTRA_Info_r16_Release:
				(*ie.Sl_ConfigDedicatedEUTRA_Info_r16).Release = &struct{}{}
				if err := d.DecodeNull(); err != nil {
					return err
				}
			case RRCReconfiguration_v1610_IEs_Sl_ConfigDedicatedEUTRA_Info_r16_Setup:
				(*ie.Sl_ConfigDedicatedEUTRA_Info_r16).Setup = new(SL_ConfigDedicatedEUTRA_Info_r16)
				if err := (*ie.Sl_ConfigDedicatedEUTRA_Info_r16).Setup.Decode(d); err != nil {
					return err
				}
			}
		}
	}

	// 13. targetCellSMTC-SCG-r16: ref
	{
		if seq.IsComponentPresent(11) {
			ie.TargetCellSMTC_SCG_r16 = new(SSB_MTC)
			if err := ie.TargetCellSMTC_SCG_r16.Decode(d); err != nil {
				return err
			}
		}
	}

	// 14. nonCriticalExtension: ref
	{
		if seq.IsComponentPresent(12) {
			ie.NonCriticalExtension = new(RRCReconfiguration_v1700_IEs)
			if err := ie.NonCriticalExtension.Decode(d); err != nil {
				return err
			}
		}
	}

	return nil
}
