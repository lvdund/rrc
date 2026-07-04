// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: FeatureSetDownlink-v1900 (line 19765).

var featureSetDownlinkV1900Constraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "twoTA-IntraCellBM-r19", Optional: true},
		{Name: "twoTA-InterCellBM-r19", Optional: true},
		{Name: "twoTA-RxTimeDiff-r19", Optional: true},
		{Name: "interCellCFRA-r19", Optional: true},
		{Name: "twoTA-IntraCellMultiTRP-r19", Optional: true},
		{Name: "srs-AntennaSwitching3T6R2SP-1Periodic-r19", Optional: true},
		{Name: "srs-AntennaSwitching3T3R2SP-1Periodic-r19", Optional: true},
		{Name: "lpwus-OOK-Connected-r19", Optional: true},
		{Name: "lpwus-OFDM-Connected-r19", Optional: true},
		{Name: "lpwus-FreqResourceOutsideBWP-r19", Optional: true},
	},
}

const (
	FeatureSetDownlink_v1900_TwoTA_IntraCellBM_r19_Supported = 0
)

var featureSetDownlinkV1900TwoTAIntraCellBMR19Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{FeatureSetDownlink_v1900_TwoTA_IntraCellBM_r19_Supported},
}

const (
	FeatureSetDownlink_v1900_TwoTA_InterCellBM_r19_Supported = 0
)

var featureSetDownlinkV1900TwoTAInterCellBMR19Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{FeatureSetDownlink_v1900_TwoTA_InterCellBM_r19_Supported},
}

const (
	FeatureSetDownlink_v1900_TwoTA_RxTimeDiff_r19_Supported = 0
)

var featureSetDownlinkV1900TwoTARxTimeDiffR19Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{FeatureSetDownlink_v1900_TwoTA_RxTimeDiff_r19_Supported},
}

const (
	FeatureSetDownlink_v1900_InterCellCFRA_r19_Supported = 0
)

var featureSetDownlinkV1900InterCellCFRAR19Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{FeatureSetDownlink_v1900_InterCellCFRA_r19_Supported},
}

const (
	FeatureSetDownlink_v1900_TwoTA_IntraCellMultiTRP_r19_Supported = 0
)

var featureSetDownlinkV1900TwoTAIntraCellMultiTRPR19Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{FeatureSetDownlink_v1900_TwoTA_IntraCellMultiTRP_r19_Supported},
}

const (
	FeatureSetDownlink_v1900_Srs_AntennaSwitching3T6R2SP_1Periodic_r19_Supported = 0
)

var featureSetDownlinkV1900SrsAntennaSwitching3T6R2SP1PeriodicR19Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{FeatureSetDownlink_v1900_Srs_AntennaSwitching3T6R2SP_1Periodic_r19_Supported},
}

const (
	FeatureSetDownlink_v1900_Srs_AntennaSwitching3T3R2SP_1Periodic_r19_Supported = 0
)

var featureSetDownlinkV1900SrsAntennaSwitching3T3R2SP1PeriodicR19Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{FeatureSetDownlink_v1900_Srs_AntennaSwitching3T3R2SP_1Periodic_r19_Supported},
}

const (
	FeatureSetDownlink_v1900_Lpwus_FreqResourceOutsideBWP_r19_Supported = 0
)

var featureSetDownlinkV1900LpwusFreqResourceOutsideBWPR19Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{FeatureSetDownlink_v1900_Lpwus_FreqResourceOutsideBWP_r19_Supported},
}

const (
	FeatureSetDownlink_v1900_Lpwus_OOK_Connected_r19_SupportedProc_r19_Option1_1 = 0
	FeatureSetDownlink_v1900_Lpwus_OOK_Connected_r19_SupportedProc_r19_Option1_2 = 1
	FeatureSetDownlink_v1900_Lpwus_OOK_Connected_r19_SupportedProc_r19_Both      = 2
)

var featureSetDownlinkV1900LpwusOOKConnectedR19SupportedProcR19Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{FeatureSetDownlink_v1900_Lpwus_OOK_Connected_r19_SupportedProc_r19_Option1_1, FeatureSetDownlink_v1900_Lpwus_OOK_Connected_r19_SupportedProc_r19_Option1_2, FeatureSetDownlink_v1900_Lpwus_OOK_Connected_r19_SupportedProc_r19_Both},
}

