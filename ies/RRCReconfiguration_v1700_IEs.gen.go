// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: RRCReconfiguration-v1700-IEs (line 1018).

var rRCReconfigurationV1700IEsConstraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "otherConfig-v1700", Optional: true},
		{Name: "sl-L2RelayUE-Config-r17", Optional: true},
		{Name: "sl-L2RemoteUE-Config-r17", Optional: true},
		{Name: "dedicatedPagingDelivery-r17", Optional: true},
		{Name: "needForGapNCSG-ConfigNR-r17", Optional: true},
		{Name: "needForGapNCSG-ConfigEUTRA-r17", Optional: true},
		{Name: "musim-GapConfig-r17", Optional: true},
		{Name: "ul-GapFR2-Config-r17", Optional: true},
		{Name: "scg-State-r17", Optional: true},
		{Name: "appLayerMeasConfig-r17", Optional: true},
		{Name: "ue-TxTEG-RequestUL-TDOA-Config-r17", Optional: true},
		{Name: "nonCriticalExtension", Optional: true},
	},
}

var rRCReconfiguration_v1700_IEsSlL2RelayUEConfigR17Constraints = per.ChoiceConstraints{
	Extensible: false,
	RootAlternatives: []per.AlternativeInfo{
		{Name: "release"},
		{Name: "setup"},
	},
}

const (
	RRCReconfiguration_v1700_IEs_Sl_L2RelayUE_Config_r17_Release = 0
	RRCReconfiguration_v1700_IEs_Sl_L2RelayUE_Config_r17_Setup   = 1
)

var rRCReconfiguration_v1700_IEsSlL2RemoteUEConfigR17Constraints = per.ChoiceConstraints{
	Extensible: false,
	RootAlternatives: []per.AlternativeInfo{
		{Name: "release"},
		{Name: "setup"},
	},
}

const (
	RRCReconfiguration_v1700_IEs_Sl_L2RemoteUE_Config_r17_Release = 0
	RRCReconfiguration_v1700_IEs_Sl_L2RemoteUE_Config_r17_Setup   = 1
)

var rRCReconfigurationV1700IEsDedicatedPagingDeliveryR17Constraints = per.SizeConstraints{}

var rRCReconfiguration_v1700_IEsNeedForGapNCSGConfigNRR17Constraints = per.ChoiceConstraints{
	Extensible: false,
	RootAlternatives: []per.AlternativeInfo{
		{Name: "release"},
		{Name: "setup"},
	},
}

const (
	RRCReconfiguration_v1700_IEs_NeedForGapNCSG_ConfigNR_r17_Release = 0
	RRCReconfiguration_v1700_IEs_NeedForGapNCSG_ConfigNR_r17_Setup   = 1
)

var rRCReconfiguration_v1700_IEsNeedForGapNCSGConfigEUTRAR17Constraints = per.ChoiceConstraints{
	Extensible: false,
	RootAlternatives: []per.AlternativeInfo{
		{Name: "release"},
		{Name: "setup"},
	},
}

const (
	RRCReconfiguration_v1700_IEs_NeedForGapNCSG_ConfigEUTRA_r17_Release = 0
	RRCReconfiguration_v1700_IEs_NeedForGapNCSG_ConfigEUTRA_r17_Setup   = 1
)

var rRCReconfiguration_v1700_IEsMusimGapConfigR17Constraints = per.ChoiceConstraints{
	Extensible: false,
	RootAlternatives: []per.AlternativeInfo{
		{Name: "release"},
		{Name: "setup"},
	},
}

const (
	RRCReconfiguration_v1700_IEs_Musim_GapConfig_r17_Release = 0
	RRCReconfiguration_v1700_IEs_Musim_GapConfig_r17_Setup   = 1
)

var rRCReconfiguration_v1700_IEsUlGapFR2ConfigR17Constraints = per.ChoiceConstraints{
	Extensible: false,
	RootAlternatives: []per.AlternativeInfo{
		{Name: "release"},
		{Name: "setup"},
	},
}

