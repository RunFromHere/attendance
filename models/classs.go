package models

import "attendance/database"

type Classs struct {
	ClasssId     int    `json:"classsId"`
	Name         string `json:"name"`
	StudentCount int    `json:"studentCount"`
}

func (cla *Classs) SearchStuCountById() error {
	stmtIn, err := database.SqlDB.Prepare(`SELECT student_count FROM classs WHERE classs_id=?`)
	if err != nil {
		return err
	}
	defer stmtIn.Close()

	err = stmtIn.QueryRow(cla.ClasssId).Scan(&cla.StudentCount)
	if err != nil {
		return err
	}
	return nil
}
