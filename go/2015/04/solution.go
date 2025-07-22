package solutions_2015_04

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
)

func PartOne(secret string) int {
    salt := 0
    isHash := false
    hasher := md5.New()

    for !isHash {
        isHash = true
        hasher.Reset()
        hasher.Write([]byte(fmt.Sprintf("%s%d", secret, salt)))
        md5Hash := hex.EncodeToString(hasher.Sum(nil))
        
        for i := range 5 {
            if md5Hash[i] != '0' {
                isHash = false
                break
            }
        }
        if isHash {
            break
        }
        salt++
    }

    return salt
}

func PartTwo(secret string) int {
    salt := 0
    isHash := false
    hasher := md5.New()

    for !isHash {
        isHash = true
        hasher.Reset()
        hasher.Write([]byte(fmt.Sprintf("%s%d", secret, salt)))
        md5Hash := hex.EncodeToString(hasher.Sum(nil))
        
        for i := range 6 {
            if md5Hash[i] != '0' {
                isHash = false
                break
            }
        }
        if isHash {
            break
        }
        salt++
    }

    return salt
}
