package main

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/machadoborges1/consome_fila_eventos/configs"
	"github.com/machadoborges1/consome_fila_eventos/internal/entity"
	"github.com/machadoborges1/consome_fila_eventos"

	go_ora "github.com/sijms/go-ora/v2"
)

func main() {

	config, err := configs.LoadConfig(".")
	if err != nil {
		panic(err)
	}

	port := config.Port
	connString := go_ora.BuildUrl(config.Host, port, config.ServiceName, config.User, config.Password, nil)
	println(connString)

	db, err := sql.Open("oracle", connString)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Conexão bem sucedida!")

	dados, err := selectTCBContrFilaEventos(db)
	if err != nil {
		log.Fatal("Erro ao selecionar dados:", err)
	}

	// Exiba os dados selecionados
	fmt.Println(dados[1])

	// idEvento := int64(25852682504738455)
	// evento, err := selectTCBContrFilaEventos(db, idEvento)
	// if err != nil {
	// 	log.Fatal("Erro ao recuperar evento:", err)
	// }

	// // Exiba os detalhes do evento
	// fmt.Println("Detalhes do Evento:")
	// fmt.Println("ID do Evento:", evento.IDEvento)
	// fmt.Println("Tipo de Ação:", evento.TipoAcao)
	// fmt.Println("Status:", evento.Status)
}

// func selectTCBContrFilaEventos(db *sql.DB, id int64) (*entity.TCBContrFilaEventos, error) {
// 	stmt, err := db.Prepare("SELECT ID_EVENTO, TIPO_ACAO, STATUS, AUDSID, DTH_INICIO_PROCESSAMENTO, NRO_ARQUIVO, NRO_LINHA, COD_PERIODO, COD_PESSOA, COD_FIP_GF, COD_GRUPO_FIN, COD_SERVICO, COD_PARCELA, MATRICULA, ANO, COD_TURMA, COD_DISC, COTA, MENS_ERRO, NRO_SEQ_FAT, NRO_SEQ_REA, COD_FIP_CAIXA, COD_CAIXA, COD_AUTENTICACAO, TIPO_BOLSA, NRO_CPF, NRO_DEPOSITO, DT_BASE, NRO_ITERACOES FROM tcb_contr_fila_eventos WHERE ID_EVENTO = :1")
// 	if err != nil {
// 		return nil, err
// 	}
// 	defer stmt.Close()

// 	var tcb entity.TCBContrFilaEventos
// 	err = stmt.QueryRow(id).Scan(&tcb.IDEvento, &tcb.TipoAcao, &tcb.Status, &tcb.AUDSID, &tcb.DtInicioProcessamento, &tcb.NroArquivo, &tcb.NroLinha, &tcb.CodPeriodo, &tcb.CodPessoa, &tcb.CodFipGf, &tcb.CodGrupoFin, &tcb.CodServico, &tcb.CodParcela, &tcb.Matricula, &tcb.Ano, &tcb.CodTurma, &tcb.CodDisc, &tcb.Cota, &tcb.MensErro, &tcb.NroSeqFat, &tcb.NroSeqRea, &tcb.CodFipCaixa, &tcb.CodCaixa, &tcb.CodAutenticacao, &tcb.TipoBolsa, &tcb.NroCPF, &tcb.NroDeposito, &tcb.DtBase, &tcb.NroIteracoes)
// 	if err != nil {
// 		return nil, err
// 	}

// 	return &tcb, nil
// }




// func PCB_CONSOME_FILA_EVENTOS(db *sql.DB) (pProcessou string, err error) {
// 	var vAchou bool
// 	var vAudSID int
// 	var vId_Atual int
// 	var vDt_Atual string
// 	var vTipo_Acao string
// }




func selectTCBContrFilaEventos(db *sql.DB) ([]entity.TCBContrFilaEventos, error) {
	rows, err := db.Query("SELECT * FROM TCB_CONTR_FILA_EVENTOS WHERE STATUS = 'A' AND TIPO_ACAO <> 'BDT_DEP_SEQ_DADO' AND NRO_ITERACOES < 10 ORDER BY ID_EVENTO")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var eventos []entity.TCBContrFilaEventos

	for rows.Next() {
		var evento entity.TCBContrFilaEventos
		err := rows.Scan(
			&evento.IDEvento,
			&evento.TipoAcao,
			&evento.Status,
			&evento.AUDSID,
			&evento.DtInicioProcessamento,
			&evento.NroArquivo,
			&evento.NroLinha,
			&evento.CodPeriodo,
			&evento.CodPessoa,
			&evento.CodFipGf,
			&evento.CodGrupoFin,
			&evento.CodServico,
			&evento.CodParcela,
			&evento.Matricula,
			&evento.Ano,
			&evento.CodTurma,
			&evento.CodDisc,
			&evento.Cota,
			&evento.MensErro,
			&evento.NroSeqFat,
			&evento.NroSeqRea,
			&evento.CodFipCaixa,
			&evento.CodCaixa,
			&evento.CodAutenticacao,
			&evento.TipoBolsa,
			&evento.NroCPF,
			&evento.NroDeposito,
			&evento.DtBase,
			&evento.NroIteracoes,
		)
		if err != nil {
			return nil, err
		}
		eventos = append(eventos, evento)
	}

	err = rows.Err()
	if err != nil {
		return nil, err
	}

	return eventos, nil
}


func selectDistinctNroSeqRea(db *sql.DB, nroSeqFat int64) ([]int64, error) {
	rows, err := db.Query("SELECT DISTINCT NRO_SEQUENCIAL AS NRO_SEQ_REA FROM TCB_DET_REALIZADO_DADOS WHERE NRO_SEQ_PREV IS NOT NULL AND NRO_SEQ_PREV = :1", nroSeqFat)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var nroSeqReaList []int64

	for rows.Next() {
		var nroSeqRea int64
		err := rows.Scan(&nroSeqRea)
		if err != nil {
			return nil, err
		}
		nroSeqReaList = append(nroSeqReaList, nroSeqRea)
	}

	err = rows.Err()
	if err != nil {
		return nil, err
	}

	return nroSeqReaList, nil
}


func travaBDTFatGR(db *sql.DB, vAudSID int64, vIdAtual int64, vDados entity.TCBContrFilaEventos) (bool, error) {
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
