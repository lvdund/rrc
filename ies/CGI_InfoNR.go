package ies

import (
	"bytes"

	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type CGI_InfoNR struct {
	Plmn_IdentityInfoList       *PLMN_IdentityInfoList                  `optional`
	FrequencyBandList           *MultiFrequencyBandListNR               `optional`
	NoSIB1                      *CGI_InfoNR_noSIB1                      `lb:0,ub:15,optional`
	Npn_IdentityInfoList_r16    *NPN_IdentityInfoList_r16               `optional,ext-1`
	CellReservedForOtherUse_r16 *CGI_InfoNR_cellReservedForOtherUse_r16 `optional,ext-2`
}

func (ie *CGI_InfoNR) Encode(w *uper.UperWriter) error {
	var err error
	hasExtensions := ie.Npn_IdentityInfoList_r16 != nil || ie.CellReservedForOtherUse_r16 != nil
	preambleBits := []bool{hasExtensions, ie.Plmn_IdentityInfoList != nil, ie.FrequencyBandList != nil, ie.NoSIB1 != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.Plmn_IdentityInfoList != nil {
		if err = ie.Plmn_IdentityInfoList.Encode(w); err != nil {
			return utils.WrapError("Encode Plmn_IdentityInfoList", err)
		}
	}
	if ie.FrequencyBandList != nil {
		if err = ie.FrequencyBandList.Encode(w); err != nil {
			return utils.WrapError("Encode FrequencyBandList", err)
		}
	}
	if ie.NoSIB1 != nil {
		if err = ie.NoSIB1.Encode(w); err != nil {
			return utils.WrapError("Encode NoSIB1", err)
		}
	}
	if hasExtensions {
		// Extension bitmap: 2 bits for 2 extension groups
		extBitmap := []bool{ie.Npn_IdentityInfoList_r16 != nil, ie.CellReservedForOtherUse_r16 != nil}
		if err := w.WriteExtBitMap(extBitmap); err != nil {
			return utils.WrapError("WriteExtBitMap CGI_InfoNR", err)
		}

		// encode extension group 1
		if extBitmap[0] {
			extBuf := new(bytes.Buffer)
			extWriter := uper.NewWriter(extBuf)

			// Write preamble bits for optional fields in extension group 1
			optionals_ext_1 := []bool{ie.Npn_IdentityInfoList_r16 != nil}
			for _, bit := range optionals_ext_1 {
				if err := extWriter.WriteBool(bit); err != nil {
					return err
				}
			}

			// encode Npn_IdentityInfoList_r16 optional
			if ie.Npn_IdentityInfoList_r16 != nil {
				if err = ie.Npn_IdentityInfoList_r16.Encode(extWriter); err != nil {
					return utils.WrapError("Encode Npn_IdentityInfoList_r16", err)
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
			optionals_ext_2 := []bool{ie.CellReservedForOtherUse_r16 != nil}
			for _, bit := range optionals_ext_2 {
				if err := extWriter.WriteBool(bit); err != nil {
					return err
				}
			}

			// encode CellReservedForOtherUse_r16 optional
			if ie.CellReservedForOtherUse_r16 != nil {
				if err = ie.CellReservedForOtherUse_r16.Encode(extWriter); err != nil {
					return utils.WrapError("Encode CellReservedForOtherUse_r16", err)
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

func (ie *CGI_InfoNR) Decode(r *uper.UperReader) error {
	var err error
	var extensionBit bool
	if extensionBit, err = r.ReadBool(); err != nil {
		return err
	}
	var Plmn_IdentityInfoListPresent bool
	if Plmn_IdentityInfoListPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var FrequencyBandListPresent bool
	if FrequencyBandListPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var NoSIB1Present bool
	if NoSIB1Present, err = r.ReadBool(); err != nil {
		return err
	}
	if Plmn_IdentityInfoListPresent {
		ie.Plmn_IdentityInfoList = new(PLMN_IdentityInfoList)
		if err = ie.Plmn_IdentityInfoList.Decode(r); err != nil {
			return utils.WrapError("Decode Plmn_IdentityInfoList", err)
		}
	}
	if FrequencyBandListPresent {
		ie.FrequencyBandList = new(MultiFrequencyBandListNR)
		if err = ie.FrequencyBandList.Decode(r); err != nil {
			return utils.WrapError("Decode FrequencyBandList", err)
		}
	}
	if NoSIB1Present {
		ie.NoSIB1 = new(CGI_InfoNR_noSIB1)
		if err = ie.NoSIB1.Decode(r); err != nil {
			return utils.WrapError("Decode NoSIB1", err)
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

			Npn_IdentityInfoList_r16Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			// decode Npn_IdentityInfoList_r16 optional
			if Npn_IdentityInfoList_r16Present {
				ie.Npn_IdentityInfoList_r16 = new(NPN_IdentityInfoList_r16)
				if err = ie.Npn_IdentityInfoList_r16.Decode(extReader); err != nil {
					return utils.WrapError("Decode Npn_IdentityInfoList_r16", err)
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

			CellReservedForOtherUse_r16Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			// decode CellReservedForOtherUse_r16 optional
			if CellReservedForOtherUse_r16Present {
				ie.CellReservedForOtherUse_r16 = new(CGI_InfoNR_cellReservedForOtherUse_r16)
				if err = ie.CellReservedForOtherUse_r16.Decode(extReader); err != nil {
					return utils.WrapError("Decode CellReservedForOtherUse_r16", err)
				}
			}
		}
	}
	return nil
}
