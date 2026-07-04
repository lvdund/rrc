// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: FeatureSetUplinkPerCC-v1900 (line 20695).

var featureSetUplinkPerCCV1900Constraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "supportedBandwidthUL-v1900", Optional: true},
		{Name: "supportedMinBandwidthUL-v1900", Optional: true},
		{Name: "support32-UL-HARQ-ProcessTN-r19", Optional: true},
		{Name: "nonCodebook-3TxPUSCH-SingleTRP-r19", Optional: true},
		{Name: "codebook-3TxPUSCH-SingleTRP-r19", Optional: true},
		{Name: "codebook-3PortPUSCH-TypeB-r19", Optional: true},
		{Name: "mTRP-PUSCH-RepetitionTypeB-3Port-r19", Optional: true},
	},
}

const (
	FeatureSetUplinkPerCC_v1900_Support32_UL_HARQ_ProcessTN_r19_Supported = 0
)

var featureSetUplinkPerCCV1900Support32ULHARQProcessTNR19Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{FeatureSetUplinkPerCC_v1900_Support32_UL_HARQ_ProcessTN_r19_Supported},
}

var featureSetUplinkPerCCV1900Codebook3PortPUSCHTypeBR19Constraints = per.Constrained(1, 2)

var featureSetUplinkPerCCV1900MTRPPUSCHRepetitionTypeB3PortR19Constraints = per.Constrained(1, 3)

type FeatureSetUplinkPerCC_v1900 struct {
	SupportedBandwidthUL_v1900         *SupportedBandwidth_v1900
	SupportedMinBandwidthUL_v1900      *SupportedBandwidth_v1900
	Support32_UL_HARQ_ProcessTN_r19    *int64
	NonCodebook_3TxPUSCH_SingleTRP_r19 *struct {
		MaxNumberLayer_r19           int64
		MaxNumberSRS_Resource_r19    int64
		MaxNumberSimultaneousSRS_r19 int64
	}
	Codebook_3TxPUSCH_SingleTRP_r19 *struct {
		MaxNumberPUSCH_MIMO_Layer_r19 int64
		MaxNumberSRS_Resource_r19     int64
	}
	Codebook_3PortPUSCH_TypeB_r19        *int64
	MTRP_PUSCH_RepetitionTypeB_3Port_r19 *int64
}

func (ie *FeatureSetUplinkPerCC_v1900) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(featureSetUplinkPerCCV1900Constraints)

	// 1. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.SupportedBandwidthUL_v1900 != nil, ie.SupportedMinBandwidthUL_v1900 != nil, ie.Support32_UL_HARQ_ProcessTN_r19 != nil, ie.NonCodebook_3TxPUSCH_SingleTRP_r19 != nil, ie.Codebook_3TxPUSCH_SingleTRP_r19 != nil, ie.Codebook_3PortPUSCH_TypeB_r19 != nil, ie.MTRP_PUSCH_RepetitionTypeB_3Port_r19 != nil}); err != nil {
		return err
	}

	// 2. supportedBandwidthUL-v1900: ref
	{
		if ie.SupportedBandwidthUL_v1900 != nil {
			if err := ie.SupportedBandwidthUL_v1900.Encode(e); err != nil {
				return err
			}
		}
	}

	// 3. supportedMinBandwidthUL-v1900: ref
	{
		if ie.SupportedMinBandwidthUL_v1900 != nil {
			if err := ie.SupportedMinBandwidthUL_v1900.Encode(e); err != nil {
				return err
			}
		}
	}

	// 4. support32-UL-HARQ-ProcessTN-r19: enumerated
	{
		if ie.Support32_UL_HARQ_ProcessTN_r19 != nil {
			if err := e.EncodeEnumerated(*ie.Support32_UL_HARQ_ProcessTN_r19, featureSetUplinkPerCCV1900Support32ULHARQProcessTNR19Constraints); err != nil {
				return err
			}
		}
	}

	// 5. nonCodebook-3TxPUSCH-SingleTRP-r19: sequence
	{
		if ie.NonCodebook_3TxPUSCH_SingleTRP_r19 != nil {
			c := ie.NonCodebook_3TxPUSCH_SingleTRP_r19
			if err := e.EncodeInteger(c.MaxNumberLayer_r19, per.Constrained(1, 3)); err != nil {
				return err
			}
			if err := e.EncodeInteger(c.MaxNumberSRS_Resource_r19, per.Constrained(1, 3)); err != nil {
				return err
			}
			if err := e.EncodeInteger(c.MaxNumberSimultaneousSRS_r19, per.Constrained(1, 3)); err != nil {
				return err
			}
		}
	}

	// 6. codebook-3TxPUSCH-SingleTRP-r19: sequence
	{
		if ie.Codebook_3TxPUSCH_SingleTRP_r19 != nil {
			c := ie.Codebook_3TxPUSCH_SingleTRP_r19
			if err := e.EncodeInteger(c.MaxNumberPUSCH_MIMO_Layer_r19, per.Constrained(1, 3)); err != nil {
				return err
			}
			if err := e.EncodeInteger(c.MaxNumberSRS_Resource_r19, per.Constrained(1, 2)); err != nil {
				return err
			}
		}
	}

	// 7. codebook-3PortPUSCH-TypeB-r19: integer
	{
		if ie.Codebook_3PortPUSCH_TypeB_r19 != nil {
			if err := e.EncodeInteger(*ie.Codebook_3PortPUSCH_TypeB_r19, featureSetUplinkPerCCV1900Codebook3PortPUSCHTypeBR19Constraints); err != nil {
				return err
			}
		}
	}

	// 8. mTRP-PUSCH-RepetitionTypeB-3Port-r19: integer
	{
		if ie.MTRP_PUSCH_RepetitionTypeB_3Port_r19 != nil {
			if err := e.EncodeInteger(*ie.MTRP_PUSCH_RepetitionTypeB_3Port_r19, featureSetUplinkPerCCV1900MTRPPUSCHRepetitionTypeB3PortR19Constraints); err != nil {
				return err
			}
		}
	}

	return nil
}

