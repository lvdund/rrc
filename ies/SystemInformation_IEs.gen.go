// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/rrc/common"
)

// Source: SystemInformation-IEs (line 2351).

var systemInformationIEsConstraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "sib-TypeAndInfo"},
		{Name: "lateNonCriticalExtension", Optional: true},
		{Name: "nonCriticalExtension", Optional: true},
	},
}

var systemInformationIEsSibTypeAndInfoConstraints = per.SizeRange(1, common.MaxSIB)

var systemInformationIEsLateNonCriticalExtensionConstraints = per.SizeConstraints{}

var systemInformationIEsSibTypeAndInfoElemConstraints = per.ChoiceConstraints{
	Extensible: true,
	RootAlternatives: []per.AlternativeInfo{
		{Name: "sib2"},
		{Name: "sib3"},
		{Name: "sib4"},
		{Name: "sib5"},
		{Name: "sib6"},
		{Name: "sib7"},
		{Name: "sib8"},
		{Name: "sib9"},
	},
	ExtAlternatives: []per.AlternativeInfo{
		{Name: "sib10-v1610"},
		{Name: "sib11-v1610"},
		{Name: "sib12-v1610"},
		{Name: "sib13-v1610"},
		{Name: "sib14-v1610"},
		{Name: "sib15-v1700"},
		{Name: "sib16-v1700"},
		{Name: "sib17-v1700"},
		{Name: "sib18-v1700"},
		{Name: "sib19-v1700"},
		{Name: "sib20-v1700"},
		{Name: "sib21-v1700"},
		{Name: "sib22-v1800"},
		{Name: "sib23-v1800"},
		{Name: "sib24-v1800"},
		{Name: "sib25-v1800"},
		{Name: "sib17bis-v1820"},
		{Name: "sib26-v1900"},
		{Name: "sib27-v1900"},
		{Name: "sib28-v1910"},
	},
}

const (
	SystemInformation_IEs_Sib_TypeAndInfo_Elem_Sib2 = 0
	SystemInformation_IEs_Sib_TypeAndInfo_Elem_Sib3 = 1
	SystemInformation_IEs_Sib_TypeAndInfo_Elem_Sib4 = 2
	SystemInformation_IEs_Sib_TypeAndInfo_Elem_Sib5 = 3
	SystemInformation_IEs_Sib_TypeAndInfo_Elem_Sib6 = 4
	SystemInformation_IEs_Sib_TypeAndInfo_Elem_Sib7 = 5
	SystemInformation_IEs_Sib_TypeAndInfo_Elem_Sib8 = 6
	SystemInformation_IEs_Sib_TypeAndInfo_Elem_Sib9 = 7
)

type SystemInformation_IEs struct {
	Sib_TypeAndInfo []struct {
		Choice int
		Sib2   *SIB2
		Sib3   *SIB3
		Sib4   *SIB4
		Sib5   *SIB5
		Sib6   *SIB6
		Sib7   *SIB7
		Sib8   *SIB8
		Sib9   *SIB9
	}
	LateNonCriticalExtension []byte
	NonCriticalExtension     *struct{}
}

