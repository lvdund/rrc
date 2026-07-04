// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: PTRS-UplinkConfig (line 11947).

var pTRSUplinkConfigConstraints = per.SequenceConstraints{
	Extensible: true,
	RootComponents: []per.ComponentInfo{
		{Name: "transformPrecoderDisabled", Optional: true},
		{Name: "transformPrecoderEnabled", Optional: true},
	},
	ExtComponents: []per.ComponentInfo{
		{Name: "extension-addition-0", Optional: true},
	},
}

const (
	PTRS_UplinkConfig_Ext_MaxNrofPorts_SDM_r18_N1 = 0
	PTRS_UplinkConfig_Ext_MaxNrofPorts_SDM_r18_N2 = 1
)

var pTRSUplinkConfigExtMaxNrofPortsSDMR18Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{PTRS_UplinkConfig_Ext_MaxNrofPorts_SDM_r18_N1, PTRS_UplinkConfig_Ext_MaxNrofPorts_SDM_r18_N2},
}

var pTRSUplinkConfigTransformPrecoderDisabledConstraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "frequencyDensity", Optional: true},
		{Name: "timeDensity", Optional: true},
		{Name: "maxNrofPorts"},
		{Name: "resourceElementOffset", Optional: true},
		{Name: "ptrs-Power"},
	},
}

var pTRSUplinkConfigTransformPrecoderDisabledFrequencyDensityConstraints = per.FixedSize(2)

var pTRSUplinkConfigTransformPrecoderDisabledTimeDensityConstraints = per.FixedSize(3)

const (
	PTRS_UplinkConfig_TransformPrecoderDisabled_MaxNrofPorts_N1 = 0
	PTRS_UplinkConfig_TransformPrecoderDisabled_MaxNrofPorts_N2 = 1
)

var pTRSUplinkConfigTransformPrecoderDisabledMaxNrofPortsConstraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{PTRS_UplinkConfig_TransformPrecoderDisabled_MaxNrofPorts_N1, PTRS_UplinkConfig_TransformPrecoderDisabled_MaxNrofPorts_N2},
}

const (
	PTRS_UplinkConfig_TransformPrecoderDisabled_ResourceElementOffset_Offset01 = 0
	PTRS_UplinkConfig_TransformPrecoderDisabled_ResourceElementOffset_Offset10 = 1
	PTRS_UplinkConfig_TransformPrecoderDisabled_ResourceElementOffset_Offset11 = 2
)

var pTRSUplinkConfigTransformPrecoderDisabledResourceElementOffsetConstraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{PTRS_UplinkConfig_TransformPrecoderDisabled_ResourceElementOffset_Offset01, PTRS_UplinkConfig_TransformPrecoderDisabled_ResourceElementOffset_Offset10, PTRS_UplinkConfig_TransformPrecoderDisabled_ResourceElementOffset_Offset11},
}

const (
	PTRS_UplinkConfig_TransformPrecoderDisabled_Ptrs_Power_P00 = 0
	PTRS_UplinkConfig_TransformPrecoderDisabled_Ptrs_Power_P01 = 1
	PTRS_UplinkConfig_TransformPrecoderDisabled_Ptrs_Power_P10 = 2
	PTRS_UplinkConfig_TransformPrecoderDisabled_Ptrs_Power_P11 = 3
)

var pTRSUplinkConfigTransformPrecoderDisabledPtrsPowerConstraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{PTRS_UplinkConfig_TransformPrecoderDisabled_Ptrs_Power_P00, PTRS_UplinkConfig_TransformPrecoderDisabled_Ptrs_Power_P01, PTRS_UplinkConfig_TransformPrecoderDisabled_Ptrs_Power_P10, PTRS_UplinkConfig_TransformPrecoderDisabled_Ptrs_Power_P11},
}

var pTRSUplinkConfigTransformPrecoderEnabledConstraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "sampleDensity"},
		{Name: "timeDensityTransformPrecoding", Optional: true},
	},
}

var pTRSUplinkConfigTransformPrecoderEnabledSampleDensityConstraints = per.FixedSize(5)

const (
	PTRS_UplinkConfig_TransformPrecoderEnabled_TimeDensityTransformPrecoding_D2 = 0
)

var pTRSUplinkConfigTransformPrecoderEnabledTimeDensityTransformPrecodingConstraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{PTRS_UplinkConfig_TransformPrecoderEnabled_TimeDensityTransformPrecoding_D2},
}

type PTRS_UplinkConfig struct {
	TransformPrecoderDisabled *struct {
		FrequencyDensity      []int64
		TimeDensity           []int64
		MaxNrofPorts          int64
		ResourceElementOffset *int64
		Ptrs_Power            int64
	}
	TransformPrecoderEnabled *struct {
		SampleDensity                 []int64
		TimeDensityTransformPrecoding *int64
	}
	MaxNrofPorts_SDM_r18 *int64
}

