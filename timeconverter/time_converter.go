package timeconverter

import "time"

/*
 * Crea una función que reciba días, horas, minutos y segundos (como enteros)
 * y retorne su resultado en milisegundos.
 */

const (
	_dayInMiliseconds    = 86400000
	_hourInMiliseconds   = 3600000
	_minuteInMiliseconds = 60000
	_secondInMiliseconds = 1000
)

func Execute(days, hour, minute, seconds int) int64 {
	daysM := daysToMilliseconds(days)
	hoursM := hoursToMilliseconds(hour)
	minutesM := minutesToMilliseconds(minute)
	secondsM := secondsToMilliseconds(seconds)

	return (daysM + hoursM + minutesM + secondsM).Milliseconds()
}

func daysToMilliseconds(days int) time.Duration {
	return time.Duration(days) * _dayInMiliseconds * time.Millisecond
}

func hoursToMilliseconds(hour int) time.Duration {
	return time.Duration(hour) * _hourInMiliseconds * time.Millisecond
}

func minutesToMilliseconds(minute int) time.Duration {
	return time.Duration(minute) * _minuteInMiliseconds * time.Millisecond
}

func secondsToMilliseconds(second int) time.Duration {
	return time.Duration(second) * _secondInMiliseconds * time.Millisecond
}
