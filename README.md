# Installation

```
go get -u github.com/cclin81922/proxyhttpsbe/cmd/proxyhttpsbe
export PATH=$PATH:~/go/bin
```

# Commmand Line Usage

```
proxyhttpsbe -host=github.com -port=443

// then you can use curl or web browser to test http://localhost:8443
```

Use curl to test

```
curl -L http://localhost:8443
```

Use web brower to test

```
open http://localhost:8443
```

# For Developer

Run all tests

```
go test github.com/cclin81922/proxyhttpsbe/cmd/proxyhttpsbe
```
