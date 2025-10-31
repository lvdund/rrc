package ies

// MAC-Parameters-v1430 ::= SEQUENCE
type MacParametersV1430 struct {
	ShortspsIntervalfddR14 *MacParametersV1430ShortspsIntervalfddR14
	ShortspsIntervaltddR14 *MacParametersV1430ShortspsIntervaltddR14
	SkipuplinkdynamicR14   *MacParametersV1430SkipuplinkdynamicR14
	SkipuplinkspsR14       *MacParametersV1430SkipuplinkspsR14
	MultipleuplinkspsR14   *MacParametersV1430MultipleuplinkspsR14
	DatainactmonR14        *MacParametersV1430DatainactmonR14
}
