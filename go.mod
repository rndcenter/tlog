module github.com/nikandfor/tlog

go 1.13

require (
	github.com/beorn7/perks v1.0.1
	github.com/gin-gonic/gin v1.5.0
	github.com/golang/glog v0.0.0-20160126235308-23def4e6c14b
	github.com/golang/protobuf v1.4.2
	github.com/nikandfor/cli v0.0.0-20200325075312-052d5b29bac6
	github.com/nikandfor/errors v0.1.0
	github.com/nikandfor/json v0.2.0
	github.com/nikandfor/quantile v0.0.0-20200824213034-5a47c65eb02b
	github.com/prometheus/client_golang v1.7.1
	github.com/prometheus/client_model v0.2.0
	github.com/stretchr/testify v1.6.1
	google.golang.org/protobuf v1.25.0
)

// replace github.com/nikandfor/cli => ../cli
// replace github.com/nikandfor/json => ../json
// replace github.com/nikandfor/xrain => ../xrain
// replace github.com/nikandfor/quantile => ../quantile
