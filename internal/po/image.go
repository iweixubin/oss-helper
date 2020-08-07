package po

type Image struct {
	Length     int64  // 文件的长度
	Tail       int64  // 取文件最后 8个byte 转成一个 int64 作为尾部值
	MD5        string // 文件的 MD5 哈希值
	SHA512     string // 文件的 SHA512 哈希值
	Signature  string // 文件标识
	BucketName string // 存储空间名称
	ObjectName string // 存储对象名称
	CreatedAt  string // 创建时间
}
