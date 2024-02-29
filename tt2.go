package main

import (
	"database/sql"

	"github.com/machadoborges1/consome_fila_eventos/internal/entity"
)

func travaBDTFaatGR(db *sql.DB, vAudSID int64, vIdAtual int64, vDados entity.TCBContrFilaEventos) (bool, error) {
	vTipoAcao := "BDT_FAT_GR"

	tx, err := db.Begin() // Inicia uma transação
	if err != nil {
		return false, err
	}

	defer func() {
		if err != nil {
			tx.Rollback() // Reverte a transação se ocorrer um erro
		}
	}()

	_, err = tx.Exec(`
		UPDATE TCB_CONTR_FILA_EVENTOS
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
			)
	`, vAudSID, vTipoAcao, vIdAtual, vDados.CodPeriodo, vDados.CodPessoa, vDados.CodFipGf, vDados.CodGrupoFin, vDados.CodServico, vDados.CodParcela)
	if err != nil {
		return false, err
	}

	var rowCount int
	err = db.QueryRow(`
		SELECT COUNT(*)
		FROM TCB_CONTR_FILA_EVENTOS
		WHERE STATUS = 'P'
			AND TIPO_ACAO = :1
			AND ID_EVENTO <= :2
			AND COD_PERIODO = :3
			AND COD_PESSOA = :4
			AND COD_FIP_GF = :5
			AND COD_GRUPO_FIN = :6
			AND COD_SERVICO = :7
			AND COD_PARCELA = :8
	`, vTipoAcao, vIdAtual, vDados.CodPeriodo, vDados.CodPessoa, vDados.CodFipGf, vDados.CodGrupoFin, vDados.CodServico, vDados.CodParcela).Scan(&rowCount)
	if err != nil {
		return false, err
	}

	if rowCount > 0 {
		err = tx.Commit() // Faz o commit da transação
		if err != nil {
			return false, err
		}
		return true, nil
	}

	_, err = tx.Exec(`
		UPDATE TCB_CONTR_FILA_EVENTOS
		SET NRO_ITERACOES = NRO_ITERACOES + 1
		WHERE STATUS = 'A'
			AND TIPO_ACAO = :1
			AND ID_EVENTO <= :2
			AND COD_PERIODO = :3
			AND COD_PESSOA = :4
			AND COD_FIP_GF = :5
			AND COD_GRUPO_FIN = :6
			AND COD_SERVICO = :7
			AND COD_PARCELA = :8
	`, vTipoAcao, vIdAtual, vDados.CodPeriodo, vDados.CodPessoa, vDados.CodFipGf, vDados.CodGrupoFin, vDados.CodServico, vDados.CodParcela)
	if err != nil {
		return false, err
	}

	err = tx.Commit() // Faz o commit da transação
	if err != nil {
		return false, err
	}

	return false, nil
}
