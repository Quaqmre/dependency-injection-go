package main

import (
	"fmt"
	"reflect"
)

type depender struct {
	depends map[string]interface{}
}

func newDepender() *depender {
	return &depender{
		depends: make(map[string]interface{}),
	}
}

func (d *depender) Adddepender(f interface{}) {
	r := d.analizeDepends(f)
	var v reflect.Value
	var bo bool
	if r != nil {
		v = d.callAddDependWithArgs(r, f)
		bo=true
	}

	if t := reflect.TypeOf(f); t.Kind() == reflect.Ptr {
		d.depends[t.Elem().Name()] = v
	} else {
		if bo {
			d.depends[t.Name()] = v.Interface()
		}else{
			d.depends[t.Name()]=f
		}
	}
}

func (d *depender) GetDepend(i interface{}) interface{} {
	if t := reflect.TypeOf(i); t.Kind() == reflect.Ptr {
		dep := d.depends[t.Elem().Name()]
		if dep != nil {
			return dep
		} else {
			panic("not exist depend")
		}
	} else {
		dep := d.depends[t.Name()]
		if dep != nil {
			return dep
		} else {
			panic("not exist depend")
		}
	}
}

type s struct {
	name string
}

func (s s) addDepend(str string) interface{} {
	s.name = str
	return s
}

func (d *depender) analizeDepends(y interface{}) []interface{} {
	x := reflect.TypeOf(y)
	a := x.Kind()
	_ = a
	if a == reflect.Func {

		numIn := x.NumIn() //Count inbound parameters
		inputs := []interface{}{}
		for i := 0; i < numIn; i++ {

			inV := x.In(i)
			if inV.Kind() == reflect.Ptr {

				e, exist := d.depends[inV.Elem().Name()]

				if !exist {
					panic(inV.Elem().Name() + " doesnt exist right now,add it before")
				}
				inputs[i] = e
			} else {
				e, exist := d.depends[inV.Name()]
				if !exist {
					panic(inV.Elem().Name() + " doesnt exist right now,add it before")
				}
				inputs=append(inputs,e)
			}

		}
		if numIn == 0 {
			return nil
		} else {
			return inputs
		}
	}
	return nil

}

func (d *depender) callAddDependWithArgs(r []interface{}, f interface{}) reflect.Value {
	x := reflect.TypeOf(f)
	in := make([]reflect.Value, x.NumIn())
	for i, v := range r {
		in[i] = reflect.ValueOf(v)
	}
	if x.Kind() == reflect.Func {
		g := reflect.ValueOf(f)
		res := g.Call(in)

		if x.NumOut() != 1 {
			panic("return count must be 1")
		}

		return res[x.NumOut()-1]
	}
	panic("something goes wrong")

}

func main() {
	//di:=newDepender()
	//di.Adddepender(addS)

	//s:=di.GetDepend(&s{})
	//fmt.Println(s)
	di := newDepender()
	di.Adddepender("test String for conf or something")
	di.Adddepender(s{}.addDepend)
	myIn:=di.GetDepend(s{}.addDepend).(s)
	fmt.Println(myIn.name)


}