const (
	RRCReconfiguration_v1700_IEs_Ul_GapFR2_Config_r17_Release = 0
	RRCReconfiguration_v1700_IEs_Ul_GapFR2_Config_r17_Setup   = 1
)

const (
	RRCReconfiguration_v1700_IEs_Scg_State_r17_Deactivated = 0
)

var rRCReconfigurationV1700IEsScgStateR17Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{RRCReconfiguration_v1700_IEs_Scg_State_r17_Deactivated},
}

var rRCReconfiguration_v1700_IEsUeTxTEGRequestULTDOAConfigR17Constraints = per.ChoiceConstraints{
	Extensible: false,
	RootAlternatives: []per.AlternativeInfo{
		{Name: "release"},
		{Name: "setup"},
	},
}

const (
	RRCReconfiguration_v1700_IEs_Ue_TxTEG_RequestUL_TDOA_Config_r17_Release = 0
	RRCReconfiguration_v1700_IEs_Ue_TxTEG_RequestUL_TDOA_Config_r17_Setup   = 1
)

type RRCReconfiguration_v1700_IEs struct {
	OtherConfig_v1700       *OtherConfig_v1700
	Sl_L2RelayUE_Config_r17 *struct {
		Choice  int
		Release *struct{}
		Setup   *SL_L2RelayUE_Config_r17
	}
	Sl_L2RemoteUE_Config_r17 *struct {
		Choice  int
		Release *struct{}
		Setup   *SL_L2RemoteUE_Config_r17
	}
	DedicatedPagingDelivery_r17 []byte
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
	Musim_GapConfig_r17 *struct {
		Choice  int
		Release *struct{}
		Setup   *MUSIM_GapConfig_r17
	}
	Ul_GapFR2_Config_r17 *struct {
		Choice  int
		Release *struct{}
		Setup   *UL_GapFR2_Config_r17
	}
	Scg_State_r17                      *int64
	AppLayerMeasConfig_r17             *AppLayerMeasConfig_r17
	Ue_TxTEG_RequestUL_TDOA_Config_r17 *struct {
		Choice  int
		Release *struct{}
		Setup   *UE_TxTEG_RequestUL_TDOA_Config_r17
	}
	NonCriticalExtension *RRCReconfiguration_v1800_IEs
}

