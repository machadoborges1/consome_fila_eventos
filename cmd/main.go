package main

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/machadoborges1/consome_fila_eventos/configs"
	"github.com/machadoborges1/consome_fila_eventos/internal/entity"

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