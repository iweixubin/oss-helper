package file_sniffer

import (
	"io/ioutil"
	"testing"
)

func Test_getExt(t *testing.T) {
	fileName := "120px-SmallFullColourGIF.gif"
	var ext string
	ext = getExt(fileName)
	if ext != "gif" {
		t.Error("获取扩展名错误！")
	}

	fileName = "120px-SmallFullColourGIF.GiF"
	ext = getExt(SamplesPath + `\` + fileName)
	if ext != "gif" {
		t.Error("获取扩展名错误！")
	}
}

func Test_getTail(t *testing.T) {
	dic := make(map[string]int64)
	dic["1"] = 49
	dic["123"] = 3224115
	dic["12345678"] = 3544952156018063160
	dic["1234567890"] = 3689632501694216496
	dic["a"] = 97
	dic["abcdef"] = 107075202213222
	dic["abcdefgh"] = 7017280452245743464
	dic["abcdefghijk"] = 7234300970759973483
	dic["A"] = 65
	dic["ABCDEF"] = 71752852194630
	dic["ABCDEFGH"] = 4702394921427289928
	dic["ABCDEFGHIJK"] = 4919415439941519947
	dic["þíþí"] = -4341817839657761875
	dic["ÿØÿà..JFIF.."] = 3327678840410877486

	for k, v := range dic {
		bytes := []byte(k)
		x := getTail(bytes)

		if v != x {
			t.Errorf("\n不匹配: \n%v \n%v \n%v", k, v, x)
			return
		}
	}

	zero := getTail([]byte{})

	if zero != 0 {
		t.Error("\n应该是数字 0")
	}

	buffer := []byte{191, 118, 253, 129, 0, 13, 245, 175, 255, 217}
	bufferValue := getTail(buffer)
	if bufferValue != -179862450161582119 {
		t.Errorf("\n不匹配: \n%v \n%v", bufferValue, -179862450161582119)
		return
	}

}

// 为什么不能换行
// go 语言 string 类型中的换行只有一个字符： 10     \n
// C# 语言 string 类型中的换行会有两个字符： 13 10  \r\n
// 不知道在线 MD5 加密是什么什么换行符，所以很难测试； 奇怪的是 SHA512 可以换行，而且和在线 SHA512 一致
// 但如果将有换行的 将进酒 保存到 txt文本中， go 与 C# 都用读取文件的 byte， 那么读到的 byte 是一致的； C# 不能先读取成 string 再转 byte， 那么是不一致的
var QiangJinJiu = `君不见黄河之水天上来，奔流到海不复回。君不见高堂明镜悲白发，朝如青丝暮成雪。人生得意须尽欢，莫使金樽空对月。天生我材必有用，千金散尽还复来。烹羊宰牛且为乐，会须一饮三百杯。岑夫子，丹丘生，将进酒，杯莫停。`

func Test_md5Hash(t *testing.T) {
	dic := make(map[string]string)
	dic["123"] = "202cb962ac59075b964b07152d234b70"
	dic["abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"] = "2ad372c377013baa4ee32ab6649d2449"
	dic[QiangJinJiu] = "3adea61cd25f7ed5a25406e5ead8eaba"

	for k, v := range dic {
		x := md5Hash([]byte(k))
		if x != v {
			t.Errorf("不匹配 %v %v %v", k, v, x)
			return
		}
	}
}

func Test_sha512Hash(t *testing.T) {
	dic := make(map[string]string)
	dic["123"] = "3c9909afec25354d551dae21590bb26e38d53f2173b8d3dc3eee4c047e7ab1c1eb8b85103e3be7ba613b31bb5c9c36214dc9f14a42fd7a2fdb84856bca5c44c2"
	dic["abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"] = "b8afc1e7b4e4d6a99a6d514a4450431dc189a5628a6777e785c0cd1540045c0eb60274c7d0a951357d2bc4c9407f212e80231eb3c12c877eb7cffda4081587ae"
	dic[QiangJinJiu] = "ac97827488bf52d3002616aef83dda2fde15d80612c84cf7fa145a7ac1e3ca6b2b7d04d468ccb44b87f280226eef192c6370b37c2064d0669779f4164dcd050d"

	for k, v := range dic {
		x := sha512Hash([]byte(k))
		if x != v {
			t.Errorf("不匹配 %v %v %v", k, v, x)
			return
		}
	}
}

func Test_md5sha512(t *testing.T) {
	// 测试加密是否会改变原本  []byte 的内容
	type temp struct {
		Key    string
		Md5    string
		SHA512 string
	}

	samples := []temp{
		temp{
			Key:    "123",
			Md5:    "202cb962ac59075b964b07152d234b70",
			SHA512: "3c9909afec25354d551dae21590bb26e38d53f2173b8d3dc3eee4c047e7ab1c1eb8b85103e3be7ba613b31bb5c9c36214dc9f14a42fd7a2fdb84856bca5c44c2",
		},
		temp{
			Key:    "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ",
			Md5:    "2ad372c377013baa4ee32ab6649d2449",
			SHA512: "b8afc1e7b4e4d6a99a6d514a4450431dc189a5628a6777e785c0cd1540045c0eb60274c7d0a951357d2bc4c9407f212e80231eb3c12c877eb7cffda4081587ae",
		},
		temp{
			Key:    QiangJinJiu,
			Md5:    "3adea61cd25f7ed5a25406e5ead8eaba",
			SHA512: "ac97827488bf52d3002616aef83dda2fde15d80612c84cf7fa145a7ac1e3ca6b2b7d04d468ccb44b87f280226eef192c6370b37c2064d0669779f4164dcd050d",
		},
	}

	for _, v := range samples {
		buffer := []byte(v.Key)

		x := md5Hash(buffer)
		if x != v.Md5 {
			t.Errorf("不匹配 %v %v %v", v.Key, v.Md5, x)
		}

		y := sha512Hash(buffer)
		if y != v.SHA512 {
			t.Errorf("不匹配 %v %v %v", v.Key, v.SHA512, y)
		}

		z := md5Hash(buffer)
		if z != v.Md5 {
			t.Errorf("不匹配 %v %v %v", v.Key, v.Md5, z)
		}
	}
}

// 文件 MD5    计算 http://www.metools.info/code/c26.html
// 文件 SHA512 计算 http://www.metools.info/code/c92.html

var SamplesPath = `F:\Docs\oss-helper\file_sniffer\samples`

func assertFileSign(t *testing.T, x, y FileSign) {
	if x.Length != y.Length {
		t.Errorf("\nLength 不一致：\n%v\n%v", x.Length, y.Length)
		return
	}

	if x.Tail != y.Tail {
		t.Errorf("\nTail 不一致：\n%v\n%v", x.Tail, y.Tail)
		return
	}

	if x.MD5 != y.MD5 {
		t.Errorf("\nMD5 不一致：\n%v\n%v", x.MD5, y.MD5)
		return
	}

	if x.SHA512 != y.SHA512 {
		t.Errorf("\nSHA512 不一致：\n%v\n%v", x.SHA512, y.SHA512)
		return
	}

	if x.Signature != y.Signature {
		t.Errorf("\nSignature 不一致：\n%v\n%v", x.Signature, y.Signature)
		return
	}

	if x.ContextType != y.ContextType {
		t.Errorf("\nContextType 不一致：\n%v\n%v", x.ContextType, y.ContextType)
		return
	}

	if len(x.Extensions) != len(y.Extensions) {
		t.Errorf("\nExtensions 不一致：\n%v\n%v", x.Extensions, y.Extensions)
		return
	}

	for i, _ := range x.Extensions {
		if x.Extensions[i] != y.Extensions[i] {
			t.Errorf("\nExtensions 不一致：\n%v\n%v", x.Extensions, y.Extensions)
			return
		}
	}
}

func Test_bmp(t *testing.T) {
	// 图片来源
}

func Test_gif(t *testing.T) {
	// 图片来源 https://en.wikipedia.org/wiki/File:SmallFullColourGIF.gif
	// 图片来源 https://en.wikipedia.org/wiki/File:Sunflower_as_gif_websafe.gif
	fsAssert := FileSign{
		Length:      31195,
		Tail:        8070524208268443707,
		MD5:         "3a22b4721c6659e499b0ddf6f2de0f2a",
		SHA512:      "4daad1e6f2b14f617604b6b9187f5786b99d77fe77f31901a0ca4e65cb01c414454aafb14e977bcf7cfc07f18420a7f7127c03af61e4b7d92ba9309b6f8f1e03",
		Extensions:  []string{"gif"},
		Signature:   "gif",
		ContextType: "image/gif",
	}

	fileName := "120px-SmallFullColourGIF.gif"
	data, err := ioutil.ReadFile(SamplesPath + `\` + fileName)

	if err != nil {
		t.Error(err)
		return
	}

	fs, err := Check(data, fileName)
	if err != nil {
		t.Error(err)
		return
	}

	assertFileSign(t, fs, fsAssert)
}

func Test_icon(t *testing.T) {
	// 图片来源 https://www.google.com/favicon.ico
	fsAssert := FileSign{
		Length:      5430,
		Tail:        -72056494530493441,
		MD5:         "f3418a443e7d841097c714d69ec4bcb8",
		SHA512:      "82d017c4b7ec8e0c46e8b75da0ca6a52fd8bce7fcf4e556cbdf16b49fc81be9953fe7e25a05f63ecd41c7272e8bb0a9fd9aedf0ac06cb6032330b096b3702563",
		Extensions:  []string{"ico", "icon"},
		Signature:   "ico",
		ContextType: "image/x-icon",
	}

	fileName := "icon.ico"
	data, err := ioutil.ReadFile(SamplesPath + `\` + fileName)

	if err != nil {
		t.Error(err)
		return
	}

	fs, err := Check(data, fileName)
	if err != nil {
		t.Error(err)
		return
	}

	assertFileSign(t, fs, fsAssert)
}

func Test_jpg(t *testing.T) {
	// 图片来源 https://en.wikipedia.org/wiki/JPEG  质量 100 那张
	fsAssert := FileSign{
		Length:      3430,
		Tail:        0,
		MD5:         "013c66fe5c38c10bf7e0646648907a17",
		SHA512:      "bc57d6975f0d3856ea2dec85a705d9261fb99f33eda201e4812a73857a5697da3553a7ea460e91d3daa2904eb1fcf79abc2f1ead0815abd32b45a23f8710811b",
		Extensions:  []string{"jpg", "jpeg", "jpe", "jif", "jfif", "jfi"},
		Signature:   "jpg",
		ContextType: "image/jpeg",
	}

	fileName := "80px-JPEG_example_JPG_RIP_100.jpg"
	data, err := ioutil.ReadFile(SamplesPath + `\` + fileName)

	if err != nil {
		t.Error(err)
		return
	}

	fs, err := Check(data, fileName)
	if err != nil {
		t.Error(err)
		return
	}

	assertFileSign(t, fs, fsAssert)
}

func Test_png(t *testing.T) {
	// 图片来源 https://en.wikipedia.org/wiki/File:PNG_transparency_demonstration_1.png
	fsAssert := FileSign{
		Length:      383589,
		Tail:        5279712195050102914,
		MD5:         "32351df8f7942ad9bacd1e5cf55f8ed2",
		SHA512:      "80735e79b4a8e15c8534627b207a334b235bb36774b4c68bd6a5e7bf9f76d001d1b9a0b548543a73be5258182b5466c6812c5262ada35e8835a389522a2dae49",
		Extensions:  []string{"png"},
		Signature:   "png",
		ContextType: "image/png",
	}

	fileName := "PNG_transparency_demonstration_1.png"
	data, err := ioutil.ReadFile(SamplesPath + `\` + fileName)

	if err != nil {
		t.Error(err)
		return
	}

	fs, err := Check(data, fileName)
	if err != nil {
		t.Error(err)
		return
	}

	assertFileSign(t, fs, fsAssert)
}
