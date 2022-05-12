package main

import (
	"database/sql"
	"encoding/base64"
	"fmt"
	"log"
	"net/http"
	"okay/crypto"

	_ "github.com/go-sql-driver/mysql"
)

type dbInfo struct {
	user     string
	pwd      string
	url      string
	engine   string
	database string
}

var db1 = dbInfo{"root", "mypassword", "localhost:3306", "mysql", "dbtest"}

func main() {

	db := dbInfo{"root", "qkswl110!@k43nuw8", "localhost:3306", "mysql", "testdb"}
	dataSoruce := db.user + ":" + db.pwd + "@tcp(" + db.url + ")/" + db.database
	conn, err := sql.Open(db.engine, dataSoruce)

	if err != nil || conn.Ping() != nil {
		log.Fatal(err)
	}

	defer conn.Close()

	/*
			_, err = conn.Exec("CREATE TABLE orangeTest(id VARCHAR(255), name VARCHAR(255))")

			if err != nil {
				log.Fatal(err)
			}
				_, err = conn.Exec("INSERT INTO orangeTest(id, name) VALUES ('GREAT', 'GREAT')")
				if err != nil {
					log.Fatal(err)
				}


		//Table에 INSERT 하고 나서 그 작업이 되었는 지 확인 방법
		result, err := conn.Exec("INSERT INTO orangetest VALUES (?, ?)", "carrot4", "orange3")
		if err != nil {
			log.Fatal(err)
		}

		n, err := result.RowsAffected()
		if n == 1 {
			fmt.Println("1 row inserted")
		}
	*/

	/*
		//Select 확인 법
		var name string
		err = conn.QueryRow("SELECT id from orangetest where id = 'carrot4'").Scan(&name)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(name)

		// 다수의 ROW에서 다음 Row로 이동하기위해 NEXT() 메서드를 사용하는데, 반복문을 사용하여 체크합니다.
		var id string
		var name1 string

		rows, err := conn.Query("SELECT id, name FROM orangetest where id = 'carrot1'")
		if err != nil {
			log.Fatal(err)
		}
		defer rows.Close()

		for rows.Next() {
			err := rows.Scan(&id, &name1)
			if err != nil {
				log.Fatal(err)
			}

			fmt.Println(id, name1)
		}
	*/
	//====================간단 암호화/복호화=======================
	//return the generateKey
	privKey, publicKey := crypto.GenerateKeyPair(1024)
	/*
		fmt.Println(privKey)

		fmt.Println(publicKey)
	*/

	//return the ciphertext and plaintext
	var secretMsg string = "hello"
	ciphertext := crypto.EncryptWithPublicKey(secretMsg, publicKey)

	//128개 배열 변환
	fmt.Println(ciphertext)
	fmt.Println(base64.StdEncoding.EncodeToString(ciphertext))

	plaintext := crypto.DecryptWithPrivateKey(ciphertext, privKey)

	fmt.Println(plaintext)
	fmt.Println(string(plaintext))

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "hello world")
	})

	http.ListenAndServe(":3000", nil)

}
