Ozdata
---------------------

This is a simple library, app and importer to handle geocoded Australian postcode data written in Go.
It is very much a work in progress and the code is probably quite shoddy so feel free to fork
and submit pull requests.

The data is provided by Datalicious:

http://blog.datalicious.com/free-download-all-australian-postcodes-geocod/

Please note, this data originally comes from Auspost and is free for non-commercial use
but in commercial applications it requires a licence.

http://www.postconnect.com.au/postcode-data

Library useage
---------------------

To include the library in your projects, go get it:

    $ go get github.com/tomjowitt/ozdata

And import it:

    import (
        "github.com/tomjowitt/ozdata/lib"
    )

To use the data within your code:

    data, err := ozdata.NewSuburbData()
    if err != nil {
        fmt.Println(err)
    }

    suburb, err := data.GetSuburbByPostcode(2041)
    if err != nil {
        fmt.Println(err)
    }

    // do something with suburb

The response is based on the following data structure and will return a Suburb type:

    type Suburb struct {
        Name       string
        Postcode   int64
        Coordinate Coordinate
        State      State
    }

    type Coordinate struct {
        Lat  float64
        Long float64
    }

    type State struct {
        Name          string
        Code          string
        Capital       string
        Country       Country
        PostcodeRange []PostcodeRange
    }

    type Country struct {
        Name string
        Code string
    }

    type PostcodeRange struct {
        Low  int64
        High int64
    }

Application useage
---------------------

To query the data using the built-in app simply pass a p (postcode) flag to the application:

    $ go run app/main.go -p 2041

LICENCE (excluding data)
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
