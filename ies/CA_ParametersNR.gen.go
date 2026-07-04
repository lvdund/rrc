// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: CA-ParametersNR (line 17270).

var cAParametersNRConstraints = per.SequenceConstraints{
	Extensible: true,
	RootComponents: []per.ComponentInfo{
		{Name: "dummy", Optional: true},
		{Name: "parallelTxSRS-PUCCH-PUSCH", Optional: true},
		{Name: "parallelTxPRACH-SRS-PUCCH-PUSCH", Optional: true},
		{Name: "simultaneousRxTxInterBandCA", Optional: true},
		{Name: "simultaneousRxTxSUL", Optional: true},
		{Name: "diffNumerologyAcrossPUCCH-Group", Optional: true},
		{Name: "diffNumerologyWithinPUCCH-GroupSmallerSCS", Optional: true},
		{Name: "supportedNumberTAG", Optional: true},
	},
}

const (
	CA_ParametersNR_Dummy_Supported = 0
)

var cAParametersNRDummyConstraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{CA_ParametersNR_Dummy_Supported},
}

const (
	CA_ParametersNR_ParallelTxSRS_PUCCH_PUSCH_Supported = 0
)

var cAParametersNRParallelTxSRSPUCCHPUSCHConstraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{CA_ParametersNR_ParallelTxSRS_PUCCH_PUSCH_Supported},
}

const (
	CA_ParametersNR_ParallelTxPRACH_SRS_PUCCH_PUSCH_Supported = 0
)

var cAParametersNRParallelTxPRACHSRSPUCCHPUSCHConstraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{CA_ParametersNR_ParallelTxPRACH_SRS_PUCCH_PUSCH_Supported},
}

const (
	CA_ParametersNR_SimultaneousRxTxInterBandCA_Supported = 0
)

var cAParametersNRSimultaneousRxTxInterBandCAConstraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{CA_ParametersNR_SimultaneousRxTxInterBandCA_Supported},
}

const (
	CA_ParametersNR_SimultaneousRxTxSUL_Supported = 0
)

var cAParametersNRSimultaneousRxTxSULConstraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{CA_ParametersNR_SimultaneousRxTxSUL_Supported},
}

const (
	CA_ParametersNR_DiffNumerologyAcrossPUCCH_Group_Supported = 0
)

var cAParametersNRDiffNumerologyAcrossPUCCHGroupConstraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{CA_ParametersNR_DiffNumerologyAcrossPUCCH_Group_Supported},
}

const (
	CA_ParametersNR_DiffNumerologyWithinPUCCH_GroupSmallerSCS_Supported = 0
)

var cAParametersNRDiffNumerologyWithinPUCCHGroupSmallerSCSConstraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{CA_ParametersNR_DiffNumerologyWithinPUCCH_GroupSmallerSCS_Supported},
}

const (
	CA_ParametersNR_SupportedNumberTAG_N2 = 0
	CA_ParametersNR_SupportedNumberTAG_N3 = 1
	CA_ParametersNR_SupportedNumberTAG_N4 = 2
)

var cAParametersNRSupportedNumberTAGConstraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{CA_ParametersNR_SupportedNumberTAG_N2, CA_ParametersNR_SupportedNumberTAG_N3, CA_ParametersNR_SupportedNumberTAG_N4},
}

type CA_ParametersNR struct {
	Dummy                                     *int64
	ParallelTxSRS_PUCCH_PUSCH                 *int64
	ParallelTxPRACH_SRS_PUCCH_PUSCH           *int64
	SimultaneousRxTxInterBandCA               *int64
	SimultaneousRxTxSUL                       *int64
	DiffNumerologyAcrossPUCCH_Group           *int64
	DiffNumerologyWithinPUCCH_GroupSmallerSCS *int64
	SupportedNumberTAG                        *int64
}

