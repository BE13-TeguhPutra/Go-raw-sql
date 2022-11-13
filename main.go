package main

import (
	entities "be13/rawsql/Entities"
	"be13/rawsql/config"
	"be13/rawsql/controller"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

// object relational mapping mengkonversi data dari data_type yang tidak kompatibel sehinggan bisa dibaca oleh bahasa pemrograman
func main() {
	//go get -u github.com/go-sql-driver/mysql
	//<username>:<password>@tcp(<hostname>:<portDB>)/<db_name>
	// var connectionString = "root:Teguh12345@tcp(127.0.0.1:3306)/db_be13_teguh"
	// db, err := sql.Open("mysql", connectionString)
	// if err != nil {
	// 	log.Fatal("error open connection", err.Error())
	// }

	// errPing := db.Ping()
	// if errPing != nil {
	// 	log.Fatal("Error connect to db", errPing.Error())
	// } else {
	// 	fmt.Println("Koneksi berhasil")
	// }

	dbConnection := config.Connection()
	//buat mekanisme baru
	fmt.Println("Menu:\n 1.Baca data \n 2.Tambah data \n 3.Update data \n 4.Delete Data \n 5.Baca data by id")
	fmt.Println("Masukan pilihan anda:")
	var pilihan int
	fmt.Scanln(&pilihan)

	switch pilihan {
	case 1:
		dataUser, errbaca := controller.BacaData(dbConnection)
		if errbaca != nil {
			log.Fatal("Error baca data")
		}

		// result, errSelect := dbConnection.Query("SELECT id,Nama,email,prhone FROM users")
		// if errSelect != nil {
		// 	log.Fatal("error select", errSelect.Error())
		// }

		// var dataUser []entities.User
		// for result.Next() {
		// 	var userrow entities.User
		// 	errScan := result.Scan(&userrow.Id, &userrow.Nama, &userrow.Email, &userrow.Phone)
		// 	if errScan != nil {
		// 		log.Fatal("eror scan", errScan.Error())
		// 	}
		// 	// fmt.Printf("id:%s  nama:%s  email: %s\n", userrow.Id, userrow.Nama, userrow.Email)
		// 	dataUser = append(dataUser, userrow)
		// }
		for _, v := range dataUser {
			fmt.Printf("id:%s  nama:%s  email: %s\n", v.Id, v.Nama, v.Email)

		}
	case 2:
		{
			newUser := entities.User{}
			fmt.Println("masukkan id user")
			fmt.Scanln(&newUser.Id)
			fmt.Println("masukkan nama user")
			fmt.Scanln(&newUser.Nama)
			fmt.Println("masukkan Email user")
			fmt.Scanln(&newUser.Email)
			fmt.Println("masukkan Phone user")
			fmt.Scanln(&newUser.Phone)
			fmt.Println("masukkan DOmisili user")
			fmt.Scanln(&newUser.Domisili)
			// newUser := entities.User{}
			err := controller.Insertdata(dbConnection, newUser)
			if err != nil {
				log.Fatal("Insert gagal maning")

			}

			// fmt.Println("masukkan id user")
			// fmt.Scanln(&newUser.Id)
			// fmt.Println("masukkan nama user")
			// fmt.Scanln(&newUser.Nama)
			// fmt.Println("masukkan Email user")
			// fmt.Scanln(&newUser.Email)
			// fmt.Println("masukkan Phone user")
			// fmt.Scanln(&newUser.Phone)
			// fmt.Println("masukkan DOmisili user")
			// fmt.Scanln(&newUser.Domisili)

			// var query = "insert into users (id,nama,email,prhone,domisili) values (?,?,?,?,?)"
			// statement, errPrepare := dbConnection.Prepare(query)
			// if errPrepare != nil {
			// 	log.Fatal("error prepare insert", errPrepare.Error())
			// }

			// result, errExec := statement.Exec(newUser.Id, newUser.Nama, newUser.Email, newUser.Phone, newUser.Domisili)
			// if errExec != nil {
			// 	log.Fatal("error exec insert", errExec.Error())
			// } else {
			// 	row, _ := result.RowsAffected()
			// 	if row > 0 {
			// 		fmt.Println("Insert berhasil")
			// 	} else {
			// 		fmt.Println("Insert gagal")
			// 	}
			// }
		}
	case 3:
		{
			updateUser := entities.User{}
			fmt.Println("masukkan id user")
			fmt.Scanln(&updateUser.Id)
			fmt.Println("masukkan nama user")
			fmt.Scanln(&updateUser.Nama)
			fmt.Println("masukkan Email user")
			fmt.Scanln(&updateUser.Email)
			fmt.Println("masukkan Phone user")
			fmt.Scanln(&updateUser.Phone)
			fmt.Println("masukkan DOmisili user")
			fmt.Scanln(&updateUser.Domisili)

			err := controller.Update(dbConnection, updateUser)
			if err != nil {
				log.Fatal("Update tidak berhasil")
			}

			// var query = "UPDATE users set nama=?,email=?,prhone=?,domisili=? where id=?"
			// statement, errPrepare := dbConnection.Prepare(query)
			// if errPrepare != nil {
			// 	log.Fatal("error prepare", errPrepare.Error())
			// }

			// result, errExec := statement.Exec(updateUser.Nama, updateUser.Email, updateUser.Phone, updateUser.Domisili, updateUser.Id)
			// if errExec != nil {
			// 	log.Fatal("error exec insert", errExec.Error())
			// } else {
			// 	row, _ := result.RowsAffected()
			// 	if row > 0 {
			// 		fmt.Println("Update  berhasil")
			// 	} else {
			// 		fmt.Println("Update gagal")
			// 	}
			// }
		}

	case 4:
		{
			var noId string

			fmt.Println("masukkan id user")
			fmt.Scanln(&noId)
			err := controller.Delete(dbConnection, noId)
			if err != nil {
				log.Fatal("Error")
			}

			// var query = "DELETE from users where id=?"
			// statement, errPrepare := dbConnection.Prepare(query)
			// if errPrepare != nil {
			// 	log.Fatal("error prepare", errPrepare.Error())
			// }

			// result, errExec := statement.Exec(noId)
			// if errExec != nil {
			// 	log.Fatal("error exec ", errExec.Error())
			// } else {
			// 	row, _ := result.RowsAffected()
			// 	if row > 0 {
			// 		fmt.Println("delete berhasil")
			// 	} else {
			// 		fmt.Println("delete gagal")
			// 	}
			// }
		}

	case 5:
		var Id string
		fmt.Println("masukkan id user")
		fmt.Scanln(&Id)

		userrow := controller.BacaByid(dbConnection, Id)

		fmt.Printf("id:%s  nama:%s  email: %s Domisili:%s\n", userrow.Id, userrow.Nama, userrow.Email, userrow.Domisili)

		// fmt.Println("baca data by id")
		// result := dbConnection.QueryRow("SELECT * FROM users where id=?", Id)

		// var userrow entities.User
		// errScan := result.Scan(&userrow.Id, &userrow.Nama, &userrow.Email, &userrow.Phone, &userrow.Domisili)
		// if errScan != nil {
		// 	if errScan == sql.ErrNoRows {
		// 		fmt.Println("ID tidak ada")
		// 	} else {
		// 		log.Fatal("eror scan", errScan.Error())
		// 	}

		// } else {
		// 	fmt.Printf("id:%s  nama:%s  email: %s Domisili:%s\n", userrow.Id, userrow.Nama, userrow.Email, userrow.Domisili)

		// }
		// fmt.Printf("id:%s  nama:%s  email: %s\n", userrow.Id, userrow.Nama, userrow.Email)

	}

	defer dbConnection.Close()

}
