package main

import (
	"fmt"
	"strings"
	"unicode"
)

func main() {
	fmt.Println("-----------------this is test Contains-----------------")
	expContains()
	fmt.Println("-----------------this is test Contains-----------------")

	fmt.Println("-----------------this is test ContainsAny-----------------")
	expContainsAny()
	fmt.Println("-----------------this is test ContainsAny-----------------")

	fmt.Println("-----------------this is test ContainsRune-----------------")
	expContainsRune()
	fmt.Println("-----------------this is test ContainsRune-----------------")

	fmt.Println("-----------------this is test Count-----------------")
	expCount()
	fmt.Println("-----------------this is test Count-----------------")

	fmt.Println("-----------------this is test EqualFold-----------------")
	expEqualFold()
	fmt.Println("-----------------this is test EqualFold-----------------")

	fmt.Println("-----------------this is test Fields-----------------")
	expFields()
	fmt.Println("-----------------this is test Fields-----------------")

	fmt.Println("-----------------this is test FieldsFunc-----------------")
	expFieldsFunc()
	fmt.Println("-----------------this is test FieldsFunc-----------------")

	fmt.Println("-----------------this is test HasPrefix-----------------")
	expHasPrefix()
	fmt.Println("-----------------this is test HasPrefix-----------------")

	fmt.Println("-----------------this is test HasPrefix-----------------")
	expHasPrefix()
	fmt.Println("-----------------this is test HasPrefix-----------------")

	fmt.Println("-----------------this is test HasSuffix-----------------")
	expHasSuffix()
	fmt.Println("-----------------this is test HasSuffix-----------------")

	fmt.Println("-----------------this is test Index-----------------")
	expIndex()
	fmt.Println("-----------------this is test Index-----------------")

	fmt.Println("-----------------this is test expLastIndex-----------------")
	expLastIndex()
	fmt.Println("-----------------this is test expLastIndex-----------------")

	fmt.Println("-----------------this is test IndexAny-----------------")
	expIndexAny()
	fmt.Println("-----------------this is test IndexAny-----------------")

	fmt.Println("-----------------this is test expLastIndexAny-----------------")
	expLastIndexAny()
	fmt.Println("-----------------this is test expLastIndexAny-----------------")

	fmt.Println("-----------------this is test IndexByte-----------------")
	expIndexByte()
	fmt.Println("-----------------this is test IndexByte-----------------")

	fmt.Println("-----------------this is test expIndexFunc-----------------")
	expIndexFunc()
	fmt.Println("-----------------this is test expIndexFunc-----------------")

	fmt.Println("-----------------this is test expIndexRune-----------------")
	expIndexRune()
	fmt.Println("-----------------this is test expIndexRune-----------------")

	fmt.Println("-----------------this is test expJoin-----------------")
	expJoin()
	fmt.Println("-----------------this is test expJoin-----------------")

	fmt.Println("-----------------this is test expRepeat-----------------")
	expRepeat()
	fmt.Println("-----------------this is test expRepeat-----------------")

	fmt.Println("-----------------this is test expReplace-----------------")
	expReplace()
	fmt.Println("-----------------this is test expReplace-----------------")

	fmt.Println("-----------------this is test expSplit-----------------")
	expSplit()
	fmt.Println("-----------------this is test expSplit-----------------")

	fmt.Println("-----------------this is test expSplitN-----------------")
	expSplitN()
	fmt.Println("-----------------this is test expSplitN-----------------")

	fmt.Println("-----------------this is test expSplitAfter-----------------")
	expSplitAfter()
	fmt.Println("-----------------this is test expSplitAfter-----------------")

	fmt.Println("-----------------this is test expSplitAfterN-----------------")
	expSplitAfterN()
	fmt.Println("-----------------this is test expSplitAfterN-----------------")

	fmt.Println("-----------------this is test expTitle-----------------")
	expTitle()
	fmt.Println("-----------------this is test expTitle-----------------")

	fmt.Println("-----------------this is test expToLower-----------------")
	expToLower()
	fmt.Println("-----------------this is test expToLower-----------------")

	fmt.Println("-----------------this is test expToTitle-----------------")
	expToTitle()
	fmt.Println("-----------------this is test expToTitle-----------------")

	fmt.Println("-----------------this is test expToUpper-----------------")
	expToUpper()
	fmt.Println("-----------------this is test expToUpper-----------------")

	fmt.Println("-----------------this is test expTrim-----------------")
	expTrim()
	fmt.Println("-----------------this is test expTrim-----------------")

	fmt.Println("-----------------this is test expTrimSpace-----------------")
	expTrimSpace()
	fmt.Println("-----------------this is test expTrimSpace-----------------")

	fmt.Println("-----------------this is test expTrimPerfix-----------------")
	expTrimPerfix()
	fmt.Println("-----------------this is test expTrimPerfix-----------------")

	fmt.Println("-----------------this is test expTrimSuffix-----------------")
	expTrimSuffix()
	fmt.Println("-----------------this is test expTrimSuffix-----------------")

}

