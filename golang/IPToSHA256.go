package main

import (
  "fmt"
  "time"
  "strings"
  "strconv"
  "crypto/sha256"
  "encoding/hex"
)

func main() {
  t := time.Now()
	for i := 0; i <= 255; i++ {
		for j :=0; j <= 255; j++ {
      for k :=0; k <= 255; k++ {
        for l :=0; l <= 255; l++ {
          var builder strings.Builder
          builder.WriteString(strconv.Itoa(i))
          builder.WriteString(".")
          builder.WriteString(strconv.Itoa(j))
          builder.WriteString(".")
          builder.WriteString(strconv.Itoa(k))
          builder.WriteString(".")
          builder.WriteString(strconv.Itoa(l))
          ip := builder.String()
          message := []byte(ip)
          hash := sha256.New()
          hash.Write(message)
          // 计算hash值，并转化为16进制
          hashCode := hex.EncodeToString(hash.Sum(nil))
          fmt.Printf("%s: %s\n", ip, hashCode)
        }
      }
    }
	}
  elapsed := time.Since(t)
  fmt.Printf("程序耗时：%s\n", elapsed)
}
