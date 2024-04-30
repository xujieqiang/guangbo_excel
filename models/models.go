package models

type Alltask struct {
	Id       int
	Taskname string // 任务的名称 如上课铃声，下课铃声等
	Jobtype  int    // 1 表示每天  2 表示每周  3 表示每月  4 表示一次性任务
	Jobmask  int    // 是否有终止日期 0表示没有 1表示有
	Medias   string // 音频路径
	Tag      int    // 1 表示日常任务， 0表示特殊 或者 不寻常任务。 2 中考铃声任务
}

func (a *Alltask) TableName() string {
	return "alltask"
}

type Fangan struct {
	//  方案列表，总的
	Id          int    // id
	Listname    string // 方案名称
	Discreption string //方案描述
	Createtime  string
	Fatherfa    int
}

func (f *Fangan) TableName() string {
	return "fangan"
}

type Groups struct {
	Id        int    // id
	Groupname string // 给那些组的数字标号取一个名字  如 初一 初二等
	Val       string // 在广播系统中真正有意义的数字组别
}

func (g *Groups) TableName() string {
	return "groups"
}

type Mediapath struct {
	Id        int
	Medianame string // 媒体文件名称  如什么音乐
	Pathval   string // 媒体存放路径
}

func (m *Mediapath) TableName() string {
	return "mediapath"
}

type Tasklist struct {

	// 生成最终的方案列表
	Id        int
	Fanganid  int    // 方案id
	Taskid    int    // 任务id
	Name      string // 就是taskname  任务名称
	Jobtype   int    //  1 表示每天  2 表示每周  3 表示每月  4 表示一次性任务
	Jobmask   int    // 是否有终止日期 0表示没有 1表示有
	Duration  int    // 功能未知 默认为0
	Starttime string // 开始时间  年-月-日  时：分：秒格式
	Stoptime  string // 结束时间  年-月-日  时：分：秒格式
	Jobdata   int    // 表示一周中有几天是需要执行打铃任务的， 如果JobType 是4 一次性任务时，该值为0，如果
	// 65663 = 1000 000 000 000 000 00   + 1111111
	// 111 1110  表示的是 周一不执行
	// 1 2 4 8 16  32 64
	Repeatnum      int    // 曲子重复播放的次数
	Playmode       int    // 未知 默认 0
	Playvol        int    // 未知 默认 0
	Medias         string // 歌曲路径
	Terms          string // 要播放的终端列表
	Areamasks      string // 未知 默认 空字符串
	Groups         string // 要播放的组列表
	Poweraheadplay int    // 未知  默认0
}

func (t *Tasklist) TableName() string {
	return "tasklist"
}

type Setting struct {
	Id   int
	Dayd string
}

func (s *Setting) TableName() string {
	return "setting"
}
