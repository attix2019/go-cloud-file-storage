package meta

import (
	"fmt"
	"sort"
	"time"
)

type FileMeta struct{
	FileSha1 string
	FileName string
	FileSize int64
	Location string
	UploadAt string
}

var fileMetas map[string]FileMeta

func init(){
	fileMetas = make(map[string]FileMeta)
}

func UpdateFileMeta(meta FileMeta){
	println("更新hash为:", meta.FileSha1," 的文件")
	fileMetas[meta.FileSha1] = meta
}

func GetFileMeta(fileSha1 string) FileMeta{
	var val FileMeta
	var ok bool
	if val, ok = fileMetas[fileSha1]; !ok{
		println("元信息中不包含哈希值为:", fileSha1, " 的文件")
	}
	return val
}

func GetFileMetas(limit int)[]FileMeta {
	metaArray := make([]FileMeta, 0)
	for _, value := range fileMetas{
		metaArray = append(metaArray, value)
		fmt.Println(value)
	}
	const baseFormat = "2022-05-13 20:18:00"
	sort.Slice(metaArray, func(i,j int) bool {
		iTime, _ := time.Parse(baseFormat,metaArray[i].UploadAt)
		jTime, _ := time.Parse(baseFormat, metaArray[i].UploadAt)
		return iTime.UnixNano() > jTime.UnixNano()
	})
	if len(metaArray) < limit {
		limit = len(metaArray)
	}
	println("取回 ", limit, " 项")
	return metaArray[0: limit]
}

func (meta FileMeta) String() string {
	return fmt.Sprintf("%s %s", meta.FileName, meta.FileSha1)
}

func RemoveItemFromCatalog(filehash string){
	delete(fileMetas, filehash)
}