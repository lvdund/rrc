// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: CRS-InterfMitigation-r17 (line 19966).

var cRSInterfMitigationR17Constraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "crs-IM-DSS-15kHzSCS-r17", Optional: true},
		{Name: "crs-IM-nonDSS-15kHzSCS-r17", Optional: true},
		{Name: "crs-IM-nonDSS-NWA-15kHzSCS-r17", Optional: true},
		{Name: "crs-IM-nonDSS-30kHzSCS-r17", Optional: true},
		{Name: "crs-IM-nonDSS-NWA-30kHzSCS-r17", Optional: true},
	},
}

const (
	CRS_InterfMitigation_r17_Crs_IM_DSS_15kHzSCS_r17_Supported = 0
)

var cRSInterfMitigationR17CrsIMDSS15kHzSCSR17Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{CRS_InterfMitigation_r17_Crs_IM_DSS_15kHzSCS_r17_Supported},
}

const (
	CRS_InterfMitigation_r17_Crs_IM_NonDSS_15kHzSCS_r17_Supported = 0
)

var cRSInterfMitigationR17CrsIMNonDSS15kHzSCSR17Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{CRS_InterfMitigation_r17_Crs_IM_NonDSS_15kHzSCS_r17_Supported},
}

const (
	CRS_InterfMitigation_r17_Crs_IM_NonDSS_NWA_15kHzSCS_r17_Supported = 0
)

var cRSInterfMitigationR17CrsIMNonDSSNWA15kHzSCSR17Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{CRS_InterfMitigation_r17_Crs_IM_NonDSS_NWA_15kHzSCS_r17_Supported},
}

const (
	CRS_InterfMitigation_r17_Crs_IM_NonDSS_30kHzSCS_r17_Supported = 0
)

var cRSInterfMitigationR17CrsIMNonDSS30kHzSCSR17Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{CRS_InterfMitigation_r17_Crs_IM_NonDSS_30kHzSCS_r17_Supported},
}

const (
	CRS_InterfMitigation_r17_Crs_IM_NonDSS_NWA_30kHzSCS_r17_Supported = 0
)

var cRSInterfMitigationR17CrsIMNonDSSNWA30kHzSCSR17Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{CRS_InterfMitigation_r17_Crs_IM_NonDSS_NWA_30kHzSCS_r17_Supported},
}

type CRS_InterfMitigation_r17 struct {
	Crs_IM_DSS_15kHzSCS_r17        *int64
	Crs_IM_NonDSS_15kHzSCS_r17     *int64
	Crs_IM_NonDSS_NWA_15kHzSCS_r17 *int64
	Crs_IM_NonDSS_30kHzSCS_r17     *int64
	Crs_IM_NonDSS_NWA_30kHzSCS_r17 *int64
}

func (ie *CRS_InterfMitigation_r17) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(cRSInterfMitigationR17Constraints)

	// 1. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.Crs_IM_DSS_15kHzSCS_r17 != nil, ie.Crs_IM_NonDSS_15kHzSCS_r17 != nil, ie.Crs_IM_NonDSS_NWA_15kHzSCS_r17 != nil, ie.Crs_IM_NonDSS_30kHzSCS_r17 != nil, ie.Crs_IM_NonDSS_NWA_30kHzSCS_r17 != nil}); err != nil {
		return err
	}

	// 2. crs-IM-DSS-15kHzSCS-r17: enumerated
	{
		if ie.Crs_IM_DSS_15kHzSCS_r17 != nil {
			if err := e.EncodeEnumerated(*ie.Crs_IM_DSS_15kHzSCS_r17, cRSInterfMitigationR17CrsIMDSS15kHzSCSR17Constraints); err != nil {
				return err
			}
		}
	}

	// 3. crs-IM-nonDSS-15kHzSCS-r17: enumerated
	{
		if ie.Crs_IM_NonDSS_15kHzSCS_r17 != nil {
			if err := e.EncodeEnumerated(*ie.Crs_IM_NonDSS_15kHzSCS_r17, cRSInterfMitigationR17CrsIMNonDSS15kHzSCSR17Constraints); err != nil {
				return err
			}
		}
	}

	// 4. crs-IM-nonDSS-NWA-15kHzSCS-r17: enumerated
	{
		if ie.Crs_IM_NonDSS_NWA_15kHzSCS_r17 != nil {
			if err := e.EncodeEnumerated(*ie.Crs_IM_NonDSS_NWA_15kHzSCS_r17, cRSInterfMitigationR17CrsIMNonDSSNWA15kHzSCSR17Constraints); err != nil {
				return err
			}
		}
	}

	// 5. crs-IM-nonDSS-30kHzSCS-r17: enumerated
	{
		if ie.Crs_IM_NonDSS_30kHzSCS_r17 != nil {
			if err := e.EncodeEnumerated(*ie.Crs_IM_NonDSS_30kHzSCS_r17, cRSInterfMitigationR17CrsIMNonDSS30kHzSCSR17Constraints); err != nil {
				return err
			}
		}
	}

	// 6. crs-IM-nonDSS-NWA-30kHzSCS-r17: enumerated
	{
		if ie.Crs_IM_NonDSS_NWA_30kHzSCS_r17 != nil {
			if err := e.EncodeEnumerated(*ie.Crs_IM_NonDSS_NWA_30kHzSCS_r17, cRSInterfMitigationR17CrsIMNonDSSNWA30kHzSCSR17Constraints); err != nil {
				return err
			}
		}
	}

	return nil
}

func (ie *CRS_InterfMitigation_r17) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(cRSInterfMitigationR17Constraints)

	// 1. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 2. crs-IM-DSS-15kHzSCS-r17: enumerated
	{
		if seq.IsComponentPresent(0) {
			idx, err := d.DecodeEnumerated(cRSInterfMitigationR17CrsIMDSS15kHzSCSR17Constraints)
			if err != nil {
				return err
			}
			ie.Crs_IM_DSS_15kHzSCS_r17 = &idx
		}
	}

	// 3. crs-IM-nonDSS-15kHzSCS-r17: enumerated
	{
		if seq.IsComponentPresent(1) {
			idx, err := d.DecodeEnumerated(cRSInterfMitigationR17CrsIMNonDSS15kHzSCSR17Constraints)
			if err != nil {
				return err
			}
			ie.Crs_IM_NonDSS_15kHzSCS_r17 = &idx
		}
	}

	// 4. crs-IM-nonDSS-NWA-15kHzSCS-r17: enumerated
	{
		if seq.IsComponentPresent(2) {
			idx, err := d.DecodeEnumerated(cRSInterfMitigationR17CrsIMNonDSSNWA15kHzSCSR17Constraints)
			if err != nil {
				return err
			}
			ie.Crs_IM_NonDSS_NWA_15kHzSCS_r17 = &idx
		}
	}

	// 5. crs-IM-nonDSS-30kHzSCS-r17: enumerated
	{
		if seq.IsComponentPresent(3) {
			idx, err := d.DecodeEnumerated(cRSInterfMitigationR17CrsIMNonDSS30kHzSCSR17Constraints)
			if err != nil {
				return err
			}
			ie.Crs_IM_NonDSS_30kHzSCS_r17 = &idx
		}
	}

	// 6. crs-IM-nonDSS-NWA-30kHzSCS-r17: enumerated
	{
		if seq.IsComponentPresent(4) {
			idx, err := d.DecodeEnumerated(cRSInterfMitigationR17CrsIMNonDSSNWA30kHzSCSR17Constraints)
			if err != nil {
				return err
			}
			ie.Crs_IM_NonDSS_NWA_30kHzSCS_r17 = &idx
		}
	}

	return nil
}
