package service

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/machadoborges1/consome_fila_eventos/internal/entity"
	go_ora "github.com/sijms/go-ora/v2"
)

func ProcessaBDTFatGR(db *sql.DB, evento entity.TCBContrFilaEventos, vDt_Atual string) {

	var vMensErro string
	var vNroSequencial int64

	fmt.Println("Processando evento tipo BDT_FAT_GR")

	tx, err := db.Begin()
	if err != nil {
		fmt.Println("Erro ao iniciar a transação:", err)
		return
	}

	var fsGetIdAlt int64
	roww := db.QueryRow("SELECT FS_GET_ID_ALT FROM DUAL")
	if err := roww.Scan(&fsGetIdAlt); err != nil {
		log.Fatal(err)
	}

	stmt, err := db.Prepare("BEGIN " +
		"PCB_GERA_PREVISTO_MOV(pDt_Base => :1, pCod_Periodo => :2, pCod_Pessoa => :3, " +
		"pCod_Fip_GF => :4, pCod_Grupo_Fin => :5, pCod_Servico => :6, pCod_Parcela => :7, " +
		"pNecessita_Existir => :8, pNro_Sequencial => :9, pMens_Erro => :10); " +
		"END;")
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()

	_, err = stmt.Exec(
		vDt_Atual,
		evento.CodPeriodo.String,
		evento.CodPessoa.Int64,
		evento.CodFipGf.Int64,
		evento.CodGrupoFin.String,
		evento.CodServico.Int64,
		evento.CodParcela.String,
		"N",
		sql.Out{Dest: &vNroSequencial},
		go_ora.Out{Dest: &vMensErro, Size: 200})
	if err != nil {
		log.Fatal(err)
	} else {
		fmt.Println(vDt_Atual)
	}

	if vMensErro == "" {
		if vNroSequencial == 0 {
			fmt.Println("Devolveu Nro_Seq_Fat nulo")
			tx.Rollback()
		} else {
			_, err := db.Exec("BEGIN FIT.PF_SET_CONTABILIZADO_FICHA_FIN("+
				":pCod_Periodo, :pCod_Pessoa, :pCod_Fip, :pCod_Grupo_Fin, :pCod_Servico, :pCod_Parcela); END;",
				evento.CodPeriodo, evento.CodPessoa, evento.CodFipGf, evento.CodGrupoFin, evento.CodServico, evento.CodParcela)
			if err != nil {
				log.Println(err)
			} else {
				fmt.Println("commitado")
				tx.Commit()
			}
		}
	} else {
		fmt.Println("Existe vMensErro")
		tx.Rollback()
	}


	if vMensErro == "" {
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
		} else {
			tx.Commit()
			fmt.Println("Deletado")
		}

		_, err = db.Exec(`INSERT INTO TCB_CONTR_FILA_EVENTOS
		(ID_EVENTO, TIPO_ACAO, STATUS, DT_BASE, NRO_SEQ_FAT)
		VALUES (:1, 'LOG_FAT', 'A', :2, :3)`, fsGetIdAlt, evento.DtBase, vNroSequencial)
		if err != nil {
			fmt.Println("erro8")
			log.Fatal(err)
		} else {
			tx.Commit()
			fmt.Println("inserido na tabela")
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
