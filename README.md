# compress
Compress the image to the specified size using a dichotomy


# how to use?
```bash
go get github.com/pzxy/compress
```

# examples

```golang
    b, err := os.ReadFile("img.jpg")
    if err != nil {
        t.Fatal(err)
    }
    b, err = Do(b, 0, 100, 200, 10)
    if err != nil {
        t.Fatal(err)
    }
    os.WriteFile("img2.jpg", b, 0644)
```