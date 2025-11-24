package ies

import (
	"bytes"

	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type RateMatchPattern struct {
	RateMatchPatternId     RateMatchPatternId            `madatory`
	PatternType            *RateMatchPattern_patternType `lb:275,ub:275,optional`
	SubcarrierSpacing      *SubcarrierSpacing            `optional,ext`
	Dummy                  RateMatchPattern_dummy        `madatory,ext`
	ControlResourceSet_r16 *ControlResourceSetId_r16     `optional,ext-1`
}

func (ie *RateMatchPattern) Encode(w *uper.UperWriter) error {
	var err error
	hasExtensions := ie.ControlResourceSet_r16 != nil
	preambleBits := []bool{hasExtensions, ie.PatternType != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if err = ie.RateMatchPatternId.Encode(w); err != nil {
		return utils.WrapError("Encode RateMatchPatternId", err)
	}
	if ie.PatternType != nil {
		if err = ie.PatternType.Encode(w); err != nil {
			return utils.WrapError("Encode PatternType", err)
		}
	}
	if hasExtensions {
		// Extension bitmap: 1 bits for 1 extension groups
		extBitmap := []bool{ie.ControlResourceSet_r16 != nil}
		if err := w.WriteExtBitMap(extBitmap); err != nil {
			return utils.WrapError("WriteExtBitMap RateMatchPattern", err)
		}

		// encode extension group 1
		if extBitmap[0] {
			extBuf := new(bytes.Buffer)
			extWriter := uper.NewWriter(extBuf)

			// Write preamble bits for optional fields in extension group 1
			optionals_ext_1 := []bool{ie.ControlResourceSet_r16 != nil}
			for _, bit := range optionals_ext_1 {
				if err := extWriter.WriteBool(bit); err != nil {
					return err
				}
			}

			// encode ControlResourceSet_r16 optional
			if ie.ControlResourceSet_r16 != nil {
				if err = ie.ControlResourceSet_r16.Encode(extWriter); err != nil {
					return utils.WrapError("Encode ControlResourceSet_r16", err)
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

func (ie *RateMatchPattern) Decode(r *uper.UperReader) error {
	var err error
	var extensionBit bool
	if extensionBit, err = r.ReadBool(); err != nil {
		return err
	}
	var PatternTypePresent bool
	if PatternTypePresent, err = r.ReadBool(); err != nil {
		return err
	}
	if err = ie.RateMatchPatternId.Decode(r); err != nil {
		return utils.WrapError("Decode RateMatchPatternId", err)
	}
	if PatternTypePresent {
		ie.PatternType = new(RateMatchPattern_patternType)
		if err = ie.PatternType.Decode(r); err != nil {
			return utils.WrapError("Decode PatternType", err)
		}
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

			extReader := uper.NewReader(bytes.NewReader(extBytes))

			ControlResourceSet_r16Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			// decode ControlResourceSet_r16 optional
			if ControlResourceSet_r16Present {
				ie.ControlResourceSet_r16 = new(ControlResourceSetId_r16)
				if err = ie.ControlResourceSet_r16.Decode(extReader); err != nil {
					return utils.WrapError("Decode ControlResourceSet_r16", err)
				}
			}
		}
	}
	return nil
}
