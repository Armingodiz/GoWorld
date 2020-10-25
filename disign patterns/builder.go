package builder


import "fmt"

type SmartPhone struct {
	Brand string
	Monitor   int
	Camera    int
	Price     int
}
type BuildProcess interface {
	SetBrand() BuildProcess
	SetCamera() BuildProcess
	SetMonitor() BuildProcess
	SetPrice() BuildProcess
	GetGadget() SmartPhone
}

type ManufacturingDirector struct {
	builder BuildProcess
}

func (m *ManufacturingDirector) SetBuilder(b BuildProcess) {
	m.builder = b
}

func (m *ManufacturingDirector) Construct() SmartPhone {
	m.builder.SetBrand().SetMonitor().SetCamera().SetPrice()
	return m.builder.GetGadget()
}

func (m *ManufacturingDirector) PrintProduct() {
	gadget := m.builder.GetGadget()
	fmt.Printf("Structure: %s \n", gadget.Brand)
	fmt.Printf("Monitor: %d \n", gadget.Monitor)
	fmt.Printf("Camera: %d \n", gadget.Camera)
	fmt.Printf("Price: %d $ \n", gadget.Price)
	fmt.Printf("===============\n")
}

////////////////////////////////////////////////////////////////////////////////////////////////////////// Laptop :
type Apple struct {
	smartPhone SmartPhone
}

func (l *Apple) SetBrand() BuildProcess {
	l.smartPhone.Brand = "Apple"
	return l
}

func (l *Apple) SetMonitor() BuildProcess {
	l.smartPhone.Monitor = 1
	return l
}

func (l *Apple) SetCamera() BuildProcess {
	l.smartPhone.Camera = 1
	return l
}

func (l *Apple) SetPrice() BuildProcess {
	l.smartPhone.Price = 600
	return l
}

func (l *Apple) GetGadget() SmartPhone {
	return l.smartPhone
}

//////////////////////////////////////////////////////////////////////////////////////////////////////////// END LAPTOP
////////////////////////////////////////////////////////////////////////////////////////////////////// SMART PHONE :
type Samsung struct {
	smartPhone SmartPhone
}

func (l *Samsung) SetBrand() BuildProcess {
	l.smartPhone.Brand = "Samsung"
	return l
}

func (l *Samsung) SetMonitor() BuildProcess {
	l.smartPhone.Monitor = 2
	return l
}

func (l *Samsung) SetCamera() BuildProcess {
	l.smartPhone.Camera = 2
	return l
}
func (l *Samsung) SetPrice() BuildProcess {
	l.smartPhone.Price = 300
	return l
}
func (l *Samsung) GetGadget() SmartPhone {
	return l.smartPhone
}

//////////////////////////////////////////////////////////////////////////////////////////////////////END SMART PHONE
func main() {
	manufacturingDirector := ManufacturingDirector{}
	laptop := &Apple{}
	manufacturingDirector.SetBuilder(laptop)
	manufacturingDirector.Construct()
	manufacturingDirector.PrintProduct()
	//Structure: Apple
	//Monitor: 1
	//Camera: 1
	//Price: 600 $
	//===============
	samsung := &Samsung{}
	manufacturingDirector.SetBuilder(samsung)
	manufacturingDirector.Construct()
	manufacturingDirector.PrintProduct()
	//Structure: Samsung
	//Monitor: 2
	//Camera: 2
	//Price: 300 $
}


