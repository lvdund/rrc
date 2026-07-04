// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: UE-BasedPerfMeas-Parameters-r16 (line 25430).

var uEBasedPerfMeasParametersR16Constraints = per.SequenceConstraints{
	Extensible: true,
	RootComponents: []per.ComponentInfo{
		{Name: "barometerMeasReport-r16", Optional: true},
		{Name: "immMeasBT-r16", Optional: true},
		{Name: "immMeasWLAN-r16", Optional: true},
		{Name: "loggedMeasBT-r16", Optional: true},
		{Name: "loggedMeasurements-r16", Optional: true},
		{Name: "loggedMeasWLAN-r16", Optional: true},
		{Name: "orientationMeasReport-r16", Optional: true},
		{Name: "speedMeasReport-r16", Optional: true},
		{Name: "gnss-Location-r16", Optional: true},
		{Name: "ulPDCP-Delay-r16", Optional: true},
	},
	ExtComponents: []per.ComponentInfo{
		{Name: "extension-addition-0", Optional: true},
		{Name: "extension-addition-1", Optional: true},
		{Name: "extension-addition-2", Optional: true},
	},
}

const (
	UE_BasedPerfMeas_Parameters_r16_BarometerMeasReport_r16_Supported = 0
)

var uEBasedPerfMeasParametersR16BarometerMeasReportR16Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{UE_BasedPerfMeas_Parameters_r16_BarometerMeasReport_r16_Supported},
}

const (
	UE_BasedPerfMeas_Parameters_r16_ImmMeasBT_r16_Supported = 0
)

var uEBasedPerfMeasParametersR16ImmMeasBTR16Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{UE_BasedPerfMeas_Parameters_r16_ImmMeasBT_r16_Supported},
}

const (
	UE_BasedPerfMeas_Parameters_r16_ImmMeasWLAN_r16_Supported = 0
)

var uEBasedPerfMeasParametersR16ImmMeasWLANR16Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{UE_BasedPerfMeas_Parameters_r16_ImmMeasWLAN_r16_Supported},
}

const (
	UE_BasedPerfMeas_Parameters_r16_LoggedMeasBT_r16_Supported = 0
)

var uEBasedPerfMeasParametersR16LoggedMeasBTR16Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{UE_BasedPerfMeas_Parameters_r16_LoggedMeasBT_r16_Supported},
}

const (
	UE_BasedPerfMeas_Parameters_r16_LoggedMeasurements_r16_Supported = 0
)

var uEBasedPerfMeasParametersR16LoggedMeasurementsR16Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{UE_BasedPerfMeas_Parameters_r16_LoggedMeasurements_r16_Supported},
}

const (
	UE_BasedPerfMeas_Parameters_r16_LoggedMeasWLAN_r16_Supported = 0
)

var uEBasedPerfMeasParametersR16LoggedMeasWLANR16Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{UE_BasedPerfMeas_Parameters_r16_LoggedMeasWLAN_r16_Supported},
}

const (
	UE_BasedPerfMeas_Parameters_r16_OrientationMeasReport_r16_Supported = 0
)

var uEBasedPerfMeasParametersR16OrientationMeasReportR16Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{UE_BasedPerfMeas_Parameters_r16_OrientationMeasReport_r16_Supported},
}

const (
	UE_BasedPerfMeas_Parameters_r16_SpeedMeasReport_r16_Supported = 0
)

var uEBasedPerfMeasParametersR16SpeedMeasReportR16Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{UE_BasedPerfMeas_Parameters_r16_SpeedMeasReport_r16_Supported},
}

const (
	UE_BasedPerfMeas_Parameters_r16_Gnss_Location_r16_Supported = 0
)

var uEBasedPerfMeasParametersR16GnssLocationR16Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{UE_BasedPerfMeas_Parameters_r16_Gnss_Location_r16_Supported},
}

const (
	UE_BasedPerfMeas_Parameters_r16_UlPDCP_Delay_r16_Supported = 0
)

var uEBasedPerfMeasParametersR16UlPDCPDelayR16Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{UE_BasedPerfMeas_Parameters_r16_UlPDCP_Delay_r16_Supported},
}

