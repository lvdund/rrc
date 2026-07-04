// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: PDSCH-ConfigPTM-r17 (line 28635).

var pDSCHConfigPTMR17Constraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "dataScramblingIdentityPDSCH-r17", Optional: true},
		{Name: "dmrs-ScramblingID0-r17", Optional: true},
		{Name: "pdsch-AggregationFactor-r17", Optional: true},
	},
}

var pDSCHConfigPTMR17DataScramblingIdentityPDSCHR17Constraints = per.Constrained(0, 1023)

var pDSCHConfigPTMR17DmrsScramblingID0R17Constraints = per.Constrained(0, 65535)

const (
	PDSCH_ConfigPTM_r17_Pdsch_AggregationFactor_r17_N2 = 0
	PDSCH_ConfigPTM_r17_Pdsch_AggregationFactor_r17_N4 = 1
	PDSCH_ConfigPTM_r17_Pdsch_AggregationFactor_r17_N8 = 2
)

var pDSCHConfigPTMR17PdschAggregationFactorR17Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{PDSCH_ConfigPTM_r17_Pdsch_AggregationFactor_r17_N2, PDSCH_ConfigPTM_r17_Pdsch_AggregationFactor_r17_N4, PDSCH_ConfigPTM_r17_Pdsch_AggregationFactor_r17_N8},
}

type PDSCH_ConfigPTM_r17 struct {
	DataScramblingIdentityPDSCH_r17 *int64
	Dmrs_ScramblingID0_r17          *int64
	Pdsch_AggregationFactor_r17     *int64
}

func (ie *PDSCH_ConfigPTM_r17) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(pDSCHConfigPTMR17Constraints)

	// 1. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.DataScramblingIdentityPDSCH_r17 != nil, ie.Dmrs_ScramblingID0_r17 != nil, ie.Pdsch_AggregationFactor_r17 != nil}); err != nil {
		return err
	}

	// 2. dataScramblingIdentityPDSCH-r17: integer
	{
		if ie.DataScramblingIdentityPDSCH_r17 != nil {
			if err := e.EncodeInteger(*ie.DataScramblingIdentityPDSCH_r17, pDSCHConfigPTMR17DataScramblingIdentityPDSCHR17Constraints); err != nil {
				return err
			}
		}
	}

	// 3. dmrs-ScramblingID0-r17: integer
	{
		if ie.Dmrs_ScramblingID0_r17 != nil {
			if err := e.EncodeInteger(*ie.Dmrs_ScramblingID0_r17, pDSCHConfigPTMR17DmrsScramblingID0R17Constraints); err != nil {
				return err
			}
		}
	}

	// 4. pdsch-AggregationFactor-r17: enumerated
	{
		if ie.Pdsch_AggregationFactor_r17 != nil {
			if err := e.EncodeEnumerated(*ie.Pdsch_AggregationFactor_r17, pDSCHConfigPTMR17PdschAggregationFactorR17Constraints); err != nil {
				return err
			}
		}
	}

	return nil
}

func (ie *PDSCH_ConfigPTM_r17) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(pDSCHConfigPTMR17Constraints)

	// 1. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 2. dataScramblingIdentityPDSCH-r17: integer
	{
		if seq.IsComponentPresent(0) {
			v, err := d.DecodeInteger(pDSCHConfigPTMR17DataScramblingIdentityPDSCHR17Constraints)
			if err != nil {
				return err
			}
			ie.DataScramblingIdentityPDSCH_r17 = &v
		}
	}

	// 3. dmrs-ScramblingID0-r17: integer
	{
		if seq.IsComponentPresent(1) {
			v, err := d.DecodeInteger(pDSCHConfigPTMR17DmrsScramblingID0R17Constraints)
			if err != nil {
				return err
			}
			ie.Dmrs_ScramblingID0_r17 = &v
		}
	}

	// 4. pdsch-AggregationFactor-r17: enumerated
	{
		if seq.IsComponentPresent(2) {
			idx, err := d.DecodeEnumerated(pDSCHConfigPTMR17PdschAggregationFactorR17Constraints)
			if err != nil {
				return err
			}
			ie.Pdsch_AggregationFactor_r17 = &idx
		}
	}

	return nil
}
