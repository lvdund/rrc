// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/rrc/common"
)

// Source: PosSystemInformation-r16-IEs (line 4809).

var posSystemInformationR16IEsConstraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "posSIB-TypeAndInfo-r16"},
		{Name: "lateNonCriticalExtension", Optional: true},
		{Name: "nonCriticalExtension", Optional: true},
	},
}

var posSystemInformationR16IEsPosSIBTypeAndInfoR16Constraints = per.SizeRange(1, common.MaxSIB)

var posSystemInformationR16IEsLateNonCriticalExtensionConstraints = per.SizeConstraints{}

var posSystemInformationR16IEsPosSIBTypeAndInfoR16ElemConstraints = per.ChoiceConstraints{
	Extensible: true,
	RootAlternatives: []per.AlternativeInfo{
		{Name: "posSib1-1-r16"},
		{Name: "posSib1-2-r16"},
		{Name: "posSib1-3-r16"},
		{Name: "posSib1-4-r16"},
		{Name: "posSib1-5-r16"},
		{Name: "posSib1-6-r16"},
		{Name: "posSib1-7-r16"},
		{Name: "posSib1-8-r16"},
		{Name: "posSib2-1-r16"},
		{Name: "posSib2-2-r16"},
		{Name: "posSib2-3-r16"},
		{Name: "posSib2-4-r16"},
		{Name: "posSib2-5-r16"},
		{Name: "posSib2-6-r16"},
		{Name: "posSib2-7-r16"},
		{Name: "posSib2-8-r16"},
		{Name: "posSib2-9-r16"},
		{Name: "posSib2-10-r16"},
		{Name: "posSib2-11-r16"},
		{Name: "posSib2-12-r16"},
		{Name: "posSib2-13-r16"},
		{Name: "posSib2-14-r16"},
		{Name: "posSib2-15-r16"},
		{Name: "posSib2-16-r16"},
		{Name: "posSib2-17-r16"},
		{Name: "posSib2-18-r16"},
		{Name: "posSib2-19-r16"},
		{Name: "posSib2-20-r16"},
		{Name: "posSib2-21-r16"},
		{Name: "posSib2-22-r16"},
		{Name: "posSib2-23-r16"},
		{Name: "posSib3-1-r16"},
		{Name: "posSib4-1-r16"},
		{Name: "posSib5-1-r16"},
		{Name: "posSib6-1-r16"},
		{Name: "posSib6-2-r16"},
		{Name: "posSib6-3-r16"},
	},
	ExtAlternatives: []per.AlternativeInfo{
		{Name: "posSib1-9-v1700"},
		{Name: "posSib1-10-v1700"},
		{Name: "posSib2-24-v1700"},
		{Name: "posSib2-25-v1700"},
		{Name: "posSib6-4-v1700"},
		{Name: "posSib6-5-v1700"},
		{Name: "posSib6-6-v1700"},
		{Name: "posSib2-17a-v1770"},
		{Name: "posSib2-18a-v1770"},
		{Name: "posSib2-20a-v1770"},
		{Name: "posSib1-11-v1800"},
		{Name: "posSib1-12-v1800"},
		{Name: "posSib2-26-v1800"},
		{Name: "posSib2-27-v1800"},
		{Name: "posSib6-7-v1800"},
		{Name: "posSib7-1-v1800"},
		{Name: "posSib7-2-v1800"},
		{Name: "posSib7-3-v1800"},
		{Name: "posSib7-4-v1800"},
	},
}

