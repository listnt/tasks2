package mymodule

import "testing"

var TestCase1flags = MyFlags{1, 0, 0, false, false, false, false, false}
var TestCase2flags = MyFlags{0, 0, 2, false, true, false, true, false}
var TestCase3flags = MyFlags{0, 0, 2, false, true, true, true, false}
var TestCase4flags = MyFlags{0, 0, 2, true, true, true, true, false}

var TestString = `drwx------ 5 user user 12288 янв 15 14:59 Downloads
drwxr-xr-x 6 user user 4096 дек 6 14:29 Android
drwxr-xr-x 8 user user 12288 янв 11 12:33 Pic+tures1Pic123
drwxr-xr-x 7 user user 4096 июн 10 2015 Sources
drwxr-xr-x 8 user user 12288 янв 11 12:33 Pictures
drwxr-xr-x 7 user user 4096 окт 31 15:08 VirtualBoxpic
drwxr-xr-x 7 user user 4096 янв 13 11:42 Lightworks
drwxr-xr-x 8 user user 12288 янв 11 12:33 Pictures1Pic`

var requestString = `pic+`

var TestString1 = `drwxr-xr-x 8 user user 12288 янв 11 12:33 Pictures
drwxr-xr-x 7 user user 4096 окт 31 15:08 VirtualBox[31mpic[0m`

var TestString2 = `drwx------ 5 user user 12288 янв 15 14:59 Downloads
drwxr-xr-x 6 user user 4096 дек 6 14:29 Android
drwxr-xr-x 8 user user 12288 янв 11 12:33 [31mPic+[0mtures1Pic123
drwxr-xr-x 7 user user 4096 июн 10 2015 Sources
drwxr-xr-x 8 user user 12288 янв 11 12:33 Pictures`

var TestString3 = `drwx------ 5 user user 12288 янв 15 14:59 Downloads
drwxr-xr-x 6 user user 4096 дек 6 14:29 Android
drwxr-xr-x 8 user user 12288 янв 11 12:33 [31mPic+[0mtures1Pic123
drwxr-xr-x 7 user user 4096 июн 10 2015 Sources
drwxr-xr-x 8 user user 12288 янв 11 12:33 Pictures
drwxr-xr-x 7 user user 4096 окт 31 15:08 VirtualBoxpic
drwxr-xr-x 7 user user 4096 янв 13 11:42 Lightworks
drwxr-xr-x 8 user user 12288 янв 11 12:33 Pictures1Pic`

var TestString4 = `7`

func TestCase1(t *testing.T) {
	grep := NewGrep()
	grep.SetFlags(TestCase1flags)
	res := grep.Grep(TestString, requestString)
	if res != TestString1 {
		t.Error("expected\n", TestString1, "\n", "got\n", res)
	}
}

func TestCase2(t *testing.T) {
	grep := NewGrep()
	grep.SetFlags(TestCase2flags)
	res := grep.Grep(TestString, requestString)
	if res != TestString2 {
		t.Error("expected\n", TestString2, "\n", "got\n", res)
	}
}

func TestCase3(t *testing.T) {
	grep := NewGrep()
	grep.SetFlags(TestCase3flags)
	res := grep.Grep(TestString, requestString)
	if res != TestString3 {
		t.Error("expected\n", TestString3, "\n", "got\n", res)
	}
}

func TestCase4(t *testing.T) {
	grep := NewGrep()
	grep.SetFlags(TestCase4flags)
	res := grep.Grep(TestString, requestString)
	if res != TestString4 {
		t.Error("expected\n", TestString4, "\n", "got\n", res)
	}
}