const (
	UE_BasedPerfMeas_Parameters_r16_Ext_SigBasedLogMDT_OverrideProtect_r17_Supported = 0
)

var uEBasedPerfMeasParametersR16ExtSigBasedLogMDTOverrideProtectR17Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{UE_BasedPerfMeas_Parameters_r16_Ext_SigBasedLogMDT_OverrideProtect_r17_Supported},
}

const (
	UE_BasedPerfMeas_Parameters_r16_Ext_MultipleCEF_Report_r17_Supported = 0
)

var uEBasedPerfMeasParametersR16ExtMultipleCEFReportR17Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{UE_BasedPerfMeas_Parameters_r16_Ext_MultipleCEF_Report_r17_Supported},
}

const (
	UE_BasedPerfMeas_Parameters_r16_Ext_ExcessPacketDelay_r17_Supported = 0
)

var uEBasedPerfMeasParametersR16ExtExcessPacketDelayR17Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{UE_BasedPerfMeas_Parameters_r16_Ext_ExcessPacketDelay_r17_Supported},
}

const (
	UE_BasedPerfMeas_Parameters_r16_Ext_EarlyMeasLog_r17_Supported = 0
)

var uEBasedPerfMeasParametersR16ExtEarlyMeasLogR17Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{UE_BasedPerfMeas_Parameters_r16_Ext_EarlyMeasLog_r17_Supported},
}

const (
	UE_BasedPerfMeas_Parameters_r16_Ext_LoggedMDT_PNI_NPN_r18_Supported = 0
)

var uEBasedPerfMeasParametersR16ExtLoggedMDTPNINPNR18Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{UE_BasedPerfMeas_Parameters_r16_Ext_LoggedMDT_PNI_NPN_r18_Supported},
}

const (
	UE_BasedPerfMeas_Parameters_r16_Ext_LoggedMDT_SNPN_r18_Supported = 0
)

var uEBasedPerfMeasParametersR16ExtLoggedMDTSNPNR18Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{UE_BasedPerfMeas_Parameters_r16_Ext_LoggedMDT_SNPN_r18_Supported},
}

const (
	UE_BasedPerfMeas_Parameters_r16_Ext_GeoAreaScopeChecking_r19_Supported = 0
)

var uEBasedPerfMeasParametersR16ExtGeoAreaScopeCheckingR19Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{UE_BasedPerfMeas_Parameters_r16_Ext_GeoAreaScopeChecking_r19_Supported},
}

const (
	UE_BasedPerfMeas_Parameters_r16_Ext_LoggedMDT_Slicing_r19_Supported = 0
)

var uEBasedPerfMeasParametersR16ExtLoggedMDTSlicingR19Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{UE_BasedPerfMeas_Parameters_r16_Ext_LoggedMDT_Slicing_r19_Supported},
}

type UE_BasedPerfMeas_Parameters_r16 struct {
	BarometerMeasReport_r16            *int64
	ImmMeasBT_r16                      *int64
	ImmMeasWLAN_r16                    *int64
	LoggedMeasBT_r16                   *int64
	LoggedMeasurements_r16             *int64
	LoggedMeasWLAN_r16                 *int64
	OrientationMeasReport_r16          *int64
	SpeedMeasReport_r16                *int64
	Gnss_Location_r16                  *int64
	UlPDCP_Delay_r16                   *int64
	SigBasedLogMDT_OverrideProtect_r17 *int64
	MultipleCEF_Report_r17             *int64
	ExcessPacketDelay_r17              *int64
	EarlyMeasLog_r17                   *int64
	LoggedMDT_PNI_NPN_r18              *int64
	LoggedMDT_SNPN_r18                 *int64
	GeoAreaScopeChecking_r19           *int64
	LoggedMDT_Slicing_r19              *int64
}

