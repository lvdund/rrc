// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: FeatureSetUplink-v1850 (line 20395).

var featureSetUplinkV1850Constraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "pusch-DMRS8Tx-r18", Optional: true},
		{Name: "additionalTime-CB-8TxPUSCH-r18", Optional: true},
		{Name: "additionalTime-NonCB-8TxPUSCH-r18", Optional: true},
	},
}

const (
	FeatureSetUplink_v1850_Pusch_DMRS8Tx_r18_Rel15 = 0
	FeatureSetUplink_v1850_Pusch_DMRS8Tx_r18_Both  = 1
)

var featureSetUplinkV1850PuschDMRS8TxR18Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{FeatureSetUplink_v1850_Pusch_DMRS8Tx_r18_Rel15, FeatureSetUplink_v1850_Pusch_DMRS8Tx_r18_Both},
}

var featureSetUplinkV1850AdditionalTimeCB8TxPUSCHR18Constraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "scs-15kHz-r18", Optional: true},
		{Name: "scs-30kHz-r18", Optional: true},
		{Name: "scs-60kHz-r18", Optional: true},
		{Name: "scs-120kHz-r18", Optional: true},
		{Name: "scs-480kHz-r18", Optional: true},
		{Name: "scs-960kHz-r18", Optional: true},
	},
}

const (
	FeatureSetUplink_v1850_AdditionalTime_CB_8TxPUSCH_r18_Scs_15kHz_r18_Sym1 = 0
	FeatureSetUplink_v1850_AdditionalTime_CB_8TxPUSCH_r18_Scs_15kHz_r18_Sym2 = 1
	FeatureSetUplink_v1850_AdditionalTime_CB_8TxPUSCH_r18_Scs_15kHz_r18_Sym4 = 2
)

var featureSetUplinkV1850AdditionalTimeCB8TxPUSCHR18Scs15kHzR18Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{FeatureSetUplink_v1850_AdditionalTime_CB_8TxPUSCH_r18_Scs_15kHz_r18_Sym1, FeatureSetUplink_v1850_AdditionalTime_CB_8TxPUSCH_r18_Scs_15kHz_r18_Sym2, FeatureSetUplink_v1850_AdditionalTime_CB_8TxPUSCH_r18_Scs_15kHz_r18_Sym4},
}

const (
	FeatureSetUplink_v1850_AdditionalTime_CB_8TxPUSCH_r18_Scs_30kHz_r18_Sym1 = 0
	FeatureSetUplink_v1850_AdditionalTime_CB_8TxPUSCH_r18_Scs_30kHz_r18_Sym2 = 1
	FeatureSetUplink_v1850_AdditionalTime_CB_8TxPUSCH_r18_Scs_30kHz_r18_Sym4 = 2
	FeatureSetUplink_v1850_AdditionalTime_CB_8TxPUSCH_r18_Scs_30kHz_r18_Sym8 = 3
)

var featureSetUplinkV1850AdditionalTimeCB8TxPUSCHR18Scs30kHzR18Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{FeatureSetUplink_v1850_AdditionalTime_CB_8TxPUSCH_r18_Scs_30kHz_r18_Sym1, FeatureSetUplink_v1850_AdditionalTime_CB_8TxPUSCH_r18_Scs_30kHz_r18_Sym2, FeatureSetUplink_v1850_AdditionalTime_CB_8TxPUSCH_r18_Scs_30kHz_r18_Sym4, FeatureSetUplink_v1850_AdditionalTime_CB_8TxPUSCH_r18_Scs_30kHz_r18_Sym8},
}

const (
	FeatureSetUplink_v1850_AdditionalTime_CB_8TxPUSCH_r18_Scs_60kHz_r18_Sym2  = 0
	FeatureSetUplink_v1850_AdditionalTime_CB_8TxPUSCH_r18_Scs_60kHz_r18_Sym4  = 1
	FeatureSetUplink_v1850_AdditionalTime_CB_8TxPUSCH_r18_Scs_60kHz_r18_Sym8  = 2
	FeatureSetUplink_v1850_AdditionalTime_CB_8TxPUSCH_r18_Scs_60kHz_r18_Sym16 = 3
)

