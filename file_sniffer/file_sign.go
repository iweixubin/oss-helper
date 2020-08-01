package file_sniffer

import (
	"crypto/md5"
	"crypto/sha512"
	"encoding/binary"
	"encoding/hex"
	"oss-helper/file_sniffer/signature"
	"path"
	"strings"
)

// FileSign 文件特征
type FileSign struct {
	Content     []byte   // 文件内容
	Length      int64    // 文件长度
	Tail        int64    // 取文件最后 8个byte 转成一个 int64 作为尾部值
	MD5         string   // 文件的 MD5 哈希值
	SHA512      string   // 文件的 SHA512 哈希值
	Signature   string   // 文件的标识(文件头标注的真实格式)
	Extensions  []string // 文件的扩展名
	ContextType string   // 文件的互联网媒体类型(MIME类型)
}

//
func Check(file []byte, fileName string) (FileSign, error) {
	// 获取扩展名
	xx := signature.Sniff(file)

	fs := FileSign{}
	fs.Content = file
	// os.Stat().FileInfo.Size() 返回的是 int64，不知道为什么不用 uint64
	fs.Length = int64(len(file))
	fs.Tail = getTail(file)
	fs.MD5 = md5Hash(file)
	fs.SHA512 = sha512Hash(file)

	fs.Signature = xx.Signature
	fs.ContextType = xx.ContentType
	fs.Extensions = xx.Extensions

	return fs, nil
}

// getExt 获取文件扩展名，小写且没有 .
func getExt(fileName string) string {
	return strings.ToLower(strings.TrimLeft(path.Ext(fileName), "."))
}

// isContains 是否包含
func isContains(ss []string, s string) bool {
	b := false
	for _, v := range ss {
		if v == s {
			b = true
			break
		}
	}

	return b
}

// getTail 取文件最后 8个byte 转成一个 int64 作为尾部值
func getTail(content []byte) int64 {
	var tailBytes = make([]byte, 8)

	indexMax := len(content) - 1
	for i := 0; i < 8; i++ {
		index := indexMax - i
		if index < 0 {
			break
		}
		tailBytes[i] = content[index]
	}
	// C# BitConverter.ToUInt64 是使用 LittleEndian
	x := binary.LittleEndian.Uint64(tailBytes)
	return int64(x)
}

// 加密(MD5);
func md5Hash(file []byte) string {
	checksum := md5.Sum(file)
	return hex.EncodeToString(checksum[:])
}

func sha512Hash(file []byte) string {
	checksum := sha512.Sum512(file)
	return hex.EncodeToString(checksum[:])
}
