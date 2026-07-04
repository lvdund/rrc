// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/rrc/common"
)

// Source: SL-U2U-Info-r18 (line 2251).

var sLU2UInfoR18Constraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "sl-U2U-Identity-r18"},
		{Name: "sl-E2E-QoS-InfoList-r18", Optional: true},
		{Name: "sl-PerHop-QoS-InfoList-r18", Optional: true},
		{Name: "sl-PerSLRB-QoS-InfoList-r18", Optional: true},
		{Name: "sl-CapabilityInformationTargetRemoteUE-r18", Optional: true},
	},
}

var sL_U2U_Info_r18SlU2UIdentityR18Constraints = per.ChoiceConstraints{
	Extensible: false,
	RootAlternatives: []per.AlternativeInfo{
		{Name: "sl-TargetUE-Identity-r18"},
		{Name: "sl-SourceUE-Identity-r18"},
	},
}

const (
	SL_U2U_Info_r18_Sl_U2U_Identity_r18_Sl_TargetUE_Identity_r18 = 0
	SL_U2U_Info_r18_Sl_U2U_Identity_r18_Sl_SourceUE_Identity_r18 = 1
)

var sLU2UInfoR18SlE2EQoSInfoListR18Constraints = per.SizeRange(1, common.MaxNrofSL_QFIsPerDest_r16)

var sLU2UInfoR18SlPerHopQoSInfoListR18Constraints = per.SizeRange(1, common.MaxNrofSL_QFIsPerDest_r16)

var sLU2UInfoR18SlPerSLRBQoSInfoListR18Constraints = per.SizeRange(1, common.MaxNrofSLRB_r16)

var sLU2UInfoR18SlCapabilityInformationTargetRemoteUER18Constraints = per.SizeConstraints{}

type SL_U2U_Info_r18 struct {
	Sl_U2U_Identity_r18 struct {
		Choice                   int
		Sl_TargetUE_Identity_r18 *SL_DestinationIdentity_r16
		Sl_SourceUE_Identity_r18 *SL_SourceIdentity_r17
	}
	Sl_E2E_QoS_InfoList_r18                    []SL_QoS_Info_r16
	Sl_PerHop_QoS_InfoList_r18                 []SL_SplitQoS_Info_r18
	Sl_PerSLRB_QoS_InfoList_r18                []SL_PerSLRB_QoS_Info_r18
	Sl_CapabilityInformationTargetRemoteUE_r18 []byte
}

func (ie *SL_U2U_Info_r18) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(sLU2UInfoR18Constraints)

	// 1. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.Sl_E2E_QoS_InfoList_r18 != nil, ie.Sl_PerHop_QoS_InfoList_r18 != nil, ie.Sl_PerSLRB_QoS_InfoList_r18 != nil, ie.Sl_CapabilityInformationTargetRemoteUE_r18 != nil}); err != nil {
		return err
	}

	// 2. sl-U2U-Identity-r18: choice
	{
		choiceEnc := e.NewChoiceEncoder(sL_U2U_Info_r18SlU2UIdentityR18Constraints)
		if err := choiceEnc.EncodeChoice(int64(ie.Sl_U2U_Identity_r18.Choice), false, nil); err != nil {
			return err
		}
		switch ie.Sl_U2U_Identity_r18.Choice {
		case SL_U2U_Info_r18_Sl_U2U_Identity_r18_Sl_TargetUE_Identity_r18:
			if err := ie.Sl_U2U_Identity_r18.Sl_TargetUE_Identity_r18.Encode(e); err != nil {
				return err
			}
		case SL_U2U_Info_r18_Sl_U2U_Identity_r18_Sl_SourceUE_Identity_r18:
			if err := ie.Sl_U2U_Identity_r18.Sl_SourceUE_Identity_r18.Encode(e); err != nil {
				return err
			}
		default:
			return &per.ConstraintViolationError{TypeName: "choice", Value: int64(ie.Sl_U2U_Identity_r18.Choice), Constraint: "index out of range"}
		}
	}

	// 3. sl-E2E-QoS-InfoList-r18: sequence-of
	{
		if ie.Sl_E2E_QoS_InfoList_r18 != nil {
			seqOf := e.NewSequenceOfEncoder(sLU2UInfoR18SlE2EQoSInfoListR18Constraints)
			if err := seqOf.EncodeLength(int64(len(ie.Sl_E2E_QoS_InfoList_r18))); err != nil {
				return err
			}
			for i := range ie.Sl_E2E_QoS_InfoList_r18 {
				if err := ie.Sl_E2E_QoS_InfoList_r18[i].Encode(e); err != nil {
					return err
				}
			}
		}
	}

	// 4. sl-PerHop-QoS-InfoList-r18: sequence-of
	{
		if ie.Sl_PerHop_QoS_InfoList_r18 != nil {
			seqOf := e.NewSequenceOfEncoder(sLU2UInfoR18SlPerHopQoSInfoListR18Constraints)
			if err := seqOf.EncodeLength(int64(len(ie.Sl_PerHop_QoS_InfoList_r18))); err != nil {
				return err
			}
			for i := range ie.Sl_PerHop_QoS_InfoList_r18 {
				if err := ie.Sl_PerHop_QoS_InfoList_r18[i].Encode(e); err != nil {
					return err
				}
			}
		}
	}

	// 5. sl-PerSLRB-QoS-InfoList-r18: sequence-of
	{
		if ie.Sl_PerSLRB_QoS_InfoList_r18 != nil {
			seqOf := e.NewSequenceOfEncoder(sLU2UInfoR18SlPerSLRBQoSInfoListR18Constraints)
			if err := seqOf.EncodeLength(int64(len(ie.Sl_PerSLRB_QoS_InfoList_r18))); err != nil {
				return err
			}
			for i := range ie.Sl_PerSLRB_QoS_InfoList_r18 {
				if err := ie.Sl_PerSLRB_QoS_InfoList_r18[i].Encode(e); err != nil {
					return err
				}
			}
		}
	}

	// 6. sl-CapabilityInformationTargetRemoteUE-r18: octet-string
	{
		if ie.Sl_CapabilityInformationTargetRemoteUE_r18 != nil {
			if err := e.EncodeOctetString(ie.Sl_CapabilityInformationTargetRemoteUE_r18, sLU2UInfoR18SlCapabilityInformationTargetRemoteUER18Constraints); err != nil {
				return err
			}
		}
	}

	return nil
}