var featureSetUplinkV1850AdditionalTimeCB8TxPUSCHR18Scs60kHzR18Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{FeatureSetUplink_v1850_AdditionalTime_CB_8TxPUSCH_r18_Scs_60kHz_r18_Sym2, FeatureSetUplink_v1850_AdditionalTime_CB_8TxPUSCH_r18_Scs_60kHz_r18_Sym4, FeatureSetUplink_v1850_AdditionalTime_CB_8TxPUSCH_r18_Scs_60kHz_r18_Sym8, FeatureSetUplink_v1850_AdditionalTime_CB_8TxPUSCH_r18_Scs_60kHz_r18_Sym16},
}

const (
	FeatureSetUplink_v1850_AdditionalTime_CB_8TxPUSCH_r18_Scs_120kHz_r18_Sym4  = 0
	FeatureSetUplink_v1850_AdditionalTime_CB_8TxPUSCH_r18_Scs_120kHz_r18_Sym8  = 1
	FeatureSetUplink_v1850_AdditionalTime_CB_8TxPUSCH_r18_Scs_120kHz_r18_Sym16 = 2
	FeatureSetUplink_v1850_AdditionalTime_CB_8TxPUSCH_r18_Scs_120kHz_r18_Sym32 = 3
)

var featureSetUplinkV1850AdditionalTimeCB8TxPUSCHR18Scs120kHzR18Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{FeatureSetUplink_v1850_AdditionalTime_CB_8TxPUSCH_r18_Scs_120kHz_r18_Sym4, FeatureSetUplink_v1850_AdditionalTime_CB_8TxPUSCH_r18_Scs_120kHz_r18_Sym8, FeatureSetUplink_v1850_AdditionalTime_CB_8TxPUSCH_r18_Scs_120kHz_r18_Sym16, FeatureSetUplink_v1850_AdditionalTime_CB_8TxPUSCH_r18_Scs_120kHz_r18_Sym32},
}

const (
	FeatureSetUplink_v1850_AdditionalTime_CB_8TxPUSCH_r18_Scs_480kHz_r18_Sym16  = 0
	FeatureSetUplink_v1850_AdditionalTime_CB_8TxPUSCH_r18_Scs_480kHz_r18_Sym32  = 1
	FeatureSetUplink_v1850_AdditionalTime_CB_8TxPUSCH_r18_Scs_480kHz_r18_Sym64  = 2
	FeatureSetUplink_v1850_AdditionalTime_CB_8TxPUSCH_r18_Scs_480kHz_r18_Sym128 = 3
)

var featureSetUplinkV1850AdditionalTimeCB8TxPUSCHR18Scs480kHzR18Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{FeatureSetUplink_v1850_AdditionalTime_CB_8TxPUSCH_r18_Scs_480kHz_r18_Sym16, FeatureSetUplink_v1850_AdditionalTime_CB_8TxPUSCH_r18_Scs_480kHz_r18_Sym32, FeatureSetUplink_v1850_AdditionalTime_CB_8TxPUSCH_r18_Scs_480kHz_r18_Sym64, FeatureSetUplink_v1850_AdditionalTime_CB_8TxPUSCH_r18_Scs_480kHz_r18_Sym128},
}

const (
	FeatureSetUplink_v1850_AdditionalTime_CB_8TxPUSCH_r18_Scs_960kHz_r18_Sym32  = 0
	FeatureSetUplink_v1850_AdditionalTime_CB_8TxPUSCH_r18_Scs_960kHz_r18_Sym64  = 1
	FeatureSetUplink_v1850_AdditionalTime_CB_8TxPUSCH_r18_Scs_960kHz_r18_Sym128 = 2
	FeatureSetUplink_v1850_AdditionalTime_CB_8TxPUSCH_r18_Scs_960kHz_r18_Sym256 = 3
)

