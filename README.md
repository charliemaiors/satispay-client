# Satispay golang client
[![Build Status](https://travis-ci.org/charliemaiors/satispay-client.svg?branch=master)](https://travis-ci.org/charliemaiors/satispay-client)

This is the implementation of a client library (written in golang) in order to interact with [Satispay online API](https://s3-eu-west-1.amazonaws.com/docs.online.satispay.com/index.html).

## Usage

In order to include this client in your project run:

```bash
go get github.com/charliemaiors/satispay-client
```

And use it in your code:

```golang
import client "github.com/charliemaiors/satispay-client"

func main(){
        client, err := client.NewClient("your-satispay-token", false)
        valid := client.CheckBearer()
        user, err := client.CreateUser("user-phone-number")
        ...
}
```