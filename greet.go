package main

type Greet interface {
	Hello() string
	HowAreYou() string
}

type Chinese struct{}

func (c *Chinese) Hello() string {
	return "安安"
}

func (c *Chinese) HowAreYou() string {
	return "你好嗎"
}

type English struct{}

func (c *English) Hello() string {
	return "hello"
}

func (c *English) HowAreYou() string {
	return "how are you"
}
