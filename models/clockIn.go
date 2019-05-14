package models

import (
	"attendance/database"
	"attendance/util"
)

type ClockIn struct {
	ClockInId int64 `json:"clockInId"`
	StartDate string `json:"startDate"`
	State     int64 `json:"state"`
	Section	`json:"section"`
	Student	`json:"student"`
}

type ClockInByStu struct {
	StartDate string `json:"startDate"`
	State     int64 `json:"state"`
}

type ClockInByStuList struct {
	ClockInByStus []ClockInByStu `json:"clockInByStus"`
}

type ClockInByStuRe struct {
	ClockInByStu `json:"clockInByStuRe"`
}

func (clo *ClockIn) UpdateClockStateBySecIdAndStuId() error {
	stmtIn, err := database.SqlDB.Prepare(`UPDATE clock_in SET state=? WHERE section_id=? and student_id=?`)
	if err != nil {
		return err
	}
	defer stmtIn.Close()

	_, err = stmtIn.Exec(clo.State, clo.Section.SectionId, clo.Student.StudentId)
	if err != nil {
		return err
	}
	return nil
}

func (clo *ClockIn) CountStuNumBySecId() error {
	stmtIn, err := database.SqlDB.Prepare(`select count(*) from clock_in WHERE section_id=?`)
	if err != nil {
		return err
	}
	defer stmtIn.Close()

	err = stmtIn.QueryRow(clo.Section.SectionId).Scan(&clo.Student.Classs.StudentCount)
	if err != nil {
		return err
	}
	return nil
}

func (clo *ClockIn) CountStuNumBySecIdAndState() (int, error) {
	var attendanceCount int
	stmtIn, err := database.SqlDB.Prepare(`select count(*) from clock_in WHERE section_id=? and state=?`)
	if err != nil {
		return 0, err
	}
	defer stmtIn.Close()

	err = stmtIn.QueryRow(clo.Section.SectionId, clo.State).Scan(&attendanceCount)
	//err = stmtIn.QueryRow(clo.Section.SectionId, clo.State)
	if err != nil {
		return 0, err
	}
	return attendanceCount, nil
}

func (clo *ClockIn) SearchSecDateAndClockStateByStuIdAndCourseId(offset, limit int) ([]ClockInByStu, error) {
	var clock ClockInByStu
	clocks := make([]ClockInByStu, 0)

	stmtIn, err := database.SqlDB.Prepare(`select sec.section_date, c.state from clock_in as c join section as sec on c.section_id = sec.section_id where course_id=? and student_id=? limit ?,?`)
	if err != nil {
		return nil, err
	}
	defer stmtIn.Close()

	rows, err := stmtIn.Query(clo.Section.Course.CourseId, clo.Student.StudentId, offset, limit)

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		//将每一行的结果都赋值到一个academy对象中
		err := rows.Scan(&clock.StartDate, &clock.State)
		if err != nil {
			return nil, err
		}
		clock.StartDate = util.ParseTime(clock.StartDate)
		//将academy追加到academies的这个数组中
		clocks = append(clocks, clock)
	}

	return clocks, nil
}

func (clo *ClockIn) CountSecDateAndClockStateByStuIdAndCourseId() (int, error) {
	var totalElement int

	stmtIn, err := database.SqlDB.Prepare(`select count(*) from clock_in as c join section as sec on c.section_id = sec.section_id where course_id=? and student_id=?`)
	if err != nil {
		return 0, err
	}
	defer stmtIn.Close()

	err = stmtIn.QueryRow(clo.Section.Course.CourseId, clo.Student.StudentId).Scan(&totalElement)
	if err != nil {
		return 0, err
	}

	return totalElement, nil
}

