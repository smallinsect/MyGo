// 角色道具

package TestMongoDB

import (
	// "fmt"

	"github.com/globalsign/mgo/bson"
	log "github.com/sirupsen/logrus"
	// "cmt1-server/data"
	// "cmt1-server/db"
	// "cmt1-server/game/cache"
	// "cmt1-server/game/errors"
	// "cmt1-server/utils"
)

type EntryI interface {
	Save() (err error)
	SaveDataStatusUpdate()  // 更新保存数据到数据库状态
	SaveUpdateStatus() bool // 是否需要保存到数据库
	Load(key interface{}) (err error)
}

type BaseM struct {
	saveStatus bool
}

func (bm *BaseM) UpdateData() {
	bm.saveStatus = true
}

func (bm *BaseM) SaveUpdateStatus() bool {
	return bm.saveStatus
}

func (bm *BaseM) SaveDataStatusUpdate() {
	bm.saveStatus = false
}

// 英雄集合
type Heros struct {
	BaseM
	Uid   bson.ObjectId   `bson:"_id" json:"_id"`
	Heros map[int64]*Hero `bson:"heros" json:"heros"`
}

// 英雄
type Hero struct {
	Goods   map[int64]*HeroGoods `bson:"goods"`
	SkillId int64                `bson:"skillId" json:"skillId"` // 英雄技能
}

// 英雄物品
type HeroGoods struct {
	Tid  int64 `bson:"tid" json:"tid"`
	Num  int64 `bson:"num" json:"num"`
	Star int64 `bson:"star" json:"star"`
}

func NewHeros() EntryI {
	return &Heros{Heros: make(map[int64]*Hero)}
}

// 保存信息
func (heros *Heros) Save() (err error) {
	if heros.Heros == nil || len(heros.Heros) == 0 {
		return
	}
	err = Upsert("heros", bson.M{"_id": heros.Uid}, heros)
	if err != nil {
		return
	}
	return
}

func (heros *Heros) Load(key interface{}) (err error) {
	switch key.(type) {
	case string:
		if key.(string) == "" {
			return nil
		}
		heros.Uid = bson.ObjectIdHex(key.(string))
	case bson.ObjectId:
		heros.Uid = key.(bson.ObjectId)
	default:
		log.WithFields(log.Fields{
			"method": "Heros_Load",
			"data":   key,
		}).Error("无效的key load")
		return nil
	}
	return heros.LoadData()
}

// 加载数据
func (heros *Heros) LoadData() (err error) {
	err = FindOne("heros", bson.M{"_id": heros.Uid}, nil, &heros)
	if err != nil {
		if err.Error() == "not found" {
			return nil
		}
		return
	}
	if heros.Heros == nil {
		heros.Heros = make(map[int64]*Hero, 0)
	}
	return
}

// // 更新数据
// func (heros *Heros) Update(val *Heros) {
// 	Cache.Put("heros", heros.Uid.Hex(), val)
// 	val.UpdateData()
// }

// // 添加英雄物品
// func (heros *Heros) AddHeroGoods(goodsT data.Good, num int64) (hero *Hero, uprop *UpdateProp) {
// 	defer heros.Update(heros)
// 	heroTid := int64(goodsT.Owner)
// 	goodstid := int64(goodsT.Id)
// 	uprop = NewUpdateProp()
// 	heroGoods := &HeroGoods{Tid: goodstid, Num: num, Star: int64(goodsT.Star)}
// 	uprop.HeroGoodsList = append(uprop.HeroGoodsList, heroGoods)

// 	if v, ok := heros.Heros[heroTid]; ok {
// 		hero = v
// 		if v1, ok := v.Goods[goodstid]; ok {
// 			v1.Num += num
// 			return
// 		}
// 		uprop.Kvs = append(uprop.Kvs, &BaseKV{Key: data.GlobalEnum.Star, Val: int64(goodsT.Star)})
// 		v.Goods[goodstid] = heroGoods
// 		return
// 	}

// 	hero = &Hero{Goods: make(map[int64]*HeroGoods, 0)}
// 	hero.Goods[goodstid] = heroGoods
// 	heros.Heros[heroTid] = hero
// 	uprop.Kvs = append(uprop.Kvs, &BaseKV{Key: data.GlobalEnum.Star, Val: int64(goodsT.Star)})
// 	return
// }

// // 获取英雄
// func (heros *Heros) GetHero(heroTid int64) (r *Hero, ok bool) {
// 	r, ok = heros.Heros[heroTid]
// 	return
// }

// // 获取英雄物品
// func (heros *Heros) GetHeroGoods(heroTid int64) (r map[int64]*HeroGoods, ok bool) {
// 	hero, ok := heros.GetHero(heroTid)
// 	if !ok {
// 		return
// 	}
// 	r = hero.Goods
// 	return
// }

// func (heros *Heros) StarNum() (star int64) {
// 	star = 0
// 	for _, item := range heros.Heros {
// 		for _, item1 := range item.Goods {
// 			star += item1.Star
// 		}
// 	}
// 	return
// }

// // 获取英雄列表
// func (heros *Heros) GetHeros() (r map[int64]*Hero) {
// 	r = heros.Heros
// 	return
// }

// // 根据物品tid获取到英雄物品信息
// func (heros *Heros) GetHeroGoodsByGoodsTid(id int64) (r *HeroGoods, ok bool) {
// 	goodTpl, ok := data.Get("Goods", int64(id))
// 	if !ok {
// 		return
// 	}
// 	goodT := goodTpl.(data.Good)
// 	mHgoods, ok := heros.GetHeroGoods(int64(goodT.Owner))
// 	if !ok {
// 		return
// 	}
// 	r, ok = mHgoods[id]
// 	return
// }

