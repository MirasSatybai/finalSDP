package main

import "fmt"

//bridge
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

//observer
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

func checkMountState() {
	player := player{
		"user",
	}
	player.update("mount")
	main()
}

//decorator
type mount interface {
	mountСonvenience() int
}

func (p *pet) mountСonvenience() int {
	return 10
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

//factory
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
	fmt.Println(" Please, choose option:")
	fmt.Println("1. Create warrior")
	fmt.Println("2. Create wizard")
	fmt.Println("3. Create mount")
	fmt.Println("4. Check mount state")
	fmt.Scan(&index)

	if index == 1 {
		createWarrior()
		weaponToWarrior()
		index := 0
		fmt.Println("Choose action :")
		fmt.Println("1. Create wizard")
		fmt.Println("2. Create mount")
		fmt.Println("3. exit")
		fmt.Scan(&index)

		if index == 1 {
			createWizard()
			weaponToWizard()
		}

		if index == 2 {
			createMount()
		}
		if index == 3 {
			main()
		}
	}
	if index == 2 {
		createWizard()
		weaponToWizard()
		index := 0
		fmt.Println("Choose action :")
		fmt.Println("1. Create warrior")
		fmt.Println("2. Create mount")
		fmt.Println("3. exit")
		fmt.Scan(&index)

		if index == 1 {
			createWarrior()
			weaponToWarrior()
		}

		if index == 2 {
			createMount()
		}
		if index == 3 {
			main()
		}
	}

	if index == 3 {
		createMount()
	}

	if index == 4 {
		checkMountState()
	}
}

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
	fmt.Println("Choose mount :")
	fmt.Println("1. mount without anything")
	fmt.Println("2. mount with saddle")
	fmt.Println("3. mount with saddle and armor")
	fmt.Println("4.exit")
	index := 0
	fmt.Scan(&index)
	if index == 1 {
		mountWithoutEquip()
	}

	if index == 2 {
		mountWithSaddle()
	}

	if index == 3 {
		mountWithSaddleAndArmor()
	}

	if index == 4 {
		main()
	}
}

func mountWithSaddle() *petWithSaddle {
	mount := &pet{}
	mountWithSaddle := &petWithSaddle{
		mount: mount,
	}
	main()
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
	main()
	return mountWithSaddleAndArmor
}

func mountWithoutEquip() *pet {
	main()
	return &pet{}
}

func createWarrior() {
	warriorOfRadiant := &warriorOfRadiant{}
	artur := &warrior{}
	artur.setCharacter(warriorOfRadiant)
	artur.attack()
	artur.def()
}

func getWarrior() *warrior {
	return &warrior{}
}

func createWizard() {
	wizardOfDire := &wizardOfDire{}
	merlin := &wizard{}
	merlin.setCharacter(wizardOfDire)
	merlin.attack()
	merlin.def()
}

func getWizard() *wizard {
	return &wizard{}
}
