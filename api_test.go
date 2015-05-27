package rage4

import (
  "os"
  "testing"
)

func TestMain(m *testing.M) {

  // setup()
    // create web server - port 9000

  retCode := m.Run()

  // teardown()
    //destroy web server

  os.Exit(retCode)
}


