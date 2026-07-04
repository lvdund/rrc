// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/rrc/common"
)

// Source: FeatureSetsPerBand (line 19426).
// FeatureSetsPerBand ::=          SEQUENCE (SIZE (1..maxFeatureSetsPerBand)) OF FeatureSet

var featureSetsPerBandSizeConstraints = per.SizeRange(1, common.MaxFeatureSetsPerBand)

type FeatureSetsPerBand []FeatureSet

func (ie *FeatureSetsPerBand) Encode(e *per.Encoder) error {
	seqOf := e.NewSequenceOfEncoder(featureSetsPerBandSizeConstraints)
	if err := seqOf.EncodeLength(int64(len(*ie))); err != nil {
		return err
	}
	for i := range *ie {
		if err := (*ie)[i].Encode(e); err != nil {
			return err
		}
	}
	return nil
}

func (ie *FeatureSetsPerBand) Decode(d *per.Decoder) error {
	seqOf := d.NewSequenceOfDecoder(featureSetsPerBandSizeConstraints)
	n, err := seqOf.DecodeLength()
	if err != nil {
		return err
	}
	*ie = make(FeatureSetsPerBand, n)
	for i := int64(0); i < n; i++ {
		if err := (*ie)[i].Decode(d); err != nil {
			return err
		}
	}
	return nil
}
