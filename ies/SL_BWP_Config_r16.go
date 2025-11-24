package ies

import (
	"bytes"

	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type SL_BWP_Config_r16 struct {
	Sl_BWP_Id                 BWP_Id                     `madatory`
	Sl_BWP_Generic_r16        *SL_BWP_Generic_r16        `optional`
	Sl_BWP_PoolConfig_r16     *SL_BWP_PoolConfig_r16     `optional`
	Sl_BWP_PoolConfigPS_r17   *SL_BWP_PoolConfig_r16     `optional,ext-1,setuprelease`
	Sl_BWP_DiscPoolConfig_r17 *SL_BWP_DiscPoolConfig_r17 `optional,ext-1,setuprelease`
}

func (ie *SL_BWP_Config_r16) Encode(w *uper.UperWriter) error {
	var err error
	hasExtensions := ie.Sl_BWP_PoolConfigPS_r17 != nil || ie.Sl_BWP_DiscPoolConfig_r17 != nil
	preambleBits := []bool{hasExtensions, ie.Sl_BWP_Generic_r16 != nil, ie.Sl_BWP_PoolConfig_r16 != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if err = ie.Sl_BWP_Id.Encode(w); err != nil {
		return utils.WrapError("Encode Sl_BWP_Id", err)
	}
	if ie.Sl_BWP_Generic_r16 != nil {
		if err = ie.Sl_BWP_Generic_r16.Encode(w); err != nil {
			return utils.WrapError("Encode Sl_BWP_Generic_r16", err)
		}
	}
	if ie.Sl_BWP_PoolConfig_r16 != nil {
		if err = ie.Sl_BWP_PoolConfig_r16.Encode(w); err != nil {
			return utils.WrapError("Encode Sl_BWP_PoolConfig_r16", err)
		}
	}
	if hasExtensions {
		// Extension bitmap: 1 bits for 1 extension groups
		extBitmap := []bool{ie.Sl_BWP_PoolConfigPS_r17 != nil || ie.Sl_BWP_DiscPoolConfig_r17 != nil}
		if err := w.WriteExtBitMap(extBitmap); err != nil {
			return utils.WrapError("WriteExtBitMap SL_BWP_Config_r16", err)
		}

		// encode extension group 1
		if extBitmap[0] {
			extBuf := new(bytes.Buffer)
			extWriter := uper.NewWriter(extBuf)

			// Write preamble bits for optional fields in extension group 1
			optionals_ext_1 := []bool{ie.Sl_BWP_PoolConfigPS_r17 != nil, ie.Sl_BWP_DiscPoolConfig_r17 != nil}
			for _, bit := range optionals_ext_1 {
				if err := extWriter.WriteBool(bit); err != nil {
					return err
				}
			}

			// encode Sl_BWP_PoolConfigPS_r17 optional
			if ie.Sl_BWP_PoolConfigPS_r17 != nil {
				tmp_Sl_BWP_PoolConfigPS_r17 := utils.SetupRelease[*SL_BWP_PoolConfig_r16]{
					Setup: ie.Sl_BWP_PoolConfigPS_r17,
				}
				if err = tmp_Sl_BWP_PoolConfigPS_r17.Encode(extWriter); err != nil {
					return utils.WrapError("Encode Sl_BWP_PoolConfigPS_r17", err)
				}
			}
			// encode Sl_BWP_DiscPoolConfig_r17 optional
			if ie.Sl_BWP_DiscPoolConfig_r17 != nil {
				tmp_Sl_BWP_DiscPoolConfig_r17 := utils.SetupRelease[*SL_BWP_DiscPoolConfig_r17]{
					Setup: ie.Sl_BWP_DiscPoolConfig_r17,
				}
				if err = tmp_Sl_BWP_DiscPoolConfig_r17.Encode(extWriter); err != nil {
					return utils.WrapError("Encode Sl_BWP_DiscPoolConfig_r17", err)
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

func (ie *SL_BWP_Config_r16) Decode(r *uper.UperReader) error {
	var err error
	var extensionBit bool
	if extensionBit, err = r.ReadBool(); err != nil {
		return err
	}
	var Sl_BWP_Generic_r16Present bool
	if Sl_BWP_Generic_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var Sl_BWP_PoolConfig_r16Present bool
	if Sl_BWP_PoolConfig_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	if err = ie.Sl_BWP_Id.Decode(r); err != nil {
		return utils.WrapError("Decode Sl_BWP_Id", err)
	}
	if Sl_BWP_Generic_r16Present {
		ie.Sl_BWP_Generic_r16 = new(SL_BWP_Generic_r16)
		if err = ie.Sl_BWP_Generic_r16.Decode(r); err != nil {
			return utils.WrapError("Decode Sl_BWP_Generic_r16", err)
		}
	}
	if Sl_BWP_PoolConfig_r16Present {
		ie.Sl_BWP_PoolConfig_r16 = new(SL_BWP_PoolConfig_r16)
		if err = ie.Sl_BWP_PoolConfig_r16.Decode(r); err != nil {
			return utils.WrapError("Decode Sl_BWP_PoolConfig_r16", err)
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

			Sl_BWP_PoolConfigPS_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			Sl_BWP_DiscPoolConfig_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			// decode Sl_BWP_PoolConfigPS_r17 optional
			if Sl_BWP_PoolConfigPS_r17Present {
				tmp_Sl_BWP_PoolConfigPS_r17 := utils.SetupRelease[*SL_BWP_PoolConfig_r16]{}
				if err = tmp_Sl_BWP_PoolConfigPS_r17.Decode(extReader); err != nil {
					return utils.WrapError("Decode Sl_BWP_PoolConfigPS_r17", err)
				}
				ie.Sl_BWP_PoolConfigPS_r17 = tmp_Sl_BWP_PoolConfigPS_r17.Setup
			}
			// decode Sl_BWP_DiscPoolConfig_r17 optional
			if Sl_BWP_DiscPoolConfig_r17Present {
				tmp_Sl_BWP_DiscPoolConfig_r17 := utils.SetupRelease[*SL_BWP_DiscPoolConfig_r17]{}
				if err = tmp_Sl_BWP_DiscPoolConfig_r17.Decode(extReader); err != nil {
					return utils.WrapError("Decode Sl_BWP_DiscPoolConfig_r17", err)
				}
				ie.Sl_BWP_DiscPoolConfig_r17 = tmp_Sl_BWP_DiscPoolConfig_r17.Setup
			}
		}
	}
	return nil
}
