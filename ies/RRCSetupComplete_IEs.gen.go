// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/rrc/common"
)

// Source: RRCSetupComplete-IEs (line 1715).

var rRCSetupCompleteIEsConstraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "selectedPLMN-Identity"},
		{Name: "registeredAMF", Optional: true},
		{Name: "guami-Type", Optional: true},
		{Name: "s-NSSAI-List", Optional: true},
		{Name: "dedicatedNAS-Message"},
		{Name: "ng-5G-S-TMSI-Value", Optional: true},
		{Name: "lateNonCriticalExtension", Optional: true},
		{Name: "nonCriticalExtension", Optional: true},
	},
}

var rRCSetupCompleteIEsSelectedPLMNIdentityConstraints = per.Constrained(1, common.MaxPLMN)

const (
	RRCSetupComplete_IEs_Guami_Type_Native = 0
	RRCSetupComplete_IEs_Guami_Type_Mapped = 1
)

var rRCSetupCompleteIEsGuamiTypeConstraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{RRCSetupComplete_IEs_Guami_Type_Native, RRCSetupComplete_IEs_Guami_Type_Mapped},
}

var rRCSetupCompleteIEsSNSSAIListConstraints = per.SizeRange(1, common.MaxNrofS_NSSAI)

var rRCSetupComplete_IEsNg5GSTMSIValueConstraints = per.ChoiceConstraints{
	Extensible: false,
	RootAlternatives: []per.AlternativeInfo{
		{Name: "ng-5G-S-TMSI"},
		{Name: "ng-5G-S-TMSI-Part2"},
	},
}

const (
	RRCSetupComplete_IEs_Ng_5G_S_TMSI_Value_Ng_5G_S_TMSI       = 0
	RRCSetupComplete_IEs_Ng_5G_S_TMSI_Value_Ng_5G_S_TMSI_Part2 = 1
)

var rRCSetupCompleteIEsLateNonCriticalExtensionConstraints = per.SizeConstraints{}

type RRCSetupComplete_IEs struct {
	SelectedPLMN_Identity int64
	RegisteredAMF         *RegisteredAMF
	Guami_Type            *int64
	S_NSSAI_List          []S_NSSAI
	DedicatedNAS_Message  DedicatedNAS_Message
	Ng_5G_S_TMSI_Value    *struct {
		Choice             int
		Ng_5G_S_TMSI       *NG_5G_S_TMSI
		Ng_5G_S_TMSI_Part2 *per.BitString
	}
	LateNonCriticalExtension []byte
	NonCriticalExtension     *RRCSetupComplete_v1610_IEs
}

func (ie *RRCSetupComplete_IEs) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(rRCSetupCompleteIEsConstraints)

	// 1. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.RegisteredAMF != nil, ie.Guami_Type != nil, ie.S_NSSAI_List != nil, ie.Ng_5G_S_TMSI_Value != nil, ie.LateNonCriticalExtension != nil, ie.NonCriticalExtension != nil}); err != nil {
		return err
	}

	// 2. selectedPLMN-Identity: integer
	{
		if err := e.EncodeInteger(ie.SelectedPLMN_Identity, rRCSetupCompleteIEsSelectedPLMNIdentityConstraints); err != nil {
			return err
		}
	}

	// 3. registeredAMF: ref
	{
		if ie.RegisteredAMF != nil {
			if err := ie.RegisteredAMF.Encode(e); err != nil {
				return err
			}
		}
	}

	// 4. guami-Type: enumerated
	{
		if ie.Guami_Type != nil {
			if err := e.EncodeEnumerated(*ie.Guami_Type, rRCSetupCompleteIEsGuamiTypeConstraints); err != nil {
				return err
			}
		}
	}

	// 5. s-NSSAI-List: sequence-of
	{
		if ie.S_NSSAI_List != nil {
			seqOf := e.NewSequenceOfEncoder(rRCSetupCompleteIEsSNSSAIListConstraints)
			if err := seqOf.EncodeLength(int64(len(ie.S_NSSAI_List))); err != nil {
				return err
			}
			for i := range ie.S_NSSAI_List {
				if err := ie.S_NSSAI_List[i].Encode(e); err != nil {
					return err
				}
			}
		}
	}

	// 6. dedicatedNAS-Message: ref
	{
		if err := ie.DedicatedNAS_Message.Encode(e); err != nil {
			return err
		}
	}

	// 7. ng-5G-S-TMSI-Value: choice
	{
		if ie.Ng_5G_S_TMSI_Value != nil {
			choiceEnc := e.NewChoiceEncoder(rRCSetupComplete_IEsNg5GSTMSIValueConstraints)
			if err := choiceEnc.EncodeChoice(int64((*ie.Ng_5G_S_TMSI_Value).Choice), false, nil); err != nil {
				return err
			}
			switch (*ie.Ng_5G_S_TMSI_Value).Choice {
			case RRCSetupComplete_IEs_Ng_5G_S_TMSI_Value_Ng_5G_S_TMSI:
				if err := (*ie.Ng_5G_S_TMSI_Value).Ng_5G_S_TMSI.Encode(e); err != nil {
					return err
				}
			case RRCSetupComplete_IEs_Ng_5G_S_TMSI_Value_Ng_5G_S_TMSI_Part2:
				if err := e.EncodeBitString((*(*ie.Ng_5G_S_TMSI_Value).Ng_5G_S_TMSI_Part2), per.FixedSize(9)); err != nil {
					return err
				}
			default:
				return &per.ConstraintViolationError{TypeName: "choice", Value: int64((*ie.Ng_5G_S_TMSI_Value).Choice), Constraint: "index out of range"}
			}
		}
	}

	// 8. lateNonCriticalExtension: octet-string
	{
		if ie.LateNonCriticalExtension != nil {
			if err := e.EncodeOctetString(ie.LateNonCriticalExtension, rRCSetupCompleteIEsLateNonCriticalExtensionConstraints); err != nil {
				return err
			}
		}
	}

	// 9. nonCriticalExtension: ref
	{
		if ie.NonCriticalExtension != nil {
			if err := ie.NonCriticalExtension.Encode(e); err != nil {
				return err
			}
		}
	}

	return nil
}

