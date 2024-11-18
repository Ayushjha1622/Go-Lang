package main
import "fmt"

type User struct{
	Name string
	Email string
	Status bool
	Age int
	
}

func(u User) GetStatus(){
	fmt.Println("Is user active: ",u.Status)
}
func main ()  {
	Ayush := User{"Ayush","Ayuhs@gmail.com",true,21}
	Ayush.GetStatus()
}


func (u User) NewMail(){
	u.Email = "test@go.dev"
	fmt.Println("Email of this user is: ",u.Email)
	
	
}