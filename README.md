Ozdata
---------------------

This is a simple library, app and importer to handle geocoded Australian postcode data written in Go.
It is very much a work in progress so feel free to fork and submit pull requests for both the code and the raw data.

The data is provided by Datalicious:

http://blog.datalicious.com/free-download-all-australian-postcodes-geocod/

Please note, this data originally comes from Auspost and is free for non-commercial use
but in commercial applications it requires a licence.

http://www.postconnect.com.au/postcode-data

Library usage
---------------------

To include the library in your projects, go get it:

```bash
$ go get github.com/tomjowitt/ozdata
```

And import it:

```go
import (
    "github.com/tomjowitt/ozdata/ozdata"
)
```

To use the data within your code:

```go
filename := "./data/test.json"
suburbs, err := ozdata.NewSuburbs(filename)
if err != nil {
    fmt.Println(err)
}

var postcode int64 = 2042

suburb, err := suburbs.GetSuburbByCode(postcode)
if err != nil {
    fmt.Println(err)
}

fmt.Println(suburb)
// prints: {Enmore 2042 {-33.899362 151.171098} {New South Wales NSW Sydney {Australia AU} [{2000 2999}]}}

// do something with suburb
```

The response is a Suburb type with the following data structure:

```go
type Suburb struct {
    Name       string
    Postcode   int64
    Coordinate struct {
        Lat  float64
        Long float64
    }
    State struct {
        Name          string
        Code          string
        Capital       string
        Country       struct {
            Name string
            Code string
        }
        PostcodeRange []struct {
            Low  int64
            High int64
        }
    }
}
```

Application usage
---------------------

To query the data using the built-in app simply pass a postcode flag to the Makefile run command:

```bash
$ make run postcode=2016
```

Importer usage
---------------------

To run the importer:

```bash
$ make import
```
This will read data from the data/import.csv spreadsheet and load it into json objects. This only
needs to be run if updates to the raw csv data have occurred.

License (excluding data)
---------------------

    The MIT License (MIT)

    Copyright (c) 2014 Tom Jowitt

    Permission is hereby granted, free of charge, to any person obtaining a copy of
    this software and associated documentation files (the "Software"), to deal in
    the Software without restriction, including without limitation the rights to
    use, copy, modify, merge, publish, distribute, sublicense, and/or sell copies of
    the Software, and to permit persons to whom the Software is furnished to do so,
    subject to the following conditions:

    The above copyright notice and this permission notice shall be included in all
    copies or substantial portions of the Software.

    THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
    IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY, FITNESS
    FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE AUTHORS OR
    COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER
    IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN
    CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.
