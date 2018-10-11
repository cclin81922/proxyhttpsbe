# Installation

```
go get -u github.com/cclin81922/proxyhttpsbe/cmd/proxyhttpsbe
export PATH=$PATH:~/go/bin
```

# Commmand Line Usage

```
proxyhttpsbe -host=www.google.com.tw -port=443

// then you can use curl to test http://localhost:8443
```

# For Developer

Run all tests

```
go test github.com/cclin81922/proxyhttpsbe/cmd/proxyhttpsbe
```
