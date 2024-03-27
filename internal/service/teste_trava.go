package service

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/machadoborges1/consome_fila_eventos/internal/entity"
)

func TravaBDTFatGRr(db *sql.DB, evento entity.TCBContrFilaEventos, audSID int64, fsGetIdAlt int64) {
	vTipoAcao := "BDT_FAT_GR"

	stmt, err := db.Prepare("UPDATE TCB_CONTR_FILA_EVENTOS " +
		"SET NRO_ITERACOES = NRO_ITERACOES + 1 " +
		"WHERE STATUS = 'A' " +
		"AND TIPO_ACAO = :1 " +
		"AND ID_EVENTO <= :2 " +
		"AND COD_PERIODO = :3 " +
		"AND COD_PESSOA = :4 " +
		"AND COD_FIP_GF = :5 " +
		"AND COD_GRUPO_FIN = :6 " +
		"AND COD_SERVICO = :7 " +
		"AND COD_PARCELA = :8")
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()

	resultado, err := stmt.Exec(vTipoAcao, fsGetIdAlt, evento.CodPeriodo.String, evento.CodPessoa.Int64,
		evento.CodFipGf.Int64, evento.CodGrupoFin.String, evento.CodServico.Int64, evento.CodParcela.String)
	if err != nil {
		log.Fatal(err)
	}
	rowsAffected, err := resultado.RowsAffected()
	if rowsAffected > 0 {
		fmt.Println(rowsAffected)
		fmt.Println("-----------------------------------------------------------------")
	}

}
