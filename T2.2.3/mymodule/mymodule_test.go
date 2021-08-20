package mymodule

import (
	"strings"
	"testing"
)

var TestString string =`drwx------ 5 user user 12288 янв 15 14:59 Downloads
drwxr-xr-x 6 user user 4096 дек 6 14:29 Android
drwxr-xr-x 7 user user 4096 июн 10 2015 Sources
drwxr-xr-x 7 user user 4096 окт 31 15:08 VirtualBox
drwxr-xr-x 7 user user 4096 янв 13 11:42 Lightworks
drwxr-xr-x 8 user user 12288 янв 11 12:33 Pictures
drwxr-xr-x 8 user user 12288 янв 11 12:33 Pictures
drwxr-xr-x 8 user user 12288 янв 11 12:33 Pictures`

var TestCase1string string = `drwx------ 5 user user 12288 янв 15 14:59 Downloads
drwxr-xr-x 6 user user 4096 дек 6 14:29 Android
drwxr-xr-x 7 user user 4096 июн 10 2015 Sources
drwxr-xr-x 7 user user 4096 окт 31 15:08 VirtualBox
drwxr-xr-x 7 user user 4096 янв 13 11:42 Lightworks
drwxr-xr-x 8 user user 12288 янв 11 12:33 Pictures
drwxr-xr-x 8 user user 12288 янв 11 12:33 Pictures
drwxr-xr-x 8 user user 12288 янв 11 12:33 Pictures`

var TestCase2string string=`drwxr-xr-x 6 user user 4096 дек 6 14:29 Android
drwx------ 5 user user 12288 янв 15 14:59 Downloads
drwxr-xr-x 7 user user 4096 янв 13 11:42 Lightworks
drwxr-xr-x 8 user user 12288 янв 11 12:33 Pictures
drwxr-xr-x 8 user user 12288 янв 11 12:33 Pictures
drwxr-xr-x 8 user user 12288 янв 11 12:33 Pictures
drwxr-xr-x 7 user user 4096 июн 10 2015 Sources
drwxr-xr-x 7 user user 4096 окт 31 15:08 VirtualBox`

var TestCase3string string=`drwxr-xr-x 6 user user 4096 дек 6 14:29 Android
drwxr-xr-x 7 user user 4096 июн 10 2015 Sources
drwxr-xr-x 8 user user 12288 янв 11 12:33 Pictures
drwxr-xr-x 8 user user 12288 янв 11 12:33 Pictures
drwxr-xr-x 8 user user 12288 янв 11 12:33 Pictures
drwxr-xr-x 7 user user 4096 янв 13 11:42 Lightworks
drwx------ 5 user user 12288 янв 15 14:59 Downloads
drwxr-xr-x 7 user user 4096 окт 31 15:08 VirtualBox`

var TestCase4string string=`drwxr-xr-x 7 user user 4096 окт 31 15:08 VirtualBox
drwx------ 5 user user 12288 янв 15 14:59 Downloads
drwxr-xr-x 7 user user 4096 янв 13 11:42 Lightworks
drwxr-xr-x 8 user user 12288 янв 11 12:33 Pictures
drwxr-xr-x 8 user user 12288 янв 11 12:33 Pictures
drwxr-xr-x 8 user user 12288 янв 11 12:33 Pictures
drwxr-xr-x 7 user user 4096 июн 10 2015 Sources
drwxr-xr-x 6 user user 4096 дек 6 14:29 Android`

var TestCase5string string=`drwxr-xr-x 7 user user 4096 окт 31 15:08 VirtualBox
drwx------ 5 user user 12288 янв 15 14:59 Downloads
drwxr-xr-x 7 user user 4096 янв 13 11:42 Lightworks
drwxr-xr-x 8 user user 12288 янв 11 12:33 Pictures
drwxr-xr-x 7 user user 4096 июн 10 2015 Sources
drwxr-xr-x 6 user user 4096 дек 6 14:29 Android`

func TestCase1(t *testing.T) {
	strings:= strings.Split( TestString,"\n")
	flags :=My_flags{N:false,K: -1,R:false,U:false}
	Sort(&strings,flags)
	res:=""
	for _,s :=range strings{
		res+=s+"\n"
	}
	res=res[:len(res)-1]
	if res!=TestCase1string{
		t.Error("Something wrong")
	}
}

func TestCase2(t *testing.T) {
	strings:= strings.Split( TestString,"\n")
	flags :=My_flags{N:false,K: 8,R:false,U:false}
	Sort(&strings,flags)
	res:=""
	for _,s :=range strings{
		res+=s+"\n"
	}
	res=res[:len(res)-1]
	if res!=TestCase2string{
		t.Error("Something wrong")
	}
}

func TestCase3(t *testing.T) {
	strings:= strings.Split( TestString,"\n")
	flags :=My_flags{N:true,K: 6,R:false,U:false}
	Sort(&strings,flags)
	res:=""
	for _,s :=range strings{
		res+=s+"\n"
	}
	res=res[:len(res)-1]
	if res!=TestCase3string{
		t.Error("Something wrong")
	}
}
func TestCase4(t *testing.T) {
	strings:= strings.Split( TestString,"\n")
	flags :=My_flags{N:true,K: 6,R:true,U:false}
	Sort(&strings,flags)
	res:=""
	for _,s :=range strings{
		res+=s+"\n"
	}
	res=res[:len(res)-1]
	if res!=TestCase4string{
		t.Error("Something wrong")
	}
}

func TestCase5(t *testing.T) {
	strings:= strings.Split( TestString,"\n")
	flags :=My_flags{N:true,K: 6,R:true,U:true}
	Sort(&strings,flags)
	res:=""
	for _,s :=range strings{
		res+=s+"\n"
	}
	res=res[:len(res)-1]
	if res!=TestCase5string{
		t.Error("Something wrong")
	}
}