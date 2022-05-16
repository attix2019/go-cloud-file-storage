package db

import(
	"filestore-server/db/mysql"
	"fmt"
	"strconv"
)

func InsetFileMeta(filehash , filename string, filesize int64, fileaddr string) bool {
	db := mysql.DbConn()

	stmt, err := db.Prepare("insert ignore into tbl_file(`file_sha1`,`file_name`,`file_size`,`file_addr`,`status`) " +
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
		fmt.Println(strconv.FormatInt(rf, 10) + " rows affected")
	}
	return true
}

type filemetaInTable struct{
	Filehash string
	Filename string
	Filesize int64
	Fileaddr string
}

func QueryFileMeta(filehash string) (*filemetaInTable, error){
	db := mysql.DbConn()

	stmt, err := db.Prepare("select file_sha1, file_name, file_size, file_addr from tbl_file where" +
		" file_sha1 = ? and status = 1")
	if err != nil{
		fmt.Println("fail to prepare sql " + err.Error())
		return nil, err
	}
	defer stmt.Close()

	filemeta := filemetaInTable{}
	err = stmt.QueryRow(filehash).Scan(&filemeta.Filehash, &filemeta.Filename, &filemeta.Filesize, &filemeta.Fileaddr)
	if err != nil{
		fmt.Println("fail to query " + err.Error())
		return &filemetaInTable{}, err
	}
	return &filemeta, nil
}