func expContains() {
	fmt.Println(strings.Contains("seafood", "foo"))
	fmt.Println(strings.Contains("seafood", "bar"))
	fmt.Println(strings.Contains("seafood", ""))
	fmt.Println(strings.Contains("", ""))
}

func expContainsAny() {
	fmt.Println(strings.ContainsAny("team", "i"))
	fmt.Println(strings.ContainsAny("failure", "u & i"))
	fmt.Println(strings.ContainsAny("foo", ""))
	fmt.Println(strings.ContainsAny("", ""))
}

//现在还不知道怎么搞
func expContainsRune() {
	fmt.Println("this is not lean")
}

func expCount() {
	fmt.Println(strings.Count("cheese", "e"))
	fmt.Println(strings.Count("five", "")) // before & after each rune
}

//忽略大小写比较
func expEqualFold() {
	fmt.Println(strings.EqualFold("Go", "go"))
}

//字段切分
func expFields() {
	fmt.Printf("Fields are: %q\n", strings.Fields("  foo bar  baz   "))
}

//切分，如果不满足函数则忽略
func expFieldsFunc() {
	f := func(c rune) bool {
		return !unicode.IsLetter(c) && !unicode.IsNumber(c)
	}
	fmt.Printf("Fields are: %q\n", strings.FieldsFunc("  foo1;bar2,baz3..123.", f))
}

func expHasPrefix() {
	fmt.Printf("the result of prefix is %v\n", strings.HasPrefix("https://www.mitix.cn", "https"))
}

func expHasSuffix() {
	fmt.Printf("the result of suffix is %v\n", strings.HasSuffix("https://www.mitix.cn", ".cn"))

}

//返回s字符串被substr配置的位置
func expIndex() {
	fmt.Println(strings.Index("chicken", "ken"))
	fmt.Println(strings.Index("chicken", "dmr"))
}

//返回s字符串substr最后的一个匹配的位置
func expLastIndex() {
	fmt.Println(strings.LastIndex("go gophergo", "go"))
	fmt.Println(strings.LastIndex("go gopher", "rodent"))
}

//返回s字符串中第一个被chars匹配的字符的位置
func expIndexAny() {
	fmt.Println(strings.IndexAny("chicken", "aeihun"))
	fmt.Println(strings.IndexAny("crwth", "aeiouy"))
}

func expLastIndexAny() {
	fmt.Println(strings.LastIndexAny("chicken", "aeihun"))
	fmt.Println(strings.LastIndexAny("crwth", "aeiouy"))
}

//返回s字符串中第一个被byte匹配的位置
func expIndexByte() {
	fmt.Println(strings.IndexByte("chicken", 'h'))
}

//返回满足函数的index
func expIndexFunc() {
	f := func(c rune) bool {
		return unicode.Is(unicode.Han, c)
	}
	fmt.Println(strings.IndexFunc("Hello, 世界", f))
	fmt.Println(strings.IndexFunc("Hello, world", f))
	expLastIndexFunc()
}

//参照IndexFunc
func expLastIndexFunc() {

}

//返回满足rune的index   还没看
func expIndexRune() {
	fmt.Println(strings.IndexRune("chicken", 'k'))
	fmt.Println(strings.IndexRune("chicken", 'd'))
}

//字符串链接
func expJoin() {
	s := []string{"foo", "bar", "baz"}
	fmt.Println(strings.Join(s, ", "))
}

//还不知道怎么搞
//todo  再学习
func expMap() {
	rot13 := func(r rune) rune {
		switch {
		case r >= 'A' && r <= 'Z':
			return 'A' + (r-'A'+13)%26
		case r >= 'a' && r <= 'z':
			return 'a' + (r-'a'+13)%26
		}
		return r
	}
	fmt.Println(strings.Map(rot13, "'Twas brillig and the slithy gopher..."))
}

//返回s字符串的计数副本拼接
func expRepeat() {
	fmt.Println("ba" + strings.Repeat("na", 2))
}

//func Replace(s, old, new string, n int) string
//替换字符串  old被替换字符串，如果为空则替换吧开头和后面每个字符之间，知道n个字符
//n替换次数，如果n<0则次数不限制
func expReplace() {
	fmt.Println(strings.Replace("oink oink oink", "", "ky", 2))
	fmt.Println(strings.Replace("oink oink oink", "k", "ky", 2))
	fmt.Println(strings.Replace("oink oink oink", "oink", "moo", -1))
}

