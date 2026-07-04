// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: CarrierAggregationVariant (line 18474).

var carrierAggregationVariantConstraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "fr1fdd-FR1TDD-CA-SpCellOnFR1FDD", Optional: true},
		{Name: "fr1fdd-FR1TDD-CA-SpCellOnFR1TDD", Optional: true},
		{Name: "fr1fdd-FR2TDD-CA-SpCellOnFR1FDD", Optional: true},
		{Name: "fr1fdd-FR2TDD-CA-SpCellOnFR2TDD", Optional: true},
		{Name: "fr1tdd-FR2TDD-CA-SpCellOnFR1TDD", Optional: true},
		{Name: "fr1tdd-FR2TDD-CA-SpCellOnFR2TDD", Optional: true},
		{Name: "fr1fdd-FR1TDD-FR2TDD-CA-SpCellOnFR1FDD", Optional: true},
		{Name: "fr1fdd-FR1TDD-FR2TDD-CA-SpCellOnFR1TDD", Optional: true},
		{Name: "fr1fdd-FR1TDD-FR2TDD-CA-SpCellOnFR2TDD", Optional: true},
	},
}

const (
	CarrierAggregationVariant_Fr1fdd_FR1TDD_CA_SpCellOnFR1FDD_Supported = 0
)

var carrierAggregationVariantFr1fddFR1TDDCASpCellOnFR1FDDConstraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{CarrierAggregationVariant_Fr1fdd_FR1TDD_CA_SpCellOnFR1FDD_Supported},
}

const (
	CarrierAggregationVariant_Fr1fdd_FR1TDD_CA_SpCellOnFR1TDD_Supported = 0
)

var carrierAggregationVariantFr1fddFR1TDDCASpCellOnFR1TDDConstraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{CarrierAggregationVariant_Fr1fdd_FR1TDD_CA_SpCellOnFR1TDD_Supported},
}

const (
	CarrierAggregationVariant_Fr1fdd_FR2TDD_CA_SpCellOnFR1FDD_Supported = 0
)

var carrierAggregationVariantFr1fddFR2TDDCASpCellOnFR1FDDConstraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{CarrierAggregationVariant_Fr1fdd_FR2TDD_CA_SpCellOnFR1FDD_Supported},
}

const (
	CarrierAggregationVariant_Fr1fdd_FR2TDD_CA_SpCellOnFR2TDD_Supported = 0
)

var carrierAggregationVariantFr1fddFR2TDDCASpCellOnFR2TDDConstraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{CarrierAggregationVariant_Fr1fdd_FR2TDD_CA_SpCellOnFR2TDD_Supported},
}

const (
	CarrierAggregationVariant_Fr1tdd_FR2TDD_CA_SpCellOnFR1TDD_Supported = 0
)

var carrierAggregationVariantFr1tddFR2TDDCASpCellOnFR1TDDConstraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{CarrierAggregationVariant_Fr1tdd_FR2TDD_CA_SpCellOnFR1TDD_Supported},
}

const (
	CarrierAggregationVariant_Fr1tdd_FR2TDD_CA_SpCellOnFR2TDD_Supported = 0
)

var carrierAggregationVariantFr1tddFR2TDDCASpCellOnFR2TDDConstraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{CarrierAggregationVariant_Fr1tdd_FR2TDD_CA_SpCellOnFR2TDD_Supported},
}

const (
	CarrierAggregationVariant_Fr1fdd_FR1TDD_FR2TDD_CA_SpCellOnFR1FDD_Supported = 0
)

var carrierAggregationVariantFr1fddFR1TDDFR2TDDCASpCellOnFR1FDDConstraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{CarrierAggregationVariant_Fr1fdd_FR1TDD_FR2TDD_CA_SpCellOnFR1FDD_Supported},
}

const (
	CarrierAggregationVariant_Fr1fdd_FR1TDD_FR2TDD_CA_SpCellOnFR1TDD_Supported = 0
)

var carrierAggregationVariantFr1fddFR1TDDFR2TDDCASpCellOnFR1TDDConstraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{CarrierAggregationVariant_Fr1fdd_FR1TDD_FR2TDD_CA_SpCellOnFR1TDD_Supported},
}

const (
	CarrierAggregationVariant_Fr1fdd_FR1TDD_FR2TDD_CA_SpCellOnFR2TDD_Supported = 0
)

var carrierAggregationVariantFr1fddFR1TDDFR2TDDCASpCellOnFR2TDDConstraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{CarrierAggregationVariant_Fr1fdd_FR1TDD_FR2TDD_CA_SpCellOnFR2TDD_Supported},
}

type CarrierAggregationVariant struct {
	Fr1fdd_FR1TDD_CA_SpCellOnFR1FDD        *int64
	Fr1fdd_FR1TDD_CA_SpCellOnFR1TDD        *int64
	Fr1fdd_FR2TDD_CA_SpCellOnFR1FDD        *int64
	Fr1fdd_FR2TDD_CA_SpCellOnFR2TDD        *int64
	Fr1tdd_FR2TDD_CA_SpCellOnFR1TDD        *int64
	Fr1tdd_FR2TDD_CA_SpCellOnFR2TDD        *int64
	Fr1fdd_FR1TDD_FR2TDD_CA_SpCellOnFR1FDD *int64
	Fr1fdd_FR1TDD_FR2TDD_CA_SpCellOnFR1TDD *int64
	Fr1fdd_FR1TDD_FR2TDD_CA_SpCellOnFR2TDD *int64
}

