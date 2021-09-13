package main

/*
使用go语言进行sha256的加密
*/

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

  message := []byte("hello world")
  hash := sha256.New()
  hash.Write(message)
  // 计算hash值，并转化为16进制
  hashCode := hex.EncodeToString(hash.Sum(nil))
  fmt.Println(hashCode)

}
