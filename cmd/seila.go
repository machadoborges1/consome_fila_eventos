package main

import (
	"database/sql"
	"fmt"
)

// Definindo uma estrutura para representar os dados da tabela TCB_CONTR_FILA_EVENTOS
type TCB_CONTR_FILA_EVENTOS struct {
	ID_EVENTO               int
	TIPO_ACAO               string
	STATUS                  string
	AUDSID                  sql.NullInt64
	DTH_INICIO_PROCESSAMENTO sql.NullTime
	NRO_ARQUIVO             sql.NullInt64
	NRO_LINHA               sql.NullInt64
	COD_PERIODO             sql.NullString
	COD_PESSOA              sql.NullInt64
	COD_FIP_GF              sql.NullInt64
	COD_GRUPO_FIN           sql.NullString
	COD_SERVICO             sql.NullInt64
	COD_PARCELA             sql.NullString
	MATRICULA               sql.NullString
	ANO                     sql.NullString
	COD_TURMA               sql.NullString
	COD_DISC                sql.NullString
	COTA                    sql.NullString
	MENS_ERRO               sql.NullString
	NRO_SEQ_FAT             sql.NullInt64
	NRO_SEQ_REA             sql.NullInt64
	COD_FIP_CAIXA           sql.NullInt64
	COD_CAIXA               sql.NullInt64
	COD_AUTENTICACAO        sql.NullInt64
	TIPO_BOLSA              sql.NullInt64
	NRO_CPF                 sql.NullString
	NRO_DEPOSITO            sql.NullInt64
	DT_BASE                 sql.NullString
	NRO_ITERACOES           int
}

// Definindo uma estrutura para representar os dados da tabela TCB_DET_REALIZADO_DADOS
type TCB_DET_REALIZADO_DADOS struct {
	NRO_SEQUENCIAL          int
	DT_ALTERACAO            string
	NRO_SEQ_DADO            int
	VALOR_MOV               float64
	NRO_SEQ_PREV            sql.NullInt64
	COD_FIP                 int
	ANO_MES                 string
	COD_PESSOA              int
	COD_GRUPO_FIN           string
	COD_PERIODO             string
	COD_SERVICO             int
	COD_PARCELA             string
	COD_TIPO_ALVO           string
	CENTRO_CUSTO            int
	COD_FIP_GF              int
	STATUS_NFSE             string
	COD_COMPENSACAO         int
	TIPO_MOVIMENTO          string
	NRO_SEQ_REA_EQUIV_TEF   int
	REG_IGUAL_AO_ANTERIOR   string
}

func PCB_CONSOME_FILA_EVENTOS(db *sql.DB) (pProcessou string, err error) {
	var vAchou bool
	var vAudSID int
	var vId_Atual int
	var vDt_Atual string
	var vTipo_Acao string

	// Consulta para obter os dados da fila de eventos
	rows, err := db.Query("SELECT * FROM TCB_CONTR_FILA_EVENTOS WHERE STATUS = 'A' AND TIPO_ACAO <> 'BDT_DEP_SEQ_DADO' AND NRO_ITERACOES < 10 ORDER BY ID_EVENTO")
	if err != nil {
		return "", err
	}
	defer rows.Close()

	for rows.Next() {
		var v_Dados TCB_CONTR_FILA_EVENTOS
		err := rows.Scan(
			&v_Dados.ID_EVENTO,
			&v_Dados.TIPO_ACAO,
			&v_Dados.STATUS,
			&v_Dados.AUDSID,
			&v_Dados.DTH_INICIO_PROCESSAMENTO,
			&v_Dados.NRO_ARQUIVO,
			&v_Dados.NRO_LINHA,
			&v_Dados.COD_PERIODO,
			&v_Dados.COD_PESSOA,
			&v_Dados.COD_FIP_GF,
			&v_Dados.COD_GRUPO_FIN,
			&v_Dados.COD_SERVICO,
			&v_Dados.COD_PARCELA,
			&v_Dados.MATRICULA,
			&v_Dados.ANO,
			&v_Dados.COD_TURMA,
			&v_Dados.COD_DISC,
			&v_Dados.COTA,
			&v_Dados.MENS_ERRO,
			&v_Dados.NRO_SEQ_FAT,
			&v_Dados.NRO_SEQ_REA,
			&v_Dados.COD_FIP_CAIXA,
			&v_Dados.COD_CAIXA,
			&v_Dados.COD_AUTENTICACAO,
			&v_Dados.TIPO_BOLSA,
			&v_Dados.NRO_CPF,
			&v_Dados.NRO_DEPOSITO,
			&v_Dados.DT_BASE,
			&v_Dados.NRO_ITERACOES,
		)
		if err != nil {
			return "", err
		}

		// Consulta para obter os dados distintos de NRO_SEQUENCIAL da tabela TCB_DET_REALIZADO_DADOS
		rows_Rea_Fat, err := db.Query("SELECT DISTINCT NRO_SEQUENCIAL, DT_ALTERACAO, NRO_SEQ_DADO, VALOR_MOV, NRO_SEQ_PREV, COD_FIP, ANO_MES, COD_PESSOA, COD_GRUPO_FIN, COD_PERIODO, COD_SERVICO, COD_PARCELA, COD_TIPO_ALVO, CENTRO_CUSTO, COD_FIP_GF, STATUS_NFSE, COD_COMPENSACAO, TIPO_MOVIMENTO, NRO_SEQ_REA_EQUIV_TEF, REG_IGUAL_AO_ANTERIOR FROM TCB_DET_REALIZADO_DADOS WHERE NRO_SEQ_PREV IS NOT NULL AND NRO_SEQ_PREV = ?", v_Dados.NRO_SEQ_FAT)
		if err != nil {
			return "", err
		}
		defer rows_Rea_Fat.Close()

		// Processamento dos resultados do cursor c_Rea_Fat
		// Processamento dos resultados do cursor c_Rea_Fat
	for rows_Rea_Fat.Next() {
		var v_Rea_Fat TCB_DET_REALIZADO_DADOS
		err := rows_Rea_Fat.Scan(&v_Rea_Fat.NRO_SEQUENCIAL, &v_Rea_Fat.DT_ALTERACAO, &v_Rea_Fat.NRO_SEQ_DADO, &v_Rea_Fat.VALOR_MOV, &v_Rea_Fat.NRO_SEQ_PREV, &v_Rea_Fat.COD_FIP, &v_Rea_Fat.ANO_MES, &v_Rea_Fat.COD_PESSOA, &v_Rea_Fat.COD_GRUPO_FIN, &v_Rea_Fat.COD_PERIODO, &v_Rea_Fat.COD_SERVICO, &v_Rea_Fat.COD_PARCELA, &v_Rea_Fat.COD_TIPO_ALVO, &v_Rea_Fat.CENTRO_CUSTO, &v_Rea_Fat.COD_FIP_GF, &v_Rea_Fat.STATUS_NFSE, &v_Rea_Fat.COD_COMPENSACAO, &v_Rea_Fat.TIPO_MOVIMENTO, &v_Rea_Fat.NRO_SEQ_REA_EQUIV_TEF, &v_Rea_Fat.REG_IGUAL_AO_ANTERIOR)
		if err != nil {
			return "", err
		}

		// Faça o processamento dos dados do cursor c_Rea_Fat conforme necessário
		fmt.Println("Resultado do cursor c_Rea_Fat:", v_Rea_Fat)
	}

	// Defina pProcessou como 'S' ou 'N' dependendo se a fila foi processada ou não
	if vAchou {
		pProcessou = "S"
	} else {
		pProcessou = "N"
	}

	return pProcessou, nil