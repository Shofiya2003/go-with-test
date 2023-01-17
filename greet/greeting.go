package main

const morningGreeting = "good morning "
const afternoonGreeting = "good afternoon "
const eveningGreeting = "good evening "

func Greet(name string, time int) string {
	greeting := morningGreeting
	if time > 12 && time < 17 {
		greeting = afternoonGreeting
	}
	if time >= 17 && time < 20 {
		greeting = eveningGreeting
	}
	return greeting + name
}