const (
	PosSystemInformation_r16_IEs_PosSIB_TypeAndInfo_r16_Elem_PosSib1_1_r16  = 0
	PosSystemInformation_r16_IEs_PosSIB_TypeAndInfo_r16_Elem_PosSib1_2_r16  = 1
	PosSystemInformation_r16_IEs_PosSIB_TypeAndInfo_r16_Elem_PosSib1_3_r16  = 2
	PosSystemInformation_r16_IEs_PosSIB_TypeAndInfo_r16_Elem_PosSib1_4_r16  = 3
	PosSystemInformation_r16_IEs_PosSIB_TypeAndInfo_r16_Elem_PosSib1_5_r16  = 4
	PosSystemInformation_r16_IEs_PosSIB_TypeAndInfo_r16_Elem_PosSib1_6_r16  = 5
	PosSystemInformation_r16_IEs_PosSIB_TypeAndInfo_r16_Elem_PosSib1_7_r16  = 6
	PosSystemInformation_r16_IEs_PosSIB_TypeAndInfo_r16_Elem_PosSib1_8_r16  = 7
	PosSystemInformation_r16_IEs_PosSIB_TypeAndInfo_r16_Elem_PosSib2_1_r16  = 8
	PosSystemInformation_r16_IEs_PosSIB_TypeAndInfo_r16_Elem_PosSib2_2_r16  = 9
	PosSystemInformation_r16_IEs_PosSIB_TypeAndInfo_r16_Elem_PosSib2_3_r16  = 10
	PosSystemInformation_r16_IEs_PosSIB_TypeAndInfo_r16_Elem_PosSib2_4_r16  = 11
	PosSystemInformation_r16_IEs_PosSIB_TypeAndInfo_r16_Elem_PosSib2_5_r16  = 12
	PosSystemInformation_r16_IEs_PosSIB_TypeAndInfo_r16_Elem_PosSib2_6_r16  = 13
	PosSystemInformation_r16_IEs_PosSIB_TypeAndInfo_r16_Elem_PosSib2_7_r16  = 14
	PosSystemInformation_r16_IEs_PosSIB_TypeAndInfo_r16_Elem_PosSib2_8_r16  = 15
	PosSystemInformation_r16_IEs_PosSIB_TypeAndInfo_r16_Elem_PosSib2_9_r16  = 16
	PosSystemInformation_r16_IEs_PosSIB_TypeAndInfo_r16_Elem_PosSib2_10_r16 = 17
	PosSystemInformation_r16_IEs_PosSIB_TypeAndInfo_r16_Elem_PosSib2_11_r16 = 18
	PosSystemInformation_r16_IEs_PosSIB_TypeAndInfo_r16_Elem_PosSib2_12_r16 = 19
	PosSystemInformation_r16_IEs_PosSIB_TypeAndInfo_r16_Elem_PosSib2_13_r16 = 20
	PosSystemInformation_r16_IEs_PosSIB_TypeAndInfo_r16_Elem_PosSib2_14_r16 = 21
	PosSystemInformation_r16_IEs_PosSIB_TypeAndInfo_r16_Elem_PosSib2_15_r16 = 22
	PosSystemInformation_r16_IEs_PosSIB_TypeAndInfo_r16_Elem_PosSib2_16_r16 = 23
	PosSystemInformation_r16_IEs_PosSIB_TypeAndInfo_r16_Elem_PosSib2_17_r16 = 24
	PosSystemInformation_r16_IEs_PosSIB_TypeAndInfo_r16_Elem_PosSib2_18_r16 = 25
	PosSystemInformation_r16_IEs_PosSIB_TypeAndInfo_r16_Elem_PosSib2_19_r16 = 26
	PosSystemInformation_r16_IEs_PosSIB_TypeAndInfo_r16_Elem_PosSib2_20_r16 = 27
	PosSystemInformation_r16_IEs_PosSIB_TypeAndInfo_r16_Elem_PosSib2_21_r16 = 28
	PosSystemInformation_r16_IEs_PosSIB_TypeAndInfo_r16_Elem_PosSib2_22_r16 = 29
	PosSystemInformation_r16_IEs_PosSIB_TypeAndInfo_r16_Elem_PosSib2_23_r16 = 30
	PosSystemInformation_r16_IEs_PosSIB_TypeAndInfo_r16_Elem_PosSib3_1_r16  = 31
	PosSystemInformation_r16_IEs_PosSIB_TypeAndInfo_r16_Elem_PosSib4_1_r16  = 32
	PosSystemInformation_r16_IEs_PosSIB_TypeAndInfo_r16_Elem_PosSib5_1_r16  = 33
	PosSystemInformation_r16_IEs_PosSIB_TypeAndInfo_r16_Elem_PosSib6_1_r16  = 34
	PosSystemInformation_r16_IEs_PosSIB_TypeAndInfo_r16_Elem_PosSib6_2_r16  = 35
	PosSystemInformation_r16_IEs_PosSIB_TypeAndInfo_r16_Elem_PosSib6_3_r16  = 36
)

type PosSystemInformation_r16_IEs struct {
	PosSIB_TypeAndInfo_r16 []struct {
		Choice         int
		PosSib1_1_r16  *SIBpos_r16
		PosSib1_2_r16  *SIBpos_r16
		PosSib1_3_r16  *SIBpos_r16
		PosSib1_4_r16  *SIBpos_r16
		PosSib1_5_r16  *SIBpos_r16
		PosSib1_6_r16  *SIBpos_r16
		PosSib1_7_r16  *SIBpos_r16
		PosSib1_8_r16  *SIBpos_r16
		PosSib2_1_r16  *SIBpos_r16
		PosSib2_2_r16  *SIBpos_r16
		PosSib2_3_r16  *SIBpos_r16
		PosSib2_4_r16  *SIBpos_r16
		PosSib2_5_r16  *SIBpos_r16
		PosSib2_6_r16  *SIBpos_r16
		PosSib2_7_r16  *SIBpos_r16
		PosSib2_8_r16  *SIBpos_r16
		PosSib2_9_r16  *SIBpos_r16
		PosSib2_10_r16 *SIBpos_r16
		PosSib2_11_r16 *SIBpos_r16
		PosSib2_12_r16 *SIBpos_r16
		PosSib2_13_r16 *SIBpos_r16
		PosSib2_14_r16 *SIBpos_r16
		PosSib2_15_r16 *SIBpos_r16
		PosSib2_16_r16 *SIBpos_r16
		PosSib2_17_r16 *SIBpos_r16
		PosSib2_18_r16 *SIBpos_r16
		PosSib2_19_r16 *SIBpos_r16
		PosSib2_20_r16 *SIBpos_r16
		PosSib2_21_r16 *SIBpos_r16
		PosSib2_22_r16 *SIBpos_r16
		PosSib2_23_r16 *SIBpos_r16
		PosSib3_1_r16  *SIBpos_r16
		PosSib4_1_r16  *SIBpos_r16
		PosSib5_1_r16  *SIBpos_r16
		PosSib6_1_r16  *SIBpos_r16
		PosSib6_2_r16  *SIBpos_r16
		PosSib6_3_r16  *SIBpos_r16
	}
	LateNonCriticalExtension []byte
	NonCriticalExtension     *struct{}
}

