package meta

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