const (
	FeatureSetDownlink_v1900_Lpwus_OOK_Connected_r19_MinimumTimeGap_r19_Ms5  = 0
	FeatureSetDownlink_v1900_Lpwus_OOK_Connected_r19_MinimumTimeGap_r19_Ms13 = 1
	FeatureSetDownlink_v1900_Lpwus_OOK_Connected_r19_MinimumTimeGap_r19_Ms37 = 2
)

var featureSetDownlinkV1900LpwusOOKConnectedR19MinimumTimeGapR19Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{FeatureSetDownlink_v1900_Lpwus_OOK_Connected_r19_MinimumTimeGap_r19_Ms5, FeatureSetDownlink_v1900_Lpwus_OOK_Connected_r19_MinimumTimeGap_r19_Ms13, FeatureSetDownlink_v1900_Lpwus_OOK_Connected_r19_MinimumTimeGap_r19_Ms37},
}

const (
	FeatureSetDownlink_v1900_Lpwus_OOK_Connected_r19_MaxNumCodepointsPerMO_r19_N2 = 0
	FeatureSetDownlink_v1900_Lpwus_OOK_Connected_r19_MaxNumCodepointsPerMO_r19_N4 = 1
	FeatureSetDownlink_v1900_Lpwus_OOK_Connected_r19_MaxNumCodepointsPerMO_r19_N6 = 2
	FeatureSetDownlink_v1900_Lpwus_OOK_Connected_r19_MaxNumCodepointsPerMO_r19_N8 = 3
)

var featureSetDownlinkV1900LpwusOOKConnectedR19MaxNumCodepointsPerMOR19Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{FeatureSetDownlink_v1900_Lpwus_OOK_Connected_r19_MaxNumCodepointsPerMO_r19_N2, FeatureSetDownlink_v1900_Lpwus_OOK_Connected_r19_MaxNumCodepointsPerMO_r19_N4, FeatureSetDownlink_v1900_Lpwus_OOK_Connected_r19_MaxNumCodepointsPerMO_r19_N6, FeatureSetDownlink_v1900_Lpwus_OOK_Connected_r19_MaxNumCodepointsPerMO_r19_N8},
}

const (
	FeatureSetDownlink_v1900_Lpwus_OFDM_Connected_r19_SupportedProc_r19_Option1_1 = 0
	FeatureSetDownlink_v1900_Lpwus_OFDM_Connected_r19_SupportedProc_r19_Option1_2 = 1
	FeatureSetDownlink_v1900_Lpwus_OFDM_Connected_r19_SupportedProc_r19_Both      = 2
)

var featureSetDownlinkV1900LpwusOFDMConnectedR19SupportedProcR19Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{FeatureSetDownlink_v1900_Lpwus_OFDM_Connected_r19_SupportedProc_r19_Option1_1, FeatureSetDownlink_v1900_Lpwus_OFDM_Connected_r19_SupportedProc_r19_Option1_2, FeatureSetDownlink_v1900_Lpwus_OFDM_Connected_r19_SupportedProc_r19_Both},
}

const (
	FeatureSetDownlink_v1900_Lpwus_OFDM_Connected_r19_MinimumTimeGap_r19_Ms5  = 0
	FeatureSetDownlink_v1900_Lpwus_OFDM_Connected_r19_MinimumTimeGap_r19_Ms13 = 1
	FeatureSetDownlink_v1900_Lpwus_OFDM_Connected_r19_MinimumTimeGap_r19_Ms37 = 2
)

var featureSetDownlinkV1900LpwusOFDMConnectedR19MinimumTimeGapR19Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{FeatureSetDownlink_v1900_Lpwus_OFDM_Connected_r19_MinimumTimeGap_r19_Ms5, FeatureSetDownlink_v1900_Lpwus_OFDM_Connected_r19_MinimumTimeGap_r19_Ms13, FeatureSetDownlink_v1900_Lpwus_OFDM_Connected_r19_MinimumTimeGap_r19_Ms37},
}

const (
	FeatureSetDownlink_v1900_Lpwus_OFDM_Connected_r19_MaxNumCodepointsPerMO_r19_N2 = 0
	FeatureSetDownlink_v1900_Lpwus_OFDM_Connected_r19_MaxNumCodepointsPerMO_r19_N4 = 1
	FeatureSetDownlink_v1900_Lpwus_OFDM_Connected_r19_MaxNumCodepointsPerMO_r19_N6 = 2
	FeatureSetDownlink_v1900_Lpwus_OFDM_Connected_r19_MaxNumCodepointsPerMO_r19_N8 = 3
)

