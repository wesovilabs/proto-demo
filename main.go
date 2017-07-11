package main

import (
	"github.com/wesovilabs/proto-demo/pb"
	"fmt"
	"encoding/json"
)

func main(){







	iban := &pb.Element_Attribute{Name:"iban", Type:pb.Element_Attribute_STRING, Value:"8904"}
	holder := &pb.Element_Attribute{Name:"holder", Type:pb.Element_Attribute_STRING, Value:"Iván Corrales Solera"}
	accountAttributes := []*pb.Element_Attribute{iban,holder}

	fnAtt := &pb.Element_Attribute{Name:"fn",Type:pb.Element_Attribute_STRING,Value:"John"}
	lnAtt := &pb.Element_Attribute{Name:"ln",Type:pb.Element_Attribute_STRING,Value:"Doe"}
	genderAtt := &pb.Element_Attribute{Name:"gender",Type:pb.Element_Attribute_BOOLEAN,Value:"false"}
	ageAtt := &pb.Element_Attribute{Name:"age",Type:pb.Element_Attribute_NUMBER,Value:"32"}
	accAtt := &pb.Element_Attribute{Name:"account", Type:pb.Element_Attribute_NODE, Children:accountAttributes}
	attributes := []*pb.Element_Attribute{fnAtt,lnAtt,genderAtt,ageAtt,accAtt}



	element := &pb.Element{Name:"person", Attributes:attributes}

	b, err := json.Marshal(element)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(b))

}
