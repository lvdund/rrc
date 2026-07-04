// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/rrc/common"
)

// Source: UE-CapabilityRequestFilterCommon (line 25482).

var uECapabilityRequestFilterCommonConstraints = per.SequenceConstraints{
	Extensible: true,
	RootComponents: []per.ComponentInfo{
		{Name: "mrdc-Request", Optional: true},
	},
	ExtComponents: []per.ComponentInfo{
		{Name: "extension-addition-0", Optional: true},
		{Name: "extension-addition-1", Optional: true},
		{Name: "extension-addition-2", Optional: true},
		{Name: "extension-addition-3", Optional: true},
	},
}

const (
	UE_CapabilityRequestFilterCommon_Ext_UplinkTxSwitchRequest_r16_True = 0
)

var uECapabilityRequestFilterCommonExtUplinkTxSwitchRequestR16Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{UE_CapabilityRequestFilterCommon_Ext_UplinkTxSwitchRequest_r16_True},
}

var uECapabilityRequestFilterCommonExtRequestedCellGroupingR16Constraints = per.SizeRange(1, common.MaxCellGroupings_r16)

const (
	UE_CapabilityRequestFilterCommon_Ext_FallbackGroupFiveRequest_r17_True = 0
)

var uECapabilityRequestFilterCommonExtFallbackGroupFiveRequestR17Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{UE_CapabilityRequestFilterCommon_Ext_FallbackGroupFiveRequest_r17_True},
}

var uECapabilityRequestFilterCommonMrdcRequestConstraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "omitEN-DC", Optional: true},
		{Name: "includeNR-DC", Optional: true},
		{Name: "includeNE-DC", Optional: true},
	},
}

const (
	UE_CapabilityRequestFilterCommon_Mrdc_Request_OmitEN_DC_True = 0
)

var uECapabilityRequestFilterCommonMrdcRequestOmitENDCConstraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{UE_CapabilityRequestFilterCommon_Mrdc_Request_OmitEN_DC_True},
}

const (
	UE_CapabilityRequestFilterCommon_Mrdc_Request_IncludeNR_DC_True = 0
)

var uECapabilityRequestFilterCommonMrdcRequestIncludeNRDCConstraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{UE_CapabilityRequestFilterCommon_Mrdc_Request_IncludeNR_DC_True},
}

const (
	UE_CapabilityRequestFilterCommon_Mrdc_Request_IncludeNE_DC_True = 0
)

var uECapabilityRequestFilterCommonMrdcRequestIncludeNEDCConstraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{UE_CapabilityRequestFilterCommon_Mrdc_Request_IncludeNE_DC_True},
}

var uECapabilityRequestFilterCommonExtCodebookTypeRequestR16Constraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "type1-SinglePanel-r16", Optional: true},
		{Name: "type1-MultiPanel-r16", Optional: true},
		{Name: "type2-r16", Optional: true},
		{Name: "type2-PortSelection-r16", Optional: true},
	},
}

const (
	UE_CapabilityRequestFilterCommon_Ext_CodebookTypeRequest_r16_Type1_SinglePanel_r16_True = 0
)

var uECapabilityRequestFilterCommonExtCodebookTypeRequestR16Type1SinglePanelR16Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{UE_CapabilityRequestFilterCommon_Ext_CodebookTypeRequest_r16_Type1_SinglePanel_r16_True},
}

const (
	UE_CapabilityRequestFilterCommon_Ext_CodebookTypeRequest_r16_Type1_MultiPanel_r16_True = 0
)

var uECapabilityRequestFilterCommonExtCodebookTypeRequestR16Type1MultiPanelR16Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{UE_CapabilityRequestFilterCommon_Ext_CodebookTypeRequest_r16_Type1_MultiPanel_r16_True},
}

const (
	UE_CapabilityRequestFilterCommon_Ext_CodebookTypeRequest_r16_Type2_r16_True = 0
)

var uECapabilityRequestFilterCommonExtCodebookTypeRequestR16Type2R16Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{UE_CapabilityRequestFilterCommon_Ext_CodebookTypeRequest_r16_Type2_r16_True},
}

