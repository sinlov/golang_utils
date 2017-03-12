# Base

[Universally unique identifier](https://en.wikipedia.org/wiki/Universally_unique_identifier)

[IETF-RFC4122](https://tools.ietf.org/html/rfc4122)

more info see http://pubs.opengroup.org/onlinepubs/9696989899/chap5.htm#tagcjh_08_02_01_01

# UUIDV4 use

```golang
var id UUIDV4 = RandV4()
fmt.Printf("uuid HexV4: %v\n", id.HexV4())
fmt.Printf("uuid RawV4: %v\n", id.RawV4())

id1, err := FromStrV4("1870747d-b26c-4507-9518-1ca62bc66e5d")
id2 := MustFromStrV4("1870747db26c450795181ca62bc66e5d")
fmt.Println(id1 == id2) // true
```

more test see [uuid_test.go](uuid_test.go)