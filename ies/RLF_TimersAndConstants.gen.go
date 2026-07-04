// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: RLF-TimersAndConstants (line 14188).

var rLFTimersAndConstantsConstraints = per.SequenceConstraints{
	Extensible: true,
	RootComponents: []per.ComponentInfo{
		{Name: "t310"},
		{Name: "n310"},
		{Name: "n311"},
	},
	ExtComponents: []per.ComponentInfo{
		{Name: "extension-addition-0", Optional: true},
	},
}

const (
	RLF_TimersAndConstants_T310_Ms0    = 0
	RLF_TimersAndConstants_T310_Ms50   = 1
	RLF_TimersAndConstants_T310_Ms100  = 2
	RLF_TimersAndConstants_T310_Ms200  = 3
	RLF_TimersAndConstants_T310_Ms500  = 4
	RLF_TimersAndConstants_T310_Ms1000 = 5
	RLF_TimersAndConstants_T310_Ms2000 = 6
	RLF_TimersAndConstants_T310_Ms4000 = 7
	RLF_TimersAndConstants_T310_Ms6000 = 8
)

var rLFTimersAndConstantsT310Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{RLF_TimersAndConstants_T310_Ms0, RLF_TimersAndConstants_T310_Ms50, RLF_TimersAndConstants_T310_Ms100, RLF_TimersAndConstants_T310_Ms200, RLF_TimersAndConstants_T310_Ms500, RLF_TimersAndConstants_T310_Ms1000, RLF_TimersAndConstants_T310_Ms2000, RLF_TimersAndConstants_T310_Ms4000, RLF_TimersAndConstants_T310_Ms6000},
}

const (
	RLF_TimersAndConstants_N310_N1  = 0
	RLF_TimersAndConstants_N310_N2  = 1
	RLF_TimersAndConstants_N310_N3  = 2
	RLF_TimersAndConstants_N310_N4  = 3
	RLF_TimersAndConstants_N310_N6  = 4
	RLF_TimersAndConstants_N310_N8  = 5
	RLF_TimersAndConstants_N310_N10 = 6
	RLF_TimersAndConstants_N310_N20 = 7
)

var rLFTimersAndConstantsN310Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{RLF_TimersAndConstants_N310_N1, RLF_TimersAndConstants_N310_N2, RLF_TimersAndConstants_N310_N3, RLF_TimersAndConstants_N310_N4, RLF_TimersAndConstants_N310_N6, RLF_TimersAndConstants_N310_N8, RLF_TimersAndConstants_N310_N10, RLF_TimersAndConstants_N310_N20},
}

const (
	RLF_TimersAndConstants_N311_N1  = 0
	RLF_TimersAndConstants_N311_N2  = 1
	RLF_TimersAndConstants_N311_N3  = 2
	RLF_TimersAndConstants_N311_N4  = 3
	RLF_TimersAndConstants_N311_N5  = 4
	RLF_TimersAndConstants_N311_N6  = 5
	RLF_TimersAndConstants_N311_N8  = 6
	RLF_TimersAndConstants_N311_N10 = 7
)

var rLFTimersAndConstantsN311Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{RLF_TimersAndConstants_N311_N1, RLF_TimersAndConstants_N311_N2, RLF_TimersAndConstants_N311_N3, RLF_TimersAndConstants_N311_N4, RLF_TimersAndConstants_N311_N5, RLF_TimersAndConstants_N311_N6, RLF_TimersAndConstants_N311_N8, RLF_TimersAndConstants_N311_N10},
}

const (
	RLF_TimersAndConstants_Ext_T311_Ms1000  = 0
	RLF_TimersAndConstants_Ext_T311_Ms3000  = 1
	RLF_TimersAndConstants_Ext_T311_Ms5000  = 2
	RLF_TimersAndConstants_Ext_T311_Ms10000 = 3
	RLF_TimersAndConstants_Ext_T311_Ms15000 = 4
	RLF_TimersAndConstants_Ext_T311_Ms20000 = 5
	RLF_TimersAndConstants_Ext_T311_Ms30000 = 6
)

var rLFTimersAndConstantsExtT311Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{RLF_TimersAndConstants_Ext_T311_Ms1000, RLF_TimersAndConstants_Ext_T311_Ms3000, RLF_TimersAndConstants_Ext_T311_Ms5000, RLF_TimersAndConstants_Ext_T311_Ms10000, RLF_TimersAndConstants_Ext_T311_Ms15000, RLF_TimersAndConstants_Ext_T311_Ms20000, RLF_TimersAndConstants_Ext_T311_Ms30000},
}

type RLF_TimersAndConstants struct {
	T310 int64
	N310 int64
	N311 int64
	T311 int64
}

func (ie *RLF_TimersAndConstants) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(rLFTimersAndConstantsConstraints)

	// Extension-group presence (one var per [[...]] group).
	hasExtGroup0 := true
	hasExtensions := hasExtGroup0

	// 1. SEQUENCE extension bit.
	if err := seq.EncodeExtensionBit(hasExtensions); err != nil {
		return err
	}

	// 2. t310: enumerated
	{
		if err := e.EncodeEnumerated(ie.T310, rLFTimersAndConstantsT310Constraints); err != nil {
			return err
		}
	}

	// 3. n310: enumerated
	{
		if err := e.EncodeEnumerated(ie.N310, rLFTimersAndConstantsN310Constraints); err != nil {
			return err
		}
	}

	// 4. n311: enumerated
	{
		if err := e.EncodeEnumerated(ie.N311, rLFTimersAndConstantsN311Constraints); err != nil {
			return err
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
					{Name: "t311"},
				},
			})
			if err := groupSeq.EncodePreamble([]bool{}); err != nil {
				return err
			}

			// T311: mandatory in group
			if err := ex.EncodeEnumerated(ie.T311, rLFTimersAndConstantsExtT311Constraints); err != nil {
				return err
			}
			extValues = append(extValues, ex.Bytes())
		}

		if err := seq.EncodeExtensionAdditions(extPresence, extValues); err != nil {
			return err
		}
	}

	return nil
}

func (ie *RLF_TimersAndConstants) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(rLFTimersAndConstantsConstraints)

	// 1. SEQUENCE extension bit.
	if err := seq.DecodeExtensionBit(); err != nil {
		return err
	}

	// 2. t310: enumerated
	{
		v0, err := d.DecodeEnumerated(rLFTimersAndConstantsT310Constraints)
		if err != nil {
			return err
		}
		ie.T310 = v0
	}

	// 3. n310: enumerated
	{
		v1, err := d.DecodeEnumerated(rLFTimersAndConstantsN310Constraints)
		if err != nil {
			return err
		}
		ie.N310 = v1
	}

	// 4. n311: enumerated
	{
		v2, err := d.DecodeEnumerated(rLFTimersAndConstantsN311Constraints)
		if err != nil {
			return err
		}
		ie.N311 = v2
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
				{Name: "t311"},
			},
		})
		if err := groupSeq.DecodePreamble(); err != nil {
			return err
		}

		// T311: mandatory in group
		v, err := dx.DecodeEnumerated(rLFTimersAndConstantsExtT311Constraints)
		if err != nil {
			return err
		}
		ie.T311 = v
	}

	return nil
}
