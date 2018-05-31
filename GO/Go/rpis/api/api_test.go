package api_test

import (
    "fmt"
    "rpis/api"
    // "rpis/api/handlers"
    "io"
    "net/http"
    "net/http/httptest"
    "strings"
    "testing"
)

var (
    server   *httptest.Server
    reader   io.Reader //Ignore this for now
    usersUrl string
)

func init() {
    server = httptest.NewServer(api.APIServer()) //Creating new server with the user handlers
    defer server.Close()
    usersUrl = fmt.Sprintf("%s/inventory/devices", server.URL) //Grab the address for the API endpoint
}

// func TestDeviceHandlerGET(t *testing.T) {

//     request, err := http.NewRequest("GET", usersUrl) //Create request with JSON body

//     res, err := http.DefaultClient.Do(request)

//     if err != nil {
//         t.Error(err) //Something is wrong while sending request
//     }

//     if res.StatusCode != 201 {
//         t.Errorf("Success expected: %d", res.StatusCode) //Uh-oh this means our test failed
//     }
// }

func TestCreateUser(t *testing.T) {
    userJson := `{"username": "dennis", "balance": 200}`

    reader = strings.NewReader(userJson) //Convert string to reader

    request, err := http.NewRequest("POST", usersUrl, reader) //Create request with JSON body

    res, err := http.DefaultClient.Do(request)

    if err != nil {
        t.Error(err) //Something is wrong while sending request
    }

    if res.StatusCode != 201 {
        t.Errorf("Success expected: %d", res.StatusCode) //Uh-oh this means our test failed
    }
}
