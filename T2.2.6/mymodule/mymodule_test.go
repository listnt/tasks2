package mymodule

import "testing"

var TestCase1Flag =My_flags{1,":",true}

var TestString=`Winter: white: snow: frost
Some string
Spring: green: grass: warm
Summer: colorful: blossom: hot
Error string
Autumn: yellow: leaves: cool
Empty string`

var TestCase1Res=`white
green
colorful
yellow`

func TestCase1(t *testing.T){
	res:=Cut(TestString,TestCase1Flag)
	if res!=TestCase1Res{
		t.Error("expected",TestCase1Res,"got",res)
	}
}