func (ie *RRCReconfiguration_v1700_IEs) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(rRCReconfigurationV1700IEsConstraints)

	// 1. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.OtherConfig_v1700 != nil, ie.Sl_L2RelayUE_Config_r17 != nil, ie.Sl_L2RemoteUE_Config_r17 != nil, ie.DedicatedPagingDelivery_r17 != nil, ie.NeedForGapNCSG_ConfigNR_r17 != nil, ie.NeedForGapNCSG_ConfigEUTRA_r17 != nil, ie.Musim_GapConfig_r17 != nil, ie.Ul_GapFR2_Config_r17 != nil, ie.Scg_State_r17 != nil, ie.AppLayerMeasConfig_r17 != nil, ie.Ue_TxTEG_RequestUL_TDOA_Config_r17 != nil, ie.NonCriticalExtension != nil}); err != nil {
		return err
	}

	// 2. otherConfig-v1700: ref
	{
		if ie.OtherConfig_v1700 != nil {
			if err := ie.OtherConfig_v1700.Encode(e); err != nil {
				return err
			}
		}
	}

	// 3. sl-L2RelayUE-Config-r17: choice
	{
		if ie.Sl_L2RelayUE_Config_r17 != nil {
			choiceEnc := e.NewChoiceEncoder(rRCReconfiguration_v1700_IEsSlL2RelayUEConfigR17Constraints)
			if err := choiceEnc.EncodeChoice(int64((*ie.Sl_L2RelayUE_Config_r17).Choice), false, nil); err != nil {
				return err
			}
			switch (*ie.Sl_L2RelayUE_Config_r17).Choice {
			case RRCReconfiguration_v1700_IEs_Sl_L2RelayUE_Config_r17_Release:
				if err := e.EncodeNull(); err != nil {
					return err
				}
			case RRCReconfiguration_v1700_IEs_Sl_L2RelayUE_Config_r17_Setup:
				if err := (*ie.Sl_L2RelayUE_Config_r17).Setup.Encode(e); err != nil {
					return err
				}
			default:
				return &per.ConstraintViolationError{TypeName: "choice", Value: int64((*ie.Sl_L2RelayUE_Config_r17).Choice), Constraint: "index out of range"}
			}
		}
	}

	// 4. sl-L2RemoteUE-Config-r17: choice
	{
		if ie.Sl_L2RemoteUE_Config_r17 != nil {
			choiceEnc := e.NewChoiceEncoder(rRCReconfiguration_v1700_IEsSlL2RemoteUEConfigR17Constraints)
			if err := choiceEnc.EncodeChoice(int64((*ie.Sl_L2RemoteUE_Config_r17).Choice), false, nil); err != nil {
				return err
			}
			switch (*ie.Sl_L2RemoteUE_Config_r17).Choice {
			case RRCReconfiguration_v1700_IEs_Sl_L2RemoteUE_Config_r17_Release:
				if err := e.EncodeNull(); err != nil {
					return err
				}
			case RRCReconfiguration_v1700_IEs_Sl_L2RemoteUE_Config_r17_Setup:
				if err := (*ie.Sl_L2RemoteUE_Config_r17).Setup.Encode(e); err != nil {
					return err
				}
			default:
				return &per.ConstraintViolationError{TypeName: "choice", Value: int64((*ie.Sl_L2RemoteUE_Config_r17).Choice), Constraint: "index out of range"}
			}
		}
	}

	// 5. dedicatedPagingDelivery-r17: octet-string
	{
		if ie.DedicatedPagingDelivery_r17 != nil {
			if err := e.EncodeOctetString(ie.DedicatedPagingDelivery_r17, rRCReconfigurationV1700IEsDedicatedPagingDeliveryR17Constraints); err != nil {
				return err
			}
		}
	}

	// 6. needForGapNCSG-ConfigNR-r17: choice
	{
		if ie.NeedForGapNCSG_ConfigNR_r17 != nil {
			choiceEnc := e.NewChoiceEncoder(rRCReconfiguration_v1700_IEsNeedForGapNCSGConfigNRR17Constraints)
			if err := choiceEnc.EncodeChoice(int64((*ie.NeedForGapNCSG_ConfigNR_r17).Choice), false, nil); err != nil {
				return err
			}
			switch (*ie.NeedForGapNCSG_ConfigNR_r17).Choice {
			case RRCReconfiguration_v1700_IEs_NeedForGapNCSG_ConfigNR_r17_Release:
				if err := e.EncodeNull(); err != nil {
					return err
				}
			case RRCReconfiguration_v1700_IEs_NeedForGapNCSG_ConfigNR_r17_Setup:
				if err := (*ie.NeedForGapNCSG_ConfigNR_r17).Setup.Encode(e); err != nil {
					return err
				}
			default:
				return &per.ConstraintViolationError{TypeName: "choice", Value: int64((*ie.NeedForGapNCSG_ConfigNR_r17).Choice), Constraint: "index out of range"}
			}
		}
	}

	// 7. needForGapNCSG-ConfigEUTRA-r17: choice
	{
		if ie.NeedForGapNCSG_ConfigEUTRA_r17 != nil {
			choiceEnc := e.NewChoiceEncoder(rRCReconfiguration_v1700_IEsNeedForGapNCSGConfigEUTRAR17Constraints)
			if err := choiceEnc.EncodeChoice(int64((*ie.NeedForGapNCSG_ConfigEUTRA_r17).Choice), false, nil); err != nil {
				return err
			}
			switch (*ie.NeedForGapNCSG_ConfigEUTRA_r17).Choice {
			case RRCReconfiguration_v1700_IEs_NeedForGapNCSG_ConfigEUTRA_r17_Release:
				if err := e.EncodeNull(); err != nil {
					return err
				}
			case RRCReconfiguration_v1700_IEs_NeedForGapNCSG_ConfigEUTRA_r17_Setup:
				if err := (*ie.NeedForGapNCSG_ConfigEUTRA_r17).Setup.Encode(e); err != nil {
					return err
				}
			default:
				return &per.ConstraintViolationError{TypeName: "choice", Value: int64((*ie.NeedForGapNCSG_ConfigEUTRA_r17).Choice), Constraint: "index out of range"}
			}
		}
	}

	// 8. musim-GapConfig-r17: choice
	{
		if ie.Musim_GapConfig_r17 != nil {
			choiceEnc := e.NewChoiceEncoder(rRCReconfiguration_v1700_IEsMusimGapConfigR17Constraints)
			if err := choiceEnc.EncodeChoice(int64((*ie.Musim_GapConfig_r17).Choice), false, nil); err != nil {
				return err
			}
			switch (*ie.Musim_GapConfig_r17).Choice {
			case RRCReconfiguration_v1700_IEs_Musim_GapConfig_r17_Release:
				if err := e.EncodeNull(); err != nil {
					return err
				}
			case RRCReconfiguration_v1700_IEs_Musim_GapConfig_r17_Setup:
				if err := (*ie.Musim_GapConfig_r17).Setup.Encode(e); err != nil {
					return err
				}
			default:
				return &per.ConstraintViolationError{TypeName: "choice", Value: int64((*ie.Musim_GapConfig_r17).Choice), Constraint: "index out of range"}
			}
		}
	}

	// 9. ul-GapFR2-Config-r17: choice
	{
		if ie.Ul_GapFR2_Config_r17 != nil {
			choiceEnc := e.NewChoiceEncoder(rRCReconfiguration_v1700_IEsUlGapFR2ConfigR17Constraints)
			if err := choiceEnc.EncodeChoice(int64((*ie.Ul_GapFR2_Config_r17).Choice), false, nil); err != nil {
				return err
			}
			switch (*ie.Ul_GapFR2_Config_r17).Choice {
			case RRCReconfiguration_v1700_IEs_Ul_GapFR2_Config_r17_Release:
				if err := e.EncodeNull(); err != nil {
					return err
				}
			case RRCReconfiguration_v1700_IEs_Ul_GapFR2_Config_r17_Setup:
				if err := (*ie.Ul_GapFR2_Config_r17).Setup.Encode(e); err != nil {
					return err
				}
			default:
				return &per.ConstraintViolationError{TypeName: "choice", Value: int64((*ie.Ul_GapFR2_Config_r17).Choice), Constraint: "index out of range"}
			}
		}
	}

	// 10. scg-State-r17: enumerated
	{
		if ie.Scg_State_r17 != nil {
			if err := e.EncodeEnumerated(*ie.Scg_State_r17, rRCReconfigurationV1700IEsScgStateR17Constraints); err != nil {
				return err
			}
		}
	}

	// 11. appLayerMeasConfig-r17: ref
	{
		if ie.AppLayerMeasConfig_r17 != nil {
			if err := ie.AppLayerMeasConfig_r17.Encode(e); err != nil {
				return err
			}
		}
	}

	// 12. ue-TxTEG-RequestUL-TDOA-Config-r17: choice
	{
		if ie.Ue_TxTEG_RequestUL_TDOA_Config_r17 != nil {
			choiceEnc := e.NewChoiceEncoder(rRCReconfiguration_v1700_IEsUeTxTEGRequestULTDOAConfigR17Constraints)
			if err := choiceEnc.EncodeChoice(int64((*ie.Ue_TxTEG_RequestUL_TDOA_Config_r17).Choice), false, nil); err != nil {
				return err
			}
			switch (*ie.Ue_TxTEG_RequestUL_TDOA_Config_r17).Choice {
			case RRCReconfiguration_v1700_IEs_Ue_TxTEG_RequestUL_TDOA_Config_r17_Release:
				if err := e.EncodeNull(); err != nil {
					return err
				}
			case RRCReconfiguration_v1700_IEs_Ue_TxTEG_RequestUL_TDOA_Config_r17_Setup:
				if err := (*ie.Ue_TxTEG_RequestUL_TDOA_Config_r17).Setup.Encode(e); err != nil {
					return err
				}
			default:
				return &per.ConstraintViolationError{TypeName: "choice", Value: int64((*ie.Ue_TxTEG_RequestUL_TDOA_Config_r17).Choice), Constraint: "index out of range"}
			}
		}
	}

	// 13. nonCriticalExtension: ref
	{
		if ie.NonCriticalExtension != nil {
			if err := ie.NonCriticalExtension.Encode(e); err != nil {
				return err
			}
		}
	}

	return nil
}

