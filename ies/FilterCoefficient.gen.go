// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: FilterCoefficient (line 8346).
// FilterCoefficient ::=       ENUMERATED { fc0, fc1, fc2, fc3, fc4, fc5, fc6, fc7, fc8, fc9, fc11, fc13, fc15, fc17, fc19, spare1, ...}

const (
	FilterCoefficient_Fc0    = 0
	FilterCoefficient_Fc1    = 1
	FilterCoefficient_Fc2    = 2
	FilterCoefficient_Fc3    = 3
	FilterCoefficient_Fc4    = 4
	FilterCoefficient_Fc5    = 5
	FilterCoefficient_Fc6    = 6
	FilterCoefficient_Fc7    = 7
	FilterCoefficient_Fc8    = 8
	FilterCoefficient_Fc9    = 9
	FilterCoefficient_Fc11   = 10
	FilterCoefficient_Fc13   = 11
	FilterCoefficient_Fc15   = 12
	FilterCoefficient_Fc17   = 13
	FilterCoefficient_Fc19   = 14
	FilterCoefficient_Spare1 = 15
)

var filterCoefficientConstraints = per.EnumeratedConstraints{
	Extensible: true,
	RootValues: []int64{FilterCoefficient_Fc0, FilterCoefficient_Fc1, FilterCoefficient_Fc2, FilterCoefficient_Fc3, FilterCoefficient_Fc4, FilterCoefficient_Fc5, FilterCoefficient_Fc6, FilterCoefficient_Fc7, FilterCoefficient_Fc8, FilterCoefficient_Fc9, FilterCoefficient_Fc11, FilterCoefficient_Fc13, FilterCoefficient_Fc15, FilterCoefficient_Fc17, FilterCoefficient_Fc19, FilterCoefficient_Spare1},
}

type FilterCoefficient struct {
	Value int64
}

func (v *FilterCoefficient) Encode(e *per.Encoder) error {
	return e.EncodeEnumerated(v.Value, filterCoefficientConstraints)
}

func (v *FilterCoefficient) Decode(d *per.Decoder) error {
	x, err := d.DecodeEnumerated(filterCoefficientConstraints)
	if err != nil {
		return err
	}
	v.Value = x
	return nil
}
