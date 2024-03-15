package service

import (
	"database/sql"
	"fmt"

	"github.com/machadoborges1/consome_fila_eventos/internal/entity"
)

func ProcessaItem(db *sql.DB, evento entity.TCBContrFilaEventos) {
	switch evento.TipoAcao.String {
	case "BDT_FAT_GR":
		ProcessaBDTFatGR(db, evento)

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
		fmt.Printf("Tipo de ação desconhecido:")
	}
}
