# pkgo
Open the document for Go package by CLI

## Installation & Updating
### Install
You can use pkgo the on only macOS
```
brew tap yuzuy/pkgo

brew install pkgo
```

### Update
```
brew upgrade pkgo
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
#### --repo, -r - Open the GitHub repository
```
pkgo --repo context   // Open "https://github.com/golang/go/tree/master/src/context"

pkgo -r crypto/bcrypt // Open "https://github.com/golang/crypto/tree/master/bcrypt"

pkgo -r yuzuy/pkgo    // Open "https://github.com/yuzuy/pkgo"
```

#### --official, -o - Open the document in golang.org
```
pkgo --official context // Open "https://golang.org/pkg/context/"

pkgo -o net/http        // Open "https://golang.org/pkg/net/http/"
```
