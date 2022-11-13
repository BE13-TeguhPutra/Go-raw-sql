package controller

import (
	entities "be13/rawsql/Entities"
	"database/sql"
	"fmt"
	"log"
)

func BacaData(db *sql.DB) ([]entities.User, error) {
	result, errSelect := db.Query("SELECT id,Nama,email,prhone FROM users")
	if errSelect != nil {
		// log.Fatal("error select", errSelect.Error())
		return nil, errSelect
	}

	var dataUser []entities.User
	for result.Next() {
		var userrow entities.User
		errScan := result.Scan(&userrow.Id, &userrow.Nama, &userrow.Email, &userrow.Phone)
		if errScan != nil {
			// log.Fatal("eror scan", errScan.Error())
			return nil, errScan
		}
		// fmt.Printf("id:%s  nama:%s  email: %s\n", userrow.Id, userrow.Nama, userrow.Email)
		dataUser = append(dataUser, userrow)
	}
	return dataUser, nil

}

func Insertdata(db *sql.DB, insert entities.User) error {

	var query = "insert into users (id,nama,email,prhone,domisili) values (?,?,?,?,?)"
	statement, errPrepare := db.Prepare(query)
	if errPrepare != nil {
		// log.Fatal("error prepare insert", errPrepare.Error())
		return errPrepare
	}

	result, errExec := statement.Exec(insert.Id, insert.Nama, insert.Email, insert.Phone, insert.Domisili)
	if errExec != nil {
		// log.Fatal("error exec insert", errExec.Error())
		return errExec
	} else {
		row, _ := result.RowsAffected()
		if row > 0 {
			fmt.Println("Insert berhasil")
		} else {
			fmt.Println("Insert gagal")
		}
	}
	return nil

}

func Update(db *sql.DB, updateUser entities.User) error {
	var query = "UPDATE users set nama=?,email=?,prhone=?,domisili=? where id=?"
	statement, errPrepare := db.Prepare(query)
	if errPrepare != nil {
		// log.Fatal("error prepare", errPrepare.Error())
		return errPrepare
	}

	result, errExec := statement.Exec(updateUser.Nama, updateUser.Email, updateUser.Phone, updateUser.Domisili, updateUser.Id)
	if errExec != nil {
		// log.Fatal("error exec insert", errExec.Error())
		return errExec
	} else {
		row, _ := result.RowsAffected()
		if row > 0 {
			fmt.Println("Update  berhasil")
		} else {
			fmt.Println("Update gagal")
		}
	}
	return nil

}

func Delete(db *sql.DB, noId string) error {

	var query = "DELETE from users where id=?"
	statement, errPrepare := db.Prepare(query)
	if errPrepare != nil {
		// log.Fatal("error prepare", errPrepare.Error())
		return errPrepare
	}

	result, errExec := statement.Exec(noId)
	if errExec != nil {
		// log.Fatal("error exec ", errExec.Error())
		return errExec
	} else {
		row, _ := result.RowsAffected()
		if row > 0 {
			fmt.Println("delete berhasil")
		} else {
			fmt.Println("delete gagal")
		}
	}
	return nil

}

func BacaByid(db *sql.DB, Id string) entities.User {
	result := db.QueryRow("SELECT * FROM users where id=?", Id)

	var userrow entities.User
	errScan := result.Scan(&userrow.Id, &userrow.Nama, &userrow.Email, &userrow.Phone, &userrow.Domisili)
	if errScan != nil {
		if errScan == sql.ErrNoRows {
			log.Fatal("Id tdk ada")
		} else {
			log.Fatal("eror scan", errScan.Error())
		}

	}
	// fmt.Printf("id:%s  nama:%s  email: %s Domisili:%s\n", userrow.Id, userrow.Nama, userrow.Email, userrow.Domisili)
	return userrow
	// }

}