var featureSetUplinkV1850AdditionalTimeCB8TxPUSCHR18Scs960kHzR18Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{FeatureSetUplink_v1850_AdditionalTime_CB_8TxPUSCH_r18_Scs_960kHz_r18_Sym32, FeatureSetUplink_v1850_AdditionalTime_CB_8TxPUSCH_r18_Scs_960kHz_r18_Sym64, FeatureSetUplink_v1850_AdditionalTime_CB_8TxPUSCH_r18_Scs_960kHz_r18_Sym128, FeatureSetUplink_v1850_AdditionalTime_CB_8TxPUSCH_r18_Scs_960kHz_r18_Sym256},
}

var featureSetUplinkV1850AdditionalTimeNonCB8TxPUSCHR18Constraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "scs-15kHz-r18", Optional: true},
		{Name: "scs-30kHz-r18", Optional: true},
		{Name: "scs-60kHz-r18", Optional: true},
		{Name: "scs-120kHz-r18", Optional: true},
		{Name: "scs-480kHz-r18", Optional: true},
		{Name: "scs-960kHz-r18", Optional: true},
	},
}

const (
	FeatureSetUplink_v1850_AdditionalTime_NonCB_8TxPUSCH_r18_Scs_15kHz_r18_Sym1 = 0
	FeatureSetUplink_v1850_AdditionalTime_NonCB_8TxPUSCH_r18_Scs_15kHz_r18_Sym2 = 1
	FeatureSetUplink_v1850_AdditionalTime_NonCB_8TxPUSCH_r18_Scs_15kHz_r18_Sym4 = 2
)

var featureSetUplinkV1850AdditionalTimeNonCB8TxPUSCHR18Scs15kHzR18Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{FeatureSetUplink_v1850_AdditionalTime_NonCB_8TxPUSCH_r18_Scs_15kHz_r18_Sym1, FeatureSetUplink_v1850_AdditionalTime_NonCB_8TxPUSCH_r18_Scs_15kHz_r18_Sym2, FeatureSetUplink_v1850_AdditionalTime_NonCB_8TxPUSCH_r18_Scs_15kHz_r18_Sym4},
}

const (
	FeatureSetUplink_v1850_AdditionalTime_NonCB_8TxPUSCH_r18_Scs_30kHz_r18_Sym1 = 0
	FeatureSetUplink_v1850_AdditionalTime_NonCB_8TxPUSCH_r18_Scs_30kHz_r18_Sym2 = 1
	FeatureSetUplink_v1850_AdditionalTime_NonCB_8TxPUSCH_r18_Scs_30kHz_r18_Sym4 = 2
	FeatureSetUplink_v1850_AdditionalTime_NonCB_8TxPUSCH_r18_Scs_30kHz_r18_Sym8 = 3
)

var featureSetUplinkV1850AdditionalTimeNonCB8TxPUSCHR18Scs30kHzR18Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{FeatureSetUplink_v1850_AdditionalTime_NonCB_8TxPUSCH_r18_Scs_30kHz_r18_Sym1, FeatureSetUplink_v1850_AdditionalTime_NonCB_8TxPUSCH_r18_Scs_30kHz_r18_Sym2, FeatureSetUplink_v1850_AdditionalTime_NonCB_8TxPUSCH_r18_Scs_30kHz_r18_Sym4, FeatureSetUplink_v1850_AdditionalTime_NonCB_8TxPUSCH_r18_Scs_30kHz_r18_Sym8},
}

