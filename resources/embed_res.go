package resources

import (
	"embed"
	"fmt"
	"io"
	"io/fs"

	"github.com/starter-go/base/safe"
)

// NewEmbed 用 embed.FS 新建资源
func NewEmbed(fs embed.FS, basepath string) Table {
	b := &embedResBuilder{}
	b.load(fs, basepath)
	return b.create()
}

////////////////////////////////////////////////////////////////////////////////

type embedResBuilder struct {
	group *embedFileGroup
	mode  safe.Mode
}

func (inst *embedResBuilder) load(fs embed.FS, basepath string) error {

	group := &embedFileGroup{}
	group.init(fs, basepath)
	inst.group = group

	return inst.loadDir(basepath, &group.fs, 0)
}

func (inst *embedResBuilder) loadDir(path string, fs *embed.FS, depth int) error {

	if depth > 32 {
		return fmt.Errorf("resource path is too deep: [%s]", path)
	}

	items, err := fs.ReadDir(path)
	if err != nil {
		return err
	}

	for _, item := range items {
		path2 := path + "/" + item.Name()
		if item.IsDir() {
			err := inst.loadDir(path2, fs, depth+1)
			if err != nil {
				return err
			}
		} else if item.Type().IsRegular() {
			inst.handleFile(path2, fs, item)
		}
	}

	return nil
}

func (inst *embedResBuilder) handleFile(path string, fs *embed.FS, item fs.DirEntry) {

	prefixLen := len(inst.group.basepath)
	shortpath := path[prefixLen:]

	file := &embedFile{}
	file.shortPath = normalizePath(shortpath)
	file.longPath = path
	file.group = inst.group
	file.name = item.Name()

	info, err := item.Info()
	if err == nil && info != nil {
		file.size = info.Size()
	}

	inst.group.add(file)
}

func (inst *embedResBuilder) create() Table {
	src := inst.group.files
	dst := make(map[string]Resource)
	for _, item := range src {
		dst[item.shortPath] = item
	}
	t := NewTable(dst, inst.mode)
	return t
}

////////////////////////////////////////////////////////////////////////////////

type embedFile struct {
	group     *embedFileGroup
	longPath  string
	shortPath string
	name      string
	size      int64
}

func (inst *embedFile) _Impl() Resource {
	return inst
}

func (inst *embedFile) Path() string {
	return inst.shortPath
}

func (inst *embedFile) SimpleName() string {
	return inst.name
}

func (inst *embedFile) Size() int64 {
	return inst.size
}

func (inst *embedFile) ReadBinary() ([]byte, error) {
	path := inst.longPath
	return inst.group.fs.ReadFile(path)
}

func (inst *embedFile) ReadText() (string, error) {
	data, err := inst.ReadBinary()
	if err != nil {
		return "", err
	}
	return string(data), nil
}

func (inst *embedFile) Open() (io.ReadCloser, error) {
	path := inst.longPath
	return inst.group.fs.Open(path)
}

////////////////////////////////////////////////////////////////////////////////

type embedFileGroup struct {
	fs       embed.FS
	basepath string
	files    map[string]*embedFile
}

func (inst *embedFileGroup) init(fs embed.FS, basepath string) {
	inst.fs = fs
	inst.basepath = basepath
	inst.files = make(map[string]*embedFile)
}

func (inst *embedFileGroup) add(item *embedFile) {
	inst.files[item.shortPath] = item
}

////////////////////////////////////////////////////////////////////////////////