var featureSetDownlinkV1900LpwusOFDMConnectedR19MaxNumCodepointsPerMOR19Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{FeatureSetDownlink_v1900_Lpwus_OFDM_Connected_r19_MaxNumCodepointsPerMO_r19_N2, FeatureSetDownlink_v1900_Lpwus_OFDM_Connected_r19_MaxNumCodepointsPerMO_r19_N4, FeatureSetDownlink_v1900_Lpwus_OFDM_Connected_r19_MaxNumCodepointsPerMO_r19_N6, FeatureSetDownlink_v1900_Lpwus_OFDM_Connected_r19_MaxNumCodepointsPerMO_r19_N8},
}

type FeatureSetDownlink_v1900 struct {
	TwoTA_IntraCellBM_r19                     *int64
	TwoTA_InterCellBM_r19                     *int64
	TwoTA_RxTimeDiff_r19                      *int64
	InterCellCFRA_r19                         *int64
	TwoTA_IntraCellMultiTRP_r19               *int64
	Srs_AntennaSwitching3T6R2SP_1Periodic_r19 *int64
	Srs_AntennaSwitching3T3R2SP_1Periodic_r19 *int64
	Lpwus_OOK_Connected_r19                   *struct {
		SupportedProc_r19         int64
		MinimumTimeGap_r19        int64
		MaxNumCodepointsPerMO_r19 int64
	}
	Lpwus_OFDM_Connected_r19 *struct {
		SupportedProc_r19         int64
		MinimumTimeGap_r19        int64
		MaxNumCodepointsPerMO_r19 int64
	}
	Lpwus_FreqResourceOutsideBWP_r19 *int64
}

