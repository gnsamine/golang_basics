package deferstatement

import "fmt"

func A() {
	fmt.Println("a worked")
}

func C() {
	fmt.Println("c worked")
}

func B() {
	defer A()
	defer C()
	fmt.Println("b worked")
// fonksiyon yukardan aşağı çalışıyor. deferlerfe geldiğinde ilk giren son çıkar prensibi uygulanır
//the function works top-down. applying the first-in, last-out principle on defers, it works
// defer works even if b fonctions has error in its code.
}