func (ie *CarrierAggregationVariant) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(carrierAggregationVariantConstraints)

	// 1. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.Fr1fdd_FR1TDD_CA_SpCellOnFR1FDD != nil, ie.Fr1fdd_FR1TDD_CA_SpCellOnFR1TDD != nil, ie.Fr1fdd_FR2TDD_CA_SpCellOnFR1FDD != nil, ie.Fr1fdd_FR2TDD_CA_SpCellOnFR2TDD != nil, ie.Fr1tdd_FR2TDD_CA_SpCellOnFR1TDD != nil, ie.Fr1tdd_FR2TDD_CA_SpCellOnFR2TDD != nil, ie.Fr1fdd_FR1TDD_FR2TDD_CA_SpCellOnFR1FDD != nil, ie.Fr1fdd_FR1TDD_FR2TDD_CA_SpCellOnFR1TDD != nil, ie.Fr1fdd_FR1TDD_FR2TDD_CA_SpCellOnFR2TDD != nil}); err != nil {
		return err
	}

	// 2. fr1fdd-FR1TDD-CA-SpCellOnFR1FDD: enumerated
	{
		if ie.Fr1fdd_FR1TDD_CA_SpCellOnFR1FDD != nil {
			if err := e.EncodeEnumerated(*ie.Fr1fdd_FR1TDD_CA_SpCellOnFR1FDD, carrierAggregationVariantFr1fddFR1TDDCASpCellOnFR1FDDConstraints); err != nil {
				return err
			}
		}
	}

	// 3. fr1fdd-FR1TDD-CA-SpCellOnFR1TDD: enumerated
	{
		if ie.Fr1fdd_FR1TDD_CA_SpCellOnFR1TDD != nil {
			if err := e.EncodeEnumerated(*ie.Fr1fdd_FR1TDD_CA_SpCellOnFR1TDD, carrierAggregationVariantFr1fddFR1TDDCASpCellOnFR1TDDConstraints); err != nil {
				return err
			}
		}
	}

	// 4. fr1fdd-FR2TDD-CA-SpCellOnFR1FDD: enumerated
	{
		if ie.Fr1fdd_FR2TDD_CA_SpCellOnFR1FDD != nil {
			if err := e.EncodeEnumerated(*ie.Fr1fdd_FR2TDD_CA_SpCellOnFR1FDD, carrierAggregationVariantFr1fddFR2TDDCASpCellOnFR1FDDConstraints); err != nil {
				return err
			}
		}
	}

	// 5. fr1fdd-FR2TDD-CA-SpCellOnFR2TDD: enumerated
	{
		if ie.Fr1fdd_FR2TDD_CA_SpCellOnFR2TDD != nil {
			if err := e.EncodeEnumerated(*ie.Fr1fdd_FR2TDD_CA_SpCellOnFR2TDD, carrierAggregationVariantFr1fddFR2TDDCASpCellOnFR2TDDConstraints); err != nil {
				return err
			}
		}
	}

	// 6. fr1tdd-FR2TDD-CA-SpCellOnFR1TDD: enumerated
	{
		if ie.Fr1tdd_FR2TDD_CA_SpCellOnFR1TDD != nil {
			if err := e.EncodeEnumerated(*ie.Fr1tdd_FR2TDD_CA_SpCellOnFR1TDD, carrierAggregationVariantFr1tddFR2TDDCASpCellOnFR1TDDConstraints); err != nil {
				return err
			}
		}
	}

	// 7. fr1tdd-FR2TDD-CA-SpCellOnFR2TDD: enumerated
	{
		if ie.Fr1tdd_FR2TDD_CA_SpCellOnFR2TDD != nil {
			if err := e.EncodeEnumerated(*ie.Fr1tdd_FR2TDD_CA_SpCellOnFR2TDD, carrierAggregationVariantFr1tddFR2TDDCASpCellOnFR2TDDConstraints); err != nil {
				return err
			}
		}
	}

	// 8. fr1fdd-FR1TDD-FR2TDD-CA-SpCellOnFR1FDD: enumerated
	{
		if ie.Fr1fdd_FR1TDD_FR2TDD_CA_SpCellOnFR1FDD != nil {
			if err := e.EncodeEnumerated(*ie.Fr1fdd_FR1TDD_FR2TDD_CA_SpCellOnFR1FDD, carrierAggregationVariantFr1fddFR1TDDFR2TDDCASpCellOnFR1FDDConstraints); err != nil {
				return err
			}
		}
	}

	// 9. fr1fdd-FR1TDD-FR2TDD-CA-SpCellOnFR1TDD: enumerated
	{
		if ie.Fr1fdd_FR1TDD_FR2TDD_CA_SpCellOnFR1TDD != nil {
			if err := e.EncodeEnumerated(*ie.Fr1fdd_FR1TDD_FR2TDD_CA_SpCellOnFR1TDD, carrierAggregationVariantFr1fddFR1TDDFR2TDDCASpCellOnFR1TDDConstraints); err != nil {
				return err
			}
		}
	}

	// 10. fr1fdd-FR1TDD-FR2TDD-CA-SpCellOnFR2TDD: enumerated
	{
		if ie.Fr1fdd_FR1TDD_FR2TDD_CA_SpCellOnFR2TDD != nil {
			if err := e.EncodeEnumerated(*ie.Fr1fdd_FR1TDD_FR2TDD_CA_SpCellOnFR2TDD, carrierAggregationVariantFr1fddFR1TDDFR2TDDCASpCellOnFR2TDDConstraints); err != nil {
				return err
			}
		}
	}

	return nil
}