func (ie *FeatureSetDownlink_v1900) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(featureSetDownlinkV1900Constraints)

	// 1. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.TwoTA_IntraCellBM_r19 != nil, ie.TwoTA_InterCellBM_r19 != nil, ie.TwoTA_RxTimeDiff_r19 != nil, ie.InterCellCFRA_r19 != nil, ie.TwoTA_IntraCellMultiTRP_r19 != nil, ie.Srs_AntennaSwitching3T6R2SP_1Periodic_r19 != nil, ie.Srs_AntennaSwitching3T3R2SP_1Periodic_r19 != nil, ie.Lpwus_OOK_Connected_r19 != nil, ie.Lpwus_OFDM_Connected_r19 != nil, ie.Lpwus_FreqResourceOutsideBWP_r19 != nil}); err != nil {
		return err
	}

	// 2. twoTA-IntraCellBM-r19: enumerated
	{
		if ie.TwoTA_IntraCellBM_r19 != nil {
			if err := e.EncodeEnumerated(*ie.TwoTA_IntraCellBM_r19, featureSetDownlinkV1900TwoTAIntraCellBMR19Constraints); err != nil {
				return err
			}
		}
	}

	// 3. twoTA-InterCellBM-r19: enumerated
	{
		if ie.TwoTA_InterCellBM_r19 != nil {
			if err := e.EncodeEnumerated(*ie.TwoTA_InterCellBM_r19, featureSetDownlinkV1900TwoTAInterCellBMR19Constraints); err != nil {
				return err
			}
		}
	}

	// 4. twoTA-RxTimeDiff-r19: enumerated
	{
		if ie.TwoTA_RxTimeDiff_r19 != nil {
			if err := e.EncodeEnumerated(*ie.TwoTA_RxTimeDiff_r19, featureSetDownlinkV1900TwoTARxTimeDiffR19Constraints); err != nil {
				return err
			}
		}
	}

	// 5. interCellCFRA-r19: enumerated
	{
		if ie.InterCellCFRA_r19 != nil {
			if err := e.EncodeEnumerated(*ie.InterCellCFRA_r19, featureSetDownlinkV1900InterCellCFRAR19Constraints); err != nil {
				return err
			}
		}
	}

	// 6. twoTA-IntraCellMultiTRP-r19: enumerated
	{
		if ie.TwoTA_IntraCellMultiTRP_r19 != nil {
			if err := e.EncodeEnumerated(*ie.TwoTA_IntraCellMultiTRP_r19, featureSetDownlinkV1900TwoTAIntraCellMultiTRPR19Constraints); err != nil {
				return err
			}
		}
	}

	// 7. srs-AntennaSwitching3T6R2SP-1Periodic-r19: enumerated
	{
		if ie.Srs_AntennaSwitching3T6R2SP_1Periodic_r19 != nil {
			if err := e.EncodeEnumerated(*ie.Srs_AntennaSwitching3T6R2SP_1Periodic_r19, featureSetDownlinkV1900SrsAntennaSwitching3T6R2SP1PeriodicR19Constraints); err != nil {
				return err
			}
		}
	}

	// 8. srs-AntennaSwitching3T3R2SP-1Periodic-r19: enumerated
	{
		if ie.Srs_AntennaSwitching3T3R2SP_1Periodic_r19 != nil {
			if err := e.EncodeEnumerated(*ie.Srs_AntennaSwitching3T3R2SP_1Periodic_r19, featureSetDownlinkV1900SrsAntennaSwitching3T3R2SP1PeriodicR19Constraints); err != nil {
				return err
			}
		}
	}

	// 9. lpwus-OOK-Connected-r19: sequence
	{
		if ie.Lpwus_OOK_Connected_r19 != nil {
			c := ie.Lpwus_OOK_Connected_r19
			if err := e.EncodeEnumerated(c.SupportedProc_r19, featureSetDownlinkV1900LpwusOOKConnectedR19SupportedProcR19Constraints); err != nil {
				return err
			}
			if err := e.EncodeEnumerated(c.MinimumTimeGap_r19, featureSetDownlinkV1900LpwusOOKConnectedR19MinimumTimeGapR19Constraints); err != nil {
				return err
			}
			if err := e.EncodeEnumerated(c.MaxNumCodepointsPerMO_r19, featureSetDownlinkV1900LpwusOOKConnectedR19MaxNumCodepointsPerMOR19Constraints); err != nil {
				return err
			}
		}
	}

	// 10. lpwus-OFDM-Connected-r19: sequence
	{
		if ie.Lpwus_OFDM_Connected_r19 != nil {
			c := ie.Lpwus_OFDM_Connected_r19
			if err := e.EncodeEnumerated(c.SupportedProc_r19, featureSetDownlinkV1900LpwusOFDMConnectedR19SupportedProcR19Constraints); err != nil {
				return err
			}
			if err := e.EncodeEnumerated(c.MinimumTimeGap_r19, featureSetDownlinkV1900LpwusOFDMConnectedR19MinimumTimeGapR19Constraints); err != nil {
				return err
			}
			if err := e.EncodeEnumerated(c.MaxNumCodepointsPerMO_r19, featureSetDownlinkV1900LpwusOFDMConnectedR19MaxNumCodepointsPerMOR19Constraints); err != nil {
				return err
			}
		}
	}

	// 11. lpwus-FreqResourceOutsideBWP-r19: enumerated
	{
		if ie.Lpwus_FreqResourceOutsideBWP_r19 != nil {
			if err := e.EncodeEnumerated(*ie.Lpwus_FreqResourceOutsideBWP_r19, featureSetDownlinkV1900LpwusFreqResourceOutsideBWPR19Constraints); err != nil {
				return err
			}
		}
	}

	return nil
}

