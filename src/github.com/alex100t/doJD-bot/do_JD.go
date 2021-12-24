package main

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
	"time"
)

/*
	Response format:
	in GT.m: 66012  09-25-2021 Saturday
	in bot: 66012 09-25-2021 Saturday
*/

const (
	USAFormat    = "01-02-2006"
	EUROFormat   = "02-01-2006"
	MonthFormat  = "02-Jan-2006"
	ShortFormat  = "2006-01-02"
	MaxJD        = 500000
	SecondsInDay = 24*60*60 - 1
)

func JD2Time(jdDate int) (t time.Time, err error) {

	dayzero, _ := time.Parse(ShortFormat, "1840-12-31")
	trgdate := dayzero.AddDate(0, 0, jdDate)

	return trgdate, nil
}

func Date2Time(strDate string) (t time.Time, err error) {

	t, err = time.Parse(USAFormat, strDate)
	if err != nil {
		t, err = time.Parse(EUROFormat, strDate)
	}
	if err != nil {
		t, err = time.Parse(ShortFormat, strDate)
	}
	if err != nil {
		t, err = time.Parse(MonthFormat, strings.Title(strings.ToLower(strDate)))
	}
	if err != nil {
		err = errors.New("invalid date format")
		return time.Time{}, err
	}

	return t, nil
}

func cvtJD(jdDate int) (res string, err error) {

	if jdDate > MaxJD {
		err = errors.New("too big number")
		return "", err
	}

	t, err := JD2Time(jdDate)
	if err != nil {
		return "", err
	}

	res, err = formatTime(t)
	if err != nil {
		return "", err
	}

	return res, nil
}

func cvtDATE(strDate string) (res string, err error) {

	t, err := Date2Time(strDate)
	if err != nil {
		return "", err
	}

	res, err = formatTime(t)
	if err != nil {
		return "", err
	}

	return res, nil
}

func cvtUNIX(uDate int) (res string, err error) {

	t := time.Unix(int64(uDate), 0)

	if err != nil {
		return "", err
	}

	res, err = formatTime(t)
	if err != nil {
		return "", err
	}

	return res, nil
}

func Time2JD(t time.Time) (jdRes int, err error) {

	dayzero, _ := time.Parse(ShortFormat, "1840-12-31")
	difference := t.Sub(dayzero)

	return int(difference.Hours() / 24), nil
}

func formatTime(t time.Time) (resp string, err error) {

	var b strings.Builder

	jdDate, err := Time2JD(t)
	if err != nil {
		return "", err
	}

	fmt.Fprintf(&b, "%d %s %s", jdDate, t.Format(USAFormat), t.Weekday())

	resp = b.String()
	return resp, nil
}

func CvtNumber2Time(inp string) (resp string) {

	var b strings.Builder

	nSec, err := strconv.Atoi(inp)
	if err != nil {
		resp = "invalid Number format"
		return resp
	}

	if nSec < 0 {
		resp = "the number of seconds cannot be less than zero"
		return resp
	}

	if nSec > SecondsInDay {

		fmt.Fprintf(&b, "%s (%d)", "the number of seconds is greater than the number of seconds in a day", SecondsInDay)
		resp = b.String()
		return resp
	}

	t := time.Date(0, 0, 0, 0, 0, nSec, 0, time.UTC)

	fmt.Fprintf(&b, "%s", t.Format("15:04:05"))

	resp = b.String()
	return resp

}

func CvtTime2Number(inp string) (resp string) {

	var b strings.Builder

	//const shortForm = "2006-Jan-02"
	//t, _ := time.Parse("15:04:05", "17:41:56")
	t, err := time.Parse("15:04:05", inp)
	if err != nil {
		resp = "invalid Time format"
		return resp
	}
	//fmt.Println(t)

	hh, mm, ss := t.Clock()

	nSec := hh*60*60 + mm*60 + ss

	if nSec > SecondsInDay {
		fmt.Fprintf(&b, "%s (%d)", "the number of seconds is greater than the number of seconds in a day", SecondsInDay)
		resp = b.String()
		return resp
	}

	if nSec < 0 {
		resp = "the number of seconds cannot be less than zero"
		return resp
	}

	fmt.Fprintf(&b, "%d", nSec)

	resp = b.String()
	return resp

}
