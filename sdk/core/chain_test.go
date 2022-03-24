package sdk

import (
	"fmt"
	"reflect"
	"testing"
)

func TestA(t *testing.T) {
	// c := Chain{Name: "aa"}

	// c.Description()

	// b := EthChain{}
	// b.Description()

	// var d Protocol

	// t.Log(d.Name)

	// d = &Chain{}
	// d = &EthChain{}
	// d.Description()
}

func TestChain(t *testing.T) {
	sdk := GetSharedManager()
	c1 := sdk.ChainWithType(ChainBinance)
	c2 := sdk.ChainWithType(ChainChainx)
	c3 := sdk.ChainWithType("xxx")

	println(c1, c2, c3)

	println(c1.AsEth())

	println(c2.AsEth())

	x := reflect.TypeOf(100)
	fmt.Printf("name = %s, kind = %s\n", x.Name(), x.Kind())
}
