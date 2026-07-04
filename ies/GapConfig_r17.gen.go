// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: GapConfig-r17 (line 9177).

var gapConfigR17Constraints = per.SequenceConstraints{
	Extensible: true,
	RootComponents: []per.ComponentInfo{
		{Name: "measGapId-r17"},
		{Name: "gapType-r17"},
		{Name: "gapOffset-r17"},
		{Name: "mgl-r17"},
		{Name: "mgrp-r17"},
		{Name: "mgta-r17"},
		{Name: "refServCellIndicator-r17", Optional: true},
		{Name: "refFR2-ServCellAsyncCA-r17", Optional: true},
		{Name: "preConfigInd-r17", Optional: true},
		{Name: "ncsgInd-r17", Optional: true},
		{Name: "gapAssociationPRS-r17", Optional: true},
		{Name: "gapSharing-r17", Optional: true},
		{Name: "gapPriority-r17", Optional: true},
	},
}

const (
	GapConfig_r17_GapType_r17_PerUE  = 0
	GapConfig_r17_GapType_r17_PerFR1 = 1
	GapConfig_r17_GapType_r17_PerFR2 = 2
)

var gapConfigR17GapTypeR17Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{GapConfig_r17_GapType_r17_PerUE, GapConfig_r17_GapType_r17_PerFR1, GapConfig_r17_GapType_r17_PerFR2},
}

var gapConfigR17GapOffsetR17Constraints = per.Constrained(0, 159)

const (
	GapConfig_r17_Mgl_r17_Ms1     = 0
	GapConfig_r17_Mgl_r17_Ms1dot5 = 1
	GapConfig_r17_Mgl_r17_Ms2     = 2
	GapConfig_r17_Mgl_r17_Ms3     = 3
	GapConfig_r17_Mgl_r17_Ms3dot5 = 4
	GapConfig_r17_Mgl_r17_Ms4     = 5
	GapConfig_r17_Mgl_r17_Ms5     = 6
	GapConfig_r17_Mgl_r17_Ms5dot5 = 7
	GapConfig_r17_Mgl_r17_Ms6     = 8
	GapConfig_r17_Mgl_r17_Ms10    = 9
	GapConfig_r17_Mgl_r17_Ms20    = 10
)

var gapConfigR17MglR17Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{GapConfig_r17_Mgl_r17_Ms1, GapConfig_r17_Mgl_r17_Ms1dot5, GapConfig_r17_Mgl_r17_Ms2, GapConfig_r17_Mgl_r17_Ms3, GapConfig_r17_Mgl_r17_Ms3dot5, GapConfig_r17_Mgl_r17_Ms4, GapConfig_r17_Mgl_r17_Ms5, GapConfig_r17_Mgl_r17_Ms5dot5, GapConfig_r17_Mgl_r17_Ms6, GapConfig_r17_Mgl_r17_Ms10, GapConfig_r17_Mgl_r17_Ms20},
}

const (
	GapConfig_r17_Mgrp_r17_Ms20  = 0
	GapConfig_r17_Mgrp_r17_Ms40  = 1
	GapConfig_r17_Mgrp_r17_Ms80  = 2
	GapConfig_r17_Mgrp_r17_Ms160 = 3
)

var gapConfigR17MgrpR17Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{GapConfig_r17_Mgrp_r17_Ms20, GapConfig_r17_Mgrp_r17_Ms40, GapConfig_r17_Mgrp_r17_Ms80, GapConfig_r17_Mgrp_r17_Ms160},
}

const (
	GapConfig_r17_Mgta_r17_Ms0      = 0
	GapConfig_r17_Mgta_r17_Ms0dot25 = 1
	GapConfig_r17_Mgta_r17_Ms0dot5  = 2
	GapConfig_r17_Mgta_r17_Ms0dot75 = 3
)

var gapConfigR17MgtaR17Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{GapConfig_r17_Mgta_r17_Ms0, GapConfig_r17_Mgta_r17_Ms0dot25, GapConfig_r17_Mgta_r17_Ms0dot5, GapConfig_r17_Mgta_r17_Ms0dot75},
}

const (
	GapConfig_r17_RefServCellIndicator_r17_PCell   = 0
	GapConfig_r17_RefServCellIndicator_r17_PSCell  = 1
	GapConfig_r17_RefServCellIndicator_r17_Mcg_FR2 = 2
)

