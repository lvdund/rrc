// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/rrc/common"
)

// Source: SL-DRX-ConfigGC-BC-r17 (line 27132).

var sLDRXConfigGCBCR17Constraints = per.SequenceConstraints{
	Extensible: true,
	RootComponents: []per.ComponentInfo{
		{Name: "sl-DRX-GC-BC-PerQoS-List-r17", Optional: true},
		{Name: "sl-DRX-GC-generic-r17", Optional: true},
		{Name: "sl-DefaultDRX-GC-BC-r17", Optional: true},
	},
}

var sLDRXConfigGCBCR17SlDRXGCBCPerQoSListR17Constraints = per.SizeRange(1, common.MaxSL_GC_BC_DRX_QoS_r17)

type SL_DRX_ConfigGC_BC_r17 struct {
	Sl_DRX_GC_BC_PerQoS_List_r17 []SL_DRX_GC_BC_QoS_r17
	Sl_DRX_GC_Generic_r17        *SL_DRX_GC_Generic_r17
	Sl_DefaultDRX_GC_BC_r17      *SL_DRX_GC_BC_QoS_r17
}

func (ie *SL_DRX_ConfigGC_BC_r17) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(sLDRXConfigGCBCR17Constraints)

	// 1. SEQUENCE extension bit.
	if err := seq.EncodeExtensionBit(false); err != nil {
		return err
	}

	// 2. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.Sl_DRX_GC_BC_PerQoS_List_r17 != nil, ie.Sl_DRX_GC_Generic_r17 != nil, ie.Sl_DefaultDRX_GC_BC_r17 != nil}); err != nil {
		return err
	}

	// 3. sl-DRX-GC-BC-PerQoS-List-r17: sequence-of
	{
		if ie.Sl_DRX_GC_BC_PerQoS_List_r17 != nil {
			seqOf := e.NewSequenceOfEncoder(sLDRXConfigGCBCR17SlDRXGCBCPerQoSListR17Constraints)
			if err := seqOf.EncodeLength(int64(len(ie.Sl_DRX_GC_BC_PerQoS_List_r17))); err != nil {
				return err
			}
			for i := range ie.Sl_DRX_GC_BC_PerQoS_List_r17 {
				if err := ie.Sl_DRX_GC_BC_PerQoS_List_r17[i].Encode(e); err != nil {
					return err
				}
			}
		}
	}

	// 4. sl-DRX-GC-generic-r17: ref
	{
		if ie.Sl_DRX_GC_Generic_r17 != nil {
			if err := ie.Sl_DRX_GC_Generic_r17.Encode(e); err != nil {
				return err
			}
		}
	}

	// 5. sl-DefaultDRX-GC-BC-r17: ref
	{
		if ie.Sl_DefaultDRX_GC_BC_r17 != nil {
			if err := ie.Sl_DefaultDRX_GC_BC_r17.Encode(e); err != nil {
				return err
			}
		}
	}

	return nil
}

func (ie *SL_DRX_ConfigGC_BC_r17) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(sLDRXConfigGCBCR17Constraints)

	// 1. SEQUENCE extension bit.
	if err := seq.DecodeExtensionBit(); err != nil {
		return err
	}

	// 2. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 3. sl-DRX-GC-BC-PerQoS-List-r17: sequence-of
	{
		if seq.IsComponentPresent(0) {
			seqOf := d.NewSequenceOfDecoder(sLDRXConfigGCBCR17SlDRXGCBCPerQoSListR17Constraints)
			n, err := seqOf.DecodeLength()
			if err != nil {
				return err
			}
			ie.Sl_DRX_GC_BC_PerQoS_List_r17 = make([]SL_DRX_GC_BC_QoS_r17, n)
			for i := int64(0); i < n; i++ {
				if err := ie.Sl_DRX_GC_BC_PerQoS_List_r17[i].Decode(d); err != nil {
					return err
				}
			}
		}
	}

	// 4. sl-DRX-GC-generic-r17: ref
	{
		if seq.IsComponentPresent(1) {
			ie.Sl_DRX_GC_Generic_r17 = new(SL_DRX_GC_Generic_r17)
			if err := ie.Sl_DRX_GC_Generic_r17.Decode(d); err != nil {
				return err
			}
		}
	}

	// 5. sl-DefaultDRX-GC-BC-r17: ref
	{
		if seq.IsComponentPresent(2) {
			ie.Sl_DefaultDRX_GC_BC_r17 = new(SL_DRX_GC_BC_QoS_r17)
			if err := ie.Sl_DefaultDRX_GC_BC_r17.Decode(d); err != nil {
				return err
			}
		}
	}

	return nil
}
