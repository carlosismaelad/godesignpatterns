package builder

import "carlosismaelad.com/godesignpatterns/builder/model"

type ClientBuilder struct {
	client model.Client
}

func NewBuilder() *ClientBuilder {
    return &ClientBuilder{}
}

func (b *ClientBuilder) Name(name string) *ClientBuilder {
	b.client.Name = name
	return b
}

func (b *ClientBuilder) Email(email string) *ClientBuilder {
	b.client.Email = email
	return b
}

func (b *ClientBuilder) Phone(phone string) *ClientBuilder {
	b.client.Phone = phone
	return b
}

func (b *ClientBuilder) Address(address string) *ClientBuilder {
	b.client.Address = address
	return b
}

func (b *ClientBuilder) Age(age int) *ClientBuilder {
	b.client.Age = age
	return b
}

func (b *ClientBuilder) IsActive(isActive bool) *ClientBuilder {
	b.client.IsActive = isActive
	return b
}

func (b *ClientBuilder) Build() model.Client {
	return b.client
}