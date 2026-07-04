// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: MRB-PDCP-ConfigMulticast-r18 (line 28584).

var mRBPDCPConfigMulticastR18Constraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "pdcp-SN-SizeDL-r18"},
		{Name: "headerCompression-r18"},
		{Name: "t-Reordering-r17", Optional: true},
	},
}

const (
	MRB_PDCP_ConfigMulticast_r18_Pdcp_SN_SizeDL_r18_Len12bits = 0
	MRB_PDCP_ConfigMulticast_r18_Pdcp_SN_SizeDL_r18_Len18bits = 1
)

var mRBPDCPConfigMulticastR18PdcpSNSizeDLR18Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{MRB_PDCP_ConfigMulticast_r18_Pdcp_SN_SizeDL_r18_Len12bits, MRB_PDCP_ConfigMulticast_r18_Pdcp_SN_SizeDL_r18_Len18bits},
}

var mRB_PDCP_ConfigMulticast_r18HeaderCompressionR18Constraints = per.ChoiceConstraints{
	Extensible: false,
	RootAlternatives: []per.AlternativeInfo{
		{Name: "notUsed"},
		{Name: "rohc"},
	},
}

const (
	MRB_PDCP_ConfigMulticast_r18_HeaderCompression_r18_NotUsed = 0
	MRB_PDCP_ConfigMulticast_r18_HeaderCompression_r18_Rohc    = 1
)

const (
	MRB_PDCP_ConfigMulticast_r18_T_Reordering_r17_Ms1    = 0
	MRB_PDCP_ConfigMulticast_r18_T_Reordering_r17_Ms10   = 1
	MRB_PDCP_ConfigMulticast_r18_T_Reordering_r17_Ms40   = 2
	MRB_PDCP_ConfigMulticast_r18_T_Reordering_r17_Ms160  = 3
	MRB_PDCP_ConfigMulticast_r18_T_Reordering_r17_Ms500  = 4
	MRB_PDCP_ConfigMulticast_r18_T_Reordering_r17_Ms1000 = 5
	MRB_PDCP_ConfigMulticast_r18_T_Reordering_r17_Ms1250 = 6
	MRB_PDCP_ConfigMulticast_r18_T_Reordering_r17_Ms2750 = 7
)

var mRBPDCPConfigMulticastR18TReorderingR17Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{MRB_PDCP_ConfigMulticast_r18_T_Reordering_r17_Ms1, MRB_PDCP_ConfigMulticast_r18_T_Reordering_r17_Ms10, MRB_PDCP_ConfigMulticast_r18_T_Reordering_r17_Ms40, MRB_PDCP_ConfigMulticast_r18_T_Reordering_r17_Ms160, MRB_PDCP_ConfigMulticast_r18_T_Reordering_r17_Ms500, MRB_PDCP_ConfigMulticast_r18_T_Reordering_r17_Ms1000, MRB_PDCP_ConfigMulticast_r18_T_Reordering_r17_Ms1250, MRB_PDCP_ConfigMulticast_r18_T_Reordering_r17_Ms2750},
}

type MRB_PDCP_ConfigMulticast_r18 struct {
	Pdcp_SN_SizeDL_r18    int64
	HeaderCompression_r18 struct {
		Choice  int
		NotUsed *struct{}
		Rohc    *struct {
			MaxCID_r18   int64
			Profiles_r18 struct {
				Profile0x0000_r18 bool
				Profile0x0001_r18 bool
				Profile0x0002_r18 bool
			}
		}
	}
	T_Reordering_r17 *int64
}

