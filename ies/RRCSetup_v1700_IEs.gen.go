// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: RRCSetup-v1700-IEs (line 1698).

var rRCSetupV1700IEsConstraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "sl-ConfigDedicatedNR-r17", Optional: true},
		{Name: "sl-L2RemoteUE-Config-r17", Optional: true},
		{Name: "nonCriticalExtension", Optional: true},
	},
}

type RRCSetup_v1700_IEs struct {
	Sl_ConfigDedicatedNR_r17 *SL_ConfigDedicatedNR_r16
	Sl_L2RemoteUE_Config_r17 *SL_L2RemoteUE_Config_r17
	NonCriticalExtension     *struct{}
}

func (ie *RRCSetup_v1700_IEs) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(rRCSetupV1700IEsConstraints)

	// 1. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.Sl_ConfigDedicatedNR_r17 != nil, ie.Sl_L2RemoteUE_Config_r17 != nil, ie.NonCriticalExtension != nil}); err != nil {
		return err
	}

	// 2. sl-ConfigDedicatedNR-r17: ref
	{
		if ie.Sl_ConfigDedicatedNR_r17 != nil {
			if err := ie.Sl_ConfigDedicatedNR_r17.Encode(e); err != nil {
				return err
			}
		}
	}

	// 3. sl-L2RemoteUE-Config-r17: ref
	{
		if ie.Sl_L2RemoteUE_Config_r17 != nil {
			if err := ie.Sl_L2RemoteUE_Config_r17.Encode(e); err != nil {
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

func (ie *RRCSetup_v1700_IEs) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(rRCSetupV1700IEsConstraints)

	// 1. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 2. sl-ConfigDedicatedNR-r17: ref
	{
		if seq.IsComponentPresent(0) {
			ie.Sl_ConfigDedicatedNR_r17 = new(SL_ConfigDedicatedNR_r16)
			if err := ie.Sl_ConfigDedicatedNR_r17.Decode(d); err != nil {
				return err
			}
		}
	}

	// 3. sl-L2RemoteUE-Config-r17: ref
	{
		if seq.IsComponentPresent(1) {
			ie.Sl_L2RemoteUE_Config_r17 = new(SL_L2RemoteUE_Config_r17)
			if err := ie.Sl_L2RemoteUE_Config_r17.Decode(d); err != nil {
				return err
			}
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
