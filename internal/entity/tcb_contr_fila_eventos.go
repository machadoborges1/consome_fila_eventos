package entity

import (
	"database/sql"
	"errors"
)

var (
	ErrTipoAcaoIsRequired     = errors.New("tipo_acao is required")
	ErrStatusIsRequired       = errors.New("status is required")
	ErrInvalidNroArquivo      = errors.New("invalid nro_arquivo")
	ErrInvalidNroLinha        = errors.New("invalid nro_linha")
	ErrInvalidCodPessoa       = errors.New("invalid cod_pessoa")
	ErrInvalidCodFipGf        = errors.New("invalid cod_fip_gf")
	ErrInvalidCodServico      = errors.New("invalid cod_servico")
	ErrInvalidNroSeqFat       = errors.New("invalid nro_seq_fat")
	ErrInvalidNroSeqRea       = errors.New("invalid nro_seq_rea")
	ErrInvalidCodFipCaixa     = errors.New("invalid cod_fip_caixa")
	ErrInvalidCodCaixa        = errors.New("invalid cod_caixa")
	ErrInvalidCodAutenticacao = errors.New("invalid cod_autenticacao")
	ErrInvalidTipoBolsa       = errors.New("invalid tipo_bolsa")
	ErrInvalidNroDeposito     = errors.New("invalid nro_deposito")
	ErrInvalidDtBase          = errors.New("invalid dt_base")
	ErrInvalidNroIteracoes    = errors.New("invalid nro_iteracoes")
)

type TCBContrFilaEventos struct {
	IDEvento              sql.NullInt64
	TipoAcao              sql.NullString
	Status                sql.NullString
	AUDSID                sql.NullInt64
	DtInicioProcessamento sql.NullTime
	NroArquivo            sql.NullInt64
	NroLinha              sql.NullInt64
	CodPeriodo            sql.NullString
	CodPessoa             sql.NullInt64
	CodFipGf              sql.NullInt64
	CodGrupoFin           sql.NullString
	CodServico            sql.NullInt64
	CodParcela            sql.NullString
	Matricula             sql.NullString
	Ano                   sql.NullString
	CodTurma              sql.NullString
	CodDisc               sql.NullString
	Cota                  sql.NullString
	MensErro              sql.NullString
	NroSeqFat             sql.NullInt64
	NroSeqRea             sql.NullInt64
	CodFipCaixa           sql.NullInt64
	CodCaixa              sql.NullInt64
	CodAutenticacao       sql.NullInt64
	TipoBolsa             sql.NullInt64
	NroCPF                sql.NullString
	NroDeposito           sql.NullInt64
	DtBase                sql.NullString
	NroIteracoes          sql.NullInt64
}

func NewTCBContrFilaEventos(
	idEvento sql.NullInt64,
	tipoAcao sql.NullString,
	status sql.NullString,
	audsid sql.NullInt64,
	dtInicioProcessamento sql.NullTime,
	nroArquivo sql.NullInt64,
	nroLinha sql.NullInt64,
	codPeriodo sql.NullString,
	codPessoa sql.NullInt64,
	codFipGf sql.NullInt64,
	codGrupoFin sql.NullString,
	codServico sql.NullInt64,
	codParcela sql.NullString,
	matricula sql.NullString,
	ano sql.NullString,
	codTurma sql.NullString,
	codDisc sql.NullString,
	cota sql.NullString,
	mensErro sql.NullString,
	nroSeqFat sql.NullInt64,
	nroSeqRea sql.NullInt64,
	codFipCaixa sql.NullInt64,
	codCaixa sql.NullInt64,
	codAutenticacao sql.NullInt64,
	tipoBolsa sql.NullInt64,
	nroCPF sql.NullString,
	nroDeposito sql.NullInt64,
	dtBase sql.NullString,
	nroIteracoes sql.NullInt64,
) *TCBContrFilaEventos {
	return &TCBContrFilaEventos{
		IDEvento:              idEvento,
		TipoAcao:              tipoAcao,
		Status:                status,
		AUDSID:                audsid,
		DtInicioProcessamento: dtInicioProcessamento,
		NroArquivo:            nroArquivo,
		NroLinha:              nroLinha,
		CodPeriodo:            codPeriodo,
		CodPessoa:             codPessoa,
		CodFipGf:              codFipGf,
		CodGrupoFin:           codGrupoFin,
		CodServico:            codServico,
		CodParcela:            codParcela,
		Matricula:             matricula,
		Ano:                   ano,
		CodTurma:              codTurma,
		CodDisc:               codDisc,
		Cota:                  cota,
		MensErro:              mensErro,
		NroSeqFat:             nroSeqFat,
		NroSeqRea:             nroSeqRea,
		CodFipCaixa:           codFipCaixa,
		CodCaixa:              codCaixa,
		CodAutenticacao:       codAutenticacao,
		TipoBolsa:             tipoBolsa,
		NroCPF:                nroCPF,
		NroDeposito:           nroDeposito,
		DtBase:                dtBase,
		NroIteracoes:          nroIteracoes,
	}
}

func (t *TCBContrFilaEventos) Validate() error {
	// Outras validações podem ser adicionadas conforme necessário
	return nil
}
