// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/rrc/common"
)

// Source: ReportSFTD-NR (line 13551).

var reportSFTDNRConstraints = per.SequenceConstraints{
	Extensible: true,
	RootComponents: []per.ComponentInfo{
		{Name: "reportSFTD-Meas"},
		{Name: "reportRSRP"},
	},
	ExtComponents: []per.ComponentInfo{
		{Name: "extension-addition-0", Optional: true},
	},
}

const (
	ReportSFTD_NR_Ext_ReportSFTD_NeighMeas_True = 0
)

var reportSFTDNRExtReportSFTDNeighMeasConstraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{ReportSFTD_NR_Ext_ReportSFTD_NeighMeas_True},
}

const (
	ReportSFTD_NR_Ext_Drx_SFTD_NeighMeas_True = 0
)

var reportSFTDNRExtDrxSFTDNeighMeasConstraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{ReportSFTD_NR_Ext_Drx_SFTD_NeighMeas_True},
}

var reportSFTDNRExtCellsForWhichToReportSFTDConstraints = per.SizeRange(1, common.MaxCellSFTD)

type ReportSFTD_NR struct {
	ReportSFTD_Meas           bool
	ReportRSRP                bool
	ReportSFTD_NeighMeas      *int64
	Drx_SFTD_NeighMeas        *int64
	CellsForWhichToReportSFTD []PhysCellId
}

func (ie *ReportSFTD_NR) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(reportSFTDNRConstraints)

	// Extension-group presence (one var per [[...]] group).
	hasExtGroup0 := ie.ReportSFTD_NeighMeas != nil || ie.Drx_SFTD_NeighMeas != nil || ie.CellsForWhichToReportSFTD != nil
	hasExtensions := hasExtGroup0

	// 1. SEQUENCE extension bit.
	if err := seq.EncodeExtensionBit(hasExtensions); err != nil {
		return err
	}

	// 2. reportSFTD-Meas: boolean
	{
		if err := e.EncodeBoolean(ie.ReportSFTD_Meas); err != nil {
			return err
		}
	}

	// 3. reportRSRP: boolean
	{
		if err := e.EncodeBoolean(ie.ReportRSRP); err != nil {
			return err
		}
	}

	if hasExtensions {
		extPresence := []bool{hasExtGroup0}
		var extValues [][]byte

		// Extension group 0:
		if hasExtGroup0 {
			ex := per.NewEncoder(e.Variant())
			groupSeq := ex.NewSequenceEncoder(per.SequenceConstraints{
				Extensible: false,
				RootComponents: []per.ComponentInfo{
					{Name: "reportSFTD-NeighMeas", Optional: true},
					{Name: "drx-SFTD-NeighMeas", Optional: true},
					{Name: "cellsForWhichToReportSFTD", Optional: true},
				},
			})
			if err := groupSeq.EncodePreamble([]bool{ie.ReportSFTD_NeighMeas != nil, ie.Drx_SFTD_NeighMeas != nil, ie.CellsForWhichToReportSFTD != nil}); err != nil {
				return err
			}

			if ie.ReportSFTD_NeighMeas != nil {
				if err := ex.EncodeEnumerated(*ie.ReportSFTD_NeighMeas, reportSFTDNRExtReportSFTDNeighMeasConstraints); err != nil {
					return err
				}
			}

			if ie.Drx_SFTD_NeighMeas != nil {
				if err := ex.EncodeEnumerated(*ie.Drx_SFTD_NeighMeas, reportSFTDNRExtDrxSFTDNeighMeasConstraints); err != nil {
					return err
				}
			}

			if ie.CellsForWhichToReportSFTD != nil {
				seqOf := ex.NewSequenceOfEncoder(reportSFTDNRExtCellsForWhichToReportSFTDConstraints)
				if err := seqOf.EncodeLength(int64(len(ie.CellsForWhichToReportSFTD))); err != nil {
					return err
				}
				for i := range ie.CellsForWhichToReportSFTD {
					if err := ie.CellsForWhichToReportSFTD[i].Encode(ex); err != nil {
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

func (ie *ReportSFTD_NR) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(reportSFTDNRConstraints)

	// 1. SEQUENCE extension bit.
	if err := seq.DecodeExtensionBit(); err != nil {
		return err
	}

	// 2. reportSFTD-Meas: boolean
	{
		v0, err := d.DecodeBoolean()
		if err != nil {
			return err
		}
		ie.ReportSFTD_Meas = v0
	}

	// 3. reportRSRP: boolean
	{
		v1, err := d.DecodeBoolean()
		if err != nil {
			return err
		}
		ie.ReportRSRP = v1
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
				{Name: "reportSFTD-NeighMeas", Optional: true},
				{Name: "drx-SFTD-NeighMeas", Optional: true},
				{Name: "cellsForWhichToReportSFTD", Optional: true},
			},
		})
		if err := groupSeq.DecodePreamble(); err != nil {
			return err
		}

		if groupSeq.IsComponentPresent(0) {
			v, err := dx.DecodeEnumerated(reportSFTDNRExtReportSFTDNeighMeasConstraints)
			if err != nil {
				return err
			}
			ie.ReportSFTD_NeighMeas = &v
		}

		if groupSeq.IsComponentPresent(1) {
			v, err := dx.DecodeEnumerated(reportSFTDNRExtDrxSFTDNeighMeasConstraints)
			if err != nil {
				return err
			}
			ie.Drx_SFTD_NeighMeas = &v
		}

		if groupSeq.IsComponentPresent(2) {
			seqOf := dx.NewSequenceOfDecoder(reportSFTDNRExtCellsForWhichToReportSFTDConstraints)
			n, err := seqOf.DecodeLength()
			if err != nil {
				return err
			}
			ie.CellsForWhichToReportSFTD = make([]PhysCellId, n)
			for i := int64(0); i < n; i++ {
				if err := ie.CellsForWhichToReportSFTD[i].Decode(dx); err != nil {
					return err
				}
			}
		}
	}

	return nil
}
