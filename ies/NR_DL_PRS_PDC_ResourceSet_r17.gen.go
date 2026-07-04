// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/rrc/common"
)

// Source: NR-DL-PRS-PDC-ResourceSet-r17 (line 10608).

var nRDLPRSPDCResourceSetR17Constraints = per.SequenceConstraints{
	Extensible: true,
	RootComponents: []per.ComponentInfo{
		{Name: "periodicityAndOffset-r17"},
		{Name: "numSymbols-r17"},
		{Name: "dl-PRS-ResourceBandwidth-r17"},
		{Name: "dl-PRS-StartPRB-r17"},
		{Name: "resourceList-r17"},
		{Name: "repFactorAndTimeGap-r17", Optional: true},
	},
}

const (
	NR_DL_PRS_PDC_ResourceSet_r17_NumSymbols_r17_N2       = 0
	NR_DL_PRS_PDC_ResourceSet_r17_NumSymbols_r17_N4       = 1
	NR_DL_PRS_PDC_ResourceSet_r17_NumSymbols_r17_N6       = 2
	NR_DL_PRS_PDC_ResourceSet_r17_NumSymbols_r17_N12      = 3
	NR_DL_PRS_PDC_ResourceSet_r17_NumSymbols_r17_N1_v1800 = 4
	NR_DL_PRS_PDC_ResourceSet_r17_NumSymbols_r17_Spare3   = 5
	NR_DL_PRS_PDC_ResourceSet_r17_NumSymbols_r17_Spare2   = 6
	NR_DL_PRS_PDC_ResourceSet_r17_NumSymbols_r17_Spare1   = 7
)

var nRDLPRSPDCResourceSetR17NumSymbolsR17Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{NR_DL_PRS_PDC_ResourceSet_r17_NumSymbols_r17_N2, NR_DL_PRS_PDC_ResourceSet_r17_NumSymbols_r17_N4, NR_DL_PRS_PDC_ResourceSet_r17_NumSymbols_r17_N6, NR_DL_PRS_PDC_ResourceSet_r17_NumSymbols_r17_N12, NR_DL_PRS_PDC_ResourceSet_r17_NumSymbols_r17_N1_v1800, NR_DL_PRS_PDC_ResourceSet_r17_NumSymbols_r17_Spare3, NR_DL_PRS_PDC_ResourceSet_r17_NumSymbols_r17_Spare2, NR_DL_PRS_PDC_ResourceSet_r17_NumSymbols_r17_Spare1},
}

var nRDLPRSPDCResourceSetR17DlPRSResourceBandwidthR17Constraints = per.Constrained(1, 63)

var nRDLPRSPDCResourceSetR17DlPRSStartPRBR17Constraints = per.Constrained(0, 2176)

var nRDLPRSPDCResourceSetR17ResourceListR17Constraints = per.SizeRange(1, common.MaxNrofPRS_ResourcesPerSet_r17)

type NR_DL_PRS_PDC_ResourceSet_r17 struct {
	PeriodicityAndOffset_r17     NR_DL_PRS_Periodicity_And_ResourceSetSlotOffset_r17
	NumSymbols_r17               int64
	Dl_PRS_ResourceBandwidth_r17 int64
	Dl_PRS_StartPRB_r17          int64
	ResourceList_r17             []NR_DL_PRS_Resource_r17
	RepFactorAndTimeGap_r17      *RepFactorAndTimeGap_r17
}

func (ie *NR_DL_PRS_PDC_ResourceSet_r17) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(nRDLPRSPDCResourceSetR17Constraints)

	// 1. SEQUENCE extension bit.
	if err := seq.EncodeExtensionBit(false); err != nil {
		return err
	}

	// 2. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.RepFactorAndTimeGap_r17 != nil}); err != nil {
		return err
	}

	// 3. periodicityAndOffset-r17: ref
	{
		if err := ie.PeriodicityAndOffset_r17.Encode(e); err != nil {
			return err
		}
	}

	// 4. numSymbols-r17: enumerated
	{
		if err := e.EncodeEnumerated(ie.NumSymbols_r17, nRDLPRSPDCResourceSetR17NumSymbolsR17Constraints); err != nil {
			return err
		}
	}

	// 5. dl-PRS-ResourceBandwidth-r17: integer
	{
		if err := e.EncodeInteger(ie.Dl_PRS_ResourceBandwidth_r17, nRDLPRSPDCResourceSetR17DlPRSResourceBandwidthR17Constraints); err != nil {
			return err
		}
	}

	// 6. dl-PRS-StartPRB-r17: integer
	{
		if err := e.EncodeInteger(ie.Dl_PRS_StartPRB_r17, nRDLPRSPDCResourceSetR17DlPRSStartPRBR17Constraints); err != nil {
			return err
		}
	}

	// 7. resourceList-r17: sequence-of
	{
		seqOf := e.NewSequenceOfEncoder(nRDLPRSPDCResourceSetR17ResourceListR17Constraints)
		if err := seqOf.EncodeLength(int64(len(ie.ResourceList_r17))); err != nil {
			return err
		}
		for i := range ie.ResourceList_r17 {
			if err := ie.ResourceList_r17[i].Encode(e); err != nil {
				return err
			}
		}
	}

	// 8. repFactorAndTimeGap-r17: ref
	{
		if ie.RepFactorAndTimeGap_r17 != nil {
			if err := ie.RepFactorAndTimeGap_r17.Encode(e); err != nil {
				return err
			}
		}
	}

	return nil
}

func (ie *NR_DL_PRS_PDC_ResourceSet_r17) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(nRDLPRSPDCResourceSetR17Constraints)

	// 1. SEQUENCE extension bit.
	if err := seq.DecodeExtensionBit(); err != nil {
		return err
	}

	// 2. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 3. periodicityAndOffset-r17: ref
	{
		if err := ie.PeriodicityAndOffset_r17.Decode(d); err != nil {
			return err
		}
	}

	// 4. numSymbols-r17: enumerated
	{
		v1, err := d.DecodeEnumerated(nRDLPRSPDCResourceSetR17NumSymbolsR17Constraints)
		if err != nil {
			return err
		}
		ie.NumSymbols_r17 = v1
	}

	// 5. dl-PRS-ResourceBandwidth-r17: integer
	{
		v2, err := d.DecodeInteger(nRDLPRSPDCResourceSetR17DlPRSResourceBandwidthR17Constraints)
		if err != nil {
			return err
		}
		ie.Dl_PRS_ResourceBandwidth_r17 = v2
	}

	// 6. dl-PRS-StartPRB-r17: integer
	{
		v3, err := d.DecodeInteger(nRDLPRSPDCResourceSetR17DlPRSStartPRBR17Constraints)
		if err != nil {
			return err
		}
		ie.Dl_PRS_StartPRB_r17 = v3
	}

	// 7. resourceList-r17: sequence-of
	{
		seqOf := d.NewSequenceOfDecoder(nRDLPRSPDCResourceSetR17ResourceListR17Constraints)
		n, err := seqOf.DecodeLength()
		if err != nil {
			return err
		}
		ie.ResourceList_r17 = make([]NR_DL_PRS_Resource_r17, n)
		for i := int64(0); i < n; i++ {
			if err := ie.ResourceList_r17[i].Decode(d); err != nil {
				return err
			}
		}
	}

	// 8. repFactorAndTimeGap-r17: ref
	{
		if seq.IsComponentPresent(5) {
			ie.RepFactorAndTimeGap_r17 = new(RepFactorAndTimeGap_r17)
			if err := ie.RepFactorAndTimeGap_r17.Decode(d); err != nil {
				return err
			}
		}
	}

	return nil
}
