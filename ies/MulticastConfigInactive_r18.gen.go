// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: MulticastConfigInactive-r18 (line 1521).

var multicastConfigInactiveR18Constraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "inactivePTM-Config-r18", Optional: true},
		{Name: "inactiveMCCH-Config-r18", Optional: true},
	},
}

var multicastConfigInactiveR18InactivePTMConfigR18Constraints = per.SizeConstraints{}

var multicastConfigInactiveR18InactiveMCCHConfigR18Constraints = per.SizeConstraints{}

type MulticastConfigInactive_r18 struct {
	InactivePTM_Config_r18  []byte
	InactiveMCCH_Config_r18 []byte
}

func (ie *MulticastConfigInactive_r18) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(multicastConfigInactiveR18Constraints)

	// 1. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.InactivePTM_Config_r18 != nil, ie.InactiveMCCH_Config_r18 != nil}); err != nil {
		return err
	}

	// 2. inactivePTM-Config-r18: octet-string
	{
		if ie.InactivePTM_Config_r18 != nil {
			if err := e.EncodeOctetString(ie.InactivePTM_Config_r18, multicastConfigInactiveR18InactivePTMConfigR18Constraints); err != nil {
				return err
			}
		}
	}

	// 3. inactiveMCCH-Config-r18: octet-string
	{
		if ie.InactiveMCCH_Config_r18 != nil {
			if err := e.EncodeOctetString(ie.InactiveMCCH_Config_r18, multicastConfigInactiveR18InactiveMCCHConfigR18Constraints); err != nil {
				return err
			}
		}
	}

	return nil
}

func (ie *MulticastConfigInactive_r18) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(multicastConfigInactiveR18Constraints)

	// 1. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 2. inactivePTM-Config-r18: octet-string
	{
		if seq.IsComponentPresent(0) {
			v, err := d.DecodeOctetString(multicastConfigInactiveR18InactivePTMConfigR18Constraints)
			if err != nil {
				return err
			}
			ie.InactivePTM_Config_r18 = v
		}
	}

	// 3. inactiveMCCH-Config-r18: octet-string
	{
		if seq.IsComponentPresent(1) {
			v, err := d.DecodeOctetString(multicastConfigInactiveR18InactiveMCCHConfigR18Constraints)
			if err != nil {
				return err
			}
			ie.InactiveMCCH_Config_r18 = v
		}
	}

	return nil
}