func (ie *PTRS_UplinkConfig) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(pTRSUplinkConfigConstraints)

	// Extension-group presence (one var per [[...]] group).
	hasExtGroup0 := ie.MaxNrofPorts_SDM_r18 != nil
	hasExtensions := hasExtGroup0

	// 1. SEQUENCE extension bit.
	if err := seq.EncodeExtensionBit(hasExtensions); err != nil {
		return err
	}

	// 2. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.TransformPrecoderDisabled != nil, ie.TransformPrecoderEnabled != nil}); err != nil {
		return err
	}

	// 3. transformPrecoderDisabled: sequence
	{
		if ie.TransformPrecoderDisabled != nil {
			c := ie.TransformPrecoderDisabled
			pTRSUplinkConfigTransformPrecoderDisabledSeq := e.NewSequenceEncoder(pTRSUplinkConfigTransformPrecoderDisabledConstraints)
			if err := pTRSUplinkConfigTransformPrecoderDisabledSeq.EncodePreamble([]bool{c.FrequencyDensity != nil, c.TimeDensity != nil, c.ResourceElementOffset != nil}); err != nil {
				return err
			}
			if c.FrequencyDensity != nil {
				seqOf := e.NewSequenceOfEncoder(pTRSUplinkConfigTransformPrecoderDisabledFrequencyDensityConstraints)
				if err := seqOf.EncodeLength(int64(len(c.FrequencyDensity))); err != nil {
					return err
				}
				for i := range c.FrequencyDensity {
					if err := e.EncodeInteger(c.FrequencyDensity[i], per.Constrained(1, 276)); err != nil {
						return err
					}
				}
			}
			if c.TimeDensity != nil {
				seqOf := e.NewSequenceOfEncoder(pTRSUplinkConfigTransformPrecoderDisabledTimeDensityConstraints)
				if err := seqOf.EncodeLength(int64(len(c.TimeDensity))); err != nil {
					return err
				}
				for i := range c.TimeDensity {
					if err := e.EncodeInteger(c.TimeDensity[i], per.Constrained(0, 29)); err != nil {
						return err
					}
				}
			}
			if err := e.EncodeEnumerated(c.MaxNrofPorts, pTRSUplinkConfigTransformPrecoderDisabledMaxNrofPortsConstraints); err != nil {
				return err
			}
			if c.ResourceElementOffset != nil {
				if err := e.EncodeEnumerated((*c.ResourceElementOffset), pTRSUplinkConfigTransformPrecoderDisabledResourceElementOffsetConstraints); err != nil {
					return err
				}
			}
			if err := e.EncodeEnumerated(c.Ptrs_Power, pTRSUplinkConfigTransformPrecoderDisabledPtrsPowerConstraints); err != nil {
				return err
			}
		}
	}

	// 4. transformPrecoderEnabled: sequence
	{
		if ie.TransformPrecoderEnabled != nil {
			c := ie.TransformPrecoderEnabled
			pTRSUplinkConfigTransformPrecoderEnabledSeq := e.NewSequenceEncoder(pTRSUplinkConfigTransformPrecoderEnabledConstraints)
			if err := pTRSUplinkConfigTransformPrecoderEnabledSeq.EncodePreamble([]bool{c.TimeDensityTransformPrecoding != nil}); err != nil {
				return err
			}
			{
				seqOf := e.NewSequenceOfEncoder(pTRSUplinkConfigTransformPrecoderEnabledSampleDensityConstraints)
				if err := seqOf.EncodeLength(int64(len(c.SampleDensity))); err != nil {
					return err
				}
				for i := range c.SampleDensity {
					if err := e.EncodeInteger(c.SampleDensity[i], per.Constrained(1, 276)); err != nil {
						return err
					}
				}
			}
			if c.TimeDensityTransformPrecoding != nil {
				if err := e.EncodeEnumerated((*c.TimeDensityTransformPrecoding), pTRSUplinkConfigTransformPrecoderEnabledTimeDensityTransformPrecodingConstraints); err != nil {
					return err
				}
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
					{Name: "maxNrofPorts-SDM-r18", Optional: true},
				},
			})
			if err := groupSeq.EncodePreamble([]bool{ie.MaxNrofPorts_SDM_r18 != nil}); err != nil {
				return err
			}

			if ie.MaxNrofPorts_SDM_r18 != nil {
				if err := ex.EncodeEnumerated(*ie.MaxNrofPorts_SDM_r18, pTRSUplinkConfigExtMaxNrofPortsSDMR18Constraints); err != nil {
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

func (ie *PTRS_UplinkConfig) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(pTRSUplinkConfigConstraints)

	// 1. SEQUENCE extension bit.
	if err := seq.DecodeExtensionBit(); err != nil {
		return err
	}

	// 2. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 3. transformPrecoderDisabled: sequence
	{
		if seq.IsComponentPresent(0) {
			ie.TransformPrecoderDisabled = &struct {
				FrequencyDensity      []int64
				TimeDensity           []int64
				MaxNrofPorts          int64
				ResourceElementOffset *int64
				Ptrs_Power            int64
			}{}
			c := ie.TransformPrecoderDisabled
			pTRSUplinkConfigTransformPrecoderDisabledSeq := d.NewSequenceDecoder(pTRSUplinkConfigTransformPrecoderDisabledConstraints)
			if err := pTRSUplinkConfigTransformPrecoderDisabledSeq.DecodePreamble(); err != nil {
				return err
			}
			if pTRSUplinkConfigTransformPrecoderDisabledSeq.IsComponentPresent(0) {
				seqOf := d.NewSequenceOfDecoder(pTRSUplinkConfigTransformPrecoderDisabledFrequencyDensityConstraints)
				n, err := seqOf.DecodeLength()
				if err != nil {
					return err
				}
				c.FrequencyDensity = make([]int64, n)
				for i := int64(0); i < n; i++ {
					v, err := d.DecodeInteger(per.Constrained(1, 276))
					if err != nil {
						return err
					}
					c.FrequencyDensity[i] = v
				}
			}
			if pTRSUplinkConfigTransformPrecoderDisabledSeq.IsComponentPresent(1) {
				seqOf := d.NewSequenceOfDecoder(pTRSUplinkConfigTransformPrecoderDisabledTimeDensityConstraints)
				n, err := seqOf.DecodeLength()
				if err != nil {
					return err
				}
				c.TimeDensity = make([]int64, n)
				for i := int64(0); i < n; i++ {
					v, err := d.DecodeInteger(per.Constrained(0, 29))
					if err != nil {
						return err
					}
					c.TimeDensity[i] = v
				}
			}
			{
				v, err := d.DecodeEnumerated(pTRSUplinkConfigTransformPrecoderDisabledMaxNrofPortsConstraints)
				if err != nil {
					return err
				}
				c.MaxNrofPorts = v
			}
			if pTRSUplinkConfigTransformPrecoderDisabledSeq.IsComponentPresent(3) {
				c.ResourceElementOffset = new(int64)
				v, err := d.DecodeEnumerated(pTRSUplinkConfigTransformPrecoderDisabledResourceElementOffsetConstraints)
				if err != nil {
					return err
				}
				(*c.ResourceElementOffset) = v
			}
			{
				v, err := d.DecodeEnumerated(pTRSUplinkConfigTransformPrecoderDisabledPtrsPowerConstraints)
				if err != nil {
					return err
				}
				c.Ptrs_Power = v
			}
		}
	}

	// 4. transformPrecoderEnabled: sequence
	{
		if seq.IsComponentPresent(1) {
			ie.TransformPrecoderEnabled = &struct {
				SampleDensity                 []int64
				TimeDensityTransformPrecoding *int64
			}{}
			c := ie.TransformPrecoderEnabled
			pTRSUplinkConfigTransformPrecoderEnabledSeq := d.NewSequenceDecoder(pTRSUplinkConfigTransformPrecoderEnabledConstraints)
			if err := pTRSUplinkConfigTransformPrecoderEnabledSeq.DecodePreamble(); err != nil {
				return err
			}
			{
				seqOf := d.NewSequenceOfDecoder(pTRSUplinkConfigTransformPrecoderEnabledSampleDensityConstraints)
				n, err := seqOf.DecodeLength()
				if err != nil {
					return err
				}
				c.SampleDensity = make([]int64, n)
				for i := int64(0); i < n; i++ {
					v, err := d.DecodeInteger(per.Constrained(1, 276))
					if err != nil {
						return err
					}
					c.SampleDensity[i] = v
				}
			}
			if pTRSUplinkConfigTransformPrecoderEnabledSeq.IsComponentPresent(1) {
				c.TimeDensityTransformPrecoding = new(int64)
				v, err := d.DecodeEnumerated(pTRSUplinkConfigTransformPrecoderEnabledTimeDensityTransformPrecodingConstraints)
				if err != nil {
					return err
				}
				(*c.TimeDensityTransformPrecoding) = v
			}
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
				{Name: "maxNrofPorts-SDM-r18", Optional: true},
			},
		})
		if err := groupSeq.DecodePreamble(); err != nil {
			return err
		}

		if groupSeq.IsComponentPresent(0) {
			v, err := dx.DecodeEnumerated(pTRSUplinkConfigExtMaxNrofPortsSDMR18Constraints)
			if err != nil {
				return err
			}
			ie.MaxNrofPorts_SDM_r18 = &v
		}
	}

	return nil
}