var gapConfigR17RefServCellIndicatorR17Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{GapConfig_r17_RefServCellIndicator_r17_PCell, GapConfig_r17_RefServCellIndicator_r17_PSCell, GapConfig_r17_RefServCellIndicator_r17_Mcg_FR2},
}

const (
	GapConfig_r17_PreConfigInd_r17_True = 0
)

var gapConfigR17PreConfigIndR17Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{GapConfig_r17_PreConfigInd_r17_True},
}

const (
	GapConfig_r17_NcsgInd_r17_True = 0
)

var gapConfigR17NcsgIndR17Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{GapConfig_r17_NcsgInd_r17_True},
}

const (
	GapConfig_r17_GapAssociationPRS_r17_True = 0
)

var gapConfigR17GapAssociationPRSR17Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{GapConfig_r17_GapAssociationPRS_r17_True},
}

type GapConfig_r17 struct {
	MeasGapId_r17              MeasGapId_r17
	GapType_r17                int64
	GapOffset_r17              int64
	Mgl_r17                    int64
	Mgrp_r17                   int64
	Mgta_r17                   int64
	RefServCellIndicator_r17   *int64
	RefFR2_ServCellAsyncCA_r17 *ServCellIndex
	PreConfigInd_r17           *int64
	NcsgInd_r17                *int64
	GapAssociationPRS_r17      *int64
	GapSharing_r17             *MeasGapSharingScheme
	GapPriority_r17            *GapPriority_r17
}

func (ie *GapConfig_r17) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(gapConfigR17Constraints)

	// 1. SEQUENCE extension bit.
	if err := seq.EncodeExtensionBit(false); err != nil {
		return err
	}

	// 2. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.RefServCellIndicator_r17 != nil, ie.RefFR2_ServCellAsyncCA_r17 != nil, ie.PreConfigInd_r17 != nil, ie.NcsgInd_r17 != nil, ie.GapAssociationPRS_r17 != nil, ie.GapSharing_r17 != nil, ie.GapPriority_r17 != nil}); err != nil {
		return err
	}

	// 3. measGapId-r17: ref
	{
		if err := ie.MeasGapId_r17.Encode(e); err != nil {
			return err
		}
	}

	// 4. gapType-r17: enumerated
	{
		if err := e.EncodeEnumerated(ie.GapType_r17, gapConfigR17GapTypeR17Constraints); err != nil {
			return err
		}
	}

	// 5. gapOffset-r17: integer
	{
		if err := e.EncodeInteger(ie.GapOffset_r17, gapConfigR17GapOffsetR17Constraints); err != nil {
			return err
		}
	}

	// 6. mgl-r17: enumerated
	{
		if err := e.EncodeEnumerated(ie.Mgl_r17, gapConfigR17MglR17Constraints); err != nil {
			return err
		}
	}

	// 7. mgrp-r17: enumerated
	{
		if err := e.EncodeEnumerated(ie.Mgrp_r17, gapConfigR17MgrpR17Constraints); err != nil {
			return err
		}
	}

	// 8. mgta-r17: enumerated
	{
		if err := e.EncodeEnumerated(ie.Mgta_r17, gapConfigR17MgtaR17Constraints); err != nil {
			return err
		}
	}

	// 9. refServCellIndicator-r17: enumerated
	{
		if ie.RefServCellIndicator_r17 != nil {
			if err := e.EncodeEnumerated(*ie.RefServCellIndicator_r17, gapConfigR17RefServCellIndicatorR17Constraints); err != nil {
				return err
			}
		}
	}

	// 10. refFR2-ServCellAsyncCA-r17: ref
	{
		if ie.RefFR2_ServCellAsyncCA_r17 != nil {
			if err := ie.RefFR2_ServCellAsyncCA_r17.Encode(e); err != nil {
				return err
			}
		}
	}

	// 11. preConfigInd-r17: enumerated
	{
		if ie.PreConfigInd_r17 != nil {
			if err := e.EncodeEnumerated(*ie.PreConfigInd_r17, gapConfigR17PreConfigIndR17Constraints); err != nil {
				return err
			}
		}
	}

	// 12. ncsgInd-r17: enumerated
	{
		if ie.NcsgInd_r17 != nil {
			if err := e.EncodeEnumerated(*ie.NcsgInd_r17, gapConfigR17NcsgIndR17Constraints); err != nil {
				return err
			}
		}
	}

	// 13. gapAssociationPRS-r17: enumerated
	{
		if ie.GapAssociationPRS_r17 != nil {
			if err := e.EncodeEnumerated(*ie.GapAssociationPRS_r17, gapConfigR17GapAssociationPRSR17Constraints); err != nil {
				return err
			}
		}
	}

	// 14. gapSharing-r17: ref
	{
		if ie.GapSharing_r17 != nil {
			if err := ie.GapSharing_r17.Encode(e); err != nil {
				return err
			}
		}
	}

	// 15. gapPriority-r17: ref
	{
		if ie.GapPriority_r17 != nil {
			if err := ie.GapPriority_r17.Encode(e); err != nil {
				return err
			}
		}
	}

	return nil
}

