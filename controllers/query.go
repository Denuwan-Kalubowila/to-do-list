package query

import (
	"context"
	"fmt"
	"time"

	mysql "github.com/go-to-do/utils"
)

type Student struct {
	id    int
	fname string
	lname string
	age   int
}

func GetAllStudent() (err error) {
	ctx, cancelFunc := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancelFunc()

	row, err := mysql.GetDbConnection().QueryContext(ctx, `SELECT * FROM student`)

	if err != nil {
		return
	}

	defer row.Close()

	var students []Student

	for row.Next() {
		var std Student
		err = row.Scan(
			&std.id,
			&std.fname,
			&std.lname,
			&std.age,
		)
		if err != nil {
			return
		}
		students = append(students, std)
	}
	if err = row.Err(); err != nil {
		return
	}
	fmt.Printf("%v \n", students)
	return

}

func GetStudentbyId(id int) (err error) {
	ctx, cancelFunc := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancelFunc()

	row := mysql.GetDbConnection().QueryRowContext(ctx, `SELECT * FROM student WHERE id=?`, id)
	if err != nil {
		return
	}

	//defer row.Close()

	var stdData Student
	err = row.Scan(
		&stdData.id,
		&stdData.fname,
		&stdData.lname,
		&stdData.age,
	)
	if err != nil {
		return
	}

	fmt.Println(stdData.fname)
	fmt.Println(stdData.lname)
	fmt.Println(stdData.age)
	return
}

func UpdateById(id int, fName string, lName string, age int) (err error) {
	ctx, cancelFunc := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancelFunc()

	result, err := mysql.GetDbConnection().ExecContext(ctx, `UPDATE student SET fname=? ,lname=?,age=? WHERE id=?`, fName, lName, age, id)
	if err != nil {
		return
	}
	rows, err := result.RowsAffected()
	if err != nil {
		return
	}
	if rows != 1 {
		fmt.Printf("expected to affect 1 row, affected %d", rows)
	}
	return
}
