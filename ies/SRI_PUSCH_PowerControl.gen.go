// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: SRI-PUSCH-PowerControl (line 12619).

var sRIPUSCHPowerControlConstraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "sri-PUSCH-PowerControlId"},
		{Name: "sri-PUSCH-PathlossReferenceRS-Id"},
		{Name: "sri-P0-PUSCH-AlphaSetId"},
		{Name: "sri-PUSCH-ClosedLoopIndex"},
	},
}

const (
	SRI_PUSCH_PowerControl_Sri_PUSCH_ClosedLoopIndex_I0 = 0
	SRI_PUSCH_PowerControl_Sri_PUSCH_ClosedLoopIndex_I1 = 1
)

var sRIPUSCHPowerControlSriPUSCHClosedLoopIndexConstraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{SRI_PUSCH_PowerControl_Sri_PUSCH_ClosedLoopIndex_I0, SRI_PUSCH_PowerControl_Sri_PUSCH_ClosedLoopIndex_I1},
}

type SRI_PUSCH_PowerControl struct {
	Sri_PUSCH_PowerControlId         SRI_PUSCH_PowerControlId
	Sri_PUSCH_PathlossReferenceRS_Id PUSCH_PathlossReferenceRS_Id
	Sri_P0_PUSCH_AlphaSetId          P0_PUSCH_AlphaSetId
	Sri_PUSCH_ClosedLoopIndex        int64
}

func (ie *SRI_PUSCH_PowerControl) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(sRIPUSCHPowerControlConstraints)
	_ = seq

	// 1. sri-PUSCH-PowerControlId: ref
	{
		if err := ie.Sri_PUSCH_PowerControlId.Encode(e); err != nil {
			return err
		}
	}

	// 2. sri-PUSCH-PathlossReferenceRS-Id: ref
	{
		if err := ie.Sri_PUSCH_PathlossReferenceRS_Id.Encode(e); err != nil {
			return err
		}
	}

	// 3. sri-P0-PUSCH-AlphaSetId: ref
	{
		if err := ie.Sri_P0_PUSCH_AlphaSetId.Encode(e); err != nil {
			return err
		}
	}

	// 4. sri-PUSCH-ClosedLoopIndex: enumerated
	{
		if err := e.EncodeEnumerated(ie.Sri_PUSCH_ClosedLoopIndex, sRIPUSCHPowerControlSriPUSCHClosedLoopIndexConstraints); err != nil {
			return err
		}
	}

	return nil
}

func (ie *SRI_PUSCH_PowerControl) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(sRIPUSCHPowerControlConstraints)
	_ = seq

	// 1. sri-PUSCH-PowerControlId: ref
	{
		if err := ie.Sri_PUSCH_PowerControlId.Decode(d); err != nil {
			return err
		}
	}

	// 2. sri-PUSCH-PathlossReferenceRS-Id: ref
	{
		if err := ie.Sri_PUSCH_PathlossReferenceRS_Id.Decode(d); err != nil {
			return err
		}
	}

	// 3. sri-P0-PUSCH-AlphaSetId: ref
	{
		if err := ie.Sri_P0_PUSCH_AlphaSetId.Decode(d); err != nil {
			return err
		}
	}

	// 4. sri-PUSCH-ClosedLoopIndex: enumerated
	{
		v3, err := d.DecodeEnumerated(sRIPUSCHPowerControlSriPUSCHClosedLoopIndexConstraints)
		if err != nil {
			return err
		}
		ie.Sri_PUSCH_ClosedLoopIndex = v3
	}

	return nil
}
