// Player
package player

type Player struct {
	Name string
	Pass string
}

func (p *Player) SetInfo(Name string, Pass string) {
	p.Name = Name
	p.Pass = Pass
}
