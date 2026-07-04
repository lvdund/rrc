// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/rrc/common"
)

// Source: SCS-SpecificCarrier (line 14334).

var sCSSpecificCarrierConstraints = per.SequenceConstraints{
	Extensible: true,
	RootComponents: []per.ComponentInfo{
		{Name: "offsetToCarrier"},
		{Name: "subcarrierSpacing"},
		{Name: "carrierBandwidth"},
	},
	ExtComponents: []per.ComponentInfo{
		{Name: "extension-addition-0", Optional: true},
		{Name: "extension-addition-1", Optional: true},
	},
}

var sCSSpecificCarrierOffsetToCarrierConstraints = per.Constrained(0, 2199)

var sCSSpecificCarrierCarrierBandwidthConstraints = per.Constrained(1, common.MaxNrofPhysicalResourceBlocks)

var sCSSpecificCarrierTxDirectCurrentLocationConstraints = per.Constrained(0, 4095)

type SCS_SpecificCarrier struct {
	OffsetToCarrier             int64
	SubcarrierSpacing           SubcarrierSpacing
	CarrierBandwidth            int64
	TxDirectCurrentLocation     *int64
	Sbfd_Subband_Allocation_r19 *SBFD_Subband_Allocation_r19
}

func (ie *SCS_SpecificCarrier) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(sCSSpecificCarrierConstraints)

	// Extension-group presence (one var per [[...]] group).
	hasExtGroup0 := ie.TxDirectCurrentLocation != nil
	hasExtGroup1 := ie.Sbfd_Subband_Allocation_r19 != nil
	hasExtensions := hasExtGroup0 || hasExtGroup1

	// 1. SEQUENCE extension bit.
	if err := seq.EncodeExtensionBit(hasExtensions); err != nil {
		return err
	}

	// 2. offsetToCarrier: integer
	{
		if err := e.EncodeInteger(ie.OffsetToCarrier, sCSSpecificCarrierOffsetToCarrierConstraints); err != nil {
			return err
		}
	}

	// 3. subcarrierSpacing: ref
	{
		if err := ie.SubcarrierSpacing.Encode(e); err != nil {
			return err
		}
	}

	// 4. carrierBandwidth: integer
	{
		if err := e.EncodeInteger(ie.CarrierBandwidth, sCSSpecificCarrierCarrierBandwidthConstraints); err != nil {
			return err
		}
	}

	if hasExtensions {
		extPresence := []bool{hasExtGroup0, hasExtGroup1}
		var extValues [][]byte

		// Extension group 0:
		if hasExtGroup0 {
			ex := per.NewEncoder(e.Variant())
			groupSeq := ex.NewSequenceEncoder(per.SequenceConstraints{
				Extensible: false,
				RootComponents: []per.ComponentInfo{
					{Name: "txDirectCurrentLocation", Optional: true},
				},
			})
			if err := groupSeq.EncodePreamble([]bool{ie.TxDirectCurrentLocation != nil}); err != nil {
				return err
			}

			if ie.TxDirectCurrentLocation != nil {
				if err := ex.EncodeInteger(*ie.TxDirectCurrentLocation, sCSSpecificCarrierTxDirectCurrentLocationConstraints); err != nil {
					return err
				}
			}
			extValues = append(extValues, ex.Bytes())
		}

		// Extension group 1:
		if hasExtGroup1 {
			ex := per.NewEncoder(e.Variant())
			groupSeq := ex.NewSequenceEncoder(per.SequenceConstraints{
				Extensible: false,
				RootComponents: []per.ComponentInfo{
					{Name: "sbfd-Subband-Allocation-r19", Optional: true},
				},
			})
			if err := groupSeq.EncodePreamble([]bool{ie.Sbfd_Subband_Allocation_r19 != nil}); err != nil {
				return err
			}

			if ie.Sbfd_Subband_Allocation_r19 != nil {
				if err := ie.Sbfd_Subband_Allocation_r19.Encode(ex); err != nil {
					return err
				}
			}
			extValues = append(extValues, ex.Bytes())
		}

		if err := seq.EncodeExtensionAdditions(extPresence, extValues); err != nil {
			return err
		}
	}

	return nil
}

func (ie *SCS_SpecificCarrier) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(sCSSpecificCarrierConstraints)

	// 1. SEQUENCE extension bit.
	if err := seq.DecodeExtensionBit(); err != nil {
		return err
	}

	// 2. offsetToCarrier: integer
	{
		v0, err := d.DecodeInteger(sCSSpecificCarrierOffsetToCarrierConstraints)
		if err != nil {
			return err
		}
		ie.OffsetToCarrier = v0
	}

	// 3. subcarrierSpacing: ref
	{
		if err := ie.SubcarrierSpacing.Decode(d); err != nil {
			return err
		}
	}

	// 4. carrierBandwidth: integer
	{
		v2, err := d.DecodeInteger(sCSSpecificCarrierCarrierBandwidthConstraints)
		if err != nil {
			return err
		}
		ie.CarrierBandwidth = v2
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
				{Name: "txDirectCurrentLocation", Optional: true},
			},
		})
		if err := groupSeq.DecodePreamble(); err != nil {
			return err
		}

		if groupSeq.IsComponentPresent(0) {
			v, err := dx.DecodeInteger(sCSSpecificCarrierTxDirectCurrentLocationConstraints)
			if err != nil {
				return err
			}
			ie.TxDirectCurrentLocation = &v
		}
	}

	// Extension group 1:
	if seq.IsExtensionPresent(1) {
		dx := per.NewDecoder(extValues[extValueIndex], d.Variant())
		extValueIndex++
		groupSeq := dx.NewSequenceDecoder(per.SequenceConstraints{
			Extensible: false,
			RootComponents: []per.ComponentInfo{
				{Name: "sbfd-Subband-Allocation-r19", Optional: true},
			},
		})
		if err := groupSeq.DecodePreamble(); err != nil {
			return err
		}

		if groupSeq.IsComponentPresent(0) {
			ie.Sbfd_Subband_Allocation_r19 = new(SBFD_Subband_Allocation_r19)
			if err := ie.Sbfd_Subband_Allocation_r19.Decode(dx); err != nil {
				return err
			}
		}
	}

	return nil
}
