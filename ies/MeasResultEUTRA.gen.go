// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: MeasResultEUTRA (line 9796).

var measResultEUTRAConstraints = per.SequenceConstraints{
	Extensible: true,
	RootComponents: []per.ComponentInfo{
		{Name: "eutra-PhysCellId"},
		{Name: "measResult"},
		{Name: "cgi-Info", Optional: true},
	},
	ExtComponents: []per.ComponentInfo{
		{Name: "extension-addition-0", Optional: true},
	},
}

const (
	MeasResultEUTRA_Ext_Hsdn_Cell_r19_True = 0
)

var measResultEUTRAExtHsdnCellR19Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{MeasResultEUTRA_Ext_Hsdn_Cell_r19_True},
}

type MeasResultEUTRA struct {
	Eutra_PhysCellId PhysCellId
	MeasResult       MeasQuantityResultsEUTRA
	Cgi_Info         *CGI_InfoEUTRA
	Hsdn_Cell_r19    *int64
}

func (ie *MeasResultEUTRA) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(measResultEUTRAConstraints)

	// Extension-group presence (one var per [[...]] group).
	hasExtGroup0 := ie.Hsdn_Cell_r19 != nil
	hasExtensions := hasExtGroup0

	// 1. SEQUENCE extension bit.
	if err := seq.EncodeExtensionBit(hasExtensions); err != nil {
		return err
	}

	// 2. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.Cgi_Info != nil}); err != nil {
		return err
	}

	// 3. eutra-PhysCellId: ref
	{
		if err := ie.Eutra_PhysCellId.Encode(e); err != nil {
			return err
		}
	}

	// 4. measResult: ref
	{
		if err := ie.MeasResult.Encode(e); err != nil {
			return err
		}
	}

	// 5. cgi-Info: ref
	{
		if ie.Cgi_Info != nil {
			if err := ie.Cgi_Info.Encode(e); err != nil {
				return err
			}
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
					{Name: "hsdn-Cell-r19", Optional: true},
				},
			})
			if err := groupSeq.EncodePreamble([]bool{ie.Hsdn_Cell_r19 != nil}); err != nil {
				return err
			}

			if ie.Hsdn_Cell_r19 != nil {
				if err := ex.EncodeEnumerated(*ie.Hsdn_Cell_r19, measResultEUTRAExtHsdnCellR19Constraints); err != nil {
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

func (ie *MeasResultEUTRA) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(measResultEUTRAConstraints)

	// 1. SEQUENCE extension bit.
	if err := seq.DecodeExtensionBit(); err != nil {
		return err
	}

	// 2. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 3. eutra-PhysCellId: ref
	{
		if err := ie.Eutra_PhysCellId.Decode(d); err != nil {
			return err
		}
	}

	// 4. measResult: ref
	{
		if err := ie.MeasResult.Decode(d); err != nil {
			return err
		}
	}

	// 5. cgi-Info: ref
	{
		if seq.IsComponentPresent(2) {
			ie.Cgi_Info = new(CGI_InfoEUTRA)
			if err := ie.Cgi_Info.Decode(d); err != nil {
				return err
			}
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
				{Name: "hsdn-Cell-r19", Optional: true},
			},
		})
		if err := groupSeq.DecodePreamble(); err != nil {
			return err
		}

		if groupSeq.IsComponentPresent(0) {
			v, err := dx.DecodeEnumerated(measResultEUTRAExtHsdnCellR19Constraints)
			if err != nil {
				return err
			}
			ie.Hsdn_Cell_r19 = &v
		}
	}

	return nil
}
