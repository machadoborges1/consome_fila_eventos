package service

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/machadoborges1/consome_fila_eventos/internal/entity"
)

func TravaBDTFatGR(db *sql.DB, evento entity.TCBContrFilaEventos) {

	vTipoAcao := "BDT_FAT_GR"

	var vAudSID int64
	roww := db.QueryRow("SELECT USERENV('SESSIONID') FROM DUAL")
	if err := roww.Scan(&vAudSID); err != nil {
		fmt.Println("erro3")
		log.Fatal(err)
	}

	var vDt_Atual string
	row := db.QueryRow("SELECT TO_CHAR(SYSDATE, 'YYYYMMDD') FROM DUAL")
	if err := row.Scan(&vDt_Atual); err != nil {
		fmt.Println("erro4")
		log.Fatal(err)
	}

	var fsGetIdAlt int64
	rowww := db.QueryRow("SELECT FS_GET_ID_ALT FROM DUAL")
	if err := rowww.Scan(&fsGetIdAlt); err != nil {
		log.Fatal(err)
	}
	fmt.Println(fsGetIdAlt)

	stmt, err := db.Prepare(`
        UPDATE TCB_CONTR_FILA_EVENTOS
        SET STATUS = 'P'
           ,AUDSID = :1
           ,DTH_INICIO_PROCESSAMENTO = :2
        WHERE STATUS = 'A'
          AND TIPO_ACAO = :3
          AND ID_EVENTO <= :4
          AND COD_PERIODO   = :5
          AND COD_PESSOA    = :6
          AND COD_FIP_GF    = :7
          AND COD_GRUPO_FIN = :8
          AND COD_SERVICO   = :9
          AND COD_PARCELA   = :10
          AND NOT EXISTS (
              SELECT *
              FROM TCB_CONTR_FILA_EVENTOS
              WHERE STATUS = 'P'
                AND AUDSID IS NOT NULL
                AND AUDSID <> :11
                AND TIPO_ACAO = :12
                AND COD_PERIODO   = :13
                AND COD_PESSOA    = :14
                AND COD_FIP_GF    = :15
                AND COD_GRUPO_FIN = :16
                AND COD_SERVICO   = :17
                AND COD_PARCELA   = :18
          )
    `)
	if err != nil {
		fmt.Println("erro1")
		log.Fatal(err)
	}
	defer stmt.Close()

	_, err = stmt.Exec(
		vAudSID, vDt_Atual, vTipoAcao,
		fsGetIdAlt, evento.CodPeriodo.String, evento.CodPessoa.Int64, evento.CodFipGf.Int64, evento.CodGrupoFin.String, evento.CodServico.Int64,
		evento.CodParcela.String, vAudSID, vTipoAcao,
		evento.CodPeriodo.String, evento.CodPessoa.Int64, evento.CodFipGf.Int64, evento.CodGrupoFin.String, evento.CodServico.Int64, evento.CodParcela.String,
		vAudSID, // Adicionando valor para :18
	)

	if err != nil {
		fmt.Println("erro2")
		log.Fatal(err)
	}

}
