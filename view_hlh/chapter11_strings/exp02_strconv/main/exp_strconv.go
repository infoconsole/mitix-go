package main

import (
	"fmt"
	"strconv"
)

//todo  还没有看这个内容
func main() {
	fmt.Println("-----------------this is test expAppendBool-----------------")
	expAppendBool()
	fmt.Println("-----------------this is test expAppendBool-----------------")

	fmt.Println("-----------------this is test expAppendFloat-----------------")
	expAppendFloat()
	fmt.Println("-----------------this is test expAppendFloat-----------------")

	fmt.Println("-----------------this is test expAppendInt-----------------")
	expAppendInt()
	fmt.Println("-----------------this is test expAppendInt-----------------")

	fmt.Println("-----------------this is test expAppendQuote-----------------")
	expAppendQuote()
	fmt.Println("-----------------this is test expAppendQuote-----------------")

	fmt.Println("-----------------this is test expAppendQuoteRune-----------------")
	expAppendQuoteRune()
	fmt.Println("-----------------this is test expAppendQuoteRune-----------------")

	fmt.Println("-----------------this is test expAppendQuoteRuneToASCII-----------------")
	expAppendQuoteRuneToASCII()
	fmt.Println("-----------------this is test expAppendQuoteRuneToASCII-----------------")

	fmt.Println("-----------------this is test expAppendQuoteToASCII-----------------")
	expAppendQuoteToASCII()
	fmt.Println("-----------------this is test expAppendQuoteToASCII-----------------")

	fmt.Println("-----------------this is test expAppendUint-----------------")
	expAppendUint()
	fmt.Println("-----------------this is test expAppendUint-----------------")

	fmt.Println("-----------------this is test expAtoi-----------------")
	expAtoi()
	fmt.Println("-----------------this is test expAtoi-----------------")

	fmt.Println("-----------------this is test expCanBackquote-----------------")
	expCanBackquote()
	fmt.Println("-----------------this is test expCanBackquote-----------------")

	fmt.Println("-----------------this is test expFormatBool-----------------")
	expFormatBool()
	fmt.Println("-----------------this is test expFormatBool-----------------")

	fmt.Println("-----------------this is test expFormatFloat-----------------")
	expFormatFloat()
	fmt.Println("-----------------this is test expFormatFloat-----------------")

	fmt.Println("-----------------this is test expFormatInt-----------------")
	expFormatInt()
	fmt.Println("-----------------this is test expFormatInt-----------------")

}

func expAppendBool() {
	var dst []byte
	dst = strconv.AppendBool(dst, true)
	dst = strconv.AppendBool(dst, true)
	fmt.Println(dst)
}

func expAppendFloat() {

}

func expAppendInt() {

}

func expAppendQuote() {

}

func expAppendQuoteRune() {

}

func expAppendQuoteRuneToASCII() {

}
func expAppendQuoteToASCII() {

}

func expAppendUint() {

}

func expCanBackquote() {

}

func expFormatBool() {

}
func expFormatFloat() {

}
func expFormatInt() {

}

func expFormatUint() {

}

func expIsPrint() {

}

func expItoa() {

}

func expAtoi() {
	var stri = "666"
	num, err := strconv.Atoi(stri)
	if err != nil {
		fmt.Println("err is %s", err)
	}
	fmt.Printf("num is %d\n", num)
}

func expParseBool() {

}

func expParseFloat() {

}

func expParseInt() {

}

func expParseUint() {

}

func expQuote() {

}

func expQuoteRune() {

}

func expQuoteRuneToASCII() {

}

func expQuoteToASCII() {

}

func expUnquote() {

}

func expUnquoteChar() {

}