func (ie *SystemInformation_IEs) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(systemInformationIEsConstraints)

	// 1. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.LateNonCriticalExtension != nil, ie.NonCriticalExtension != nil}); err != nil {
		return err
	}

	// 2. sib-TypeAndInfo: sequence-of
	{
		seqOf := e.NewSequenceOfEncoder(systemInformationIEsSibTypeAndInfoConstraints)
		if err := seqOf.EncodeLength(int64(len(ie.Sib_TypeAndInfo))); err != nil {
			return err
		}
		for i := range ie.Sib_TypeAndInfo {
			choiceEnc := e.NewChoiceEncoder(systemInformationIEsSibTypeAndInfoElemConstraints)
			if err := choiceEnc.EncodeChoice(int64(ie.Sib_TypeAndInfo[i].Choice), false, nil); err != nil {
				return err
			}
			switch ie.Sib_TypeAndInfo[i].Choice {
			case SystemInformation_IEs_Sib_TypeAndInfo_Elem_Sib2:
				if err := ie.Sib_TypeAndInfo[i].Sib2.Encode(e); err != nil {
					return err
				}
			case SystemInformation_IEs_Sib_TypeAndInfo_Elem_Sib3:
				if err := ie.Sib_TypeAndInfo[i].Sib3.Encode(e); err != nil {
					return err
				}
			case SystemInformation_IEs_Sib_TypeAndInfo_Elem_Sib4:
				if err := ie.Sib_TypeAndInfo[i].Sib4.Encode(e); err != nil {
					return err
				}
			case SystemInformation_IEs_Sib_TypeAndInfo_Elem_Sib5:
				if err := ie.Sib_TypeAndInfo[i].Sib5.Encode(e); err != nil {
					return err
				}
			case SystemInformation_IEs_Sib_TypeAndInfo_Elem_Sib6:
				if err := ie.Sib_TypeAndInfo[i].Sib6.Encode(e); err != nil {
					return err
				}
			case SystemInformation_IEs_Sib_TypeAndInfo_Elem_Sib7:
				if err := ie.Sib_TypeAndInfo[i].Sib7.Encode(e); err != nil {
					return err
				}
			case SystemInformation_IEs_Sib_TypeAndInfo_Elem_Sib8:
				if err := ie.Sib_TypeAndInfo[i].Sib8.Encode(e); err != nil {
					return err
				}
			case SystemInformation_IEs_Sib_TypeAndInfo_Elem_Sib9:
				if err := ie.Sib_TypeAndInfo[i].Sib9.Encode(e); err != nil {
					return err
				}
			}
		}
	}

	// 3. lateNonCriticalExtension: octet-string
	{
		if ie.LateNonCriticalExtension != nil {
			if err := e.EncodeOctetString(ie.LateNonCriticalExtension, systemInformationIEsLateNonCriticalExtensionConstraints); err != nil {
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

func (ie *SystemInformation_IEs) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(systemInformationIEsConstraints)

	// 1. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 2. sib-TypeAndInfo: sequence-of
	{
		seqOf := d.NewSequenceOfDecoder(systemInformationIEsSibTypeAndInfoConstraints)
		n, err := seqOf.DecodeLength()
		if err != nil {
			return err
		}
		ie.Sib_TypeAndInfo = make([]struct {
			Choice int
			Sib2   *SIB2
			Sib3   *SIB3
			Sib4   *SIB4
			Sib5   *SIB5
			Sib6   *SIB6
			Sib7   *SIB7
			Sib8   *SIB8
			Sib9   *SIB9
		}, n)
		for i := int64(0); i < n; i++ {
			choiceDec := d.NewChoiceDecoder(systemInformationIEsSibTypeAndInfoElemConstraints)
			index, _, _, err := choiceDec.DecodeChoice()
			if err != nil {
				return err
			}
			ie.Sib_TypeAndInfo[i].Choice = int(index)
			switch index {
			case SystemInformation_IEs_Sib_TypeAndInfo_Elem_Sib2:
				ie.Sib_TypeAndInfo[i].Sib2 = new(SIB2)
				if err := ie.Sib_TypeAndInfo[i].Sib2.Decode(d); err != nil {
					return err
				}
			case SystemInformation_IEs_Sib_TypeAndInfo_Elem_Sib3:
				ie.Sib_TypeAndInfo[i].Sib3 = new(SIB3)
				if err := ie.Sib_TypeAndInfo[i].Sib3.Decode(d); err != nil {
					return err
				}
			case SystemInformation_IEs_Sib_TypeAndInfo_Elem_Sib4:
				ie.Sib_TypeAndInfo[i].Sib4 = new(SIB4)
				if err := ie.Sib_TypeAndInfo[i].Sib4.Decode(d); err != nil {
					return err
				}
			case SystemInformation_IEs_Sib_TypeAndInfo_Elem_Sib5:
				ie.Sib_TypeAndInfo[i].Sib5 = new(SIB5)
				if err := ie.Sib_TypeAndInfo[i].Sib5.Decode(d); err != nil {
					return err
				}
			case SystemInformation_IEs_Sib_TypeAndInfo_Elem_Sib6:
				ie.Sib_TypeAndInfo[i].Sib6 = new(SIB6)
				if err := ie.Sib_TypeAndInfo[i].Sib6.Decode(d); err != nil {
					return err
				}
			case SystemInformation_IEs_Sib_TypeAndInfo_Elem_Sib7:
				ie.Sib_TypeAndInfo[i].Sib7 = new(SIB7)
				if err := ie.Sib_TypeAndInfo[i].Sib7.Decode(d); err != nil {
					return err
				}
			case SystemInformation_IEs_Sib_TypeAndInfo_Elem_Sib8:
				ie.Sib_TypeAndInfo[i].Sib8 = new(SIB8)
				if err := ie.Sib_TypeAndInfo[i].Sib8.Decode(d); err != nil {
					return err
				}
			case SystemInformation_IEs_Sib_TypeAndInfo_Elem_Sib9:
				ie.Sib_TypeAndInfo[i].Sib9 = new(SIB9)
				if err := ie.Sib_TypeAndInfo[i].Sib9.Decode(d); err != nil {
					return err
				}
			}
		}
	}

	// 3. lateNonCriticalExtension: octet-string
	{
		if seq.IsComponentPresent(1) {
			v, err := d.DecodeOctetString(systemInformationIEsLateNonCriticalExtensionConstraints)
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
