package ies

import (
	"bytes"

	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type CellAccessRelatedInfo struct {
	Plmn_IdentityInfoList        PLMN_IdentityInfoList                               `madatory`
	CellReservedForOtherUse      *CellAccessRelatedInfo_cellReservedForOtherUse      `optional`
	CellReservedForFutureUse_r16 *CellAccessRelatedInfo_cellReservedForFutureUse_r16 `optional,ext-1`
	Npn_IdentityInfoList_r16     *NPN_IdentityInfoList_r16                           `optional,ext-1`
	Snpn_AccessInfoList_r17      []SNPN_AccessInfo_r17                               `lb:1,ub:maxNPN_r16,optional,ext-2`
}

func (ie *CellAccessRelatedInfo) Encode(w *uper.UperWriter) error {
	var err error
	hasExtensions := ie.CellReservedForFutureUse_r16 != nil || ie.Npn_IdentityInfoList_r16 != nil || len(ie.Snpn_AccessInfoList_r17) > 0
	preambleBits := []bool{hasExtensions, ie.CellReservedForOtherUse != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if err = ie.Plmn_IdentityInfoList.Encode(w); err != nil {
		return utils.WrapError("Encode Plmn_IdentityInfoList", err)
	}
	if ie.CellReservedForOtherUse != nil {
		if err = ie.CellReservedForOtherUse.Encode(w); err != nil {
			return utils.WrapError("Encode CellReservedForOtherUse", err)
		}
	}
	if hasExtensions {
		// Extension bitmap: 2 bits for 2 extension groups
		extBitmap := []bool{ie.CellReservedForFutureUse_r16 != nil || ie.Npn_IdentityInfoList_r16 != nil, len(ie.Snpn_AccessInfoList_r17) > 0}
		if err := w.WriteExtBitMap(extBitmap); err != nil {
			return utils.WrapError("WriteExtBitMap CellAccessRelatedInfo", err)
		}

		// encode extension group 1
		if extBitmap[0] {
			extBuf := new(bytes.Buffer)
			extWriter := uper.NewWriter(extBuf)

			// Write preamble bits for optional fields in extension group 1
			optionals_ext_1 := []bool{ie.CellReservedForFutureUse_r16 != nil, ie.Npn_IdentityInfoList_r16 != nil}
			for _, bit := range optionals_ext_1 {
				if err := extWriter.WriteBool(bit); err != nil {
					return err
				}
			}

			// encode CellReservedForFutureUse_r16 optional
			if ie.CellReservedForFutureUse_r16 != nil {
				if err = ie.CellReservedForFutureUse_r16.Encode(extWriter); err != nil {
					return utils.WrapError("Encode CellReservedForFutureUse_r16", err)
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
			optionals_ext_2 := []bool{len(ie.Snpn_AccessInfoList_r17) > 0}
			for _, bit := range optionals_ext_2 {
				if err := extWriter.WriteBool(bit); err != nil {
					return err
				}
			}

			// encode Snpn_AccessInfoList_r17 optional
			if len(ie.Snpn_AccessInfoList_r17) > 0 {
				tmp_Snpn_AccessInfoList_r17 := utils.NewSequence[*SNPN_AccessInfo_r17]([]*SNPN_AccessInfo_r17{}, uper.Constraint{Lb: 1, Ub: maxNPN_r16}, false)
				for _, i := range ie.Snpn_AccessInfoList_r17 {
					tmp_Snpn_AccessInfoList_r17.Value = append(tmp_Snpn_AccessInfoList_r17.Value, &i)
				}
				if err = tmp_Snpn_AccessInfoList_r17.Encode(extWriter); err != nil {
					return utils.WrapError("Encode Snpn_AccessInfoList_r17", err)
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

func (ie *CellAccessRelatedInfo) Decode(r *uper.UperReader) error {
	var err error
	var extensionBit bool
	if extensionBit, err = r.ReadBool(); err != nil {
		return err
	}
	var CellReservedForOtherUsePresent bool
	if CellReservedForOtherUsePresent, err = r.ReadBool(); err != nil {
		return err
	}
	if err = ie.Plmn_IdentityInfoList.Decode(r); err != nil {
		return utils.WrapError("Decode Plmn_IdentityInfoList", err)
	}
	if CellReservedForOtherUsePresent {
		ie.CellReservedForOtherUse = new(CellAccessRelatedInfo_cellReservedForOtherUse)
		if err = ie.CellReservedForOtherUse.Decode(r); err != nil {
			return utils.WrapError("Decode CellReservedForOtherUse", err)
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

			CellReservedForFutureUse_r16Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			Npn_IdentityInfoList_r16Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			// decode CellReservedForFutureUse_r16 optional
			if CellReservedForFutureUse_r16Present {
				ie.CellReservedForFutureUse_r16 = new(CellAccessRelatedInfo_cellReservedForFutureUse_r16)
				if err = ie.CellReservedForFutureUse_r16.Decode(extReader); err != nil {
					return utils.WrapError("Decode CellReservedForFutureUse_r16", err)
				}
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

			Snpn_AccessInfoList_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			// decode Snpn_AccessInfoList_r17 optional
			if Snpn_AccessInfoList_r17Present {
				tmp_Snpn_AccessInfoList_r17 := utils.NewSequence[*SNPN_AccessInfo_r17]([]*SNPN_AccessInfo_r17{}, uper.Constraint{Lb: 1, Ub: maxNPN_r16}, false)
				fn_Snpn_AccessInfoList_r17 := func() *SNPN_AccessInfo_r17 {
					return new(SNPN_AccessInfo_r17)
				}
				if err = tmp_Snpn_AccessInfoList_r17.Decode(extReader, fn_Snpn_AccessInfoList_r17); err != nil {
					return utils.WrapError("Decode Snpn_AccessInfoList_r17", err)
				}
				ie.Snpn_AccessInfoList_r17 = []SNPN_AccessInfo_r17{}
				for _, i := range tmp_Snpn_AccessInfoList_r17.Value {
					ie.Snpn_AccessInfoList_r17 = append(ie.Snpn_AccessInfoList_r17, *i)
				}
			}
		}
	}
	return nil
}
