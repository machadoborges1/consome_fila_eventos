package service

import (
	"fmt"
	"database/sql"
)
func processaBDTFatGR(db *sql.DB, evento *entity.TCBContrFilaEventos) {

	var vMensErro string
	var vNroSequencial int64

	fmt.Println("Processando evento tipo BDT_FAT_GR")

	// Defindo a transação para que possa ser revertida em caso de erro
	tx, err := db.Begin()
	if err != nil {
		return err
	}

	var vDt_Atual string
	row := db.QueryRow("SELECT TO_CHAR(SYSDATE, 'YYYYMMDD') FROM DUAL")
	if err := row.Scan(&vDt_Atual); err != nil {
		log.Fatal(err)
	}


	result, err := db.Exec("BEGIN "+
    "PCB_GERA_PREVISTO_MOV(pDt_Base => :1, pCod_Periodo => :2, pCod_Pessoa => :3, "+
    "pCod_Fip_GF => :4, pCod_Grupo_Fin => :5, pCod_Servico => :6, pCod_Parcela => :7, "+
    "pNecessita_Existir => :8, pNro_Sequencial => :9, pMens_Erro => :10); "+
    "END;",
    vDt_Atual, evento.codPeriodo, evento.codPessoa, evento.codFipGf, evento.COD_GRUPO_FIN,
    evento.COD_SERVICO, evento.COD_PARCELA, "N", vNro_Sequencial, vMens_Erro)
	if err != nil {
		log.Fatal(err)
	}

	if vMens_Erro != nil {
		if	vNro_Sequencial == nil{
			tx.Rollback()
			fmt.Println("numero sequencial nulo")
		}
		else {
			_, err := db.Exec("BEGIN FIT.PF_SET_CONTABILIZADO_FICHA_FIN("+
				":pCod_Periodo, :pCod_Pessoa, :pCod_Fip, :pCod_Grupo_Fin, :pCod_Servico, :pCod_Parcela); END;",
				evento.COD_PERIODO, evento.COD_PESSOA, evento.COD_FIP_GF, evento.COD_GRUPO_FIN, evento.COD_SERVICO, evento.COD_PARCELA)
			if err != nil {
				log.Println(err)
			} else { 
				tx.Commit() 
			}
		}
	} else { 
		tx.Rollback() 
	}

	if vMensErro == "" {
		_, err := db.Exec(`DELETE FROM TCB_CONTR_FILA_EVENTOS
			WHERE STATUS IN ('P','E')
			AND TIPO_ACAO = ?
			AND ID_EVENTO <= ?
			AND COD_PERIODO = ?
			AND COD_PESSOA = ?
			AND COD_FIP_GF = ?
			AND COD_GRUPO_FIN = ?
			AND COD_SERVICO = ?
			AND COD_PARCELA = ?`, evento.TipoAcao, evento.IDEvento, evento.CodPeriodo, evento.CodPessoa, evento.CodFipGf, evento.CodGrupoFin, evento.CodServico, evento.CodParcela)
		if err != nil {
			return err
		}

		_, err = db.Exec(`INSERT INTO TCB_CONTR_FILA_EVENTOS
			(ID_EVENTO, TIPO_ACAO, STATUS, DT_BASE, NRO_SEQ_FAT)
			VALUES (?, 'LOG_FAT', 'A', ?, ?)`, fsGetIdAlt, evento.DtBase, vNroSequencial)
		if err != nil {
			return err
		}
	} else {
		_, err := db.Exec(`UPDATE TCB_CONTR_FILA_EVENTOS
			SET STATUS = 'E', MENS_ERRO = ?
			WHERE STATUS = 'P'
			AND TIPO_ACAO = ?
			AND ID_EVENTO <= ?
			AND COD_PERIODO = ?
			AND COD_PESSOA = ?
			AND COD_FIP_GF = ?
			AND COD_GRUPO_FIN = ?
			AND COD_SERVICO = ?
			AND COD_PARCELA = ?`, vMensErro, evento.TipoAcao, evento.IDEvento, evento.CodPeriodo, evento.CodPessoa, evento.CodFipGf, evento.CodGrupoFin, evento.CodServico, evento.CodParcela)
		if err != nil {
			return err
		}
	}



}


func processaBDTFatGR(db *sql.DB, evento *entity.TCBContrFilaEventos) error {
	var vMensErro string
	var vNroSequencial int64

	fmt.Println("Processando evento tipo BDT_FAT_GR")

	// Defina a transação para que possa ser revertida em caso de erro
	tx, err := db.Begin()
	if err != nil {
		return err
	}

	// Chame a função PCB_GERA_PREVISTO_MOV
	err = pcbGeraPrevistoMov(db, evento)
	if err != nil {
		tx.Rollback() // Reverta a transação em caso de erro
		return err
	}

	if vMensErro == "" {
		if vNroSequencial == 0 {
			tx.Rollback()
			return fmt.Errorf("Devolveu Nro_Seq_Fat nulo")
		}

		// Chame a função FIT.PF_SET_CONTABILIZADO_FICHA_FIN
		err = fitPfSetContabilizadoFichaFin(db, evento)
		if err != nil {
			tx.Rollback()
			return err
		}

		tx.Commit() // Commit se não houver erros
	} else {
		tx.Rollback() // Reverta a transação em caso de erro
	}

	// Se não houver erros, execute a exclusão e inserção conforme necessário
	if vMensErro == "" {
		_, err := db.Exec(`DELETE FROM TCB_CONTR_FILA_EVENTOS
			WHERE STATUS IN ('P','E')
			AND TIPO_ACAO = ?
			AND ID_EVENTO <= ?
			AND COD_PERIODO = ?
			AND COD_PESSOA = ?
			AND COD_FIP_GF = ?
			AND COD_GRUPO_FIN = ?
			AND COD_SERVICO = ?
			AND COD_PARCELA = ?`, evento.TipoAcao, evento.IDEvento, evento.CodPeriodo, evento.CodPessoa, evento.CodFipGf, evento.CodGrupoFin, evento.CodServico, evento.CodParcela)
		if err != nil {
			return err
		}

		_, err = db.Exec(`INSERT INTO TCB_CONTR_FILA_EVENTOS
			(ID_EVENTO, TIPO_ACAO, STATUS, DT_BASE, NRO_SEQ_FAT)
			VALUES (?, 'LOG_FAT', 'A', ?, ?)`, fsGetIdAlt, evento.DtBase, vNroSequencial)
		if err != nil {
			return err
		}
	} else {
		_, err := db.Exec(`UPDATE TCB_CONTR_FILA_EVENTOS
			SET STATUS = 'E', MENS_ERRO = ?
			WHERE STATUS = 'P'
			AND TIPO_ACAO = ?
			AND ID_EVENTO <= ?
			AND COD_PERIODO = ?
			AND COD_PESSOA = ?
			AND COD_FIP_GF = ?
			AND COD_GRUPO_FIN = ?
			AND COD_SERVICO = ?
			AND COD_PARCELA = ?`, vMensErro, evento.TipoAcao, evento.IDEvento, evento.CodPeriodo, evento.CodPessoa, evento.CodFipGf, evento.CodGrupoFin, evento.CodServico, evento.CodParcela)
		if err != nil {
			return err
		}
	}

	return nil
}
