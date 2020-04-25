/*
  author='du'
  date='2020/4/25 14:45'
*/
package rpcdemo

type DemoService struct{}

type Args struct {
	A, B int
}

func (DemoService) Add(args Args, result *float64) error {
	*result = float64(args.A) + float64(args.B)
	return nil
}
