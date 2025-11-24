package ies

import (
	"bytes"

	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type QuantityConfig struct {
	QuantityConfigNR_List      []QuantityConfigNR          `lb:1,ub:maxNrofQuantityConfig,optional`
	QuantityConfigEUTRA        *FilterConfig               `optional,ext-1`
	QuantityConfigUTRA_FDD_r16 *QuantityConfigUTRA_FDD_r16 `optional,ext-2`
	QuantityConfigCLI_r16      *FilterConfigCLI_r16        `optional,ext-2`
}

func (ie *QuantityConfig) Encode(w *uper.UperWriter) error {
	var err error
	hasExtensions := ie.QuantityConfigEUTRA != nil || ie.QuantityConfigUTRA_FDD_r16 != nil || ie.QuantityConfigCLI_r16 != nil
	preambleBits := []bool{hasExtensions, len(ie.QuantityConfigNR_List) > 0}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if len(ie.QuantityConfigNR_List) > 0 {
		tmp_QuantityConfigNR_List := utils.NewSequence[*QuantityConfigNR]([]*QuantityConfigNR{}, uper.Constraint{Lb: 1, Ub: maxNrofQuantityConfig}, false)
		for _, i := range ie.QuantityConfigNR_List {
			tmp_QuantityConfigNR_List.Value = append(tmp_QuantityConfigNR_List.Value, &i)
		}
		if err = tmp_QuantityConfigNR_List.Encode(w); err != nil {
			return utils.WrapError("Encode QuantityConfigNR_List", err)
		}
	}
	if hasExtensions {
		// Extension bitmap: 2 bits for 2 extension groups
		extBitmap := []bool{ie.QuantityConfigEUTRA != nil, ie.QuantityConfigUTRA_FDD_r16 != nil || ie.QuantityConfigCLI_r16 != nil}
		if err := w.WriteExtBitMap(extBitmap); err != nil {
			return utils.WrapError("WriteExtBitMap QuantityConfig", err)
		}

		// encode extension group 1
		if extBitmap[0] {
			extBuf := new(bytes.Buffer)
			extWriter := uper.NewWriter(extBuf)

			// Write preamble bits for optional fields in extension group 1
			optionals_ext_1 := []bool{ie.QuantityConfigEUTRA != nil}
			for _, bit := range optionals_ext_1 {
				if err := extWriter.WriteBool(bit); err != nil {
					return err
				}
			}

			// encode QuantityConfigEUTRA optional
			if ie.QuantityConfigEUTRA != nil {
				if err = ie.QuantityConfigEUTRA.Encode(extWriter); err != nil {
					return utils.WrapError("Encode QuantityConfigEUTRA", err)
				}
			}

			if err := extWriter.Close(); err != nil {
				return err
			}

			if err := w.WriteOpenType(extBuf.Bytes()); err != nil {
				return err
			}
		}

		// encode extension group 2
		if extBitmap[1] {
			extBuf := new(bytes.Buffer)
			extWriter := uper.NewWriter(extBuf)

			// Write preamble bits for optional fields in extension group 2
			optionals_ext_2 := []bool{ie.QuantityConfigUTRA_FDD_r16 != nil, ie.QuantityConfigCLI_r16 != nil}
			for _, bit := range optionals_ext_2 {
				if err := extWriter.WriteBool(bit); err != nil {
					return err
				}
			}

			// encode QuantityConfigUTRA_FDD_r16 optional
			if ie.QuantityConfigUTRA_FDD_r16 != nil {
				if err = ie.QuantityConfigUTRA_FDD_r16.Encode(extWriter); err != nil {
					return utils.WrapError("Encode QuantityConfigUTRA_FDD_r16", err)
				}
			}
			// encode QuantityConfigCLI_r16 optional
			if ie.QuantityConfigCLI_r16 != nil {
				if err = ie.QuantityConfigCLI_r16.Encode(extWriter); err != nil {
					return utils.WrapError("Encode QuantityConfigCLI_r16", err)
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

func (ie *QuantityConfig) Decode(r *uper.UperReader) error {
	var err error
	var extensionBit bool
	if extensionBit, err = r.ReadBool(); err != nil {
		return err
	}
	var QuantityConfigNR_ListPresent bool
	if QuantityConfigNR_ListPresent, err = r.ReadBool(); err != nil {
		return err
	}
	if QuantityConfigNR_ListPresent {
		tmp_QuantityConfigNR_List := utils.NewSequence[*QuantityConfigNR]([]*QuantityConfigNR{}, uper.Constraint{Lb: 1, Ub: maxNrofQuantityConfig}, false)
		fn_QuantityConfigNR_List := func() *QuantityConfigNR {
			return new(QuantityConfigNR)
		}
		if err = tmp_QuantityConfigNR_List.Decode(r, fn_QuantityConfigNR_List); err != nil {
			return utils.WrapError("Decode QuantityConfigNR_List", err)
		}
		ie.QuantityConfigNR_List = []QuantityConfigNR{}
		for _, i := range tmp_QuantityConfigNR_List.Value {
			ie.QuantityConfigNR_List = append(ie.QuantityConfigNR_List, *i)
		}
	}

	if extensionBit {
		// Read extension bitmap: 2 bits for 2 extension groups
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

			QuantityConfigEUTRAPresent, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			// decode QuantityConfigEUTRA optional
			if QuantityConfigEUTRAPresent {
				ie.QuantityConfigEUTRA = new(FilterConfig)
				if err = ie.QuantityConfigEUTRA.Decode(extReader); err != nil {
					return utils.WrapError("Decode QuantityConfigEUTRA", err)
				}
			}
		}
		// decode extension group 2
		if len(extBitmap) > 1 && extBitmap[1] {
			extBytes, err := r.ReadOpenType()
			if err != nil {
				return err
			}

			extReader := uper.NewReader(bytes.NewReader(extBytes))

			QuantityConfigUTRA_FDD_r16Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			QuantityConfigCLI_r16Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			// decode QuantityConfigUTRA_FDD_r16 optional
			if QuantityConfigUTRA_FDD_r16Present {
				ie.QuantityConfigUTRA_FDD_r16 = new(QuantityConfigUTRA_FDD_r16)
				if err = ie.QuantityConfigUTRA_FDD_r16.Decode(extReader); err != nil {
					return utils.WrapError("Decode QuantityConfigUTRA_FDD_r16", err)
				}
			}
			// decode QuantityConfigCLI_r16 optional
			if QuantityConfigCLI_r16Present {
				ie.QuantityConfigCLI_r16 = new(FilterConfigCLI_r16)
				if err = ie.QuantityConfigCLI_r16.Decode(extReader); err != nil {
					return utils.WrapError("Decode QuantityConfigCLI_r16", err)
				}
			}
		}
	}
	return nil
}
