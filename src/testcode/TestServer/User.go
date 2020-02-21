package testserver

import (
	"errors"
	"sync"

	"github.com/globalsign/mgo/bson"
	log "github.com/sirupsen/logrus"
)

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

type BaseKV struct {
	Key int64 `json:"key"`
	Val int64 `json:"val"`
}

type User struct {
	BaseM
	Uid                bson.ObjectId `bson:"_id" json:"_uid"`
	AccountId          string        `bson:"accountId" json:"accountId"`
	Number             int64         `bson:"number" json:"number"` //玩家编号
	Name               string        `bson:"name" json:"name"`
	Head               string        `bson:"head" json:"head"`
	Energy             int64         `bson:"energy" json:"energy"`                   //体力
	Gold               int64         `bson:"gold" json:"gold"`                       //金币
	Diamon             int64         `bson:"diamon" json:"diamon"`                   //砖石
	Shield             int64         `bson:"shield" json:"shield"`                   //盾牌
	Pos                *PosInfo      `bson:"pos,omitempty" json:"pos"`               // 关卡位置信息
	MaxPos             *PosInfo      `bson:"maxPos,omitempty" json:"maxPos"`         // 最大关卡信息
	TheftNum           int64         `bson:"theftNum" json:"theftNum"`               // 被偷窃次数
	FireNum            int64         `bson:"fireNum" json:"fireNum"`                 // 放火次数
	LastLoginTime      int64         `bson:"last_login_time" json:"lastLoginTime"`   // 最后登录时间
	OfflineTime        int64         `bson:"offline_time" json:"offlineTime"`        // 离线时间
	Robot              bool          `bson:"robot" json:"robot"`                     // 是否是机器人
	EnergyTime         int64         `bson:"energyTime" json:"energyTime"`           // 给体力的时间
	NextRecoverTime    int64         `bson:"nextRecoverTime" json:"nextRecoverTime"` // 恢复体力时间差
	Status             int64         `bson:"status" json:"status"`                   // 状态 0 正常 1 正在被放火
	BanTime            int64         `bson:"banTime" json:"banTime"`                 // 解禁时间
	SlotNoticeUser     string        `bson:"slotNoticeUser" json:"slotNoticeUser"`   // 老虎机通知 玩家对象
	openFireGrid       bool
	FireProtec         bool            `bson:"fireProtec" json:"fireProtec"`         // 放火保护
	SlotLv             int64           `bson:"slotLv" json:"slotLv"`                 // 老虎机等级
	Star               int64           `bson:"star" json:"star"`                     // 星星数量
	StarSum            int64           `bson:"starSum" json:"starSum"`               // 星星总数量
	StarChangeTime     int64           `bson:"starChangeTime" json:"starChangeTime"` // 星星改变时间
	UnionId            string          `bson:"unionId" json:"unionId"`               // 公会id
	UnionScore         int64           `bson:"unionScore" json:"unionScore"`         // 公会积分
	UnionEnergyReqTime int64           `bson:"unionEnergyReqTime" json:"unionEnergyReqTime"`
	HandleGoldRate     int64           `bson:"handleGoldRate" json:"handleGoldRate"` // 拾取金币收益
	TheftGoldRate      int64           `bson:"theftGoldRate" json:"theftGoldRate"`   // 被偷金币损失
	FireGridRate       int64           `bson:"fireGridRate" json:"fireGridRate"`     // 挖焦土消耗金币
	RegisterTime       int64           `bson:"registerTime" json:"registerTime"`     // 注册时间
	Platform           int64           `bson:"platform" json:"platform"`             // 登录平台
	UsedName           string          `bson:"usedName" json:"usedName"`             // 曾用名
	AvatarFrame        int64           `bson:"avatarFrame" json:"avatarFrame"`       // 头像框ID-使用中
	AvatarFrameArr     map[int64]int64 `bson:"avatarFrameArr" json:"avatarFrameArr"` // 头像框ID列表-持有
	rw                 sync.RWMutex
}

// 关卡位置信息
type PosInfo struct {
	LockBlockTid     int64 `bson:"lockBlockTid" json:"lockBlockTid"`
	LayerId          int64 `bson:"layerId" json:"layerId"`
	LayerOpenGridNum int64 `bson:"layerOpenGridNum" json:"layerOpenGridNum"`
}

// 物品
type Goods struct {
	Tid int64 `bson:"tid" json:"tid"`
	Num int64 `bson:"num" json:"num"`
}

