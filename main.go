package main

import (
	"fmt"
	"strconv"
	"strings"
	"time"
)

func main() {
	now := time.Now()
	var dia, mes, ano, hora, minutos, segundos string
	var dataString string = fmt.Sprintf("%v", now)
	dataLimitada := dataString[:19]
	var dataFormatada []string = strings.Split(dataLimitada, "-")
	var horaLista []string = strings.Split(dataFormatada[2], " ")
	dia = horaLista[0]
	mes = dataFormatada[1]
	ano = dataFormatada[0]
	hora = horaLista[1]
	fmt.Println(dia, "/", mes, "/", ano, " ", hora)
	fmt.Println(ano, "-", mes, "-", dia, "t", hora)
	var horaSep []string = strings.Split(hora, ":")
	hora = horaSep[0]
	minutos = horaSep[1]
	segundos = horaSep[2]
	var diaInt, mesInt, anoInt, horaInt, minutosInt, segInt, contadorAno, contadorBissexto, somaDiaMes int
	diaInt, _ = strconv.Atoi(dia)
	mesInt, _ = strconv.Atoi(mes)
	anoInt, _ = strconv.Atoi(ano)
	horaInt, _ = strconv.Atoi(hora)
	minutosInt, _ = strconv.Atoi(minutos)
	segInt, _ = strconv.Atoi(segundos)
	for anoInicio := 1970; anoInicio <= anoInt; anoInicio++ {
		if anoInicio%4 == 0 && anoInicio%100 != 0 || anoInicio%400 == 0 {
			contadorBissexto = contadorBissexto + 366
		} else if anoInicio == anoInt {
			somaDiaMes = ((mesInt - 1) * 2629743) + ((diaInt - 1) * 86400)
		} else {
			contadorAno = contadorAno + 365
		}
	}
	var dataEpoch int = ((contadorAno + contadorBissexto) * 86400) + (horaInt * 3600) + (minutosInt * 60) + segInt + somaDiaMes + (3 * 3600)
	fmt.Println(dataEpoch)
	fmt.Println(now.Unix())
}
