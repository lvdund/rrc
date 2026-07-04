// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: AdvancedReceiver-MU-MIMO-r18 (line 4952).

var advancedReceiverMUMIMOR18Constraints = per.SequenceConstraints{
	Extensible: true,
	RootComponents: []per.ComponentInfo{
		{Name: "precodingAndResourceAllocation-r18", Optional: true},
		{Name: "pdsch-TimeDomainAllocation-r18", Optional: true},
		{Name: "mcs-Table-r18", Optional: true},
		{Name: "advReceiver-MU-MIMO-DCI-1-1-r18", Optional: true},
	},
}

const (
	AdvancedReceiver_MU_MIMO_r18_Mcs_Table_r18_Qam1024 = 0
	AdvancedReceiver_MU_MIMO_r18_Mcs_Table_r18_Qam256  = 1
	AdvancedReceiver_MU_MIMO_r18_Mcs_Table_r18_Qam64   = 2
	AdvancedReceiver_MU_MIMO_r18_Mcs_Table_r18_Spare1  = 3
)

var advancedReceiverMUMIMOR18McsTableR18Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{AdvancedReceiver_MU_MIMO_r18_Mcs_Table_r18_Qam1024, AdvancedReceiver_MU_MIMO_r18_Mcs_Table_r18_Qam256, AdvancedReceiver_MU_MIMO_r18_Mcs_Table_r18_Qam64, AdvancedReceiver_MU_MIMO_r18_Mcs_Table_r18_Spare1},
}

const (
	AdvancedReceiver_MU_MIMO_r18_AdvReceiver_MU_MIMO_DCI_1_1_r18_Enabled = 0
)

var advancedReceiverMUMIMOR18AdvReceiverMUMIMODCI11R18Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{AdvancedReceiver_MU_MIMO_r18_AdvReceiver_MU_MIMO_DCI_1_1_r18_Enabled},
}

type AdvancedReceiver_MU_MIMO_r18 struct {
	PrecodingAndResourceAllocation_r18 *bool
	Pdsch_TimeDomainAllocation_r18     *bool
	Mcs_Table_r18                      *int64
	AdvReceiver_MU_MIMO_DCI_1_1_r18    *int64
}

func (ie *AdvancedReceiver_MU_MIMO_r18) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(advancedReceiverMUMIMOR18Constraints)

	// 1. SEQUENCE extension bit.
	if err := seq.EncodeExtensionBit(false); err != nil {
		return err
	}

	// 2. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.PrecodingAndResourceAllocation_r18 != nil, ie.Pdsch_TimeDomainAllocation_r18 != nil, ie.Mcs_Table_r18 != nil, ie.AdvReceiver_MU_MIMO_DCI_1_1_r18 != nil}); err != nil {
		return err
	}

	// 3. precodingAndResourceAllocation-r18: boolean
	{
		if ie.PrecodingAndResourceAllocation_r18 != nil {
			if err := e.EncodeBoolean(*ie.PrecodingAndResourceAllocation_r18); err != nil {
				return err
			}
		}
	}

	// 4. pdsch-TimeDomainAllocation-r18: boolean
	{
		if ie.Pdsch_TimeDomainAllocation_r18 != nil {
			if err := e.EncodeBoolean(*ie.Pdsch_TimeDomainAllocation_r18); err != nil {
				return err
			}
		}
	}

	// 5. mcs-Table-r18: enumerated
	{
		if ie.Mcs_Table_r18 != nil {
			if err := e.EncodeEnumerated(*ie.Mcs_Table_r18, advancedReceiverMUMIMOR18McsTableR18Constraints); err != nil {
				return err
			}
		}
	}

	// 6. advReceiver-MU-MIMO-DCI-1-1-r18: enumerated
	{
		if ie.AdvReceiver_MU_MIMO_DCI_1_1_r18 != nil {
			if err := e.EncodeEnumerated(*ie.AdvReceiver_MU_MIMO_DCI_1_1_r18, advancedReceiverMUMIMOR18AdvReceiverMUMIMODCI11R18Constraints); err != nil {
				return err
			}
		}
	}

	return nil
}

func (ie *AdvancedReceiver_MU_MIMO_r18) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(advancedReceiverMUMIMOR18Constraints)

	// 1. SEQUENCE extension bit.
	if err := seq.DecodeExtensionBit(); err != nil {
		return err
	}

	// 2. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 3. precodingAndResourceAllocation-r18: boolean
	{
		if seq.IsComponentPresent(0) {
			v, err := d.DecodeBoolean()
			if err != nil {
				return err
			}
			ie.PrecodingAndResourceAllocation_r18 = &v
		}
	}

	// 4. pdsch-TimeDomainAllocation-r18: boolean
	{
		if seq.IsComponentPresent(1) {
			v, err := d.DecodeBoolean()
			if err != nil {
				return err
			}
			ie.Pdsch_TimeDomainAllocation_r18 = &v
		}
	}

	// 5. mcs-Table-r18: enumerated
	{
		if seq.IsComponentPresent(2) {
			idx, err := d.DecodeEnumerated(advancedReceiverMUMIMOR18McsTableR18Constraints)
			if err != nil {
				return err
			}
			ie.Mcs_Table_r18 = &idx
		}
	}

	// 6. advReceiver-MU-MIMO-DCI-1-1-r18: enumerated
	{
		if seq.IsComponentPresent(3) {
			idx, err := d.DecodeEnumerated(advancedReceiverMUMIMOR18AdvReceiverMUMIMODCI11R18Constraints)
			if err != nil {
				return err
			}
			ie.AdvReceiver_MU_MIMO_DCI_1_1_r18 = &idx
		}
	}

	return nil
}
