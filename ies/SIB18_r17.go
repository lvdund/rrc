package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

type SIB18_r17 struct {
	Gin_ElementList_r17      []GIN_Element_r17  `lb:1,ub:maxGIN_r17,optional`
	Gins_PerSNPN_List_r17    []GINs_PerSNPN_r17 `lb:1,ub:maxNPN_r16,optional`
	LateNonCriticalExtension *[]byte            `optional`
}

func (ie *SIB18_r17) Encode(w *aper.AperWriter) error {
	var err error
	preambleBits := []bool{len(ie.Gin_ElementList_r17) > 0, len(ie.Gins_PerSNPN_List_r17) > 0, ie.LateNonCriticalExtension != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if len(ie.Gin_ElementList_r17) > 0 {
		tmp_Gin_ElementList_r17 := utils.NewSequence[*GIN_Element_r17]([]*GIN_Element_r17{}, aper.Constraint{Lb: 1, Ub: maxGIN_r17}, false)
		for _, i := range ie.Gin_ElementList_r17 {
			tmp_Gin_ElementList_r17.Value = append(tmp_Gin_ElementList_r17.Value, &i)
		}
		if err = tmp_Gin_ElementList_r17.Encode(w); err != nil {
			return utils.WrapError("Encode Gin_ElementList_r17", err)
		}
	}
	if len(ie.Gins_PerSNPN_List_r17) > 0 {
		tmp_Gins_PerSNPN_List_r17 := utils.NewSequence[*GINs_PerSNPN_r17]([]*GINs_PerSNPN_r17{}, aper.Constraint{Lb: 1, Ub: maxNPN_r16}, false)
		for _, i := range ie.Gins_PerSNPN_List_r17 {
			tmp_Gins_PerSNPN_List_r17.Value = append(tmp_Gins_PerSNPN_List_r17.Value, &i)
		}
		if err = tmp_Gins_PerSNPN_List_r17.Encode(w); err != nil {
			return utils.WrapError("Encode Gins_PerSNPN_List_r17", err)
		}
	}
	if ie.LateNonCriticalExtension != nil {
		if err = w.WriteOctetString(*ie.LateNonCriticalExtension, nil, false); err != nil {
			return utils.WrapError("Encode LateNonCriticalExtension", err)
		}
	}
	return nil
}

func (ie *SIB18_r17) Decode(r *aper.AperReader) error {
	var err error
	var Gin_ElementList_r17Present bool
	if Gin_ElementList_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var Gins_PerSNPN_List_r17Present bool
	if Gins_PerSNPN_List_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var LateNonCriticalExtensionPresent bool
	if LateNonCriticalExtensionPresent, err = r.ReadBool(); err != nil {
		return err
	}
	if Gin_ElementList_r17Present {
		tmp_Gin_ElementList_r17 := utils.NewSequence[*GIN_Element_r17]([]*GIN_Element_r17{}, aper.Constraint{Lb: 1, Ub: maxGIN_r17}, false)
		fn_Gin_ElementList_r17 := func() *GIN_Element_r17 {
			return new(GIN_Element_r17)
		}
		if err = tmp_Gin_ElementList_r17.Decode(r, fn_Gin_ElementList_r17); err != nil {
			return utils.WrapError("Decode Gin_ElementList_r17", err)
		}
		ie.Gin_ElementList_r17 = []GIN_Element_r17{}
		for _, i := range tmp_Gin_ElementList_r17.Value {
			ie.Gin_ElementList_r17 = append(ie.Gin_ElementList_r17, *i)
		}
	}
	if Gins_PerSNPN_List_r17Present {
		tmp_Gins_PerSNPN_List_r17 := utils.NewSequence[*GINs_PerSNPN_r17]([]*GINs_PerSNPN_r17{}, aper.Constraint{Lb: 1, Ub: maxNPN_r16}, false)
		fn_Gins_PerSNPN_List_r17 := func() *GINs_PerSNPN_r17 {
			return new(GINs_PerSNPN_r17)
		}
		if err = tmp_Gins_PerSNPN_List_r17.Decode(r, fn_Gins_PerSNPN_List_r17); err != nil {
			return utils.WrapError("Decode Gins_PerSNPN_List_r17", err)
		}
		ie.Gins_PerSNPN_List_r17 = []GINs_PerSNPN_r17{}
		for _, i := range tmp_Gins_PerSNPN_List_r17.Value {
			ie.Gins_PerSNPN_List_r17 = append(ie.Gins_PerSNPN_List_r17, *i)
		}
	}
	if LateNonCriticalExtensionPresent {
		var tmp_os_LateNonCriticalExtension []byte
		if tmp_os_LateNonCriticalExtension, err = r.ReadOctetString(nil, false); err != nil {
			return utils.WrapError("Decode LateNonCriticalExtension", err)
		}
		ie.LateNonCriticalExtension = &tmp_os_LateNonCriticalExtension
	}
	return nil
}
