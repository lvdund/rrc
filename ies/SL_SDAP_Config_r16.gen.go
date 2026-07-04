// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/rrc/common"
)

// Source: SL-SDAP-Config-r16 (line 28230).

var sLSDAPConfigR16Constraints = per.SequenceConstraints{
	Extensible: true,
	RootComponents: []per.ComponentInfo{
		{Name: "sl-SDAP-Header-r16"},
		{Name: "sl-DefaultRB-r16"},
		{Name: "sl-MappedQoS-Flows-r16", Optional: true},
		{Name: "sl-CastType-r16", Optional: true},
	},
}

const (
	SL_SDAP_Config_r16_Sl_SDAP_Header_r16_Present = 0
	SL_SDAP_Config_r16_Sl_SDAP_Header_r16_Absent  = 1
)

var sLSDAPConfigR16SlSDAPHeaderR16Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{SL_SDAP_Config_r16_Sl_SDAP_Header_r16_Present, SL_SDAP_Config_r16_Sl_SDAP_Header_r16_Absent},
}

var sL_SDAP_Config_r16SlMappedQoSFlowsR16Constraints = per.ChoiceConstraints{
	Extensible: false,
	RootAlternatives: []per.AlternativeInfo{
		{Name: "sl-MappedQoS-FlowsList-r16"},
		{Name: "sl-MappedQoS-FlowsListDedicated-r16"},
	},
}

const (
	SL_SDAP_Config_r16_Sl_MappedQoS_Flows_r16_Sl_MappedQoS_FlowsList_r16          = 0
	SL_SDAP_Config_r16_Sl_MappedQoS_Flows_r16_Sl_MappedQoS_FlowsListDedicated_r16 = 1
)

const (
	SL_SDAP_Config_r16_Sl_CastType_r16_Broadcast = 0
	SL_SDAP_Config_r16_Sl_CastType_r16_Groupcast = 1
	SL_SDAP_Config_r16_Sl_CastType_r16_Unicast   = 2
	SL_SDAP_Config_r16_Sl_CastType_r16_Spare1    = 3
)

var sLSDAPConfigR16SlCastTypeR16Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{SL_SDAP_Config_r16_Sl_CastType_r16_Broadcast, SL_SDAP_Config_r16_Sl_CastType_r16_Groupcast, SL_SDAP_Config_r16_Sl_CastType_r16_Unicast, SL_SDAP_Config_r16_Sl_CastType_r16_Spare1},
}

var sLSDAPConfigR16SlMappedQoSFlowsR16SlMappedQoSFlowsListR16Constraints = per.SizeRange(1, common.MaxNrofSL_QFIs_r16)

type SL_SDAP_Config_r16 struct {
	Sl_SDAP_Header_r16     int64
	Sl_DefaultRB_r16       bool
	Sl_MappedQoS_Flows_r16 *struct {
		Choice                              int
		Sl_MappedQoS_FlowsList_r16          *[]SL_QoS_Profile_r16
		Sl_MappedQoS_FlowsListDedicated_r16 *SL_MappedQoS_FlowsListDedicated_r16
	}
	Sl_CastType_r16 *int64
}

func (ie *SL_SDAP_Config_r16) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(sLSDAPConfigR16Constraints)

	// 1. SEQUENCE extension bit.
	if err := seq.EncodeExtensionBit(false); err != nil {
		return err
	}

	// 2. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.Sl_MappedQoS_Flows_r16 != nil, ie.Sl_CastType_r16 != nil}); err != nil {
		return err
	}

	// 3. sl-SDAP-Header-r16: enumerated
	{
		if err := e.EncodeEnumerated(ie.Sl_SDAP_Header_r16, sLSDAPConfigR16SlSDAPHeaderR16Constraints); err != nil {
			return err
		}
	}

	// 4. sl-DefaultRB-r16: boolean
	{
		if err := e.EncodeBoolean(ie.Sl_DefaultRB_r16); err != nil {
			return err
		}
	}

	// 5. sl-MappedQoS-Flows-r16: choice
	{
		if ie.Sl_MappedQoS_Flows_r16 != nil {
			choiceEnc := e.NewChoiceEncoder(sL_SDAP_Config_r16SlMappedQoSFlowsR16Constraints)
			if err := choiceEnc.EncodeChoice(int64((*ie.Sl_MappedQoS_Flows_r16).Choice), false, nil); err != nil {
				return err
			}
			switch (*ie.Sl_MappedQoS_Flows_r16).Choice {
			case SL_SDAP_Config_r16_Sl_MappedQoS_Flows_r16_Sl_MappedQoS_FlowsList_r16:
				seqOf := e.NewSequenceOfEncoder(sLSDAPConfigR16SlMappedQoSFlowsR16SlMappedQoSFlowsListR16Constraints)
				if err := seqOf.EncodeLength(int64(len((*(*ie.Sl_MappedQoS_Flows_r16).Sl_MappedQoS_FlowsList_r16)))); err != nil {
					return err
				}
				for i := range *(*ie.Sl_MappedQoS_Flows_r16).Sl_MappedQoS_FlowsList_r16 {
					if err := (*(*ie.Sl_MappedQoS_Flows_r16).Sl_MappedQoS_FlowsList_r16)[i].Encode(e); err != nil {
						return err
					}
				}
			case SL_SDAP_Config_r16_Sl_MappedQoS_Flows_r16_Sl_MappedQoS_FlowsListDedicated_r16:
				if err := (*ie.Sl_MappedQoS_Flows_r16).Sl_MappedQoS_FlowsListDedicated_r16.Encode(e); err != nil {
					return err
				}
			default:
				return &per.ConstraintViolationError{TypeName: "choice", Value: int64((*ie.Sl_MappedQoS_Flows_r16).Choice), Constraint: "index out of range"}
			}
		}
	}

	// 6. sl-CastType-r16: enumerated
	{
		if ie.Sl_CastType_r16 != nil {
			if err := e.EncodeEnumerated(*ie.Sl_CastType_r16, sLSDAPConfigR16SlCastTypeR16Constraints); err != nil {
				return err
			}
		}
	}

	return nil
}

