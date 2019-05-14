package models

import "attendance/database"

type Course struct {
	CourseId int64    `json:"courseId" binding:"required"`
	Name     string `json:"name"`
	PassWord string `json:"password"`
	Teacher  `json:"teacher"`
	Classs   `json:"classs"`
	Academy  `json:"academy"`
}

func (cou *Course) SearchdClasssIdById() error {
	stmtIn, err := database.SqlDB.Prepare(`SELECT classs_id FROM course WHERE course_id=?`)
	if err != nil {
		return err
	}
	defer stmtIn.Close()

	err = stmtIn.QueryRow(cou.CourseId).Scan(&cou.Classs.ClasssId)
	if err != nil {
		return err
	}
	return nil
}