//func Split(s, sep string) []string
//字符串切分 返回字符串数组
func expSplit() {
	fmt.Printf("%q\n", strings.Split("a,b,c", ","))
	fmt.Printf("%q\n", strings.Split("a man a plan a canal panama ", "a "))
	fmt.Printf("%q\n", strings.Split(" xyz ", ""))
	fmt.Printf("%q\n", strings.Split("xyz", ""))
	fmt.Printf("%q\n", strings.Split("", "Bernardo O'Higgins"))
}

//func Split(s, sep string) []string
//字符串切分 返回字符串数组
func expSplitN() {
	fmt.Printf("%q\n", strings.SplitN("a,b,c", ",", 2))
	fmt.Printf("%q\n", strings.SplitN("a man a plan a canal panama ", "a ", 3))
	fmt.Printf("%q\n", strings.SplitN(" xyz ", "", 2))
	fmt.Printf("%q\n", strings.SplitN("", "Bernardo O'Higgins", 2))
}

//func SplitAfter(s, sep string) []string
//切分字符串  保留分隔符在前面的一个数组上
func expSplitAfter() {
	fmt.Printf("%q\n", strings.SplitAfter("a,b,c", ","))
	fmt.Printf("%q\n", strings.SplitAfter("a man a plan a canal panama ", "a "))

}

//func SplitAfterN(s, sep string, n int) []string
//切分字符串  保留分隔符在前面的一个数组上
func expSplitAfterN() {
	fmt.Printf("%q\n", strings.SplitAfterN("a,b,c", ",", 2))
	fmt.Printf("%q\n", strings.SplitAfterN("a man a plan a canal panama ", "a ", 3))

}

//func Title(s string) string
//返回每个单词首字母大写的字符串
func expTitle() {
	fmt.Println(strings.Title("her royal highness"))
}

//func ToLower(s string) string
//返回小写字符串
func expToLower() {
	fmt.Println(strings.ToLower("Gopher"))
}

//func ToLowerSpecial(_case unicode.SpecialCase, s string) string
//todo 后续学习
func expToLowerSpecial() {

}

//所有的Unicode字符都映射到标题大小写
func expToTitle() {
	fmt.Println(strings.ToTitle("loud noises"))
	fmt.Println(strings.ToTitle("хлеб"))
}

//func ToTitleSpecial(_case unicode.SpecialCase, s string) string
//todo 后续学习
func expToTitleSpecial() {

}

//其中所有Unicode字母都映射到它们的大写。
//func ToUpper(s string) string
func expToUpper() {
	fmt.Println(strings.ToUpper("Gopher"))
}

//func ToUpperSpecial(_case unicode.SpecialCase, s string) string
//todo 后续学习
func expToUpperSpecial() {

}

//func Trim(s string, cutset string) string
//Trim返回字符串s的一个片段，删除cutset中包含的所有前导和后导Unicode代码点。
func expTrim() {
	fmt.Printf("[%q]", strings.Trim(" !!! Achtung! Achtung! !!! ", "! "))
}

//func TrimLeft(s string, cutset string) string
////Trim返回字符串s的一个片段，删除cutset中包含的所有前导Unicode代码点。
func expTrimLeft() {
	fmt.Printf("[%q]", strings.TrimLeft(" !!! Achtung! Achtung! !!! ", "! "))
}

//func TrimRight(s string, cutset string) string
////Trim返回字符串s的一个片段，删除cutset中包含的所有后导Unicode代码点。
func expTrimRight() {
	fmt.Printf("[%q]", strings.TrimRight(" !!! Achtung! Achtung! !!! ", "! "))
}

//func TrimFunc(s string, f func(rune) bool) string
//todo 后续学习
func expTrimFunc() {

}

//func TrimLeftFunc(s string, f func(rune) bool) string
//todo 后续学习
func expTrimLeftFunc() {

}

//func TrimRightFunc(s string, f func(rune) bool) string
//todo 后续学习
func expTrimRightFunc() {

}

//func TrimSpace(s string) string
//去掉前后面的空白信息
func expTrimSpace() {
	fmt.Println(strings.TrimSpace(" \t\n a lone gopher \n\t\r\n"))
}

//func TrimPrefix(s, prefix string) string
//返回去掉前缀的字符串
func expTrimPerfix() {
	var s = "Goodbye,, world!"
	s = strings.TrimPrefix(s, "Goodbye,")
	s = strings.TrimPrefix(s, "Howdy,")
	fmt.Println("Hello" + s)
}

//func TrimSuffix(s, suffix string) string
//返回去掉后缀的字符串
func expTrimSuffix() {
	var s = "Hello, goodbye, etc!"
	s = strings.TrimSuffix(s, "goodbye, etc!")
	s = strings.TrimSuffix(s, "planet")
	fmt.Println(s, "world!")
}
