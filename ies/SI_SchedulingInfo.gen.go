// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/rrc/common"
)

// Source: SI-SchedulingInfo (line 15051).

var sISchedulingInfoConstraints = per.SequenceConstraints{
	Extensible: true,
	RootComponents: []per.ComponentInfo{
		{Name: "schedulingInfoList"},
		{Name: "si-WindowLength"},
		{Name: "si-RequestConfig", Optional: true},
		{Name: "si-RequestConfigSUL", Optional: true},
		{Name: "systemInformationAreaID", Optional: true},
	},
}

var sISchedulingInfoSchedulingInfoListConstraints = per.SizeRange(1, common.MaxSI_Message)

const (
	SI_SchedulingInfo_Si_WindowLength_S5          = 0
	SI_SchedulingInfo_Si_WindowLength_S10         = 1
	SI_SchedulingInfo_Si_WindowLength_S20         = 2
	SI_SchedulingInfo_Si_WindowLength_S40         = 3
	SI_SchedulingInfo_Si_WindowLength_S80         = 4
	SI_SchedulingInfo_Si_WindowLength_S160        = 5
	SI_SchedulingInfo_Si_WindowLength_S320        = 6
	SI_SchedulingInfo_Si_WindowLength_S640        = 7
	SI_SchedulingInfo_Si_WindowLength_S1280       = 8
	SI_SchedulingInfo_Si_WindowLength_S2560_v1710 = 9
	SI_SchedulingInfo_Si_WindowLength_S5120_v1710 = 10
)

var sISchedulingInfoSiWindowLengthConstraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{SI_SchedulingInfo_Si_WindowLength_S5, SI_SchedulingInfo_Si_WindowLength_S10, SI_SchedulingInfo_Si_WindowLength_S20, SI_SchedulingInfo_Si_WindowLength_S40, SI_SchedulingInfo_Si_WindowLength_S80, SI_SchedulingInfo_Si_WindowLength_S160, SI_SchedulingInfo_Si_WindowLength_S320, SI_SchedulingInfo_Si_WindowLength_S640, SI_SchedulingInfo_Si_WindowLength_S1280, SI_SchedulingInfo_Si_WindowLength_S2560_v1710, SI_SchedulingInfo_Si_WindowLength_S5120_v1710},
}

var sISchedulingInfoSystemInformationAreaIDConstraints = per.FixedSize(24)

type SI_SchedulingInfo struct {
	SchedulingInfoList      []SchedulingInfo
	Si_WindowLength         int64
	Si_RequestConfig        *SI_RequestConfig
	Si_RequestConfigSUL     *SI_RequestConfig
	SystemInformationAreaID *per.BitString
}

func (ie *SI_SchedulingInfo) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(sISchedulingInfoConstraints)

	// 1. SEQUENCE extension bit.
	if err := seq.EncodeExtensionBit(false); err != nil {
		return err
	}

	// 2. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.Si_RequestConfig != nil, ie.Si_RequestConfigSUL != nil, ie.SystemInformationAreaID != nil}); err != nil {
		return err
	}

	// 3. schedulingInfoList: sequence-of
	{
		seqOf := e.NewSequenceOfEncoder(sISchedulingInfoSchedulingInfoListConstraints)
		if err := seqOf.EncodeLength(int64(len(ie.SchedulingInfoList))); err != nil {
			return err
		}
		for i := range ie.SchedulingInfoList {
			if err := ie.SchedulingInfoList[i].Encode(e); err != nil {
				return err
			}
		}
	}

	// 4. si-WindowLength: enumerated
	{
		if err := e.EncodeEnumerated(ie.Si_WindowLength, sISchedulingInfoSiWindowLengthConstraints); err != nil {
			return err
		}
	}

	// 5. si-RequestConfig: ref
	{
		if ie.Si_RequestConfig != nil {
			if err := ie.Si_RequestConfig.Encode(e); err != nil {
				return err
			}
		}
	}

	// 6. si-RequestConfigSUL: ref
	{
		if ie.Si_RequestConfigSUL != nil {
			if err := ie.Si_RequestConfigSUL.Encode(e); err != nil {
				return err
			}
		}
	}

	// 7. systemInformationAreaID: bit-string
	{
		if ie.SystemInformationAreaID != nil {
			if err := e.EncodeBitString(*ie.SystemInformationAreaID, sISchedulingInfoSystemInformationAreaIDConstraints); err != nil {
				return err
			}
		}
	}

	return nil
}

func (ie *SI_SchedulingInfo) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(sISchedulingInfoConstraints)

	// 1. SEQUENCE extension bit.
	if err := seq.DecodeExtensionBit(); err != nil {
		return err
	}

	// 2. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 3. schedulingInfoList: sequence-of
	{
		seqOf := d.NewSequenceOfDecoder(sISchedulingInfoSchedulingInfoListConstraints)
		n, err := seqOf.DecodeLength()
		if err != nil {
			return err
		}
		ie.SchedulingInfoList = make([]SchedulingInfo, n)
		for i := int64(0); i < n; i++ {
			if err := ie.SchedulingInfoList[i].Decode(d); err != nil {
				return err
			}
		}
	}

	// 4. si-WindowLength: enumerated
	{
		v1, err := d.DecodeEnumerated(sISchedulingInfoSiWindowLengthConstraints)
		if err != nil {
			return err
		}
		ie.Si_WindowLength = v1
	}

	// 5. si-RequestConfig: ref
	{
		if seq.IsComponentPresent(2) {
			ie.Si_RequestConfig = new(SI_RequestConfig)
			if err := ie.Si_RequestConfig.Decode(d); err != nil {
				return err
			}
		}
	}

	// 6. si-RequestConfigSUL: ref
	{
		if seq.IsComponentPresent(3) {
			ie.Si_RequestConfigSUL = new(SI_RequestConfig)
			if err := ie.Si_RequestConfigSUL.Decode(d); err != nil {
				return err
			}
		}
	}

	// 7. systemInformationAreaID: bit-string
	{
		if seq.IsComponentPresent(4) {
			v, err := d.DecodeBitString(sISchedulingInfoSystemInformationAreaIDConstraints)
			if err != nil {
				return err
			}
			ie.SystemInformationAreaID = &v
		}
	}

	return nil
}
