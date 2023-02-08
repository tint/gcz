package main

import (
	"os"
	"path/filepath"
	"strings"
)

func isGitRepo() bool {
	dir, err := os.Getwd()
	if err != nil {
		throw(err)
	}
	file, err := os.Stat(filepath.Join(dir, ".git"))
	if err != nil {
		if os.IsNotExist(err) {
			return false
		}
		throw(err)
	}
	if !file.IsDir() {
		return false
	}
	if _, err = git("status"); err != nil {
		throw(err)
	}
	return true
}

func initRepo() error {
	_, err := git("init")
	if err == nil {
		_, err = git("branch", "-M", "main")
	}
	return err
}

type gitFile struct {
	Path string
	flag uint8
}

func (f *gitFile) Status() string {
	switch f.flag {
	case 'D':
		return "删除"
	case 'M':
		return "修改"
	case 'R':
		return "重命名"
	case 'T':
		return "修改文件类型"
	case 'C':
		return "复制"
	case 'U':
		return "更新但未合并"
	case '?':
		return "新增"
	}
	return ""
}

func lsRepoFiles(fs *[]*gitFile) error {
	stdout, err := git("status", "-s", "-uall")
	if err != nil {
		return err
	}
	lines := strings.Split(stdout.String(), "\n")
	var files []*gitFile
	for _, line := range lines {
		if len(line) < 3 {
			continue
		}
		// 目录表示里面文件全名新增
		if strings.HasSuffix(line, "/") {
			err = scanDir(strings.TrimSpace(line[2:]), func(s string) error {
				stdout, err := git("check-ignore", s)
				if err != nil {
					return err
				}
				if strings.TrimSpace(stdout.String()) == s {
					return nil
				}
				files = append(files, &gitFile{
					Path: "",
					flag: 0,
				})
				return nil
			})
			if err != nil {
				return err
			}
		} else {
			files = append(files, &gitFile{
				Path: strings.TrimSpace(line[2:]),
				flag: line[1],
			})
		}
	}
	*fs = files
	return nil
}
