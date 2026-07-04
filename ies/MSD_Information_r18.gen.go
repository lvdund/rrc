// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: MSD-Information-r18 (line 24681).

var mSDInformationR18Constraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "msd-Type-r18"},
		{Name: "msd-PowerClass-r18"},
		{Name: "msd-Class-r18"},
	},
}

const (
	MSD_Information_r18_Msd_Type_r18_Harmonic           = 0
	MSD_Information_r18_Msd_Type_r18_HarmonicMixing     = 1
	MSD_Information_r18_Msd_Type_r18_CrossBandIsolation = 2
	MSD_Information_r18_Msd_Type_r18_Imd2               = 3
	MSD_Information_r18_Msd_Type_r18_Imd3               = 4
	MSD_Information_r18_Msd_Type_r18_Imd4               = 5
	MSD_Information_r18_Msd_Type_r18_Imd5               = 6
	MSD_Information_r18_Msd_Type_r18_All                = 7
	MSD_Information_r18_Msd_Type_r18_Spare8             = 8
	MSD_Information_r18_Msd_Type_r18_Spare7             = 9
	MSD_Information_r18_Msd_Type_r18_Spare6             = 10
	MSD_Information_r18_Msd_Type_r18_Spare5             = 11
	MSD_Information_r18_Msd_Type_r18_Spare4             = 12
	MSD_Information_r18_Msd_Type_r18_Spare3             = 13
	MSD_Information_r18_Msd_Type_r18_Spare2             = 14
	MSD_Information_r18_Msd_Type_r18_Spare1             = 15
)

var mSDInformationR18MsdTypeR18Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{MSD_Information_r18_Msd_Type_r18_Harmonic, MSD_Information_r18_Msd_Type_r18_HarmonicMixing, MSD_Information_r18_Msd_Type_r18_CrossBandIsolation, MSD_Information_r18_Msd_Type_r18_Imd2, MSD_Information_r18_Msd_Type_r18_Imd3, MSD_Information_r18_Msd_Type_r18_Imd4, MSD_Information_r18_Msd_Type_r18_Imd5, MSD_Information_r18_Msd_Type_r18_All, MSD_Information_r18_Msd_Type_r18_Spare8, MSD_Information_r18_Msd_Type_r18_Spare7, MSD_Information_r18_Msd_Type_r18_Spare6, MSD_Information_r18_Msd_Type_r18_Spare5, MSD_Information_r18_Msd_Type_r18_Spare4, MSD_Information_r18_Msd_Type_r18_Spare3, MSD_Information_r18_Msd_Type_r18_Spare2, MSD_Information_r18_Msd_Type_r18_Spare1},
}

const (
	MSD_Information_r18_Msd_PowerClass_r18_Pc1dot5 = 0
	MSD_Information_r18_Msd_PowerClass_r18_Pc2     = 1
	MSD_Information_r18_Msd_PowerClass_r18_Pc3     = 2
)

var mSDInformationR18MsdPowerClassR18Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{MSD_Information_r18_Msd_PowerClass_r18_Pc1dot5, MSD_Information_r18_Msd_PowerClass_r18_Pc2, MSD_Information_r18_Msd_PowerClass_r18_Pc3},
}

const (
	MSD_Information_r18_Msd_Class_r18_ClassI    = 0
	MSD_Information_r18_Msd_Class_r18_ClassII   = 1
	MSD_Information_r18_Msd_Class_r18_ClassIII  = 2
	MSD_Information_r18_Msd_Class_r18_ClassIV   = 3
	MSD_Information_r18_Msd_Class_r18_ClassV    = 4
	MSD_Information_r18_Msd_Class_r18_ClassVI   = 5
	MSD_Information_r18_Msd_Class_r18_ClassVII  = 6
	MSD_Information_r18_Msd_Class_r18_ClassVIII = 7
)

var mSDInformationR18MsdClassR18Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{MSD_Information_r18_Msd_Class_r18_ClassI, MSD_Information_r18_Msd_Class_r18_ClassII, MSD_Information_r18_Msd_Class_r18_ClassIII, MSD_Information_r18_Msd_Class_r18_ClassIV, MSD_Information_r18_Msd_Class_r18_ClassV, MSD_Information_r18_Msd_Class_r18_ClassVI, MSD_Information_r18_Msd_Class_r18_ClassVII, MSD_Information_r18_Msd_Class_r18_ClassVIII},
}

type MSD_Information_r18 struct {
	Msd_Type_r18       int64
	Msd_PowerClass_r18 int64
	Msd_Class_r18      int64
}

func (ie *MSD_Information_r18) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(mSDInformationR18Constraints)
	_ = seq

	// 1. msd-Type-r18: enumerated
	{
		if err := e.EncodeEnumerated(ie.Msd_Type_r18, mSDInformationR18MsdTypeR18Constraints); err != nil {
			return err
		}
	}

	// 2. msd-PowerClass-r18: enumerated
	{
		if err := e.EncodeEnumerated(ie.Msd_PowerClass_r18, mSDInformationR18MsdPowerClassR18Constraints); err != nil {
			return err
		}
	}

	// 3. msd-Class-r18: enumerated
	{
		if err := e.EncodeEnumerated(ie.Msd_Class_r18, mSDInformationR18MsdClassR18Constraints); err != nil {
			return err
		}
	}

	return nil
}

func (ie *MSD_Information_r18) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(mSDInformationR18Constraints)
	_ = seq

	// 1. msd-Type-r18: enumerated
	{
		v0, err := d.DecodeEnumerated(mSDInformationR18MsdTypeR18Constraints)
		if err != nil {
			return err
		}
		ie.Msd_Type_r18 = v0
	}

	// 2. msd-PowerClass-r18: enumerated
	{
		v1, err := d.DecodeEnumerated(mSDInformationR18MsdPowerClassR18Constraints)
		if err != nil {
			return err
		}
		ie.Msd_PowerClass_r18 = v1
	}

	// 3. msd-Class-r18: enumerated
	{
		v2, err := d.DecodeEnumerated(mSDInformationR18MsdClassR18Constraints)
		if err != nil {
			return err
		}
		ie.Msd_Class_r18 = v2
	}

	return nil
}