// 宝箱内的物品
type BoxRes struct {
	Tid int64
	Res []*Goods
}

// 浅赋值
func (u *User) ShallowSet(u1 User) {
	defer u.Update(u)
	u.Energy = u1.Energy
	u.Gold = u1.Gold
	u.Diamon = u1.Diamon
	u.Shield = u1.Shield
	u.SlotLv = u1.SlotLv
	u.Star = u1.Star
	u.EnergyTime = u1.EnergyTime
	u.NextRecoverTime = u1.NextRecoverTime
	u.UnionScore = u1.UnionScore
	u.UnionEnergyReqTime = u1.UnionEnergyReqTime
	u.HandleGoldRate = u1.HandleGoldRate
	u.TheftGoldRate = u1.TheftGoldRate
	u.FireGridRate = u1.FireGridRate
}

func NewUser() EntryI {
	return new(User)
}

func (u *User) Save() (err error) {
	if u.Robot {
		return
	}
	// u.StarSum = u.StarVal()
	// err = db.Upsert("user", bson.M{"_id": u.Uid}, u)
	// if err != nil {
	// 	return
	// }
	return
}

func (u *User) Load(key interface{}) (err error) {
	switch key.(type) {
	case string:
		if key.(string) == "" || !bson.IsObjectIdHex(key.(string)) {
			log.WithFields(log.Fields{
				"method": "User_Load",
				"data":   key,
			}).Error("无效的key load")
			return errors.New("无效的key load")
		}
		u.Uid = bson.ObjectIdHex(key.(string))
	case bson.ObjectId:
		u.Uid = key.(bson.ObjectId)
	default:
		log.WithFields(log.Fields{
			"method": "User_Load",
			"data":   key,
		}).Error("无效的key load")
		return errors.New("无效的key load")
	}
	return u.LoadData()
}

// 加载数据
func (u *User) LoadData() (err error) {
	// err = db.FindOne("user", bson.M{"_id": u.Uid}, nil, u)
	err = nil
	// if err != nil {
	// 	if err.Error() == "not found" {
	// 		err = errors.NewErrcode(data.Errcode.UserNotFound, "角色不存在")
	// 		return err
	// 	}
	// 	return
	// }
	// if u.AvatarFrameArr == nil {
	// 	u.AvatarFrameArr = make(map[int64]int64, 0)
	// 	u.AvatarFrameArr[data.AvatarFrameEnum.DefaultFrame] = data.AvatarFrameEnum.DefaultFrame //默认头像框
	// }
	// if u.AvatarFrame == 0 {
	// 	u.AvatarFrame = data.AvatarFrameEnum.DefaultFrame //默认头像框
	// }
	return
}

// 更新数据
func (u *User) Update(val *User) {
	cache.Put("user", u.Uid.Hex(), val)
	val.UpdateData()
}

// 初始化
// func (u *User) Init() (err error) {
// 	u.CaculEnergy()
// 	u.HideShop().InitHideShop()
// 	return
// }

// func (u *User) FindByName(name string) (err error) {
// 	err = db.FindOne("user", bson.M{"name": name}, nil, u)
// 	if err != nil {
// 		err = errors.NewErrcode(data.Errcode.UserNotFound, err.Error())
// 		return
// 	}
// 	return
// }

// func (u *User) FindByAccountId(accountId string) (err error) {
// 	err = db.FindOne("user", bson.M{"accountId": accountId}, nil, u)
// 	if err != nil {
// 		err = errors.NewErrcode(data.Errcode.UserNotFound, err.Error())
// 		return
// 	}
// 	return
// }

// 获取玩家数量
// func (u *User) Count() (count int64, err error) {
// 	rcount, err := db.Count("user", nil)
// 	if err != nil {
// 		return
// 	}
// 	count = int64(rcount)
// 	return
// }

// 设置关卡信息
// func (u *User) SetLockBlocks(val *LockBlocks) {
// 	Cache.Put("lockBlock", u.Uid.Hex(), val)
// 	val.UpdateData()
// }

// func (u *User) AddGold(from string, val int64) (r *BaseKV) {
// 	u.rw.Lock()
// 	defer u.rw.Unlock()
// 	logData := make(map[string]interface{}, 0)
// 	logData["val"] = val
// 	logData["oldGold"] = u.Gold
// 	u.Gold += val
// 	r = &BaseKV{Key: data.GlobalEnum.Gold, Val: val}
// 	logData["newGold"] = u.Gold
// 	if from != "CheckCost" {
// 		log.WithFields(log.Fields{
// 			"from":   from,
// 			"uid":    u.Uid.Hex(),
// 			"method": "AddGold",
// 			"data":   logData,
// 		}).Debug("基础数据变更")
// 	}

