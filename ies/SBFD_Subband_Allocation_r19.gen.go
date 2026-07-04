// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: SBFD-Subband-Allocation-r19 (line 14347).

var sBFDSubbandAllocationR19Constraints = per.SequenceConstraints{
	Extensible: true,
	RootComponents: []per.ComponentInfo{
		{Name: "ul-subbandlocationAndBandwidth-r19", Optional: true},
		{Name: "firstDL-subbandlocationAndBandwidth-r19", Optional: true},
		{Name: "secondDL-subbandlocationAndBandwidth-r19", Optional: true},
	},
}

var sBFDSubbandAllocationR19UlSubbandlocationAndBandwidthR19Constraints = per.Constrained(0, 37949)

var sBFDSubbandAllocationR19FirstDLSubbandlocationAndBandwidthR19Constraints = per.Constrained(0, 37949)

var sBFDSubbandAllocationR19SecondDLSubbandlocationAndBandwidthR19Constraints = per.Constrained(0, 37949)

type SBFD_Subband_Allocation_r19 struct {
	Ul_SubbandlocationAndBandwidth_r19       *int64
	FirstDL_SubbandlocationAndBandwidth_r19  *int64
	SecondDL_SubbandlocationAndBandwidth_r19 *int64
}

func (ie *SBFD_Subband_Allocation_r19) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(sBFDSubbandAllocationR19Constraints)

	// 1. SEQUENCE extension bit.
	if err := seq.EncodeExtensionBit(false); err != nil {
		return err
	}

	// 2. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.Ul_SubbandlocationAndBandwidth_r19 != nil, ie.FirstDL_SubbandlocationAndBandwidth_r19 != nil, ie.SecondDL_SubbandlocationAndBandwidth_r19 != nil}); err != nil {
		return err
	}

	// 3. ul-subbandlocationAndBandwidth-r19: integer
	{
		if ie.Ul_SubbandlocationAndBandwidth_r19 != nil {
			if err := e.EncodeInteger(*ie.Ul_SubbandlocationAndBandwidth_r19, sBFDSubbandAllocationR19UlSubbandlocationAndBandwidthR19Constraints); err != nil {
				return err
			}
		}
	}

	// 4. firstDL-subbandlocationAndBandwidth-r19: integer
	{
		if ie.FirstDL_SubbandlocationAndBandwidth_r19 != nil {
			if err := e.EncodeInteger(*ie.FirstDL_SubbandlocationAndBandwidth_r19, sBFDSubbandAllocationR19FirstDLSubbandlocationAndBandwidthR19Constraints); err != nil {
				return err
			}
		}
	}

	// 5. secondDL-subbandlocationAndBandwidth-r19: integer
	{
		if ie.SecondDL_SubbandlocationAndBandwidth_r19 != nil {
			if err := e.EncodeInteger(*ie.SecondDL_SubbandlocationAndBandwidth_r19, sBFDSubbandAllocationR19SecondDLSubbandlocationAndBandwidthR19Constraints); err != nil {
				return err
			}
		}
	}

	return nil
}

func (ie *SBFD_Subband_Allocation_r19) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(sBFDSubbandAllocationR19Constraints)

	// 1. SEQUENCE extension bit.
	if err := seq.DecodeExtensionBit(); err != nil {
		return err
	}

	// 2. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 3. ul-subbandlocationAndBandwidth-r19: integer
	{
		if seq.IsComponentPresent(0) {
			v, err := d.DecodeInteger(sBFDSubbandAllocationR19UlSubbandlocationAndBandwidthR19Constraints)
			if err != nil {
				return err
			}
			ie.Ul_SubbandlocationAndBandwidth_r19 = &v
		}
	}

	// 4. firstDL-subbandlocationAndBandwidth-r19: integer
	{
		if seq.IsComponentPresent(1) {
			v, err := d.DecodeInteger(sBFDSubbandAllocationR19FirstDLSubbandlocationAndBandwidthR19Constraints)
			if err != nil {
				return err
			}
			ie.FirstDL_SubbandlocationAndBandwidth_r19 = &v
		}
	}

	// 5. secondDL-subbandlocationAndBandwidth-r19: integer
	{
		if seq.IsComponentPresent(2) {
			v, err := d.DecodeInteger(sBFDSubbandAllocationR19SecondDLSubbandlocationAndBandwidthR19Constraints)
			if err != nil {
				return err
			}
			ie.SecondDL_SubbandlocationAndBandwidth_r19 = &v
		}
	}

	return nil
}