const (
	FeatureSetUplink_v1850_AdditionalTime_NonCB_8TxPUSCH_r18_Scs_60kHz_r18_Sym2  = 0
	FeatureSetUplink_v1850_AdditionalTime_NonCB_8TxPUSCH_r18_Scs_60kHz_r18_Sym4  = 1
	FeatureSetUplink_v1850_AdditionalTime_NonCB_8TxPUSCH_r18_Scs_60kHz_r18_Sym8  = 2
	FeatureSetUplink_v1850_AdditionalTime_NonCB_8TxPUSCH_r18_Scs_60kHz_r18_Sym16 = 3
)

var featureSetUplinkV1850AdditionalTimeNonCB8TxPUSCHR18Scs60kHzR18Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{FeatureSetUplink_v1850_AdditionalTime_NonCB_8TxPUSCH_r18_Scs_60kHz_r18_Sym2, FeatureSetUplink_v1850_AdditionalTime_NonCB_8TxPUSCH_r18_Scs_60kHz_r18_Sym4, FeatureSetUplink_v1850_AdditionalTime_NonCB_8TxPUSCH_r18_Scs_60kHz_r18_Sym8, FeatureSetUplink_v1850_AdditionalTime_NonCB_8TxPUSCH_r18_Scs_60kHz_r18_Sym16},
}

const (
	FeatureSetUplink_v1850_AdditionalTime_NonCB_8TxPUSCH_r18_Scs_120kHz_r18_Sym4  = 0
	FeatureSetUplink_v1850_AdditionalTime_NonCB_8TxPUSCH_r18_Scs_120kHz_r18_Sym8  = 1
	FeatureSetUplink_v1850_AdditionalTime_NonCB_8TxPUSCH_r18_Scs_120kHz_r18_Sym16 = 2
	FeatureSetUplink_v1850_AdditionalTime_NonCB_8TxPUSCH_r18_Scs_120kHz_r18_Sym32 = 3
)

var featureSetUplinkV1850AdditionalTimeNonCB8TxPUSCHR18Scs120kHzR18Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{FeatureSetUplink_v1850_AdditionalTime_NonCB_8TxPUSCH_r18_Scs_120kHz_r18_Sym4, FeatureSetUplink_v1850_AdditionalTime_NonCB_8TxPUSCH_r18_Scs_120kHz_r18_Sym8, FeatureSetUplink_v1850_AdditionalTime_NonCB_8TxPUSCH_r18_Scs_120kHz_r18_Sym16, FeatureSetUplink_v1850_AdditionalTime_NonCB_8TxPUSCH_r18_Scs_120kHz_r18_Sym32},
}

const (
	FeatureSetUplink_v1850_AdditionalTime_NonCB_8TxPUSCH_r18_Scs_480kHz_r18_Sym16  = 0
	FeatureSetUplink_v1850_AdditionalTime_NonCB_8TxPUSCH_r18_Scs_480kHz_r18_Sym32  = 1
	FeatureSetUplink_v1850_AdditionalTime_NonCB_8TxPUSCH_r18_Scs_480kHz_r18_Sym64  = 2
	FeatureSetUplink_v1850_AdditionalTime_NonCB_8TxPUSCH_r18_Scs_480kHz_r18_Sym128 = 3
)

var featureSetUplinkV1850AdditionalTimeNonCB8TxPUSCHR18Scs480kHzR18Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{FeatureSetUplink_v1850_AdditionalTime_NonCB_8TxPUSCH_r18_Scs_480kHz_r18_Sym16, FeatureSetUplink_v1850_AdditionalTime_NonCB_8TxPUSCH_r18_Scs_480kHz_r18_Sym32, FeatureSetUplink_v1850_AdditionalTime_NonCB_8TxPUSCH_r18_Scs_480kHz_r18_Sym64, FeatureSetUplink_v1850_AdditionalTime_NonCB_8TxPUSCH_r18_Scs_480kHz_r18_Sym128},
}

