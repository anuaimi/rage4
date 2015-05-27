package main 

import (
  "fmt"
  "github.com/anuaimi/rage4"
  "net/http"
  "os"
  "github.com/spf13/viper"
)

//
// main - start of app
//
func main() {

  // get API values from config file
  viper.SetConfigName("config")     // can be config.json, config.yaml
  viper.AddConfigPath("./")
  viper.SetDefault("URL", "https://secure.rage4.com/rapi/")
  viper.ReadInConfig()

  accountKey := viper.GetString("AccountKey")
  email := viper.GetString("Email")

  //make sure we were able to read values
  if (len(accountKey) == 0) || (len(email) == 0) {
    fmt.Print("could not read rag4 config values\n")
    os.Exit(-1)
  }

  // create connection to Rage4
  client := rage4.Client{
    AccountKey: viper.GetString("AccountKey"),
    Email: viper.GetString("Email"),
    URL:  viper.GetString("URL") ,
    Http:  http.DefaultClient,
  }

  // GetDomains
  fmt.Print("GetDomains\n")
  domains, _ := client.GetDomains()
  for _, domain := range domains {
    fmt.Printf("  %s (%d)\n", domain.Name, domain.Id)
  }

  os.Exit(0)

  // //ShowCurrentGlobalUsage
  // fmt.Print("\nShowCurrentGlobalUsage\n")
  // usage,  _ := client.ShowCurrentGlobalUsage()
  // fmt.Printf("  usage from %s: %d\n", usage.Date, usage.Value)

  // //ListRecordTypes
  // fmt.Print("\nListRecordTypes\n")
  // recordTypes, err := client.ListRecordTypes()
  // if (recordTypes == nil) {
  //   fmt.Printf("error: %s", err)
  // } else {
  //   for _, recordType := range recordTypes {
  //     fmt.Printf("  %s: %d\n", recordType.Name, recordType.Value)
  //   }
  // }

  // //ListGeoRegions
  // fmt.Print("\nListGeoRegions\n")
  // regions, err := client.ListGeoRegions()
  // if (regions == nil) {
  //   fmt.Printf("error: %s", err)
  // } else {
  //   for _, region := range regions {
  //     fmt.Printf("  %s: %d\n", region.Name, region.Value)
  //   }
  // }


  //CreateRegularDomain
  var domainName = "blabla.com"

  fmt.Print("\nCreateRegularDomain\n")
  status, err := client.CreateRegularDomain( domainName, "admin@blabla.com")
  if (status.Id == 0) {
    fmt.Printf("%s\n", status.Error)
    os.Exit(-1)
  } else {
    fmt.Printf("  status: %t  id: %d err:%s\n", status.Status, status.Id, status.Error)
  }

  var domainId = status.Id

  //GetDomain
  fmt.Print("\nGetDomain\n")
  domain, _ := client.GetDomain( domainId)
  fmt.Printf("  %d: %s %s\n", domainId, domain.Name, domain.Email)

  //GetDomainByName
  fmt.Print("\nGetDomainByName\n")
  domain, _ = client.GetDomainByName( domainName)
  fmt.Printf("  %s: %d\n", domainName, domain.Id)

  //UpdateDomain
  fmt.Print("\nUpdateDomain\n")
  status, err = client.UpdateDomain( domainId, "admin2@test.com")
  if (err == nil) {
    fmt.Printf("  status: %t  id: %d  error: %s\n", status.Status, status.Id, status.Error)
  } else {
    fmt.Printf("%s", err)
  }

  //ShowCurrentUsage
  fmt.Print("\nShowCurrentUsage\n")
  fmt.Printf("  for: %d\n", domainId)
  domainUsage,  _ := client.ShowCurrentUsage( domainId)
  for _, dailyUsage := range domainUsage {
    fmt.Printf("  from %s:  -  ", dailyUsage.Date)
    fmt.Printf("  Total: %d", dailyUsage.Total)
    fmt.Printf("  EU: %d", dailyUsage.EUTotal)
    fmt.Printf("  US: %d", dailyUsage.USTotal)
    fmt.Printf("  SA: %d", dailyUsage.SATotal)
    fmt.Printf("  AP: %d", dailyUsage.APTotal)
    fmt.Printf("  AF: %d\n", dailyUsage.AFTotal)
  }


  //GetRecords
  fmt.Print("\nGetRecords\n")
  records, err := client.GetRecords( domainId)
  if (err != nil) {
    fmt.Printf("%s", err)
  } else {
    for _, record := range records {
      fmt.Printf("  %d: %s %s\n", record.Id, record.Type, record.Content)
    }
  }

  // recordId := 1231889

  // //DeleteRecord
  // fmt.Print("\nDeleteRecord\n")
  // status,  err = client.DeleteRecord( recordId)
  // if (err == nil) {
  //   fmt.Printf("%", err)
  // } else {
  //   fmt.Printf("  status: %t  id: %d  error: %s\n", status.Status, status.Id, status.Error)
  // }

  //DeleteDomain
  fmt.Print("\nDeleteDomain\n")
  status, err = client.DeleteDomain( domainId)
  if (err == nil) {
    fmt.Printf("%s", err)
  } else {
    fmt.Printf("  status: %t  id: %d  error: %s\n", status.Status, status.Id, status.Error)
  }

}