func (ie *SL_SDAP_Config_r16) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(sLSDAPConfigR16Constraints)

	// 1. SEQUENCE extension bit.
	if err := seq.DecodeExtensionBit(); err != nil {
		return err
	}

	// 2. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 3. sl-SDAP-Header-r16: enumerated
	{
		v0, err := d.DecodeEnumerated(sLSDAPConfigR16SlSDAPHeaderR16Constraints)
		if err != nil {
			return err
		}
		ie.Sl_SDAP_Header_r16 = v0
	}

	// 4. sl-DefaultRB-r16: boolean
	{
		v1, err := d.DecodeBoolean()
		if err != nil {
			return err
		}
		ie.Sl_DefaultRB_r16 = v1
	}

	// 5. sl-MappedQoS-Flows-r16: choice
	{
		if seq.IsComponentPresent(2) {
			ie.Sl_MappedQoS_Flows_r16 = &struct {
				Choice                              int
				Sl_MappedQoS_FlowsList_r16          *[]SL_QoS_Profile_r16
				Sl_MappedQoS_FlowsListDedicated_r16 *SL_MappedQoS_FlowsListDedicated_r16
			}{}
			choiceDec := d.NewChoiceDecoder(sL_SDAP_Config_r16SlMappedQoSFlowsR16Constraints)
			index, isExt, _, err := choiceDec.DecodeChoice()
			if err != nil {
				return err
			}
			(*ie.Sl_MappedQoS_Flows_r16).Choice = int(index)
			if isExt {
				return &per.DecodeError{TypeName: "choice", Reason: "extension alternative not supported"}
			}
			switch index {
			case SL_SDAP_Config_r16_Sl_MappedQoS_Flows_r16_Sl_MappedQoS_FlowsList_r16:
				(*ie.Sl_MappedQoS_Flows_r16).Sl_MappedQoS_FlowsList_r16 = new([]SL_QoS_Profile_r16)
				seqOf := d.NewSequenceOfDecoder(sLSDAPConfigR16SlMappedQoSFlowsR16SlMappedQoSFlowsListR16Constraints)
				n, err := seqOf.DecodeLength()
				if err != nil {
					return err
				}
				(*(*ie.Sl_MappedQoS_Flows_r16).Sl_MappedQoS_FlowsList_r16) = make([]SL_QoS_Profile_r16, n)
				for i := int64(0); i < n; i++ {
					if err := (*(*ie.Sl_MappedQoS_Flows_r16).Sl_MappedQoS_FlowsList_r16)[i].Decode(d); err != nil {
						return err
					}
				}
			case SL_SDAP_Config_r16_Sl_MappedQoS_Flows_r16_Sl_MappedQoS_FlowsListDedicated_r16:
				(*ie.Sl_MappedQoS_Flows_r16).Sl_MappedQoS_FlowsListDedicated_r16 = new(SL_MappedQoS_FlowsListDedicated_r16)
				if err := (*ie.Sl_MappedQoS_Flows_r16).Sl_MappedQoS_FlowsListDedicated_r16.Decode(d); err != nil {
					return err
				}
			}
		}
	}

	// 6. sl-CastType-r16: enumerated
	{
		if seq.IsComponentPresent(3) {
			idx, err := d.DecodeEnumerated(sLSDAPConfigR16SlCastTypeR16Constraints)
			if err != nil {
				return err
			}
			ie.Sl_CastType_r16 = &idx
		}
	}

	return nil
}
