
# text2emoji

This is a simple Golang script that converts text into emojis. It uses the emoji library to convert text into emojis.

## Usage
```Go
func main() {
    s1 := "hello world ABC 123 你好"
    es1 := Encrypt(s1)
    fmt.Println(es1)
    s2 := Decrypt(es1)
    fmt.Println(s2)
}
```
Output :
```text
 😀🍇😂🍐🚎🍓😂💰🏈🍒⭐💴⌚🍓😂🐻🏈🐻🍎😇😁😊😁😆⌚🐭⌚😁🏉🚑🍇😁🎱🚕🍓📲💷
hello world ABC 123 你好
```