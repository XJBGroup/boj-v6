package unittest

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"gopkg.in/yaml.v2"
	"io"
	"os"
	"syscall"
)

type TMetaHTTPHeader = StringMap
type TMetaHTTPMethod = string
type TMetaHTTPData = DataBody
type TMetaHTTPEncoding = string

const (
	MetaHTTPHeader   = "HTTPHeader"
	MetaHTTPMethod   = "HTTPMethod"
	MetaHTTPEncoding = "HTTPEncoding"
	MetaUri          = "HTTPUri"
	MetaMethod       = "Method"
	MetaEncoding     = "Encoding"
	MetaHeader       = "Header"
	MetaData         = "Data"
)

func load() {
	f, err := os.Open("test.yaml")
	if err != nil {
		panic(err)
	}
	var spec SpecV1
	err = yaml.NewDecoder(f).Decode(&spec)
	if err != nil {
		panic(err)
	}

	gd, err := generateCaseV1(&spec)
	if err != nil {
		panic(err)
	}
	for _, x := range gd.TestCases {
		fmt.Println("Name:", x.Name)
		fmt.Println("Path:", x.Path)
		fmt.Println("Method:", x.Meta[MetaMethod])
		fmt.Println("Encoding:", x.Meta[MetaEncoding])
		fmt.Println("Data:", x.Meta[MetaData])
		fmt.Println("Header:", x.Meta[MetaHeader])
		fmt.Println("HTTPUri:", x.Meta[MetaUri])
		fmt.Println("HTTPMethod:", x.Meta[MetaHTTPMethod])
		fmt.Println("HTTPEncoding:", x.Meta[MetaHTTPEncoding])
		fmt.Println("HTTPHeader:", x.Meta[MetaHTTPHeader])
		fmt.Println("Assertion:", x.Assertions)
		fmt.Println("----------------------------------------------------------------------")
	}
}

func Load(filepath string, withCache bool) *GoDynamicTestData {
	f, err := os.Open(filepath)
	if err != nil {
		panic(err)
	}

	var fileCharacteristic string
	if withCache {
		fs, err := f.Stat()
		if err != nil {
			panic(err)
		}

		var hbSize int64 = 4 * 1024
		if fs.Size() < hbSize {
			hbSize = fs.Size()
		}

		var hashBlock = make([]byte, hbSize)
		var md5HashDriver = md5.New()
		var teeStash = io.TeeReader(f, md5HashDriver)
		for err == nil {
			_, err = teeStash.Read(hashBlock)
			if err == io.EOF {
				err = nil
				break
			}
		}

		fileCharacteristic = hex.EncodeToString(md5HashDriver.Sum(nil))
		if _, err := os.Stat(filepath + ".cache"); err == nil {
			g, err := loadCache(filepath + ".cache")
			if err != nil {
				panic(err)
			}
			if g.Cache == fileCharacteristic {
				return g
			}
		}
	}
	var off int64
	off, err = f.Seek(0, syscall.FILE_BEGIN)
	if err != nil {
		panic(err)
	}
	if off != 0 {
		panic("reset offset not zero")
	}
	var spec SpecV1
	err = yaml.NewDecoder(f).Decode(&spec)
	if err != nil {
		panic(err)
	}

	gd, err := generateCaseV1(&spec)
	if err != nil {
		panic(err)
	}

	if withCache {
		gd.Cache = fileCharacteristic
		err = saveCache(filepath+".cache", gd)
		if err != nil {
			panic(err)
		}
	}

	return gd
}
