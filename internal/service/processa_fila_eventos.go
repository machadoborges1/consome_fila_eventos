package service

import (
	"fmt"

	"github.com/machadoborges1/consome_fila_eventos/internal/entity"
)

func ProcessaItem(evento *entity.TCBContrFilaEventos) error {
	switch evento.TipoAcao.String {
	case "BDT_FAT_GR":
		processaBDTFatGR(evento)

	case "BDT_FAT_PG":
		processaBDTFatPG()

	case "BDT_FAT_CL":
		processaBDTFatCL()

	case "BDT_FAT_NE":
		processaBDTFatNE()

	case "BDT_BCO":
		processaBDTBco()

	case "BDT_DEL_BCO":
		processaBDTDelBco()

	case "BDT_CXA_DEP":
		processaBDTCxaDep()

	case "BDT_CXA":
		processaBDTCxa()

	case "LOG_FAT":
		processaLogFat()

	case "FATURAMENTO":
		processaFaturamento()

	case "ATU_REA":
		processaAtuRea()

	case "RECEBIMENTO":
		processaRecebimento()

	case "CONF_FAT":
		processaConfFat()

	case "CONF_REA":
		processaConfRea()

	default:
		return fmt.Errorf("Tipo de ação desconhecido: %s", evento.TipoAcao.String)
	}
	return nil
}
