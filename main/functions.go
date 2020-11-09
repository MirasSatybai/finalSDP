package main

import "fmt"

func weaponToWarrior() {
	weaponFactory, _ := getWeaponFactory("melee")
	arturs := weaponFactory.setWeapon()
	println("I take weapon :", arturs.getName(), " with damage", arturs.getDamage())
}

func weaponToWizard() {
	weaponFactory, _ := getWeaponFactory("ranged")
	arturs := weaponFactory.setWeapon()
	println("I take weapon :", arturs.getName(), " with damage", arturs.getDamage())
}

func createMount() {

}

func mountWithSaddle() *petWithSaddle {
	mount := &pet{}
	mountWithSaddle := &petWithSaddle{
		mount: mount,
	}
	return mountWithSaddle
}

func mountWithSaddleAndArmor() *petWithSaddleAndArmor {
	mount := &pet{}

	mountWithSaddle := &petWithSaddle{
		mount: mount,
	}

	mountWithSaddleAndArmor := &petWithSaddleAndArmor{
		mount: mountWithSaddle.mount,
	}

	return mountWithSaddleAndArmor
}

func mountWithoutEquip() *pet {
	return &pet{}
}

func createWarrior() {
	warriorOfRadiant := &warriorOfRadiant{}
	artur := &warrior{}
	artur.setCharacter(warriorOfRadiant)
}

func getWarrior() *warrior {
	return &warrior{}
}

func createWizard() {
	wizardOfDire := &wizardOfDire{}
	merlin := &wizard{}
	merlin.setCharacter(wizardOfDire)
}

func getWizard() *wizard {
	return &wizard{}
}

type actions interface {
	attack()
	def()
	setCharacter(character)
}

type wizard struct {
	character  character
	attackType string
}

func (w *wizard) attack() {
	fmt.Println("There are enemies!")
	w.character.attackEnemy()
}

func (w *wizard) def() {
	fmt.Println("They are attack")
	w.character.defence()
}

func (w *wizard) setCharacter(c character) {
	w.character = c
}

type warrior struct {
	character character
}

func (w *warrior) attack() {
	fmt.Println("There are enemies!")
	w.character.attackEnemy()
}

func (w *warrior) def() {
	fmt.Println("They are attack")
	w.character.defence()
}

func (w *warrior) setCharacter(c character) {
	w.character = c
}

type character interface {
	attackEnemy()
	defence()
}

type wizardOfDire struct {
	character character
}

func (w *wizardOfDire) attackEnemy() {
	fmt.Println("pooooof!")
}

func (w *wizardOfDire) defence() {
	fmt.Println(" I`m use glimmer")
}

type warriorOfRadiant struct {
	character character
}

func (w *warriorOfRadiant) attackEnemy() {
	fmt.Println("Bzhuuuuuh")
}

func (w *warriorOfRadiant) defence() {
	fmt.Println(" I`m use shild")
}

type game interface {
	typeName(Observer observer)
	exit(Observer observer)
	notifyAll()
}
type pet struct {
	observerList []observer
	name         string
	isHungry     bool
	convenience  int
}

func newPet(name string) *pet {
	return &pet{
		name: name,
	}
}

func (p *pet) updateNotification() {
	fmt.Printf("%s is now hungry\n", p.name)
	p.isHungry = true
	p.notifyAll()
}

func (p *pet) typeName(o observer) {
	p.observerList = append(p.observerList, o)
}

func (p *pet) notifyAll() {
	for _, observer := range p.observerList {
		observer.update(p.name)
	}
}

func (p *pet) mountСonvenience() int {
	return 10
}

type observer interface {
	update(string)
	getPlayerName() string
}

type player struct {
	playerName string
}

func (pl *player) update(petName string) {
	fmt.Printf("%s , your %s is hungry, you should feed\n", pl.playerName, petName)
}

func (pl *player) getPlayerName() string {
	return pl.playerName
}

type mount interface {
	mountСonvenience() int
}

type petWithSaddle struct {
	mount mount
}

func (p *petWithSaddle) mountConvenience() int {
	mountConv := p.mount.mountСonvenience()
	return mountConv + 20
}

type petWithSaddleAndArmor struct {
	mount mount
}

func (p *petWithSaddleAndArmor) mountConvenience() int {
	mountConv := p.mount.mountСonvenience()
	return mountConv - 10
}

type iWeaponFactory interface {
	setWeapon() iWeapone
}

func getWeaponFactory(attackType string) (iWeaponFactory, error) {
	if attackType == "melee" {
		return &warrior{}, nil
	}

	if attackType == "ranged" {
		return &wizard{}, nil
	}

	return nil, fmt.Errorf("wrong attack type")
}

func (w *warrior) setWeapon() iWeapone {
	return &meleeForm{
		weapone{
			name:   "Crystalis",
			damage: 18,
		},
	}
}

type ranged struct {
}

func (w *wizard) setWeapon() iWeapone {
	return &rangedForm{
		weapone{
			"Dagon",
			15,
		},
	}
}

type iWeapone interface {
	setName(name string)
	setDamage(damage int)
	getName() string
	getDamage() int
}

type weapone struct {
	name   string
	damage int
}

func (w *weapone) setName(name string) {
	w.name = name
}

func (w *weapone) getName() string {
	return w.name
}

func (w *weapone) setDamage(damage int) {
	w.damage = damage
}

func (w *weapone) getDamage() int {
	return w.damage
}

type rangedForm struct {
	weapone
}

type meleeForm struct {
	weapone
}

func main() {
	index := 0
	fmt.Println("Welcome! please choose option:")
	fmt.Println("1. Create warrior")
	fmt.Println("2. Create wizard")
	fmt.Println("3. Create mount")
	fmt.Scan(&index)

	if index == 1 {
		createWarrior()
		weaponToWarrior()

	}
	if index == 2 {
		createWizard()
		weaponToWizard()
	}

	if index == 3 {

	}
}