// 	return
// }

// func (u *User) AddDiamond(from string, val int64) (r *BaseKV) {
// 	u.rw.Lock()
// 	defer u.rw.Unlock()
// 	logData := make(map[string]interface{}, 0)
// 	logData["val"] = val
// 	logData["oldDiamon"] = u.Diamon
// 	u.Diamon += val
// 	r = &BaseKV{Key: data.GlobalEnum.Diamond, Val: val}
// 	logData["newDiamon"] = u.Diamon
// 	if from != "CheckCost" {
// 		log.WithFields(log.Fields{
// 			"from":   from,
// 			"uid":    u.Uid.Hex(),
// 			"method": "AddDiamond",
// 			"data":   logData,
// 		}).Debug("基础数据变更")
// 	}
// 	return
// }

// func (u *User) AddEnergy(from string, val int64) (r []*BaseKV) {
// 	u.rw.Lock()
// 	defer u.rw.Unlock()
// 	logData := make(map[string]interface{}, 0)
// 	logData["val"] = val
// 	logData["oldEnergy"] = u.Energy

// 	//当前体力>=50，或者当前体力值+val>=50时，体力时间（起始时间）改为当前时间
// 	if u.Energy >= 50 || u.Energy+val >= 50 {
// 		u.EnergyTime = time.Now().Unix()
// 		u.NextRecoverTime = int64(3600)
// 	}
// 	//每小时恢复体力
// 	nextRecoverTime := int64(3600)
// 	// 宠物加快体力恢复速度
// 	rateVal := u.GetUserPet().GetSkillVal(uint64(data.PetSkillEffectEnum.FastRestEnergy))
// 	if rateVal > 0 {
// 		nextRecoverTime -= int64(float64(3600) * float64(rateVal) / float64(100))
// 	}

// 	//恢复体力方法请求时，需重新计算恢复时间差 || 当体力>=50 && 体力+val<50，需计算恢复体力时间差
// 	if (from == "RestoreEnergyRequest") || (nextRecoverTime < u.NextRecoverTime && u.Energy >= 50 && u.Energy+val < 50) {
// 		u.NextRecoverTime = nextRecoverTime
// 	}
// 	u.Energy += val
// 	r = append(r, &BaseKV{Key: data.GlobalEnum.Energy, Val: val})
// 	r = append(r, &BaseKV{Key: data.GlobalEnum.EnergyTime, Val: u.EnergyTime})
// 	r = append(r, &BaseKV{Key: data.GlobalEnum.NextRecoverTime, Val: u.NextRecoverTime})

// 	if from != "CheckCost" {
// 		logData["newEnergy"] = u.Energy
// 		log.WithFields(log.Fields{
// 			"from":   from,
// 			"uid":    u.Uid.Hex(),
// 			"method": "AddEnergy",
// 			"data":   logData,
// 		}).Debug("基础数据变更")
// 	}
// 	return
// }

// func (u *User) AddShield(from string, val int64) (r []*BaseKV) {
// 	if val > 0 && u.Shield >= 3 {
// 		return u.AddEnergy(from+"_AddShield", val)
// 	}
// 	if u.Shield+val > 3 && val > 0 {
// 		r = append(r, u.AddEnergy(from+"_AddShield", val-(3-u.Shield))...)
// 		val = 3 - u.Shield
// 	}
// 	logData := make(map[string]interface{}, 0)
// 	logData["val"] = val
// 	logData["oldShield"] = u.Shield
// 	u.Shield += val
// 	r = append(r, &BaseKV{Key: data.GlobalEnum.Shield, Val: val})
// 	logData["newShield"] = u.Shield
// 	if from != "CheckCost" {
// 		log.WithFields(log.Fields{
// 			"from":   from,
// 			"uid":    u.Uid.Hex(),
// 			"method": "AddShield",
// 			"data":   logData,
// 		}).Debug("基础数据变更")
// 	}
// 	return
// }

// func (u *User) AddSlotLv(val int64) (r *BaseKV) {
// 	u.rw.Lock()
// 	defer u.rw.Unlock()
// 	u.SlotLv += val
// 	r = &BaseKV{Key: data.GlobalEnum.SlotLv, Val: val}
// 	return
// }

