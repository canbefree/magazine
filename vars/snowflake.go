package vars

type SnowflakeIFace interface {
	GenerateID() uint64
}

var (
	Snowflake SnowflakeIFace
)
