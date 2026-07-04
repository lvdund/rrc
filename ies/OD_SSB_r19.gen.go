// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/rrc/common"
)

// Source: OD-SSB-r19 (line 5850).

var oDSSBR19Constraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "od-SSB-AbsoluteFrequency-r19", Optional: true},
		{Name: "od-SSB-SubcarrierSpacing-r19", Optional: true},
		{Name: "od-SSB-PBCH-BlockPower-r19", Optional: true},
		{Name: "od-SSB-ConfigToAddModList-r19", Optional: true},
		{Name: "od-SSB-ConfigToReleaseList-r19", Optional: true},
	},
}

var oDSSBR19OdSSBPBCHBlockPowerR19Constraints = per.Constrained(-60, 50)

var oDSSBR19OdSSBConfigToAddModListR19Constraints = per.SizeRange(1, common.MaxNrofOD_SSB_r19)

var oDSSBR19OdSSBConfigToReleaseListR19Constraints = per.SizeRange(1, common.MaxNrofOD_SSB_r19)

type OD_SSB_r19 struct {
	Od_SSB_AbsoluteFrequency_r19   *ARFCN_ValueNR
	Od_SSB_SubcarrierSpacing_r19   *SubcarrierSpacing
	Od_SSB_PBCH_BlockPower_r19     *int64
	Od_SSB_ConfigToAddModList_r19  []OD_SSB_Config_r19
	Od_SSB_ConfigToReleaseList_r19 []OD_SSB_ConfigId_r19
}

func (ie *OD_SSB_r19) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(oDSSBR19Constraints)

	// 1. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.Od_SSB_AbsoluteFrequency_r19 != nil, ie.Od_SSB_SubcarrierSpacing_r19 != nil, ie.Od_SSB_PBCH_BlockPower_r19 != nil, ie.Od_SSB_ConfigToAddModList_r19 != nil, ie.Od_SSB_ConfigToReleaseList_r19 != nil}); err != nil {
		return err
	}

	// 2. od-SSB-AbsoluteFrequency-r19: ref
	{
		if ie.Od_SSB_AbsoluteFrequency_r19 != nil {
			if err := ie.Od_SSB_AbsoluteFrequency_r19.Encode(e); err != nil {
				return err
			}
		}
	}

	// 3. od-SSB-SubcarrierSpacing-r19: ref
	{
		if ie.Od_SSB_SubcarrierSpacing_r19 != nil {
			if err := ie.Od_SSB_SubcarrierSpacing_r19.Encode(e); err != nil {
				return err
			}
		}
	}

	// 4. od-SSB-PBCH-BlockPower-r19: integer
	{
		if ie.Od_SSB_PBCH_BlockPower_r19 != nil {
			if err := e.EncodeInteger(*ie.Od_SSB_PBCH_BlockPower_r19, oDSSBR19OdSSBPBCHBlockPowerR19Constraints); err != nil {
				return err
			}
		}
	}

	// 5. od-SSB-ConfigToAddModList-r19: sequence-of
	{
		if ie.Od_SSB_ConfigToAddModList_r19 != nil {
			seqOf := e.NewSequenceOfEncoder(oDSSBR19OdSSBConfigToAddModListR19Constraints)
			if err := seqOf.EncodeLength(int64(len(ie.Od_SSB_ConfigToAddModList_r19))); err != nil {
				return err
			}
			for i := range ie.Od_SSB_ConfigToAddModList_r19 {
				if err := ie.Od_SSB_ConfigToAddModList_r19[i].Encode(e); err != nil {
					return err
				}
			}
		}
	}

	// 6. od-SSB-ConfigToReleaseList-r19: sequence-of
	{
		if ie.Od_SSB_ConfigToReleaseList_r19 != nil {
			seqOf := e.NewSequenceOfEncoder(oDSSBR19OdSSBConfigToReleaseListR19Constraints)
			if err := seqOf.EncodeLength(int64(len(ie.Od_SSB_ConfigToReleaseList_r19))); err != nil {
				return err
			}
			for i := range ie.Od_SSB_ConfigToReleaseList_r19 {
				if err := ie.Od_SSB_ConfigToReleaseList_r19[i].Encode(e); err != nil {
					return err
				}
			}
		}
	}

	return nil
}

func (ie *OD_SSB_r19) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(oDSSBR19Constraints)

	// 1. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 2. od-SSB-AbsoluteFrequency-r19: ref
	{
		if seq.IsComponentPresent(0) {
			ie.Od_SSB_AbsoluteFrequency_r19 = new(ARFCN_ValueNR)
			if err := ie.Od_SSB_AbsoluteFrequency_r19.Decode(d); err != nil {
				return err
			}
		}
	}

	// 3. od-SSB-SubcarrierSpacing-r19: ref
	{
		if seq.IsComponentPresent(1) {
			ie.Od_SSB_SubcarrierSpacing_r19 = new(SubcarrierSpacing)
			if err := ie.Od_SSB_SubcarrierSpacing_r19.Decode(d); err != nil {
				return err
			}
		}
	}

	// 4. od-SSB-PBCH-BlockPower-r19: integer
	{
		if seq.IsComponentPresent(2) {
			v, err := d.DecodeInteger(oDSSBR19OdSSBPBCHBlockPowerR19Constraints)
			if err != nil {
				return err
			}
			ie.Od_SSB_PBCH_BlockPower_r19 = &v
		}
	}

	// 5. od-SSB-ConfigToAddModList-r19: sequence-of
	{
		if seq.IsComponentPresent(3) {
			seqOf := d.NewSequenceOfDecoder(oDSSBR19OdSSBConfigToAddModListR19Constraints)
			n, err := seqOf.DecodeLength()
			if err != nil {
				return err
			}
			ie.Od_SSB_ConfigToAddModList_r19 = make([]OD_SSB_Config_r19, n)
			for i := int64(0); i < n; i++ {
				if err := ie.Od_SSB_ConfigToAddModList_r19[i].Decode(d); err != nil {
					return err
				}
			}
		}
	}

	// 6. od-SSB-ConfigToReleaseList-r19: sequence-of
	{
		if seq.IsComponentPresent(4) {
			seqOf := d.NewSequenceOfDecoder(oDSSBR19OdSSBConfigToReleaseListR19Constraints)
			n, err := seqOf.DecodeLength()
			if err != nil {
				return err
			}
			ie.Od_SSB_ConfigToReleaseList_r19 = make([]OD_SSB_ConfigId_r19, n)
			for i := int64(0); i < n; i++ {
				if err := ie.Od_SSB_ConfigToReleaseList_r19[i].Decode(d); err != nil {
					return err
				}
			}
		}
	}

	return nil
}