// func (u *User) AddStar(val int64) (r *BaseKV) {
// 	u.rw.Lock()
// 	defer u.rw.Unlock()
// 	u.Star += val
// 	r = &BaseKV{Key: data.GlobalEnum.Star, Val: val}
// 	return
// }

// func (u *User) AddUnionScore(val int64) (r *BaseKV) {
// 	u.rw.Lock()
// 	defer u.rw.Unlock()
// 	u.UnionScore += val
// 	r = &BaseKV{Key: data.GlobalEnum.Union, Val: val}
// 	return
// }

// 增加玩家持有的头像(不替换) 传入参数，头像id; needCallClient 玩家数据发生变更，需要推给玩家
// func (u *User) AddAvatarFrame(val int64) (needCallClient bool, curFrame int64, gainFrameArr map[int64]int64) {
// 	needCallClient = false
// 	curFrame = u.AvatarFrame
// 	u.rw.Lock()
// 	defer u.rw.Unlock()
// 	if u.AvatarFrameArr == nil {
// 		u.AvatarFrameArr = make(map[int64]int64, 0)
// 		u.AvatarFrameArr[data.AvatarFrameEnum.DefaultFrame] = data.AvatarFrameEnum.DefaultFrame //默认头像框
// 		needCallClient = true
// 	}
// 	// 检查是否是配表有的头像ID
// 	// 读表
// 	_, findTheConfig := data.Get("AvatarFrames", uint64(val))
// 	if findTheConfig == false {
// 		log.WithFields(log.Fields{
// 			"method": "AddAvatarFrame",
// 			"buffId": val,
// 		}).Error("AvatarFrames读表失败,没有找到该ID的头像框配置")
// 		return needCallClient, u.AvatarFrame, u.AvatarFrameArr
// 	}
// 	//
// 	_, findFrame := u.AvatarFrameArr[val]
// 	if findFrame == false {
// 		u.AvatarFrameArr[val] = val
// 		u.Update(u)
// 		needCallClient = true
// 	}
// 	return needCallClient, u.AvatarFrame, u.AvatarFrameArr
// }

// 删除玩家指定的头像，如果玩家当前正在使用删除头像，则换成默认头像;needCallClient 玩家数据发生变更，需要推给玩家
// func (u *User) RemoveAvatarFrame(val int64) (needCallClient bool, curFrame int64, gainFrameArr map[int64]int64) {
// 	needCallClient = false
// 	u.rw.Lock()
// 	defer u.rw.Unlock()
// 	if u.AvatarFrameArr == nil {
// 		u.AvatarFrameArr = make(map[int64]int64, 0)
// 		u.AvatarFrameArr[data.AvatarFrameEnum.DefaultFrame] = data.AvatarFrameEnum.DefaultFrame //默认头像框
// 		needCallClient = true
// 		return needCallClient, u.AvatarFrame, u.AvatarFrameArr
// 	}
// 	if val == data.AvatarFrameEnum.DefaultFrame { // 不能删除默认头像框
// 		return needCallClient, u.AvatarFrame, u.AvatarFrameArr
// 	}
// 	_, findFrame := u.AvatarFrameArr[val]
// 	if findFrame {
// 		delete(u.AvatarFrameArr, val)
// 		if u.AvatarFrame == val {
// 			//玩家当前正在使用删除头像，则换成默认头像
// 			u.AvatarFrame = data.AvatarFrameEnum.DefaultFrame
// 		}
// 		needCallClient = true
// 		u.Update(u)
// 	}
// 	return needCallClient, u.AvatarFrame, u.AvatarFrameArr
// }

// 更改玩家的头像框，传入参数为id
// func (u *User) ChangeAvatarFrame(targetFrameId int64) {
// 	u.rw.Lock()
// 	defer u.rw.Unlock()
// 	_, findFrame := u.AvatarFrameArr[targetFrameId]
// 	if findFrame {
// 		u.AvatarFrame = targetFrameId
// 		u.Update(u)
// 	}
// 	return
// }

// 设置为默认头像
// func (u *User) SetDefaultAvatarFrame() {
// 	u.rw.Lock()
// 	defer u.rw.Unlock()
// 	if u.AvatarFrameArr == nil {
// 		u.AvatarFrameArr = make(map[int64]int64, 0)
// 	}
// 	defaultFrame := data.AvatarFrameEnum.DefaultFrame
// 	u.AvatarFrameArr[defaultFrame] = defaultFrame //默认头像框
// 	u.AvatarFrame = defaultFrame
// 	u.Update(u)
// }