func (ie *RRCReconfiguration_v1700_IEs) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(rRCReconfigurationV1700IEsConstraints)

	// 1. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 2. otherConfig-v1700: ref
	{
		if seq.IsComponentPresent(0) {
			ie.OtherConfig_v1700 = new(OtherConfig_v1700)
			if err := ie.OtherConfig_v1700.Decode(d); err != nil {
				return err
			}
		}
	}

	// 3. sl-L2RelayUE-Config-r17: choice
	{
		if seq.IsComponentPresent(1) {
			ie.Sl_L2RelayUE_Config_r17 = &struct {
				Choice  int
				Release *struct{}
				Setup   *SL_L2RelayUE_Config_r17
			}{}
			choiceDec := d.NewChoiceDecoder(rRCReconfiguration_v1700_IEsSlL2RelayUEConfigR17Constraints)
			index, isExt, _, err := choiceDec.DecodeChoice()
			if err != nil {
				return err
			}
			(*ie.Sl_L2RelayUE_Config_r17).Choice = int(index)
			if isExt {
				return &per.DecodeError{TypeName: "choice", Reason: "extension alternative not supported"}
			}
			switch index {
			case RRCReconfiguration_v1700_IEs_Sl_L2RelayUE_Config_r17_Release:
				(*ie.Sl_L2RelayUE_Config_r17).Release = &struct{}{}
				if err := d.DecodeNull(); err != nil {
					return err
				}
			case RRCReconfiguration_v1700_IEs_Sl_L2RelayUE_Config_r17_Setup:
				(*ie.Sl_L2RelayUE_Config_r17).Setup = new(SL_L2RelayUE_Config_r17)
				if err := (*ie.Sl_L2RelayUE_Config_r17).Setup.Decode(d); err != nil {
					return err
				}
			}
		}
	}

	// 4. sl-L2RemoteUE-Config-r17: choice
	{
		if seq.IsComponentPresent(2) {
			ie.Sl_L2RemoteUE_Config_r17 = &struct {
				Choice  int
				Release *struct{}
				Setup   *SL_L2RemoteUE_Config_r17
			}{}
			choiceDec := d.NewChoiceDecoder(rRCReconfiguration_v1700_IEsSlL2RemoteUEConfigR17Constraints)
			index, isExt, _, err := choiceDec.DecodeChoice()
			if err != nil {
				return err
			}
			(*ie.Sl_L2RemoteUE_Config_r17).Choice = int(index)
			if isExt {
				return &per.DecodeError{TypeName: "choice", Reason: "extension alternative not supported"}
			}
			switch index {
			case RRCReconfiguration_v1700_IEs_Sl_L2RemoteUE_Config_r17_Release:
				(*ie.Sl_L2RemoteUE_Config_r17).Release = &struct{}{}
				if err := d.DecodeNull(); err != nil {
					return err
				}
			case RRCReconfiguration_v1700_IEs_Sl_L2RemoteUE_Config_r17_Setup:
				(*ie.Sl_L2RemoteUE_Config_r17).Setup = new(SL_L2RemoteUE_Config_r17)
				if err := (*ie.Sl_L2RemoteUE_Config_r17).Setup.Decode(d); err != nil {
					return err
				}
			}
		}
	}

	// 5. dedicatedPagingDelivery-r17: octet-string
	{
		if seq.IsComponentPresent(3) {
			v, err := d.DecodeOctetString(rRCReconfigurationV1700IEsDedicatedPagingDeliveryR17Constraints)
			if err != nil {
				return err
			}
			ie.DedicatedPagingDelivery_r17 = v
		}
	}

	// 6. needForGapNCSG-ConfigNR-r17: choice
	{
		if seq.IsComponentPresent(4) {
			ie.NeedForGapNCSG_ConfigNR_r17 = &struct {
				Choice  int
				Release *struct{}
				Setup   *NeedForGapNCSG_ConfigNR_r17
			}{}
			choiceDec := d.NewChoiceDecoder(rRCReconfiguration_v1700_IEsNeedForGapNCSGConfigNRR17Constraints)
			index, isExt, _, err := choiceDec.DecodeChoice()
			if err != nil {
				return err
			}
			(*ie.NeedForGapNCSG_ConfigNR_r17).Choice = int(index)
			if isExt {
				return &per.DecodeError{TypeName: "choice", Reason: "extension alternative not supported"}
			}
			switch index {
			case RRCReconfiguration_v1700_IEs_NeedForGapNCSG_ConfigNR_r17_Release:
				(*ie.NeedForGapNCSG_ConfigNR_r17).Release = &struct{}{}
				if err := d.DecodeNull(); err != nil {
					return err
				}
			case RRCReconfiguration_v1700_IEs_NeedForGapNCSG_ConfigNR_r17_Setup:
				(*ie.NeedForGapNCSG_ConfigNR_r17).Setup = new(NeedForGapNCSG_ConfigNR_r17)
				if err := (*ie.NeedForGapNCSG_ConfigNR_r17).Setup.Decode(d); err != nil {
					return err
				}
			}
		}
	}

	// 7. needForGapNCSG-ConfigEUTRA-r17: choice
	{
		if seq.IsComponentPresent(5) {
			ie.NeedForGapNCSG_ConfigEUTRA_r17 = &struct {
				Choice  int
				Release *struct{}
				Setup   *NeedForGapNCSG_ConfigEUTRA_r17
			}{}
			choiceDec := d.NewChoiceDecoder(rRCReconfiguration_v1700_IEsNeedForGapNCSGConfigEUTRAR17Constraints)
			index, isExt, _, err := choiceDec.DecodeChoice()
			if err != nil {
				return err
			}
			(*ie.NeedForGapNCSG_ConfigEUTRA_r17).Choice = int(index)
			if isExt {
				return &per.DecodeError{TypeName: "choice", Reason: "extension alternative not supported"}
			}
			switch index {
			case RRCReconfiguration_v1700_IEs_NeedForGapNCSG_ConfigEUTRA_r17_Release:
				(*ie.NeedForGapNCSG_ConfigEUTRA_r17).Release = &struct{}{}
				if err := d.DecodeNull(); err != nil {
					return err
				}
			case RRCReconfiguration_v1700_IEs_NeedForGapNCSG_ConfigEUTRA_r17_Setup:
				(*ie.NeedForGapNCSG_ConfigEUTRA_r17).Setup = new(NeedForGapNCSG_ConfigEUTRA_r17)
				if err := (*ie.NeedForGapNCSG_ConfigEUTRA_r17).Setup.Decode(d); err != nil {
					return err
				}
			}
		}
	}

	// 8. musim-GapConfig-r17: choice
	{
		if seq.IsComponentPresent(6) {
			ie.Musim_GapConfig_r17 = &struct {
				Choice  int
				Release *struct{}
				Setup   *MUSIM_GapConfig_r17
			}{}
			choiceDec := d.NewChoiceDecoder(rRCReconfiguration_v1700_IEsMusimGapConfigR17Constraints)
			index, isExt, _, err := choiceDec.DecodeChoice()
			if err != nil {
				return err
			}
			(*ie.Musim_GapConfig_r17).Choice = int(index)
			if isExt {
				return &per.DecodeError{TypeName: "choice", Reason: "extension alternative not supported"}
			}
			switch index {
			case RRCReconfiguration_v1700_IEs_Musim_GapConfig_r17_Release:
				(*ie.Musim_GapConfig_r17).Release = &struct{}{}
				if err := d.DecodeNull(); err != nil {
					return err
				}
			case RRCReconfiguration_v1700_IEs_Musim_GapConfig_r17_Setup:
				(*ie.Musim_GapConfig_r17).Setup = new(MUSIM_GapConfig_r17)
				if err := (*ie.Musim_GapConfig_r17).Setup.Decode(d); err != nil {
					return err
				}
			}
		}
	}

	// 9. ul-GapFR2-Config-r17: choice
	{
		if seq.IsComponentPresent(7) {
			ie.Ul_GapFR2_Config_r17 = &struct {
				Choice  int
				Release *struct{}
				Setup   *UL_GapFR2_Config_r17
			}{}
			choiceDec := d.NewChoiceDecoder(rRCReconfiguration_v1700_IEsUlGapFR2ConfigR17Constraints)
			index, isExt, _, err := choiceDec.DecodeChoice()
			if err != nil {
				return err
			}
			(*ie.Ul_GapFR2_Config_r17).Choice = int(index)
			if isExt {
				return &per.DecodeError{TypeName: "choice", Reason: "extension alternative not supported"}
			}
			switch index {
			case RRCReconfiguration_v1700_IEs_Ul_GapFR2_Config_r17_Release:
				(*ie.Ul_GapFR2_Config_r17).Release = &struct{}{}
				if err := d.DecodeNull(); err != nil {
					return err
				}
			case RRCReconfiguration_v1700_IEs_Ul_GapFR2_Config_r17_Setup:
				(*ie.Ul_GapFR2_Config_r17).Setup = new(UL_GapFR2_Config_r17)
				if err := (*ie.Ul_GapFR2_Config_r17).Setup.Decode(d); err != nil {
					return err
				}
			}
		}
	}

	// 10. scg-State-r17: enumerated
	{
		if seq.IsComponentPresent(8) {
			idx, err := d.DecodeEnumerated(rRCReconfigurationV1700IEsScgStateR17Constraints)
			if err != nil {
				return err
			}
			ie.Scg_State_r17 = &idx
		}
	}

	// 11. appLayerMeasConfig-r17: ref
	{
		if seq.IsComponentPresent(9) {
			ie.AppLayerMeasConfig_r17 = new(AppLayerMeasConfig_r17)
			if err := ie.AppLayerMeasConfig_r17.Decode(d); err != nil {
				return err
			}
		}
	}

	// 12. ue-TxTEG-RequestUL-TDOA-Config-r17: choice
	{
		if seq.IsComponentPresent(10) {
			ie.Ue_TxTEG_RequestUL_TDOA_Config_r17 = &struct {
				Choice  int
				Release *struct{}
				Setup   *UE_TxTEG_RequestUL_TDOA_Config_r17
			}{}
			choiceDec := d.NewChoiceDecoder(rRCReconfiguration_v1700_IEsUeTxTEGRequestULTDOAConfigR17Constraints)
			index, isExt, _, err := choiceDec.DecodeChoice()
			if err != nil {
				return err
			}
			(*ie.Ue_TxTEG_RequestUL_TDOA_Config_r17).Choice = int(index)
			if isExt {
				return &per.DecodeError{TypeName: "choice", Reason: "extension alternative not supported"}
			}
			switch index {
			case RRCReconfiguration_v1700_IEs_Ue_TxTEG_RequestUL_TDOA_Config_r17_Release:
				(*ie.Ue_TxTEG_RequestUL_TDOA_Config_r17).Release = &struct{}{}
				if err := d.DecodeNull(); err != nil {
					return err
				}
			case RRCReconfiguration_v1700_IEs_Ue_TxTEG_RequestUL_TDOA_Config_r17_Setup:
				(*ie.Ue_TxTEG_RequestUL_TDOA_Config_r17).Setup = new(UE_TxTEG_RequestUL_TDOA_Config_r17)
				if err := (*ie.Ue_TxTEG_RequestUL_TDOA_Config_r17).Setup.Decode(d); err != nil {
					return err
				}
			}
		}
	}

	// 13. nonCriticalExtension: ref
	{
		if seq.IsComponentPresent(11) {
			ie.NonCriticalExtension = new(RRCReconfiguration_v1800_IEs)
			if err := ie.NonCriticalExtension.Decode(d); err != nil {
				return err
			}
		}
	}

	return nil
}