// // 扣除英雄物品数量
// func (heros *Heros) CostHeroGoods(id, num int64) (ok bool) {
// 	hgoods, ok := heros.GetHeroGoodsByGoodsTid(id)
// 	if !ok {
// 		return
// 	}
// 	defer heros.Update(heros)
// 	hgoods.Num -= num
// 	return
// }

// // 玩家英雄列表信息
// func (u *User) Heros() (r *Heros) {
// 	v1, _ := Cache.Get("heros", u.Uid.Hex())
// 	r = v1.(*Heros)
// 	return
// }

// // 检查英雄是否集满物品
// func (hero *Hero) CheckFullGoods(heroTid uint64) (ok bool) {
// 	list, ok := data.Filters("Goods", func(val interface{}) bool {
// 		val1 := val.(data.Good)
// 		if val1.Owner == heroTid {
// 			return true
// 		}
// 		return false
// 	})
// 	if !ok {
// 		return false
// 	}

// 	for _, item := range list {
// 		val := item.(data.Good)
// 		if _, ok := hero.Goods[int64(val.Id)]; !ok {
// 			return false
// 		}
// 	}
// 	return true
// }

// func (u *User) AddHeroGoods(goodsT data.Good, num int64) (uprop *UpdateProp, err error) {
// 	if num == 0 {
// 		err = errors.NewErrcode(data.Errcode.ServerErr, fmt.Sprintf("AddHeroGoods goods 数量为0goods tid:%d", goodsT.Id))
// 		return
// 	}
// 	hero, uprop := u.Heros().AddHeroGoods(goodsT, num)
// 	if hero.SkillId != 0 {
// 		return
// 	}

// 	if !hero.CheckFullGoods(goodsT.Owner) {
// 		fmt.Println("英雄道具没满")
// 		return uprop, nil
// 	}

// 	herosTpl, ok := data.Get("Heros", goodsT.Owner)
// 	if !ok {
// 		err = errors.NewErrcode(data.Errcode.ConfigNotFound, fmt.Sprintf("没有找到英雄配置:heros id:%d", goodsT.Owner))
// 		return
// 	}

// 	heroT := herosTpl.(data.Hero)

// 	heroSkillTpl, ok := data.Get("HeroSkills", heroT.Skill)
// 	if !ok {
// 		err = errors.NewErrcode(data.Errcode.ConfigNotFound, fmt.Sprintf("没有找到英雄技能配置:heros skill id:%d", heroT.Skill))
// 		return
// 	}

// 	skillT := heroSkillTpl.(data.HeroSkill)
// 	hero.SkillId = int64(skillT.Id)

// 	uprop.HeroUprops = append(uprop.HeroUprops, &HeroUprop{Tid: int64(goodsT.Owner), SkillId: hero.SkillId})

// 	switch int64(skillT.Type) {
// 	case data.HeroSkillTypeEnum.PetGet:
// 		fmt.Println("获得宠物")
// 		kv, pet, upropPetInfo, err := u.GetUserPet().AddPet(utils.Int64(skillT.Effect))
// 		if err != nil {
// 			return uprop, err
// 		}
// 		if pet != nil {
// 			uprop.Pets = append(uprop.Pets, pet)
// 		}
// 		if upropPetInfo != nil {
// 			uprop.Petinfo = upropPetInfo
// 		}
// 		if kv != nil {
// 			uprop.Kvs = append(uprop.Kvs, kv)
// 		}
// 	case data.HeroSkillTypeEnum.PetPlus:
// 		fmt.Println("宠物进阶")
// 		switch skillT.Effect.(type) {
// 		case map[string]interface{}:
// 			mEffect := skillT.Effect.(map[string]interface{})
// 			key := mEffect["key"]
// 			val := mEffect["value"]
// 			if key == 0 && val == 0 {
// 				return nil, nil
// 			}

// 			uprop1, pet, err := u.GetUserPet().AddPetStep(utils.Int64(key), utils.Int64(val))
// 			if err != nil {
// 				return uprop, err
// 			}
// 			if pet != nil {
// 				uprop.Pets = append(uprop.Pets, pet)
// 			}

// 			if uprop1 != nil {
// 				uprop.Merge(uprop1)
// 			}

// 		default:
// 			fmt.Println("宠物进阶 effect : ", goodsT.Effect)
// 		}
// 	case data.HeroSkillTypeEnum.GridCost:
// 		fmt.Println("挖矿相关辅助技能")
// 		switch int64(skillT.ExtType) {
// 		case data.HeroSkillExtTypeEnum.HandleGoldRate:
// 			u.HandleGoldRate += utils.Int64(skillT.Effect)
// 			return
// 		case data.HeroSkillExtTypeEnum.TheftGoldRate:
// 			u.TheftGoldRate += utils.Int64(skillT.Effect)
// 			return
// 		case data.HeroSkillExtTypeEnum.FireGridRate:
// 			u.FireGridRate += utils.Int64(skillT.Effect)
// 			uprop.Kvs = append(uprop.Kvs, &BaseKV{Key: data.GlobalEnum.FireGridRate, Val: utils.Int64(skillT.Effect)})
// 			return
// 		default:
// 			fmt.Println("无效的辅助类型 ： ", skillT.ExtType)
// 			return
// 		}
// 	default:
// 		fmt.Println("无效的类型 ： ", skillT.Type)
// 	}
// 	return
// }
