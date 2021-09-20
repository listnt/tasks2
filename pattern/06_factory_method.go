package pattern

import "fmt"

type iGun interface {
	Shoot()
	Reload()
}

type Ak47 struct {
}

func NewAk47() iGun {
	return &Ak47{}
}

func (ak47 *Ak47) Shoot() {
	fmt.Println("ТРА ТА ТА ТА ТА")
}

func (ak47 *Ak47) Reload() {
	fmt.Println("ПЕРЕЗАРЯЖАЮСЬ")
}

type M14 struct {
}

func NewM14() iGun {
	return &M14{}
}

func (m14 *M14) Shoot() {
	fmt.Println("PEW PEW PEW PEW")
}

func (m14 *M14) Reload() {
	fmt.Println("RELOADING")
}

func getGun(typeOfGun string) (iGun, error) {
	switch typeOfGun {
	case "ak47":
		return NewAk47(), nil
	case "m14":
		return NewM14(), nil
	default:
		return nil, fmt.Errorf("error")
	}
}

func driverCode() {
	ak47, err := getGun("ak47")
	if err != nil {
		fmt.Println(err)
		return
	}
	m14, err := getGun("m14")
	if err != nil {
		fmt.Println(err)
		return
	}
	ak47.Shoot()
	ak47.Reload()
	m14.Shoot()
	m14.Reload()
}
