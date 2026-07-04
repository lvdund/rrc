// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: MRB-PDCP-ConfigBroadcast-r17 (line 28530).

var mRBPDCPConfigBroadcastR17Constraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "pdcp-SN-SizeDL-r17", Optional: true},
		{Name: "headerCompression-r17"},
		{Name: "t-Reordering-r17", Optional: true},
	},
}

const (
	MRB_PDCP_ConfigBroadcast_r17_Pdcp_SN_SizeDL_r17_Len12bits = 0
)

var mRBPDCPConfigBroadcastR17PdcpSNSizeDLR17Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{MRB_PDCP_ConfigBroadcast_r17_Pdcp_SN_SizeDL_r17_Len12bits},
}

var mRB_PDCP_ConfigBroadcast_r17HeaderCompressionR17Constraints = per.ChoiceConstraints{
	Extensible: false,
	RootAlternatives: []per.AlternativeInfo{
		{Name: "notUsed"},
		{Name: "rohc"},
	},
}

const (
	MRB_PDCP_ConfigBroadcast_r17_HeaderCompression_r17_NotUsed = 0
	MRB_PDCP_ConfigBroadcast_r17_HeaderCompression_r17_Rohc    = 1
)

const (
	MRB_PDCP_ConfigBroadcast_r17_T_Reordering_r17_Ms1    = 0
	MRB_PDCP_ConfigBroadcast_r17_T_Reordering_r17_Ms10   = 1
	MRB_PDCP_ConfigBroadcast_r17_T_Reordering_r17_Ms40   = 2
	MRB_PDCP_ConfigBroadcast_r17_T_Reordering_r17_Ms160  = 3
	MRB_PDCP_ConfigBroadcast_r17_T_Reordering_r17_Ms500  = 4
	MRB_PDCP_ConfigBroadcast_r17_T_Reordering_r17_Ms1000 = 5
	MRB_PDCP_ConfigBroadcast_r17_T_Reordering_r17_Ms1250 = 6
	MRB_PDCP_ConfigBroadcast_r17_T_Reordering_r17_Ms2750 = 7
)

var mRBPDCPConfigBroadcastR17TReorderingR17Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{MRB_PDCP_ConfigBroadcast_r17_T_Reordering_r17_Ms1, MRB_PDCP_ConfigBroadcast_r17_T_Reordering_r17_Ms10, MRB_PDCP_ConfigBroadcast_r17_T_Reordering_r17_Ms40, MRB_PDCP_ConfigBroadcast_r17_T_Reordering_r17_Ms160, MRB_PDCP_ConfigBroadcast_r17_T_Reordering_r17_Ms500, MRB_PDCP_ConfigBroadcast_r17_T_Reordering_r17_Ms1000, MRB_PDCP_ConfigBroadcast_r17_T_Reordering_r17_Ms1250, MRB_PDCP_ConfigBroadcast_r17_T_Reordering_r17_Ms2750},
}

type MRB_PDCP_ConfigBroadcast_r17 struct {
	Pdcp_SN_SizeDL_r17    *int64
	HeaderCompression_r17 struct {
		Choice  int
		NotUsed *struct{}
		Rohc    *struct {
			MaxCID_r17   int64
			Profiles_r17 struct {
				Profile0x0000_r17 bool
				Profile0x0001_r17 bool
				Profile0x0002_r17 bool
			}
		}
	}
	T_Reordering_r17 *int64
}

