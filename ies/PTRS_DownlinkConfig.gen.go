// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: PTRS-DownlinkConfig (line 11932).

var pTRSDownlinkConfigConstraints = per.SequenceConstraints{
	Extensible: true,
	RootComponents: []per.ComponentInfo{
		{Name: "frequencyDensity", Optional: true},
		{Name: "timeDensity", Optional: true},
		{Name: "epre-Ratio", Optional: true},
		{Name: "resourceElementOffset", Optional: true},
	},
	ExtComponents: []per.ComponentInfo{
		{Name: "extension-addition-0", Optional: true},
	},
}

var pTRSDownlinkConfigFrequencyDensityConstraints = per.FixedSize(2)

var pTRSDownlinkConfigTimeDensityConstraints = per.FixedSize(3)

var pTRSDownlinkConfigEpreRatioConstraints = per.Constrained(0, 3)

const (
	PTRS_DownlinkConfig_ResourceElementOffset_Offset01 = 0
	PTRS_DownlinkConfig_ResourceElementOffset_Offset10 = 1
	PTRS_DownlinkConfig_ResourceElementOffset_Offset11 = 2
)

var pTRSDownlinkConfigResourceElementOffsetConstraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{PTRS_DownlinkConfig_ResourceElementOffset_Offset01, PTRS_DownlinkConfig_ResourceElementOffset_Offset10, PTRS_DownlinkConfig_ResourceElementOffset_Offset11},
}

const (
	PTRS_DownlinkConfig_Ext_MaxNrofPorts_r16_N1 = 0
	PTRS_DownlinkConfig_Ext_MaxNrofPorts_r16_N2 = 1
)

var pTRSDownlinkConfigExtMaxNrofPortsR16Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{PTRS_DownlinkConfig_Ext_MaxNrofPorts_r16_N1, PTRS_DownlinkConfig_Ext_MaxNrofPorts_r16_N2},
}

type PTRS_DownlinkConfig struct {
	FrequencyDensity      []int64
	TimeDensity           []int64
	Epre_Ratio            *int64
	ResourceElementOffset *int64
	MaxNrofPorts_r16      *int64
}

func (ie *PTRS_DownlinkConfig) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(pTRSDownlinkConfigConstraints)

	// Extension-group presence (one var per [[...]] group).
	hasExtGroup0 := ie.MaxNrofPorts_r16 != nil
	hasExtensions := hasExtGroup0

	// 1. SEQUENCE extension bit.
	if err := seq.EncodeExtensionBit(hasExtensions); err != nil {
		return err
	}

	// 2. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.FrequencyDensity != nil, ie.TimeDensity != nil, ie.Epre_Ratio != nil, ie.ResourceElementOffset != nil}); err != nil {
		return err
	}

	// 3. frequencyDensity: sequence-of
	{
		if ie.FrequencyDensity != nil {
			seqOf := e.NewSequenceOfEncoder(pTRSDownlinkConfigFrequencyDensityConstraints)
			if err := seqOf.EncodeLength(int64(len(ie.FrequencyDensity))); err != nil {
				return err
			}
			for i := range ie.FrequencyDensity {
				if err := e.EncodeInteger(ie.FrequencyDensity[i], per.Constrained(1, 276)); err != nil {
					return err
				}
			}
		}
	}

	// 4. timeDensity: sequence-of
	{
		if ie.TimeDensity != nil {
			seqOf := e.NewSequenceOfEncoder(pTRSDownlinkConfigTimeDensityConstraints)
			if err := seqOf.EncodeLength(int64(len(ie.TimeDensity))); err != nil {
				return err
			}
			for i := range ie.TimeDensity {
				if err := e.EncodeInteger(ie.TimeDensity[i], per.Constrained(0, 29)); err != nil {
					return err
				}
			}
		}
	}

	// 5. epre-Ratio: integer
	{
		if ie.Epre_Ratio != nil {
			if err := e.EncodeInteger(*ie.Epre_Ratio, pTRSDownlinkConfigEpreRatioConstraints); err != nil {
				return err
			}
		}
	}

	// 6. resourceElementOffset: enumerated
	{
		if ie.ResourceElementOffset != nil {
			if err := e.EncodeEnumerated(*ie.ResourceElementOffset, pTRSDownlinkConfigResourceElementOffsetConstraints); err != nil {
				return err
			}
		}
	}

	if hasExtensions {
		extPresence := []bool{hasExtGroup0}
		var extValues [][]byte

		// Extension group 0:
		if hasExtGroup0 {
			ex := per.NewEncoder(e.Variant())
			groupSeq := ex.NewSequenceEncoder(per.SequenceConstraints{
				Extensible: false,
				RootComponents: []per.ComponentInfo{
					{Name: "maxNrofPorts-r16", Optional: true},
				},
			})
			if err := groupSeq.EncodePreamble([]bool{ie.MaxNrofPorts_r16 != nil}); err != nil {
				return err
			}

			if ie.MaxNrofPorts_r16 != nil {
				if err := ex.EncodeEnumerated(*ie.MaxNrofPorts_r16, pTRSDownlinkConfigExtMaxNrofPortsR16Constraints); err != nil {
					return err
				}
			}
			extValues = append(extValues, ex.Bytes())
		}

		if err := seq.EncodeExtensionAdditions(extPresence, extValues); err != nil {
			return err
		}
	}

	return nil
}