const (
	FeatureSetUplink_v1850_AdditionalTime_NonCB_8TxPUSCH_r18_Scs_960kHz_r18_Sym32  = 0
	FeatureSetUplink_v1850_AdditionalTime_NonCB_8TxPUSCH_r18_Scs_960kHz_r18_Sym64  = 1
	FeatureSetUplink_v1850_AdditionalTime_NonCB_8TxPUSCH_r18_Scs_960kHz_r18_Sym128 = 2
	FeatureSetUplink_v1850_AdditionalTime_NonCB_8TxPUSCH_r18_Scs_960kHz_r18_Sym256 = 3
)

var featureSetUplinkV1850AdditionalTimeNonCB8TxPUSCHR18Scs960kHzR18Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{FeatureSetUplink_v1850_AdditionalTime_NonCB_8TxPUSCH_r18_Scs_960kHz_r18_Sym32, FeatureSetUplink_v1850_AdditionalTime_NonCB_8TxPUSCH_r18_Scs_960kHz_r18_Sym64, FeatureSetUplink_v1850_AdditionalTime_NonCB_8TxPUSCH_r18_Scs_960kHz_r18_Sym128, FeatureSetUplink_v1850_AdditionalTime_NonCB_8TxPUSCH_r18_Scs_960kHz_r18_Sym256},
}

type FeatureSetUplink_v1850 struct {
	Pusch_DMRS8Tx_r18              *int64
	AdditionalTime_CB_8TxPUSCH_r18 *struct {
		Scs_15kHz_r18  *int64
		Scs_30kHz_r18  *int64
		Scs_60kHz_r18  *int64
		Scs_120kHz_r18 *int64
		Scs_480kHz_r18 *int64
		Scs_960kHz_r18 *int64
	}
	AdditionalTime_NonCB_8TxPUSCH_r18 *struct {
		Scs_15kHz_r18  *int64
		Scs_30kHz_r18  *int64
		Scs_60kHz_r18  *int64
		Scs_120kHz_r18 *int64
		Scs_480kHz_r18 *int64
		Scs_960kHz_r18 *int64
	}
}

