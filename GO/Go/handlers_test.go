package device

import "testing"

func TestCreateDevice(t *testing.T) {

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

