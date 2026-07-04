// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: CGI-InfoNR (line 5938).

var cGIInfoNRConstraints = per.SequenceConstraints{
	Extensible: true,
	RootComponents: []per.ComponentInfo{
		{Name: "plmn-IdentityInfoList", Optional: true},
		{Name: "frequencyBandList", Optional: true},
		{Name: "noSIB1", Optional: true},
	},
	ExtComponents: []per.ComponentInfo{
		{Name: "extension-addition-0", Optional: true},
		{Name: "extension-addition-1", Optional: true},
		{Name: "extension-addition-2", Optional: true},
	},
}

const (
	CGI_InfoNR_Ext_CellReservedForOtherUse_r16_True = 0
)

var cGIInfoNRExtCellReservedForOtherUseR16Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{CGI_InfoNR_Ext_CellReservedForOtherUse_r16_True},
}

const (
	CGI_InfoNR_Ext_Hsdn_Cell_r19_True = 0
)

var cGIInfoNRExtHsdnCellR19Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{CGI_InfoNR_Ext_Hsdn_Cell_r19_True},
}

type CGI_InfoNR struct {
	Plmn_IdentityInfoList *PLMN_IdentityInfoList
	FrequencyBandList     *MultiFrequencyBandListNR
	NoSIB1                *struct {
		Ssb_SubcarrierOffset int64
		Pdcch_ConfigSIB1     PDCCH_ConfigSIB1
	}
	Npn_IdentityInfoList_r16    *NPN_IdentityInfoList_r16
	CellReservedForOtherUse_r16 *int64
	Hsdn_Cell_r19               *int64
}

