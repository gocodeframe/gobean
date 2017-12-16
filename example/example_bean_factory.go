package example

import "gobean/bean"

var ExampleBeanFactory = bean.NewFactory()

func init() {
	ExampleBeanFactory.AddBean("person1", &Person{Phone: &ApplePhone{}})
	ExampleBeanFactory.AddBean("person2", &Person{Phone: &VivoPhone{}})
}
