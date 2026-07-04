// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: Orbital-r17 (line 8276).

var orbitalR17Constraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "semiMajorAxis-r17"},
		{Name: "eccentricity-r17"},
		{Name: "periapsis-r17"},
		{Name: "longitude-r17"},
		{Name: "inclination-r17"},
		{Name: "meanAnomaly-r17"},
	},
}

var orbitalR17SemiMajorAxisR17Constraints = per.Constrained(0, 8589934591)

var orbitalR17EccentricityR17Constraints = per.Constrained(0, 1048575)

var orbitalR17PeriapsisR17Constraints = per.Constrained(0, 268435455)

var orbitalR17LongitudeR17Constraints = per.Constrained(0, 268435455)

var orbitalR17InclinationR17Constraints = per.Constrained(-67108864, 67108863)

var orbitalR17MeanAnomalyR17Constraints = per.Constrained(0, 268435455)

type Orbital_r17 struct {
	SemiMajorAxis_r17 int64
	Eccentricity_r17  int64
	Periapsis_r17     int64
	Longitude_r17     int64
	Inclination_r17   int64
	MeanAnomaly_r17   int64
}

func (ie *Orbital_r17) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(orbitalR17Constraints)
	_ = seq

	// 1. semiMajorAxis-r17: integer
	{
		if err := e.EncodeInteger(ie.SemiMajorAxis_r17, orbitalR17SemiMajorAxisR17Constraints); err != nil {
			return err
		}
	}

	// 2. eccentricity-r17: integer
	{
		if err := e.EncodeInteger(ie.Eccentricity_r17, orbitalR17EccentricityR17Constraints); err != nil {
			return err
		}
	}

	// 3. periapsis-r17: integer
	{
		if err := e.EncodeInteger(ie.Periapsis_r17, orbitalR17PeriapsisR17Constraints); err != nil {
			return err
		}
	}

	// 4. longitude-r17: integer
	{
		if err := e.EncodeInteger(ie.Longitude_r17, orbitalR17LongitudeR17Constraints); err != nil {
			return err
		}
	}

	// 5. inclination-r17: integer
	{
		if err := e.EncodeInteger(ie.Inclination_r17, orbitalR17InclinationR17Constraints); err != nil {
			return err
		}
	}

	// 6. meanAnomaly-r17: integer
	{
		if err := e.EncodeInteger(ie.MeanAnomaly_r17, orbitalR17MeanAnomalyR17Constraints); err != nil {
			return err
		}
	}

	return nil
}

func (ie *Orbital_r17) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(orbitalR17Constraints)
	_ = seq

	// 1. semiMajorAxis-r17: integer
	{
		v0, err := d.DecodeInteger(orbitalR17SemiMajorAxisR17Constraints)
		if err != nil {
			return err
		}
		ie.SemiMajorAxis_r17 = v0
	}

	// 2. eccentricity-r17: integer
	{
		v1, err := d.DecodeInteger(orbitalR17EccentricityR17Constraints)
		if err != nil {
			return err
		}
		ie.Eccentricity_r17 = v1
	}

	// 3. periapsis-r17: integer
	{
		v2, err := d.DecodeInteger(orbitalR17PeriapsisR17Constraints)
		if err != nil {
			return err
		}
		ie.Periapsis_r17 = v2
	}

	// 4. longitude-r17: integer
	{
		v3, err := d.DecodeInteger(orbitalR17LongitudeR17Constraints)
		if err != nil {
			return err
		}
		ie.Longitude_r17 = v3
	}

	// 5. inclination-r17: integer
	{
		v4, err := d.DecodeInteger(orbitalR17InclinationR17Constraints)
		if err != nil {
			return err
		}
		ie.Inclination_r17 = v4
	}

	// 6. meanAnomaly-r17: integer
	{
		v5, err := d.DecodeInteger(orbitalR17MeanAnomalyR17Constraints)
		if err != nil {
			return err
		}
		ie.MeanAnomaly_r17 = v5
	}

	return nil
}
