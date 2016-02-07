// Code generated by protoc-gen-gogo.
// source: dessert.proto
// DO NOT EDIT!

/*
Package graphql is a generated protocol buffer package.

It is generated from these files:
	dessert.proto

It has these top-level messages:
	Dessert
*/
package graphql

import testing "testing"
import math_rand "math/rand"
import time "time"
import proto "github.com/gogo/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "github.com/gogo/protobuf/gogoproto"
import _ "github.com/opsee/protobuf/gogogqlproto"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

func TestDessertGraphQL(t *testing.T) {
	popr := math_rand.New(math_rand.NewSource(time.Now().UnixNano()))
	_ = NewPopulatedDessert(popr, false)
	objdesc := "A delicious dessert dish on the menu"
	pdesc := DessertType.PrivateDescription
	if pdesc != objdesc {
		t.Fatalf("String want %v got %v", objdesc, pdesc)
	}
}

//These tests are generated by github.com/opsee/protobuf/plugin/graphql