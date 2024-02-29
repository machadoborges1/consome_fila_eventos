package entity

import (
	"database/sql"
	"errors"
)

var (
	ErrInvalidDtAlteracao        = errors.New("invalid dt_alteracao")
	ErrInvalidValorMov           = errors.New("invalid valor_mov")
	ErrInvalidCodFip             = errors.New("invalid cod_fip")
	ErrInvalidCentroCusto        = errors.New("invalid centro_custo")
	ErrInvalidStatusNfse         = errors.New("invalid status_nfse")
	ErrInvalidTipoMovimento      = errors.New("invalid tipo_movimento")
	ErrInvalidRegIgualAoAnterior = errors.New("invalid reg_igual_ao_anterior")
)

type TCBDetRealizadoDados struct {
	NroSequencial      sql.NullInt64
	DtAlteracao        sql.NullString
	NroSeqDado         sql.NullInt64
	ValorMov           sql.NullFloat64
	NroSeqPrev         sql.NullInt64
	CodFip             sql.NullInt64
	AnoMes             sql.NullString
	CodPessoa          sql.NullInt64
	CodGrupoFin        sql.NullString
	CodPeriodo         sql.NullString
	CodServico         sql.NullInt64
	CodParcela         sql.NullString
	CodTipoAlvo        sql.NullString
	CentroCusto        sql.NullInt64
	CodFipGf           sql.NullInt64
	StatusNfse         sql.NullString
	CodCompensacao     sql.NullInt64
	TipoMovimento      sql.NullString
	NroSeqReaEquivTef  sql.NullInt64
	RegIgualAoAnterior sql.NullString
}

func NewTCBDetRealizadoDados(
	nroSequencial sql.NullInt64,
	dtAlteracao sql.NullString,
	nroSeqDado sql.NullInt64,
	valorMov sql.NullFloat64,
	nroSeqPrev sql.NullInt64,
	codFip sql.NullInt64,
	anoMes sql.NullString,
	codPessoa sql.NullInt64,
	codGrupoFin sql.NullString,
	codPeriodo sql.NullString,
	codServico sql.NullInt64,
	codParcela sql.NullString,
	codTipoAlvo sql.NullString,
	centroCusto sql.NullInt64,
	codFipGf sql.NullInt64,
	statusNfse sql.NullString,
	codCompensacao sql.NullInt64,
	tipoMovimento sql.NullString,
	nroSeqReaEquivTef sql.NullInt64,
	regIgualAoAnterior sql.NullString,
) *TCBDetRealizadoDados {
	return &TCBDetRealizadoDados{
		NroSequencial:      nroSequencial,
		DtAlteracao:        dtAlteracao,
		NroSeqDado:         nroSeqDado,
		ValorMov:           valorMov,
		NroSeqPrev:         nroSeqPrev,
		CodFip:             codFip,
		AnoMes:             anoMes,
		CodPessoa:          codPessoa,
		CodGrupoFin:        codGrupoFin,
		CodPeriodo:         codPeriodo,
		CodServico:         codServico,
		CodParcela:         codParcela,
		CodTipoAlvo:        codTipoAlvo,
		CentroCusto:        centroCusto,
		CodFipGf:           codFipGf,
		StatusNfse:         statusNfse,
		CodCompensacao:     codCompensacao,
		TipoMovimento:      tipoMovimento,
		NroSeqReaEquivTef:  nroSeqReaEquivTef,
		RegIgualAoAnterior: regIgualAoAnterior,
	}
}

func (t *TCBDetRealizadoDados) Validate() error {
	// Outras validações podem ser adicionadas conforme necessário
	return nil
}
