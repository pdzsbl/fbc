package sw

import (
	"fmt"
	"strconv"
	"testing"
)

/*func TestWBObfECDSASignWithIndex(t *testing.T) {
	msg := []byte("hello world")

	priv , _ := keygen()

	generateTableWithIndex(80,30,priv,0)

	r,s := WBObfECDSASignWithIndex(msg,80,0);

	x := ecdsa.Verify(priv.Public().(*ecdsa.PublicKey),msg,r,s)
	fmt.Print(x)
}*/
var ecdsapriv,_ = keygen()
func TestWrite(t *testing.T){
	//priv := new(big.Int).SetUint64(uint64(500))
	tableIndex:=generateTableIfNewPriv(ecdsapriv.D,128,40)//如果是新的私钥，创建表，最终都返回表的index
	fmt.Println(strconv.Itoa(tableIndex))
}