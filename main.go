package main

import "fmt"

// func removeAnElement[T comparable](slice []T, val T) []T { //Met this problem .\main.go:48:39: Observer does not implement comparable
//
//    for i, str := range slice {
//       if str == val {
//          slice = append(slice[:i], slice[i+1:]...)
//          break
//       }
//    }
//    return slice
// }
func removeAnElement(slice []Observer, val Observer) []Observer {
	for i, str := range slice {
		if str == val {
			slice = append(slice[:i], slice[i+1:]...)
			break
		}
	}
	return slice
}

func removeAnElementStringType(slice []string, val string) []string { // could not create one func with generic type
	for i, str := range slice {
		if str == val {
			slice = append(slice[:i], slice[i+1:]...)
			break
		}
	}
	return slice
}

type Observable interface {
	subscribe(Observable)
	unsubscribe(Observable)
	sendAll()
}
type Observer interface {
	handleEvent([]string)
}

func subscribe(jobSite *JobSite, observer Observer) {
	jobSite.subscribers = append(jobSite.subscribers, observer)
}

func unsubscribe(jobSite *JobSite, observer Observer) {
	jobSite.subscribers = removeAnElement(jobSite.subscribers, observer)
}

type JobSite struct {
	subscribers []Observer
	vacancies   []string
}

type Person struct {
	name string
}

func (p Person) handleEvent(vacancies []string) {
	fmt.Print("Hello Mr/Mrs ")
	fmt.Println(p.name)
	fmt.Println("Updated job list:")
	for _, v := range vacancies {
		fmt.Println(v)
	}
}

func sendAll(jobSite *JobSite) {
	for _, subscriber := range jobSite.subscribers {
		subscriber.handleEvent(jobSite.vacancies)
	}
}
func addVacancy(jobSite *JobSite, vacancy string) {
	jobSite.vacancies = append(jobSite.vacancies, vacancy)
	sendAll(jobSite)
}
func removeVacancy(jobSite *JobSite, vacancy string) {
	jobSite.vacancies = removeAnElementStringType(jobSite.vacancies, vacancy)
	sendAll(jobSite)
}

func main() {
	fmt.Println()

	var jobSite JobSite
	//Person 1
	jason := Person{name: "jason"}
	subscribe(&jobSite, jason)
	//Job 1
	addVacancy(&jobSite, "FrontEnd")

	//Person 2
	madam := Person{name: "madam"}
	subscribe(&jobSite, madam)
	//Job 2
	fmt.Println("================================")
	addVacancy(&jobSite, "BackEnd")

	//Unsubscribe and add Vacancy
	fmt.Println("================================")
	unsubscribe(&jobSite, jason)
	addVacancy(&jobSite, "Go developer")

	//Remove Vacancy
	fmt.Println("================================")
	removeVacancy(&jobSite, "FrontEnd")
}
