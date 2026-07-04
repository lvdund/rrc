// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: EUTRA-MultiBandInfoAerial-r18 (line 26170).

var eUTRAMultiBandInfoAerialR18Constraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "eutra-FreqBandIndicator-r18"},
		{Name: "eutra-NS-PmaxListAerial-r18", Optional: true},
	},
}

type EUTRA_MultiBandInfoAerial_r18 struct {
	Eutra_FreqBandIndicator_r18 FreqBandIndicatorEUTRA
	Eutra_NS_PmaxListAerial_r18 *EUTRA_NS_PmaxList
}

func (ie *EUTRA_MultiBandInfoAerial_r18) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(eUTRAMultiBandInfoAerialR18Constraints)

	// 1. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.Eutra_NS_PmaxListAerial_r18 != nil}); err != nil {
		return err
	}

	// 2. eutra-FreqBandIndicator-r18: ref
	{
		if err := ie.Eutra_FreqBandIndicator_r18.Encode(e); err != nil {
			return err
		}
	}

	// 3. eutra-NS-PmaxListAerial-r18: ref
	{
		if ie.Eutra_NS_PmaxListAerial_r18 != nil {
			if err := ie.Eutra_NS_PmaxListAerial_r18.Encode(e); err != nil {
				return err
			}
		}
	}

	return nil
}

func (ie *EUTRA_MultiBandInfoAerial_r18) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(eUTRAMultiBandInfoAerialR18Constraints)

	// 1. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 2. eutra-FreqBandIndicator-r18: ref
	{
		if err := ie.Eutra_FreqBandIndicator_r18.Decode(d); err != nil {
			return err
		}
	}

	// 3. eutra-NS-PmaxListAerial-r18: ref
	{
		if seq.IsComponentPresent(1) {
			ie.Eutra_NS_PmaxListAerial_r18 = new(EUTRA_NS_PmaxList)
			if err := ie.Eutra_NS_PmaxListAerial_r18.Decode(d); err != nil {
				return err
			}
		}
	}

	return nil
}
