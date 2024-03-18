package service

import (
	"database/sql"

	"github.com/machadoborges1/consome_fila_eventos/internal/entity"
)

func SelectFirstTCBContrFilaEvento(db *sql.DB) (entity.TCBContrFilaEventos, error) {
	row := db.QueryRow("SELECT * FROM TCB_CONTR_FILA_EVENTOS WHERE STATUS = 'A' AND TIPO_ACAO <> 'BDT_DEP_SEQ_DADO' AND NRO_ITERACOES < 10 ORDER BY ID_EVENTO")

	var evento entity.TCBContrFilaEventos
	err := row.Scan(
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
		if err == sql.ErrNoRows {
			return entity.TCBContrFilaEventos{}, nil // Retorna um objeto vazio se não houver linhas
		}
		return entity.TCBContrFilaEventos{}, err
	}

	return evento, nil
}
