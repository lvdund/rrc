// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/rrc/common"
)

// Source: DCP-Config-r16 (line 11728).

var dCPConfigR16Constraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "ps-RNTI-r16"},
		{Name: "ps-Offset-r16"},
		{Name: "sizeDCI-2-6-r16"},
		{Name: "ps-PositionDCI-2-6-r16"},
		{Name: "ps-WakeUp-r16", Optional: true},
		{Name: "ps-TransmitPeriodicL1-RSRP-r16", Optional: true},
		{Name: "ps-TransmitOtherPeriodicCSI-r16", Optional: true},
	},
}

var dCPConfigR16PsOffsetR16Constraints = per.Constrained(1, 120)

var dCPConfigR16SizeDCI26R16Constraints = per.Constrained(1, common.MaxDCI_2_6_Size_r16)

var dCPConfigR16PsPositionDCI26R16Constraints = per.Constrained(0, common.MaxDCI_2_6_Size_1_r16)

const (
	DCP_Config_r16_Ps_WakeUp_r16_True = 0
)

var dCPConfigR16PsWakeUpR16Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{DCP_Config_r16_Ps_WakeUp_r16_True},
}

const (
	DCP_Config_r16_Ps_TransmitPeriodicL1_RSRP_r16_True = 0
)

var dCPConfigR16PsTransmitPeriodicL1RSRPR16Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{DCP_Config_r16_Ps_TransmitPeriodicL1_RSRP_r16_True},
}

const (
	DCP_Config_r16_Ps_TransmitOtherPeriodicCSI_r16_True = 0
)

var dCPConfigR16PsTransmitOtherPeriodicCSIR16Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{DCP_Config_r16_Ps_TransmitOtherPeriodicCSI_r16_True},
}

type DCP_Config_r16 struct {
	Ps_RNTI_r16                     RNTI_Value
	Ps_Offset_r16                   int64
	SizeDCI_2_6_r16                 int64
	Ps_PositionDCI_2_6_r16          int64
	Ps_WakeUp_r16                   *int64
	Ps_TransmitPeriodicL1_RSRP_r16  *int64
	Ps_TransmitOtherPeriodicCSI_r16 *int64
}

func (ie *DCP_Config_r16) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(dCPConfigR16Constraints)

	// 1. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.Ps_WakeUp_r16 != nil, ie.Ps_TransmitPeriodicL1_RSRP_r16 != nil, ie.Ps_TransmitOtherPeriodicCSI_r16 != nil}); err != nil {
		return err
	}

	// 2. ps-RNTI-r16: ref
	{
		if err := ie.Ps_RNTI_r16.Encode(e); err != nil {
			return err
		}
	}

	// 3. ps-Offset-r16: integer
	{
		if err := e.EncodeInteger(ie.Ps_Offset_r16, dCPConfigR16PsOffsetR16Constraints); err != nil {
			return err
		}
	}

	// 4. sizeDCI-2-6-r16: integer
	{
		if err := e.EncodeInteger(ie.SizeDCI_2_6_r16, dCPConfigR16SizeDCI26R16Constraints); err != nil {
			return err
		}
	}

	// 5. ps-PositionDCI-2-6-r16: integer
	{
		if err := e.EncodeInteger(ie.Ps_PositionDCI_2_6_r16, dCPConfigR16PsPositionDCI26R16Constraints); err != nil {
			return err
		}
	}

	// 6. ps-WakeUp-r16: enumerated
	{
		if ie.Ps_WakeUp_r16 != nil {
			if err := e.EncodeEnumerated(*ie.Ps_WakeUp_r16, dCPConfigR16PsWakeUpR16Constraints); err != nil {
				return err
			}
		}
	}

	// 7. ps-TransmitPeriodicL1-RSRP-r16: enumerated
	{
		if ie.Ps_TransmitPeriodicL1_RSRP_r16 != nil {
			if err := e.EncodeEnumerated(*ie.Ps_TransmitPeriodicL1_RSRP_r16, dCPConfigR16PsTransmitPeriodicL1RSRPR16Constraints); err != nil {
				return err
			}
		}
	}

	// 8. ps-TransmitOtherPeriodicCSI-r16: enumerated
	{
		if ie.Ps_TransmitOtherPeriodicCSI_r16 != nil {
			if err := e.EncodeEnumerated(*ie.Ps_TransmitOtherPeriodicCSI_r16, dCPConfigR16PsTransmitOtherPeriodicCSIR16Constraints); err != nil {
				return err
			}
		}
	}

	return nil
}

func (ie *DCP_Config_r16) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(dCPConfigR16Constraints)

	// 1. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 2. ps-RNTI-r16: ref
	{
		if err := ie.Ps_RNTI_r16.Decode(d); err != nil {
			return err
		}
	}

	// 3. ps-Offset-r16: integer
	{
		v1, err := d.DecodeInteger(dCPConfigR16PsOffsetR16Constraints)
		if err != nil {
			return err
		}
		ie.Ps_Offset_r16 = v1
	}

	// 4. sizeDCI-2-6-r16: integer
	{
		v2, err := d.DecodeInteger(dCPConfigR16SizeDCI26R16Constraints)
		if err != nil {
			return err
		}
		ie.SizeDCI_2_6_r16 = v2
	}

	// 5. ps-PositionDCI-2-6-r16: integer
	{
		v3, err := d.DecodeInteger(dCPConfigR16PsPositionDCI26R16Constraints)
		if err != nil {
			return err
		}
		ie.Ps_PositionDCI_2_6_r16 = v3
	}

	// 6. ps-WakeUp-r16: enumerated
	{
		if seq.IsComponentPresent(4) {
			idx, err := d.DecodeEnumerated(dCPConfigR16PsWakeUpR16Constraints)
			if err != nil {
				return err
			}
			ie.Ps_WakeUp_r16 = &idx
		}
	}

	// 7. ps-TransmitPeriodicL1-RSRP-r16: enumerated
	{
		if seq.IsComponentPresent(5) {
			idx, err := d.DecodeEnumerated(dCPConfigR16PsTransmitPeriodicL1RSRPR16Constraints)
			if err != nil {
				return err
			}
			ie.Ps_TransmitPeriodicL1_RSRP_r16 = &idx
		}
	}

	// 8. ps-TransmitOtherPeriodicCSI-r16: enumerated
	{
		if seq.IsComponentPresent(6) {
			idx, err := d.DecodeEnumerated(dCPConfigR16PsTransmitOtherPeriodicCSIR16Constraints)
			if err != nil {
				return err
			}
			ie.Ps_TransmitOtherPeriodicCSI_r16 = &idx
		}
	}

	return nil
}
