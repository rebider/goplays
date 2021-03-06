package data

import "gopkg.in/mgo.v2/bson"

//设置变量
//key             value
const (
	ENV1  = "regist_diamond"  //注册赠送钻石
	ENV2  = "regist_coin"     //注册赠送金币
	ENV3  = "regist_chip"     //注册赠送筹码
	ENV4  = "regist_card"     //注册赠送房卡
	ENV5  = "build"           //绑定赠送
	ENV6  = "first_pay_multi" //首充送n倍
	ENV7  = "first_pay_coin"  //首充送金币
	ENV8  = "relieve"         //救济金次数
	ENV9  = "prizedraw"       //转盘抽奖次数
	ENV10 = "bankrupt_coin"   //破产金额
	ENV11 = "relieve_coin"    //救济金额
	ENV12 = "robot_num"       //虚假人数
	ENV13 = "robot_allot1"    //机器人分配规则1
	ENV14 = "robot_allot2"    //机器人分配规则2
	ENV15 = "robot_bet"       //机器人下注AI
)

type Env struct {
	Key   string `bson:"_id" json:"key"`     //key
	Value int32  `bson:"value" json:"value"` //value
}

func GetEnvList() []Env {
	var list []Env
	ListByQ(Envs, nil, &list)
	return list
}

func (this *Env) DelEnv() bool {
	return Delete(Envs, this)
}

func (this *Env) SetEnv() bool {
	return Upsert(Envs, bson.M{"_id": this.Key}, this)
}
