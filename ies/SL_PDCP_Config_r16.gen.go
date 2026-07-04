// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: SL-PDCP-Config-r16 (line 27596).

var sLPDCPConfigR16Constraints = per.SequenceConstraints{
	Extensible: true,
	RootComponents: []per.ComponentInfo{
		{Name: "sl-DiscardTimer-r16", Optional: true},
		{Name: "sl-PDCP-SN-Size-r16", Optional: true},
		{Name: "sl-OutOfOrderDelivery", Optional: true},
	},
}

const (
	SL_PDCP_Config_r16_Sl_DiscardTimer_r16_Ms3      = 0
	SL_PDCP_Config_r16_Sl_DiscardTimer_r16_Ms10     = 1
	SL_PDCP_Config_r16_Sl_DiscardTimer_r16_Ms20     = 2
	SL_PDCP_Config_r16_Sl_DiscardTimer_r16_Ms25     = 3
	SL_PDCP_Config_r16_Sl_DiscardTimer_r16_Ms30     = 4
	SL_PDCP_Config_r16_Sl_DiscardTimer_r16_Ms40     = 5
	SL_PDCP_Config_r16_Sl_DiscardTimer_r16_Ms50     = 6
	SL_PDCP_Config_r16_Sl_DiscardTimer_r16_Ms60     = 7
	SL_PDCP_Config_r16_Sl_DiscardTimer_r16_Ms75     = 8
	SL_PDCP_Config_r16_Sl_DiscardTimer_r16_Ms100    = 9
	SL_PDCP_Config_r16_Sl_DiscardTimer_r16_Ms150    = 10
	SL_PDCP_Config_r16_Sl_DiscardTimer_r16_Ms200    = 11
	SL_PDCP_Config_r16_Sl_DiscardTimer_r16_Ms250    = 12
	SL_PDCP_Config_r16_Sl_DiscardTimer_r16_Ms300    = 13
	SL_PDCP_Config_r16_Sl_DiscardTimer_r16_Ms500    = 14
	SL_PDCP_Config_r16_Sl_DiscardTimer_r16_Ms750    = 15
	SL_PDCP_Config_r16_Sl_DiscardTimer_r16_Ms1500   = 16
	SL_PDCP_Config_r16_Sl_DiscardTimer_r16_Infinity = 17
)

var sLPDCPConfigR16SlDiscardTimerR16Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{SL_PDCP_Config_r16_Sl_DiscardTimer_r16_Ms3, SL_PDCP_Config_r16_Sl_DiscardTimer_r16_Ms10, SL_PDCP_Config_r16_Sl_DiscardTimer_r16_Ms20, SL_PDCP_Config_r16_Sl_DiscardTimer_r16_Ms25, SL_PDCP_Config_r16_Sl_DiscardTimer_r16_Ms30, SL_PDCP_Config_r16_Sl_DiscardTimer_r16_Ms40, SL_PDCP_Config_r16_Sl_DiscardTimer_r16_Ms50, SL_PDCP_Config_r16_Sl_DiscardTimer_r16_Ms60, SL_PDCP_Config_r16_Sl_DiscardTimer_r16_Ms75, SL_PDCP_Config_r16_Sl_DiscardTimer_r16_Ms100, SL_PDCP_Config_r16_Sl_DiscardTimer_r16_Ms150, SL_PDCP_Config_r16_Sl_DiscardTimer_r16_Ms200, SL_PDCP_Config_r16_Sl_DiscardTimer_r16_Ms250, SL_PDCP_Config_r16_Sl_DiscardTimer_r16_Ms300, SL_PDCP_Config_r16_Sl_DiscardTimer_r16_Ms500, SL_PDCP_Config_r16_Sl_DiscardTimer_r16_Ms750, SL_PDCP_Config_r16_Sl_DiscardTimer_r16_Ms1500, SL_PDCP_Config_r16_Sl_DiscardTimer_r16_Infinity},
}

const (
	SL_PDCP_Config_r16_Sl_PDCP_SN_Size_r16_Len12bits = 0
	SL_PDCP_Config_r16_Sl_PDCP_SN_Size_r16_Len18bits = 1
)

var sLPDCPConfigR16SlPDCPSNSizeR16Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{SL_PDCP_Config_r16_Sl_PDCP_SN_Size_r16_Len12bits, SL_PDCP_Config_r16_Sl_PDCP_SN_Size_r16_Len18bits},
}

const (
	SL_PDCP_Config_r16_Sl_OutOfOrderDelivery_True = 0
)

var sLPDCPConfigR16SlOutOfOrderDeliveryConstraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{SL_PDCP_Config_r16_Sl_OutOfOrderDelivery_True},
}

type SL_PDCP_Config_r16 struct {
	Sl_DiscardTimer_r16   *int64
	Sl_PDCP_SN_Size_r16   *int64
	Sl_OutOfOrderDelivery *int64
}

func (ie *SL_PDCP_Config_r16) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(sLPDCPConfigR16Constraints)

	// 1. SEQUENCE extension bit.
	if err := seq.EncodeExtensionBit(false); err != nil {
		return err
	}

	// 2. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.Sl_DiscardTimer_r16 != nil, ie.Sl_PDCP_SN_Size_r16 != nil, ie.Sl_OutOfOrderDelivery != nil}); err != nil {
		return err
	}

	// 3. sl-DiscardTimer-r16: enumerated
	{
		if ie.Sl_DiscardTimer_r16 != nil {
			if err := e.EncodeEnumerated(*ie.Sl_DiscardTimer_r16, sLPDCPConfigR16SlDiscardTimerR16Constraints); err != nil {
				return err
			}
		}
	}

	// 4. sl-PDCP-SN-Size-r16: enumerated
	{
		if ie.Sl_PDCP_SN_Size_r16 != nil {
			if err := e.EncodeEnumerated(*ie.Sl_PDCP_SN_Size_r16, sLPDCPConfigR16SlPDCPSNSizeR16Constraints); err != nil {
				return err
			}
		}
	}

	// 5. sl-OutOfOrderDelivery: enumerated
	{
		if ie.Sl_OutOfOrderDelivery != nil {
			if err := e.EncodeEnumerated(*ie.Sl_OutOfOrderDelivery, sLPDCPConfigR16SlOutOfOrderDeliveryConstraints); err != nil {
				return err
			}
		}
	}

	return nil
}

func (ie *SL_PDCP_Config_r16) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(sLPDCPConfigR16Constraints)

	// 1. SEQUENCE extension bit.
	if err := seq.DecodeExtensionBit(); err != nil {
		return err
	}

	// 2. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 3. sl-DiscardTimer-r16: enumerated
	{
		if seq.IsComponentPresent(0) {
			idx, err := d.DecodeEnumerated(sLPDCPConfigR16SlDiscardTimerR16Constraints)
			if err != nil {
				return err
			}
			ie.Sl_DiscardTimer_r16 = &idx
		}
	}

	// 4. sl-PDCP-SN-Size-r16: enumerated
	{
		if seq.IsComponentPresent(1) {
			idx, err := d.DecodeEnumerated(sLPDCPConfigR16SlPDCPSNSizeR16Constraints)
			if err != nil {
				return err
			}
			ie.Sl_PDCP_SN_Size_r16 = &idx
		}
	}

	// 5. sl-OutOfOrderDelivery: enumerated
	{
		if seq.IsComponentPresent(2) {
			idx, err := d.DecodeEnumerated(sLPDCPConfigR16SlOutOfOrderDeliveryConstraints)
			if err != nil {
				return err
			}
			ie.Sl_OutOfOrderDelivery = &idx
		}
	}

	return nil
}
