package models

import "attendance/database"

type Teacher struct {
	TeacherId  int    `json:"teacherId"`
	TeacherNum string `json:"teacherNum"`
	TrueName   string `json:"truename"`
	PassWord   string `json:"password"`
	Academy    `json:"academy"`
}

func (tea *Teacher) UpdateTeaPsdById() error {
	stmtIn, err := database.SqlDB.Prepare(`UPDATE teacher SET password=? WHERE teacher_id=?`)
	if err != nil {
		return err
	}
	defer stmtIn.Close()

	_, err = stmtIn.Exec(tea.PassWord, tea.TeacherId)
	if err != nil {
		return err
	}
	return nil
}

func (tea *Teacher) SearchTeaPsdById() error {
	stmtIn, err := database.SqlDB.Prepare(`SELECT password FROM teacher WHERE teacher_id=?`)
	if err != nil {
		return err
	}
	defer stmtIn.Close()

	err = stmtIn.QueryRow(tea.TeacherId).Scan(&tea.PassWord)
	if err != nil {
		return err
	}
	return nil
}

func (tea *Teacher) SearchTeaIdAndTrueNameAndPsdByTeaNum() error {
	stmtIn, err := database.SqlDB.Prepare(`SELECT teacher_id, truename, password FROM teacher WHERE teacher_num=?`)
	if err != nil {
		return err
	}
	defer stmtIn.Close()

	err = stmtIn.QueryRow(tea.TeacherNum).Scan(&tea.TeacherId, &tea.TrueName, &tea.PassWord)
	if err != nil {
		return err
	}
	return nil
}
