package service

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/machadoborges1/consome_fila_eventos/internal/entity"
)

func TravaBDTFatGR(db *sql.DB, evento entity.TCBContrFilaEventos, audSID int64, fsGetIdAlt int64) {
	vTipoAcao := "BDT_FAT_GR"

	tx, err := db.Begin()
	if err != nil {
		fmt.Println("Erro ao iniciar a transação:", err)
		return
	}

	stmt, err := db.Prepare(`
        UPDATE TCB_CONTR_FILA_EVENTOS
        SET STATUS = 'P'
           ,AUDSID = :1
           ,DTH_INICIO_PROCESSAMENTO = SYSDATE
        WHERE STATUS = 'A'
          AND TIPO_ACAO = :2
          AND ID_EVENTO <= :3
          AND COD_PERIODO   = :4
          AND COD_PESSOA    = :5
          AND COD_FIP_GF    = :6
          AND COD_GRUPO_FIN = :7
          AND COD_SERVICO   = :8
          AND COD_PARCELA   = :9
          AND NOT EXISTS (
              SELECT *
              FROM TCB_CONTR_FILA_EVENTOS
              WHERE STATUS = 'P'
                AND AUDSID IS NOT NULL
                AND AUDSID <> :10
                AND TIPO_ACAO = :11
                AND COD_PERIODO   = :12
                AND COD_PESSOA    = :13
                AND COD_FIP_GF    = :14
                AND COD_GRUPO_FIN = :15
                AND COD_SERVICO   = :16
                AND COD_PARCELA   = :17
          )
    `)
	if err != nil {
		fmt.Println("erro1")
		log.Fatal(err)
	}
	defer stmt.Close()

    fmt.Println(		audSID, vTipoAcao,
		fsGetIdAlt,
		evento.CodPeriodo.String,
		evento.CodPessoa.Int64,
		evento.CodFipGf.Int64,
		evento.CodGrupoFin.String,
		evento.CodServico.Int64,
		evento.CodParcela.String,
		audSID, vTipoAcao,
		evento.CodPeriodo.String,
		evento.CodPessoa.Int64,
		evento.CodFipGf.Int64,
		evento.CodGrupoFin.String,
		evento.CodServico.Int64,
		evento.CodParcela.String,)

	result, err := stmt.Exec(
		audSID, vTipoAcao,
		fsGetIdAlt,
		evento.CodPeriodo.String,
		evento.CodPessoa.Int64,
		evento.CodFipGf.Int64,
		evento.CodGrupoFin.String,
		evento.CodServico.Int64,
		evento.CodParcela.String,
		audSID, vTipoAcao,
		evento.CodPeriodo.String,
		evento.CodPessoa.Int64,
		evento.CodFipGf.Int64,
		evento.CodGrupoFin.String,
		evento.CodServico.Int64,
		evento.CodParcela.String,
	)
	if err != nil {
		fmt.Println("erro2")
		log.Fatal(err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		log.Fatal(err)
	}

	if rowsAffected > 0 {
		err = tx.Commit()
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("Commit realizado. Linhas afetadas:", rowsAffected)
	} else {
		err = tx.Rollback()
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("Rollback realizado. Nenhuma linha afetada.")

        fmt.Println(vTipoAcao, fsGetIdAlt, evento.CodPeriodo.String, evento.CodPessoa.Int64,
			evento.CodFipGf.Int64, evento.CodGrupoFin.String, evento.CodServico.Int64, evento.CodParcela.String)

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

		_, err = stmt.Exec(vTipoAcao, fsGetIdAlt, evento.CodPeriodo.String, evento.CodPessoa.Int64,
			evento.CodFipGf.Int64, evento.CodGrupoFin.String, evento.CodServico.Int64, evento.CodParcela.String)
		if err != nil {
			log.Fatal(err)
		}
	}
}
