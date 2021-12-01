package startup

import (
	"github.com/canbefree/magazine/utils/snowflake"
	"github.com/canbefree/magazine/vars"
)

func init() {
	vars.Snowflake = snowflake.NewSnowflakeIP()
}