func (ie *PosSystemInformation_r16_IEs) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(posSystemInformationR16IEsConstraints)

	// 1. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.LateNonCriticalExtension != nil, ie.NonCriticalExtension != nil}); err != nil {
		return err
	}

	// 2. posSIB-TypeAndInfo-r16: sequence-of
	{
		seqOf := e.NewSequenceOfEncoder(posSystemInformationR16IEsPosSIBTypeAndInfoR16Constraints)
		if err := seqOf.EncodeLength(int64(len(ie.PosSIB_TypeAndInfo_r16))); err != nil {
			return err
		}
		for i := range ie.PosSIB_TypeAndInfo_r16 {
			choiceEnc := e.NewChoiceEncoder(posSystemInformationR16IEsPosSIBTypeAndInfoR16ElemConstraints)
			if err := choiceEnc.EncodeChoice(int64(ie.PosSIB_TypeAndInfo_r16[i].Choice), false, nil); err != nil {
				return err
			}
			switch ie.PosSIB_TypeAndInfo_r16[i].Choice {
			case PosSystemInformation_r16_IEs_PosSIB_TypeAndInfo_r16_Elem_PosSib1_1_r16:
				if err := ie.PosSIB_TypeAndInfo_r16[i].PosSib1_1_r16.Encode(e); err != nil {
					return err
				}
			case PosSystemInformation_r16_IEs_PosSIB_TypeAndInfo_r16_Elem_PosSib1_2_r16:
				if err := ie.PosSIB_TypeAndInfo_r16[i].PosSib1_2_r16.Encode(e); err != nil {
					return err
				}
			case PosSystemInformation_r16_IEs_PosSIB_TypeAndInfo_r16_Elem_PosSib1_3_r16:
				if err := ie.PosSIB_TypeAndInfo_r16[i].PosSib1_3_r16.Encode(e); err != nil {
					return err
				}
			case PosSystemInformation_r16_IEs_PosSIB_TypeAndInfo_r16_Elem_PosSib1_4_r16:
				if err := ie.PosSIB_TypeAndInfo_r16[i].PosSib1_4_r16.Encode(e); err != nil {
					return err
				}
			case PosSystemInformation_r16_IEs_PosSIB_TypeAndInfo_r16_Elem_PosSib1_5_r16:
				if err := ie.PosSIB_TypeAndInfo_r16[i].PosSib1_5_r16.Encode(e); err != nil {
					return err
				}
			case PosSystemInformation_r16_IEs_PosSIB_TypeAndInfo_r16_Elem_PosSib1_6_r16:
				if err := ie.PosSIB_TypeAndInfo_r16[i].PosSib1_6_r16.Encode(e); err != nil {
					return err
				}
			case PosSystemInformation_r16_IEs_PosSIB_TypeAndInfo_r16_Elem_PosSib1_7_r16:
				if err := ie.PosSIB_TypeAndInfo_r16[i].PosSib1_7_r16.Encode(e); err != nil {
					return err
				}
			case PosSystemInformation_r16_IEs_PosSIB_TypeAndInfo_r16_Elem_PosSib1_8_r16:
				if err := ie.PosSIB_TypeAndInfo_r16[i].PosSib1_8_r16.Encode(e); err != nil {
					return err
				}
			case PosSystemInformation_r16_IEs_PosSIB_TypeAndInfo_r16_Elem_PosSib2_1_r16:
				if err := ie.PosSIB_TypeAndInfo_r16[i].PosSib2_1_r16.Encode(e); err != nil {
					return err
				}
			case PosSystemInformation_r16_IEs_PosSIB_TypeAndInfo_r16_Elem_PosSib2_2_r16:
				if err := ie.PosSIB_TypeAndInfo_r16[i].PosSib2_2_r16.Encode(e); err != nil {
					return err
				}
			case PosSystemInformation_r16_IEs_PosSIB_TypeAndInfo_r16_Elem_PosSib2_3_r16:
				if err := ie.PosSIB_TypeAndInfo_r16[i].PosSib2_3_r16.Encode(e); err != nil {
					return err
				}
			case PosSystemInformation_r16_IEs_PosSIB_TypeAndInfo_r16_Elem_PosSib2_4_r16:
				if err := ie.PosSIB_TypeAndInfo_r16[i].PosSib2_4_r16.Encode(e); err != nil {
					return err
				}
			case PosSystemInformation_r16_IEs_PosSIB_TypeAndInfo_r16_Elem_PosSib2_5_r16:
				if err := ie.PosSIB_TypeAndInfo_r16[i].PosSib2_5_r16.Encode(e); err != nil {
					return err
				}
			case PosSystemInformation_r16_IEs_PosSIB_TypeAndInfo_r16_Elem_PosSib2_6_r16:
				if err := ie.PosSIB_TypeAndInfo_r16[i].PosSib2_6_r16.Encode(e); err != nil {
					return err
				}
			case PosSystemInformation_r16_IEs_PosSIB_TypeAndInfo_r16_Elem_PosSib2_7_r16:
				if err := ie.PosSIB_TypeAndInfo_r16[i].PosSib2_7_r16.Encode(e); err != nil {
					return err
				}
			case PosSystemInformation_r16_IEs_PosSIB_TypeAndInfo_r16_Elem_PosSib2_8_r16:
				if err := ie.PosSIB_TypeAndInfo_r16[i].PosSib2_8_r16.Encode(e); err != nil {
					return err
				}
			case PosSystemInformation_r16_IEs_PosSIB_TypeAndInfo_r16_Elem_PosSib2_9_r16:
				if err := ie.PosSIB_TypeAndInfo_r16[i].PosSib2_9_r16.Encode(e); err != nil {
					return err
				}
			case PosSystemInformation_r16_IEs_PosSIB_TypeAndInfo_r16_Elem_PosSib2_10_r16:
				if err := ie.PosSIB_TypeAndInfo_r16[i].PosSib2_10_r16.Encode(e); err != nil {
					return err
				}
			case PosSystemInformation_r16_IEs_PosSIB_TypeAndInfo_r16_Elem_PosSib2_11_r16:
				if err := ie.PosSIB_TypeAndInfo_r16[i].PosSib2_11_r16.Encode(e); err != nil {
					return err
				}
			case PosSystemInformation_r16_IEs_PosSIB_TypeAndInfo_r16_Elem_PosSib2_12_r16:
				if err := ie.PosSIB_TypeAndInfo_r16[i].PosSib2_12_r16.Encode(e); err != nil {
					return err
				}
			case PosSystemInformation_r16_IEs_PosSIB_TypeAndInfo_r16_Elem_PosSib2_13_r16:
				if err := ie.PosSIB_TypeAndInfo_r16[i].PosSib2_13_r16.Encode(e); err != nil {
					return err
				}
			case PosSystemInformation_r16_IEs_PosSIB_TypeAndInfo_r16_Elem_PosSib2_14_r16:
				if err := ie.PosSIB_TypeAndInfo_r16[i].PosSib2_14_r16.Encode(e); err != nil {
					return err
				}
			case PosSystemInformation_r16_IEs_PosSIB_TypeAndInfo_r16_Elem_PosSib2_15_r16:
				if err := ie.PosSIB_TypeAndInfo_r16[i].PosSib2_15_r16.Encode(e); err != nil {
					return err
				}
			case PosSystemInformation_r16_IEs_PosSIB_TypeAndInfo_r16_Elem_PosSib2_16_r16:
				if err := ie.PosSIB_TypeAndInfo_r16[i].PosSib2_16_r16.Encode(e); err != nil {
					return err
				}
			case PosSystemInformation_r16_IEs_PosSIB_TypeAndInfo_r16_Elem_PosSib2_17_r16:
				if err := ie.PosSIB_TypeAndInfo_r16[i].PosSib2_17_r16.Encode(e); err != nil {
					return err
				}
			case PosSystemInformation_r16_IEs_PosSIB_TypeAndInfo_r16_Elem_PosSib2_18_r16:
				if err := ie.PosSIB_TypeAndInfo_r16[i].PosSib2_18_r16.Encode(e); err != nil {
					return err
				}
			case PosSystemInformation_r16_IEs_PosSIB_TypeAndInfo_r16_Elem_PosSib2_19_r16:
				if err := ie.PosSIB_TypeAndInfo_r16[i].PosSib2_19_r16.Encode(e); err != nil {
					return err
				}
			case PosSystemInformation_r16_IEs_PosSIB_TypeAndInfo_r16_Elem_PosSib2_20_r16:
				if err := ie.PosSIB_TypeAndInfo_r16[i].PosSib2_20_r16.Encode(e); err != nil {
					return err
				}
			case PosSystemInformation_r16_IEs_PosSIB_TypeAndInfo_r16_Elem_PosSib2_21_r16:
				if err := ie.PosSIB_TypeAndInfo_r16[i].PosSib2_21_r16.Encode(e); err != nil {
					return err
				}
			case PosSystemInformation_r16_IEs_PosSIB_TypeAndInfo_r16_Elem_PosSib2_22_r16:
				if err := ie.PosSIB_TypeAndInfo_r16[i].PosSib2_22_r16.Encode(e); err != nil {
					return err
				}
			case PosSystemInformation_r16_IEs_PosSIB_TypeAndInfo_r16_Elem_PosSib2_23_r16:
				if err := ie.PosSIB_TypeAndInfo_r16[i].PosSib2_23_r16.Encode(e); err != nil {
					return err
				}
			case PosSystemInformation_r16_IEs_PosSIB_TypeAndInfo_r16_Elem_PosSib3_1_r16:
				if err := ie.PosSIB_TypeAndInfo_r16[i].PosSib3_1_r16.Encode(e); err != nil {
					return err
				}
			case PosSystemInformation_r16_IEs_PosSIB_TypeAndInfo_r16_Elem_PosSib4_1_r16:
				if err := ie.PosSIB_TypeAndInfo_r16[i].PosSib4_1_r16.Encode(e); err != nil {
					return err
				}
			case PosSystemInformation_r16_IEs_PosSIB_TypeAndInfo_r16_Elem_PosSib5_1_r16:
				if err := ie.PosSIB_TypeAndInfo_r16[i].PosSib5_1_r16.Encode(e); err != nil {
					return err
				}
			case PosSystemInformation_r16_IEs_PosSIB_TypeAndInfo_r16_Elem_PosSib6_1_r16:
				if err := ie.PosSIB_TypeAndInfo_r16[i].PosSib6_1_r16.Encode(e); err != nil {
					return err
				}
			case PosSystemInformation_r16_IEs_PosSIB_TypeAndInfo_r16_Elem_PosSib6_2_r16:
				if err := ie.PosSIB_TypeAndInfo_r16[i].PosSib6_2_r16.Encode(e); err != nil {
					return err
				}
			case PosSystemInformation_r16_IEs_PosSIB_TypeAndInfo_r16_Elem_PosSib6_3_r16:
				if err := ie.PosSIB_TypeAndInfo_r16[i].PosSib6_3_r16.Encode(e); err != nil {
					return err
				}
			}
		}
	}

	// 3. lateNonCriticalExtension: octet-string
	{
		if ie.LateNonCriticalExtension != nil {
			if err := e.EncodeOctetString(ie.LateNonCriticalExtension, posSystemInformationR16IEsLateNonCriticalExtensionConstraints); err != nil {
				return err
			}
		}
	}

	// 4. nonCriticalExtension: sequence
	{
		if ie.NonCriticalExtension != nil {
			// empty SEQUENCE {}: no content to encode.
		}
	}

	return nil
}

