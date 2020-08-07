package dal

import (
	"fmt"
	"oss-helper/internal/po"
	"strconv"
)

func pickTableName(t *po.Image) string {
	remainder := int(t.Length % 10)
	return "image_" + strconv.Itoa(remainder)
}

func ImageCount() (uint64, error) {
	_sql := `
SELECT 
(SELECT COUNT(*) FROM image_0) + 
(SELECT COUNT(*) FROM image_1) + 
(SELECT COUNT(*) FROM image_2) + 
(SELECT COUNT(*) FROM image_3) + 
(SELECT COUNT(*) FROM image_4) + 
(SELECT COUNT(*) FROM image_5) + 
(SELECT COUNT(*) FROM image_6) + 
(SELECT COUNT(*) FROM image_7) + 
(SELECT COUNT(*) FROM image_8) + 
(SELECT COUNT(*) FROM image_9) 
`
	var count uint64
	err := db.QueryRow(_sql).Scan(&count)

	return count, err
}

// 查是否存在
func ImageQuery(t *po.Image) ([]*po.Image, error) {
	tn := pickTableName(t)
	_sql := fmt.Sprintf(`
SELECT
	* 
FROM
	%v 
WHERE
	Length = ? 
	AND Tail = ? 
	AND MD5 = ? 
	AND SHA512 = ? 
	AND Signature = ?
`, tn)

	rows, err := db.Query(_sql, t.Length, t.Tail, t.MD5, t.SHA512, t.Signature)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var ps []*po.Image

	for rows.Next() {
		p := &po.Image{}
		err = rows.Scan(&p.Length, &p.Tail, &p.MD5, &p.SHA512, &p.Signature, &p.BucketName, &p.ObjectName, &p.CreatedAt)
		if err != nil {
			return ps, err
		}
		ps = append(ps, p)
	}

	return ps, nil
}

// 插入
func ImageInsert(t *po.Image) error {
	t.CreatedAt = createdAt()
	tn := pickTableName(t)
	_sql := fmt.Sprintf("INSERT INTO %v VALUES(?, ?, ?, ?, ?, ?, ?, ?)", tn)
	result, err := db.Exec(_sql, t.Length, t.Tail, t.MD5, t.SHA512, t.Signature, t.BucketName, t.ObjectName, t.CreatedAt)
	_ = result
	return err
}
