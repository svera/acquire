package bots

type NullBot struct{}

func (s *NullBot) Update(st interface{}) {}
func (s *NullBot) Play() interface{} {
	return nil
}
