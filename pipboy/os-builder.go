package pipboy


import (
	"github.com/rivo/tview"
)

type Screen interface {
  BuildScreen() 
  GetScreen() *tview.Primitive
  GetName() *string
}

type PipBoyOS struct {
  Screens    map[string]*tview.Primitive
  HomeScreen *tview.Primitive
}

func NewPipBoyOS() *PipBoyOS {
  pb := &PipBoyOS{
    Screens: map[string]*tview.Primitive{},
    HomeScreen: nil,
  }

  var screens []*Screen

  rs := newRadioScreen()
  rs.BuildScreen()

  screens = append(screens, &rs)

  for _, screen := range screens {
    pb.AddScreen(*screen)
  }

  pb.UpdateHomeScreen(rs)

  return pb
}

func (pb *PipBoyOS) Run() error {
  
	if err := tview.NewApplication().SetRoot(*pb.HomeScreen, true).SetFocus(*pb.HomeScreen).Run(); err != nil {
		return err
	}

  return nil
}

func (pb *PipBoyOS) AddScreen(screen Screen) {
  pb.Screens[*screen.GetName()] = screen.GetScreen()
}

func (pb *PipBoyOS) UpdateHomeScreen(screen Screen) {
  pb.HomeScreen = screen.GetScreen()
}
