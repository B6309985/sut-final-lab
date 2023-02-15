package entity

import (
	"testing"

	"github.com/asaskevich/govalidator"
	. "github.com/onsi/gomega"
	"gorm.io/gorm"
	//"github.com/B6309985/sut-final-lab/entity"
)

type Customer struct {
	gorm.Model
	Name       string `valid:"required~name not blank"` // ต้องไม่เป็นค่าว่าง
	Email      string
	CustomerID string // รหัสลูกค้าขึ8นต้นด้วย L หรือ M หรือ H แล้วตามด้วยตัวเลขจํานวน 7 ตัว
}

func TestNameCantNotBeBlank(t *testing.T) {

	g := NewGomegaWithT(t)
	t.Run("check Name not blank", func(t *testing.T) {
		customer := Customer{
			Name:       "",
			Email:      "g@gmail.com",
			CustomerID: "H1236598",
		}

		ok, err := govalidator.ValidateStruct(customer)

		g.Expect(ok).ToNot(BeTrue())
		g.Expect(err).ToNot(BeNil())
		g.Expect(err.Error()).To(Equal("name not blank"))

	})

}
