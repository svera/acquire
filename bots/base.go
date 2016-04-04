package bots

type base struct {
	status Status
}

func (b *base) Update(st interface{}) {
	if st, ok := st.(Status); ok {
		b.status = st
	}
}
