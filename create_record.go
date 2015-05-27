package rage4

import (
  "fmt"
  // "net/url"
)

func (c *Client) CreateRecord(domainId int, record Record) (status Status, err error) {


//convert type to int
// type=<record type [int]>
// priority=<record priority [int, nullable]
// ttl=<ttl [int]

// active=<record online/offline [bool (true|false), nullable]
// failover=<enable/disable failover [bool (true|false)]
// failovercontent=<record failover value [string]>
// geozone=<geo region id [long]
// geolock=<geo lock coordinates bool (true|false)
// geolat=<geo latitude [double, nullable]
// geolong=<geo longitude [double, nullable]
// udplimit=<enable/disable result set limit [bool (true|false)

  // create http request
  parameters := fmt.Sprintf("%d?name=%s&content=%s", domainId, record.Name, record.Content)
  endpoint := fmt.Sprintf("/createrecord/%s", parameters)
  req, err := c.NewRequest(nil, "GET", endpoint)
  if err != nil {
    return Status{}, err
  }

  // issue the API request
  resp, err := c.Http.Do(req)
  if err != nil {
    return Status{}, err
  }
  defer resp.Body.Close()

  // parse the response
  getStatusResponse := Status{}
  err = decode(resp.Body, &getStatusResponse)
  if err != nil {
    return Status{}, err
  }

  status = getStatusResponse
  
  return status, nil
}






