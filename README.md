# pkgo
Open the document for Go package by CLI

## Installation & Updating
### Install
```
go get github.com/yuzuy/pkgo
```

### Update
```
go get -u github.com/yuzuy/pkgo
```

## Usage
Open the document for the specified package in "pkg.go.dev"

### Examples
```
pkgo context       // Open "https://pkg.go.dev/context"

pkgo crypto/bcrypt // Open "https://pkg.go.dev/golang.org/x/crypto/bcrypt"

pkgo yuzuy/pkgo    // Open "https://pkg.go.dev/github.com/yuzuy/pkgo"
```

### Flags
```
pkgo --repo context // Open "https://github.com/golang/go/context"
```