func (ie *CGI_InfoNR) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(cGIInfoNRConstraints)

	// Extension-group presence (one var per [[...]] group).
	hasExtGroup0 := ie.Npn_IdentityInfoList_r16 != nil
	hasExtGroup1 := ie.CellReservedForOtherUse_r16 != nil
	hasExtGroup2 := ie.Hsdn_Cell_r19 != nil
	hasExtensions := hasExtGroup0 || hasExtGroup1 || hasExtGroup2

	// 1. SEQUENCE extension bit.
	if err := seq.EncodeExtensionBit(hasExtensions); err != nil {
		return err
	}

	// 2. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.Plmn_IdentityInfoList != nil, ie.FrequencyBandList != nil, ie.NoSIB1 != nil}); err != nil {
		return err
	}

	// 3. plmn-IdentityInfoList: ref
	{
		if ie.Plmn_IdentityInfoList != nil {
			if err := ie.Plmn_IdentityInfoList.Encode(e); err != nil {
				return err
			}
		}
	}

	// 4. frequencyBandList: ref
	{
		if ie.FrequencyBandList != nil {
			if err := ie.FrequencyBandList.Encode(e); err != nil {
				return err
			}
		}
	}

	// 5. noSIB1: sequence
	{
		if ie.NoSIB1 != nil {
			c := ie.NoSIB1
			if err := e.EncodeInteger(c.Ssb_SubcarrierOffset, per.Constrained(0, 15)); err != nil {
				return err
			}
			if err := c.Pdcch_ConfigSIB1.Encode(e); err != nil {
				return err
			}
		}
	}

	if hasExtensions {
		extPresence := []bool{hasExtGroup0, hasExtGroup1, hasExtGroup2}
		var extValues [][]byte

		// Extension group 0:
		if hasExtGroup0 {
			ex := per.NewEncoder(e.Variant())
			groupSeq := ex.NewSequenceEncoder(per.SequenceConstraints{
				Extensible: false,
				RootComponents: []per.ComponentInfo{
					{Name: "npn-IdentityInfoList-r16", Optional: true},
				},
			})
			if err := groupSeq.EncodePreamble([]bool{ie.Npn_IdentityInfoList_r16 != nil}); err != nil {
				return err
			}

			if ie.Npn_IdentityInfoList_r16 != nil {
				if err := ie.Npn_IdentityInfoList_r16.Encode(ex); err != nil {
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
					{Name: "cellReservedForOtherUse-r16", Optional: true},
				},
			})
			if err := groupSeq.EncodePreamble([]bool{ie.CellReservedForOtherUse_r16 != nil}); err != nil {
				return err
			}

			if ie.CellReservedForOtherUse_r16 != nil {
				if err := ex.EncodeEnumerated(*ie.CellReservedForOtherUse_r16, cGIInfoNRExtCellReservedForOtherUseR16Constraints); err != nil {
					return err
				}
			}
			extValues = append(extValues, ex.Bytes())
		}

		// Extension group 2:
		if hasExtGroup2 {
			ex := per.NewEncoder(e.Variant())
			groupSeq := ex.NewSequenceEncoder(per.SequenceConstraints{
				Extensible: false,
				RootComponents: []per.ComponentInfo{
					{Name: "hsdn-Cell-r19", Optional: true},
				},
			})
			if err := groupSeq.EncodePreamble([]bool{ie.Hsdn_Cell_r19 != nil}); err != nil {
				return err
			}

			if ie.Hsdn_Cell_r19 != nil {
				if err := ex.EncodeEnumerated(*ie.Hsdn_Cell_r19, cGIInfoNRExtHsdnCellR19Constraints); err != nil {
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

func (ie *CGI_InfoNR) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(cGIInfoNRConstraints)

	// 1. SEQUENCE extension bit.
	if err := seq.DecodeExtensionBit(); err != nil {
		return err
	}

	// 2. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 3. plmn-IdentityInfoList: ref
	{
		if seq.IsComponentPresent(0) {
			ie.Plmn_IdentityInfoList = new(PLMN_IdentityInfoList)
			if err := ie.Plmn_IdentityInfoList.Decode(d); err != nil {
				return err
			}
		}
	}

	// 4. frequencyBandList: ref
	{
		if seq.IsComponentPresent(1) {
			ie.FrequencyBandList = new(MultiFrequencyBandListNR)
			if err := ie.FrequencyBandList.Decode(d); err != nil {
				return err
			}
		}
	}

	// 5. noSIB1: sequence
	{
		if seq.IsComponentPresent(2) {
			ie.NoSIB1 = &struct {
				Ssb_SubcarrierOffset int64
				Pdcch_ConfigSIB1     PDCCH_ConfigSIB1
			}{}
			c := ie.NoSIB1
			{
				v, err := d.DecodeInteger(per.Constrained(0, 15))
				if err != nil {
					return err
				}
				c.Ssb_SubcarrierOffset = v
			}
			{
				if err := c.Pdcch_ConfigSIB1.Decode(d); err != nil {
					return err
				}
			}
		}
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
				{Name: "npn-IdentityInfoList-r16", Optional: true},
			},
		})
		if err := groupSeq.DecodePreamble(); err != nil {
			return err
		}

		if groupSeq.IsComponentPresent(0) {
			ie.Npn_IdentityInfoList_r16 = new(NPN_IdentityInfoList_r16)
			if err := ie.Npn_IdentityInfoList_r16.Decode(dx); err != nil {
				return err
			}
		}
	}

	// Extension group 1:
	if seq.IsExtensionPresent(1) {
		dx := per.NewDecoder(extValues[extValueIndex], d.Variant())
		extValueIndex++
		groupSeq := dx.NewSequenceDecoder(per.SequenceConstraints{
			Extensible: false,
			RootComponents: []per.ComponentInfo{
				{Name: "cellReservedForOtherUse-r16", Optional: true},
			},
		})
		if err := groupSeq.DecodePreamble(); err != nil {
			return err
		}

		if groupSeq.IsComponentPresent(0) {
			v, err := dx.DecodeEnumerated(cGIInfoNRExtCellReservedForOtherUseR16Constraints)
			if err != nil {
				return err
			}
			ie.CellReservedForOtherUse_r16 = &v
		}
	}

	// Extension group 2:
	if seq.IsExtensionPresent(2) {
		dx := per.NewDecoder(extValues[extValueIndex], d.Variant())
		extValueIndex++
		groupSeq := dx.NewSequenceDecoder(per.SequenceConstraints{
			Extensible: false,
			RootComponents: []per.ComponentInfo{
				{Name: "hsdn-Cell-r19", Optional: true},
			},
		})
		if err := groupSeq.DecodePreamble(); err != nil {
			return err
		}

		if groupSeq.IsComponentPresent(0) {
			v, err := dx.DecodeEnumerated(cGIInfoNRExtHsdnCellR19Constraints)
			if err != nil {
				return err
			}
			ie.Hsdn_Cell_r19 = &v
		}
	}

	return nil
}
