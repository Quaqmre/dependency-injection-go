package main

import (
	"fmt"
	"reflect"
)

type depender struct {
	depends map[string]interface{}
}

func newDepender()*depender{
	return &depender{
		depends: make(map[string]interface{}),
	}
}

func (d *depender)Adddepender(f func()interface{}){
	a:=f()
	if t := reflect.TypeOf(a); t.Kind() == reflect.Ptr {
		d.depends[t.Elem().Name()]=a
	} else {
		panic(t.Name()+ " dependency error")
	}
}

func (d *depender) GetDepend(i interface{}) interface{}{
	if t := reflect.TypeOf(i); t.Kind() == reflect.Ptr {
		dep:=d.depends[t.Elem().Name()]
		if dep!=nil{
			return dep
		}else{
			panic("not exist depend")
		}
	} else {
		panic(t.Name()+ "dependency error")
	}
}
type s struct{
	name string
}

func addS()interface{}{
	return &s{"ilk dependency"}
}

func main(){
	di:=newDepender()
	di.Adddepender(addS)

	s:=di.GetDepend(&s{})
	fmt.Println(s)

}

