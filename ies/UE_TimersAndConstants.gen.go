// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: UE-TimersAndConstants (line 16220).

var uETimersAndConstantsConstraints = per.SequenceConstraints{
	Extensible: true,
	RootComponents: []per.ComponentInfo{
		{Name: "t300"},
		{Name: "t301"},
		{Name: "t310"},
		{Name: "n310"},
		{Name: "t311"},
		{Name: "n311"},
		{Name: "t319"},
	},
}

const (
	UE_TimersAndConstants_T300_Ms100  = 0
	UE_TimersAndConstants_T300_Ms200  = 1
	UE_TimersAndConstants_T300_Ms300  = 2
	UE_TimersAndConstants_T300_Ms400  = 3
	UE_TimersAndConstants_T300_Ms600  = 4
	UE_TimersAndConstants_T300_Ms1000 = 5
	UE_TimersAndConstants_T300_Ms1500 = 6
	UE_TimersAndConstants_T300_Ms2000 = 7
)

var uETimersAndConstantsT300Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{UE_TimersAndConstants_T300_Ms100, UE_TimersAndConstants_T300_Ms200, UE_TimersAndConstants_T300_Ms300, UE_TimersAndConstants_T300_Ms400, UE_TimersAndConstants_T300_Ms600, UE_TimersAndConstants_T300_Ms1000, UE_TimersAndConstants_T300_Ms1500, UE_TimersAndConstants_T300_Ms2000},
}

const (
	UE_TimersAndConstants_T301_Ms100  = 0
	UE_TimersAndConstants_T301_Ms200  = 1
	UE_TimersAndConstants_T301_Ms300  = 2
	UE_TimersAndConstants_T301_Ms400  = 3
	UE_TimersAndConstants_T301_Ms600  = 4
	UE_TimersAndConstants_T301_Ms1000 = 5
	UE_TimersAndConstants_T301_Ms1500 = 6
	UE_TimersAndConstants_T301_Ms2000 = 7
)

var uETimersAndConstantsT301Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{UE_TimersAndConstants_T301_Ms100, UE_TimersAndConstants_T301_Ms200, UE_TimersAndConstants_T301_Ms300, UE_TimersAndConstants_T301_Ms400, UE_TimersAndConstants_T301_Ms600, UE_TimersAndConstants_T301_Ms1000, UE_TimersAndConstants_T301_Ms1500, UE_TimersAndConstants_T301_Ms2000},
}

const (
	UE_TimersAndConstants_T310_Ms0    = 0
	UE_TimersAndConstants_T310_Ms50   = 1
	UE_TimersAndConstants_T310_Ms100  = 2
	UE_TimersAndConstants_T310_Ms200  = 3
	UE_TimersAndConstants_T310_Ms500  = 4
	UE_TimersAndConstants_T310_Ms1000 = 5
	UE_TimersAndConstants_T310_Ms2000 = 6
)

var uETimersAndConstantsT310Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{UE_TimersAndConstants_T310_Ms0, UE_TimersAndConstants_T310_Ms50, UE_TimersAndConstants_T310_Ms100, UE_TimersAndConstants_T310_Ms200, UE_TimersAndConstants_T310_Ms500, UE_TimersAndConstants_T310_Ms1000, UE_TimersAndConstants_T310_Ms2000},
}

const (
	UE_TimersAndConstants_N310_N1  = 0
	UE_TimersAndConstants_N310_N2  = 1
	UE_TimersAndConstants_N310_N3  = 2
	UE_TimersAndConstants_N310_N4  = 3
	UE_TimersAndConstants_N310_N6  = 4
	UE_TimersAndConstants_N310_N8  = 5
	UE_TimersAndConstants_N310_N10 = 6
	UE_TimersAndConstants_N310_N20 = 7
)

var uETimersAndConstantsN310Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{UE_TimersAndConstants_N310_N1, UE_TimersAndConstants_N310_N2, UE_TimersAndConstants_N310_N3, UE_TimersAndConstants_N310_N4, UE_TimersAndConstants_N310_N6, UE_TimersAndConstants_N310_N8, UE_TimersAndConstants_N310_N10, UE_TimersAndConstants_N310_N20},
}

const (
	UE_TimersAndConstants_T311_Ms1000  = 0
	UE_TimersAndConstants_T311_Ms3000  = 1
	UE_TimersAndConstants_T311_Ms5000  = 2
	UE_TimersAndConstants_T311_Ms10000 = 3
	UE_TimersAndConstants_T311_Ms15000 = 4
	UE_TimersAndConstants_T311_Ms20000 = 5
	UE_TimersAndConstants_T311_Ms30000 = 6
)

var uETimersAndConstantsT311Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{UE_TimersAndConstants_T311_Ms1000, UE_TimersAndConstants_T311_Ms3000, UE_TimersAndConstants_T311_Ms5000, UE_TimersAndConstants_T311_Ms10000, UE_TimersAndConstants_T311_Ms15000, UE_TimersAndConstants_T311_Ms20000, UE_TimersAndConstants_T311_Ms30000},
}

