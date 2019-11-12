package utilclient

import (
	"github.com/sirupsen/logrus"
	"os"
	"strconv"
)

// QueryTimeoutSeconds is the timeout for the client query in seconds (default to 3 minutes)
var QueryTimeoutSeconds int64
// MedCoConnectorURL is the URL of the MedCo connector this client is attached to
var MedCoConnectorURL string

// OidcReqTokenURL is the URL from which the JWT is retrieved
var OidcReqTokenURL string

func init() {
	var err error

	QueryTimeoutSeconds, err = strconv.ParseInt(os.Getenv("CLIENT_QUERY_TIMEOUT_SECONDS"), 10, 64)
	if err != nil || QueryTimeoutSeconds < 0 {
		logrus.Warn("invalid client query timeout")
		QueryTimeoutSeconds = 3 * 60
	}

	MedCoConnectorURL = os.Getenv("MEDCO_CONNECTOR_URL")

	OidcReqTokenURL = os.Getenv("OIDC_REQ_TOKEN_URL")
}