func (ie *FeatureSetUplink_v1850) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(featureSetUplinkV1850Constraints)

	// 1. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.Pusch_DMRS8Tx_r18 != nil, ie.AdditionalTime_CB_8TxPUSCH_r18 != nil, ie.AdditionalTime_NonCB_8TxPUSCH_r18 != nil}); err != nil {
		return err
	}

	// 2. pusch-DMRS8Tx-r18: enumerated
	{
		if ie.Pusch_DMRS8Tx_r18 != nil {
			if err := e.EncodeEnumerated(*ie.Pusch_DMRS8Tx_r18, featureSetUplinkV1850PuschDMRS8TxR18Constraints); err != nil {
				return err
			}
		}
	}

	// 3. additionalTime-CB-8TxPUSCH-r18: sequence
	{
		if ie.AdditionalTime_CB_8TxPUSCH_r18 != nil {
			c := ie.AdditionalTime_CB_8TxPUSCH_r18
			featureSetUplinkV1850AdditionalTimeCB8TxPUSCHR18Seq := e.NewSequenceEncoder(featureSetUplinkV1850AdditionalTimeCB8TxPUSCHR18Constraints)
			if err := featureSetUplinkV1850AdditionalTimeCB8TxPUSCHR18Seq.EncodePreamble([]bool{c.Scs_15kHz_r18 != nil, c.Scs_30kHz_r18 != nil, c.Scs_60kHz_r18 != nil, c.Scs_120kHz_r18 != nil, c.Scs_480kHz_r18 != nil, c.Scs_960kHz_r18 != nil}); err != nil {
				return err
			}
			if c.Scs_15kHz_r18 != nil {
				if err := e.EncodeEnumerated((*c.Scs_15kHz_r18), featureSetUplinkV1850AdditionalTimeCB8TxPUSCHR18Scs15kHzR18Constraints); err != nil {
					return err
				}
			}
			if c.Scs_30kHz_r18 != nil {
				if err := e.EncodeEnumerated((*c.Scs_30kHz_r18), featureSetUplinkV1850AdditionalTimeCB8TxPUSCHR18Scs30kHzR18Constraints); err != nil {
					return err
				}
			}
			if c.Scs_60kHz_r18 != nil {
				if err := e.EncodeEnumerated((*c.Scs_60kHz_r18), featureSetUplinkV1850AdditionalTimeCB8TxPUSCHR18Scs60kHzR18Constraints); err != nil {
					return err
				}
			}
			if c.Scs_120kHz_r18 != nil {
				if err := e.EncodeEnumerated((*c.Scs_120kHz_r18), featureSetUplinkV1850AdditionalTimeCB8TxPUSCHR18Scs120kHzR18Constraints); err != nil {
					return err
				}
			}
			if c.Scs_480kHz_r18 != nil {
				if err := e.EncodeEnumerated((*c.Scs_480kHz_r18), featureSetUplinkV1850AdditionalTimeCB8TxPUSCHR18Scs480kHzR18Constraints); err != nil {
					return err
				}
			}
			if c.Scs_960kHz_r18 != nil {
				if err := e.EncodeEnumerated((*c.Scs_960kHz_r18), featureSetUplinkV1850AdditionalTimeCB8TxPUSCHR18Scs960kHzR18Constraints); err != nil {
					return err
				}
			}
		}
	}

	// 4. additionalTime-NonCB-8TxPUSCH-r18: sequence
	{
		if ie.AdditionalTime_NonCB_8TxPUSCH_r18 != nil {
			c := ie.AdditionalTime_NonCB_8TxPUSCH_r18
			featureSetUplinkV1850AdditionalTimeNonCB8TxPUSCHR18Seq := e.NewSequenceEncoder(featureSetUplinkV1850AdditionalTimeNonCB8TxPUSCHR18Constraints)
			if err := featureSetUplinkV1850AdditionalTimeNonCB8TxPUSCHR18Seq.EncodePreamble([]bool{c.Scs_15kHz_r18 != nil, c.Scs_30kHz_r18 != nil, c.Scs_60kHz_r18 != nil, c.Scs_120kHz_r18 != nil, c.Scs_480kHz_r18 != nil, c.Scs_960kHz_r18 != nil}); err != nil {
				return err
			}
			if c.Scs_15kHz_r18 != nil {
				if err := e.EncodeEnumerated((*c.Scs_15kHz_r18), featureSetUplinkV1850AdditionalTimeNonCB8TxPUSCHR18Scs15kHzR18Constraints); err != nil {
					return err
				}
			}
			if c.Scs_30kHz_r18 != nil {
				if err := e.EncodeEnumerated((*c.Scs_30kHz_r18), featureSetUplinkV1850AdditionalTimeNonCB8TxPUSCHR18Scs30kHzR18Constraints); err != nil {
					return err
				}
			}
			if c.Scs_60kHz_r18 != nil {
				if err := e.EncodeEnumerated((*c.Scs_60kHz_r18), featureSetUplinkV1850AdditionalTimeNonCB8TxPUSCHR18Scs60kHzR18Constraints); err != nil {
					return err
				}
			}
			if c.Scs_120kHz_r18 != nil {
				if err := e.EncodeEnumerated((*c.Scs_120kHz_r18), featureSetUplinkV1850AdditionalTimeNonCB8TxPUSCHR18Scs120kHzR18Constraints); err != nil {
					return err
				}
			}
			if c.Scs_480kHz_r18 != nil {
				if err := e.EncodeEnumerated((*c.Scs_480kHz_r18), featureSetUplinkV1850AdditionalTimeNonCB8TxPUSCHR18Scs480kHzR18Constraints); err != nil {
					return err
				}
			}
			if c.Scs_960kHz_r18 != nil {
				if err := e.EncodeEnumerated((*c.Scs_960kHz_r18), featureSetUplinkV1850AdditionalTimeNonCB8TxPUSCHR18Scs960kHzR18Constraints); err != nil {
					return err
				}
			}
		}
	}

	return nil
}