func (ie *MRB_PDCP_ConfigMulticast_r18) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(mRBPDCPConfigMulticastR18Constraints)

	// 1. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.T_Reordering_r17 != nil}); err != nil {
		return err
	}

	// 2. pdcp-SN-SizeDL-r18: enumerated
	{
		if err := e.EncodeEnumerated(ie.Pdcp_SN_SizeDL_r18, mRBPDCPConfigMulticastR18PdcpSNSizeDLR18Constraints); err != nil {
			return err
		}
	}

	// 3. headerCompression-r18: choice
	{
		choiceEnc := e.NewChoiceEncoder(mRB_PDCP_ConfigMulticast_r18HeaderCompressionR18Constraints)
		if err := choiceEnc.EncodeChoice(int64(ie.HeaderCompression_r18.Choice), false, nil); err != nil {
			return err
		}
		switch ie.HeaderCompression_r18.Choice {
		case MRB_PDCP_ConfigMulticast_r18_HeaderCompression_r18_NotUsed:
			if err := e.EncodeNull(); err != nil {
				return err
			}
		case MRB_PDCP_ConfigMulticast_r18_HeaderCompression_r18_Rohc:
			c := (*ie.HeaderCompression_r18.Rohc)
			if err := e.EncodeInteger(c.MaxCID_r18, per.Constrained(1, 16)); err != nil {
				return err
			}
			{
				c := &c.Profiles_r18
				if err := e.EncodeBoolean(c.Profile0x0000_r18); err != nil {
					return err
				}
				if err := e.EncodeBoolean(c.Profile0x0001_r18); err != nil {
					return err
				}
				if err := e.EncodeBoolean(c.Profile0x0002_r18); err != nil {
					return err
				}
			}
		default:
			return &per.ConstraintViolationError{TypeName: "choice", Value: int64(ie.HeaderCompression_r18.Choice), Constraint: "index out of range"}
		}
	}

	// 4. t-Reordering-r17: enumerated
	{
		if ie.T_Reordering_r17 != nil {
			if err := e.EncodeEnumerated(*ie.T_Reordering_r17, mRBPDCPConfigMulticastR18TReorderingR17Constraints); err != nil {
				return err
			}
		}
	}

	return nil
}

func (ie *MRB_PDCP_ConfigMulticast_r18) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(mRBPDCPConfigMulticastR18Constraints)

	// 1. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 2. pdcp-SN-SizeDL-r18: enumerated
	{
		v0, err := d.DecodeEnumerated(mRBPDCPConfigMulticastR18PdcpSNSizeDLR18Constraints)
		if err != nil {
			return err
		}
		ie.Pdcp_SN_SizeDL_r18 = v0
	}

	// 3. headerCompression-r18: choice
	{
		choiceDec := d.NewChoiceDecoder(mRB_PDCP_ConfigMulticast_r18HeaderCompressionR18Constraints)
		index, isExt, _, err := choiceDec.DecodeChoice()
		if err != nil {
			return err
		}
		ie.HeaderCompression_r18.Choice = int(index)
		if isExt {
			return &per.DecodeError{TypeName: "choice", Reason: "extension alternative not supported"}
		}
		switch index {
		case MRB_PDCP_ConfigMulticast_r18_HeaderCompression_r18_NotUsed:
			ie.HeaderCompression_r18.NotUsed = &struct{}{}
			if err := d.DecodeNull(); err != nil {
				return err
			}
		case MRB_PDCP_ConfigMulticast_r18_HeaderCompression_r18_Rohc:
			ie.HeaderCompression_r18.Rohc = &struct {
				MaxCID_r18   int64
				Profiles_r18 struct {
					Profile0x0000_r18 bool
					Profile0x0001_r18 bool
					Profile0x0002_r18 bool
				}
			}{}
			c := (*ie.HeaderCompression_r18.Rohc)
			{
				v, err := d.DecodeInteger(per.Constrained(1, 16))
				if err != nil {
					return err
				}
				c.MaxCID_r18 = v
			}
			{
				c := &c.Profiles_r18
				{
					v, err := d.DecodeBoolean()
					if err != nil {
						return err
					}
					c.Profile0x0000_r18 = v
				}
				{
					v, err := d.DecodeBoolean()
					if err != nil {
						return err
					}
					c.Profile0x0001_r18 = v
				}
				{
					v, err := d.DecodeBoolean()
					if err != nil {
						return err
					}
					c.Profile0x0002_r18 = v
				}
			}
		}
	}

	// 4. t-Reordering-r17: enumerated
	{
		if seq.IsComponentPresent(2) {
			idx, err := d.DecodeEnumerated(mRBPDCPConfigMulticastR18TReorderingR17Constraints)
			if err != nil {
				return err
			}
			ie.T_Reordering_r17 = &idx
		}
	}

	return nil
}