const (
	UE_CapabilityRequestFilterCommon_Ext_CodebookTypeRequest_r16_Type2_PortSelection_r16_True = 0
)

var uECapabilityRequestFilterCommonExtCodebookTypeRequestR16Type2PortSelectionR16Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{UE_CapabilityRequestFilterCommon_Ext_CodebookTypeRequest_r16_Type2_PortSelection_r16_True},
}

var uECapabilityRequestFilterCommonExtLowerMSDRequestR18Constraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "pc1dot5-r18", Optional: true},
		{Name: "pc2-r18", Optional: true},
		{Name: "pc3-r18", Optional: true},
	},
}

const (
	UE_CapabilityRequestFilterCommon_Ext_LowerMSDRequest_r18_Pc1dot5_r18_True = 0
)

var uECapabilityRequestFilterCommonExtLowerMSDRequestR18Pc1dot5R18Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{UE_CapabilityRequestFilterCommon_Ext_LowerMSDRequest_r18_Pc1dot5_r18_True},
}

const (
	UE_CapabilityRequestFilterCommon_Ext_LowerMSDRequest_r18_Pc2_r18_True = 0
)

var uECapabilityRequestFilterCommonExtLowerMSDRequestR18Pc2R18Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{UE_CapabilityRequestFilterCommon_Ext_LowerMSDRequest_r18_Pc2_r18_True},
}

const (
	UE_CapabilityRequestFilterCommon_Ext_LowerMSDRequest_r18_Pc3_r18_True = 0
)

var uECapabilityRequestFilterCommonExtLowerMSDRequestR18Pc3R18Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{UE_CapabilityRequestFilterCommon_Ext_LowerMSDRequest_r18_Pc3_r18_True},
}

type UE_CapabilityRequestFilterCommon struct {
	Mrdc_Request *struct {
		OmitEN_DC    *int64
		IncludeNR_DC *int64
		IncludeNE_DC *int64
	}
	CodebookTypeRequest_r16 *struct {
		Type1_SinglePanel_r16   *int64
		Type1_MultiPanel_r16    *int64
		Type2_r16               *int64
		Type2_PortSelection_r16 *int64
	}
	UplinkTxSwitchRequest_r16    *int64
	RequestedCellGrouping_r16    []CellGrouping_r16
	FallbackGroupFiveRequest_r17 *int64
	LowerMSDRequest_r18          *struct {
		Pc1dot5_r18 *int64
		Pc2_r18     *int64
		Pc3_r18     *int64
	}
}

