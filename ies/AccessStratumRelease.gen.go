// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: AccessStratumRelease (line 16471).

const (
	AccessStratumRelease_Rel15  = 0
	AccessStratumRelease_Rel16  = 1
	AccessStratumRelease_Rel17  = 2
	AccessStratumRelease_Rel18  = 3
	AccessStratumRelease_Rel19  = 4
	AccessStratumRelease_Spare3 = 5
	AccessStratumRelease_Spare2 = 6
	AccessStratumRelease_Spare1 = 7
)

var accessStratumReleaseConstraints = per.EnumeratedConstraints{
	Extensible: true,
	RootValues: []int64{AccessStratumRelease_Rel15, AccessStratumRelease_Rel16, AccessStratumRelease_Rel17, AccessStratumRelease_Rel18, AccessStratumRelease_Rel19, AccessStratumRelease_Spare3, AccessStratumRelease_Spare2, AccessStratumRelease_Spare1},
}

type AccessStratumRelease struct {
	Value int64
}

func (v *AccessStratumRelease) Encode(e *per.Encoder) error {
	return e.EncodeEnumerated(v.Value, accessStratumReleaseConstraints)
}

func (v *AccessStratumRelease) Decode(d *per.Decoder) error {
	x, err := d.DecodeEnumerated(accessStratumReleaseConstraints)
	if err != nil {
		return err
	}
	v.Value = x
	return nil
}
