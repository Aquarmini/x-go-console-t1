package providers

type Config struct {
	id   int64
	name string
}

func (this *Config)Output() (r int64, err error) {
	r = this.id + 1
	return
}
