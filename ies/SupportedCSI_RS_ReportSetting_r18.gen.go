// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: SupportedCSI-RS-ReportSetting-r18 (line 19340).

var supportedCSIRSReportSettingR18Constraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "maxN4-r18"},
		{Name: "maxNumberTxPortsPerResource-r18"},
		{Name: "maxNumberResourcesPerBand-r18"},
		{Name: "totalNumberTxPortsPerBand-r18"},
	},
}

const (
	SupportedCSI_RS_ReportSetting_r18_MaxN4_r18_N1 = 0
	SupportedCSI_RS_ReportSetting_r18_MaxN4_r18_N2 = 1
	SupportedCSI_RS_ReportSetting_r18_MaxN4_r18_N4 = 2
	SupportedCSI_RS_ReportSetting_r18_MaxN4_r18_N8 = 3
)

var supportedCSIRSReportSettingR18MaxN4R18Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{SupportedCSI_RS_ReportSetting_r18_MaxN4_r18_N1, SupportedCSI_RS_ReportSetting_r18_MaxN4_r18_N2, SupportedCSI_RS_ReportSetting_r18_MaxN4_r18_N4, SupportedCSI_RS_ReportSetting_r18_MaxN4_r18_N8},
}

const (
	SupportedCSI_RS_ReportSetting_r18_MaxNumberTxPortsPerResource_r18_P2  = 0
	SupportedCSI_RS_ReportSetting_r18_MaxNumberTxPortsPerResource_r18_P4  = 1
	SupportedCSI_RS_ReportSetting_r18_MaxNumberTxPortsPerResource_r18_P8  = 2
	SupportedCSI_RS_ReportSetting_r18_MaxNumberTxPortsPerResource_r18_P12 = 3
	SupportedCSI_RS_ReportSetting_r18_MaxNumberTxPortsPerResource_r18_P16 = 4
	SupportedCSI_RS_ReportSetting_r18_MaxNumberTxPortsPerResource_r18_P24 = 5
	SupportedCSI_RS_ReportSetting_r18_MaxNumberTxPortsPerResource_r18_P32 = 6
)

var supportedCSIRSReportSettingR18MaxNumberTxPortsPerResourceR18Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{SupportedCSI_RS_ReportSetting_r18_MaxNumberTxPortsPerResource_r18_P2, SupportedCSI_RS_ReportSetting_r18_MaxNumberTxPortsPerResource_r18_P4, SupportedCSI_RS_ReportSetting_r18_MaxNumberTxPortsPerResource_r18_P8, SupportedCSI_RS_ReportSetting_r18_MaxNumberTxPortsPerResource_r18_P12, SupportedCSI_RS_ReportSetting_r18_MaxNumberTxPortsPerResource_r18_P16, SupportedCSI_RS_ReportSetting_r18_MaxNumberTxPortsPerResource_r18_P24, SupportedCSI_RS_ReportSetting_r18_MaxNumberTxPortsPerResource_r18_P32},
}

var supportedCSIRSReportSettingR18MaxNumberResourcesPerBandR18Constraints = per.Constrained(1, 64)

var supportedCSIRSReportSettingR18TotalNumberTxPortsPerBandR18Constraints = per.Constrained(2, 256)

type SupportedCSI_RS_ReportSetting_r18 struct {
	MaxN4_r18                       int64
	MaxNumberTxPortsPerResource_r18 int64
	MaxNumberResourcesPerBand_r18   int64
	TotalNumberTxPortsPerBand_r18   int64
}

func (ie *SupportedCSI_RS_ReportSetting_r18) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(supportedCSIRSReportSettingR18Constraints)
	_ = seq

	// 1. maxN4-r18: enumerated
	{
		if err := e.EncodeEnumerated(ie.MaxN4_r18, supportedCSIRSReportSettingR18MaxN4R18Constraints); err != nil {
			return err
		}
	}

	// 2. maxNumberTxPortsPerResource-r18: enumerated
	{
		if err := e.EncodeEnumerated(ie.MaxNumberTxPortsPerResource_r18, supportedCSIRSReportSettingR18MaxNumberTxPortsPerResourceR18Constraints); err != nil {
			return err
		}
	}

	// 3. maxNumberResourcesPerBand-r18: integer
	{
		if err := e.EncodeInteger(ie.MaxNumberResourcesPerBand_r18, supportedCSIRSReportSettingR18MaxNumberResourcesPerBandR18Constraints); err != nil {
			return err
		}
	}

	// 4. totalNumberTxPortsPerBand-r18: integer
	{
		if err := e.EncodeInteger(ie.TotalNumberTxPortsPerBand_r18, supportedCSIRSReportSettingR18TotalNumberTxPortsPerBandR18Constraints); err != nil {
			return err
		}
	}

	return nil
}

func (ie *SupportedCSI_RS_ReportSetting_r18) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(supportedCSIRSReportSettingR18Constraints)
	_ = seq

	// 1. maxN4-r18: enumerated
	{
		v0, err := d.DecodeEnumerated(supportedCSIRSReportSettingR18MaxN4R18Constraints)
		if err != nil {
			return err
		}
		ie.MaxN4_r18 = v0
	}

	// 2. maxNumberTxPortsPerResource-r18: enumerated
	{
		v1, err := d.DecodeEnumerated(supportedCSIRSReportSettingR18MaxNumberTxPortsPerResourceR18Constraints)
		if err != nil {
			return err
		}
		ie.MaxNumberTxPortsPerResource_r18 = v1
	}

	// 3. maxNumberResourcesPerBand-r18: integer
	{
		v2, err := d.DecodeInteger(supportedCSIRSReportSettingR18MaxNumberResourcesPerBandR18Constraints)
		if err != nil {
			return err
		}
		ie.MaxNumberResourcesPerBand_r18 = v2
	}

	// 4. totalNumberTxPortsPerBand-r18: integer
	{
		v3, err := d.DecodeInteger(supportedCSIRSReportSettingR18TotalNumberTxPortsPerBandR18Constraints)
		if err != nil {
			return err
		}
		ie.TotalNumberTxPortsPerBand_r18 = v3
	}

	return nil
}
