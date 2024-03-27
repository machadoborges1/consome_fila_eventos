package main

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/machadoborges1/consome_fila_eventos/configs"
	"github.com/machadoborges1/consome_fila_eventos/internal/entity"
	"github.com/machadoborges1/consome_fila_eventos/internal/service"

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
	service.PcbConsomeFilaEventos(db)

	// var count int
	// err = db.QueryRow("SELECT COUNT(*) FROM Exemplo").Scan(&count)
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// fmt.Printf("Saída da procedure: %s\n", saida)

	// // Chame a função alterarPorID
	// id := 12                         // ID da linha que deseja alterar
	// novoNome := "NovoNooooooooooome" // Novo nome para atribuir à linha

	// rowsAffected, err := alterarPorID(db, id, novoNome)
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// fmt.Printf("Número de linhas modificadas: %d\n", rowsAffected)

	// Exiba os dados selecionados

	//service.ProcessaBDTFatGR(db, dado)
	//service.TravaBDTFatGR(db, dado)

	/////////////////////////////////////////////////////////////

	// var saida string

	// stmt, err := db.Prepare("BEGIN InserirValor(:1, :2, :3); END;")
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// defer stmt.Close()

	// _, err = stmt.Exec(55, "Exempl013333", go_ora.Out{Dest: &saida, Size: 200})
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// fmt.Println("Valor de saida após a execução da procedure:", saida)

	//////////////////////////////////////////////////////////////

	//var vMensErro string
	//var vNroSequencial int64

	// fmt.Println(vDt_Atual,
	// 	dado.CodPeriodo.String,
	// 	dado.CodPessoa.Int64,
	// 	dado.CodFipGf.Int64,
	// 	dado.CodGrupoFin.String,
	// 	dado.CodServico.Int64,
	// 	dado.CodParcela.String,
	// 	"N",
	// 	vNroSequencial,
	// 	vMensErro)

	// result, err := db.Exec("CALL PCB_GERA_PREVISTO_MOV(:1, :2, :3, :4, :5, :6, :7, :8, :9, :10)",
	// 	vDt_Atual,
	// 	dado.CodPeriodo.String,
	// 	dado.CodPessoa.Int64,
	// 	dado.CodFipGf.Int64,
	// 	dado.CodGrupoFin.String,
	// 	dado.CodServico.Int64,
	// 	dado.CodParcela.String,
	// 	"N",
	// 	vNroSequencial,
	// 	vMensErro)
	// fmt.Println(vMensErro)
	// if err != nil {
	// 	fmt.Println("erro1")
	// 	log.Fatal(err)
	// } else {
	// 	fmt.Println(result)
	// 	fmt.Printf("okkkkkk")
	// }

	// var num int
	// num = 7
	// var sauda string

	// oi, err := db.Exec("BEGIN VerificaNumero(:1, :2);END;", num, sauda)
	// if err != nil {
	// 	fmt.Print(oi)
	// 	log.Fatal(err)
	// }

	// // Imprima a saudação
	// fmt.Println("Saudação:", sauda)
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

