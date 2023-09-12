package utils

import (
	"hash/crc32"
	"io"
	"os"
)

// IsExist check file or directory is exists
func IsExist(path string) bool {
	_, err := os.Stat(path)
	if err != nil {
		return os.IsExist(err)
	}
	return true
}

// GetCrc32 get file crc32
func GetCrc32(path string) (uint32, error) {
	file, err := os.Open(path)
	if err != nil {
		return 0, err
	}
	defer file.Close()
	crc32Hash := crc32.NewIEEE()
	// 将文件内容传递给计算器
	_, err = io.Copy(crc32Hash, file)
	if err != nil {
		return 0, err
	}
	// 计算 CRC32 值
	return crc32Hash.Sum32(), nil
}
