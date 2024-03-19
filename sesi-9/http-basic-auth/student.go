package main

var student = []*Student{}

type Student struct {
	Id    string
	Name  string
	Grade int
}

func GetStudents() []*Student {
	return student
}

func SelectStudent(id string) *Student {
	for _, each := range student {
		if each.Id == id {
			return each
		}
	}
	return nil
}

func init() {
	student = append(student, &Student{Id: "s001", Name: "wick", Grade: 2})
	student = append(student, &Student{Id: "s002", Name: "ethan", Grade: 2})
	student = append(student, &Student{Id: "s003", Name: "bourne", Grade: 3})
	student = append(student, &Student{Id: "s004", Name: "bond", Grade: 4})
}