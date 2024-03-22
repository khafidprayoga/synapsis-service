package seed

import (
	"github.com/go-faker/faker/v4"
	"github.com/google/uuid"
	"github.com/khafidprayoga/synapsis-service/src/gen/synapsis/v1"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func NewUser() []*synapsisv1.User {
	return []*synapsisv1.User{
		{
			FullName: faker.FirstNameMale(),
			Email:    faker.Email(),
			Password: faker.Password(),
			Dob:      timestamppb.Now(),
		},
		{
			Id:       uuid.New().String(),
			FullName: faker.Name(),
			Email:    faker.Email(),
			Password: faker.Password(),
			Dob:      timestamppb.Now(),
		},
	}
}