func (ie *CA_ParametersNR) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(cAParametersNRConstraints)

	// 1. SEQUENCE extension bit.
	if err := seq.EncodeExtensionBit(false); err != nil {
		return err
	}

	// 2. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.Dummy != nil, ie.ParallelTxSRS_PUCCH_PUSCH != nil, ie.ParallelTxPRACH_SRS_PUCCH_PUSCH != nil, ie.SimultaneousRxTxInterBandCA != nil, ie.SimultaneousRxTxSUL != nil, ie.DiffNumerologyAcrossPUCCH_Group != nil, ie.DiffNumerologyWithinPUCCH_GroupSmallerSCS != nil, ie.SupportedNumberTAG != nil}); err != nil {
		return err
	}

	// 3. dummy: enumerated
	{
		if ie.Dummy != nil {
			if err := e.EncodeEnumerated(*ie.Dummy, cAParametersNRDummyConstraints); err != nil {
				return err
			}
		}
	}

	// 4. parallelTxSRS-PUCCH-PUSCH: enumerated
	{
		if ie.ParallelTxSRS_PUCCH_PUSCH != nil {
			if err := e.EncodeEnumerated(*ie.ParallelTxSRS_PUCCH_PUSCH, cAParametersNRParallelTxSRSPUCCHPUSCHConstraints); err != nil {
				return err
			}
		}
	}

	// 5. parallelTxPRACH-SRS-PUCCH-PUSCH: enumerated
	{
		if ie.ParallelTxPRACH_SRS_PUCCH_PUSCH != nil {
			if err := e.EncodeEnumerated(*ie.ParallelTxPRACH_SRS_PUCCH_PUSCH, cAParametersNRParallelTxPRACHSRSPUCCHPUSCHConstraints); err != nil {
				return err
			}
		}
	}

	// 6. simultaneousRxTxInterBandCA: enumerated
	{
		if ie.SimultaneousRxTxInterBandCA != nil {
			if err := e.EncodeEnumerated(*ie.SimultaneousRxTxInterBandCA, cAParametersNRSimultaneousRxTxInterBandCAConstraints); err != nil {
				return err
			}
		}
	}

	// 7. simultaneousRxTxSUL: enumerated
	{
		if ie.SimultaneousRxTxSUL != nil {
			if err := e.EncodeEnumerated(*ie.SimultaneousRxTxSUL, cAParametersNRSimultaneousRxTxSULConstraints); err != nil {
				return err
			}
		}
	}

	// 8. diffNumerologyAcrossPUCCH-Group: enumerated
	{
		if ie.DiffNumerologyAcrossPUCCH_Group != nil {
			if err := e.EncodeEnumerated(*ie.DiffNumerologyAcrossPUCCH_Group, cAParametersNRDiffNumerologyAcrossPUCCHGroupConstraints); err != nil {
				return err
			}
		}
	}

	// 9. diffNumerologyWithinPUCCH-GroupSmallerSCS: enumerated
	{
		if ie.DiffNumerologyWithinPUCCH_GroupSmallerSCS != nil {
			if err := e.EncodeEnumerated(*ie.DiffNumerologyWithinPUCCH_GroupSmallerSCS, cAParametersNRDiffNumerologyWithinPUCCHGroupSmallerSCSConstraints); err != nil {
				return err
			}
		}
	}

	// 10. supportedNumberTAG: enumerated
	{
		if ie.SupportedNumberTAG != nil {
			if err := e.EncodeEnumerated(*ie.SupportedNumberTAG, cAParametersNRSupportedNumberTAGConstraints); err != nil {
				return err
			}
		}
	}

	return nil
}

func (ie *CA_ParametersNR) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(cAParametersNRConstraints)

	// 1. SEQUENCE extension bit.
	if err := seq.DecodeExtensionBit(); err != nil {
		return err
	}

	// 2. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 3. dummy: enumerated
	{
		if seq.IsComponentPresent(0) {
			idx, err := d.DecodeEnumerated(cAParametersNRDummyConstraints)
			if err != nil {
				return err
			}
			ie.Dummy = &idx
		}
	}

	// 4. parallelTxSRS-PUCCH-PUSCH: enumerated
	{
		if seq.IsComponentPresent(1) {
			idx, err := d.DecodeEnumerated(cAParametersNRParallelTxSRSPUCCHPUSCHConstraints)
			if err != nil {
				return err
			}
			ie.ParallelTxSRS_PUCCH_PUSCH = &idx
		}
	}

	// 5. parallelTxPRACH-SRS-PUCCH-PUSCH: enumerated
	{
		if seq.IsComponentPresent(2) {
			idx, err := d.DecodeEnumerated(cAParametersNRParallelTxPRACHSRSPUCCHPUSCHConstraints)
			if err != nil {
				return err
			}
			ie.ParallelTxPRACH_SRS_PUCCH_PUSCH = &idx
		}
	}

	// 6. simultaneousRxTxInterBandCA: enumerated
	{
		if seq.IsComponentPresent(3) {
			idx, err := d.DecodeEnumerated(cAParametersNRSimultaneousRxTxInterBandCAConstraints)
			if err != nil {
				return err
			}
			ie.SimultaneousRxTxInterBandCA = &idx
		}
	}

	// 7. simultaneousRxTxSUL: enumerated
	{
		if seq.IsComponentPresent(4) {
			idx, err := d.DecodeEnumerated(cAParametersNRSimultaneousRxTxSULConstraints)
			if err != nil {
				return err
			}
			ie.SimultaneousRxTxSUL = &idx
		}
	}

	// 8. diffNumerologyAcrossPUCCH-Group: enumerated
	{
		if seq.IsComponentPresent(5) {
			idx, err := d.DecodeEnumerated(cAParametersNRDiffNumerologyAcrossPUCCHGroupConstraints)
			if err != nil {
				return err
			}
			ie.DiffNumerologyAcrossPUCCH_Group = &idx
		}
	}

	// 9. diffNumerologyWithinPUCCH-GroupSmallerSCS: enumerated
	{
		if seq.IsComponentPresent(6) {
			idx, err := d.DecodeEnumerated(cAParametersNRDiffNumerologyWithinPUCCHGroupSmallerSCSConstraints)
			if err != nil {
				return err
			}
			ie.DiffNumerologyWithinPUCCH_GroupSmallerSCS = &idx
		}
	}

	// 10. supportedNumberTAG: enumerated
	{
		if seq.IsComponentPresent(7) {
			idx, err := d.DecodeEnumerated(cAParametersNRSupportedNumberTAGConstraints)
			if err != nil {
				return err
			}
			ie.SupportedNumberTAG = &idx
		}
	}

	return nil
}
