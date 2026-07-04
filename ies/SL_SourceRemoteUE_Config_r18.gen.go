// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: SL-SourceRemoteUE-Config-r18 (line 27460).

var sLSourceRemoteUEConfigR18Constraints = per.SequenceConstraints{
	Extensible: true,
	RootComponents: []per.ComponentInfo{
		{Name: "sl-SourceUE-Identity-r18"},
		{Name: "sl-SRAP-ConfigU2U-r18"},
	},
}

type SL_SourceRemoteUE_Config_r18 struct {
	Sl_SourceUE_Identity_r18 SL_SourceIdentity_r17
	Sl_SRAP_ConfigU2U_r18    SL_SRAP_ConfigU2U_r18
}

func (ie *SL_SourceRemoteUE_Config_r18) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(sLSourceRemoteUEConfigR18Constraints)

	// 1. SEQUENCE extension bit.
	if err := seq.EncodeExtensionBit(false); err != nil {
		return err
	}

	// 2. sl-SourceUE-Identity-r18: ref
	{
		if err := ie.Sl_SourceUE_Identity_r18.Encode(e); err != nil {
			return err
		}
	}

	// 3. sl-SRAP-ConfigU2U-r18: ref
	{
		if err := ie.Sl_SRAP_ConfigU2U_r18.Encode(e); err != nil {
			return err
		}
	}

	return nil
}

func (ie *SL_SourceRemoteUE_Config_r18) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(sLSourceRemoteUEConfigR18Constraints)

	// 1. SEQUENCE extension bit.
	if err := seq.DecodeExtensionBit(); err != nil {
		return err
	}

	// 2. sl-SourceUE-Identity-r18: ref
	{
		if err := ie.Sl_SourceUE_Identity_r18.Decode(d); err != nil {
			return err
		}
	}

	// 3. sl-SRAP-ConfigU2U-r18: ref
	{
		if err := ie.Sl_SRAP_ConfigU2U_r18.Decode(d); err != nil {
			return err
		}
	}

	return nil
}
