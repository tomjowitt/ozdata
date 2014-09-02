Ozdata
---------------------

This is a simple library, app and importer to handle Australian postcode data written in Go.

The data is provided by Datalicious:

http://blog.datalicious.com/free-download-all-australian-postcodes-geocod/

Please note, this data originally comes from Auspost and is free for non-commercial use
but use in commercial applications required a licence.

http://www.postconnect.com.au/postcode-data

Useage
---------------------

To include the library in your projects, go get it and import it:

`go get github.com/tomjowitt/ozdata`

`import (
    "github.com/tomjowitt/ozdata/lib"
)`

To query the data using the built-in app simply pass a p (postcode) flag to the application:

`go run app/main.go -p 2041`
