package service

import (
	"database/sql"
	"fmt"
	"log"
)

func PcbConsomeFilaEventos(db *sql.DB) bool {

	dado, err := SelectFirstTCBContrFilaEvento(db)
	if err != nil {
		log.Fatal("Erro ao selecionar dados:", err)
		return false
	}
	fmt.Println()

	var vDt_Atual string
	row := db.QueryRow("SELECT TO_CHAR(SYSDATE, 'YYYYMMDD') FROM DUAL")
	if err := row.Scan(&vDt_Atual); err != nil {
		log.Fatal(err)
	}

	var vDt_Atuall sql.NullTime
	rowwww := db.QueryRow("SELECT SYSDATE FROM DUAL")
	if err := rowwww.Scan(&vDt_Atuall); err != nil {
		fmt.Println("erro4")
		log.Fatal(err)
	}

	var fsGetIdAlt int64
	rowww := db.QueryRow("SELECT FS_GET_ID_ALT FROM DUAL")
	if err := rowww.Scan(&fsGetIdAlt); err != nil {
		log.Fatal(err)
	}

	var vAudSID int64
	roww := db.QueryRow("SELECT USERENV('SESSIONID') FROM DUAL")
	if err := roww.Scan(&vAudSID); err != nil {
		fmt.Println("erro3")
		log.Fatal(err)
	}

	fmt.Println(dado, vAudSID, vDt_Atual, fsGetIdAlt)
	TravaBDTFatGR(db, dado, vAudSID, fsGetIdAlt)

	return true

}
