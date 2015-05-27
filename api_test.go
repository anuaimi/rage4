package rage4

import (
  "os"
  "testing"
  "net/http"
  "github.com/spf13/viper"
)

func TestMain(m *testing.M) {

  // setup()
  // create web server - port 9000
  baseURL := ""
  http.HandleFunc("/", apiHandler)
  http.ListenAndServer(":9000", nil)

  // create client to test API calls
  client := rage4.Client{
    AccountKey: viper.GetString("AccountKey"),
    Email: viper.GetString("Email"),
    URL:  viper.GetString("URL") ,
    Http:  http.DefaultClient,
  }

  retCode := m.Run()

  // teardown()
    //destroy web server

  os.Exit(retCode)
}