func (ie *MRB_PDCP_ConfigBroadcast_r17) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(mRBPDCPConfigBroadcastR17Constraints)

	// 1. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.Pdcp_SN_SizeDL_r17 != nil, ie.T_Reordering_r17 != nil}); err != nil {
		return err
	}

	// 2. pdcp-SN-SizeDL-r17: enumerated
	{
		if ie.Pdcp_SN_SizeDL_r17 != nil {
			if err := e.EncodeEnumerated(*ie.Pdcp_SN_SizeDL_r17, mRBPDCPConfigBroadcastR17PdcpSNSizeDLR17Constraints); err != nil {
				return err
			}
		}
	}

	// 3. headerCompression-r17: choice
	{
		choiceEnc := e.NewChoiceEncoder(mRB_PDCP_ConfigBroadcast_r17HeaderCompressionR17Constraints)
		if err := choiceEnc.EncodeChoice(int64(ie.HeaderCompression_r17.Choice), false, nil); err != nil {
			return err
		}
		switch ie.HeaderCompression_r17.Choice {
		case MRB_PDCP_ConfigBroadcast_r17_HeaderCompression_r17_NotUsed:
			if err := e.EncodeNull(); err != nil {
				return err
			}
		case MRB_PDCP_ConfigBroadcast_r17_HeaderCompression_r17_Rohc:
			c := (*ie.HeaderCompression_r17.Rohc)
			if err := e.EncodeInteger(c.MaxCID_r17, per.Constrained(1, 16)); err != nil {
				return err
			}
			{
				c := &c.Profiles_r17
				if err := e.EncodeBoolean(c.Profile0x0000_r17); err != nil {
					return err
				}
				if err := e.EncodeBoolean(c.Profile0x0001_r17); err != nil {
					return err
				}
				if err := e.EncodeBoolean(c.Profile0x0002_r17); err != nil {
					return err
				}
			}
		default:
			return &per.ConstraintViolationError{TypeName: "choice", Value: int64(ie.HeaderCompression_r17.Choice), Constraint: "index out of range"}
		}
	}

	// 4. t-Reordering-r17: enumerated
	{
		if ie.T_Reordering_r17 != nil {
			if err := e.EncodeEnumerated(*ie.T_Reordering_r17, mRBPDCPConfigBroadcastR17TReorderingR17Constraints); err != nil {
				return err
			}
		}
	}

	return nil
}

func (ie *MRB_PDCP_ConfigBroadcast_r17) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(mRBPDCPConfigBroadcastR17Constraints)

	// 1. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 2. pdcp-SN-SizeDL-r17: enumerated
	{
		if seq.IsComponentPresent(0) {
			idx, err := d.DecodeEnumerated(mRBPDCPConfigBroadcastR17PdcpSNSizeDLR17Constraints)
			if err != nil {
				return err
			}
			ie.Pdcp_SN_SizeDL_r17 = &idx
		}
	}

	// 3. headerCompression-r17: choice
	{
		choiceDec := d.NewChoiceDecoder(mRB_PDCP_ConfigBroadcast_r17HeaderCompressionR17Constraints)
		index, isExt, _, err := choiceDec.DecodeChoice()
		if err != nil {
			return err
		}
		ie.HeaderCompression_r17.Choice = int(index)
		if isExt {
			return &per.DecodeError{TypeName: "choice", Reason: "extension alternative not supported"}
		}
		switch index {
		case MRB_PDCP_ConfigBroadcast_r17_HeaderCompression_r17_NotUsed:
			ie.HeaderCompression_r17.NotUsed = &struct{}{}
			if err := d.DecodeNull(); err != nil {
				return err
			}
		case MRB_PDCP_ConfigBroadcast_r17_HeaderCompression_r17_Rohc:
			ie.HeaderCompression_r17.Rohc = &struct {
				MaxCID_r17   int64
				Profiles_r17 struct {
					Profile0x0000_r17 bool
					Profile0x0001_r17 bool
					Profile0x0002_r17 bool
				}
			}{}
			c := (*ie.HeaderCompression_r17.Rohc)
			{
				v, err := d.DecodeInteger(per.Constrained(1, 16))
				if err != nil {
					return err
				}
				c.MaxCID_r17 = v
			}
			{
				c := &c.Profiles_r17
				{
					v, err := d.DecodeBoolean()
					if err != nil {
						return err
					}
					c.Profile0x0000_r17 = v
				}
				{
					v, err := d.DecodeBoolean()
					if err != nil {
						return err
					}
					c.Profile0x0001_r17 = v
				}
				{
					v, err := d.DecodeBoolean()
					if err != nil {
						return err
					}
					c.Profile0x0002_r17 = v
				}
			}
		}
	}

	// 4. t-Reordering-r17: enumerated
	{
		if seq.IsComponentPresent(2) {
			idx, err := d.DecodeEnumerated(mRBPDCPConfigBroadcastR17TReorderingR17Constraints)
			if err != nil {
				return err
			}
			ie.T_Reordering_r17 = &idx
		}
	}

	return nil
}