func (ie *CarrierAggregationVariant) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(carrierAggregationVariantConstraints)

	// 1. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 2. fr1fdd-FR1TDD-CA-SpCellOnFR1FDD: enumerated
	{
		if seq.IsComponentPresent(0) {
			idx, err := d.DecodeEnumerated(carrierAggregationVariantFr1fddFR1TDDCASpCellOnFR1FDDConstraints)
			if err != nil {
				return err
			}
			ie.Fr1fdd_FR1TDD_CA_SpCellOnFR1FDD = &idx
		}
	}

	// 3. fr1fdd-FR1TDD-CA-SpCellOnFR1TDD: enumerated
	{
		if seq.IsComponentPresent(1) {
			idx, err := d.DecodeEnumerated(carrierAggregationVariantFr1fddFR1TDDCASpCellOnFR1TDDConstraints)
			if err != nil {
				return err
			}
			ie.Fr1fdd_FR1TDD_CA_SpCellOnFR1TDD = &idx
		}
	}

	// 4. fr1fdd-FR2TDD-CA-SpCellOnFR1FDD: enumerated
	{
		if seq.IsComponentPresent(2) {
			idx, err := d.DecodeEnumerated(carrierAggregationVariantFr1fddFR2TDDCASpCellOnFR1FDDConstraints)
			if err != nil {
				return err
			}
			ie.Fr1fdd_FR2TDD_CA_SpCellOnFR1FDD = &idx
		}
	}

	// 5. fr1fdd-FR2TDD-CA-SpCellOnFR2TDD: enumerated
	{
		if seq.IsComponentPresent(3) {
			idx, err := d.DecodeEnumerated(carrierAggregationVariantFr1fddFR2TDDCASpCellOnFR2TDDConstraints)
			if err != nil {
				return err
			}
			ie.Fr1fdd_FR2TDD_CA_SpCellOnFR2TDD = &idx
		}
	}

	// 6. fr1tdd-FR2TDD-CA-SpCellOnFR1TDD: enumerated
	{
		if seq.IsComponentPresent(4) {
			idx, err := d.DecodeEnumerated(carrierAggregationVariantFr1tddFR2TDDCASpCellOnFR1TDDConstraints)
			if err != nil {
				return err
			}
			ie.Fr1tdd_FR2TDD_CA_SpCellOnFR1TDD = &idx
		}
	}

	// 7. fr1tdd-FR2TDD-CA-SpCellOnFR2TDD: enumerated
	{
		if seq.IsComponentPresent(5) {
			idx, err := d.DecodeEnumerated(carrierAggregationVariantFr1tddFR2TDDCASpCellOnFR2TDDConstraints)
			if err != nil {
				return err
			}
			ie.Fr1tdd_FR2TDD_CA_SpCellOnFR2TDD = &idx
		}
	}

	// 8. fr1fdd-FR1TDD-FR2TDD-CA-SpCellOnFR1FDD: enumerated
	{
		if seq.IsComponentPresent(6) {
			idx, err := d.DecodeEnumerated(carrierAggregationVariantFr1fddFR1TDDFR2TDDCASpCellOnFR1FDDConstraints)
			if err != nil {
				return err
			}
			ie.Fr1fdd_FR1TDD_FR2TDD_CA_SpCellOnFR1FDD = &idx
		}
	}

	// 9. fr1fdd-FR1TDD-FR2TDD-CA-SpCellOnFR1TDD: enumerated
	{
		if seq.IsComponentPresent(7) {
			idx, err := d.DecodeEnumerated(carrierAggregationVariantFr1fddFR1TDDFR2TDDCASpCellOnFR1TDDConstraints)
			if err != nil {
				return err
			}
			ie.Fr1fdd_FR1TDD_FR2TDD_CA_SpCellOnFR1TDD = &idx
		}
	}

	// 10. fr1fdd-FR1TDD-FR2TDD-CA-SpCellOnFR2TDD: enumerated
	{
		if seq.IsComponentPresent(8) {
			idx, err := d.DecodeEnumerated(carrierAggregationVariantFr1fddFR1TDDFR2TDDCASpCellOnFR2TDDConstraints)
			if err != nil {
				return err
			}
			ie.Fr1fdd_FR1TDD_FR2TDD_CA_SpCellOnFR2TDD = &idx
		}
	}

	return nil
}
