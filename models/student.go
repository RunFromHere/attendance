package models

import "attendance/database"

type Student struct {
	StudentId  int64    `json:"studentId"`
	StudentNum string `json:"studentNum"`
	TrueName   string `json:"truename"`
	PassWord   string `json:"password"`
	Academy    `json:"academy"`
	Classs     `json:"classs"`
}

func (stu *Student) UpdateStuPsdById() error {
	stmtIn, err := database.SqlDB.Prepare(`UPDATE student SET password=? WHERE student_id=?`)
	if err != nil {
		return err
	}
	defer stmtIn.Close()

	_, err = stmtIn.Exec(stu.PassWord, stu.StudentId)
	if err != nil {
		return err
	}
	return nil
}

func (stu *Student) SearchStuPsdById() error {
	stmtIn, err := database.SqlDB.Prepare(`SELECT password FROM student WHERE student_id=?`)
	if err != nil {
		return err
	}
	defer stmtIn.Close()

	err = stmtIn.QueryRow(stu.StudentId).Scan(&stu.PassWord)
	if err != nil {
		return err
	}
	return nil
}

func (stu *Student) SearchStuNumAndNameById() error {
	stmtIn, err := database.SqlDB.Prepare(`SELECT student_num, truename FROM student WHERE student_id=?`)
	if err != nil {
		return err
	}
	defer stmtIn.Close()

	err = stmtIn.QueryRow(stu.StudentId).Scan(&stu.StudentNum, &stu.TrueName)
	if err != nil {
		return err
	}
	return nil
}

func (stu *Student) SearchStuIdAndNameByStuNum() error {
	stmtIn, err := database.SqlDB.Prepare(`SELECT student_id, truename FROM student WHERE student_num=?`)
	if err != nil {
		return err
	}
	defer stmtIn.Close()

	err = stmtIn.QueryRow(stu.StudentNum).Scan(&stu.StudentId, &stu.TrueName)
	if err != nil {
		return err
	}
	return nil
}

func (stu *Student) SearchStuIdAndNameAndPsdByStuNum() error {
	stmtIn, err := database.SqlDB.Prepare(`SELECT student_id, truename, password FROM student WHERE student_num=?`)
	if err != nil {
		return err
	}
	defer stmtIn.Close()

	err = stmtIn.QueryRow(stu.StudentNum).Scan(&stu.StudentId, &stu.TrueName, &stu.PassWord)
	if err != nil {
		return err
	}
	return nil
}



