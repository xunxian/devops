package main

import (
  "fmt"
  "crypto/hmac"
  "crypto/sha256"
  "encoding/hex"
)

func main(){

  // 有key加密
  s := []byte("你好世界")
  key := []byte("dlongtz")
  m := hmac.New(sha256.New, key)
  m.Write(s)
  signature := hex.EncodeToString(m.Sum(nil))
  fmt.Print(signature + "\n\r")

  // 无key加密
  mm := sha256.New()
  mm.Write([]byte("你好世界"))
  res := hex.EncodeToString(mm.Sum(nil))
  fmt.Println(res)
}