func (ie *FeatureSetUplinkPerCC_v1900) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(featureSetUplinkPerCCV1900Constraints)

	// 1. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 2. supportedBandwidthUL-v1900: ref
	{
		if seq.IsComponentPresent(0) {
			ie.SupportedBandwidthUL_v1900 = new(SupportedBandwidth_v1900)
			if err := ie.SupportedBandwidthUL_v1900.Decode(d); err != nil {
				return err
			}
		}
	}

	// 3. supportedMinBandwidthUL-v1900: ref
	{
		if seq.IsComponentPresent(1) {
			ie.SupportedMinBandwidthUL_v1900 = new(SupportedBandwidth_v1900)
			if err := ie.SupportedMinBandwidthUL_v1900.Decode(d); err != nil {
				return err
			}
		}
	}

	// 4. support32-UL-HARQ-ProcessTN-r19: enumerated
	{
		if seq.IsComponentPresent(2) {
			idx, err := d.DecodeEnumerated(featureSetUplinkPerCCV1900Support32ULHARQProcessTNR19Constraints)
			if err != nil {
				return err
			}
			ie.Support32_UL_HARQ_ProcessTN_r19 = &idx
		}
	}

	// 5. nonCodebook-3TxPUSCH-SingleTRP-r19: sequence
	{
		if seq.IsComponentPresent(3) {
			ie.NonCodebook_3TxPUSCH_SingleTRP_r19 = &struct {
				MaxNumberLayer_r19           int64
				MaxNumberSRS_Resource_r19    int64
				MaxNumberSimultaneousSRS_r19 int64
			}{}
			c := ie.NonCodebook_3TxPUSCH_SingleTRP_r19
			{
				v, err := d.DecodeInteger(per.Constrained(1, 3))
				if err != nil {
					return err
				}
				c.MaxNumberLayer_r19 = v
			}
			{
				v, err := d.DecodeInteger(per.Constrained(1, 3))
				if err != nil {
					return err
				}
				c.MaxNumberSRS_Resource_r19 = v
			}
			{
				v, err := d.DecodeInteger(per.Constrained(1, 3))
				if err != nil {
					return err
				}
				c.MaxNumberSimultaneousSRS_r19 = v
			}
		}
	}

	// 6. codebook-3TxPUSCH-SingleTRP-r19: sequence
	{
		if seq.IsComponentPresent(4) {
			ie.Codebook_3TxPUSCH_SingleTRP_r19 = &struct {
				MaxNumberPUSCH_MIMO_Layer_r19 int64
				MaxNumberSRS_Resource_r19     int64
			}{}
			c := ie.Codebook_3TxPUSCH_SingleTRP_r19
			{
				v, err := d.DecodeInteger(per.Constrained(1, 3))
				if err != nil {
					return err
				}
				c.MaxNumberPUSCH_MIMO_Layer_r19 = v
			}
			{
				v, err := d.DecodeInteger(per.Constrained(1, 2))
				if err != nil {
					return err
				}
				c.MaxNumberSRS_Resource_r19 = v
			}
		}
	}

	// 7. codebook-3PortPUSCH-TypeB-r19: integer
	{
		if seq.IsComponentPresent(5) {
			v, err := d.DecodeInteger(featureSetUplinkPerCCV1900Codebook3PortPUSCHTypeBR19Constraints)
			if err != nil {
				return err
			}
			ie.Codebook_3PortPUSCH_TypeB_r19 = &v
		}
	}

	// 8. mTRP-PUSCH-RepetitionTypeB-3Port-r19: integer
	{
		if seq.IsComponentPresent(6) {
			v, err := d.DecodeInteger(featureSetUplinkPerCCV1900MTRPPUSCHRepetitionTypeB3PortR19Constraints)
			if err != nil {
				return err
			}
			ie.MTRP_PUSCH_RepetitionTypeB_3Port_r19 = &v
		}
	}

	return nil
}
