package main

import (
	"errors"
	"fmt"
	"github.com/tint/inquirer"
	"github.com/tint/inquirer/terminal"
)

var typ *Type
var emojis []*Emoji
var scope string
var subject string
var files []*gitFile
var config *Config

func init() {
	config = LoadConfig()
}

func selectFiles() error {
	var fs []*gitFile
	if err := lsRepoFiles(&fs); err != nil {
		return err
	}
	if len(fs) == 0 {
		return nil
	}
	var opts []string
	var defs []string
	for _, file := range fs {
		opts = append(opts, file.Path)
		if file.flag == 'A' {
			defs = append(defs, file.Path)
		}
	}
	var selects []string
	for {
		err := inquirer.AskOne(&inquirer.MultiSelect{
			Message:     "请选择提交的文件",
			Options:     opts,
			Default:     defs,
			Description: func(value string, index int) string { return fs[index].Status() },
		}, &selects)
		if err != nil {
			return err
		}
		if len(fs) == len(selects) {
			files = fs
			return nil
		}
		files = []*gitFile{}
		for _, f := range fs {
			for _, path := range selects {
				if f.Path == path {
					files = append(files, f)
				}
			}
		}
		if len(files) > 0 {
			return nil
		}
	}
}

func selectType() error {
	var options []string
	for _, t := range config.Types {
		options = append(options, t.Name)
	}
	var val string
	err := inquirer.AskOne(&inquirer.Select{
		Message:     "请选择提交类型：",
		Options:     options,
		Description: func(value string, index int) string { return config.Types[index].Desc },
	}, &val, nil)
	if err != nil {
		return err
	}
	for _, t := range config.Types {
		if t.Name == val {
			typ = t
			return nil
		}
	}
	throw(errors.New("选择了错误类型"))
	return nil
}

func selectEmojis() error {
	if typ == nil {
		err := selectType()
		if err != nil {
			return err
		}
	}
	// 没有找到关联的 emoji 列表
	if len(typ.Emojis) == 0 || len(config.Emojis) == 0 {
		return nil
	}
	var es []*Emoji
	var options []string
	for _, e := range config.Emojis {
		for _, s := range typ.Emojis {
			if e.Code == s {
				es = append(es, e)
				options = append(options, e.Text)
				break
			}
		}
	}
	if len(es) == 0 {
		return nil
	}
	var val []string
	err := inquirer.AskOne(&inquirer.MultiSelect{
		Message:     "请选择 Emoji 符号:",
		Options:     options,
		Description: func(value string, index int) string { return es[index].Desc },
	}, &val)
	if err != nil {
		return err
	}
	emojis = []*Emoji{}
	for _, s := range val {
		for _, emj := range es {
			if s == emj.Text {
				emojis = append(emojis, emj)
			}
		}
	}
	return nil
}

func inputScope() error {
	var val string
	err := inquirer.AskOne(&inquirer.Input{
		Message: "请输入影响范围：",
		Default: "*",
	}, &val)
	if err != nil {
		return err
	}
	scope = val
	return nil
}

func inputSubject() error {
	var val string
	for {
		err := inquirer.AskOne(&inquirer.Input{
			Message: "请输入日志内容，建议不超过50个字符：",
		}, &val)
		if err != nil {
			return err
		}
		if len(val) > 0 {
			subject = val
			return nil
		}
	}
}

func ask() error {
	if err := selectFiles(); err != nil {
		return err
	}
	if err := selectType(); err != nil {
		return err
	}
	if err := selectEmojis(); err != nil {
		return err
	}
	if err := inputScope(); err != nil {
		return err
	}
	if err := inputSubject(); err != nil {
		return err
	}
	return nil
}

func preview() string {
	var str string
	if len(files) >= 8 {
		for i, file := range files {
			if i > 0 {
				str += "\n"
			}
			if i > 6 {
				str += fmt.Sprintf("  └─...剩余%d个文件", len(files)-6)
				break
			} else {
				str += fmt.Sprintf("  ├─%s (%s)", file.Path, file.Status())
			}
		}
	} else {
		for i, file := range files {
			if i > 0 {
				str += "\n"
			}
			str += fmt.Sprintf("  ├─%s (%s)", file.Path, file.Status())
		}
	}
	if len(str) > 0 {
		str = fmt.Sprintf("选择的文件（共%d个）：\n%s\n", len(files), str)
	}
	str += "提交日志：\n"
	str += "  " + getMessage(false)
	return str
}

func getMessage(code bool) string {
	var str string
	if typ != nil {
		str += typ.Name
		if len(scope) > 0 {
			str += "(" + scope + ")"
		} else {
			str += "(*)"
		}
		str += ": "
	}
	if len(emojis) > 0 {
		for _, emoji := range emojis {
			if code {
				str += ":" + emoji.Code + ":"
			} else {
				str += emoji.Text
			}
		}
		str += " "
	}
	str += subject
	return str
}

func confirm() (bool, error) {
	if typ == nil {
		err := ask()
		if err != nil {
			return false, err
		}
	}
	ok := false
	err := inquirer.AskOne(&inquirer.Confirm{
		Message: "输入结果如下：\n" + preview() + "\n立即提交？",
	}, &ok)
	if err != nil {
		return false, err
	}
	if !ok {
		return false, nil
	}
	return true, nil
}

func commit() error {
	var news []string
	for _, file := range files {
		if file.flag == '?' {
			news = append(news, "./"+file.Path)
		}
	}
	if len(news) > 0 {
		_, err := git("add", news...)
		if err != nil {
			return err
		}
	}
	args := []string{"-o"}
	for _, f := range files {
		args = append(args, "./"+f.Path)
	}
	args = append(args, "-m", getMessage(true))
	_, err := git("commit", args...)
	return err
}

func main() {
	if !isGitRepo() {
		err := initRepo()
		if err != nil {
			throw(err)
		}
	}
	ok, err := confirm()
	if err == nil && ok {
		err = commit()
	}
	if err != nil {
		if errors.Is(err, terminal.InterruptErr) {
			throw(errors.New("Ctrl+C pressed in Terminal"))
		}
		throw(err)
	}
}
