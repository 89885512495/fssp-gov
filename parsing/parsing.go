// Read csvfile and add info in DB
package parsing

import (
	"bytes"
	"database/sql"
	"encoding/csv"
	"fmt"
	"io"
	"io/ioutil"
	"log"

	"github.com/dennis/sendmsg"
	_ "github.com/go-sql-driver/mysql"
)

func AddDB(filename string) {

	content, _ := ioutil.ReadFile(filename)

	reader := csv.NewReader(bytes.NewBuffer(content))
	_, err := reader.Read() // skip first line

	if err != nil {
		if err != io.EOF {
			sendmsg.SendMsg("Error when reading file.")
			log.Fatalln(err)
		}
	}

	db, err := sql.Open("mysql", "root:Dunaewlad88!@/fsspGov")
	if err != nil {
		sendmsg.SendMsg("Error when opening database.")
		log.Fatalln(err)
	}
	defer db.Close()

	for i := 0; i < 10; i++ {

		line, err := reader.Read()
		if err != nil {
			if err == io.EOF {
				fmt.Println(err)
				break
			}

		}

		// Prepare statement for inserting data
		stmtIns, err := db.Prepare(
			`INSERT INTO debtors(
				debtor, address, realaddress, enforcementNumber, 
				dateOfArousal, executiveSummaryNumber, typeofExecutiveDocument,
				executiveDate, executiveDocumentNumber, executiveDocumentRequirements,
				subjectOfExecution, amountToBePaid, balanceOwed, departmentOfBailiffs, 
				bailiffDepartmentAddress) VALUES(?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)`)

		if err != nil {
			sendmsg.SendMsg("Error while preparing statement.")
			panic(err.Error()) // proper error handling instead of panic in your app
		}
		defer stmtIns.Close() // Close the statement when we leave main() / the program terminates

		_, er := stmtIns.Exec(line[0], line[1], line[2],
			line[3], line[4], line[5],
			line[6], line[7], line[8],
			line[9], line[10], line[11],
			line[12], line[13], line[14])
		if er != nil {
			sendmsg.SendMsg("Error while insetr data in db.")
			panic(er.Error())
		}
	}
}