func (ie *UE_BasedPerfMeas_Parameters_r16) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(uEBasedPerfMeasParametersR16Constraints)

	// Extension-group presence (one var per [[...]] group).
	hasExtGroup0 := ie.SigBasedLogMDT_OverrideProtect_r17 != nil || ie.MultipleCEF_Report_r17 != nil || ie.ExcessPacketDelay_r17 != nil || ie.EarlyMeasLog_r17 != nil
	hasExtGroup1 := ie.LoggedMDT_PNI_NPN_r18 != nil || ie.LoggedMDT_SNPN_r18 != nil
	hasExtGroup2 := ie.GeoAreaScopeChecking_r19 != nil || ie.LoggedMDT_Slicing_r19 != nil
	hasExtensions := hasExtGroup0 || hasExtGroup1 || hasExtGroup2

	// 1. SEQUENCE extension bit.
	if err := seq.EncodeExtensionBit(hasExtensions); err != nil {
		return err
	}

	// 2. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.BarometerMeasReport_r16 != nil, ie.ImmMeasBT_r16 != nil, ie.ImmMeasWLAN_r16 != nil, ie.LoggedMeasBT_r16 != nil, ie.LoggedMeasurements_r16 != nil, ie.LoggedMeasWLAN_r16 != nil, ie.OrientationMeasReport_r16 != nil, ie.SpeedMeasReport_r16 != nil, ie.Gnss_Location_r16 != nil, ie.UlPDCP_Delay_r16 != nil}); err != nil {
		return err
	}

	// 3. barometerMeasReport-r16: enumerated
	{
		if ie.BarometerMeasReport_r16 != nil {
			if err := e.EncodeEnumerated(*ie.BarometerMeasReport_r16, uEBasedPerfMeasParametersR16BarometerMeasReportR16Constraints); err != nil {
				return err
			}
		}
	}

	// 4. immMeasBT-r16: enumerated
	{
		if ie.ImmMeasBT_r16 != nil {
			if err := e.EncodeEnumerated(*ie.ImmMeasBT_r16, uEBasedPerfMeasParametersR16ImmMeasBTR16Constraints); err != nil {
				return err
			}
		}
	}

	// 5. immMeasWLAN-r16: enumerated
	{
		if ie.ImmMeasWLAN_r16 != nil {
			if err := e.EncodeEnumerated(*ie.ImmMeasWLAN_r16, uEBasedPerfMeasParametersR16ImmMeasWLANR16Constraints); err != nil {
				return err
			}
		}
	}

	// 6. loggedMeasBT-r16: enumerated
	{
		if ie.LoggedMeasBT_r16 != nil {
			if err := e.EncodeEnumerated(*ie.LoggedMeasBT_r16, uEBasedPerfMeasParametersR16LoggedMeasBTR16Constraints); err != nil {
				return err
			}
		}
	}

	// 7. loggedMeasurements-r16: enumerated
	{
		if ie.LoggedMeasurements_r16 != nil {
			if err := e.EncodeEnumerated(*ie.LoggedMeasurements_r16, uEBasedPerfMeasParametersR16LoggedMeasurementsR16Constraints); err != nil {
				return err
			}
		}
	}

	// 8. loggedMeasWLAN-r16: enumerated
	{
		if ie.LoggedMeasWLAN_r16 != nil {
			if err := e.EncodeEnumerated(*ie.LoggedMeasWLAN_r16, uEBasedPerfMeasParametersR16LoggedMeasWLANR16Constraints); err != nil {
				return err
			}
		}
	}

	// 9. orientationMeasReport-r16: enumerated
	{
		if ie.OrientationMeasReport_r16 != nil {
			if err := e.EncodeEnumerated(*ie.OrientationMeasReport_r16, uEBasedPerfMeasParametersR16OrientationMeasReportR16Constraints); err != nil {
				return err
			}
		}
	}

	// 10. speedMeasReport-r16: enumerated
	{
		if ie.SpeedMeasReport_r16 != nil {
			if err := e.EncodeEnumerated(*ie.SpeedMeasReport_r16, uEBasedPerfMeasParametersR16SpeedMeasReportR16Constraints); err != nil {
				return err
			}
		}
	}

	// 11. gnss-Location-r16: enumerated
	{
		if ie.Gnss_Location_r16 != nil {
			if err := e.EncodeEnumerated(*ie.Gnss_Location_r16, uEBasedPerfMeasParametersR16GnssLocationR16Constraints); err != nil {
				return err
			}
		}
	}

	// 12. ulPDCP-Delay-r16: enumerated
	{
		if ie.UlPDCP_Delay_r16 != nil {
			if err := e.EncodeEnumerated(*ie.UlPDCP_Delay_r16, uEBasedPerfMeasParametersR16UlPDCPDelayR16Constraints); err != nil {
				return err
			}
		}
	}

	if hasExtensions {
		extPresence := []bool{hasExtGroup0, hasExtGroup1, hasExtGroup2}
		var extValues [][]byte

		// Extension group 0:
		if hasExtGroup0 {
			ex := per.NewEncoder(e.Variant())
			groupSeq := ex.NewSequenceEncoder(per.SequenceConstraints{
				Extensible: false,
				RootComponents: []per.ComponentInfo{
					{Name: "sigBasedLogMDT-OverrideProtect-r17", Optional: true},
					{Name: "multipleCEF-Report-r17", Optional: true},
					{Name: "excessPacketDelay-r17", Optional: true},
					{Name: "earlyMeasLog-r17", Optional: true},
				},
			})
			if err := groupSeq.EncodePreamble([]bool{ie.SigBasedLogMDT_OverrideProtect_r17 != nil, ie.MultipleCEF_Report_r17 != nil, ie.ExcessPacketDelay_r17 != nil, ie.EarlyMeasLog_r17 != nil}); err != nil {
				return err
			}

			if ie.SigBasedLogMDT_OverrideProtect_r17 != nil {
				if err := ex.EncodeEnumerated(*ie.SigBasedLogMDT_OverrideProtect_r17, uEBasedPerfMeasParametersR16ExtSigBasedLogMDTOverrideProtectR17Constraints); err != nil {
					return err
				}
			}

			if ie.MultipleCEF_Report_r17 != nil {
				if err := ex.EncodeEnumerated(*ie.MultipleCEF_Report_r17, uEBasedPerfMeasParametersR16ExtMultipleCEFReportR17Constraints); err != nil {
					return err
				}
			}

			if ie.ExcessPacketDelay_r17 != nil {
				if err := ex.EncodeEnumerated(*ie.ExcessPacketDelay_r17, uEBasedPerfMeasParametersR16ExtExcessPacketDelayR17Constraints); err != nil {
					return err
				}
			}

			if ie.EarlyMeasLog_r17 != nil {
				if err := ex.EncodeEnumerated(*ie.EarlyMeasLog_r17, uEBasedPerfMeasParametersR16ExtEarlyMeasLogR17Constraints); err != nil {
					return err
				}
			}
			extValues = append(extValues, ex.Bytes())
		}

		// Extension group 1:
		if hasExtGroup1 {
			ex := per.NewEncoder(e.Variant())
			groupSeq := ex.NewSequenceEncoder(per.SequenceConstraints{
				Extensible: false,
				RootComponents: []per.ComponentInfo{
					{Name: "loggedMDT-PNI-NPN-r18", Optional: true},
					{Name: "loggedMDT-SNPN-r18", Optional: true},
				},
			})
			if err := groupSeq.EncodePreamble([]bool{ie.LoggedMDT_PNI_NPN_r18 != nil, ie.LoggedMDT_SNPN_r18 != nil}); err != nil {
				return err
			}

			if ie.LoggedMDT_PNI_NPN_r18 != nil {
				if err := ex.EncodeEnumerated(*ie.LoggedMDT_PNI_NPN_r18, uEBasedPerfMeasParametersR16ExtLoggedMDTPNINPNR18Constraints); err != nil {
					return err
				}
			}

			if ie.LoggedMDT_SNPN_r18 != nil {
				if err := ex.EncodeEnumerated(*ie.LoggedMDT_SNPN_r18, uEBasedPerfMeasParametersR16ExtLoggedMDTSNPNR18Constraints); err != nil {
					return err
				}
			}
			extValues = append(extValues, ex.Bytes())
		}

		// Extension group 2:
		if hasExtGroup2 {
			ex := per.NewEncoder(e.Variant())
			groupSeq := ex.NewSequenceEncoder(per.SequenceConstraints{
				Extensible: false,
				RootComponents: []per.ComponentInfo{
					{Name: "geoAreaScopeChecking-r19", Optional: true},
					{Name: "loggedMDT-Slicing-r19", Optional: true},
				},
			})
			if err := groupSeq.EncodePreamble([]bool{ie.GeoAreaScopeChecking_r19 != nil, ie.LoggedMDT_Slicing_r19 != nil}); err != nil {
				return err
			}

			if ie.GeoAreaScopeChecking_r19 != nil {
				if err := ex.EncodeEnumerated(*ie.GeoAreaScopeChecking_r19, uEBasedPerfMeasParametersR16ExtGeoAreaScopeCheckingR19Constraints); err != nil {
					return err
				}
			}

			if ie.LoggedMDT_Slicing_r19 != nil {
				if err := ex.EncodeEnumerated(*ie.LoggedMDT_Slicing_r19, uEBasedPerfMeasParametersR16ExtLoggedMDTSlicingR19Constraints); err != nil {
					return err
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

func (ie *UE_BasedPerfMeas_Parameters_r16) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(uEBasedPerfMeasParametersR16Constraints)

	// 1. SEQUENCE extension bit.
	if err := seq.DecodeExtensionBit(); err != nil {
		return err
	}

	// 2. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 3. barometerMeasReport-r16: enumerated
	{
		if seq.IsComponentPresent(0) {
			idx, err := d.DecodeEnumerated(uEBasedPerfMeasParametersR16BarometerMeasReportR16Constraints)
			if err != nil {
				return err
			}
			ie.BarometerMeasReport_r16 = &idx
		}
	}

	// 4. immMeasBT-r16: enumerated
	{
		if seq.IsComponentPresent(1) {
			idx, err := d.DecodeEnumerated(uEBasedPerfMeasParametersR16ImmMeasBTR16Constraints)
			if err != nil {
				return err
			}
			ie.ImmMeasBT_r16 = &idx
		}
	}

	// 5. immMeasWLAN-r16: enumerated
	{
		if seq.IsComponentPresent(2) {
			idx, err := d.DecodeEnumerated(uEBasedPerfMeasParametersR16ImmMeasWLANR16Constraints)
			if err != nil {
				return err
			}
			ie.ImmMeasWLAN_r16 = &idx
		}
	}

	// 6. loggedMeasBT-r16: enumerated
	{
		if seq.IsComponentPresent(3) {
			idx, err := d.DecodeEnumerated(uEBasedPerfMeasParametersR16LoggedMeasBTR16Constraints)
			if err != nil {
				return err
			}
			ie.LoggedMeasBT_r16 = &idx
		}
	}

	// 7. loggedMeasurements-r16: enumerated
	{
		if seq.IsComponentPresent(4) {
			idx, err := d.DecodeEnumerated(uEBasedPerfMeasParametersR16LoggedMeasurementsR16Constraints)
			if err != nil {
				return err
			}
			ie.LoggedMeasurements_r16 = &idx
		}
	}

	// 8. loggedMeasWLAN-r16: enumerated
	{
		if seq.IsComponentPresent(5) {
			idx, err := d.DecodeEnumerated(uEBasedPerfMeasParametersR16LoggedMeasWLANR16Constraints)
			if err != nil {
				return err
			}
			ie.LoggedMeasWLAN_r16 = &idx
		}
	}

	// 9. orientationMeasReport-r16: enumerated
	{
		if seq.IsComponentPresent(6) {
			idx, err := d.DecodeEnumerated(uEBasedPerfMeasParametersR16OrientationMeasReportR16Constraints)
			if err != nil {
				return err
			}
			ie.OrientationMeasReport_r16 = &idx
		}
	}

	// 10. speedMeasReport-r16: enumerated
	{
		if seq.IsComponentPresent(7) {
			idx, err := d.DecodeEnumerated(uEBasedPerfMeasParametersR16SpeedMeasReportR16Constraints)
			if err != nil {
				return err
			}
			ie.SpeedMeasReport_r16 = &idx
		}
	}

	// 11. gnss-Location-r16: enumerated
	{
		if seq.IsComponentPresent(8) {
			idx, err := d.DecodeEnumerated(uEBasedPerfMeasParametersR16GnssLocationR16Constraints)
			if err != nil {
				return err
			}
			ie.Gnss_Location_r16 = &idx
		}
	}

	// 12. ulPDCP-Delay-r16: enumerated
	{
		if seq.IsComponentPresent(9) {
			idx, err := d.DecodeEnumerated(uEBasedPerfMeasParametersR16UlPDCPDelayR16Constraints)
			if err != nil {
				return err
			}
			ie.UlPDCP_Delay_r16 = &idx
		}
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
				{Name: "sigBasedLogMDT-OverrideProtect-r17", Optional: true},
				{Name: "multipleCEF-Report-r17", Optional: true},
				{Name: "excessPacketDelay-r17", Optional: true},
				{Name: "earlyMeasLog-r17", Optional: true},
			},
		})
		if err := groupSeq.DecodePreamble(); err != nil {
			return err
		}

		if groupSeq.IsComponentPresent(0) {
			v, err := dx.DecodeEnumerated(uEBasedPerfMeasParametersR16ExtSigBasedLogMDTOverrideProtectR17Constraints)
			if err != nil {
				return err
			}
			ie.SigBasedLogMDT_OverrideProtect_r17 = &v
		}

		if groupSeq.IsComponentPresent(1) {
			v, err := dx.DecodeEnumerated(uEBasedPerfMeasParametersR16ExtMultipleCEFReportR17Constraints)
			if err != nil {
				return err
			}
			ie.MultipleCEF_Report_r17 = &v
		}

		if groupSeq.IsComponentPresent(2) {
			v, err := dx.DecodeEnumerated(uEBasedPerfMeasParametersR16ExtExcessPacketDelayR17Constraints)
			if err != nil {
				return err
			}
			ie.ExcessPacketDelay_r17 = &v
		}

		if groupSeq.IsComponentPresent(3) {
			v, err := dx.DecodeEnumerated(uEBasedPerfMeasParametersR16ExtEarlyMeasLogR17Constraints)
			if err != nil {
				return err
			}
			ie.EarlyMeasLog_r17 = &v
		}
	}

	// Extension group 1:
	if seq.IsExtensionPresent(1) {
		dx := per.NewDecoder(extValues[extValueIndex], d.Variant())
		extValueIndex++
		groupSeq := dx.NewSequenceDecoder(per.SequenceConstraints{
			Extensible: false,
			RootComponents: []per.ComponentInfo{
				{Name: "loggedMDT-PNI-NPN-r18", Optional: true},
				{Name: "loggedMDT-SNPN-r18", Optional: true},
			},
		})
		if err := groupSeq.DecodePreamble(); err != nil {
			return err
		}

		if groupSeq.IsComponentPresent(0) {
			v, err := dx.DecodeEnumerated(uEBasedPerfMeasParametersR16ExtLoggedMDTPNINPNR18Constraints)
			if err != nil {
				return err
			}
			ie.LoggedMDT_PNI_NPN_r18 = &v
		}

		if groupSeq.IsComponentPresent(1) {
			v, err := dx.DecodeEnumerated(uEBasedPerfMeasParametersR16ExtLoggedMDTSNPNR18Constraints)
			if err != nil {
				return err
			}
			ie.LoggedMDT_SNPN_r18 = &v
		}
	}

	// Extension group 2:
	if seq.IsExtensionPresent(2) {
		dx := per.NewDecoder(extValues[extValueIndex], d.Variant())
		extValueIndex++
		groupSeq := dx.NewSequenceDecoder(per.SequenceConstraints{
			Extensible: false,
			RootComponents: []per.ComponentInfo{
				{Name: "geoAreaScopeChecking-r19", Optional: true},
				{Name: "loggedMDT-Slicing-r19", Optional: true},
			},
		})
		if err := groupSeq.DecodePreamble(); err != nil {
			return err
		}

		if groupSeq.IsComponentPresent(0) {
			v, err := dx.DecodeEnumerated(uEBasedPerfMeasParametersR16ExtGeoAreaScopeCheckingR19Constraints)
			if err != nil {
				return err
			}
			ie.GeoAreaScopeChecking_r19 = &v
		}

		if groupSeq.IsComponentPresent(1) {
			v, err := dx.DecodeEnumerated(uEBasedPerfMeasParametersR16ExtLoggedMDTSlicingR19Constraints)
			if err != nil {
				return err
			}
			ie.LoggedMDT_Slicing_r19 = &v
		}
	}

	return nil
}
