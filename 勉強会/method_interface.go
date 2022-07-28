// オンライン ストアを管理するためのパッケージを作成する
// カスタム パッケージを使用してオンライン ストアのアカウントを管理するプログラムを作成します。 課題には、次の 4 つの要素が含まれます。

// アカウント所有者の姓と名を含む、Account という名前のカスタム型を作成します。 型には、ChangeName の機能も含まれている必要があります。

// クレジット数を float64 型として格納するための変数を含み Account オブジェクトを埋め込む、Employee という名前の別のカスタム型を作成します。
// 型には AddCredits、RemoveCredits、CheckCredits の機能も含まれている必要があります。 Employee オブジェクトを使用してアカウント名を変更できることを示す必要があります。

// Employee 名が姓と名を含む形式で出力されるように、Account オブジェクトに Stringer メソッドを記述します。

// 最後に、作成したパッケージを使用するプログラムを作成して、この課題に示されているすべての機能をテストします。
// つまり、メイン プログラムでは名前の変更、名前の出力、クレジットの追加、クレジットの削除、残高の確認などを行う必要があります。

package main

import (
	"fmt"
)

type Account struct {
	FirstName string
	LastName string
}

type Employee struct {
	*Account
	Credits float64
}

func (e Employee) String() string {
	return fmt.Sprintf("Name: %s %s\nCredits: %.2f\n", e.FirstName, e.LastName, e.Credits)
}

func (a *Account) ChangeName(firstName string) {
	a.FirstName = firstName
}

func CreateEmployee(firstName, lastName string, credits float64) (*Employee, error) {
	return &Employee{Account{firstName, lastName}, credits}, nil
}

func (e *Employee) AddCredits(amount float64) {
	e.Credits += amount
	fmt.Printf("Your current account balance is %v\n", e.Credits)
}

func (e *Employee) RemoveCredits() {
	e.Credits = 0
	fmt.Printf("Your current account balance is %v\n", e.Credits)
}

func (e *Employee) CheckCredits() {
	fmt.Printf("Your current account balance is %v\n", e.Credits)
}

func main()  {
	account := &Account{FirstName: "Kawasaki", LastName: "Takudai"}
	employee := &Employee{Credits: 500.3, Account: account}
	employee.ChangeName("Bautista")
	employee.AddCredits(3.9)
	employee.RemoveCredits()
	employee.CheckCredits()
}