// testada e correta ---------------------------------------------------------------------------------------------------------------------
func selectTCBContrFilaEventos(db *sql.DB) ([]entity.TCBContrFilaEventos, error) {
	rows, err := db.Query("SELECT * FROM TCB_CONTR_FILA_EVENTOS WHERE STATUS = 'A' AND TIPO_ACAO <> 'BDT_DEP_SEQ_DADO' AND NRO_ITERACOES < 10 AND ROWNUM = 1 ORDER BY ID_EVENTO")
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

// func travaBDTFatGR(db *sql.DB, vAudSID int64, vIdAtual int64, vDados entity.TCBContrFilaEventos) (bool, error) {
// 	vTipoAcao := "BDT_FAT_GR"

// 	_, err := db.Exec(`UPDATE TCB_CONTR_FILA_EVENTOS
// 	SET STATUS = 'P',
// 		AUDSID = :1,
// 		DTH_INICIO_PROCESSAMENTO = SYSDATE
// 	WHERE STATUS = 'A'
// 	  AND TIPO_ACAO = :2
// 	  AND ID_EVENTO <= :3
// 	  AND COD_PERIODO = :4
// 	  AND COD_PESSOA = :5
// 	  AND COD_FIP_GF = :6
// 	  AND COD_GRUPO_FIN = :7
// 	  AND COD_SERVICO = :8
// 	  AND COD_PARCELA = :9
// 	  AND NOT EXISTS (
// 		SELECT *
// 		FROM TCB_CONTR_FILA_EVENTOS
// 		WHERE STATUS = 'P'
// 		  AND AUDSID IS NOT NULL
// 		  AND AUDSID <> :1
// 		  AND TIPO_ACAO = :2
// 		  AND COD_PERIODO = :4
// 		  AND COD_PESSOA = :5
// 		  AND COD_FIP_GF = :6
// 		  AND COD_GRUPO_FIN = :7
// 		  AND COD_SERVICO = :8
// 		  AND COD_PARCELA = :9
// 	  )`, vAudSID, vTipoAcao, vIdAtual, vDados.CodPeriodo, vDados.CodPessoa, vDados.CodFipGf, vDados.CodGrupoFin, vDados.CodServico, vDados.CodParcela)
// 	if err != nil {
// 		return false, err
// 	}

// 	rowCount, err := db.Exec(`SELECT COUNT(*) FROM TCB_CONTR_FILA_EVENTOS WHERE STATUS = 'P' AND TIPO_ACAO = :1 AND ID_EVENTO <= :2 AND COD_PERIODO = :3 AND COD_PESSOA = :4 AND COD_FIP_GF = :5 AND COD_GRUPO_FIN = :6 AND COD_SERVICO = :7 AND COD_PARCELA = :8`, vTipoAcao, vIdAtual, vDados.CodPeriodo, vDados.CodPessoa, vDados.CodFipGf, vDados.CodGrupoFin, vDados.CodServico, vDados.CodParcela)
// 	if err != nil {
// 		return false, err
// 	}

// 	var rowCountInt int64
// 	err = rowCount.Scan(&rowCountInt)
// 	if err != nil {
// 		return false, err
// 	}

// 	var retorno bool
// 	if rowCountInt > 0 {
// 		err := db.Commit()
// 		if err != nil {
// 			return false, err
// 		}
// 		retorno = true
// 	} else {
// 		err := db.Rollback()
// 		if err != nil {
// 			return false, err
// 		}

// 		_, err = db.Exec(`UPDATE TCB_CONTR_FILA_EVENTOS
// 		SET NRO_ITERACOES = NRO_ITERACOES + 1
// 		WHERE STATUS = 'A'
// 		  AND TIPO_ACAO = :1
// 		  AND ID_EVENTO <= :2
// 		  AND COD_PERIODO = :3
// 		  AND COD_PESSOA = :4
// 		  AND COD_FIP_GF = :5
// 		  AND COD_GRUPO_FIN = :6
// 		  AND COD_SERVICO = :7
// 		  AND COD_PARCELA = :8`, vTipoAcao, vIdAtual, vDados.CodPeriodo, vDados.CodPessoa, vDados.CodFipGf, vDados.CodGrupoFin, vDados.CodServico, vDados.CodParcela)
// 		if err != nil {
// 			return false, err
// 		}

// 		err = db.Commit()
// 		if err != nil {
// 			return false, err
// 		}

// 		retorno = false
// 	}

// 	return retorno, nil
// }

func chamarProcedure(db *sql.DB, id int, nome string) (string, error) {
	var saida string

	// Preparar a chamada da procedure
	stmt, err := db.Prepare("BEGIN InserirValor(:1, :2, :3); END;")
	if err != nil {
		return "", err
	}
	defer stmt.Close()

	// Executar a procedure
	_, err = stmt.Exec(id, nome, go_ora.Out{Dest: &saida, Size: 200})
	if err != nil {
		return "", err
	}

	return saida, nil
}

func alterarPorID(db *sql.DB, id int, novoNome string) (int64, error) {
	// Preparar a declaração SQL para atualização
	stmt, err := db.Prepare("UPDATE Exemplo SET Nome = :1 WHERE ID = :2")
	if err != nil {
		return 0, err
	}
	defer stmt.Close()

	// Executar a declaração SQL
	result, err := stmt.Exec(novoNome, id)
	if err != nil {
		return 0, err
	}

	// Obter o número de linhas afetadas
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return 0, err
	}

	return rowsAffected, nil
}
