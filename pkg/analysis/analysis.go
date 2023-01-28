package analysis

import (
	"math"
	"math/cmplx"
	"strconv"
)

const (
	frequencia_da_rede = 60
	H                  = 5.0
	V1                 = 1.05
	V2                 = 1.0
	potencia_media     = 0.8
)

var systemData = make(map[string]float64)

func SystemAnalysis(X_1_2_pre, X_1_2_pos, Xd string) map[string]float64 {
	X_1_2_pre_falta, _ := strconv.ParseComplex(X_1_2_pre, 64)
	X_1_2_pos_falta, _ := strconv.ParseComplex(X_1_2_pos, 64)
	Xd_gerador, _ := strconv.ParseComplex(Xd, 64)

	var Ws float64 = 2 * math.Pi * frequencia_da_rede
	var X_equivalente_pre_falta float64 = cmplx.Abs(X_1_2_pre_falta + Xd_gerador)
	var X_equivalente_pos_falta float64 = cmplx.Abs(X_1_2_pos_falta + Xd_gerador)

	systemData["angulo_tensao_barra_curto"] = math.Asin((potencia_media * cmplx.Abs(X_1_2_pre_falta)) / (V1 * V2))

	corrente_saida_gerador := ((cmplx.Rect(V1, systemData["angulo_tensao_barra_curto"]) - complex(V2, 0)) / X_1_2_pre_falta)
	systemData["modulo_corrente_de_saida_gerador"], systemData["angulo_corrente_de_saida_gerador"] = cmplx.Polar(corrente_saida_gerador)
	systemData["modulo_tensao_de_saida_gerador"], systemData["angulo_tensao_de_saida_gerador"] = cmplx.Polar(cmplx.Rect(V1, systemData["angulo_tensao_barra_curto"]) + Xd_gerador*corrente_saida_gerador)

	systemData["potencia_media_pre_barra_gerador"] = systemData["modulo_tensao_de_saida_gerador"] * V1 / cmplx.Abs(Xd_gerador)
	systemData["potencia_media_pre_barra_curto"] = systemData["modulo_tensao_de_saida_gerador"] * V2 / X_equivalente_pre_falta

	systemData["potencia_media_pos_barra_curto"] = systemData["modulo_tensao_de_saida_gerador"] * V2 / X_equivalente_pos_falta

	systemData["angulo_delta_maximo"] = math.Pi - math.Asin(potencia_media/systemData["potencia_media_pos_barra_curto"])

	systemData["angulo_delta_critico"] = math.Acos(
		(potencia_media*(systemData["angulo_delta_maximo"]-systemData["angulo_tensao_de_saida_gerador"]) +
			systemData["potencia_media_pos_barra_curto"]*math.Cos(systemData["angulo_delta_maximo"])) / systemData["potencia_media_pos_barra_curto"])

	systemData["tempo_maximo"] = math.Sqrt((4 * H * (systemData["angulo_delta_critico"] - systemData["angulo_tensao_de_saida_gerador"])) / (Ws * potencia_media))

	return systemData
}
