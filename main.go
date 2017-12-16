package main

import "gobean/example"

func main() {
	example.ExampleBeanFactory.GetBean("person1").(*example.Person).PhoneType()
	example.ExampleBeanFactory.GetBean("person2").(*example.Person).PhoneType()
}
