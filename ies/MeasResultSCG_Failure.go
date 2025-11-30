package ies

import (
	"bytes"

	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

type MeasResultSCG_Failure struct {
	MeasResultPerMOList MeasResultList2NR `madatory`
	LocationInfo_r16    *LocationInfo_r16 `optional,ext-1`
}

func (ie *MeasResultSCG_Failure) Encode(w *aper.AperWriter) error {
	var err error
	hasExtensions := ie.LocationInfo_r16 != nil
	preambleBits := []bool{hasExtensions}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if err = ie.MeasResultPerMOList.Encode(w); err != nil {
		return utils.WrapError("Encode MeasResultPerMOList", err)
	}
	if hasExtensions {
		// Extension bitmap: 1 bits for 1 extension groups
		extBitmap := []bool{ie.LocationInfo_r16 != nil}
		if err := w.WriteExtBitMap(extBitmap); err != nil {
			return utils.WrapError("WriteExtBitMap MeasResultSCG_Failure", err)
		}

		// encode extension group 1
		if extBitmap[0] {
			extBuf := new(bytes.Buffer)
			extWriter := aper.NewWriter(extBuf)

			// Write preamble bits for optional fields in extension group 1
			optionals_ext_1 := []bool{ie.LocationInfo_r16 != nil}
			for _, bit := range optionals_ext_1 {
				if err := extWriter.WriteBool(bit); err != nil {
					return err
				}
			}

			// encode LocationInfo_r16 optional
			if ie.LocationInfo_r16 != nil {
				if err = ie.LocationInfo_r16.Encode(extWriter); err != nil {
					return utils.WrapError("Encode LocationInfo_r16", err)
				}
			}

			if err := extWriter.Close(); err != nil {
				return err
			}

			if err := w.WriteOpenType(extBuf.Bytes()); err != nil {
				return err
			}
		}
	}
	return nil
}

func (ie *MeasResultSCG_Failure) Decode(r *aper.AperReader) error {
	var err error
	var extensionBit bool
	if extensionBit, err = r.ReadBool(); err != nil {
		return err
	}
	if err = ie.MeasResultPerMOList.Decode(r); err != nil {
		return utils.WrapError("Decode MeasResultPerMOList", err)
	}

	if extensionBit {
		// Read extension bitmap: 1 bits for 1 extension groups
		extBitmap, err := r.ReadExtBitMap()
		if err != nil {
			return err
		}

		// decode extension group 1
		if len(extBitmap) > 0 && extBitmap[0] {
			extBytes, err := r.ReadOpenType()
			if err != nil {
				return err
			}

			extReader := aper.NewReader(bytes.NewReader(extBytes))

			LocationInfo_r16Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			// decode LocationInfo_r16 optional
			if LocationInfo_r16Present {
				ie.LocationInfo_r16 = new(LocationInfo_r16)
				if err = ie.LocationInfo_r16.Decode(extReader); err != nil {
					return utils.WrapError("Decode LocationInfo_r16", err)
				}
			}
		}
	}
	return nil
}