func (ie *FeatureSetDownlink_v1900) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(featureSetDownlinkV1900Constraints)

	// 1. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 2. twoTA-IntraCellBM-r19: enumerated
	{
		if seq.IsComponentPresent(0) {
			idx, err := d.DecodeEnumerated(featureSetDownlinkV1900TwoTAIntraCellBMR19Constraints)
			if err != nil {
				return err
			}
			ie.TwoTA_IntraCellBM_r19 = &idx
		}
	}

	// 3. twoTA-InterCellBM-r19: enumerated
	{
		if seq.IsComponentPresent(1) {
			idx, err := d.DecodeEnumerated(featureSetDownlinkV1900TwoTAInterCellBMR19Constraints)
			if err != nil {
				return err
			}
			ie.TwoTA_InterCellBM_r19 = &idx
		}
	}

	// 4. twoTA-RxTimeDiff-r19: enumerated
	{
		if seq.IsComponentPresent(2) {
			idx, err := d.DecodeEnumerated(featureSetDownlinkV1900TwoTARxTimeDiffR19Constraints)
			if err != nil {
				return err
			}
			ie.TwoTA_RxTimeDiff_r19 = &idx
		}
	}

	// 5. interCellCFRA-r19: enumerated
	{
		if seq.IsComponentPresent(3) {
			idx, err := d.DecodeEnumerated(featureSetDownlinkV1900InterCellCFRAR19Constraints)
			if err != nil {
				return err
			}
			ie.InterCellCFRA_r19 = &idx
		}
	}

	// 6. twoTA-IntraCellMultiTRP-r19: enumerated
	{
		if seq.IsComponentPresent(4) {
			idx, err := d.DecodeEnumerated(featureSetDownlinkV1900TwoTAIntraCellMultiTRPR19Constraints)
			if err != nil {
				return err
			}
			ie.TwoTA_IntraCellMultiTRP_r19 = &idx
		}
	}

	// 7. srs-AntennaSwitching3T6R2SP-1Periodic-r19: enumerated
	{
		if seq.IsComponentPresent(5) {
			idx, err := d.DecodeEnumerated(featureSetDownlinkV1900SrsAntennaSwitching3T6R2SP1PeriodicR19Constraints)
			if err != nil {
				return err
			}
			ie.Srs_AntennaSwitching3T6R2SP_1Periodic_r19 = &idx
		}
	}

	// 8. srs-AntennaSwitching3T3R2SP-1Periodic-r19: enumerated
	{
		if seq.IsComponentPresent(6) {
			idx, err := d.DecodeEnumerated(featureSetDownlinkV1900SrsAntennaSwitching3T3R2SP1PeriodicR19Constraints)
			if err != nil {
				return err
			}
			ie.Srs_AntennaSwitching3T3R2SP_1Periodic_r19 = &idx
		}
	}

	// 9. lpwus-OOK-Connected-r19: sequence
	{
		if seq.IsComponentPresent(7) {
			ie.Lpwus_OOK_Connected_r19 = &struct {
				SupportedProc_r19         int64
				MinimumTimeGap_r19        int64
				MaxNumCodepointsPerMO_r19 int64
			}{}
			c := ie.Lpwus_OOK_Connected_r19
			{
				v, err := d.DecodeEnumerated(featureSetDownlinkV1900LpwusOOKConnectedR19SupportedProcR19Constraints)
				if err != nil {
					return err
				}
				c.SupportedProc_r19 = v
			}
			{
				v, err := d.DecodeEnumerated(featureSetDownlinkV1900LpwusOOKConnectedR19MinimumTimeGapR19Constraints)
				if err != nil {
					return err
				}
				c.MinimumTimeGap_r19 = v
			}
			{
				v, err := d.DecodeEnumerated(featureSetDownlinkV1900LpwusOOKConnectedR19MaxNumCodepointsPerMOR19Constraints)
				if err != nil {
					return err
				}
				c.MaxNumCodepointsPerMO_r19 = v
			}
		}
	}

	// 10. lpwus-OFDM-Connected-r19: sequence
	{
		if seq.IsComponentPresent(8) {
			ie.Lpwus_OFDM_Connected_r19 = &struct {
				SupportedProc_r19         int64
				MinimumTimeGap_r19        int64
				MaxNumCodepointsPerMO_r19 int64
			}{}
			c := ie.Lpwus_OFDM_Connected_r19
			{
				v, err := d.DecodeEnumerated(featureSetDownlinkV1900LpwusOFDMConnectedR19SupportedProcR19Constraints)
				if err != nil {
					return err
				}
				c.SupportedProc_r19 = v
			}
			{
				v, err := d.DecodeEnumerated(featureSetDownlinkV1900LpwusOFDMConnectedR19MinimumTimeGapR19Constraints)
				if err != nil {
					return err
				}
				c.MinimumTimeGap_r19 = v
			}
			{
				v, err := d.DecodeEnumerated(featureSetDownlinkV1900LpwusOFDMConnectedR19MaxNumCodepointsPerMOR19Constraints)
				if err != nil {
					return err
				}
				c.MaxNumCodepointsPerMO_r19 = v
			}
		}
	}

	// 11. lpwus-FreqResourceOutsideBWP-r19: enumerated
	{
		if seq.IsComponentPresent(9) {
			idx, err := d.DecodeEnumerated(featureSetDownlinkV1900LpwusFreqResourceOutsideBWPR19Constraints)
			if err != nil {
				return err
			}
			ie.Lpwus_FreqResourceOutsideBWP_r19 = &idx
		}
	}

	return nil
}
