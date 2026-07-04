// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: PTRS-DensityRecommendationUL (line 22506).

var pTRSDensityRecommendationULConstraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "frequencyDensity1"},
		{Name: "frequencyDensity2"},
		{Name: "timeDensity1"},
		{Name: "timeDensity2"},
		{Name: "timeDensity3"},
		{Name: "sampleDensity1"},
		{Name: "sampleDensity2"},
		{Name: "sampleDensity3"},
		{Name: "sampleDensity4"},
		{Name: "sampleDensity5"},
	},
}

var pTRSDensityRecommendationULFrequencyDensity1Constraints = per.Constrained(1, 276)

var pTRSDensityRecommendationULFrequencyDensity2Constraints = per.Constrained(1, 276)

var pTRSDensityRecommendationULTimeDensity1Constraints = per.Constrained(0, 29)

var pTRSDensityRecommendationULTimeDensity2Constraints = per.Constrained(0, 29)

var pTRSDensityRecommendationULTimeDensity3Constraints = per.Constrained(0, 29)

var pTRSDensityRecommendationULSampleDensity1Constraints = per.Constrained(1, 276)

var pTRSDensityRecommendationULSampleDensity2Constraints = per.Constrained(1, 276)

var pTRSDensityRecommendationULSampleDensity3Constraints = per.Constrained(1, 276)

var pTRSDensityRecommendationULSampleDensity4Constraints = per.Constrained(1, 276)

var pTRSDensityRecommendationULSampleDensity5Constraints = per.Constrained(1, 276)

type PTRS_DensityRecommendationUL struct {
	FrequencyDensity1 int64
	FrequencyDensity2 int64
	TimeDensity1      int64
	TimeDensity2      int64
	TimeDensity3      int64
	SampleDensity1    int64
	SampleDensity2    int64
	SampleDensity3    int64
	SampleDensity4    int64
	SampleDensity5    int64
}

func (ie *PTRS_DensityRecommendationUL) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(pTRSDensityRecommendationULConstraints)
	_ = seq

	// 1. frequencyDensity1: integer
	{
		if err := e.EncodeInteger(ie.FrequencyDensity1, pTRSDensityRecommendationULFrequencyDensity1Constraints); err != nil {
			return err
		}
	}

	// 2. frequencyDensity2: integer
	{
		if err := e.EncodeInteger(ie.FrequencyDensity2, pTRSDensityRecommendationULFrequencyDensity2Constraints); err != nil {
			return err
		}
	}

	// 3. timeDensity1: integer
	{
		if err := e.EncodeInteger(ie.TimeDensity1, pTRSDensityRecommendationULTimeDensity1Constraints); err != nil {
			return err
		}
	}

	// 4. timeDensity2: integer
	{
		if err := e.EncodeInteger(ie.TimeDensity2, pTRSDensityRecommendationULTimeDensity2Constraints); err != nil {
			return err
		}
	}

	// 5. timeDensity3: integer
	{
		if err := e.EncodeInteger(ie.TimeDensity3, pTRSDensityRecommendationULTimeDensity3Constraints); err != nil {
			return err
		}
	}

	// 6. sampleDensity1: integer
	{
		if err := e.EncodeInteger(ie.SampleDensity1, pTRSDensityRecommendationULSampleDensity1Constraints); err != nil {
			return err
		}
	}

	// 7. sampleDensity2: integer
	{
		if err := e.EncodeInteger(ie.SampleDensity2, pTRSDensityRecommendationULSampleDensity2Constraints); err != nil {
			return err
		}
	}

	// 8. sampleDensity3: integer
	{
		if err := e.EncodeInteger(ie.SampleDensity3, pTRSDensityRecommendationULSampleDensity3Constraints); err != nil {
			return err
		}
	}

	// 9. sampleDensity4: integer
	{
		if err := e.EncodeInteger(ie.SampleDensity4, pTRSDensityRecommendationULSampleDensity4Constraints); err != nil {
			return err
		}
	}

	// 10. sampleDensity5: integer
	{
		if err := e.EncodeInteger(ie.SampleDensity5, pTRSDensityRecommendationULSampleDensity5Constraints); err != nil {
			return err
		}
	}

	return nil
}

func (ie *PTRS_DensityRecommendationUL) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(pTRSDensityRecommendationULConstraints)
	_ = seq

	// 1. frequencyDensity1: integer
	{
		v0, err := d.DecodeInteger(pTRSDensityRecommendationULFrequencyDensity1Constraints)
		if err != nil {
			return err
		}
		ie.FrequencyDensity1 = v0
	}

	// 2. frequencyDensity2: integer
	{
		v1, err := d.DecodeInteger(pTRSDensityRecommendationULFrequencyDensity2Constraints)
		if err != nil {
			return err
		}
		ie.FrequencyDensity2 = v1
	}

	// 3. timeDensity1: integer
	{
		v2, err := d.DecodeInteger(pTRSDensityRecommendationULTimeDensity1Constraints)
		if err != nil {
			return err
		}
		ie.TimeDensity1 = v2
	}

	// 4. timeDensity2: integer
	{
		v3, err := d.DecodeInteger(pTRSDensityRecommendationULTimeDensity2Constraints)
		if err != nil {
			return err
		}
		ie.TimeDensity2 = v3
	}

	// 5. timeDensity3: integer
	{
		v4, err := d.DecodeInteger(pTRSDensityRecommendationULTimeDensity3Constraints)
		if err != nil {
			return err
		}
		ie.TimeDensity3 = v4
	}

	// 6. sampleDensity1: integer
	{
		v5, err := d.DecodeInteger(pTRSDensityRecommendationULSampleDensity1Constraints)
		if err != nil {
			return err
		}
		ie.SampleDensity1 = v5
	}

	// 7. sampleDensity2: integer
	{
		v6, err := d.DecodeInteger(pTRSDensityRecommendationULSampleDensity2Constraints)
		if err != nil {
			return err
		}
		ie.SampleDensity2 = v6
	}

	// 8. sampleDensity3: integer
	{
		v7, err := d.DecodeInteger(pTRSDensityRecommendationULSampleDensity3Constraints)
		if err != nil {
			return err
		}
		ie.SampleDensity3 = v7
	}

	// 9. sampleDensity4: integer
	{
		v8, err := d.DecodeInteger(pTRSDensityRecommendationULSampleDensity4Constraints)
		if err != nil {
			return err
		}
		ie.SampleDensity4 = v8
	}

	// 10. sampleDensity5: integer
	{
		v9, err := d.DecodeInteger(pTRSDensityRecommendationULSampleDensity5Constraints)
		if err != nil {
			return err
		}
		ie.SampleDensity5 = v9
	}

	return nil
}