func (ie *SL_U2U_Info_r18) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(sLU2UInfoR18Constraints)

	// 1. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 2. sl-U2U-Identity-r18: choice
	{
		choiceDec := d.NewChoiceDecoder(sL_U2U_Info_r18SlU2UIdentityR18Constraints)
		index, isExt, _, err := choiceDec.DecodeChoice()
		if err != nil {
			return err
		}
		ie.Sl_U2U_Identity_r18.Choice = int(index)
		if isExt {
			return &per.DecodeError{TypeName: "choice", Reason: "extension alternative not supported"}
		}
		switch index {
		case SL_U2U_Info_r18_Sl_U2U_Identity_r18_Sl_TargetUE_Identity_r18:
			ie.Sl_U2U_Identity_r18.Sl_TargetUE_Identity_r18 = new(SL_DestinationIdentity_r16)
			if err := ie.Sl_U2U_Identity_r18.Sl_TargetUE_Identity_r18.Decode(d); err != nil {
				return err
			}
		case SL_U2U_Info_r18_Sl_U2U_Identity_r18_Sl_SourceUE_Identity_r18:
			ie.Sl_U2U_Identity_r18.Sl_SourceUE_Identity_r18 = new(SL_SourceIdentity_r17)
			if err := ie.Sl_U2U_Identity_r18.Sl_SourceUE_Identity_r18.Decode(d); err != nil {
				return err
			}
		}
	}

	// 3. sl-E2E-QoS-InfoList-r18: sequence-of
	{
		if seq.IsComponentPresent(1) {
			seqOf := d.NewSequenceOfDecoder(sLU2UInfoR18SlE2EQoSInfoListR18Constraints)
			n, err := seqOf.DecodeLength()
			if err != nil {
				return err
			}
			ie.Sl_E2E_QoS_InfoList_r18 = make([]SL_QoS_Info_r16, n)
			for i := int64(0); i < n; i++ {
				if err := ie.Sl_E2E_QoS_InfoList_r18[i].Decode(d); err != nil {
					return err
				}
			}
		}
	}

	// 4. sl-PerHop-QoS-InfoList-r18: sequence-of
	{
		if seq.IsComponentPresent(2) {
			seqOf := d.NewSequenceOfDecoder(sLU2UInfoR18SlPerHopQoSInfoListR18Constraints)
			n, err := seqOf.DecodeLength()
			if err != nil {
				return err
			}
			ie.Sl_PerHop_QoS_InfoList_r18 = make([]SL_SplitQoS_Info_r18, n)
			for i := int64(0); i < n; i++ {
				if err := ie.Sl_PerHop_QoS_InfoList_r18[i].Decode(d); err != nil {
					return err
				}
			}
		}
	}

	// 5. sl-PerSLRB-QoS-InfoList-r18: sequence-of
	{
		if seq.IsComponentPresent(3) {
			seqOf := d.NewSequenceOfDecoder(sLU2UInfoR18SlPerSLRBQoSInfoListR18Constraints)
			n, err := seqOf.DecodeLength()
			if err != nil {
				return err
			}
			ie.Sl_PerSLRB_QoS_InfoList_r18 = make([]SL_PerSLRB_QoS_Info_r18, n)
			for i := int64(0); i < n; i++ {
				if err := ie.Sl_PerSLRB_QoS_InfoList_r18[i].Decode(d); err != nil {
					return err
				}
			}
		}
	}

	// 6. sl-CapabilityInformationTargetRemoteUE-r18: octet-string
	{
		if seq.IsComponentPresent(4) {
			v, err := d.DecodeOctetString(sLU2UInfoR18SlCapabilityInformationTargetRemoteUER18Constraints)
			if err != nil {
				return err
			}
			ie.Sl_CapabilityInformationTargetRemoteUE_r18 = v
		}
	}

	return nil
}
