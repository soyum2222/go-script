package communication

type Chan interface {
	In(interface{}) error
	Out() error
}