func (ie *PTRS_DownlinkConfig) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(pTRSDownlinkConfigConstraints)

	// 1. SEQUENCE extension bit.
	if err := seq.DecodeExtensionBit(); err != nil {
		return err
	}

	// 2. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 3. frequencyDensity: sequence-of
	{
		if seq.IsComponentPresent(0) {
			seqOf := d.NewSequenceOfDecoder(pTRSDownlinkConfigFrequencyDensityConstraints)
			n, err := seqOf.DecodeLength()
			if err != nil {
				return err
			}
			ie.FrequencyDensity = make([]int64, n)
			for i := int64(0); i < n; i++ {
				v, err := d.DecodeInteger(per.Constrained(1, 276))
				if err != nil {
					return err
				}
				ie.FrequencyDensity[i] = v
			}
		}
	}

	// 4. timeDensity: sequence-of
	{
		if seq.IsComponentPresent(1) {
			seqOf := d.NewSequenceOfDecoder(pTRSDownlinkConfigTimeDensityConstraints)
			n, err := seqOf.DecodeLength()
			if err != nil {
				return err
			}
			ie.TimeDensity = make([]int64, n)
			for i := int64(0); i < n; i++ {
				v, err := d.DecodeInteger(per.Constrained(0, 29))
				if err != nil {
					return err
				}
				ie.TimeDensity[i] = v
			}
		}
	}

	// 5. epre-Ratio: integer
	{
		if seq.IsComponentPresent(2) {
			v, err := d.DecodeInteger(pTRSDownlinkConfigEpreRatioConstraints)
			if err != nil {
				return err
			}
			ie.Epre_Ratio = &v
		}
	}

	// 6. resourceElementOffset: enumerated
	{
		if seq.IsComponentPresent(3) {
			idx, err := d.DecodeEnumerated(pTRSDownlinkConfigResourceElementOffsetConstraints)
			if err != nil {
				return err
			}
			ie.ResourceElementOffset = &idx
		}
	}

	// Extension additions as open types.
	extValues, err := seq.DecodeExtensionAdditions()
	if err != nil {
		return err
	}
	extValueIndex := 0

	// Extension group 0:
	if seq.IsExtensionPresent(0) {
		dx := per.NewDecoder(extValues[extValueIndex], d.Variant())
		extValueIndex++
		groupSeq := dx.NewSequenceDecoder(per.SequenceConstraints{
			Extensible: false,
			RootComponents: []per.ComponentInfo{
				{Name: "maxNrofPorts-r16", Optional: true},
			},
		})
		if err := groupSeq.DecodePreamble(); err != nil {
			return err
		}

		if groupSeq.IsComponentPresent(0) {
			v, err := dx.DecodeEnumerated(pTRSDownlinkConfigExtMaxNrofPortsR16Constraints)
			if err != nil {
				return err
			}
			ie.MaxNrofPorts_r16 = &v
		}
	}

	return nil
}
