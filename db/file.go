package db

import(
	"filestore-server/db/mysql"
	"fmt"
)

func InsetFileMeta(filehash , filename string, filesize int64, fileaddr string) bool {
	db := mysql.DbConn()
	defer db.Close()

	stmt, err := db.Prepare("insert into tbl_file(`file_sha1`,`file_name`,`file_size`,`file_addr`,`status`) " +
		"values(?,?,?,?,1)")
	if err != nil {
		fmt.Println("fail to prepare statement "  + err.Error())
		return false
	}
	defer stmt.Close()

	res, err := stmt.Exec(filehash, filename, filesize, fileaddr )
	if err != nil{
		fmt.Println("fail to execute statement " + err.Error())
		return false
	}

	rf, _ := res.RowsAffected()
	if rf <= 0{
		fmt.Println(string(rf) + "rows affected")
	}
	return true
}