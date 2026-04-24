package usecase

type DemoUsecase struct{}

func NewDemoUsecase() *DemoUsecase {
	return &DemoUsecase{}
}

func (d *DemoUsecase) Query() any {

	return "data19"
}
