// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: PTRS-DensityRecommendationDL (line 22498).

var pTRSDensityRecommendationDLConstraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "frequencyDensity1"},
		{Name: "frequencyDensity2"},
		{Name: "timeDensity1"},
		{Name: "timeDensity2"},
		{Name: "timeDensity3"},
	},
}

var pTRSDensityRecommendationDLFrequencyDensity1Constraints = per.Constrained(1, 276)

var pTRSDensityRecommendationDLFrequencyDensity2Constraints = per.Constrained(1, 276)

var pTRSDensityRecommendationDLTimeDensity1Constraints = per.Constrained(0, 29)

var pTRSDensityRecommendationDLTimeDensity2Constraints = per.Constrained(0, 29)

var pTRSDensityRecommendationDLTimeDensity3Constraints = per.Constrained(0, 29)

type PTRS_DensityRecommendationDL struct {
	FrequencyDensity1 int64
	FrequencyDensity2 int64
	TimeDensity1      int64
	TimeDensity2      int64
	TimeDensity3      int64
}

func (ie *PTRS_DensityRecommendationDL) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(pTRSDensityRecommendationDLConstraints)
	_ = seq

	// 1. frequencyDensity1: integer
	{
		if err := e.EncodeInteger(ie.FrequencyDensity1, pTRSDensityRecommendationDLFrequencyDensity1Constraints); err != nil {
			return err
		}
	}

	// 2. frequencyDensity2: integer
	{
		if err := e.EncodeInteger(ie.FrequencyDensity2, pTRSDensityRecommendationDLFrequencyDensity2Constraints); err != nil {
			return err
		}
	}

	// 3. timeDensity1: integer
	{
		if err := e.EncodeInteger(ie.TimeDensity1, pTRSDensityRecommendationDLTimeDensity1Constraints); err != nil {
			return err
		}
	}

	// 4. timeDensity2: integer
	{
		if err := e.EncodeInteger(ie.TimeDensity2, pTRSDensityRecommendationDLTimeDensity2Constraints); err != nil {
			return err
		}
	}

	// 5. timeDensity3: integer
	{
		if err := e.EncodeInteger(ie.TimeDensity3, pTRSDensityRecommendationDLTimeDensity3Constraints); err != nil {
			return err
		}
	}

	return nil
}

func (ie *PTRS_DensityRecommendationDL) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(pTRSDensityRecommendationDLConstraints)
	_ = seq

	// 1. frequencyDensity1: integer
	{
		v0, err := d.DecodeInteger(pTRSDensityRecommendationDLFrequencyDensity1Constraints)
		if err != nil {
			return err
		}
		ie.FrequencyDensity1 = v0
	}

	// 2. frequencyDensity2: integer
	{
		v1, err := d.DecodeInteger(pTRSDensityRecommendationDLFrequencyDensity2Constraints)
		if err != nil {
			return err
		}
		ie.FrequencyDensity2 = v1
	}

	// 3. timeDensity1: integer
	{
		v2, err := d.DecodeInteger(pTRSDensityRecommendationDLTimeDensity1Constraints)
		if err != nil {
			return err
		}
		ie.TimeDensity1 = v2
	}

	// 4. timeDensity2: integer
	{
		v3, err := d.DecodeInteger(pTRSDensityRecommendationDLTimeDensity2Constraints)
		if err != nil {
			return err
		}
		ie.TimeDensity2 = v3
	}

	// 5. timeDensity3: integer
	{
		v4, err := d.DecodeInteger(pTRSDensityRecommendationDLTimeDensity3Constraints)
		if err != nil {
			return err
		}
		ie.TimeDensity3 = v4
	}

	return nil
}
