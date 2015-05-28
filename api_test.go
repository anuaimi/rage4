package rage4

import (
  "fmt"
  "os"
  "testing"
  "net/http"
  "github.com/spf13/viper"
)

var client Client

func TestMain(m *testing.M) {

  // setup mock rage4 api server()
  go serverWebPages()

  // get settings for rage4
  viper.SetConfigName("testing")     // can be testing.json, testing.yaml
  viper.AddConfigPath("./example")
  viper.ReadInConfig()

  accountKey := viper.GetString("AccountKey")
  email := viper.GetString("Email")
  url :=  viper.GetString("URL") 
  fmt.Printf("testing rage4 api at %s using account %s\n", url, accountKey)

  // create client to test API calls
  client = Client{
    AccountKey: accountKey,
    Email: email,
    URL: url,
    Http:  http.DefaultClient,
  }

  if (client == Client{}) {
    os.Exit(-1)
  }

  retCode := m.Run()
  fmt.Printf("result = %d", retCode)

  // teardown()

  os.Exit(retCode)
}

func serverWebPages() {

  // create web server - port 9000
  http.HandleFunc("/testAuth", testAuthHandler)

  baseURL := "rapi"
  http.HandleFunc( baseURL + "/getdomains", getDomainsHandler)
  http.HandleFunc( baseURL + "/getdomain", getDomainHandler)
  http.HandleFunc( baseURL + "/getdomainbyname", getDomainsHandler)
  http.HandleFunc( baseURL + "/createregulardomain", createRegularDomainHandler)
  http.HandleFunc( baseURL + "/createregulardomainext", createRegularDomainExtHandler)
  http.HandleFunc( baseURL + "/createreversedomain4", createReverseDomain4Handler)
  http.HandleFunc( baseURL + "/createreversedomain6", createReverseDomain6Handler)
  http.HandleFunc( baseURL + "/updatedomain", updateDomainHandler)
  http.HandleFunc( baseURL + "/deletedomain", deleteDomainHandler)
  http.HandleFunc( baseURL + "/showcurrentusage", showCurrentUsageHandler)
  http.HandleFunc( baseURL + "/showcurrentglobalusage", showCurrentGlobalUsageHandler)
  http.HandleFunc( baseURL + "/listrecordtypes", listRecordTypesHandler)
  http.HandleFunc( baseURL + "/listgeoregions", listGeoRegionsHandler)
  http.HandleFunc( baseURL + "/getrecords", getRecordsHandler)
  http.HandleFunc( baseURL + "/createrecord", createRecordHandler)
  http.HandleFunc( baseURL + "/updaterecord", updateRecordHandler)
  http.HandleFunc( baseURL + "/deleterecord", deleteRecordHandler)

  http.ListenAndServe(":9000", nil)

}

func testAuthHandler(w http.ResponseWriter, r *http.Request) {
  // test run completed
  os.Exit(0)
}


func getDomainsHandler(w http.ResponseWriter, r *http.Request) {
  // no input required
  // return array of domain info
}

func getDomainHandler(w http.ResponseWriter, r *http.Request) {
  // no input required
  // return array of domain info
}

func createRegularDomainHandler(w http.ResponseWriter, r *http.Request) {
  // no input required
  // return array of domain info
}

func createRegularDomainExtHandler(w http.ResponseWriter, r *http.Request) {
  // no input required
  // return array of domain info
}

func createReverseDomain4Handler(w http.ResponseWriter, r *http.Request) {
  // no input required
  // return array of domain info
}

func createReverseDomain6Handler(w http.ResponseWriter, r *http.Request) {
  // no input required
  // return array of domain info
}

func updateDomainHandler(w http.ResponseWriter, r *http.Request) {
  // no input required
  // return array of domain info
}

func deleteDomainHandler(w http.ResponseWriter, r *http.Request) {
  // no input required
  // return array of domain info
}

func showCurrentUsageHandler(w http.ResponseWriter, r *http.Request) {
  // no input required
  // return array of domain info
}
func showCurrentGlobalUsageHandler(w http.ResponseWriter, r *http.Request) {
  // no input required
  // return array of domain info
}
func listRecordTypesHandler(w http.ResponseWriter, r *http.Request) {
  // no input required
  // return array of domain info
}
func listGeoRegionsHandler(w http.ResponseWriter, r *http.Request) {
  // no input required
  // return array of domain info
}
func getRecordsHandler(w http.ResponseWriter, r *http.Request) {
  // no input required
  // return array of domain info
}
func createRecordHandler(w http.ResponseWriter, r *http.Request) {
  // no input required
  // return array of domain info
}
func updateRecordHandler(w http.ResponseWriter, r *http.Request) {
  // no input required
  // return array of domain info
}
func deleteRecordHandler(w http.ResponseWriter, r *http.Request) {
  // no input required
  // return array of domain info
}

/////////////

func TestAuthentication(t *testing.T) {
  //make sure username & password passed in and 
}

func TestGetDomains(t *testing.T) {

  fmt.Print("testing getdomains")
  // domains, err := client.GetDomains()
  client.GetDomains()

}

