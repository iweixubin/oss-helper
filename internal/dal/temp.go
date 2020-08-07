package dal

import (
	"oss-helper/internal/po"
)

func TempInsert(t *po.Temp) error {
	_sql := "INSERT INTO temp VALUES(?, ?)"
	_, err := db.Exec(_sql, t.Id, t.FilePath)
	return err
}
