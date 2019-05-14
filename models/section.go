package models

import "attendance/database"

type Section struct {
	SectionId   int64    `json:"sectionId"`
	SectionDate string `json:"sectionDate"`
	Site        string `json:"site"`
	State       int64    `json:"state"`
	Course	`json:"course"`
}

func (sec *Section) SearchCourseIddById() error {
	stmtIn, err := database.SqlDB.Prepare(`SELECT course_id FROM section WHERE section_id=?`)
	if err != nil {
		return err
	}
	defer stmtIn.Close()

	err = stmtIn.QueryRow(sec.SectionId).Scan(&sec.Course.CourseId)
	if err != nil {
		return err
	}
	return nil
}

