// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: FrequencyComponent-r17 (line 16414).
// FrequencyComponent-r17 ::=  ENUMERATED {activeCarrier,configuredCarrier,activeBWP,configuredBWP}

const (
	FrequencyComponent_r17_ActiveCarrier     = 0
	FrequencyComponent_r17_ConfiguredCarrier = 1
	FrequencyComponent_r17_ActiveBWP         = 2
	FrequencyComponent_r17_ConfiguredBWP     = 3
)

var frequencyComponentR17Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{FrequencyComponent_r17_ActiveCarrier, FrequencyComponent_r17_ConfiguredCarrier, FrequencyComponent_r17_ActiveBWP, FrequencyComponent_r17_ConfiguredBWP},
}

type FrequencyComponent_r17 struct {
	Value int64
}

func (v *FrequencyComponent_r17) Encode(e *per.Encoder) error {
	return e.EncodeEnumerated(v.Value, frequencyComponentR17Constraints)
}

func (v *FrequencyComponent_r17) Decode(d *per.Decoder) error {
	x, err := d.DecodeEnumerated(frequencyComponentR17Constraints)
	if err != nil {
		return err
	}
	v.Value = x
	return nil
}