func (ie *UE_CapabilityRequestFilterCommon) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(uECapabilityRequestFilterCommonConstraints)

	// Extension-group presence (one var per [[...]] group).
	hasExtGroup0 := ie.CodebookTypeRequest_r16 != nil || ie.UplinkTxSwitchRequest_r16 != nil
	hasExtGroup1 := ie.RequestedCellGrouping_r16 != nil
	hasExtGroup2 := ie.FallbackGroupFiveRequest_r17 != nil
	hasExtGroup3 := ie.LowerMSDRequest_r18 != nil
	hasExtensions := hasExtGroup0 || hasExtGroup1 || hasExtGroup2 || hasExtGroup3

	// 1. SEQUENCE extension bit.
	if err := seq.EncodeExtensionBit(hasExtensions); err != nil {
		return err
	}

	// 2. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.Mrdc_Request != nil}); err != nil {
		return err
	}

	// 3. mrdc-Request: sequence
	{
		if ie.Mrdc_Request != nil {
			c := ie.Mrdc_Request
			uECapabilityRequestFilterCommonMrdcRequestSeq := e.NewSequenceEncoder(uECapabilityRequestFilterCommonMrdcRequestConstraints)
			if err := uECapabilityRequestFilterCommonMrdcRequestSeq.EncodePreamble([]bool{c.OmitEN_DC != nil, c.IncludeNR_DC != nil, c.IncludeNE_DC != nil}); err != nil {
				return err
			}
			if c.OmitEN_DC != nil {
				if err := e.EncodeEnumerated((*c.OmitEN_DC), uECapabilityRequestFilterCommonMrdcRequestOmitENDCConstraints); err != nil {
					return err
				}
			}
			if c.IncludeNR_DC != nil {
				if err := e.EncodeEnumerated((*c.IncludeNR_DC), uECapabilityRequestFilterCommonMrdcRequestIncludeNRDCConstraints); err != nil {
					return err
				}
			}
			if c.IncludeNE_DC != nil {
				if err := e.EncodeEnumerated((*c.IncludeNE_DC), uECapabilityRequestFilterCommonMrdcRequestIncludeNEDCConstraints); err != nil {
					return err
				}
			}
		}
	}

	if hasExtensions {
		extPresence := []bool{hasExtGroup0, hasExtGroup1, hasExtGroup2, hasExtGroup3}
		var extValues [][]byte

		// Extension group 0:
		if hasExtGroup0 {
			ex := per.NewEncoder(e.Variant())
			groupSeq := ex.NewSequenceEncoder(per.SequenceConstraints{
				Extensible: false,
				RootComponents: []per.ComponentInfo{
					{Name: "codebookTypeRequest-r16", Optional: true},
					{Name: "uplinkTxSwitchRequest-r16", Optional: true},
				},
			})
			if err := groupSeq.EncodePreamble([]bool{ie.CodebookTypeRequest_r16 != nil, ie.UplinkTxSwitchRequest_r16 != nil}); err != nil {
				return err
			}

			if ie.CodebookTypeRequest_r16 != nil {
				c := ie.CodebookTypeRequest_r16
				uECapabilityRequestFilterCommonExtCodebookTypeRequestR16Seq := ex.NewSequenceEncoder(uECapabilityRequestFilterCommonExtCodebookTypeRequestR16Constraints)
				if err := uECapabilityRequestFilterCommonExtCodebookTypeRequestR16Seq.EncodePreamble([]bool{c.Type1_SinglePanel_r16 != nil, c.Type1_MultiPanel_r16 != nil, c.Type2_r16 != nil, c.Type2_PortSelection_r16 != nil}); err != nil {
					return err
				}
				if c.Type1_SinglePanel_r16 != nil {
					if err := ex.EncodeEnumerated((*c.Type1_SinglePanel_r16), uECapabilityRequestFilterCommonExtCodebookTypeRequestR16Type1SinglePanelR16Constraints); err != nil {
						return err
					}
				}
				if c.Type1_MultiPanel_r16 != nil {
					if err := ex.EncodeEnumerated((*c.Type1_MultiPanel_r16), uECapabilityRequestFilterCommonExtCodebookTypeRequestR16Type1MultiPanelR16Constraints); err != nil {
						return err
					}
				}
				if c.Type2_r16 != nil {
					if err := ex.EncodeEnumerated((*c.Type2_r16), uECapabilityRequestFilterCommonExtCodebookTypeRequestR16Type2R16Constraints); err != nil {
						return err
					}
				}
				if c.Type2_PortSelection_r16 != nil {
					if err := ex.EncodeEnumerated((*c.Type2_PortSelection_r16), uECapabilityRequestFilterCommonExtCodebookTypeRequestR16Type2PortSelectionR16Constraints); err != nil {
						return err
					}
				}
			}

			if ie.UplinkTxSwitchRequest_r16 != nil {
				if err := ex.EncodeEnumerated(*ie.UplinkTxSwitchRequest_r16, uECapabilityRequestFilterCommonExtUplinkTxSwitchRequestR16Constraints); err != nil {
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
					{Name: "requestedCellGrouping-r16", Optional: true},
				},
			})
			if err := groupSeq.EncodePreamble([]bool{ie.RequestedCellGrouping_r16 != nil}); err != nil {
				return err
			}

			if ie.RequestedCellGrouping_r16 != nil {
				seqOf := ex.NewSequenceOfEncoder(uECapabilityRequestFilterCommonExtRequestedCellGroupingR16Constraints)
				if err := seqOf.EncodeLength(int64(len(ie.RequestedCellGrouping_r16))); err != nil {
					return err
				}
				for i := range ie.RequestedCellGrouping_r16 {
					if err := ie.RequestedCellGrouping_r16[i].Encode(ex); err != nil {
						return err
					}
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
					{Name: "fallbackGroupFiveRequest-r17", Optional: true},
				},
			})
			if err := groupSeq.EncodePreamble([]bool{ie.FallbackGroupFiveRequest_r17 != nil}); err != nil {
				return err
			}

			if ie.FallbackGroupFiveRequest_r17 != nil {
				if err := ex.EncodeEnumerated(*ie.FallbackGroupFiveRequest_r17, uECapabilityRequestFilterCommonExtFallbackGroupFiveRequestR17Constraints); err != nil {
					return err
				}
			}
			extValues = append(extValues, ex.Bytes())
		}

		// Extension group 3:
		if hasExtGroup3 {
			ex := per.NewEncoder(e.Variant())
			groupSeq := ex.NewSequenceEncoder(per.SequenceConstraints{
				Extensible: false,
				RootComponents: []per.ComponentInfo{
					{Name: "lowerMSDRequest-r18", Optional: true},
				},
			})
			if err := groupSeq.EncodePreamble([]bool{ie.LowerMSDRequest_r18 != nil}); err != nil {
				return err
			}

			if ie.LowerMSDRequest_r18 != nil {
				c := ie.LowerMSDRequest_r18
				uECapabilityRequestFilterCommonExtLowerMSDRequestR18Seq := ex.NewSequenceEncoder(uECapabilityRequestFilterCommonExtLowerMSDRequestR18Constraints)
				if err := uECapabilityRequestFilterCommonExtLowerMSDRequestR18Seq.EncodePreamble([]bool{c.Pc1dot5_r18 != nil, c.Pc2_r18 != nil, c.Pc3_r18 != nil}); err != nil {
					return err
				}
				if c.Pc1dot5_r18 != nil {
					if err := ex.EncodeEnumerated((*c.Pc1dot5_r18), uECapabilityRequestFilterCommonExtLowerMSDRequestR18Pc1dot5R18Constraints); err != nil {
						return err
					}
				}
				if c.Pc2_r18 != nil {
					if err := ex.EncodeEnumerated((*c.Pc2_r18), uECapabilityRequestFilterCommonExtLowerMSDRequestR18Pc2R18Constraints); err != nil {
						return err
					}
				}
				if c.Pc3_r18 != nil {
					if err := ex.EncodeEnumerated((*c.Pc3_r18), uECapabilityRequestFilterCommonExtLowerMSDRequestR18Pc3R18Constraints); err != nil {
						return err
					}
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

func (ie *UE_CapabilityRequestFilterCommon) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(uECapabilityRequestFilterCommonConstraints)

	// 1. SEQUENCE extension bit.
	if err := seq.DecodeExtensionBit(); err != nil {
		return err
	}

	// 2. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 3. mrdc-Request: sequence
	{
		if seq.IsComponentPresent(0) {
			ie.Mrdc_Request = &struct {
				OmitEN_DC    *int64
				IncludeNR_DC *int64
				IncludeNE_DC *int64
			}{}
			c := ie.Mrdc_Request
			uECapabilityRequestFilterCommonMrdcRequestSeq := d.NewSequenceDecoder(uECapabilityRequestFilterCommonMrdcRequestConstraints)
			if err := uECapabilityRequestFilterCommonMrdcRequestSeq.DecodePreamble(); err != nil {
				return err
			}
			if uECapabilityRequestFilterCommonMrdcRequestSeq.IsComponentPresent(0) {
				c.OmitEN_DC = new(int64)
				v, err := d.DecodeEnumerated(uECapabilityRequestFilterCommonMrdcRequestOmitENDCConstraints)
				if err != nil {
					return err
				}
				(*c.OmitEN_DC) = v
			}
			if uECapabilityRequestFilterCommonMrdcRequestSeq.IsComponentPresent(1) {
				c.IncludeNR_DC = new(int64)
				v, err := d.DecodeEnumerated(uECapabilityRequestFilterCommonMrdcRequestIncludeNRDCConstraints)
				if err != nil {
					return err
				}
				(*c.IncludeNR_DC) = v
			}
			if uECapabilityRequestFilterCommonMrdcRequestSeq.IsComponentPresent(2) {
				c.IncludeNE_DC = new(int64)
				v, err := d.DecodeEnumerated(uECapabilityRequestFilterCommonMrdcRequestIncludeNEDCConstraints)
				if err != nil {
					return err
				}
				(*c.IncludeNE_DC) = v
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
				{Name: "codebookTypeRequest-r16", Optional: true},
				{Name: "uplinkTxSwitchRequest-r16", Optional: true},
			},
		})
		if err := groupSeq.DecodePreamble(); err != nil {
			return err
		}

		if groupSeq.IsComponentPresent(0) {
			ie.CodebookTypeRequest_r16 = &struct {
				Type1_SinglePanel_r16   *int64
				Type1_MultiPanel_r16    *int64
				Type2_r16               *int64
				Type2_PortSelection_r16 *int64
			}{}
			c := ie.CodebookTypeRequest_r16
			uECapabilityRequestFilterCommonExtCodebookTypeRequestR16Seq := dx.NewSequenceDecoder(uECapabilityRequestFilterCommonExtCodebookTypeRequestR16Constraints)
			if err := uECapabilityRequestFilterCommonExtCodebookTypeRequestR16Seq.DecodePreamble(); err != nil {
				return err
			}
			if uECapabilityRequestFilterCommonExtCodebookTypeRequestR16Seq.IsComponentPresent(0) {
				c.Type1_SinglePanel_r16 = new(int64)
				v, err := dx.DecodeEnumerated(uECapabilityRequestFilterCommonExtCodebookTypeRequestR16Type1SinglePanelR16Constraints)
				if err != nil {
					return err
				}
				(*c.Type1_SinglePanel_r16) = v
			}
			if uECapabilityRequestFilterCommonExtCodebookTypeRequestR16Seq.IsComponentPresent(1) {
				c.Type1_MultiPanel_r16 = new(int64)
				v, err := dx.DecodeEnumerated(uECapabilityRequestFilterCommonExtCodebookTypeRequestR16Type1MultiPanelR16Constraints)
				if err != nil {
					return err
				}
				(*c.Type1_MultiPanel_r16) = v
			}
			if uECapabilityRequestFilterCommonExtCodebookTypeRequestR16Seq.IsComponentPresent(2) {
				c.Type2_r16 = new(int64)
				v, err := dx.DecodeEnumerated(uECapabilityRequestFilterCommonExtCodebookTypeRequestR16Type2R16Constraints)
				if err != nil {
					return err
				}
				(*c.Type2_r16) = v
			}
			if uECapabilityRequestFilterCommonExtCodebookTypeRequestR16Seq.IsComponentPresent(3) {
				c.Type2_PortSelection_r16 = new(int64)
				v, err := dx.DecodeEnumerated(uECapabilityRequestFilterCommonExtCodebookTypeRequestR16Type2PortSelectionR16Constraints)
				if err != nil {
					return err
				}
				(*c.Type2_PortSelection_r16) = v
			}
		}

		if groupSeq.IsComponentPresent(1) {
			v, err := dx.DecodeEnumerated(uECapabilityRequestFilterCommonExtUplinkTxSwitchRequestR16Constraints)
			if err != nil {
				return err
			}
			ie.UplinkTxSwitchRequest_r16 = &v
		}
	}

	// Extension group 1:
	if seq.IsExtensionPresent(1) {
		dx := per.NewDecoder(extValues[extValueIndex], d.Variant())
		extValueIndex++
		groupSeq := dx.NewSequenceDecoder(per.SequenceConstraints{
			Extensible: false,
			RootComponents: []per.ComponentInfo{
				{Name: "requestedCellGrouping-r16", Optional: true},
			},
		})
		if err := groupSeq.DecodePreamble(); err != nil {
			return err
		}

		if groupSeq.IsComponentPresent(0) {
			seqOf := dx.NewSequenceOfDecoder(uECapabilityRequestFilterCommonExtRequestedCellGroupingR16Constraints)
			n, err := seqOf.DecodeLength()
			if err != nil {
				return err
			}
			ie.RequestedCellGrouping_r16 = make([]CellGrouping_r16, n)
			for i := int64(0); i < n; i++ {
				if err := ie.RequestedCellGrouping_r16[i].Decode(dx); err != nil {
					return err
				}
			}
		}
	}

	// Extension group 2:
	if seq.IsExtensionPresent(2) {
		dx := per.NewDecoder(extValues[extValueIndex], d.Variant())
		extValueIndex++
		groupSeq := dx.NewSequenceDecoder(per.SequenceConstraints{
			Extensible: false,
			RootComponents: []per.ComponentInfo{
				{Name: "fallbackGroupFiveRequest-r17", Optional: true},
			},
		})
		if err := groupSeq.DecodePreamble(); err != nil {
			return err
		}

		if groupSeq.IsComponentPresent(0) {
			v, err := dx.DecodeEnumerated(uECapabilityRequestFilterCommonExtFallbackGroupFiveRequestR17Constraints)
			if err != nil {
				return err
			}
			ie.FallbackGroupFiveRequest_r17 = &v
		}
	}

	// Extension group 3:
	if seq.IsExtensionPresent(3) {
		dx := per.NewDecoder(extValues[extValueIndex], d.Variant())
		extValueIndex++
		groupSeq := dx.NewSequenceDecoder(per.SequenceConstraints{
			Extensible: false,
			RootComponents: []per.ComponentInfo{
				{Name: "lowerMSDRequest-r18", Optional: true},
			},
		})
		if err := groupSeq.DecodePreamble(); err != nil {
			return err
		}

		if groupSeq.IsComponentPresent(0) {
			ie.LowerMSDRequest_r18 = &struct {
				Pc1dot5_r18 *int64
				Pc2_r18     *int64
				Pc3_r18     *int64
			}{}
			c := ie.LowerMSDRequest_r18
			uECapabilityRequestFilterCommonExtLowerMSDRequestR18Seq := dx.NewSequenceDecoder(uECapabilityRequestFilterCommonExtLowerMSDRequestR18Constraints)
			if err := uECapabilityRequestFilterCommonExtLowerMSDRequestR18Seq.DecodePreamble(); err != nil {
				return err
			}
			if uECapabilityRequestFilterCommonExtLowerMSDRequestR18Seq.IsComponentPresent(0) {
				c.Pc1dot5_r18 = new(int64)
				v, err := dx.DecodeEnumerated(uECapabilityRequestFilterCommonExtLowerMSDRequestR18Pc1dot5R18Constraints)
				if err != nil {
					return err
				}
				(*c.Pc1dot5_r18) = v
			}
			if uECapabilityRequestFilterCommonExtLowerMSDRequestR18Seq.IsComponentPresent(1) {
				c.Pc2_r18 = new(int64)
				v, err := dx.DecodeEnumerated(uECapabilityRequestFilterCommonExtLowerMSDRequestR18Pc2R18Constraints)
				if err != nil {
					return err
				}
				(*c.Pc2_r18) = v
			}
			if uECapabilityRequestFilterCommonExtLowerMSDRequestR18Seq.IsComponentPresent(2) {
				c.Pc3_r18 = new(int64)
				v, err := dx.DecodeEnumerated(uECapabilityRequestFilterCommonExtLowerMSDRequestR18Pc3R18Constraints)
				if err != nil {
					return err
				}
				(*c.Pc3_r18) = v
			}
		}
	}

	return nil
}
