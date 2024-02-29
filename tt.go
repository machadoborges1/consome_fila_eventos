package main

import (
	"database/sql"
	"log"
)

func travaBDTFatGR(db *sql.DB, vAudSID int64, vIdAtual int64, vDados TCBContrFilaEventos) (bool, error) {
	vTipoAcao := "BDT_FAT_GR"

	_, err := db.Exec(`UPDATE TCB_CONTR_FILA_EVENTOS
	SET STATUS = 'P',
		AUDSID = :1,
		DTH_INICIO_PROCESSAMENTO = SYSDATE
	WHERE STATUS = 'A'
	  AND TIPO_ACAO = :2
	  AND ID_EVENTO <= :3
	  AND COD_PERIODO = :4
	  AND COD_PESSOA = :5
	  AND COD_FIP_GF = :6
	  AND COD_GRUPO_FIN = :7
	  AND COD_SERVICO = :8
	  AND COD_PARCELA = :9
	  AND NOT EXISTS (
		SELECT *
		FROM TCB_CONTR_FILA_EVENTOS
		WHERE STATUS = 'P'
		  AND AUDSID IS NOT NULL
		  AND AUDSID <> :1
		  AND TIPO_ACAO = :2
		  AND COD_PERIODO = :4
		  AND COD_PESSOA = :5
		  AND COD_FIP_GF = :6
		  AND COD_GRUPO_FIN = :7
		  AND COD_SERVICO = :8
		  AND COD_PARCELA = :9
	  )`, vAudSID, vTipoAcao, vIdAtual, vDados.CodPeriodo, vDados.CodPessoa, vDados.CodFipGf, vDados.CodGrupoFin, vDados.CodServico, vDados.CodParcela)
	if err != nil {
		return false, err
	}

	rowCount, err := db.Exec(`SELECT COUNT(*) FROM TCB_CONTR_FILA_EVENTOS WHERE STATUS = 'P' AND TIPO_ACAO = :1 AND ID_EVENTO <= :2 AND COD_PERIODO = :3 AND COD_PESSOA = :4 AND COD_FIP_GF = :5 AND COD_GRUPO_FIN = :6 AND COD_SERVICO = :7 AND COD_PARCELA = :8`, vTipoAcao, vIdAtual, vDados.CodPeriodo, vDados.CodPessoa, vDados.CodFipGf, vDados.CodGrupoFin, vDados.CodServico, vDados.CodParcela)
	if err != nil {
		return false, err
	}

	var rowCountInt int64
	err = rowCount.Scan(&rowCountInt)
	if err != nil {
		return false, err
	}

	var retorno bool
	if rowCountInt > 0 {
		err := db.Commit()
		if err != nil {
			return false, err
		}
		retorno = true
	} else {
		err := db.Rollback()
		if err != nil {
			return false, err
		}

		_, err = db.Exec(`UPDATE TCB_CONTR_FILA_EVENTOS
		SET NRO_ITERACOES = NRO_ITERACOES + 1
		WHERE STATUS = 'A'
		  AND TIPO_ACAO = :1
		  AND ID_EVENTO <= :2
		  AND COD_PERIODO = :3
		  AND COD_PESSOA = :4
		  AND COD_FIP_GF = :5
		  AND COD_GRUPO_FIN = :6
		  AND COD_SERVICO = :7
		  AND COD_PARCELA = :8`, vTipoAcao, vIdAtual, vDados.CodPeriodo, vDados.CodPessoa, vDados.CodFipGf, vDados.CodGrupoFin, vDados.CodServico, vDados.CodParcela)
		if err != nil {
			return false, err
		}

		err = db.Commit()
		if err != nil {
			return false, err
		}

		retorno = false
	}

	return retorno, nil
}