func (ie *GapConfig_r17) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(gapConfigR17Constraints)

	// 1. SEQUENCE extension bit.
	if err := seq.DecodeExtensionBit(); err != nil {
		return err
	}

	// 2. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 3. measGapId-r17: ref
	{
		if err := ie.MeasGapId_r17.Decode(d); err != nil {
			return err
		}
	}

	// 4. gapType-r17: enumerated
	{
		v1, err := d.DecodeEnumerated(gapConfigR17GapTypeR17Constraints)
		if err != nil {
			return err
		}
		ie.GapType_r17 = v1
	}

	// 5. gapOffset-r17: integer
	{
		v2, err := d.DecodeInteger(gapConfigR17GapOffsetR17Constraints)
		if err != nil {
			return err
		}
		ie.GapOffset_r17 = v2
	}

	// 6. mgl-r17: enumerated
	{
		v3, err := d.DecodeEnumerated(gapConfigR17MglR17Constraints)
		if err != nil {
			return err
		}
		ie.Mgl_r17 = v3
	}

	// 7. mgrp-r17: enumerated
	{
		v4, err := d.DecodeEnumerated(gapConfigR17MgrpR17Constraints)
		if err != nil {
			return err
		}
		ie.Mgrp_r17 = v4
	}

	// 8. mgta-r17: enumerated
	{
		v5, err := d.DecodeEnumerated(gapConfigR17MgtaR17Constraints)
		if err != nil {
			return err
		}
		ie.Mgta_r17 = v5
	}

	// 9. refServCellIndicator-r17: enumerated
	{
		if seq.IsComponentPresent(6) {
			idx, err := d.DecodeEnumerated(gapConfigR17RefServCellIndicatorR17Constraints)
			if err != nil {
				return err
			}
			ie.RefServCellIndicator_r17 = &idx
		}
	}

	// 10. refFR2-ServCellAsyncCA-r17: ref
	{
		if seq.IsComponentPresent(7) {
			ie.RefFR2_ServCellAsyncCA_r17 = new(ServCellIndex)
			if err := ie.RefFR2_ServCellAsyncCA_r17.Decode(d); err != nil {
				return err
			}
		}
	}

	// 11. preConfigInd-r17: enumerated
	{
		if seq.IsComponentPresent(8) {
			idx, err := d.DecodeEnumerated(gapConfigR17PreConfigIndR17Constraints)
			if err != nil {
				return err
			}
			ie.PreConfigInd_r17 = &idx
		}
	}

	// 12. ncsgInd-r17: enumerated
	{
		if seq.IsComponentPresent(9) {
			idx, err := d.DecodeEnumerated(gapConfigR17NcsgIndR17Constraints)
			if err != nil {
				return err
			}
			ie.NcsgInd_r17 = &idx
		}
	}

	// 13. gapAssociationPRS-r17: enumerated
	{
		if seq.IsComponentPresent(10) {
			idx, err := d.DecodeEnumerated(gapConfigR17GapAssociationPRSR17Constraints)
			if err != nil {
				return err
			}
			ie.GapAssociationPRS_r17 = &idx
		}
	}

	// 14. gapSharing-r17: ref
	{
		if seq.IsComponentPresent(11) {
			ie.GapSharing_r17 = new(MeasGapSharingScheme)
			if err := ie.GapSharing_r17.Decode(d); err != nil {
				return err
			}
		}
	}

	// 15. gapPriority-r17: ref
	{
		if seq.IsComponentPresent(12) {
			ie.GapPriority_r17 = new(GapPriority_r17)
			if err := ie.GapPriority_r17.Decode(d); err != nil {
				return err
			}
		}
	}

	return nil
}
