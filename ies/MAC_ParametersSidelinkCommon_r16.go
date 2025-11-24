package ies

import (
	"bytes"

	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type MAC_ParametersSidelinkCommon_r16 struct {
	Lcp_RestrictionSidelink_r16          *MAC_ParametersSidelinkCommon_r16_lcp_RestrictionSidelink_r16          `optional`
	MultipleConfiguredGrantsSidelink_r16 *MAC_ParametersSidelinkCommon_r16_multipleConfiguredGrantsSidelink_r16 `optional`
	Drx_OnSidelink_r17                   *MAC_ParametersSidelinkCommon_r16_drx_OnSidelink_r17                   `optional,ext-1`
}

func (ie *MAC_ParametersSidelinkCommon_r16) Encode(w *uper.UperWriter) error {
	var err error
	hasExtensions := ie.Drx_OnSidelink_r17 != nil
	preambleBits := []bool{hasExtensions, ie.Lcp_RestrictionSidelink_r16 != nil, ie.MultipleConfiguredGrantsSidelink_r16 != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.Lcp_RestrictionSidelink_r16 != nil {
		if err = ie.Lcp_RestrictionSidelink_r16.Encode(w); err != nil {
			return utils.WrapError("Encode Lcp_RestrictionSidelink_r16", err)
		}
	}
	if ie.MultipleConfiguredGrantsSidelink_r16 != nil {
		if err = ie.MultipleConfiguredGrantsSidelink_r16.Encode(w); err != nil {
			return utils.WrapError("Encode MultipleConfiguredGrantsSidelink_r16", err)
		}
	}
	if hasExtensions {
		// Extension bitmap: 1 bits for 1 extension groups
		extBitmap := []bool{ie.Drx_OnSidelink_r17 != nil}
		if err := w.WriteExtBitMap(extBitmap); err != nil {
			return utils.WrapError("WriteExtBitMap MAC_ParametersSidelinkCommon_r16", err)
		}

		// encode extension group 1
		if extBitmap[0] {
			extBuf := new(bytes.Buffer)
			extWriter := uper.NewWriter(extBuf)

			// Write preamble bits for optional fields in extension group 1
			optionals_ext_1 := []bool{ie.Drx_OnSidelink_r17 != nil}
			for _, bit := range optionals_ext_1 {
				if err := extWriter.WriteBool(bit); err != nil {
					return err
				}
			}

			// encode Drx_OnSidelink_r17 optional
			if ie.Drx_OnSidelink_r17 != nil {
				if err = ie.Drx_OnSidelink_r17.Encode(extWriter); err != nil {
					return utils.WrapError("Encode Drx_OnSidelink_r17", err)
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

func (ie *MAC_ParametersSidelinkCommon_r16) Decode(r *uper.UperReader) error {
	var err error
	var extensionBit bool
	if extensionBit, err = r.ReadBool(); err != nil {
		return err
	}
	var Lcp_RestrictionSidelink_r16Present bool
	if Lcp_RestrictionSidelink_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var MultipleConfiguredGrantsSidelink_r16Present bool
	if MultipleConfiguredGrantsSidelink_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	if Lcp_RestrictionSidelink_r16Present {
		ie.Lcp_RestrictionSidelink_r16 = new(MAC_ParametersSidelinkCommon_r16_lcp_RestrictionSidelink_r16)
		if err = ie.Lcp_RestrictionSidelink_r16.Decode(r); err != nil {
			return utils.WrapError("Decode Lcp_RestrictionSidelink_r16", err)
		}
	}
	if MultipleConfiguredGrantsSidelink_r16Present {
		ie.MultipleConfiguredGrantsSidelink_r16 = new(MAC_ParametersSidelinkCommon_r16_multipleConfiguredGrantsSidelink_r16)
		if err = ie.MultipleConfiguredGrantsSidelink_r16.Decode(r); err != nil {
			return utils.WrapError("Decode MultipleConfiguredGrantsSidelink_r16", err)
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

			Drx_OnSidelink_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			// decode Drx_OnSidelink_r17 optional
			if Drx_OnSidelink_r17Present {
				ie.Drx_OnSidelink_r17 = new(MAC_ParametersSidelinkCommon_r16_drx_OnSidelink_r17)
				if err = ie.Drx_OnSidelink_r17.Decode(extReader); err != nil {
					return utils.WrapError("Decode Drx_OnSidelink_r17", err)
				}
			}
		}
	}
	return nil
}
