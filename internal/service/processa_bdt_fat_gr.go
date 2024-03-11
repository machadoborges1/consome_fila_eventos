package service

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/machadoborges1/consome_fila_eventos/internal/entity"
)

func ProcessaBDTFatGR(db *sql.DB, evento entity.TCBContrFilaEventos) {

	var vMensErro string
	var vNroSequencial int64

	fmt.Println("Processando evento tipo BDT_FAT_GR")

	// Defindo a transação para que possa ser revertida em caso de erro
	tx, err := db.Begin()
	if err != nil {
		fmt.Println("Erro ao iniciar a transação:", err)
		// Trate o erro de acordo com o seu caso, por exemplo, registrando-o, lançando um panic ou retornando
		return
	}

	var vDt_Atual string
	row := db.QueryRow("SELECT TO_CHAR(SYSDATE, 'YYYYMMDD') FROM DUAL")
	if err := row.Scan(&vDt_Atual); err != nil {
		log.Fatal(err)
	}

	var fsGetIdAlt int64
	roww := db.QueryRow("SELECT FS_GET_ID_ALT FROM DUAL")
	if err := roww.Scan(&fsGetIdAlt); err != nil {
		log.Fatal(err)
	}

	_, err = db.Exec("BEGIN "+
		"PCB_GERA_PREVISTO_MOV(pDt_Base => :1, pCod_Periodo => :2, pCod_Pessoa => :3, "+
		"pCod_Fip_GF => :4, pCod_Grupo_Fin => :5, pCod_Servico => :6, pCod_Parcela => :7, "+
		"pNecessita_Existir => :8, pNro_Sequencial => :9, pMens_Erro => :10); "+
		"END;",
		vDt_Atual,
		evento.CodPeriodo.String,
		evento.CodPessoa.Int64,
		evento.CodFipGf.Int64,
		evento.CodGrupoFin.String,
		evento.CodServico.Int64,
		evento.CodParcela.String,
		"N",
		vNroSequencial,
		vMensErro)
	if err != nil {
		fmt.Println("erro1")
		log.Fatal(err)
	} else {
		fmt.Printf("okkkkkk")
	}

	if vMensErro != "" {
		if vNroSequencial == 0 {
			fmt.Println("erro2")
			tx.Rollback()
			fmt.Println("numero sequencial nulo")
		}
		_, err := db.Exec("BEGIN FIT.PF_SET_CONTABILIZADO_FICHA_FIN("+
			":pCod_Periodo, :pCod_Pessoa, :pCod_Fip, :pCod_Grupo_Fin, :pCod_Servico, :pCod_Parcela); END;",
			evento.CodPeriodo, evento.CodPessoa, evento.CodFipGf, evento.CodGrupoFin, evento.CodServico, evento.CodParcela)
		if err != nil {
			fmt.Println("erro3")
			log.Println(err)
		} else {
			fmt.Println("erro4")
			tx.Commit()
		}
	} else {
		fmt.Println("erro5")
		tx.Rollback()
	}

	if vMensErro == "" {
		fmt.Println("erro6")
		_, err := db.Exec(`DELETE FROM TCB_CONTR_FILA_EVENTOS
			WHERE STATUS IN ('P','E')
			AND TIPO_ACAO = :1
			AND ID_EVENTO <= :2
			AND COD_PERIODO = :3
			AND COD_PESSOA = :4
			AND COD_FIP_GF = :5
			AND COD_GRUPO_FIN = :6
			AND COD_SERVICO = :7
			AND COD_PARCELA = :8`,
			evento.TipoAcao, evento.IDEvento, evento.CodPeriodo, evento.CodPessoa, evento.CodFipGf, evento.CodGrupoFin, evento.CodServico, evento.CodParcela)
		if err != nil {
			fmt.Println("erro7")
			log.Fatal(err)
		}
		_, err = db.Exec(`INSERT INTO TCB_CONTR_FILA_EVENTOS
		(ID_EVENTO, TIPO_ACAO, STATUS, DT_BASE, NRO_SEQ_FAT)
		VALUES (:1, 'LOG_FAT', 'A', :2, :3)`, fsGetIdAlt, evento.DtBase, vNroSequencial)
		if err != nil {
			fmt.Println("erro8")
			log.Fatal(err)
		}

	} else {
		fmt.Println("erro9")
		_, err := db.Exec(`UPDATE TCB_CONTR_FILA_EVENTOS
			SET STATUS = 'E', MENS_ERRO = :1
			WHERE STATUS = 'P'
			AND TIPO_ACAO = :2
			AND ID_EVENTO <= :3
			AND COD_PERIODO = :4
			AND COD_PESSOA = :5
			AND COD_FIP_GF = :6
			AND COD_GRUPO_FIN = :7
			AND COD_SERVICO = :8
			AND COD_PARCELA = :9`,
			vMensErro, evento.TipoAcao, evento.IDEvento, evento.CodPeriodo, evento.CodPessoa, evento.CodFipGf, evento.CodGrupoFin, evento.CodServico, evento.CodParcela)
		if err != nil {
			fmt.Println("erro10")
			log.Fatal(err)
		}
	}
}

