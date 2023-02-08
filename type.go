package main

type Type struct {
	Name     string   `json:"name"`             // 名称
	Desc     string   `json:"desc"`             // 描述
	Emojis   []string `json:"emojis,omitempty"` // 有效表情
	Platform bool     `json:"platform"`         // 平台关联
}

var DefaultTypes = []*Type{
	{
		Name:   "tada",
		Desc:   "创世提交",
		Emojis: []string{"tada"},
	},
	{
		Name:     "feat",
		Desc:     "实现新的功能",
		Emojis:   []string{"new", "sparkles", "alembic", "egg", "bento", "green_heart", "rotating_light"},
		Platform: true,
	},
	{
		Name:     "fix",
		Desc:     "自动修复问题 (适合于一次提交直接修复问题)",
		Emojis:   []string{"bug", "ambulance", "globe_with_meridians", "wrench", "pencil2", "green_heart", "lock", "fire"},
		Platform: true,
	},
	{
		Name:     "to",
		Desc:     "不自动修复问题 (适合于多次提交，最终修复问题提交时使用fix)",
		Emojis:   []string{"bug", "ambulance", "globe_with_meridians", "wrench", "pencil2", "construction", "green_heart", "lock", "fire"},
		Platform: true,
	},
	{
		Name: "docs",
		Desc: "文档变更",
		Emojis: []string{
			"globe_with_meridians", "lipstick", "clapper", "wrench", "heavy_plus_sign",
			"heavy_minus_sign", "zap", "racehorse", "memo", "book", "construction", "sparkles",
			"bulb", "fire",
		},
		Platform: true,
	},
	{
		Name:   "style",
		Desc:   "格式样式美化（不影响代码运行的变动）",
		Emojis: []string{"tada", "globe_with_meridians", "lipstick", "clapper", "art", "wheelchair", "fire"},
	},
	{
		Name:     "refactor",
		Desc:     "代码重构（即没有新增功能，也没有修改bug）",
		Emojis:   []string{"rotating_light", "heavy_plus_sign", "heavy_minus_sign", "art", "wastebasket", "construction_worker"},
		Platform: true,
	},
	{
		Name: "perf",
		Desc: "优化性能或体验等",
		Emojis: []string{
			"globe_with_meridians", "lipstick", "arrow_up", "arrow_down", "zap", "racehorse",
			"chart_with_upwards_trend", "hammer", "construction_worker", "wheelchair", "children_crossing",
		},
		Platform: true,
	},
	{
		Name: "test",
		Desc: "新增测试用例或是变动现有测试",
		Emojis: []string{
			"globe_with_meridians", "wrench", "heavy_plus_sign", "heavy_minus_sign", "fire",
			"arrow_up", "arrow_down", "wheelchair", "construction_worker", "white_check_mark",
		},
		Platform: true,
	},
	{
		Name: "review",
		Desc: "代码审查相关",
		Emojis: []string{
			"ok_hand", "fire", "art", "bug", "sparkles", "pencil", "rocket", "lipstick", "arrow_down", "arrow_up",
			"pushpin", "chart_with_upwards_trend", "whale", "wrench", "globe_with_meridians", "pencil2", "poop",
			"wheelchair", "bulb", "loud_sound", "mute", "building_construction", "mag", "wheel_of_dharma", "label",
		},
		Platform: true,
	},
	{
		Name:     "build",
		Desc:     "项目构建系统，如构建过程或构建工具的变动",
		Emojis:   []string{"bookmark", "rocket", "fire"},
		Platform: true,
	},
	{
		Name:     "ci",
		Desc:     "持续集成工具变动，如 Travis，Jenkins，GitLab CI，Circle 等",
		Emojis:   []string{"construction_worker", "chart_with_upwards_trend", "green_heart", "rocket", "fire"},
		Platform: true,
	},
	{
		Name:     "depend",
		Desc:     "依赖库或依赖工具的变动",
		Emojis:   []string{"heavy_minus_sign", "heavy_plus_sign", "arrow_up", "arrow_down", "pushpin", "rocket"},
		Platform: true,
	},
	{
		Name: "chore",
		Desc: "工程依赖或辅助工具的变动",
		Emojis: []string{
			"globe_with_meridians", "heavy_plus_sign", "heavy_minus_sign", "fire",
			"arrow_up", "arrow_down", "wheelchair", "label", "mag", "see_no_evil", "rocket",
			"busts_in_silhouette", "building_construction",
			"boom", "truck", "alien", "package", "poop", "pencil2", "recycle",
			"chart_with_upwards_trend", "iphone", "clown_face", "camera_flash",
		},
		Platform: true,
	},
	{
		Name:   "LICENSE",
		Desc:   "许可协议相关",
		Emojis: []string{"page_facing_up", "fire"},
	},
	{
		Name:     "revert",
		Desc:     "回滚到上一个版本",
		Emojis:   []string{"bookmark", "rotating_light", "rewind", "rocket"},
		Platform: true,
	},
	{
		Name:   "merge",
		Desc:   "代码合并",
		Emojis: []string{"bookmark", "twisted_rightwards_arrows", "rocket"},
	},
}
