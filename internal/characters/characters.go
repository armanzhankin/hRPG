package characters

type Character interface {
	Attack(target Character)
	TakeDamage(damage float64)
	IsAlive() bool
}

type BaseCharacter struct {
	Name   string
	Health float64
	Power  float64
}

func (b *BaseCharacter) Attack(target Character) {
	target.TakeDamage(b.Power)
}

func (b *BaseCharacter) TakeDamage(damage float64) {
	b.Health -= damage
	if b.Health < 0 {
		b.Health = 0
	}
}

func (b *BaseCharacter) IsAlive() bool {
	return b.Health > 0
}

type SwordMan struct {
	BaseCharacter
}

func NewSwordMan(power float64) *SwordMan {
	return &SwordMan{BaseCharacter{
		Name:   "SwordMan",
		Health: 100,
		Power:  power,
	}}
}

type Mage struct {
	BaseCharacter
}

func NewMage(power float64) *Mage {
	return &Mage{BaseCharacter{
		Name:   "Mage",
		Health: 100,
		Power:  power,
	}}
}
