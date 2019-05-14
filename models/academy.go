package models

import (
	"attendance/database"
	"fmt"
)

type Academy struct {
	AcademyId int64    `json:"academyId"`
	Name      string `json:"name"`
}

func (aca *Academy) AddAcademy() error {
	stmtIn, err := database.SqlDB.Prepare("INSERT INTO academy (name) values (?)")
	if err != nil {
		return err
	}
	_, err = stmtIn.Exec(aca.Name)
	if err != nil {
		return err
	}

	return nil
}

func (aca *Academy) UpdateAcademyById() error {
	stmtIn, err := database.SqlDB.Prepare(`UPDATE academy SET name=? WHERE academy_id=?`)
	if err != nil {
		return err
	}
	defer stmtIn.Close()

	_, err = stmtIn.Exec(aca.Name, aca.AcademyId)
	if err != nil {
		return err
	}
	return nil
}

func (aca *Academy) DeleteAcademyById() error {
	stmtIn, err := database.SqlDB.Prepare(`DELETE FROM academy WHERE academy_id=?`)
	if err != nil {
		return err
	}
	defer stmtIn.Close()

	_, err = stmtIn.Exec(aca.AcademyId)
	if err != nil {
		return err
	}
	return nil
}

func (aca *Academy) SearchAcademyById() error {
	stmtIn, err := database.SqlDB.Prepare(`SELECT name FROM academy WHERE academy_id=?`)
	if err != nil {
		return err
	}
	defer stmtIn.Close()

	err = stmtIn.QueryRow(aca.AcademyId).Scan(&aca.Name)
	if err != nil {
		return err
	}
	return nil
}

func (aca *Academy) SearchAllAcademy(offset, limit int) ([]Academy, error) {
	var academy Academy
	academies := make([]Academy, 0)

	//stmtIn, err := database.SqlDB.Prepare(`SELECT academy_id, name FROM academy order by name limit ?,?`)
	stmtIn, err := database.SqlDB.Prepare(`SELECT academy_id, name FROM academy order by name limit ?,?`)
	if err != nil {
		return nil, err
	}
	defer stmtIn.Close()

	rows, err := stmtIn.Query(offset, limit)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	//循环读取结果
	for rows.Next() {
		//将每一行的结果都赋值到一个academy对象中
		err := rows.Scan(&academy.AcademyId, &academy.Name)
		if err != nil {
			return nil, err
		}
		fmt.Println("单个数据： ", academy)
		//将academy追加到academies的这个数组中
		academies = append(academies, academy)
	}
	fmt.Println("duo个数据： ", academies)

	return academies, nil
}

func (aca *Academy) SearchAllAcademyLikeName(offset, limit int) ([]Academy, error) {
	var academy Academy
	academies := make([]Academy, 0)

	stmtIn, err := database.SqlDB.Prepare("SELECT academy_id, name FROM academy WHERE name like ? limit ?,?")
	if err != nil {
		return nil, err
	}
	defer stmtIn.Close()

	aca.Name = database.ParseLikeMatchForString(aca.Name)
	rows, err := stmtIn.Query(aca.Name, offset, limit)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	//循环读取结果
	for rows.Next() {
		//将每一行的结果都赋值到一个academy对象中
		err := rows.Scan(&academy.AcademyId, &academy.Name)
		if err != nil {
			return nil, err
		}
		//将academy追加到academies的这个数组中
		academies = append(academies, academy)
		//fmt.Println(academy.AcademyId, "\t", academy.Name)
	}

	return academies, nil
}
