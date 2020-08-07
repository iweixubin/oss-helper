package dal

import (
	"oss-helper/internal/config"
	"oss-helper/internal/po"
	"testing"
)

func Test_ImageQuery(t *testing.T) {
	db := InitDb(config.LoadCfg())
	defer db.Close()

	p := &po.Image{}
	p.Length = 383589
	p.Tail = 5279712195050102914
	p.MD5 = "32351df8f7942ad9bacd1e5cf55f8ed2"
	p.SHA512 = "80735e79b4a8e15c8534627b207a334b235bb36774b4c68bd6a5e7bf9f76d001d1b9a0b548543a73be5258182b5466c6812c5262ada35e8835a389522a2dae49"
	p.Signature = "png"

	ps, err := ImageQuery(p)
	if err == nil {
		_ = ps
	} else {
		t.Error(err)
	}
}

func Test_ImageInsert(t *testing.T) {
	db := InitDb(config.LoadCfg())
	defer db.Close()

	p := &po.Image{}
	p.Length = 383589
	p.Tail = 5279712195050102914
	p.MD5 = "32351df8f7942ad9bacd1e5cf55f8ed2"
	p.SHA512 = "80735e79b4a8e15c8534627b207a334b235bb36774b4c68bd6a5e7bf9f76d001d1b9a0b548543a73be5258182b5466c6812c5262ada35e8835a389522a2dae49"
	p.Signature = "png"
	p.BucketName = "img"
	p.ObjectName = "x01/d5202738ceb64ac0abac9c1baa5f0e22.png"

	err := ImageInsert(p)
	if err != nil {
		t.Error(err)
	}
}

func Test_ImageCount(t *testing.T) {
	db := InitDb(config.LoadCfg())
	defer db.Close()

	count, err := ImageCount()
	if err != nil {
		t.Error(err)
	} else {
		_ = count
	}
}