// func processaBDTFatGR(db *sql.DB, evento *entity.TCBContrFilaEventos) error {
// 	var vMensErro string
// 	var vNroSequencial int64

// 	fmt.Println("Processando evento tipo BDT_FAT_GR")

// 	// Defina a transação para que possa ser revertida em caso de erro
// 	tx, err := db.Begin()
// 	if err != nil {
// 		return err
// 	}

// 	// Chame a função PCB_GERA_PREVISTO_MOV
// 	err = pcbGeraPrevistoMov(db, evento)
// 	if err != nil {
// 		tx.Rollback() // Reverta a transação em caso de erro
// 		return err
// 	}

// 	if vMensErro == "" {
// 		if vNroSequencial == 0 {
// 			tx.Rollback()
// 			return fmt.Errorf("Devolveu Nro_Seq_Fat nulo")
// 		}

// 		// Chame a função FIT.PF_SET_CONTABILIZADO_FICHA_FIN
// 		err = fitPfSetContabilizadoFichaFin(db, evento)
// 		if err != nil {
// 			tx.Rollback()
// 			return err
// 		}

// 		tx.Commit() // Commit se não houver erros
// 	} else {
// 		tx.Rollback() // Reverta a transação em caso de erro
// 	}

// 	// Se não houver erros, execute a exclusão e inserção conforme necessário
// 	if vMensErro == "" {
// 		_, err := db.Exec(`DELETE FROM TCB_CONTR_FILA_EVENTOS
// 			WHERE STATUS IN ('P','E')
// 			AND TIPO_ACAO = ?
// 			AND ID_EVENTO <= ?
// 			AND COD_PERIODO = ?
// 			AND COD_PESSOA = ?
// 			AND COD_FIP_GF = ?
// 			AND COD_GRUPO_FIN = ?
// 			AND COD_SERVICO = ?
// 			AND COD_PARCELA = ?`, evento.TipoAcao, evento.IDEvento, evento.CodPeriodo, evento.CodPessoa, evento.CodFipGf, evento.CodGrupoFin, evento.CodServico, evento.CodParcela)
// 		if err != nil {
// 			return err
// 		}

// 		_, err = db.Exec(`INSERT INTO TCB_CONTR_FILA_EVENTOS
// 			(ID_EVENTO, TIPO_ACAO, STATUS, DT_BASE, NRO_SEQ_FAT)
// 			VALUES (?, 'LOG_FAT', 'A', ?, ?)`, fsGetIdAlt, evento.DtBase, vNroSequencial)
// 		if err != nil {
// 			return err
// 		}
// 	} else {
// 		_, err := db.Exec(`UPDATE TCB_CONTR_FILA_EVENTOS
// 			SET STATUS = 'E', MENS_ERRO = ?
// 			WHERE STATUS = 'P'
// 			AND TIPO_ACAO = ?
// 			AND ID_EVENTO <= ?
// 			AND COD_PERIODO = ?
// 			AND COD_PESSOA = ?
// 			AND COD_FIP_GF = ?
// 			AND COD_GRUPO_FIN = ?
// 			AND COD_SERVICO = ?
// 			AND COD_PARCELA = ?`, vMensErro, evento.TipoAcao, evento.IDEvento, evento.CodPeriodo, evento.CodPessoa, evento.CodFipGf, evento.CodGrupoFin, evento.CodServico, evento.CodParcela)
// 		if err != nil {
// 			return err
// 		}
// 	}

// 	return nil
// }
