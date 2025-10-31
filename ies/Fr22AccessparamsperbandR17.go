package ies

// FR2-2-AccessParamsPerBand-r17 ::= SEQUENCE
// Extensible
type Fr22AccessparamsperbandR17 struct {
	DlFr22Scs120khzR17                  *Fr22AccessparamsperbandR17DlFr22Scs120khzR17
	UlFr22Scs120khzR17                  *Fr22AccessparamsperbandR17UlFr22Scs120khzR17
	Initialaccessssb120khzR17           *Fr22AccessparamsperbandR17Initialaccessssb120khzR17
	WidebandprachScs120khzR17           *Fr22AccessparamsperbandR17WidebandprachScs120khzR17
	MultirbPucchScs120khzR17            *Fr22AccessparamsperbandR17MultirbPucchScs120khzR17
	MultipdschSingledciFr22Scs120khzR17 *Fr22AccessparamsperbandR17MultipdschSingledciFr22Scs120khzR17
	MultipuschSingledciFr22Scs120khzR17 *Fr22AccessparamsperbandR17MultipuschSingledciFr22Scs120khzR17
	DlFr22Scs480khzR17                  *Fr22AccessparamsperbandR17DlFr22Scs480khzR17
	UlFr22Scs480khzR17                  *Fr22AccessparamsperbandR17UlFr22Scs480khzR17
	Initialaccessssb480khzR17           *Fr22AccessparamsperbandR17Initialaccessssb480khzR17
	WidebandprachScs480khzR17           *Fr22AccessparamsperbandR17WidebandprachScs480khzR17
	MultirbPucchScs480khzR17            *Fr22AccessparamsperbandR17MultirbPucchScs480khzR17
	EnhancedpdcchMonitoringscs480khzR17 *Fr22AccessparamsperbandR17EnhancedpdcchMonitoringscs480khzR17
	DlFr22Scs960khzR17                  *Fr22AccessparamsperbandR17DlFr22Scs960khzR17
	UlFr22Scs960khzR17                  *Fr22AccessparamsperbandR17UlFr22Scs960khzR17
	MultirbPucchScs960khzR17            *Fr22AccessparamsperbandR17MultirbPucchScs960khzR17
	EnhancedpdcchMonitoringscs960khzR17 *Fr22AccessparamsperbandR17EnhancedpdcchMonitoringscs960khzR17
	Type1ChannelaccessFr22R17           *Fr22AccessparamsperbandR17Type1ChannelaccessFr22R17
	Type2ChannelaccessFr22R17           *Fr22AccessparamsperbandR17Type2ChannelaccessFr22R17
	ReducedBeamswitchtimingFr22R17      *Fr22AccessparamsperbandR17ReducedBeamswitchtimingFr22R17
	Support32DlHarqProcessperscsR17     *Fr22AccessparamsperbandR17Support32DlHarqProcessperscsR17
	Support32UlHarqProcessperscsR17     *Fr22AccessparamsperbandR17Support32UlHarqProcessperscsR17
	Modulation64QamPuschFr22R17         *Fr22AccessparamsperbandR17Modulation64QamPuschFr22R17 `ext`
}
