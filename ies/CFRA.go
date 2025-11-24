package ies

import (
	"bytes"

	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type CFRA struct {
	Occasions                 *CFRA_occasions `optional`
	Resources                 CFRA_resources  `lb:1,ub:maxRA_SSB_Resources,madatory`
	TotalNumberOfRA_Preambles *int64          `lb:1,ub:63,optional,ext-1`
}

func (ie *CFRA) Encode(w *uper.UperWriter) error {
	var err error
	hasExtensions := ie.TotalNumberOfRA_Preambles != nil
	preambleBits := []bool{hasExtensions, ie.Occasions != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.Occasions != nil {
		if err = ie.Occasions.Encode(w); err != nil {
			return utils.WrapError("Encode Occasions", err)
		}
	}
	if err = ie.Resources.Encode(w); err != nil {
		return utils.WrapError("Encode Resources", err)
	}
	if hasExtensions {
		// Extension bitmap: 1 bits for 1 extension groups
		extBitmap := []bool{ie.TotalNumberOfRA_Preambles != nil}
		if err := w.WriteExtBitMap(extBitmap); err != nil {
			return utils.WrapError("WriteExtBitMap CFRA", err)
		}

		// encode extension group 1
		if extBitmap[0] {
			extBuf := new(bytes.Buffer)
			extWriter := uper.NewWriter(extBuf)

			// Write preamble bits for optional fields in extension group 1
			optionals_ext_1 := []bool{ie.TotalNumberOfRA_Preambles != nil}
			for _, bit := range optionals_ext_1 {
				if err := extWriter.WriteBool(bit); err != nil {
					return err
				}
			}

			// encode TotalNumberOfRA_Preambles optional
			if ie.TotalNumberOfRA_Preambles != nil {
				if err = extWriter.WriteInteger(*ie.TotalNumberOfRA_Preambles, &uper.Constraint{Lb: 1, Ub: 63}, false); err != nil {
					return utils.WrapError("Encode TotalNumberOfRA_Preambles", err)
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

func (ie *CFRA) Decode(r *uper.UperReader) error {
	var err error
	var extensionBit bool
	if extensionBit, err = r.ReadBool(); err != nil {
		return err
	}
	var OccasionsPresent bool
	if OccasionsPresent, err = r.ReadBool(); err != nil {
		return err
	}
	if OccasionsPresent {
		ie.Occasions = new(CFRA_occasions)
		if err = ie.Occasions.Decode(r); err != nil {
			return utils.WrapError("Decode Occasions", err)
		}
	}
	if err = ie.Resources.Decode(r); err != nil {
		return utils.WrapError("Decode Resources", err)
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

			TotalNumberOfRA_PreamblesPresent, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			// decode TotalNumberOfRA_Preambles optional
			if TotalNumberOfRA_PreamblesPresent {
				var tmp_int_TotalNumberOfRA_Preambles int64
				if tmp_int_TotalNumberOfRA_Preambles, err = extReader.ReadInteger(&uper.Constraint{Lb: 1, Ub: 63}, false); err != nil {
					return utils.WrapError("Decode TotalNumberOfRA_Preambles", err)
				}
				ie.TotalNumberOfRA_Preambles = &tmp_int_TotalNumberOfRA_Preambles
			}
		}
	}
	return nil
}