const (
	UE_TimersAndConstants_N311_N1  = 0
	UE_TimersAndConstants_N311_N2  = 1
	UE_TimersAndConstants_N311_N3  = 2
	UE_TimersAndConstants_N311_N4  = 3
	UE_TimersAndConstants_N311_N5  = 4
	UE_TimersAndConstants_N311_N6  = 5
	UE_TimersAndConstants_N311_N8  = 6
	UE_TimersAndConstants_N311_N10 = 7
)

var uETimersAndConstantsN311Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{UE_TimersAndConstants_N311_N1, UE_TimersAndConstants_N311_N2, UE_TimersAndConstants_N311_N3, UE_TimersAndConstants_N311_N4, UE_TimersAndConstants_N311_N5, UE_TimersAndConstants_N311_N6, UE_TimersAndConstants_N311_N8, UE_TimersAndConstants_N311_N10},
}

const (
	UE_TimersAndConstants_T319_Ms100  = 0
	UE_TimersAndConstants_T319_Ms200  = 1
	UE_TimersAndConstants_T319_Ms300  = 2
	UE_TimersAndConstants_T319_Ms400  = 3
	UE_TimersAndConstants_T319_Ms600  = 4
	UE_TimersAndConstants_T319_Ms1000 = 5
	UE_TimersAndConstants_T319_Ms1500 = 6
	UE_TimersAndConstants_T319_Ms2000 = 7
)

var uETimersAndConstantsT319Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{UE_TimersAndConstants_T319_Ms100, UE_TimersAndConstants_T319_Ms200, UE_TimersAndConstants_T319_Ms300, UE_TimersAndConstants_T319_Ms400, UE_TimersAndConstants_T319_Ms600, UE_TimersAndConstants_T319_Ms1000, UE_TimersAndConstants_T319_Ms1500, UE_TimersAndConstants_T319_Ms2000},
}

type UE_TimersAndConstants struct {
	T300 int64
	T301 int64
	T310 int64
	N310 int64
	T311 int64
	N311 int64
	T319 int64
}

func (ie *UE_TimersAndConstants) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(uETimersAndConstantsConstraints)

	// 1. SEQUENCE extension bit.
	if err := seq.EncodeExtensionBit(false); err != nil {
		return err
	}

	// 2. t300: enumerated
	{
		if err := e.EncodeEnumerated(ie.T300, uETimersAndConstantsT300Constraints); err != nil {
			return err
		}
	}

	// 3. t301: enumerated
	{
		if err := e.EncodeEnumerated(ie.T301, uETimersAndConstantsT301Constraints); err != nil {
			return err
		}
	}

	// 4. t310: enumerated
	{
		if err := e.EncodeEnumerated(ie.T310, uETimersAndConstantsT310Constraints); err != nil {
			return err
		}
	}

	// 5. n310: enumerated
	{
		if err := e.EncodeEnumerated(ie.N310, uETimersAndConstantsN310Constraints); err != nil {
			return err
		}
	}

	// 6. t311: enumerated
	{
		if err := e.EncodeEnumerated(ie.T311, uETimersAndConstantsT311Constraints); err != nil {
			return err
		}
	}

	// 7. n311: enumerated
	{
		if err := e.EncodeEnumerated(ie.N311, uETimersAndConstantsN311Constraints); err != nil {
			return err
		}
	}

	// 8. t319: enumerated
	{
		if err := e.EncodeEnumerated(ie.T319, uETimersAndConstantsT319Constraints); err != nil {
			return err
		}
	}

	return nil
}

func (ie *UE_TimersAndConstants) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(uETimersAndConstantsConstraints)

	// 1. SEQUENCE extension bit.
	if err := seq.DecodeExtensionBit(); err != nil {
		return err
	}

	// 2. t300: enumerated
	{
		v0, err := d.DecodeEnumerated(uETimersAndConstantsT300Constraints)
		if err != nil {
			return err
		}
		ie.T300 = v0
	}

	// 3. t301: enumerated
	{
		v1, err := d.DecodeEnumerated(uETimersAndConstantsT301Constraints)
		if err != nil {
			return err
		}
		ie.T301 = v1
	}

	// 4. t310: enumerated
	{
		v2, err := d.DecodeEnumerated(uETimersAndConstantsT310Constraints)
		if err != nil {
			return err
		}
		ie.T310 = v2
	}

	// 5. n310: enumerated
	{
		v3, err := d.DecodeEnumerated(uETimersAndConstantsN310Constraints)
		if err != nil {
			return err
		}
		ie.N310 = v3
	}

	// 6. t311: enumerated
	{
		v4, err := d.DecodeEnumerated(uETimersAndConstantsT311Constraints)
		if err != nil {
			return err
		}
		ie.T311 = v4
	}

	// 7. n311: enumerated
	{
		v5, err := d.DecodeEnumerated(uETimersAndConstantsN311Constraints)
		if err != nil {
			return err
		}
		ie.N311 = v5
	}

	// 8. t319: enumerated
	{
		v6, err := d.DecodeEnumerated(uETimersAndConstantsT319Constraints)
		if err != nil {
			return err
		}
		ie.T319 = v6
	}

	return nil
}
