package main

type author struct{
  name string
  branch string
  salary int
  }

func (a author) show() {

  fmt.Println("Author name:",a.name)
  fmt.Println("Branch name:",a.branch)
  fmt.Println("Salary name:",a.salary)

  }

func main(){
  result:=author{
    name:"Ranjan",
    branch:"CSE",
    salary:20000,
    }
  result.show()
  }