func (ie *FeatureSetUplink_v1850) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(featureSetUplinkV1850Constraints)

	// 1. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 2. pusch-DMRS8Tx-r18: enumerated
	{
		if seq.IsComponentPresent(0) {
			idx, err := d.DecodeEnumerated(featureSetUplinkV1850PuschDMRS8TxR18Constraints)
			if err != nil {
				return err
			}
			ie.Pusch_DMRS8Tx_r18 = &idx
		}
	}

	// 3. additionalTime-CB-8TxPUSCH-r18: sequence
	{
		if seq.IsComponentPresent(1) {
			ie.AdditionalTime_CB_8TxPUSCH_r18 = &struct {
				Scs_15kHz_r18  *int64
				Scs_30kHz_r18  *int64
				Scs_60kHz_r18  *int64
				Scs_120kHz_r18 *int64
				Scs_480kHz_r18 *int64
				Scs_960kHz_r18 *int64
			}{}
			c := ie.AdditionalTime_CB_8TxPUSCH_r18
			featureSetUplinkV1850AdditionalTimeCB8TxPUSCHR18Seq := d.NewSequenceDecoder(featureSetUplinkV1850AdditionalTimeCB8TxPUSCHR18Constraints)
			if err := featureSetUplinkV1850AdditionalTimeCB8TxPUSCHR18Seq.DecodePreamble(); err != nil {
				return err
			}
			if featureSetUplinkV1850AdditionalTimeCB8TxPUSCHR18Seq.IsComponentPresent(0) {
				c.Scs_15kHz_r18 = new(int64)
				v, err := d.DecodeEnumerated(featureSetUplinkV1850AdditionalTimeCB8TxPUSCHR18Scs15kHzR18Constraints)
				if err != nil {
					return err
				}
				(*c.Scs_15kHz_r18) = v
			}
			if featureSetUplinkV1850AdditionalTimeCB8TxPUSCHR18Seq.IsComponentPresent(1) {
				c.Scs_30kHz_r18 = new(int64)
				v, err := d.DecodeEnumerated(featureSetUplinkV1850AdditionalTimeCB8TxPUSCHR18Scs30kHzR18Constraints)
				if err != nil {
					return err
				}
				(*c.Scs_30kHz_r18) = v
			}
			if featureSetUplinkV1850AdditionalTimeCB8TxPUSCHR18Seq.IsComponentPresent(2) {
				c.Scs_60kHz_r18 = new(int64)
				v, err := d.DecodeEnumerated(featureSetUplinkV1850AdditionalTimeCB8TxPUSCHR18Scs60kHzR18Constraints)
				if err != nil {
					return err
				}
				(*c.Scs_60kHz_r18) = v
			}
			if featureSetUplinkV1850AdditionalTimeCB8TxPUSCHR18Seq.IsComponentPresent(3) {
				c.Scs_120kHz_r18 = new(int64)
				v, err := d.DecodeEnumerated(featureSetUplinkV1850AdditionalTimeCB8TxPUSCHR18Scs120kHzR18Constraints)
				if err != nil {
					return err
				}
				(*c.Scs_120kHz_r18) = v
			}
			if featureSetUplinkV1850AdditionalTimeCB8TxPUSCHR18Seq.IsComponentPresent(4) {
				c.Scs_480kHz_r18 = new(int64)
				v, err := d.DecodeEnumerated(featureSetUplinkV1850AdditionalTimeCB8TxPUSCHR18Scs480kHzR18Constraints)
				if err != nil {
					return err
				}
				(*c.Scs_480kHz_r18) = v
			}
			if featureSetUplinkV1850AdditionalTimeCB8TxPUSCHR18Seq.IsComponentPresent(5) {
				c.Scs_960kHz_r18 = new(int64)
				v, err := d.DecodeEnumerated(featureSetUplinkV1850AdditionalTimeCB8TxPUSCHR18Scs960kHzR18Constraints)
				if err != nil {
					return err
				}
				(*c.Scs_960kHz_r18) = v
			}
		}
	}

	// 4. additionalTime-NonCB-8TxPUSCH-r18: sequence
	{
		if seq.IsComponentPresent(2) {
			ie.AdditionalTime_NonCB_8TxPUSCH_r18 = &struct {
				Scs_15kHz_r18  *int64
				Scs_30kHz_r18  *int64
				Scs_60kHz_r18  *int64
				Scs_120kHz_r18 *int64
				Scs_480kHz_r18 *int64
				Scs_960kHz_r18 *int64
			}{}
			c := ie.AdditionalTime_NonCB_8TxPUSCH_r18
			featureSetUplinkV1850AdditionalTimeNonCB8TxPUSCHR18Seq := d.NewSequenceDecoder(featureSetUplinkV1850AdditionalTimeNonCB8TxPUSCHR18Constraints)
			if err := featureSetUplinkV1850AdditionalTimeNonCB8TxPUSCHR18Seq.DecodePreamble(); err != nil {
				return err
			}
			if featureSetUplinkV1850AdditionalTimeNonCB8TxPUSCHR18Seq.IsComponentPresent(0) {
				c.Scs_15kHz_r18 = new(int64)
				v, err := d.DecodeEnumerated(featureSetUplinkV1850AdditionalTimeNonCB8TxPUSCHR18Scs15kHzR18Constraints)
				if err != nil {
					return err
				}
				(*c.Scs_15kHz_r18) = v
			}
			if featureSetUplinkV1850AdditionalTimeNonCB8TxPUSCHR18Seq.IsComponentPresent(1) {
				c.Scs_30kHz_r18 = new(int64)
				v, err := d.DecodeEnumerated(featureSetUplinkV1850AdditionalTimeNonCB8TxPUSCHR18Scs30kHzR18Constraints)
				if err != nil {
					return err
				}
				(*c.Scs_30kHz_r18) = v
			}
			if featureSetUplinkV1850AdditionalTimeNonCB8TxPUSCHR18Seq.IsComponentPresent(2) {
				c.Scs_60kHz_r18 = new(int64)
				v, err := d.DecodeEnumerated(featureSetUplinkV1850AdditionalTimeNonCB8TxPUSCHR18Scs60kHzR18Constraints)
				if err != nil {
					return err
				}
				(*c.Scs_60kHz_r18) = v
			}
			if featureSetUplinkV1850AdditionalTimeNonCB8TxPUSCHR18Seq.IsComponentPresent(3) {
				c.Scs_120kHz_r18 = new(int64)
				v, err := d.DecodeEnumerated(featureSetUplinkV1850AdditionalTimeNonCB8TxPUSCHR18Scs120kHzR18Constraints)
				if err != nil {
					return err
				}
				(*c.Scs_120kHz_r18) = v
			}
			if featureSetUplinkV1850AdditionalTimeNonCB8TxPUSCHR18Seq.IsComponentPresent(4) {
				c.Scs_480kHz_r18 = new(int64)
				v, err := d.DecodeEnumerated(featureSetUplinkV1850AdditionalTimeNonCB8TxPUSCHR18Scs480kHzR18Constraints)
				if err != nil {
					return err
				}
				(*c.Scs_480kHz_r18) = v
			}
			if featureSetUplinkV1850AdditionalTimeNonCB8TxPUSCHR18Seq.IsComponentPresent(5) {
				c.Scs_960kHz_r18 = new(int64)
				v, err := d.DecodeEnumerated(featureSetUplinkV1850AdditionalTimeNonCB8TxPUSCHR18Scs960kHzR18Constraints)
				if err != nil {
					return err
				}
				(*c.Scs_960kHz_r18) = v
			}
		}
	}

	return nil
}