func (clo *ClockIn) SearchSecDateAndClockStateByStuIdAndCourseIdSortBySecDateDesc(offset, limit int) ([]ClockInByStu, error) {
	var clock ClockInByStu
	clocks := make([]ClockInByStu, 0)

	stmtIn, err := database.SqlDB.Prepare(`select section.section_date, clock_in.state from clock_in join section on clock_in.section_id = section.section_id where course_id=? and student_id=? order by section.section_date desc limit ?,?`)
	if err != nil {
		return nil, err
	}
	defer stmtIn.Close()

	rows, err := stmtIn.Query(clo.Section.Course.CourseId, clo.Student.StudentId, offset, limit)

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		//将每一行的结果都赋值到一个academy对象中
		err := rows.Scan(&clock.StartDate, &clock.State)
		if err != nil {
			return nil, err
		}
		clock.StartDate = util.ParseTime(clock.StartDate)
		//将academy追加到academies的这个数组中
		clocks = append(clocks, clock)
	}
	return clocks, nil
}

func (clo *ClockIn) SearchSecDateAndClockStateByStuIdAndCourseIdSortBySecDateAsc(offset, limit int) ([]ClockInByStu, error) {
	var clock ClockInByStu
	clocks := make([]ClockInByStu, 0)

	stmtIn, err := database.SqlDB.Prepare(`select section.section_date, clock_in.state from clock_in join section on clock_in.section_id = section.section_id where course_id=? and student_id=? order by section.section_date asc limit ?,?`)
	if err != nil {
		return nil, err
	}
	defer stmtIn.Close()

	rows, err := stmtIn.Query(clo.Section.Course.CourseId, clo.Student.StudentId, offset, limit)

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		//将每一行的结果都赋值到一个academy对象中
		err := rows.Scan(&clock.StartDate, &clock.State)
		if err != nil {
			return nil, err
		}
		clock.StartDate = util.ParseTime(clock.StartDate)
		//将academy追加到academies的这个数组中
		clocks = append(clocks, clock)
	}
	return clocks, nil
}

func (clo *ClockIn) SearchSecDateAndClockStateByStuIdAndCourseIdSortByClockStateDesc(offset, limit int) ([]ClockInByStu, error) {
	var clock ClockInByStu
	clocks := make([]ClockInByStu, 0)

	stmtIn, err := database.SqlDB.Prepare(`select section.section_date, clock_in.state from clock_in join section on clock_in.section_id = section.section_id where course_id=? and student_id=? order by clock_in.state desc limit ?,?`)
	if err != nil {
		return nil, err
	}
	defer stmtIn.Close()

	rows, err := stmtIn.Query(clo.Section.Course.CourseId, clo.Student.StudentId, offset, limit)

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		//将每一行的结果都赋值到一个academy对象中
		err := rows.Scan(&clock.StartDate, &clock.State)
		if err != nil {
			return nil, err
		}
		clock.StartDate = util.ParseTime(clock.StartDate)
		//将academy追加到academies的这个数组中

		clocks = append(clocks, clock)
	}

	return clocks, nil
}

func (clo *ClockIn) SearchSecDateAndClockStateByStuIdAndCourseIdSortByClockStateAsc(offset, limit int) ([]ClockInByStu, error) {
	var clock ClockInByStu
	clocks := make([]ClockInByStu, 0)

	stmtIn, err := database.SqlDB.Prepare(`select section.section_date, clock_in.state from clock_in join section on clock_in.section_id = section.section_id where course_id=? and student_id=? order by clock_in.state asc limit ?,?`)
	if err != nil {
		return nil, err
	}
	defer stmtIn.Close()

	rows, err := stmtIn.Query(clo.Section.Course.CourseId, clo.Student.StudentId, offset, limit)

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		//将每一行的结果都赋值到一个academy对象中
		err := rows.Scan(&clock.StartDate, &clock.State)
		if err != nil {
			return nil, err
		}
		clock.StartDate = util.ParseTime(clock.StartDate)
		//将academy追加到academies的这个数组中
		clocks = append(clocks, clock)
	}
	return clocks, nil
}
