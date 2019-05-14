package models

import (
	"attendance/database"
)

type Admin struct {
	AdminId  int    `json:"adminId"`
	Username string `json:"username"`
	PassWord string	`json:"password"`
}

func (adm *Admin) UpdateAdmPasswordById() error {
	stmtIn, err := database.SqlDB.Prepare(`UPDATE admin SET password=? WHERE admin_id=?`)
	if err != nil {
		return err
	}
	defer stmtIn.Close()

	_, err = stmtIn.Exec(adm.PassWord, adm.AdminId)
	if err != nil {
		return err
	}
	return nil
}

func (adm *Admin) SearchAdmPasswordById() error {
	stmtIn, err := database.SqlDB.Prepare(`SELECT password FROM admin WHERE admin_id=?`)
	if err != nil {
		return err
	}
	defer stmtIn.Close()

	err = stmtIn.QueryRow(adm.AdminId).Scan(&adm.PassWord)
	if err != nil {
		return err
	}
	return nil
}

func (adm *Admin) SearchAdmIdAndPsdByUsername() error {
	stmtIn, err := database.SqlDB.Prepare(`SELECT admin_id, password FROM admin WHERE username=?`)
	if err != nil {
		return err
	}
	defer stmtIn.Close()

	err = stmtIn.QueryRow(adm.Username).Scan(&adm.AdminId, &adm.PassWord)
	if err != nil {
		return err
	}
	return nil
}
