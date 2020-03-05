package mockgopher

import "syreclabs.com/go/faker"

// Faker it's just a struct thats wraps faker packages interfaces
type Faker struct {
	Address     faker.FakeAddress
	App         faker.FakeApp
	Avatar      faker.FakeAvatar
	Bitcoin     faker.FakeBitcoin
	Business    faker.FakeBusiness
	Code        faker.FakeCode
	Commerce    faker.FakeCommerce
	Company     faker.FakeCompany
	Date        faker.FakeDate
	Finance     faker.FakeFinance
	Hacker      faker.FakeHacker
	Internet    faker.FakeInternet
	Lorem       faker.FakeLorem
	Name        faker.FakeName
	Number      faker.FakeNumber
	PhoneNumber faker.FakePhoneNumber
	Team        faker.FakeTeam
	Time        faker.FakeTime
}

// NewFaker creates a Faker intance using the original library implementation
func NewFaker() *Faker {
	return &Faker{
		Address:     faker.Address(),
		App:         faker.App(),
		Avatar:      faker.Avatar(),
		Bitcoin:     faker.Bitcoin(),
		Business:    faker.Business(),
		Code:        faker.Code(),
		Commerce:    faker.Commerce(),
		Company:     faker.Company(),
		Date:        faker.Date(),
		Finance:     faker.Finance(),
		Hacker:      faker.Hacker(),
		Internet:    faker.Internet(),
		Lorem:       faker.Lorem(),
		Name:        faker.Name(),
		Number:      faker.Number(),
		PhoneNumber: faker.PhoneNumber(),
		Team:        faker.Team(),
		Time:        faker.Time(),
	}
}