func (ie *PosSystemInformation_r16_IEs) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(posSystemInformationR16IEsConstraints)

	// 1. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 2. posSIB-TypeAndInfo-r16: sequence-of
	{
		seqOf := d.NewSequenceOfDecoder(posSystemInformationR16IEsPosSIBTypeAndInfoR16Constraints)
		n, err := seqOf.DecodeLength()
		if err != nil {
			return err
		}
		ie.PosSIB_TypeAndInfo_r16 = make([]struct {
			Choice         int
			PosSib1_1_r16  *SIBpos_r16
			PosSib1_2_r16  *SIBpos_r16
			PosSib1_3_r16  *SIBpos_r16
			PosSib1_4_r16  *SIBpos_r16
			PosSib1_5_r16  *SIBpos_r16
			PosSib1_6_r16  *SIBpos_r16
			PosSib1_7_r16  *SIBpos_r16
			PosSib1_8_r16  *SIBpos_r16
			PosSib2_1_r16  *SIBpos_r16
			PosSib2_2_r16  *SIBpos_r16
			PosSib2_3_r16  *SIBpos_r16
			PosSib2_4_r16  *SIBpos_r16
			PosSib2_5_r16  *SIBpos_r16
			PosSib2_6_r16  *SIBpos_r16
			PosSib2_7_r16  *SIBpos_r16
			PosSib2_8_r16  *SIBpos_r16
			PosSib2_9_r16  *SIBpos_r16
			PosSib2_10_r16 *SIBpos_r16
			PosSib2_11_r16 *SIBpos_r16
			PosSib2_12_r16 *SIBpos_r16
			PosSib2_13_r16 *SIBpos_r16
			PosSib2_14_r16 *SIBpos_r16
			PosSib2_15_r16 *SIBpos_r16
			PosSib2_16_r16 *SIBpos_r16
			PosSib2_17_r16 *SIBpos_r16
			PosSib2_18_r16 *SIBpos_r16
			PosSib2_19_r16 *SIBpos_r16
			PosSib2_20_r16 *SIBpos_r16
			PosSib2_21_r16 *SIBpos_r16
			PosSib2_22_r16 *SIBpos_r16
			PosSib2_23_r16 *SIBpos_r16
			PosSib3_1_r16  *SIBpos_r16
			PosSib4_1_r16  *SIBpos_r16
			PosSib5_1_r16  *SIBpos_r16
			PosSib6_1_r16  *SIBpos_r16
			PosSib6_2_r16  *SIBpos_r16
			PosSib6_3_r16  *SIBpos_r16
		}, n)
		for i := int64(0); i < n; i++ {
			choiceDec := d.NewChoiceDecoder(posSystemInformationR16IEsPosSIBTypeAndInfoR16ElemConstraints)
			index, _, _, err := choiceDec.DecodeChoice()
			if err != nil {
				return err
			}
			ie.PosSIB_TypeAndInfo_r16[i].Choice = int(index)
			switch index {
			case PosSystemInformation_r16_IEs_PosSIB_TypeAndInfo_r16_Elem_PosSib1_1_r16:
				ie.PosSIB_TypeAndInfo_r16[i].PosSib1_1_r16 = new(SIBpos_r16)
				if err := ie.PosSIB_TypeAndInfo_r16[i].PosSib1_1_r16.Decode(d); err != nil {
					return err
				}
			case PosSystemInformation_r16_IEs_PosSIB_TypeAndInfo_r16_Elem_PosSib1_2_r16:
				ie.PosSIB_TypeAndInfo_r16[i].PosSib1_2_r16 = new(SIBpos_r16)
				if err := ie.PosSIB_TypeAndInfo_r16[i].PosSib1_2_r16.Decode(d); err != nil {
					return err
				}
			case PosSystemInformation_r16_IEs_PosSIB_TypeAndInfo_r16_Elem_PosSib1_3_r16:
				ie.PosSIB_TypeAndInfo_r16[i].PosSib1_3_r16 = new(SIBpos_r16)
				if err := ie.PosSIB_TypeAndInfo_r16[i].PosSib1_3_r16.Decode(d); err != nil {
					return err
				}
			case PosSystemInformation_r16_IEs_PosSIB_TypeAndInfo_r16_Elem_PosSib1_4_r16:
				ie.PosSIB_TypeAndInfo_r16[i].PosSib1_4_r16 = new(SIBpos_r16)
				if err := ie.PosSIB_TypeAndInfo_r16[i].PosSib1_4_r16.Decode(d); err != nil {
					return err
				}
			case PosSystemInformation_r16_IEs_PosSIB_TypeAndInfo_r16_Elem_PosSib1_5_r16:
				ie.PosSIB_TypeAndInfo_r16[i].PosSib1_5_r16 = new(SIBpos_r16)
				if err := ie.PosSIB_TypeAndInfo_r16[i].PosSib1_5_r16.Decode(d); err != nil {
					return err
				}
			case PosSystemInformation_r16_IEs_PosSIB_TypeAndInfo_r16_Elem_PosSib1_6_r16:
				ie.PosSIB_TypeAndInfo_r16[i].PosSib1_6_r16 = new(SIBpos_r16)
				if err := ie.PosSIB_TypeAndInfo_r16[i].PosSib1_6_r16.Decode(d); err != nil {
					return err
				}
			case PosSystemInformation_r16_IEs_PosSIB_TypeAndInfo_r16_Elem_PosSib1_7_r16:
				ie.PosSIB_TypeAndInfo_r16[i].PosSib1_7_r16 = new(SIBpos_r16)
				if err := ie.PosSIB_TypeAndInfo_r16[i].PosSib1_7_r16.Decode(d); err != nil {
					return err
				}
			case PosSystemInformation_r16_IEs_PosSIB_TypeAndInfo_r16_Elem_PosSib1_8_r16:
				ie.PosSIB_TypeAndInfo_r16[i].PosSib1_8_r16 = new(SIBpos_r16)
				if err := ie.PosSIB_TypeAndInfo_r16[i].PosSib1_8_r16.Decode(d); err != nil {
					return err
				}
			case PosSystemInformation_r16_IEs_PosSIB_TypeAndInfo_r16_Elem_PosSib2_1_r16:
				ie.PosSIB_TypeAndInfo_r16[i].PosSib2_1_r16 = new(SIBpos_r16)
				if err := ie.PosSIB_TypeAndInfo_r16[i].PosSib2_1_r16.Decode(d); err != nil {
					return err
				}
			case PosSystemInformation_r16_IEs_PosSIB_TypeAndInfo_r16_Elem_PosSib2_2_r16:
				ie.PosSIB_TypeAndInfo_r16[i].PosSib2_2_r16 = new(SIBpos_r16)
				if err := ie.PosSIB_TypeAndInfo_r16[i].PosSib2_2_r16.Decode(d); err != nil {
					return err
				}
			case PosSystemInformation_r16_IEs_PosSIB_TypeAndInfo_r16_Elem_PosSib2_3_r16:
				ie.PosSIB_TypeAndInfo_r16[i].PosSib2_3_r16 = new(SIBpos_r16)
				if err := ie.PosSIB_TypeAndInfo_r16[i].PosSib2_3_r16.Decode(d); err != nil {
					return err
				}
			case PosSystemInformation_r16_IEs_PosSIB_TypeAndInfo_r16_Elem_PosSib2_4_r16:
				ie.PosSIB_TypeAndInfo_r16[i].PosSib2_4_r16 = new(SIBpos_r16)
				if err := ie.PosSIB_TypeAndInfo_r16[i].PosSib2_4_r16.Decode(d); err != nil {
					return err
				}
			case PosSystemInformation_r16_IEs_PosSIB_TypeAndInfo_r16_Elem_PosSib2_5_r16:
				ie.PosSIB_TypeAndInfo_r16[i].PosSib2_5_r16 = new(SIBpos_r16)
				if err := ie.PosSIB_TypeAndInfo_r16[i].PosSib2_5_r16.Decode(d); err != nil {
					return err
				}
			case PosSystemInformation_r16_IEs_PosSIB_TypeAndInfo_r16_Elem_PosSib2_6_r16:
				ie.PosSIB_TypeAndInfo_r16[i].PosSib2_6_r16 = new(SIBpos_r16)
				if err := ie.PosSIB_TypeAndInfo_r16[i].PosSib2_6_r16.Decode(d); err != nil {
					return err
				}
			case PosSystemInformation_r16_IEs_PosSIB_TypeAndInfo_r16_Elem_PosSib2_7_r16:
				ie.PosSIB_TypeAndInfo_r16[i].PosSib2_7_r16 = new(SIBpos_r16)
				if err := ie.PosSIB_TypeAndInfo_r16[i].PosSib2_7_r16.Decode(d); err != nil {
					return err
				}
			case PosSystemInformation_r16_IEs_PosSIB_TypeAndInfo_r16_Elem_PosSib2_8_r16:
				ie.PosSIB_TypeAndInfo_r16[i].PosSib2_8_r16 = new(SIBpos_r16)
				if err := ie.PosSIB_TypeAndInfo_r16[i].PosSib2_8_r16.Decode(d); err != nil {
					return err
				}
			case PosSystemInformation_r16_IEs_PosSIB_TypeAndInfo_r16_Elem_PosSib2_9_r16:
				ie.PosSIB_TypeAndInfo_r16[i].PosSib2_9_r16 = new(SIBpos_r16)
				if err := ie.PosSIB_TypeAndInfo_r16[i].PosSib2_9_r16.Decode(d); err != nil {
					return err
				}
			case PosSystemInformation_r16_IEs_PosSIB_TypeAndInfo_r16_Elem_PosSib2_10_r16:
				ie.PosSIB_TypeAndInfo_r16[i].PosSib2_10_r16 = new(SIBpos_r16)
				if err := ie.PosSIB_TypeAndInfo_r16[i].PosSib2_10_r16.Decode(d); err != nil {
					return err
				}
			case PosSystemInformation_r16_IEs_PosSIB_TypeAndInfo_r16_Elem_PosSib2_11_r16:
				ie.PosSIB_TypeAndInfo_r16[i].PosSib2_11_r16 = new(SIBpos_r16)
				if err := ie.PosSIB_TypeAndInfo_r16[i].PosSib2_11_r16.Decode(d); err != nil {
					return err
				}
			case PosSystemInformation_r16_IEs_PosSIB_TypeAndInfo_r16_Elem_PosSib2_12_r16:
				ie.PosSIB_TypeAndInfo_r16[i].PosSib2_12_r16 = new(SIBpos_r16)
				if err := ie.PosSIB_TypeAndInfo_r16[i].PosSib2_12_r16.Decode(d); err != nil {
					return err
				}
			case PosSystemInformation_r16_IEs_PosSIB_TypeAndInfo_r16_Elem_PosSib2_13_r16:
				ie.PosSIB_TypeAndInfo_r16[i].PosSib2_13_r16 = new(SIBpos_r16)
				if err := ie.PosSIB_TypeAndInfo_r16[i].PosSib2_13_r16.Decode(d); err != nil {
					return err
				}
			case PosSystemInformation_r16_IEs_PosSIB_TypeAndInfo_r16_Elem_PosSib2_14_r16:
				ie.PosSIB_TypeAndInfo_r16[i].PosSib2_14_r16 = new(SIBpos_r16)
				if err := ie.PosSIB_TypeAndInfo_r16[i].PosSib2_14_r16.Decode(d); err != nil {
					return err
				}
			case PosSystemInformation_r16_IEs_PosSIB_TypeAndInfo_r16_Elem_PosSib2_15_r16:
				ie.PosSIB_TypeAndInfo_r16[i].PosSib2_15_r16 = new(SIBpos_r16)
				if err := ie.PosSIB_TypeAndInfo_r16[i].PosSib2_15_r16.Decode(d); err != nil {
					return err
				}
			case PosSystemInformation_r16_IEs_PosSIB_TypeAndInfo_r16_Elem_PosSib2_16_r16:
				ie.PosSIB_TypeAndInfo_r16[i].PosSib2_16_r16 = new(SIBpos_r16)
				if err := ie.PosSIB_TypeAndInfo_r16[i].PosSib2_16_r16.Decode(d); err != nil {
					return err
				}
			case PosSystemInformation_r16_IEs_PosSIB_TypeAndInfo_r16_Elem_PosSib2_17_r16:
				ie.PosSIB_TypeAndInfo_r16[i].PosSib2_17_r16 = new(SIBpos_r16)
				if err := ie.PosSIB_TypeAndInfo_r16[i].PosSib2_17_r16.Decode(d); err != nil {
					return err
				}
			case PosSystemInformation_r16_IEs_PosSIB_TypeAndInfo_r16_Elem_PosSib2_18_r16:
				ie.PosSIB_TypeAndInfo_r16[i].PosSib2_18_r16 = new(SIBpos_r16)
				if err := ie.PosSIB_TypeAndInfo_r16[i].PosSib2_18_r16.Decode(d); err != nil {
					return err
				}
			case PosSystemInformation_r16_IEs_PosSIB_TypeAndInfo_r16_Elem_PosSib2_19_r16:
				ie.PosSIB_TypeAndInfo_r16[i].PosSib2_19_r16 = new(SIBpos_r16)
				if err := ie.PosSIB_TypeAndInfo_r16[i].PosSib2_19_r16.Decode(d); err != nil {
					return err
				}
			case PosSystemInformation_r16_IEs_PosSIB_TypeAndInfo_r16_Elem_PosSib2_20_r16:
				ie.PosSIB_TypeAndInfo_r16[i].PosSib2_20_r16 = new(SIBpos_r16)
				if err := ie.PosSIB_TypeAndInfo_r16[i].PosSib2_20_r16.Decode(d); err != nil {
					return err
				}
			case PosSystemInformation_r16_IEs_PosSIB_TypeAndInfo_r16_Elem_PosSib2_21_r16:
				ie.PosSIB_TypeAndInfo_r16[i].PosSib2_21_r16 = new(SIBpos_r16)
				if err := ie.PosSIB_TypeAndInfo_r16[i].PosSib2_21_r16.Decode(d); err != nil {
					return err
				}
			case PosSystemInformation_r16_IEs_PosSIB_TypeAndInfo_r16_Elem_PosSib2_22_r16:
				ie.PosSIB_TypeAndInfo_r16[i].PosSib2_22_r16 = new(SIBpos_r16)
				if err := ie.PosSIB_TypeAndInfo_r16[i].PosSib2_22_r16.Decode(d); err != nil {
					return err
				}
			case PosSystemInformation_r16_IEs_PosSIB_TypeAndInfo_r16_Elem_PosSib2_23_r16:
				ie.PosSIB_TypeAndInfo_r16[i].PosSib2_23_r16 = new(SIBpos_r16)
				if err := ie.PosSIB_TypeAndInfo_r16[i].PosSib2_23_r16.Decode(d); err != nil {
					return err
				}
			case PosSystemInformation_r16_IEs_PosSIB_TypeAndInfo_r16_Elem_PosSib3_1_r16:
				ie.PosSIB_TypeAndInfo_r16[i].PosSib3_1_r16 = new(SIBpos_r16)
				if err := ie.PosSIB_TypeAndInfo_r16[i].PosSib3_1_r16.Decode(d); err != nil {
					return err
				}
			case PosSystemInformation_r16_IEs_PosSIB_TypeAndInfo_r16_Elem_PosSib4_1_r16:
				ie.PosSIB_TypeAndInfo_r16[i].PosSib4_1_r16 = new(SIBpos_r16)
				if err := ie.PosSIB_TypeAndInfo_r16[i].PosSib4_1_r16.Decode(d); err != nil {
					return err
				}
			case PosSystemInformation_r16_IEs_PosSIB_TypeAndInfo_r16_Elem_PosSib5_1_r16:
				ie.PosSIB_TypeAndInfo_r16[i].PosSib5_1_r16 = new(SIBpos_r16)
				if err := ie.PosSIB_TypeAndInfo_r16[i].PosSib5_1_r16.Decode(d); err != nil {
					return err
				}
			case PosSystemInformation_r16_IEs_PosSIB_TypeAndInfo_r16_Elem_PosSib6_1_r16:
				ie.PosSIB_TypeAndInfo_r16[i].PosSib6_1_r16 = new(SIBpos_r16)
				if err := ie.PosSIB_TypeAndInfo_r16[i].PosSib6_1_r16.Decode(d); err != nil {
					return err
				}
			case PosSystemInformation_r16_IEs_PosSIB_TypeAndInfo_r16_Elem_PosSib6_2_r16:
				ie.PosSIB_TypeAndInfo_r16[i].PosSib6_2_r16 = new(SIBpos_r16)
				if err := ie.PosSIB_TypeAndInfo_r16[i].PosSib6_2_r16.Decode(d); err != nil {
					return err
				}
			case PosSystemInformation_r16_IEs_PosSIB_TypeAndInfo_r16_Elem_PosSib6_3_r16:
				ie.PosSIB_TypeAndInfo_r16[i].PosSib6_3_r16 = new(SIBpos_r16)
				if err := ie.PosSIB_TypeAndInfo_r16[i].PosSib6_3_r16.Decode(d); err != nil {
					return err
				}
			}
		}
	}

	// 3. lateNonCriticalExtension: octet-string
	{
		if seq.IsComponentPresent(1) {
			v, err := d.DecodeOctetString(posSystemInformationR16IEsLateNonCriticalExtensionConstraints)
			if err != nil {
				return err
			}
			ie.LateNonCriticalExtension = v
		}
	}

	// 4. nonCriticalExtension: sequence
	{
		if seq.IsComponentPresent(2) {
			ie.NonCriticalExtension = &struct{}{}
		}
	}

	return nil
}
