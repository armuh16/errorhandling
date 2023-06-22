package main

func main() {
	if err := userJustShitTheirPants(); err != nil {
		changeDiapers()
	}
	if err := Bar(); err != nil {
		// handler
	}
}

func Foo() error {
	return nil
}

func Bar() error {
	if err := Foo(); err != nil {
		return err
	}
	return nil
}

func changeDiapers() (int, error) {
	return 10, nil
}

func userJustFarted() (int, error) {
	return 10, nil
}

func userJustShitTheirPants() error {
	return nil
}