func (ie *RRCSetupComplete_IEs) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(rRCSetupCompleteIEsConstraints)

	// 1. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 2. selectedPLMN-Identity: integer
	{
		v0, err := d.DecodeInteger(rRCSetupCompleteIEsSelectedPLMNIdentityConstraints)
		if err != nil {
			return err
		}
		ie.SelectedPLMN_Identity = v0
	}

	// 3. registeredAMF: ref
	{
		if seq.IsComponentPresent(1) {
			ie.RegisteredAMF = new(RegisteredAMF)
			if err := ie.RegisteredAMF.Decode(d); err != nil {
				return err
			}
		}
	}

	// 4. guami-Type: enumerated
	{
		if seq.IsComponentPresent(2) {
			idx, err := d.DecodeEnumerated(rRCSetupCompleteIEsGuamiTypeConstraints)
			if err != nil {
				return err
			}
			ie.Guami_Type = &idx
		}
	}

	// 5. s-NSSAI-List: sequence-of
	{
		if seq.IsComponentPresent(3) {
			seqOf := d.NewSequenceOfDecoder(rRCSetupCompleteIEsSNSSAIListConstraints)
			n, err := seqOf.DecodeLength()
			if err != nil {
				return err
			}
			ie.S_NSSAI_List = make([]S_NSSAI, n)
			for i := int64(0); i < n; i++ {
				if err := ie.S_NSSAI_List[i].Decode(d); err != nil {
					return err
				}
			}
		}
	}

	// 6. dedicatedNAS-Message: ref
	{
		if err := ie.DedicatedNAS_Message.Decode(d); err != nil {
			return err
		}
	}

	// 7. ng-5G-S-TMSI-Value: choice
	{
		if seq.IsComponentPresent(5) {
			ie.Ng_5G_S_TMSI_Value = &struct {
				Choice             int
				Ng_5G_S_TMSI       *NG_5G_S_TMSI
				Ng_5G_S_TMSI_Part2 *per.BitString
			}{}
			choiceDec := d.NewChoiceDecoder(rRCSetupComplete_IEsNg5GSTMSIValueConstraints)
			index, isExt, _, err := choiceDec.DecodeChoice()
			if err != nil {
				return err
			}
			(*ie.Ng_5G_S_TMSI_Value).Choice = int(index)
			if isExt {
				return &per.DecodeError{TypeName: "choice", Reason: "extension alternative not supported"}
			}
			switch index {
			case RRCSetupComplete_IEs_Ng_5G_S_TMSI_Value_Ng_5G_S_TMSI:
				(*ie.Ng_5G_S_TMSI_Value).Ng_5G_S_TMSI = new(NG_5G_S_TMSI)
				if err := (*ie.Ng_5G_S_TMSI_Value).Ng_5G_S_TMSI.Decode(d); err != nil {
					return err
				}
			case RRCSetupComplete_IEs_Ng_5G_S_TMSI_Value_Ng_5G_S_TMSI_Part2:
				(*ie.Ng_5G_S_TMSI_Value).Ng_5G_S_TMSI_Part2 = new(per.BitString)
				v, err := d.DecodeBitString(per.FixedSize(9))
				if err != nil {
					return err
				}
				(*(*ie.Ng_5G_S_TMSI_Value).Ng_5G_S_TMSI_Part2) = v
			}
		}
	}

	// 8. lateNonCriticalExtension: octet-string
	{
		if seq.IsComponentPresent(6) {
			v, err := d.DecodeOctetString(rRCSetupCompleteIEsLateNonCriticalExtensionConstraints)
			if err != nil {
				return err
			}
			ie.LateNonCriticalExtension = v
		}
	}

	// 9. nonCriticalExtension: ref
	{
		if seq.IsComponentPresent(7) {
			ie.NonCriticalExtension = new(RRCSetupComplete_v1610_IEs)
			if err := ie.NonCriticalExtension.Decode(d); err != nil {
				return err
			}
		}
	}

	return nil
}