// 更改用户名称，头像
// func (u *User) ChangeNameAndHead(newName string, newHead string) {
// 	u.rw.Lock()
// 	defer u.rw.Unlock()
// 	usedName := u.Name
// 	u.Name = newName
// 	u.Head = newHead
// 	u.UsedName = usedName
// 	u.Update(u)
// 	return
// }

// 更改用户登录平台类型
// func (u *User) ChangePlatform(newPlatform int64) {
// 	u.rw.Lock()
// 	defer u.rw.Unlock()
// 	u.Platform = newPlatform
// 	u.Update(u)
// 	return
// }

// 更改登录平台，同时更改名字和头像
// func (u *User) ChangePlatformAndNameHead(newPlatform int64, newName string, newHead string) {
// 	u.rw.Lock()
// 	defer u.rw.Unlock()
// 	usedName := u.Name
// 	u.Name = newName
// 	u.Head = newHead
// 	u.UsedName = usedName
// 	u.Platform = newPlatform
// 	u.Update(u)
// 	return
// }

// 计算并恢复 的体力 初始化用户数据才调用，用户登录计算体力。
// func (u *User) CaculEnergy() {
// 	//初始化恢复体力时间差
// 	u.NextRecoverTime = int64(3600)

// 	now := time.Now().Unix()
// 	tnum := now - u.EnergyTime //现在时间-体力时间=时间差
// 	if tnum <= 0 {
// 		return
// 	}
// 	pret := tnum / 3600 //几个小时
// 	if pret <= 0 {
// 		return
// 	}
// 	num := pret * 5     //每小时增加5点体力
// 	if u.Energy >= 50 { //当前体力>=50 不恢复体力
// 		u.EnergyTime = now
// 		return
// 	}
// 	if u.Energy+num > 50 { //当前体力+时间恢复体力>50
// 		u.EnergyTime = now
// 		num = 50 - u.Energy
// 	} else { //当前体力+时间恢复体力<=50
// 		lessTime := tnum % 3600       //增加体力用掉的时间，计算未用的时间
// 		u.EnergyTime = now - lessTime //当前时间-未用掉的时间
// 	}
// 	u.AddEnergy("CaculEnergy", num)
// }

// func (u *User) StarVal() (star int64) {
// 	star = u.Star
// 	for _, item := range u.LockBlocks().LockBlocks {
// 		star += item.Star
// 	}
// 	star += u.Heros().StarNum()
// 	star += u.GetUserPet().SumPetStar()
// 	return
// }

// 玩家被放火
// func (u *User) FireIng() {
// 	defer u.Update(u)
// 	u.Status = 1
// 	u.BanTime = time.Now().Add(time.Minute * 5).Unix()
// }

// 解除放火
// func (u *User) UnFire() {
// 	defer u.Update(u)
// 	u.Status = 0
// 	u.BanTime = 0
// }

//增加被偷盗次数
// func (u *User) AddTheftNum() {
// 	defer u.Update(u)
// 	u.TheftNum += 1
// }

// 增加被放火次数
// func (u *User) AddFireNum() {
// 	defer u.Update(u)
// 	u.FireNum += 1
// }

// 设置老虎机通知对象
// func (u *User) SetSlotNoticeUser(uid string) {
// 	defer u.Update(u)
// 	u.SlotNoticeUser = uid
// }

// 上线
// func (u *User) OnLine() {
// 	defer u.Update(u)
// 	u.OfflineTime = 0
// 	u.LastLoginTime = time.Now().Unix()
// 	u.TheftNum = 0
// 	u.FireNum = 0
// }

// 离线
// func (u *User) OffLine() {
// 	defer u.Update(u)
// 	u.OfflineTime = time.Now().Unix()
// }

// 玩家是否离线
// func (u *User) IsOffLine() bool {
// 	return u.OfflineTime > 0
// }

//玩家是否在线
// func (u *User) IsOnLine() bool {
// 	return u.OfflineTime == 0
// }

//
// func (u *User) OpenFireGrid() {
// 	u.openFireGrid = true
// }

// 是否开启了 格子
// func (u *User) IsOpenFireGrid() bool {
// 	return u.openFireGrid
// }

// 移除放火保护
// func (u *User) RemoveFireProtec() {
// 	defer u.Update(u)
// 	u.FireProtec = false
// }

// 增加放火保护
// func (u *User) AddFireProtec() {
// 	defer u.Update(u)
// 	u.FireProtec = true
// }
