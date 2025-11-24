package ies

import (
	"bytes"

	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type PLMN_IdentityInfo struct {
	Plmn_IdentityList          []PLMN_Identity                              `lb:1,ub:maxPLMN,madatory`
	TrackingAreaCode           *TrackingAreaCode                            `optional`
	Ranac                      *RAN_AreaCode                                `optional`
	CellIdentity               CellIdentity                                 `madatory`
	CellReservedForOperatorUse PLMN_IdentityInfo_cellReservedForOperatorUse `madatory`
	Iab_Support_r16            *PLMN_IdentityInfo_iab_Support_r16           `optional,ext-1`
	TrackingAreaList_r17       []TrackingAreaCode                           `lb:1,ub:maxTAC_r17,optional,ext-2`
	GNB_ID_Length_r17          *int64                                       `lb:22,ub:32,optional,ext-2`
}

func (ie *PLMN_IdentityInfo) Encode(w *uper.UperWriter) error {
	var err error
	hasExtensions := ie.Iab_Support_r16 != nil || len(ie.TrackingAreaList_r17) > 0 || ie.GNB_ID_Length_r17 != nil
	preambleBits := []bool{hasExtensions, ie.TrackingAreaCode != nil, ie.Ranac != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	tmp_Plmn_IdentityList := utils.NewSequence[*PLMN_Identity]([]*PLMN_Identity{}, uper.Constraint{Lb: 1, Ub: maxPLMN}, false)
	for _, i := range ie.Plmn_IdentityList {
		tmp_Plmn_IdentityList.Value = append(tmp_Plmn_IdentityList.Value, &i)
	}
	if err = tmp_Plmn_IdentityList.Encode(w); err != nil {
		return utils.WrapError("Encode Plmn_IdentityList", err)
	}
	if ie.TrackingAreaCode != nil {
		if err = ie.TrackingAreaCode.Encode(w); err != nil {
			return utils.WrapError("Encode TrackingAreaCode", err)
		}
	}
	if ie.Ranac != nil {
		if err = ie.Ranac.Encode(w); err != nil {
			return utils.WrapError("Encode Ranac", err)
		}
	}
	if err = ie.CellIdentity.Encode(w); err != nil {
		return utils.WrapError("Encode CellIdentity", err)
	}
	if err = ie.CellReservedForOperatorUse.Encode(w); err != nil {
		return utils.WrapError("Encode CellReservedForOperatorUse", err)
	}
	if hasExtensions {
		// Extension bitmap: 2 bits for 2 extension groups
		extBitmap := []bool{ie.Iab_Support_r16 != nil, len(ie.TrackingAreaList_r17) > 0 || ie.GNB_ID_Length_r17 != nil}
		if err := w.WriteExtBitMap(extBitmap); err != nil {
			return utils.WrapError("WriteExtBitMap PLMN_IdentityInfo", err)
		}

		// encode extension group 1
		if extBitmap[0] {
			extBuf := new(bytes.Buffer)
			extWriter := uper.NewWriter(extBuf)

			// Write preamble bits for optional fields in extension group 1
			optionals_ext_1 := []bool{ie.Iab_Support_r16 != nil}
			for _, bit := range optionals_ext_1 {
				if err := extWriter.WriteBool(bit); err != nil {
					return err
				}
			}

			// encode Iab_Support_r16 optional
			if ie.Iab_Support_r16 != nil {
				if err = ie.Iab_Support_r16.Encode(extWriter); err != nil {
					return utils.WrapError("Encode Iab_Support_r16", err)
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
			optionals_ext_2 := []bool{len(ie.TrackingAreaList_r17) > 0, ie.GNB_ID_Length_r17 != nil}
			for _, bit := range optionals_ext_2 {
				if err := extWriter.WriteBool(bit); err != nil {
					return err
				}
			}

			// encode TrackingAreaList_r17 optional
			if len(ie.TrackingAreaList_r17) > 0 {
				tmp_TrackingAreaList_r17 := utils.NewSequence[*TrackingAreaCode]([]*TrackingAreaCode{}, uper.Constraint{Lb: 1, Ub: maxTAC_r17}, false)
				for _, i := range ie.TrackingAreaList_r17 {
					tmp_TrackingAreaList_r17.Value = append(tmp_TrackingAreaList_r17.Value, &i)
				}
				if err = tmp_TrackingAreaList_r17.Encode(extWriter); err != nil {
					return utils.WrapError("Encode TrackingAreaList_r17", err)
				}
			}
			// encode GNB_ID_Length_r17 optional
			if ie.GNB_ID_Length_r17 != nil {
				if err = extWriter.WriteInteger(*ie.GNB_ID_Length_r17, &uper.Constraint{Lb: 22, Ub: 32}, false); err != nil {
					return utils.WrapError("Encode GNB_ID_Length_r17", err)
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

func (ie *PLMN_IdentityInfo) Decode(r *uper.UperReader) error {
	var err error
	var extensionBit bool
	if extensionBit, err = r.ReadBool(); err != nil {
		return err
	}
	var TrackingAreaCodePresent bool
	if TrackingAreaCodePresent, err = r.ReadBool(); err != nil {
		return err
	}
	var RanacPresent bool
	if RanacPresent, err = r.ReadBool(); err != nil {
		return err
	}
	tmp_Plmn_IdentityList := utils.NewSequence[*PLMN_Identity]([]*PLMN_Identity{}, uper.Constraint{Lb: 1, Ub: maxPLMN}, false)
	fn_Plmn_IdentityList := func() *PLMN_Identity {
		return new(PLMN_Identity)
	}
	if err = tmp_Plmn_IdentityList.Decode(r, fn_Plmn_IdentityList); err != nil {
		return utils.WrapError("Decode Plmn_IdentityList", err)
	}
	ie.Plmn_IdentityList = []PLMN_Identity{}
	for _, i := range tmp_Plmn_IdentityList.Value {
		ie.Plmn_IdentityList = append(ie.Plmn_IdentityList, *i)
	}
	if TrackingAreaCodePresent {
		ie.TrackingAreaCode = new(TrackingAreaCode)
		if err = ie.TrackingAreaCode.Decode(r); err != nil {
			return utils.WrapError("Decode TrackingAreaCode", err)
		}
	}
	if RanacPresent {
		ie.Ranac = new(RAN_AreaCode)
		if err = ie.Ranac.Decode(r); err != nil {
			return utils.WrapError("Decode Ranac", err)
		}
	}
	if err = ie.CellIdentity.Decode(r); err != nil {
		return utils.WrapError("Decode CellIdentity", err)
	}
	if err = ie.CellReservedForOperatorUse.Decode(r); err != nil {
		return utils.WrapError("Decode CellReservedForOperatorUse", err)
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

			Iab_Support_r16Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			// decode Iab_Support_r16 optional
			if Iab_Support_r16Present {
				ie.Iab_Support_r16 = new(PLMN_IdentityInfo_iab_Support_r16)
				if err = ie.Iab_Support_r16.Decode(extReader); err != nil {
					return utils.WrapError("Decode Iab_Support_r16", err)
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

			TrackingAreaList_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			GNB_ID_Length_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			// decode TrackingAreaList_r17 optional
			if TrackingAreaList_r17Present {
				tmp_TrackingAreaList_r17 := utils.NewSequence[*TrackingAreaCode]([]*TrackingAreaCode{}, uper.Constraint{Lb: 1, Ub: maxTAC_r17}, false)
				fn_TrackingAreaList_r17 := func() *TrackingAreaCode {
					return new(TrackingAreaCode)
				}
				if err = tmp_TrackingAreaList_r17.Decode(extReader, fn_TrackingAreaList_r17); err != nil {
					return utils.WrapError("Decode TrackingAreaList_r17", err)
				}
				ie.TrackingAreaList_r17 = []TrackingAreaCode{}
				for _, i := range tmp_TrackingAreaList_r17.Value {
					ie.TrackingAreaList_r17 = append(ie.TrackingAreaList_r17, *i)
				}
			}
			// decode GNB_ID_Length_r17 optional
			if GNB_ID_Length_r17Present {
				var tmp_int_GNB_ID_Length_r17 int64
				if tmp_int_GNB_ID_Length_r17, err = extReader.ReadInteger(&uper.Constraint{Lb: 22, Ub: 32}, false); err != nil {
					return utils.WrapError("Decode GNB_ID_Length_r17", err)
				}
				ie.GNB_ID_Length_r17 = &tmp_int_GNB_ID_Length_r17
			}
		}
	}
	return nil
}
