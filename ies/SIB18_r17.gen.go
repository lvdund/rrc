// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/rrc/common"
)

// Source: SIB18-r17 (line 4490).

var sIB18R17Constraints = per.SequenceConstraints{
	Extensible: true,
	RootComponents: []per.ComponentInfo{
		{Name: "gin-ElementList-r17", Optional: true},
		{Name: "gins-PerSNPN-List-r17", Optional: true},
		{Name: "lateNonCriticalExtension", Optional: true},
	},
}

var sIB18R17GinElementListR17Constraints = per.SizeRange(1, common.MaxGIN_r17)

var sIB18R17GinsPerSNPNListR17Constraints = per.SizeRange(1, common.MaxNPN_r16)

var sIB18R17LateNonCriticalExtensionConstraints = per.SizeConstraints{}

type SIB18_r17 struct {
	Gin_ElementList_r17      []GIN_Element_r17
	Gins_PerSNPN_List_r17    []GINs_PerSNPN_r17
	LateNonCriticalExtension []byte
}

func (ie *SIB18_r17) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(sIB18R17Constraints)

	// 1. SEQUENCE extension bit.
	if err := seq.EncodeExtensionBit(false); err != nil {
		return err
	}

	// 2. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.Gin_ElementList_r17 != nil, ie.Gins_PerSNPN_List_r17 != nil, ie.LateNonCriticalExtension != nil}); err != nil {
		return err
	}

	// 3. gin-ElementList-r17: sequence-of
	{
		if ie.Gin_ElementList_r17 != nil {
			seqOf := e.NewSequenceOfEncoder(sIB18R17GinElementListR17Constraints)
			if err := seqOf.EncodeLength(int64(len(ie.Gin_ElementList_r17))); err != nil {
				return err
			}
			for i := range ie.Gin_ElementList_r17 {
				if err := ie.Gin_ElementList_r17[i].Encode(e); err != nil {
					return err
				}
			}
		}
	}

	// 4. gins-PerSNPN-List-r17: sequence-of
	{
		if ie.Gins_PerSNPN_List_r17 != nil {
			seqOf := e.NewSequenceOfEncoder(sIB18R17GinsPerSNPNListR17Constraints)
			if err := seqOf.EncodeLength(int64(len(ie.Gins_PerSNPN_List_r17))); err != nil {
				return err
			}
			for i := range ie.Gins_PerSNPN_List_r17 {
				if err := ie.Gins_PerSNPN_List_r17[i].Encode(e); err != nil {
					return err
				}
			}
		}
	}

	// 5. lateNonCriticalExtension: octet-string
	{
		if ie.LateNonCriticalExtension != nil {
			if err := e.EncodeOctetString(ie.LateNonCriticalExtension, sIB18R17LateNonCriticalExtensionConstraints); err != nil {
				return err
			}
		}
	}

	return nil
}

func (ie *SIB18_r17) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(sIB18R17Constraints)

	// 1. SEQUENCE extension bit.
	if err := seq.DecodeExtensionBit(); err != nil {
		return err
	}

	// 2. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 3. gin-ElementList-r17: sequence-of
	{
		if seq.IsComponentPresent(0) {
			seqOf := d.NewSequenceOfDecoder(sIB18R17GinElementListR17Constraints)
			n, err := seqOf.DecodeLength()
			if err != nil {
				return err
			}
			ie.Gin_ElementList_r17 = make([]GIN_Element_r17, n)
			for i := int64(0); i < n; i++ {
				if err := ie.Gin_ElementList_r17[i].Decode(d); err != nil {
					return err
				}
			}
		}
	}

	// 4. gins-PerSNPN-List-r17: sequence-of
	{
		if seq.IsComponentPresent(1) {
			seqOf := d.NewSequenceOfDecoder(sIB18R17GinsPerSNPNListR17Constraints)
			n, err := seqOf.DecodeLength()
			if err != nil {
				return err
			}
			ie.Gins_PerSNPN_List_r17 = make([]GINs_PerSNPN_r17, n)
			for i := int64(0); i < n; i++ {
				if err := ie.Gins_PerSNPN_List_r17[i].Decode(d); err != nil {
					return err
				}
			}
		}
	}

	// 5. lateNonCriticalExtension: octet-string
	{
		if seq.IsComponentPresent(2) {
			v, err := d.DecodeOctetString(sIB18R17LateNonCriticalExtensionConstraints)
			if err != nil {
				return err
			}
			ie.LateNonCriticalExtension = v
		}
	}

	return nil
}
